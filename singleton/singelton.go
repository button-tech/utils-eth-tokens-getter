package singleton

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)



var (
	EthEndpoints    []string
	ContractAddress string
	Client          *ethclient.Client
	GinServer       *gin.Engine
	GinPort         string
)
