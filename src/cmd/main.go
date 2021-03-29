package main

import (
	"github.com/shijting/goft/src/classes"
	"github.com/shijting/goft/src/goft"
	"github.com/shijting/goft/src/middleware"
)

const (
	Post = ":8080"
)
func main()  {
	goft.Ignite(Post).
		Attach(middleware.NewUserMid()).
		Mont("v1",classes.NewIndexClass()).
		Mont("v2", classes.NewUserClass()).
		Launch()
}
