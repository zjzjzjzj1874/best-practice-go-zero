package gbk

import (
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
)

func encodeSimpleChinese(name string) {
	encoder := simplifiedchinese.GB18030.NewEncoder()
	name, err := encoder.String(name)
	if err != nil {
		fmt.Println("encoder Error:", err.Error())
		return
	}

	fmt.Printf("name=%s\n", name)

	dec := simplifiedchinese.GB18030.NewDecoder()
	deName, err := dec.String(name)
	if err != nil {
		fmt.Println("decode Error:", err.Error())
		return
	}
	fmt.Println("deName == ", deName)
}
