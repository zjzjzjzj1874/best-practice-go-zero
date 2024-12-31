/*
* @Author: zjzjzjzj1874 zjzjzjzj1874@gmail.com
* @Date: 2024-12-30 17:35:30
 * @LastEditors: zjzjzjzj1874 zjzjzjzj1874@gmail.com
 * @LastEditTime: 2024-12-31 10:14:30
* @FilePath: /best-practice-go-zero/__test__/ips/main.go
* @Description: ips
*/
package main

import (
	"fmt"
	"time"

	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
)

// 根据IP来获取国家和公司名
func main() {
	var dbPath = "/Users/zjzjzjzj1874/githubProject/ip2region/data/ip2region.xdb"
	searcher, err := xdb.NewWithFileOnly(dbPath)
	if err != nil {
		fmt.Printf("failed to create searcher: %s\n", err.Error())
		return
	}

	defer searcher.Close()

	// do the search
	var ip = "124.70.5.204"
	// var ip = "38.238.107.103"
	var tStart = time.Now()
	region, err := searcher.SearchByStr(ip)
	if err != nil {
		fmt.Printf("failed to SearchIP(%s): %s\n", ip, err)
		return
	}

	fmt.Printf("{region: %s, took: %s}\n", region, time.Since(tStart))

	// 备注：并发使用，每个 goroutine 需要创建一个独立的 searcher 对象。
}
