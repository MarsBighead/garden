package api

import (
	"garden"
)

type Advertiser struct {
	Directory []string
}

func (a *Advertiser) GetOperationType() string {
	return "API"
}

func (a *Advertiser) Description() string {
	return `a mock advertiser list`
}

func (a *Advertiser) Router() string {
	return `/api/json/advertiver`
}

func (a *Advertiser) Transform() {

}
func init() {
	Add("advertiser.go", func() garden.Input {
		return &Advertiser{}
	})
}
