package dingtalk

import (
	"github.com/tech-botao/atom"
	"testing"
)

func init() {
	atom.SetEnvPath("/etc/go")
	atom.LoadEnv("dingtalk.env")
}

func TestMessage_Send(t *testing.T) {
	m := NewMessage()
	type args struct {
		MsgType  string
		Text  string
		Title string
	}
	tests := []struct {
		name   string
		args args
	}{
		{"ok", args{ MsgType: "markdown", Text:    "test", Title:   "title", }},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m.Type(tt.args.MsgType).Title(tt.args.Title).Text(tt.args.Text).Send()
		})
	}
}

func TestMessage_Text(t *testing.T) {
	t.Skip("Text() is simple, dont test")
}

func TestMessage_Title(t *testing.T) {
	t.Skip("Title() is simple, dont test")
}

func TestMessage_Type(t *testing.T) {
	t.Skip("Type() is simple, dont test")
}

func TestNewMessage(t *testing.T) {
	t.Skip("NewMessage() is simple, dont test")
}

func TestSetUrl(t *testing.T) {
	t.Skip("SetUrl() is simple, dont test")
}

func TestLoadEnv(t *testing.T) {
	t.Skip("Text() is simple, dont test")
}

