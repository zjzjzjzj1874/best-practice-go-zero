package main

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/trace"
	"log"
	"os"
	"os/signal"
)

func main() {
	l := log.New(os.Stdout, "", 0)

	// Write telemetry data to a file.
	f, err := os.Create("traces.txt")
	if err != nil {
		l.Fatal(err)
	}
	defer f.Close()

	exp, err := newExporter(f)
	if err != nil {
		l.Fatal(err)
	}

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exp),
		trace.WithResource(newResource()),
	)
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			l.Fatal(err)
		}
	}()
	otel.SetTracerProvider(tp)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	errCh := make(chan error)
	app := NewApp(os.Stdin, l)
	go func() {
		errCh <- app.Run(context.Background())
	}()

	select {
	case <-sigCh:
		l.Println("\ngoodbye")
		return
	case err := <-errCh:
		if err != nil {
			l.Fatal(err)
		}
	}
}

// Fibonacci returns the n-th fibonacci number.
func Fibonacci(n uint) (uint64, error) {
	if n <= 1 {
		return uint64(n), nil
	}

	var n2, n1 uint64 = 0, 1
	for i := uint(2); i < n; i++ {
		n2, n1 = n1, n1+n2
	}

	return n2 + n1, nil
}
