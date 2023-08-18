package zlog

import (
	"encoding/json"
	"fmt"
	"github.com/zhiyunai/zy-util/dc"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

var (
	logUrl  = "http://logs.zhiyunai.com.cn/api/default/%s/_json"
	_config *Config // 配置项
)

type ZLog struct {
}

// Config 配置项
type Config struct {
	ServerName string // 应用名称(openobserve 仓库名)
	Version    string // 应用版本
	ConsoleLog bool   // 是否打印到控制台
	IsUpload   bool   // 是否上传
}

// Message Open-observe 实体
type Message struct {
	Project   string `json:"project"`
	Version   string `json:"version"`
	Content   string `json:"content"`
	Level     string `json:"level"`
	IP        string `json:"ip"`
	TimeStamp string `json:"timestamp"`
}

// New 创建 zlog 实例
func New(config *Config) *ZLog {
	if config == nil {
		log.Fatal("zlog must has config.")
	}

	_config = config
	if _config.ServerName == "" {
		_config.ServerName = "default"
	}

	logUrl = fmt.Sprintf(logUrl, _config.ServerName)
	return &ZLog{}
}

func (z *ZLog) Info(a ...interface{}) {
	msg := z.getAnyString(a)
	printf("info", msg)
	if !_config.IsUpload {
		return
	}
	ip := dc.GetIP()
	arr := make([]Message, 0)
	param := Message{
		IP:        ip,
		Content:   msg,
		Project:   _config.ServerName,
		Version:   _config.Version,
		Level:     "info",
		TimeStamp: time.Now().Format("2006/01/02 15:04:05"),
	}
	arr = append(arr, param)

	post(logUrl, arr)
}

func (z *ZLog) Debug(a ...interface{}) {
	msg := z.getAnyString(a)
	printf("debug", msg)
	if !_config.IsUpload {
		return
	}
	ip := dc.GetIP()
	arr := make([]Message, 0)
	param := Message{
		IP:        ip,
		Content:   msg,
		Project:   _config.ServerName,
		Version:   _config.Version,
		Level:     "debug",
		TimeStamp: time.Now().Format("2006/01/02 15:04:05"),
	}
	arr = append(arr, param)

	post(logUrl, arr)
}

func (z *ZLog) Warn(a ...interface{}) {
	msg := z.getAnyString(a)
	printf("warn", msg)
	if !_config.IsUpload {
		return
	}
	ip := dc.GetIP()
	arr := make([]Message, 0)
	param := Message{
		IP:        ip,
		Content:   msg,
		Project:   _config.ServerName,
		Version:   _config.Version,
		Level:     "warn",
		TimeStamp: time.Now().Format("2006/01/02 15:04:05"),
	}
	arr = append(arr, param)

	post(logUrl, arr)
}

func (z *ZLog) Error(a ...interface{}) {
	msg := z.getAnyString(a)
	printf("error", msg)
	if !_config.IsUpload {
		return
	}
	ip := dc.GetIP()
	arr := make([]Message, 0)
	param := Message{
		IP:        ip,
		Content:   msg,
		Project:   _config.ServerName,
		Version:   _config.Version,
		Level:     "error",
		TimeStamp: time.Now().Format("2006/01/02 15:04:05"),
	}
	arr = append(arr, param)

	post(logUrl, arr)
}

func (z *ZLog) getAnyString(a ...interface{}) string {
	str := ""

	for _, arg := range a {
		for i2, arg2 := range arg.([]interface{}) {
			if i2 > 0 {
				str += " "
			}

			str += fmt.Sprintf("%v", arg2)
		}
	}

	return str
}

// post HttpPost
func post(url string, data interface{}) {
	method := "POST"
	jsonBytes, _ := json.Marshal(data)
	//fmt.Println(string(jsonBytes))
	payload := strings.NewReader(string(jsonBytes))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("User-Agent", "Apifox/1.0.0 (https://apifox.com)")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic MTIyNTg0MjkwNUBxcS5jb206S0ZmdUZjRGZhUXlYOFFKeQ==")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Host", "logs.zhiyunai.com.cn")
	req.Header.Add("Connection", "keep-alive")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(string(body))

	return
}

func printf(a ...interface{}) {
	if _config.ConsoleLog {
		fmt.Println(a)
	}
}
