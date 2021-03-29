package goft

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

var ResponderList []Responder

func init() {
	ResponderList = []Responder{new(StringResponder)}
}

type Responder interface {
	RespondTo() gin.HandlerFunc
}

func Convert(handler interface{}) gin.HandlerFunc {
	h_ref := reflect.ValueOf(handler)
	for _, r := range ResponderList {
		r_ref := reflect.ValueOf(r).Elem()
		// 判断两个类型是否相等
		if h_ref.Type().ConvertibleTo(r_ref.Type()) {
			r_ref.Set(h_ref)
			// 调用接口
			return r_ref.Interface().(Responder).RespondTo()
		}
	}
	return nil
}

type StringResponder func(*gin.Context) string

func (this StringResponder) RespondTo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, this(ctx))
	}
}
