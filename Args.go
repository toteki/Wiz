package wiz

import (
	"os"
)

// Returns the command line arguments passed to the program. Omits the first element of os.Args, which is the program's filename.
func Args() []string {
	return os.Args[1:] //Omit the first, which is program name
}
