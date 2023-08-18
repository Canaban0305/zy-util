package zlog

import (
	"encoding/json"
	"fmt"
	"github.com/zhiyunai/zy-util/dc"
	"golang.org/x/exp/slog"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

var (
	ServerName = "ZLog"  // ServerName 应用名称
	Version    = "1.0.0" // Version 应用版本
	LogsUrl    = "http://logs.zhiyunai.com.cn/api/default/%s/_json"
)

type ZLog struct {
	log *slog.Logger
}

type Message struct {
	Project string `json:"project,omitempty"`
	Version string `json:"version,omitempty"`
	Content string `json:"content,omitempty"`
	Level   string `json:"level,omitempty"`
	IP      string `json:"IP,omitempty"`
}

// New 创建 zlog 实例
func New(serverName, ver string) *ZLog {
	ServerName = serverName
	Version = ver
	LogsUrl = fmt.Sprintf(LogsUrl, ServerName)
	return &ZLog{}
}

func (z *ZLog) Info(msg string) {
	log.Println(time.Now().Format("2006-01-02 15:04:05"), "[INFO]", msg)
	ip := dc.GetIP()
	arr := make([]Message, 0)
	param := Message{
		IP:      ip,
		Content: msg,
		Project: ServerName,
		Version: Version,
		Level:   "INFO",
	}
	arr = append(arr, param)

	post(LogsUrl, arr)
}

func (z *ZLog) Debug(msg string) {
	log.Println(time.Now().Format("2006-01-02 15:04:05"), "[DEBUG]", msg)
	ip := dc.GetIP()
	arr := make([]Message, 0)
	param := Message{
		IP:      ip,
		Content: msg,
		Project: ServerName,
		Version: Version,
		Level:   "DEBUG",
	}
	arr = append(arr, param)

	post(LogsUrl, arr)
}

func (z *ZLog) Warn(msg string) {
	log.Println(time.Now().Format("2006-01-02 15:04:05"), "[WARN]", msg)
	ip := dc.GetIP()
	arr := make([]Message, 0)
	param := Message{
		IP:      ip,
		Content: msg,
		Project: ServerName,
		Version: Version,
		Level:   "WARN",
	}
	arr = append(arr, param)

	post(LogsUrl, arr)
}

func (z *ZLog) Error(msg string) {
	log.Println(time.Now().Format("2006-01-02 15:04:05"), "[ERROR]", msg)
	ip := dc.GetIP()
	arr := make([]Message, 0)
	param := Message{
		IP:      ip,
		Content: msg,
		Project: ServerName,
		Version: Version,
		Level:   "ERROR",
	}
	arr = append(arr, param)

	post(LogsUrl, arr)
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

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))

	return
}
