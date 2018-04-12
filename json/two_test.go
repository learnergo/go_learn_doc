package doc_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Server struct {
	ServerName string
	ServerIP   string
}

type Serverslice struct {
	Servers []Server
}

func Test_rule(t *testing.T) {
	var s Serverslice
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)
	//结果{[{Shanghai_VPN 127.0.0.1} {Beijing_VPN 127.0.0.2}]}
	//输出了正确结果，根据serverName找到ServerName
}
