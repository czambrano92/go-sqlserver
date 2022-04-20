package main

import (
	"fmt"
	"gosqlserver/db"
)

type Result struct {
	ID   int
	Name string
	Age  int
}

func main() {

	f := db.Database()
	fmt.Println("ffff >> ", f)

}
