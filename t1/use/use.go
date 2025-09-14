package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// 使用go语言封装调用第三方接口
func main() {
	// 假设这是一个第三方接口的URL
	url := "https://api.example.com/data"

	// 构建请求体（JSON格式）
	requestBody := map[string]string{
		"param1": "value1",
		"param2": "value2",
	}
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// 创建POST请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// 设置请求头（根据需要设置，例如Content-Type）
	req.Header.Set("Content-Type", "application/json")

	// 发送请求并接收响应
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// 处理响应
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: received status code %d", resp.StatusCode)
		return
	}

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// 打印响应体（或进行其他处理）
	fmt.Println(string(body))
}
