package hidlisten

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os/exec"
	"time"
)

const hidListenReadyLine = "Listening:\n"

var ErrTimeout = errors.New("timed out while waiting for hid_listen to be ready")

type HidListen struct {
	cmd *exec.Cmd

	StdOutputPipe io.ReadCloser
	StdErrPipe    io.ReadCloser
}

func New() (*HidListen, error) {
	cmd := exec.Command("hid_listen")

	outputPipe, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	stdErrPipe, err := cmd.StderrPipe()
	if err != nil {
		return nil, err
	}

	return &HidListen{
		cmd:           cmd,
		StdOutputPipe: outputPipe,
		StdErrPipe:    stdErrPipe,
	}, nil
}

// waitUntilReady reads the outputPipe looking for the "Listening" line that indicates that hid_listen is ready.
// If it doesn't see it within a certain time frame it will timeout and return an error.
func (h *HidListen) waitUntilReady(waitLimit time.Duration) error {
	reader := bufio.NewReader(h.StdOutputPipe)

	timer := time.NewTimer(waitLimit)
	defer timer.Stop()

	isReadyChan := make(chan bool)
	readErrChan := make(chan error)

	go func() {
		for {
			// If this returns error nothing will ever be sent to the isReady channel and the error will
			// be picked up bellow on the timeout block of the select statement
			line, err := reader.ReadString('\n')
			if err != nil {
				readErrChan <- err

				return
			}

			if line == hidListenReadyLine {
				isReadyChan <- true

				return
			}
		}
	}()

	var err error

	select {
	case <-isReadyChan:
		err = nil
	case readErr := <-readErrChan:
		err = readErr
	case <-timer.C:
		err = ErrTimeout
	}

	return err
}

// wait waits for hid_listen to exit, which should never happen since we expect it to run forever.
// So if it does happen we will notify that as an error through the specified error channel
func (h *HidListen) wait(errChan chan<- error) {
	exitErr := fmt.Errorf("hid_listen terminated unexpectedly")

	err := h.cmd.Wait()
	if err != nil {
		exitErr = fmt.Errorf("error while waiting for hid_listen: %s", err)
	}

	errChan <- exitErr
}

func (h *HidListen) watchStdErr(errChan chan<- error) {
	scanner := bufio.NewScanner(h.StdErrPipe)
	for scanner.Scan() {
		errChan <- errors.New(scanner.Text())
	}
}

// Start will start the HidListen.cmd command and wait until it is ready before returning.
// It will return a error in case the command fails to start or get ready. It will also return
// a channel that will receive errors that can happen after hid_listen starts running
func (h *HidListen) Start() (<-chan error, error) {
	errChan := make(chan error)

	go h.watchStdErr(errChan)

	err := h.cmd.Start()
	if err != nil {
		return nil, err
	}

	err = h.waitUntilReady(10 * time.Second)
	if err != nil {
		return nil, err
	}

	go h.wait(errChan)

	return errChan, nil
}
