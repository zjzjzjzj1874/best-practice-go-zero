package main

import (
	"log"
	"testing"
)

func TestCheckNude(t *testing.T) {
	t.Run("#Nude", func(t *testing.T) {
		isNude, err := CheckNude("./nude.jpeg")
		if err != nil {
			log.Fatal(err)
		}

		t.Log("Result:", isNude)
	})

	t.Run("#Normal", func(t *testing.T) {
		isNude, err := CheckNude("./normal.jpg")
		if err != nil {
			log.Fatal(err)
		}

		t.Log("Result:", isNude)
	})
}
