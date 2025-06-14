package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	var data []byte
	var path string
	var err error

	// Check if stdin is a pipe
	stdinStat, _ := os.Stdin.Stat()
	isPipe := (stdinStat.Mode() & os.ModeCharDevice) == 0

	if isPipe {
		// Read from stdin
		data, err = io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println("Error reading from stdin:", err)
			os.Exit(-1)
		}

		// Path is the first argument if provided
		if len(os.Args) > 1 {
			path = os.Args[1]
		} else {
			path = ""
		}
	} else {
		// Original file-based functionality
		if len(os.Args) < 2 {
			fmt.Println("Usage: gojq <jsonfile> [path] or pipe JSON data and provide [path] as argument")
			return
		}

		jsonFile := os.Args[1]
		if len(os.Args) == 2 {
			path = ""
		} else {
			path = os.Args[2]
		}

		data, err = os.ReadFile(jsonFile)
		if err != nil {
			fmt.Println("Error reading JSON file:", err)
			os.Exit(-1)
		}
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
	if path == "" {
		return data, nil
	}
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
