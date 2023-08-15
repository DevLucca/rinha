package controller

import (
	"fmt"

	"github.com/DevLucca/rinha/application/dto"
	"github.com/DevLucca/rinha/application/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PersonController struct {
	svc *service.PersonService
}

func NewPersonController(svc *service.PersonService) *PersonController {
	return &PersonController{
		svc: svc,
	}
}

func (c *PersonController) Get(ctx *gin.Context) {

}

func (c *PersonController) Count(ctx *gin.Context) {
	count := c.svc.Count(ctx)
	ctx.JSON(200, count)
}

func (c *PersonController) Retrieve(ctx *gin.Context) {
	id := ctx.Param("id")
	parsedID, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(404, nil)
		return
	}

	person, err := c.svc.Retrieve(ctx, parsedID)
	if err != nil {
		ctx.JSON(404, nil)
		fmt.Println(err)
		return
	}

	dto := dto.PersonResponseDTO{
		ID:        person.ID,
		Name:      person.Name,
		Nickname:  person.Nickname,
		Birthdate: person.Birthdate,
		Stack:     person.Stack,
	}

	ctx.JSON(200, dto)
}

func (c *PersonController) Create(ctx *gin.Context) {
	var dto dto.PersonRequestDTO
	err := ctx.Bind(&dto)
	if err != nil {
		ctx.JSON(422, gin.H{"error": err})
	}

	id, err := c.svc.Create(ctx, dto)
	if err != nil {
		ctx.JSON(422, gin.H{"error": err})
		return
	}

	ctx.Header("Location", fmt.Sprintf("/pessoas/%s", id.String()))
	ctx.JSON(201, nil)
}
