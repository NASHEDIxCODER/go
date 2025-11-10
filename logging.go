package main

import (
	"bytes"
	"fmt"
	"log"
	"log/slog"
	"os"
)

// Go provides standard library straight forward tools for outputting logs from go programs.

func main() {
// simply invoking functions like Println from the log package uses the standard logger, 
//which is already pre configure for reasonable logging output to os.Stderr .
//additional methods like Fetal* or Panic* will exit the program after logging.

	log.Println("standard logger")

	//Logger can be configured with flags to set their output format.
	//by defaut , standard logger has the log.Lfate and log.Ltime flag set,and these are collected in log.LstdFlags.
	//we can changes its flags to emit time with microsecond accurecy.

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with micro")

	//it also supports emitting the file name and line from which the log function is called.

	log.SetFlags(log.LstdFlags |log.Lshortfile)
	log.Println("whith file/line")

	// it may be useful to create custom logger and pass it around.
	//when creating a new logger, we can set a prefix to distinguish its output from other logger.
	mylog:= log.New(os.Stdout,"my:",log.LstdFlags)
	mylog.Println("from mylog")

	//it mayu be useful to create a custom loogger and pass it around. when creating a new logger,
	//we can ser a prefix to distinguish its output from other logger.
	mylog.SetPrefix("prefixx:")
	mylog.Println("from mylog")
	
	//logger can have custom output targets; any io.writterworks.

	var buf bytes.Buffer 
	buflog := log.New(&buf, "buf:",log.LstdFlags)
	
	//this call writtes log output into buf.
	buflog.Println("hello")

	//this will actually show it on standard output.
	fmt.Println("from buflog:",buf.String())

	// the slog package provides structures log output eg. JSON
	//slog output can contain an arbitrary number of key value pairs.

	jsonHandler := slog.NewJSONHandler(os.Stderr,nil)
	myslog:= slog.New(jsonHandler)

	myslog.Info("hi there")
	myslog.Info("hello again", "key", "val","age",25)


}