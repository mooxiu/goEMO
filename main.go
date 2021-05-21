package main

import "goemo/metadata"

func main() {
	err := metadata.InitConvertMap(metadata.MathBold)
	if err != nil {
		panic(err)
	}
	metadata.ShowConvertMap()
}
