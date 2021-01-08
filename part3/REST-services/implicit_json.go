package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var ks = []byte(`{
"firstName": "Jean",
"lastName": "Batik",
"age": 86,
"education": [
{
"institution": "Northwest Missouri State Teachers College",
"degree": "Bachelor of Science in Mathematics"
},
{
"institution": "University of Pennsylvania",
"degree": "Masters in English"
}
],
"spouse": "William Batik",
"children": [
"Timothy John Batik",
"Jane Helen Batik",
"Mary Ruth Batik"
]
}`)

func main() {
	var f interface{}
	err := json.Unmarshal(ks, &f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	printJSON(f)
}

func printJSON(v interface{}) {
	switch vv := v.(type) {
	case string:
		fmt.Println("is string", vv)
	case float64:
		fmt.Println("is float64", vv)
	case []interface{}:
		fmt.Println("is an array:")
		for i, u := range vv {
			fmt.Print(i, " ")
			printJSON(u)
		}
	case map[string]interface{}:
		fmt.Println("is an object:")
		for i, u := range vv {
			fmt.Print(i, " ")
			printJSON(u)
		}
	default:
		fmt.Println("Unknown type")
	}
}
