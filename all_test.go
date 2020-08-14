package wiz

import (
	"testing"
)

func TestAll(t *testing.T) {
	Purple("Testing . . .")
	Green("Args", Args())
	Green("Executable", Executable())
	Green("ProgramName", ProgramName())
	defer Purple(". . . Tested")
}
