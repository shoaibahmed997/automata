package utils

import (
	"fmt"
	"os"
)

func WriteJsonToFile(filename string, data []byte) {
	err := os.WriteFile("macros/"+filename, data, 0644)
	// handle this error
	if err != nil {
		// print it out
		fmt.Println(err)
	}
}

func ReadJsonFromFile(filename string) []byte {
	jsonData, err := os.ReadFile("macros/" + filename)
	if err != nil {
		fmt.Println(err)
	}
	return jsonData
}

func AllMacros() []string {
	allMacros := []string{}
	dirs, err := os.ReadDir("macros")
	if err != nil {
		fmt.Println("Error", err)
	}
	for _, item := range dirs {
		allMacros = append(allMacros, item.Name())
	}
	return allMacros
}

func DeleteMacro() {

}
