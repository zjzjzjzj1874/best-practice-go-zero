package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	// 假设的数据
	AKId := "60e475259130210dbb102375c5c50518"
	timestamp := fmt.Sprintf("%d", time.Now().UnixMilli())
	nonce := fmt.Sprintf("%d", rand.Intn(10000))

	// 连接所有字符串
	combined := "2e72b684bcc62da9bb254e6f6ac0c9f4" + timestamp + nonce

	// 创建一个新的哈希
	h := sha1.New()
	h.Write([]byte(combined))
	bs := h.Sum(nil)
	sign := hex.EncodeToString(bs)

	fmt.Println("sign ======= ", sign)

	req, _ := http.NewRequest(http.MethodGet, "https://openapi.dun.163.com/openapi/v2/antispam/label/query", nil)
	req.Header.Set("X-YD-SECRETID", AKId)
	req.Header.Set("X-YD-TIMESTAMP", timestamp)
	req.Header.Set("X-YD-NONCE", nonce)
	req.Header.Set("X-YD-SIGN", sign)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}

	fmt.Println(string(body))
}
