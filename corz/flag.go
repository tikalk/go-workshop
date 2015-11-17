// Shows how to read command line arguments
//
// Author: Dmitri Krasnenko

package corz

import (
	"fmt"
	flg "flag"
)

func Flag() {
	var name = flg.String("u","Anonymous","User name")

	flg.Parse()

	fmt.Printf("Hello %s!\n", *name)

}
