package object

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

//Recurlyservers recurly servers
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

//ObjXML XML object operate samples
func ObjXML() {
	file, err := os.Open("servers.xml") // For read access.
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	v := Recurlyservers{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	//fmt.Printf("XML result:\n%v\n",string(data))
	for i := range v.Svs {
		s := v.Svs[i]
		fmt.Printf("Server result:%v\n", s.XMLName)
		fmt.Printf("Server Name:%s\n", s.ServerName)
		fmt.Printf("Server IP  :%s\n", s.ServerIP)
	}
	//fmt.Printf("XML result:\n%v\n",v)
}
