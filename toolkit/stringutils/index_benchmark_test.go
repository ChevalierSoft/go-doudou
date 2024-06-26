// SPDX-License-Identifier: MIT OR Unlicense

package stringutils

import (
	"regexp"
	"testing"
)

var testMatchEndCase = `1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890 1test`
var testUnicodeMatchEndCase = `ȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺ Ⱥtest`

var testMatchEndCaseLarge = `1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890 1test`
var testUnicodeMatchEndCaseLarge = `ȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺȺ Ⱥtest`

func BenchmarkFindAllIndex(b *testing.B) {
	r := regexp.MustCompile(`test`)
	haystack := []byte(testMatchEndCase)

	for i := 0; i < b.N; i++ {
		matches := r.FindAllIndex(haystack, -1)
		if len(matches) != 1 {
			b.Error("Expected single match")
		}
	}
}

func BenchmarkIndexAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		matches := IndexAll(testMatchEndCase, "test", -1)
		if len(matches) != 1 {
			b.Error("Expected single match")
		}
	}
}

func BenchmarkFindAllIndexLarge(b *testing.B) {
	r := regexp.MustCompile(`test`)
	haystack := []byte(testMatchEndCaseLarge)

	for i := 0; i < b.N; i++ {
		matches := r.FindAllIndex(haystack, -1)
		if len(matches) != 1 {
			b.Error("Expected single match")
		}
	}
}

func BenchmarkIndexAllLarge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		matches := IndexAll(testMatchEndCaseLarge, "test", -1)
		if len(matches) != 1 {
			b.Error("Expected single match")
		}
	}
}

func BenchmarkFindAllIndexUnicode(b *testing.B) {
	r := regexp.MustCompile(`test`)
	haystack := []byte(testUnicodeMatchEndCase)

	for i := 0; i < b.N; i++ {
		matches := r.FindAllIndex(haystack, -1)
		if len(matches) != 1 {
			b.Error("Expected single match")
		}
	}
}

func BenchmarkIndexAllUnicode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		matches := IndexAll(testUnicodeMatchEndCase, "test", -1)
		if len(matches) != 1 {
			b.Error("Expected single match")
		}
	}
}

func BenchmarkFindAllIndexUnicodeLarge(b *testing.B) {
	r := regexp.MustCompile(`test`)
	haystack := []byte(testUnicodeMatchEndCaseLarge)

	for i := 0; i < b.N; i++ {
		matches := r.FindAllIndex(haystack, -1)
		if len(matches) != 1 {
			b.Error("Expected single match")
		}
	}
}

func BenchmarkIndexAllUnicodeLarge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		matches := IndexAll(testUnicodeMatchEndCaseLarge, "test", -1)
		if len(matches) != 1 {
			b.Error("Expected single match")
		}
	}
}

// This benchmark simulates a bad case of there being many
// partial matches where the first character in the needle
// can be found throughout the haystack
func BenchmarkFindAllIndexManyPartialMatches(b *testing.B) {
	r := regexp.MustCompile(`1test`)
	haystack := []byte(testMatchEndCase)

	for i := 0; i < b.N; i++ {
		matches := r.FindAllIndex(haystack, -1)
		if len(matches) != 1 {
			b.Error("Expected single match")
		}
	}
}

// This benchmark simulates a bad case of there being many
// partial matches where the first character in the needle
// can be found throughout the haystack
func BenchmarkIndexAllManyPartialMatches(b *testing.B) {
	for i := 0; i < b.N; i++ {
		matches := IndexAll(testMatchEndCase, "1test", -1)
		if len(matches) != 1 {
			b.Error("Expected single match")
		}
	}
}

// This benchmark simulates a bad case of there being many
// partial matches where the first character in the needle
// can be found throughout the haystack
func BenchmarkFindAllIndexUnicodeManyPartialMatches(b *testing.B) {
	r := regexp.MustCompile(`Ⱥtest`)
	haystack := []byte(testUnicodeMatchEndCase)

	for i := 0; i < b.N; i++ {
		matches := r.FindAllIndex(haystack, -1)
		if len(matches) != 1 {
			b.Error("Expected single match")
		}
	}
}

// This benchmark simulates a bad case of there being many
// partial matches where the first character in the needle
// can be found throughout the haystack
func BenchmarkIndexAllUnicodeManyPartialMatches(b *testing.B) {
	for i := 0; i < b.N; i++ {
		matches := IndexAll(testUnicodeMatchEndCase, "Ⱥtest", -1)
		if len(matches) != 1 {
			b.Error("Expected single match")
		}
	}
}

func BenchmarkFindAllIndexUnicodeManyPartialMatchesVeryLarge(b *testing.B) {
	r := regexp.MustCompile(`Ⱥtest`)

	var large string
	for i := 0; i <= 100; i++ {
		large += testUnicodeMatchEndCaseLarge
	}

	haystack := []byte(large)

	for i := 0; i < b.N; i++ {
		matches := r.FindAllIndex(haystack, -1)
		if len(matches) != 101 {
			b.Error("Expected 101 match got", len(matches))
		}
	}
}

func BenchmarkIndexAllUnicodeManyPartialMatchesVeryLarge(b *testing.B) {
	var large string
	for i := 0; i <= 100; i++ {
		large += testUnicodeMatchEndCaseLarge
	}

	for i := 0; i < b.N; i++ {
		matches := IndexAll(large, "Ⱥtest", -1)
		if len(matches) != 101 {
			b.Error("Expected 101 match got", len(matches))
		}
	}
}

func BenchmarkFindAllIndexUnicodeManyPartialMatchesSuperLarge(b *testing.B) {
	r := regexp.MustCompile(`Ⱥtest`)

	var large string
	for i := 0; i <= 500; i++ {
		large += testUnicodeMatchEndCaseLarge
	}

	haystack := []byte(large)

	for i := 0; i < b.N; i++ {
		matches := r.FindAllIndex(haystack, -1)
		if len(matches) != 501 {
			b.Error("Expected 501 match got", len(matches))
		}
	}
}

func BenchmarkIndexAllUnicodeManyPartialMatchesSuperLarge(b *testing.B) {
	var large string
	for i := 0; i <= 500; i++ {
		large += testUnicodeMatchEndCaseLarge
	}

	for i := 0; i < b.N; i++ {
		matches := IndexAll(large, "Ⱥtest", -1)
		if len(matches) != 501 {
			b.Error("Expected 501 match got", len(matches))
		}
	}
}
