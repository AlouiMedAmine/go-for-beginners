//    ***** Go Comments ******

// ***** 1-Go Comments

// This is a comment
/*
package main

import "fmt"

func main() {

	// This is a comment1
	fmt.Println("Hello World!") // This is a comment2
}
*/

// **** 2-Go Multi-line Comments
/*
package main

import "fmt"

func main() {
	/* The code below will print Hello World
	to the screen, and it is amazing */

//fmt.Println("Hello World!")
//}

// ******* 3-Comment to Prevent Code Execution

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	//fmt.Println("This line does not execute")
}
