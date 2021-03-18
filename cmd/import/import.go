package main

import (
	"fmt"
	"github.com/abusizhishen/ruleEngine/src"
)

var rule string
var data map[string]interface{}
var result interface{}

func main() {
	rule = `a*b-c`

	data = map[string]interface{}{
		"a":float64(6),
		"b":float64(500),
		"c":float64(3),
	}

	result = src.Run(rule,data)
	fmt.Println(result)
	//result=2997
}
