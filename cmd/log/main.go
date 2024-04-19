package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("../codes.json")
	if err != nil {
		println("Could not open \"codes.json\".")
		os.Exit(1)
	}

	readJson, err := io.ReadAll(f)
	if err != nil {
		println("Could not read \"codes.json\"")
		os.Exit(1)
	}

	var logCodes LdLogCodesJson
	err = json.Unmarshal(readJson, &logCodes)
	if err != nil {
		println("Could not read unmarshal \"codes.json\"")
		os.Exit(1)
	}
	fmt.Printf("%+v", logCodes)

	// TODO: Add log class.
	// TODO: Add log system.
	// TODO: Add log code.
	os.Exit(0)
}
