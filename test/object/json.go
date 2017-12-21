package object

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string
	ServerIP   string
	Location   interface{}
}

type Loc1 struct {
	City string
	Zip  string
}
type Loc2 struct {
	City     string
	Province string
	Zip      string
}
type S struct {
	Servers []Server
}

// ObjJson  Test interface node which will bring different object in
func ObjJson() {
	var s S
	s.Servers = append(s.Servers, Server{ServerName: "Shanghai_VPN", ServerIP: "127.0.0.1"})
	loc1 := Loc1{
		City: "Shanghai",
		Zip:  "200001",
	}
	s.Servers[0].Location = loc1
	s.Servers = append(s.Servers, Server{ServerName: "Beijing_VPN", ServerIP: "127.0.0.2"})
	loc2 := &Loc2{
		Province: "Hebei",
		City:     "Xingtai",
		Zip:      "005400",
	}
	s.Servers[1].Location = loc2
	// General json Marshal
	// b, err := json.Marshal(s)
	// Beautiful json Marshal
	b, err := json.MarshalIndent(s, "", "    ")
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
}
