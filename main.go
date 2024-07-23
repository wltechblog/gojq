package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <jsonfile> <path>")
		return
	}

	jsonFile := os.Args[1]
	path := os.Args[2]

	data, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	var jsonData interface{}
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	result, err := getValueAtPath(jsonData, path)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	switch result := result.(type) {
	case map[string]interface{}:
		for key := range result {
			fmt.Println(key)
		}
	case []interface{}:
		for key, _ := range result {
			fmt.Println(key)
		}
	case float64:
		fmt.Printf("%d\n", int64(result))
	default:
		fmt.Println(result)
	}
}

func getValueAtPath(data interface{}, path string) (interface{}, error) {
	segments := strings.Split(path, ".")
	var err error
	for _, segment := range segments {
		switch current := data.(type) {
		case map[string]interface{}:
			d, ok := current[segment]
			if !ok {
				return nil, fmt.Errorf("key not found: %s", segment)
			}
			data = d
		case []interface{}:
			index, err := strconv.Atoi(segment)
			if err != nil {
				return nil, fmt.Errorf("invalid array index: %s", segment)
			}
			if index < 0 || index >= len(current) {
				return nil, fmt.Errorf("index out of range: %s", segment)
			}
			data = current[index]
		default:
			return nil, fmt.Errorf("invalid path segment: %s", segment)
		}
	}
	return data, err
}
