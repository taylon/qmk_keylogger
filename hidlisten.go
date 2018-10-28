package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os/exec"
	"time"
)

type HidListen struct {
	cmd *exec.Cmd

	StdOutputPipe io.ReadCloser
	StdErrPipe    io.ReadCloser
}

func NewHidListen() (*HidListen, error) {
	cmd := exec.Command("./hid_listen")

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
func (h *HidListen) waitUntilReady() error {
	reader := bufio.NewReader(h.StdOutputPipe)

	timer := time.NewTimer(10 * time.Second)
	defer timer.Stop()

	isReady := make(chan bool)
	var err error

	go func() {
		var line string

		for {
			// If this returns error nothing will ever be sent to the isReady channel and the error will
			// be picked up bellow on the timeout block of the select statement
			line, err = reader.ReadString('\n')
			if err != nil {
				return
			}

			if line == "Listening:\n" {
				isReady <- true
				return
			}
		}
	}()

	select {
	case <-isReady:
		return nil
	case <-timer.C:
		timeoutMsg := "timed out while waiting for hid_listen to be ready"

		if err != nil {
			timeoutMsg = fmt.Sprintf("%s: %s", timeoutMsg, err)
		}

		return fmt.Errorf("%s", timeoutMsg)
	}
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
func (h *HidListen) Start() (error, <-chan error) {
	errChan := make(chan error)

	go h.watchStdErr(errChan)

	err := h.cmd.Start()
	if err != nil {
		return fmt.Errorf("could not start hid_listen: %s", err), nil
	}

	err = h.waitUntilReady()
	if err != nil {
		return err, nil
	}

	go h.wait(errChan)

	return nil, errChan
}
