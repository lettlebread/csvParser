package main

import (
	"flag"
	"fmt"

	"example.com/my/dataLoader"
	"example.com/my/dataOutputer"
	"example.com/my/dataParser"
	"example.com/my/formatConverter"
	"example.com/my/utils"
)

var flagDict = map[string][]string{
	"f": []string{"string", "", "path of source file"},
	"m": []string{"string", "", "format of source data"},
	"o": []string{"string", "", "path of output data"},
	"e": []string{"string", "", "format of output data"},
}

func main() {
	flagVals := getFlags()

	// get stage instances
	loader, err := dataLoader.NewDataLoader(flagVals)
	utils.CheckError(err)

	parser, err := dataParser.NewDataParser(flagVals)
	utils.CheckError(err)

	converter, err := formatConverter.NewFormatConverter(flagVals)
	utils.CheckError(err)

	outputer, err := dataOutputer.NewDataOutputer(flagVals)
	utils.CheckError(err)

	// exec
	rawString, err := loader.LoadDataString()
	utils.CheckError(err)

	mapArray, err := parser.ParseToMapArray(rawString)
	utils.CheckError(err)

	formatedStr, err := converter.Convert(mapArray)
	utils.CheckError(err)

	err = outputer.OutputData(formatedStr)
	utils.CheckError(err)

	fmt.Printf("Convert success")
}

func getFlags() map[string]string {
	flagVal := map[string]string{}
	flagPtr := map[string]*string{}

	for key, flagInfo := range flagDict {
		flagPtr[key] = flag.String(key, flagInfo[1], flagInfo[2])
	}

	flag.Parse()

	for key, ptr := range flagPtr {
		flagVal[key] = *ptr
	}

	return flagVal
}
