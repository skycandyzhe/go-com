package command

import (
	"context"
	"os/exec"
	"time"
)

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
