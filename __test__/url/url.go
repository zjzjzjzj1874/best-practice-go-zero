package main

import (
	"fmt"
	"net/url"
)

func main() {
	// 示例 URL
	link := "https://example.com/path?query=123"

	// 解析 URL
	parsedUrl, err := url.Parse(link)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	// 判断协议
	if parsedUrl.Scheme == "http" {
		parsedUrl.Scheme = "https"
	} else if parsedUrl.Scheme == "https" {
		parsedUrl.Scheme = "http"
	}
	if parsedUrl.Host == "example.com" {
		parsedUrl.Host = "examples.com"
	}

	query := parsedUrl.Query()
	query.Set("req_content", "html")
	parsedUrl.RawQuery = query.Encode()

	fmt.Println(parsedUrl.String())
}
