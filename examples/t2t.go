package main

import (
	"fmt"
	"github.com/gohouse/converter"
)

func main() {
	t2t := converter.NewTable2Struct().Config(&converter.T2tConfig{
		StructNameToHump: true,
	})

	err := t2t.
		//SavePath("/home/go/project/model/model.go").
		Dsn("root:248655@tcp(localhost:3306)/janson?charset=utf8").
		RealNameMethod("TableName").
		Run()
	fmt.Println(err)
}
