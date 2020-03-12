package dingtalk

import (
	"bytes"
	"encoding/json"
	"github.com/tech-botao/logger"
	"github.com/tech-botao/network/rest"
	"net/http"
	"os"
)

type (
	Markdown struct {
		Title string `json:"title"`
		Text  string `json:"text"`
	}

	Message struct {
		MsgType  string `json:"msgtype"`
		Markdown Markdown `json:"markdown"`
	}
)

var _url = ""
var client = rest.NewHttpClient(http.DefaultClient)

func SetUrl(u string) {
	_url = u
}

func LoadEnv() {
	_url = os.Getenv("DINGTALK_URL")
	if len(_url) == 0 {
		logger.Info("[dingtalk] env = DINGTALK_URL", nil)
	}
}

func NewMessage() *Message {
	if len(_url) == 0 {
		LoadEnv()
	}
	return &Message{
		MsgType: "markdown",
		Markdown:Markdown{
			Title: "default",
			Text:  "test",
		},
	}
}

func (m *Message) Type(s string) *Message {
	m.MsgType = s
	return m
}

func (m *Message) Title(s string) *Message {
	m.Markdown.Title = s
	return m
}

func (m *Message) Text(s string) *Message {
	m.Markdown.Text = s
	return m
}

// Send 发送一个通知, 只记录Log， 不报错
func (m *Message) Send() {
	if len(_url) == 0 {
		logger.Warn("[dingtalk] env = DINGTALK_URL, or use SetUrl()", nil)
		return
	}
	b, err := json.Marshal(m)
	if err != nil {
		logger.Warn("[dingtalk] encoding error, ", err)
		return
	}
	req, err := http.NewRequest("POST", _url, bytes.NewReader(b))
	if err != nil {
		logger.Warn("[dingtalk] new request error, ", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	err = client.Do(req, nil)
	if err != nil {
		logger.Warn("[dingtalk] Send() error, ", err)
		return
	}
}

