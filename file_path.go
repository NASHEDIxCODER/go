package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

//file path package provides function to parse and construct file path in way that is portable between operating systems:
//dir/file on linux vs. dir\file on windows,for example.


func main(){
	p:= filepath.Join("dirl","dir2","filename")
	fmt.Println("p:",p)
//join should be used to construct path in a portableway. it takes anu number of argument and constructs a hierarical path form them.

	fmt.Println(filepath.Join("dir1//","filename"))
	fmt.Println(filepath.Join("dir1/./dir1","filename"))
	//we should always use join instead of concatenating /s or \s manually. 
	// in addition to providing portability, join will also normalize path by removing superfluous sepraters and directory changes

	fmt.Println("Dir(p):",filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))
	//dir and base can be used to split a path to the directory and the file.
	//alternatively,split will return both in same call.
	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))
	//we can check weather path is absolute.
	filename:= "sonfig.json"

	ext:= filepath.Ext(filename)
	fmt.Println(ext)//some file names have extension following a dot we can slpit the extension out of such names with Ext.
	fmt.Println(strings.TrimSuffix(filename,ext)) //to find the file name with the extension removed, use string.TrimSuffix.

	rel, err:= filepath.Rel("a/b","a/b/t/file")
	if err != nil{
		panic(err)
	}
	fmt.Println(rel)
//rel find a path between a base and a target.
//it returns an error if the target cannot be made relative to base.
	rel1,err:= filepath.Rel("a/b","a/c/t/file")
	if err!= nil{
		panic(err)
	}
	fmt.Println(rel1)

	
}