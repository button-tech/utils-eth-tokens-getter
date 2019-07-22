package singleton

import (
	"github.com/gin-gonic/gin"
)

var (
	ContractAddress string
	GinServer       *gin.Engine
	GinPort         string
)
