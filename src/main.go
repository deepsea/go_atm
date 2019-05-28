package main

import (
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	fmt.Println("Atms rule the world!")

	a := App{}

	a.Initialize("mysql", "rlawlor:password@tcp(localhost:3306)/atmdemo?charset=utf8&parseTime=True&loc=Local")

	a.Run(":9996")

}
