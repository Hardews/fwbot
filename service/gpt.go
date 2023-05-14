/**
 * @Author: Hardews
 * @Date: 2023/5/13 23:29
 * @Description:
**/

package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"fwbot/config"
	"io"
	"net/http"
	"strings"
	"time"
)

const (
	url      = "https://api.openai.com/v1/chat/completions"
	gptModel = "gpt-3.5-turbo"
	role     = "user"
)

type send struct {
	Model    string
	Messages []struct {
		Role    string
		Content string
	}
}

type accept struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func GptSend(msg string) (result string, err error) {
	var c = config.Config
	// 用户设置 或 默认值
	var r = role
	if c.Role != "" {
		r = c.Role
	}

	var model = gptModel
	if c.Model != "" {
		model = c.Model
	}

	// 封装需要发送的 json 字段
	var s = send{
		Model: model,
		Messages: []struct {
			Role    string
			Content string
		}{{r, msg}},
	}

	var (
		req    *http.Request
		resp   *http.Response
		client = &http.Client{Timeout: 10 * time.Minute} // 10分钟过期
	)

	// 解析 body
	bodyByte, err := json.Marshal(s)
	if err != nil {
		err = errors.New("marshal body failed,err:" + err.Error())
		return
	}

	body := strings.NewReader(string(bodyByte))

	// 设置请求头
	req, err = http.NewRequest(http.MethodPost, url, body)
	req.Header.Set("Authorization", "Bearer "+c.ApiKey)
	req.Header.Set("Accept", "application/json")
	if c.Organization != "" {
		req.Header.Set("OpenAI-Organization", c.Organization)
	}

	// 发送请求
	resp, err = client.Do(req)
	if err != nil {
		err = errors.New("send req failed,err:" + err.Error())
		return
	}

	defer resp.Body.Close()

	// 读取响应字段
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		err = errors.New("read resp body failed,err:" + err.Error())
		return
	}

	var respMsg accept
	fmt.Println(string(respBody))
	// 解析
	json.Unmarshal(respBody, &respMsg)

	fmt.Println(respMsg)
	if len(respMsg.Choices) == 0 {
		return "报错咯", nil
	}
	// 返回
	return respMsg.Choices[0].Message.Content, nil
}
