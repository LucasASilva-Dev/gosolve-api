package version

import "fmt"

const MAJOR uint = 0
const MINOR uint = 0
const PATCH uint = 1

var COMMIT string = ""
var IDENTIFIER string = ""

func Version() string {
	var suffix string = ""
	if len(IDENTIFIER) > 0 {
		suffix = fmt.Sprintf("-%s", IDENTIFIER)
	}

	if len(COMMIT) > 0 {
		suffix = suffix + "+"
	}

	if len(COMMIT) > 0 {
		suffix = fmt.Sprintf("%s"+"commit.%s", suffix, COMMIT)
	}

	return fmt.Sprintf("%d.%d.%d%s", MAJOR, MINOR, PATCH, suffix)
}
