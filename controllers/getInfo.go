package controllers

import (
	"cooding/graphql"

	"github.com/gin-gonic/gin"
)

func GetInfo(ctx *gin.Context) {

	a, _ := graphql.Query(ctx, "{me{name}}")

	ctx.JSON(200, gin.H{
		"message": "Hello World",
		"data":    a,
	})

}
