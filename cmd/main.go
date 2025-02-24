package main

import (
	"fmt"

	"github.com/osamikoyo/IM-review/internal/app"
)

func main(){
	app, err := app.Init()
	if err != nil{
		fmt.Println(err)
	}
	if err := app.Run();err != nil{
		fmt.Println(err)
	}	
}