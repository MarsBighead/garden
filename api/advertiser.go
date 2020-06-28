package api

import (
	"garden"
	"io/ioutil"
	"net/http"
)

//Advertiser will show a advertiser via http json API form t HTTP Server
type Advertiser struct {
	Pattern     string
	Environment map[string]string
}

func (a *Advertiser) Category() string {
	return "API"
}

func (a *Advertiser) Description() string {
	return `a mock advertiser list`
}

func (a *Advertiser) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadFile(a.Environment["DATA"] + "/ad-mock.json")
	if err != nil {
		panic(err)
	}
	_, err = w.Write(body)
	if err != nil {
		panic(err)
	}
}

//AddEnv add environment variables to the object
func (a *Advertiser) AddEnv(env map[string]string) {
	a.Environment = env
}

func init() {
	garden.Add("advertiser", func() garden.Input {
		return &Advertiser{}
	})
}
