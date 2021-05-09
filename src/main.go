package main

import (
	"flag"
	"fmt"

	"example.com/my/parser/constants"
	"example.com/my/parser/dataLoader"
	"example.com/my/parser/dataOutputer"
	"example.com/my/parser/dataParser"
	"example.com/my/parser/formatConverter"
	"example.com/my/parser/utils"
)

func main() {
	flagVals, err := getFlags(constants.FlagDict)
	utils.CheckError(err)

	// get stage instances
	loader, err := dataLoader.New(flagVals)
	utils.CheckError(err)

	parser, err := dataParser.New(flagVals)
	utils.CheckError(err)

	converter, err := formatConverter.New(flagVals)
	utils.CheckError(err)

	outputer, err := dataOutputer.New(flagVals)
	utils.CheckError(err)

	// exec
	rawString, err := loader.LoadDataString()
	utils.CheckError(err)

	recordArray, err := parser.ParseToMapArray(rawString)
	utils.CheckError(err)

	formatedStr, err := converter.Convert(recordArray)
	utils.CheckError(err)

	err = outputer.OutputData(formatedStr)
	utils.CheckError(err)

	fmt.Printf("Convert success")
}

func getFlags(flagDict map[string][3]string) (map[string]string, error) {
	flagVal := map[string]string{}
	flagPtr := map[string]*string{}

	for key, flagInfo := range flagDict {
		flagPtr[key] = flag.String(key, flagInfo[1], flagInfo[2])
	}

	flag.Parse()

	for key, ptr := range flagPtr {
		flagVal[key] = *ptr
	}

	return flagVal, nil
}
