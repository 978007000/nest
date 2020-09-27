package node

import (
	"errors"
	"github.com/gogf/gf/util/gvalid"
	"nest/app/model/label/node"
)

// 子级节点列表输入参数
type GetNodeListInput struct {
	GetNodeTreeInput
	Page  int `v:"required#页码不能为空"`
	Limit int `v:"required|max:10#分页数不能为空|最大分页数不能超过10"`
}

// 子级节点树输入参数
type GetNodeTreeInput struct {
	ParentId string `v:"required|in:1,2#父级编号不能为空|父级编号必须为1或2" json:"pid"`
}

// 根据父级编号获取下级有序节点列表
func GetNodeList(data *GetNodeListInput) ([]*node.SecondaryEntity, error) {
	// 输入参数检查
	if e := gvalid.CheckStruct(data, nil); e != nil {
		return nil, errors.New(e.FirstString())
	}
	entities, err := node.GetNodeList(data.ParentId, data.Page, data.Limit)
	return entities, err
}

// 根据父级编号获取下级有序节点树
func GetNodeTree(data *GetNodeTreeInput) ([]*node.SecondaryNode, error) {
	// 输入参数检查
	if e := gvalid.CheckStruct(data, nil); e != nil {
		return nil, errors.New(e.FirstString())
	}
	entities, err := node.GetNodeTree(data.ParentId)
	return entities, err
}
