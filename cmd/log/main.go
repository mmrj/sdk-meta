package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func addSystem(codes *LdLogCodesJson, name string, description string) error {
	_, present := (*codes)[name]
	if present {
		return fmt.Errorf("system name already exists. Please choose a new name or using the existing specifier")
	}

	maxSpecifier := -1.0
	for _, system := range *codes {
		if system.Specifier > maxSpecifier {
			maxSpecifier = system.Specifier
		}
	}
	newSpecifier := maxSpecifier + 1
	(*codes)[name] = System{
		Description: description,
		Specifier:   newSpecifier,
	}
	return nil
}

func addClass(codes *LdLogCodesJson, name string, description string) error {
	classNames := map[string]bool{}
	maxSpecifier := -1.0
	for _, system := range *codes {
		for className, class := range system.Classes {
			classNames[className] = true
			if class.Specifier > maxSpecifier {
				maxSpecifier = class.Specifier
			}
		}
	}
	_, present := classNames[name]

	if present {
		return fmt.Errorf("class name already exists. Please choose a new name or using the existing specifier")
	}

	return nil
}

func addCode(codes *LdLogCodesJson, name string, description string) {

}

func main() {
	f, err := os.Open("logs/data/codes.json")
	if err != nil {
		println(fmt.Errorf("Could not open \"codes.json\": %w", err).Error())
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

	// TODO: Add log class.
	// TODO: Add log system.
	// TODO: Add log code.
	os.Exit(0)
}
