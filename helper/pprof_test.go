package helper

import (
	_ "net/http/pprof"
	"testing"
)

func TestOpenPPROF(t *testing.T) {
	// open pprof see detail in http://localhost:9999/debug/pprof
	t.Run("#Open pprof", func(t *testing.T) {
		OpenPPROF(PprofConf{Debug: true, DebugPort: 9999})
		select {}
	})
}
