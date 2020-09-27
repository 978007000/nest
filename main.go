package main

import (
	"github.com/gogf/gf/frame/g"
	_ "nest/boot"
	_ "nest/router"
)

func main() {
	g.Server().Run()
}
