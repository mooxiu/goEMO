package main

import (
	"fmt"
	"goemo/app"
	"goemo/metadata"
)

var testText = `package beatles

var John = "John"

func foo() {
	Paul := "Paul"
	println(John, Paul)
}
`

func main() {
	err := metadata.InitConvertMap(metadata.MathBold)
	if err != nil {
		panic(err)
	}
	// metadata.ShowConvertMap()
	replaced := app.ConvertEmo(testText, "beatles", 8, metadata.MyConvertMap)
	fmt.Println(replaced)
}
