package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os/exec"
	"runtime"
	"strings"
	"sync"
	"syscall"
	"time"
)

var cmdArg [2]string

func init() {
	if runtime.GOOS == "windows" {
		cmdArg[0] = "cmd"
		cmdArg[1] = "/c"
	} else {
		cmdArg[0] = "/bin/bash"
		cmdArg[1] = "-c"
	}
}

type LineReader interface {
	Read(content string, numOfLines int)
}

type logHolder struct {
	lock    *sync.Mutex
	content *bytes.Buffer
}

func (l *logHolder) Read(content string, numOfLines int) {
	l.lock.Lock()
	l.content.WriteString(content)
	l.lock.Unlock()
}

func (l *logHolder) String() string {
	l.lock.Lock()
	defer l.lock.Unlock()
	return l.content.String()
}

func (l *logHolder) Reset() *logHolder {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.content.Reset()
	return l
}

func NewLogHolder() *logHolder {
	return &logHolder{content: &bytes.Buffer{}, lock: &sync.Mutex{}}
}

func readOut(outReader *bufio.Reader, lineReader LineReader) {
	const batchOfLines = 10
	lines := make([]string, 0, batchOfLines)
	for {
		var line string
		var err error
		line, err = outReader.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		} else {
			lines = append(lines, line)
		}

		numOfLines := len(lines)
		if numOfLines >= batchOfLines {
			originalLogs := strings.Join(lines, "")
			if lineReader != nil {
				lineReader.Read(originalLogs, numOfLines)
			}
			lines = make([]string, 0, batchOfLines)
		}
	}
	if len(lines) >= 0 {
		if lineReader != nil {
			lineReader.Read(strings.Join(lines, "\n"), len(lines))
		}
		lines = make([]string, 0, batchOfLines)
	}
}

// RunWithCmd 同步函数, 执行命令直到完成或超时
func RunWithCmd(cmd *exec.Cmd, lineReader LineReader, timeout time.Duration) (int, error) {
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return 0, err
	}
	defer stdout.Close()
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return 0, err
	}
	defer stderr.Close()

	err = cmd.Start()
	if err != nil {
		return 0, err
	}
	stdoutReader := bufio.NewReader(stdout)
	stderrReader := bufio.NewReader(stderr)

	go readOut(stdoutReader, lineReader)
	go readOut(stderrReader, lineReader)
	err, exitCode := waitForComplete(cmd, timeout)
	return exitCode, err
}

func waitForComplete(cmd *exec.Cmd, timeout time.Duration) (error, int) {
	var err error
	timer := time.NewTimer(timeout)
	done := make(chan error)
	go func() {
		defer func() {
			innerErr := recover()
			if innerErr != nil {
				if err, ok := innerErr.(error); ok {
					done <- err
				} else if errStr, ok := innerErr.(string); ok {
					done <- errors.New(errStr)
				} else {
					done <- errors.New("panic on executing command")
				}
			}
			close(done)
		}()
		done <- cmd.Wait()
	}()

	var exitCode int
	select {
	case err = <-done:
		timer.Stop()
		if cmd.ProcessState != nil {
			exitCode = cmd.ProcessState.ExitCode()
		}
	case <-timer.C:
		if err = syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL); err != nil {
			err = fmt.Errorf("command timed out and killing process fail: %s", err.Error())
		} else {
			// wait for the command to return after killing it
			<-done
			err = errors.New("command timed out")
		}
	}

	return err, exitCode
}

func PrepareCmd(scriptFile string, dir string, env []string, args ...string) *exec.Cmd {
	scriptArgs := make([]string, 0, len(args)+1)
	scriptArgs = append(scriptArgs, scriptFile)
	if len(args) > 0 {
		scriptArgs = append(scriptArgs, args...)
	}
	cmd := exec.Command(cmdArg[0], scriptArgs...)
	cmd.Dir = dir
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	if env != nil {
		cmd.Env = env
	}
	return cmd
}
