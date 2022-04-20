package main

import (
	"fmt"
	"gosqlserver/db"
)

type Result struct {
	ID   int
	Name string
	Age  string
}

func main() {

	var result Result
	db.Database().Raw("SELECT TOP (1000) [ALC_ANO] as ID ,[ALC_SEM] as Name,[ALC_CLASE] as Age FROM [patentes].[dbo].[PC_ALCOHOL] ").Scan(&result)

	fmt.Println("ffff >> ", result)

}
