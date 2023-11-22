package controller

import (
	"api-with-sql-nativ/model"
	"api-with-sql-nativ/usecase"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type exampleController struct {
	rg          *gin.RouterGroup
	usecaseExmp usecase.UsecaseExample
}

func (c *exampleController) createExampleData(ctx *gin.Context) {
	var examplePayload model.ExampeModel

	if err := ctx.ShouldBind(&examplePayload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	example, err := c.usecaseExmp.CreateExampleData(examplePayload)
	fmt.Println("hallo")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, example)
}

func (c *exampleController) getDataExample(ctx *gin.Context) {
	example, err := c.usecaseExmp.GetDataExample()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusCreated, example)
}

func (c *exampleController) Router() {
	v2 := c.rg.Group("example")

	v2.POST("", c.createExampleData)
	v2.GET("", c.getDataExample)
}

func NewControllerExample(rg *gin.RouterGroup, usecaseExmp usecase.UsecaseExample) *exampleController {
	return &exampleController{rg: rg, usecaseExmp: usecaseExmp}
}
