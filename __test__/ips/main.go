/*
* @Author: zjzjzjzj1874 zjzjzjzj1874@gmail.com
* @Date: 2024-12-30 17:35:30
 * @LastEditors: zjzjzjzj1874 zjzjzjzj1874@gmail.com
 * @LastEditTime: 2025-01-17 17:56:18
* @FilePath: /best-practice-go-zero/__test__/ips/main.go
* @Description: ips
*/
package main

import (
	"fmt"
	"net"
	"time"

	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
)

// 根据IP来获取国家和公司名
func main() {
	var domain = "baidu.com"
	// var domain = "www.baidu.com"
	// var domain = "www.toor.fun"
	// var domain = "toor.fun"
	// var domain = "http://www.toor.fun"
	var ipStrs []string
	ips, _ := net.LookupIP(domain)
	for _, ip := range ips {
		ipStrs = append(ipStrs, ip.String())
	}
	fmt.Println("ips:", ipStrs)

	var dbPath = "/Users/zjzjzjzj1874/githubProject/ip2region/data/ip2region.xdb"
	searcher, err := xdb.NewWithFileOnly(dbPath)
	if err != nil {
		fmt.Printf("failed to create searcher: %s\n", err.Error())
		return
	}

	defer searcher.Close()

	// for _, ip := range ipStrs {
	// 	var tStart = time.Now()
	// 	region, err := searcher.SearchByStr(ip)
	// 	if err != nil {
	// 		fmt.Printf("failed to SearchIP(%s): %s\n", ip, err)
	// 		return
	// 	}

	// 	fmt.Printf("{region: %s, took: %s}\n", region, time.Since(tStart))
	// }
	// do the search
	var ip = "124.221.82.130"
	// var ip = "38.238.107.103"
	var tStart = time.Now()
	region, err := searcher.SearchByStr(ip)
	if err != nil {
		fmt.Printf("failed to SearchIP(%s): %s\n", ip, err)
		return
	}

	fmt.Printf("{ip:%s, region: %s, took: %s}\n", ip, region, time.Since(tStart))

	// 备注：并发使用，每个 goroutine 需要创建一个独立的 searcher 对象。
}
