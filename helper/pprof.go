package helper

import (
	"fmt"
	"log"
	"net/http"
)

func OpenPPROF(port int) {
	fmt.Printf("listen on %d ...\n", port)

	err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), nil)
	if err != nil {
		log.Fatalf("open pprof failure:[err:%s]", err.Error())
	}
}
