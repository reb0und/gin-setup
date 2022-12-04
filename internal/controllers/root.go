package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SampleData struct {
	Foo string `json:"foo"`
	Bar string `json:"bar"`
}

func RootController(c *gin.Context) {
	c.JSON(http.StatusOK, SampleData{Foo: "f", Bar: "b"})
}