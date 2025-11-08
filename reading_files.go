package main

//reading and writing files are basic tasks needed for many. Go program

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func check(e error){ //reading file requires checking most calls for error.
	if e !=nil{
		panic(e)
	}
}

func main(){
	dat, err := os.ReadFile("test/data.txt")
//the most basic file reaading task is sluring a file's entire contents into memory.
	check(err)
	fmt.Println(string(dat))
	//if we want more control over and what parts of the file are read for there tasks.
	//start by opening a file to obtain an os.file value.
	f, err:= os.Open("test/data1.txt")
	check(err)

	b1:= make([]byte, 5) // read 5 bytes from the begining.
	n1, err:=f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n",n1,string(b1[:n1]))
	
	o2,err:= f.Seek(6, io.SeekStart)
	//seek known location from a file and read from here.
	check(err)
	b2 :=make([]byte,2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n",string(b2[:n2]))

	//other methods of seeking are relative to current cursor position.
	_, err = f.Seek(2, io.SeekCurrent)
	check(err)
	_, err = f.Seek(-4, io.SeekEnd)
	check(err)
	o3 ,err := f.Seek(6, io.SeekStart)
	check(err)
	//the io package provide some functions that may be helpful for file reading.
	//for example , reads like the ones above can be more robustly implemented with ReadLeast.

	b3:= make([]byte, 3)
	n3, err:= io.ReadAtLeast(f,b3,3)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3,o3, string(b3))
//no bult-in in rewind but seek(0, io.seekStart)
	_, err = f.Seek(0, io.SeekStart)
	check(err)
//the bufio package implements a buffred reader that may be useful bith for its efficiencywith many small reads and becuse of the additional reading it provides.

	r4 := bufio.NewReader(f)

	b4, err:= r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n",string(b4))
	f.Close()




}