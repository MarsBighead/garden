package adv

import (
	"fmt"
	"garden/config"
	"garden/pkg/router"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Advertiser will show a advertiser via http json API form t HTTP Server
type Advertiser struct {
	Environment *config.Environment
}

func (r *Advertiser) Get(c *gin.Context) {
	fmt.Printf("%v\n", r.Environment.Directory)
	body, err := ioutil.ReadFile(r.Environment.Directory.Data + "/ad-mock.json")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
	}
	_, err = c.Writer.Write(body)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err})
	}

}

func (r *Advertiser) Post(c *gin.Context) {
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "post method not support for API advertiser"})
}

func init() {
	router.Add("/advertiser", func(env *config.Environment) router.Act {
		return &Advertiser{
			Environment: env,
		}
	})
}
