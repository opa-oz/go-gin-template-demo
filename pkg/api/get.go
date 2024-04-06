package api

import (
	"server-template/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetResponse struct {
	Name  string `json:"name" swaggertype:"string" example:"test"`
	Index int32  `json:"index" swaggertype:"integer" example:"1234"`
}

// @BasePath /api

// Get godoc
// @Summary get something
// @Schemes
// @Param name path string true "Name"
// @Param index query int true "Index" default(0)
// @Description Get name and index
// @Tags api
// @Accept json
// @Produce json
// @Success 200 {object} api.GetResponse
// @Router /get/{name} [get]
func Get(c *gin.Context) {
	//ctx := c.Request.Context()
	name := c.Param("name")
	index := c.DefaultQuery("index", "0") // 0 - unlimited

	indexValue, err := strconv.ParseInt(index, 10, 64)
	if err != nil {
		utils.Unirror(c, err)
		return
	}

	utils.Unisponse(c, name, indexValue)
}
