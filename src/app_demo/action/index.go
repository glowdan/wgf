package action

import (
	"wgf/sapi"
)

type IndexAction struct {
	sapi.Action
}

//方法主体
func (p *IndexAction) Execute() error {
	p.Sapi.Print("hello world\n")
	p.Sapi.Println(p)
	return nil
}

//将action注册进wgf
func init() {
	sapi.RegisterAction("index", func() sapi.ActionInterface { return &IndexAction{} })
}
