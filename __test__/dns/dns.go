package main

import (
	"fmt"
	"net"
	"regexp"
	"strings"
)

func main() {
	//net.LookupMX()
}

func validateEmail(email string) {
	if !isValidEmail(email) {
		fmt.Println("邮箱格式有误，请重新输入：", email)
		return
	}

	domain := email[strings.LastIndex(email, "@")+1:]
	mx, err := net.LookupMX(domain)
	if err != nil {
		fmt.Println("邮箱域名有误：", domain)
		return
	}
	for _, m := range mx {
		fmt.Println(m.Host)
	}
}

func isValidEmail(email string) bool {
	//re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$`)
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}
