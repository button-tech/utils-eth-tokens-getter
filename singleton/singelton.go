package singleton

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

var (
	// add network type and port
	ContractAddress string
	EthEndpoint     string
	Client          *ethclient.Client
	GinServer       *gin.Engine
)
