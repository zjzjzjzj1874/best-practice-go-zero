package file

import _ "embed"

//go:embed test.json
var testJson string

var TestJson = testJson

//go:embed test_slice.json
var testSlice string

var TestSlice = testSlice
