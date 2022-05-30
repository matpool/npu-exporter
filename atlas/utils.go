package main

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
)

const (
	_ = iota // ignore first value by assigning to blank identifier
	// KiB kilobyte
	KiB = 1 << (10 * iota) // 1 << (10*1)
	// MiB megabyte
	MiB // 1 << (10*2)
	// GiB gigabyte
	GiB // 1 << (10*3)
	// TiB terabyte
	TiB // 1 << (10*4)
)

// ExecBashShell exec bash shell command
func ExecBashShell(arg string) (string, error) {
	return ExecBashShellWithCtx(context.Background(), arg)
}

// ExecBashShellWithCtx exec bash shell with ctx
// when ctx is done it'll kill corresponding process
func ExecBashShellWithCtx(ctx context.Context, arg string) (string, error) {
	cmd := exec.CommandContext(ctx, "bash", "-c", arg)
	return execShell(cmd)
}

// ExecPowerShell exec power shell command
func ExecPowerShell(arg string) (string, error) {
	return ExecPowerShellWithCtx(context.Background(), arg)
}

// ExecPowerShellWithCtx exec power shell command with ctx
// when ctx is done it'll kill corresponding process
func ExecPowerShellWithCtx(ctx context.Context, arg string) (string, error) {
	cmd := exec.CommandContext(ctx, "powershell", arg)
	return execShell(cmd)
}

func execShell(cmd *exec.Cmd) (string, error) {
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("exec shell failed, out_msg:%s, err_msg:%s, err:%s", stdout.String(), stderr.String(), err.Error())
	}
	return stdout.String(), nil
}
