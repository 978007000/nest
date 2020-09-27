package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"nest/app/api/label"
)

func init() {
	s := g.Server()
	s.Group("/label", func(group *ghttp.RouterGroup) {
		group.GET("/node/list", label.GetNodeList)
		group.GET("/node/tree", label.GetNodeTree)
	})
}
