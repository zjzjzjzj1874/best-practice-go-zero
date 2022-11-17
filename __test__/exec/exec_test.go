package exec

import (
	"bytes"
	"log"
	"os/exec"
	"testing"
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

}
