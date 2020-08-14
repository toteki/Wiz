package wiz

import (
	"os"
)

//		*	*	*	*	*	*	*	*	*	*	*	*	*	*	*	*
//		*	*	*	*	*	*	*	*	*	*	*	*	*	*	*	*

//		Exposed functions:
//			Args() []string
//				Returns command line arguments passed to program
//			Executable() string
//				Returns full path to currently running executable
//			ProgramName() string
//				Returns just the name of the running executable (no path)
//			Dir() string
//				Returns the absolute path of the directory containing the executable

//		*	*	*	*	*	*	*	*	*	*	*	*	*	*	*	*
//		*	*	*	*	*	*	*	*	*	*	*	*	*	*	*	*

func Args() []string {
	return os.Args[1:] //Omit the first, which is program name
}
