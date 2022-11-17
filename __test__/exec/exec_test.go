package exec

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os/exec"
	"testing"
	"time"
)

// detail see this article: https://colobu.com/2020/12/27/go-with-os-exec/
func TestCMD(t *testing.T) {
	t.Run("#catch cmd output", func(t *testing.T) {
		cmd := exec.Command("ls", "-lah")
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			log.Fatalf("failed to call cmd.Run(): %v", err)
		}
		t.Logf(out.String())
	})

	t.Run("#show downloading process", func(t *testing.T) {
		cmd := exec.Command("curl", "https://dl.google.com/go/go1.15.6.linux-amd64.tar.gz")
		var stdoutProcessStatus bytes.Buffer
		cmd.Stdout = io.MultiWriter(ioutil.Discard, &stdoutProcessStatus)
		done := make(chan struct{})
		go func() {
			tick := time.NewTicker(time.Second)
			defer tick.Stop()
			for {
				select {
				case <-done:
					return
				case <-tick.C:
					log.Printf("downloaded: %d KB", stdoutProcessStatus.Len()/1024)
					log.Printf("downloaded: %d MB", stdoutProcessStatus.Len()/1024/1024)
				}
			}
		}()
		err := cmd.Run()
		if err != nil {
			log.Fatalf("failed to call Run(): %v", err)
		}
		close(done)
	})

}
