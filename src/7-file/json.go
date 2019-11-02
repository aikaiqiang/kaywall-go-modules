package main

import (
	"encoding/json"
	"fmt"

	. "github.com/bitly/go-simplejson"
)

type Server struct {
	ServerName string
	ServerIP   string
}

type Serverslice struct {
	Servers []Server
}

func readJson() {
	var s Serverslice
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)
}

func simplejson() {
	js, err := NewJson([]byte(`{
		"test": {
			"array": [1, "2", 3],
			"int": 10,
			"float": 5.150,
			"bignum": 9223372036854775807,
			"string": "simplejson",
			"bool": true
		}
	}`))
	if err != nil {
		fmt.Printf("11-error: %v", err)
	}

	arr, _ := js.Get("test").Get("array").Array()
	fmt.Println(arr)

	i, _ := js.Get("test").Get("int").Int()
	fmt.Println(i)

	ms := js.Get("test").Get("string").MustString()
	fmt.Println(ms)
}

func writeJson() {
	var s Serverslice
	s.Servers = append(s.Servers, Server{ServerName: "Shanghai_VPN", ServerIP: "127.0.0.1"})
	s.Servers = append(s.Servers, Server{ServerName: "Beijing_VPN", ServerIP: "127.0.0.2"})
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
}

func main() {

	readJson()

	simplejson()

	writeJson()
}
