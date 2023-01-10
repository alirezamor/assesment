package flights

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type flight struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

func GetStartAndEnd(c *gin.Context) {
	data, ok := c.GetQuery("path")
	var flights []flight
	var arr [][]string
	start := ""
	end := ""

	err := json.Unmarshal([]byte(data), &arr)

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing path queryparam."})
		return
	}

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "data is not valid"})
		return
	}

	for _, i := range arr {
		f := flight{
			Start: i[0],
			End:   i[1],
		}

		flights = append(flights, f)
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
