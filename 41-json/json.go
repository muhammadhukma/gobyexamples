package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// go offers built-in support for json encoding and decoding, including to and form built-in and custom data types.

// we'll use these two structs to demonstrate encoding and decoding of coustom type below.
type response1 struct {
	Page   int
	Fruits []string
}

// only exported fields will be encoded/decoded in JSON.
// Fields must start with capital letter to be exported.
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	// first we'll look at encoding basic data to JSON strings. Here are some example for atomic values.
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(12.2)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	// and here are some for slices and maps, which encode to JSON arrays and objects as you'd expect.
	slcD := []string{"apple", "peach", "peer"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	/*
		The JSON package can automatically encode your custom data types.
		It will only include exported fields in the encode output and will by default use those names as the JSON keys.
	*/
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	/*
		You can use tags on struct field declarations to customize the encode JSON key names.
		Check the definition of response2 above to see an example of such tags.
	*/
	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	byt := []byte(`{"num": 6.13, "strs":["a", "b"]}`) // Now lets look at decoding JSON data Go values. Here an example for a generic data structur.

	// we need to provide a variable where the JSON package can put the decode data. this map[string]interface{} will hold a map of strings to arbitrary data types.
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	// in order to use the value in the dexoded map, we'll need to convert them to their approptiate type. For example here we convert the value in num to the expected float64 type.
	num := dat["num"].(float64)
	fmt.Println(num)

	// accessing nested data requires a series of conversions.
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	/*
		We can also decode JSON into custom data types.
		This has the advantages of adding additional type-safety to our programs and eliminating the need for type assertions when accessing the decode data.
	*/
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	/*
		In the examples above we alwas used bytes and string as intermediates between the data and JSON representation on standard out.
		We can also stream JSON encodings directly to os.Writers like os.Stdout or even HTTP response bodies.
	*/

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)

	dec := json.NewDocoder(strings.NewReader(str))
	res1 := response2{}
	dec.Decode(&res1)
	fmt.Println(res1)
}
