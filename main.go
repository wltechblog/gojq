package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Take arguments from command line
	// if node is omitted, we are looking for the root node
	// if node is provided, we are looking for the node

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s filename [node]", os.Args[0])
		os.Exit(1)
	}

	// set filename and attribute node
	filename := os.Args[1]
	node := ""
	if len(os.Args) < 3 {
		node = ""
	} else {
		node = os.Args[2]
	}
	// Read json file and unmarshal
	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	var data map[string]interface{}
	err = json.Unmarshal(file, &data)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		os.Exit(1)
	}
	// Find the requested node
	value, ok := findNode(data, node)
	if !ok {
		os.Exit(1)
	}

	// Print the value of the node
	fmt.Println(value)
}

func findNode(data map[string]interface{}, node string) (interface{}, bool) {
	// Split the node into parts
	parts := strings.Split(node, ".")
	// Find the value of the node
	value, ok := findNodeValue(data, parts)
	return value, ok
}

func findNodeValue(data interface{}, parts []string) (interface{}, bool) {
	// Check if the data is a map
	if dataMap, ok := data.(map[string]interface{}); ok {
		// if node is empty, return the root node
		if len(parts) == 1 && parts[0] == "" {
			// We only want the keys at this node
			var keys []string
			for key := range dataMap {
				keys = append(keys, key)
			}

			return strings.Join(keys, "\n"), true
		}

		// Check if the node is in the map
		if value, ok := dataMap[parts[0]]; ok {
			// if value is a string, reutrn the value
			if _, ok := value.(string); ok {
				return value, true
			}
			// Check if the node is the last part
			if len(parts) == 1 {
				var keys []string
				for key := range value.(map[string]interface{}) {
					keys = append(keys, key)
				}

				return strings.Join(keys, "\n"), true

			}
			// Find the value of the next part
			return findNodeValue(value, parts[1:])
		}
	}
	// Check if the data is a slice
	if dataSlice, ok := data.([]interface{}); ok {
		// Check if the node is an index
		if index, ok := parseIndex(parts[0], len(dataSlice)); ok {
			// Check if the index is in the slice
			if index < len(dataSlice) {
				// Check if the node is the last part
				if len(parts) == 1 {
					var keys []string
					for key := range dataSlice[index].(map[string]interface{}) {
						keys = append(keys, key)
					}

					return strings.Join(keys, "\n"), true
				}
			}
		}
	}
	return nil, false
}

func parseIndex(part string, length int) (int, bool) {
	// Check if the part is an index
	if index, err := strconv.Atoi(part); err == nil {
		// Check if the index is in the range
		if index >= 0 && index < length {
			return index, true
		}
	}
	return 0, false
}
