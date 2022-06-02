package test

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func ShowGinVersion() {
	fmt.Println(gin.Version)
}
