package main

import (
	"fmt"
	"kufa/Route"
	"path/filepath"
)

func main() {
	Route.Route()

	//glob, err := filepath.Glob("View/Dashboard/*.gohtml")
	//if err != nil {
	//	log.Fatalln(err.Error())
	//}
	//fmt.Println(glob)

	files := []string{}
	file := append(files, layoutFiles()...)

	fmt.Println(file)

}
func layoutFiles() []string {
	files, err := filepath.Glob("View/Dashboard/*.gohtml")
	if err != nil {
		panic(err)
	}
	return files
}
