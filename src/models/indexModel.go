package models

import "fmt"

type IndexModel struct {
	Id   int
	Name string
}

func (this *IndexModel)String() string {
	return fmt.Sprintf("ID:%d Name: %s", this.Id, this.Name)
}
