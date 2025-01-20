package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: jq <jsonfile> <path>")
		return
	}

	jsonFile := os.Args[1]
	var path string
	if len(os.Args) == 2 {
		path = ""
	} else {
		path = os.Args[2]
	}
	data, err := os.ReadFile(jsonFile)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		os.Exit(-1)
	}

	var jsonData interface{}
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		os.Exit(-1)
		return
	}

	result, err := getValueAtPath(jsonData, path)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
		return
	}
	switch result := result.(type) {
	case map[string]interface{}:
		for key := range result {
			fmt.Println(key)
		}
	case []interface{}:
		for key, value := range result {
			fmt.Printf("%v:%s\n", key, value)
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
