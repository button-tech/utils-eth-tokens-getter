package singleton

import (
	"github.com/button-tech/utils-eth-tokens-getter/token-list"
	"github.com/gin-gonic/gin"
)

var (
	TokenList       []token_list.Tokens
	ContractAddress string
	GinServer       *gin.Engine
	GinPort         string
)
