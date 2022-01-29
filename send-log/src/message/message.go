package message

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func WxSend(message string) {
	body := make(map[string]interface{})
	body["msgtype"] = "text"
	body["text"] = map[string]string{
		"content": message,
	}
	b, _ := json.Marshal(body)
	http.DefaultClient.Post("http://10.1.250.78:8092/access/testString", "application/json", bytes.NewBuffer(b))
}
