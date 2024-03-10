package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

type InputJSON map[string]interface{}
type OutputJSON []map[string]interface{}

func main() {
	inputJSON := InputJSON{}
	err := json.NewDecoder(os.Stdin).Decode(&inputJSON)
	if err != nil {
		log.Fatalf("Error decoding input JSON: %v", err)
	}

	outputJSON := transformInputJSON(inputJSON)

	outputBytes, err := json.Marshal(outputJSON)
	if err != nil {
		log.Fatalf("Error encoding output JSON: %v", err)
	}

	fmt.Println(string(outputBytes))
}

func transformInputJSON(input InputJSON) OutputJSON {
	output := make(OutputJSON, 0)

	for key, value := range input {
		key = strings.TrimSpace(key)
		if key == "" {
			continue
		}

		transformedValue := transformValue(value)
		if transformedValue != nil {
			output = append(output, map[string]interface{}{key: transformedValue})
		}
	}

	sort.Slice(output, func(i, j int) bool {
		return strings.Compare(getMapKey(output[i]), getMapKey(output[j])) < 0
	})

	return output
}

func transformValue(value interface{}) interface{} {
	switch v := value.(type) {
	case map[string]interface{}:
		return transformMap(v)
	case string:
		return transformString(v)
	case float64:
		return transformFloat(v)
	case bool:
		return transformBool(v)
	default:
		return nil
	}
}

func transformMap(input map[string]interface{}) map[string]interface{} {
	output := make(map[string]interface{})

	for key, value := range input {
		key = strings.TrimSpace(key)
		if key == "" {
			continue
		}

		transformedValue := transformValue(value)
		if transformedValue != nil {
			output[key] = transformedValue
		}
	}

	return output
}

func transformString(input string) interface{} {
	input = strings.TrimSpace(input)
	if input == "" {
		return nil
	}

	// RFC3339 time format
	t, err := time.Parse(time.RFC3339, input)
	if err == nil {
		return t.Unix()
	}

	return input
}

func transformFloat(input float64) interface{} {
	// Remove leading zeros
	return int(input)
}

func transformBool(input bool) interface{} {
	return input
}

func getMapKey(m map[string]interface{}) string {
	for k := range m {
		return k
	}
	return ""
}
