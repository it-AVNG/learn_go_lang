// a package is a way to group functions, made of files in the same dir
package main

// import a package
import "fmt"
import "rsc.io/quote"

// main function executes as default when we run the main package
func main() {
	fmt.Println("Hello World!")
  fmt.Println(quote.Go())
}
