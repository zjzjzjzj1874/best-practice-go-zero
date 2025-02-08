/*
 * @Author: zjzjzjzj1874 zjzjzjzj1874@gmail.com
 * @Date: 2025-01-23 14:46:03
 * @LastEditors: zjzjzjzj1874 zjzjzjzj1874@gmail.com
 * @LastEditTime: 2025-01-24 11:50:03
 * @FilePath: /best-practice-go-zero/__test__/chromedp/main.go
 * @Description: chromedp 测试
 */
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/chromedp/chromedp"
)

func main() {
	// 创建带取消功能的上下文
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// 创建变量存储结果
	var outerHTML string
	var innerHtml string

	// 执行操作
	err := chromedp.Run(ctx,
		// 导航到目标页面
		chromedp.Navigate("http://www.people.com.cn/"),
		// 等待页面加载完成
		chromedp.WaitReady("body"),
		// 获取 body 的 outerHTML
		chromedp.OuterHTML("body", &outerHTML),
		// 获取 body 的 innerHtml
		chromedp.InnerHTML("body", &innerHtml),
	)

	if err != nil {
		log.Fatal("执行失败:", err)
	}

	// 输出结果
	// fmt.Println(outerHTML)
	fmt.Println(innerHtml)
}
