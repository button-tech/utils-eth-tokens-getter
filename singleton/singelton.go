package singleton

import (
	"github.com/gin-gonic/gin"
	"github.com/button-tech/utils-eth-tokens-getter/token-list"
)

var (
	TokenList       []token_list.Tokens
	ContractAddress string
	GinServer       *gin.Engine
	GinPort         string
)
