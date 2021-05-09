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
	loaderType, loaderArgs := getDataLoaderArgs(flagVals)
	parserType, parserArgs := getDataParserArgs(flagVals)
	converterType, converterArgs := getFormatConverterArgs(flagVals)
	outputerType, outputerArgs := getDataOutputerArgs(flagVals)

	loader, err := dataLoader.NewDataLoader(loaderType, loaderArgs)
	utils.CheckError(err)

	parser, err := dataParser.NewDataParser(parserType, parserArgs)
	utils.CheckError(err)

	converter, err := formatConverter.NewFormatConverter(converterType, converterArgs)
	utils.CheckError(err)

	outputer, err := dataOutputer.NewDataOutputer(outputerType, outputerArgs)
	utils.CheckError(err)

	rawString, err := loader.Exec()
	utils.CheckError(err)

	mapArray, err := parser.ParseToMapArray(rawString)
	utils.CheckError(err)

	formatedStr, err := converter.Exec(mapArray)
	utils.CheckError(err)

	err = outputer.Exec(formatedStr)
	utils.CheckError(err)

	fmt.Printf("Convert success")
}

func getDataLoaderArgs(args map[string]string) (string, []string) {
	loaderType := ""
	loaderArgs := []string{}

	if val, ok := args["f"]; ok {
		loaderType = "file"
		loaderArgs = append(loaderArgs, val)
	}

	return loaderType, loaderArgs
}

func getDataParserArgs(args map[string]string) (string, []string) {
	parserType := ""
	parserArgs := []string{}

	if val, ok := args["m"]; ok {
		parserType = val
	}

	return parserType, parserArgs
}

func getFormatConverterArgs(args map[string]string) (string, []string) {
	converterType := ""
	converterArgs := []string{}

	if val, ok := args["e"]; ok {
		converterType = val
	}

	return converterType, converterArgs
}

func getDataOutputerArgs(args map[string]string) (string, []string) {
	outputerType := ""
	outputerArgs := []string{}

	if val, ok := args["o"]; ok {
		outputerType = "file"
		outputerArgs = append(outputerArgs, val)
	}

	return outputerType, outputerArgs
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
