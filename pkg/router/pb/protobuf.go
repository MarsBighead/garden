package pb

import (
	"garden/config"
	"garden/marble/pbt"
	"garden/pkg/router"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
)

type Pb struct {
	Environment *config.Environment
	message     proto.Message
}

// Pbt Test binary protobuf data protocol
func (p *Pb) Get(c *gin.Context) {
	_, err := c.Writer.WriteString(proto.MarshalTextString(p.message))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
	}
}

// RebuildPbt Test Marshal/Unmarshal protobuf
func (p *Pb) Post(c *gin.Context) {
	data := make([]byte, c.Request.ContentLength)
	n, err := c.Request.Body.Read(data)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
	}
	body := new(pbt.Test)
	err = proto.Unmarshal(data, body)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
	}
	c.IndentedJSON(http.StatusOK, gin.H{"format": "protobuf", "length": n})
}

func genMessage() (msg proto.Message) {
	msg = &pbt.Test{
		Label: proto.String("hello"),
		Type:  proto.Int32(18),
		Reps:  []int64{1, 2, 3},
		Optionalgroup: &pbt.Test_OptionalGroup{
			RequiredField: proto.String("good bye"),
		},
	}
	return
}

func init() {
	router.Add("/pb", func(env *config.Environment) router.Act {
		return &Pb{
			message:     genMessage(),
			Environment: env,
		}
	})
}
