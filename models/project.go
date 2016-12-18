package models

type Project struct {
	Name string //name
	ID string   //uuid
	Type string //github,gitlab
	Url string //仓库地址
}

//func (obj *Project) WebHook() {
//
//}
