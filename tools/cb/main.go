package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"

	couchbase "github.com/couchbase/go-couchbase"
	_ "github.com/go-sql-driver/mysql"
)

var regexpURI = regexp.MustCompile(`(\S+://)?(\S+\:\S+@)`)

func main() {
	dir, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadFile(dir + "/config.json")
	if err != nil {
		log.Fatal(err)
	}
	conf := new(Configure)
	err = json.Unmarshal(body, conf)
	if err != nil {
		log.Fatal(err)
	}
	if conf.Port == "" {
		conf.Port = "8091"
	}
	addr := fmt.Sprintf("http://%s:%s@%s:%s", conf.Username, conf.Password, conf.Host, conf.Port)
	client, err := couchbase.Connect(addr)
	if err != nil {
		log.Fatal(err)
	}
	pool, err := client.GetPool("default")
	if err != nil {
		log.Fatal(err)
	}

	obj, _ := json.MarshalIndent(pool, "", "    ")
	fmt.Printf("pool %s\n%s\n", pool.BucketURL, string(obj))
	for i := 0; i < len(pool.Nodes); i++ {
		node := pool.Nodes[i]
		//tags := map[string]string{"cluster": regexpURI.ReplaceAllString(addr, "${1}"), "hostname": node.Hostname}
		fields := make(map[string]interface{})
		fields["memory_free"] = node.MemoryFree
		fields["memory_total"] = node.MemoryTotal
		//fmt.Printf("%#v\n", tags)
		//body, _ := json.MarshalIndent(node, "", "    ")
		//fmt.Printf("node %s\n%s\n", node.Hostname)
		//fmt.Printf("%#v\n", node)
	}
	for bucketName := range pool.BucketMap {
		tags := map[string]string{"cluster": regexpURI.ReplaceAllString(addr, "${1}"), "bucket": bucketName}
		fmt.Printf("%#v\n", tags)

		bs := pool.BucketMap[bucketName].BasicStats
		fields := make(map[string]interface{})
		fields["quota_percent_used"] = bs["quotaPercentUsed"]
		fields["ops_per_sec"] = bs["opsPerSec"]
		fields["disk_fetches"] = bs["diskFetches"]
		fields["item_count"] = bs["itemCount"]
		fields["disk_used"] = bs["diskUsed"]
		fields["data_used"] = bs["dataUsed"]
		fields["mem_used"] = bs["memUsed"]
		//fmt.Printf("%#v\n", pool.BucketMap[bucketName].NodesJSON)
		//body, _ := json.MarshalIndent(pool.BucketMap[bucketName], "", "    ")
		//fmt.Printf("%s\n", string(body))
	}
}

type Configure struct {
	Driver   string `yaml:"driver"    json:"driver"`
	Host     string `yaml:"host"      json:"host"`
	Username string `yaml:"username"  json:"username"`
	Password string `yaml:"password"  json:"password"`
	Bucket   string `yaml:"bucket"    json:"bucket"`
	Switch   string `yaml:"switch"    json:"switch"`
	Port     string `yaml:"-"         json:"-"`
}
