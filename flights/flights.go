package flights

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type flight struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

func GetStartAndEnd(c *gin.Context) {
	var flights []flight

	if err := c.BindJSON(&flights); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Body is not valid"})
		return
	}
	starts := make(map[string]bool)
	ends := make(map[string]bool)

	for _, i := range flights {
		starts[i.Start] = false
		ends[i.End] = false
	}

	for _, i := range flights {
		if _, ok := starts[i.End]; ok {
			starts[i.End] = true
		}
		if _, ok := ends[i.Start]; ok {
			ends[i.Start] = true
		}
	}

	start := ""
	end := ""

	for i := range starts {
		if !starts[i] {
			start = i
		}
	}

	for i := range ends {
		if !ends[i] {
			end = i
		}
	}
	res := flight{
		Start: start,
		End:   end,
	}

	c.IndentedJSON(http.StatusOK, res)
}
