/*
How do we run the code in our project:

Go CLI: 
go build: Compiles a bunch of go source code files
go run: Compiles and excute one or two files
go fmt: Formats all the code in each file in the current directory 
go install: Compules and "installs" a package.
go get: Downloads the raw source code of someone else's package
go test: Runs any tests associated with the current project
*/
package main //Package = Project = Workspace  - every file belongs to this package need this 
//Type of package:1. Excuteable: Generates a file that we can run; 
//                2. Reusable: Code used as "help[ers. good place to put resuable logic
//
import "fmt" //Give our package access to some code insode other 
//Standard library: main  -> fmt/math/encoding/crypto/io/debug

func main() { //function main 
	fmt.Println("Hi There!")
}
