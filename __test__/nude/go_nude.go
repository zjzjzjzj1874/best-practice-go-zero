package main

import (
	"fmt"
	"log"

	"github.com/koyachi/go-nude"
)

func main() {
	imagePath := "./__test__/nude/nude.jpeg"

	isNude, err := nude.IsNude(imagePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("isNude = %v\n", isNude)
}

// CheckNude 根据皮肤暴露的比例来判断是否裸露的
func CheckNude(path string) (bool, error) {
	isNude, err := nude.IsNude(path)
	if err != nil {
		return false, err
	}

	return isNude, nil
}
