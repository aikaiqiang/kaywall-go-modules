package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Recurlyservers struct {
	XMLName     xml.Name `xml:"servers"`
	Version     string   `xml:"version,attr"`
	Svs         []server `xml:"server"`
	Description string   `xml:",innerxml"`
}

type server struct {
	XMLName    xml.Name `xml:"server"`
	ServerName string   `xml:"serverName"`
	ServerIP   string   `xml:"serverIP"`
}

// read xml file
func readXML() {
	file, err := os.Open("./template/file/servers.xml")
	if err != nil {
		fmt.Printf("11-error: %v\n", err)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("11-error: %v\n", err)
		return
	}
	v := Recurlyservers{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("11-error : %v\n", err)
	}
	fmt.Println(v)

}

type Servers struct {
	XMLName xml.Name  `xml:"servers"`
	Version string    `xml:"version,attr"`
	Svs     []server2 `xml:"server"`
}

type server2 struct {
	ServerName string `xml:"serverName"`
	ServerIP   string `xml:"serverIP"`
}

// generate xml file
func writeXML() {

	v := &Servers{Version: "1"}
	v.Svs = append(v.Svs, server2{"Shanghai", "127.0.0.1"})
	v.Svs = append(v.Svs, server2{"Beijing", "127.0.0.2"})

	output, err := xml.MarshalIndent(v, "  ", "    ")
	if err != nil {
		fmt.Printf("11-error: %v", err)
	}

	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(output)

}

func main() {
	readXML()
	writeXML()
}
