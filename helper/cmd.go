package helper

import (
	"bytes"
	"context"
	"os/exec"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
)

// CommandWithTimeOut 带超时时间的exec 执行器
func CommandWithTimeOut(timeout time.Duration, name string, args ...string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	cmd := exec.CommandContext(ctx, name, args...)
	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf

	if err := cmd.Start(); err != nil {
		logrus.Errorf("CommandContext Start err=%s", err.Error())
		return string(buf.Bytes()), err
	}
	logrus.Infof("cmd.String = %s,cmd.Stderr=%s", cmd.String(), cmd.Stderr)
	waitChan := make(chan struct{}, 1)
	defer close(waitChan)
	// 超时杀掉进程组 或正常退出
	go func() {
		select {
		case <-ctx.Done():
			logrus.Errorf("timeout kill ppid:%d", cmd.Process.Pid)
			syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
		case <-waitChan:
		}
	}()
	if err := cmd.Wait(); err != nil {
		logrus.Errorf("CommandContext Wait err=%s,Stdout=%s,Stderr=%s", err.Error(), cmd.Stdout, cmd.Stderr)
		return string(buf.Bytes()), err
	}

	return string(buf.Bytes()), nil
}
