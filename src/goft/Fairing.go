package goft

import "github.com/gin-gonic/gin"

// Fairing 整流罩， 用于保护卫星的(中间件的接口)
type Fairing interface {
	OnRequest(ctx *gin.Context) error
}
