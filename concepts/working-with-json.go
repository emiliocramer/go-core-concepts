package concepts

import (
	"encoding/json"
	"fmt"
	"os"
)

// There are many things that we can encode and decode in json using the default package
// Lets walk through some of those things

type responseType struct {
	Index  int
	Values []string
}

func JsonEncoding() {

	// Here are some ways we can encode to JSON

	// Lets encode basic data types to JSON
	stringEnc, _ := json.Marshal("template string to json")
	fmt.Println(string(stringEnc))

	intEnc, _ := json.Marshal(12)
	fmt.Println(string(intEnc))

	floatEnc, _ := json.Marshal(1.52)
	fmt.Println(string(floatEnc))

	slice := []string{"slices", "can", "be", "encoded", "the", "same"}
	sliceEnc, _ := json.Marshal(slice)
	fmt.Println(string(sliceEnc))

	mapForEnc := map[string]int{"firstKey": 3, "secondKey": 1}
	mapEnc, _ := json.Marshal(mapForEnc)
	fmt.Println(string(mapEnc))

	// We can also encode custom types
	// Field names have to be capitalized to be able to be seen
	responseOne := &responseType{
		Index:  1,
		Values: []string{"custom", "types", "work", "the", "same"},
	}
	typeEnc, _ := json.Marshal(responseOne)
	fmt.Println(string(typeEnc))
}

func JsonDecoding() {

	// Here are some ways we can decode from JSON

	// Here we're setting up a standard data set that we might be encoding
	// We also set up where we are going to put our encoded data (variable)
	jsonData := []byte(`{"num": 13.32, "vals": [12, 31]}`)
	var encoded map[string]interface{}

	// Once we've done that we have to decode the data & error check
	if err := json.Unmarshal(jsonData, &encoded); err != nil {
		panic(err)
	}
	fmt.Println(encoded)

	// We can then get the values from the encoded data
	num := encoded["num"].(float64)
	fmt.Println(num)

	vals := encoded["vals"].([]interface{})
	val1 := vals[0]
	fmt.Println(val1)

	// If the data that we are decoding fit the category of a custom type we can decode to a custom type
	str := `{"Index": 2, "Values": ["decode", "to", "custom", "types"]}`
	response := responseType{}
	if err := json.Unmarshal([]byte(str), &response); err != nil {
		panic(err)
	}
	fmt.Println(response)
	fmt.Println(response.Values[1])

	// Instead of dealing with bytes and Strings we can write to os.Writers, like os.Stdout
	encoder := json.NewEncoder(os.Stdout)
	m := map[string]int{"values": 3, "repeats": 1}
	encoder.Encode(m)
}
