// Shows how variables and packages can be initialized.
//
// Author: Dmitri Krasnenko

package corz

import "fmt"

//Always comes before init
var INIT = tierUp()

func init() {
	//Always comes before main
	fmt.Println("Initialize cor here!")
}

func tierUp() int {
	fmt.Println("Tier up!")
	return 5
}
