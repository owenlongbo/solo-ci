package models

type Build struct {
	Id      int   `orm:"pk;auto;unique" json:"id"` //主键
	Name    string `json:"name" json:"name"`
	Result  string `json:"result" json:"result"`
	Project *Project `orm:"rel(fk)" json:"project"`
}
