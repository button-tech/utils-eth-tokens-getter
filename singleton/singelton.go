package singleton

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

var ContractAddress string
var Client *ethclient.Client
var GinServer *gin.Engine
