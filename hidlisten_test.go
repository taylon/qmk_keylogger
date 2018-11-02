package main

import (
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWaitUntilReady(t *testing.T) {
	hidListen := &HidListen{
		StdOutputPipe: ioutil.NopCloser(strings.NewReader(hidListenReadyLine)),
	}

	err := hidListen.waitUntilReady(10 * time.Second)

	assert.NoError(t, err)
}

func TestWaitUntilReadyTimeout(t *testing.T) {
	// We build a big input because we want to make sure that it times out. A small input
	// would read too fast and return io.EOF (since all inputs would be read) before timing out
	giganticInput := "no_ready_line\n"
	for i := 1; i <= 10; i++ {
		giganticInput = giganticInput + giganticInput + giganticInput
	}

	hidListen := &HidListen{
		StdOutputPipe: ioutil.NopCloser(strings.NewReader(giganticInput)),
	}

	err := hidListen.waitUntilReady(1 * time.Nanosecond)

	assert.Equal(t, err, ErrTimeout)
}

func TestWaitUntilReadyReadError(t *testing.T) {
	hidListen := &HidListen{
		StdOutputPipe: ioutil.NopCloser(strings.NewReader("no_ready_line")),
	}

	err := hidListen.waitUntilReady(10 * time.Second)

	assert.Equal(t, err, io.EOF)
}

func TestWatchStdErr(t *testing.T) {
	expectedErrString := "error"

	hidListen := &HidListen{
		StdErrPipe: ioutil.NopCloser(strings.NewReader(expectedErrString)),
	}

	errChan := make(chan error)
	go hidListen.watchStdErr(errChan)

	err := <-errChan

	assert.EqualError(t, err, expectedErrString)
}

func TestWaitFailsWhenCmdWaitReturnsError(t *testing.T) {
	hidListen := &HidListen{
		cmd: exec.Command("kill_mx_browns"),
	}

	errChan := make(chan error)
	go hidListen.wait(errChan)

	err := <-errChan

	assert.Error(t, err)
}

func TestWaitReturnsErrorIfCommandEverExits(t *testing.T) {
	hidListen := &HidListen{
		cmd: exec.Command(os.Args[0], "-test.run=TestHelperProcess", "--"),
	}

	errChan := make(chan error)
	hidListen.cmd.Start()
	go hidListen.wait(errChan)

	err := <-errChan

	assert.Error(t, err)
}

func TestStartFailsWhenCmdDoesNotExist(t *testing.T) {
	hidListen := &HidListen{
		cmd:           exec.Command("kill_mx_browns"),
		StdErrPipe:    ioutil.NopCloser(strings.NewReader("")),
		StdOutputPipe: ioutil.NopCloser(strings.NewReader(hidListenReadyLine)),
	}

	_, err := hidListen.Start()

	assert.Error(t, err)
}
