package label

import (
	"github.com/gogf/gf/net/ghttp"
	"nest/app/service/label/node"
	response "nest/library"
)

// 下级后台标签列表请求参数，用于前后端交互参数格式约定
type NodeListRequest struct {
	node.GetNodeListInput
}
// 下级后台标签树请求参数，用于前后端交互参数格式约定
type NodeTreeRequest struct {
	node.GetNodeTreeInput
}

// GetNodeList 获取节点列表
func GetNodeList(r *ghttp.Request) {
	var data *NodeListRequest
	if err := r.GetStruct(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	entities, err := node.GetNodeList(&data.GetNodeListInput)
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "ok", entities)
}

// GetNodeTree 获取节点树
func GetNodeTree(r *ghttp.Request) {
	var data *NodeTreeRequest
	if err := r.GetStruct(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	entities, err := node.GetNodeTree(&data.GetNodeTreeInput)
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "ok", entities)
}
