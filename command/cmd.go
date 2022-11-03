package command

import (
	"bufio"
	"context"
	"os/exec"
	"time"

	"github.com/skycandyzhe/go-com/logger"
)

func StartCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	return cmd.Start()
}

// sync execute
func ExecCommand(name string, args ...string) ([]byte, error) {
	return ExecCommandWithTimeout(10*time.Second, name, args...)
}

// execute
func ExecCommandWithTimeout(timeout time.Duration, name string, args ...string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, name, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		if err == context.DeadlineExceeded || err == exec.ErrNotFound {
			return nil, err
		}
	}

	return out, nil
}

func SyncCmdExec(log logger.MyLoggerInterface, name string, args ...string) (out *bufio.Scanner, err *bufio.Scanner) {
	// go func() {
	// 	cmd := exec.Command(
	// 		name, args...,
	// 	)

	// 	reader, err := cmd.StdoutPipe()
	// 	if err != nil {
	// 		log.Errorf("cmd bind stdout err:%v", err)
	// 		return
	// 	}
	// 	defer reader.Close()
	// 	errReader, err := cmd.StderrPipe()
	// 	if err != nil {
	// 		log.Errorf("cmd bind stderr err:%v", err)
	// 		return
	// 	}
	// 	defer errReader.Close()
	// 	err = cmd.Start()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	out = bufio.NewScanner(reader)
	// 	err = bufio.NewScanner(errReader)

	// }()
	return nil, nil
}
