package flights

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type flight struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type airports struct {
	Paths string `json:"paths" binding:"required"`
}

func GetStartAndEnd(c *gin.Context) {
	var data airports
	var flights []flight
	var arr [][]string
	start := ""
	end := ""

	if err := c.ShouldBind(&data); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "you have to provide appropriate body."})
		fmt.Println(err)
		return
	}

	fmt.Println(data)

	// if data.paths == "" {
	// 	c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "body can not be empty."})
	// 	return
	// }

	if err := json.Unmarshal([]byte(data.Paths), &arr); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "data is not valid"})
		fmt.Println(err)
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
