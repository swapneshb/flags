package common

import (
	"flag"
	"fmt"
)

type CliOption struct {
	CliString     string // The string accepted by the commandline for this option
	ExpectedValue string // for e.g. (true/false) for bool, 'hostname' for -name
	UsageMessage  string // The message to show the user to explaing usage for this option
	URLPathString string // The corresponding string in the URL
	DefaultValue  string
	FilterOption  bool   // is this an option a filter parameter
	ID            bool   // does this option provide an ID field (like uuid/name/id)
	ParentPath    string // any parent option for this option (example server/2/vm/ ...)
	Value         string
	Mandatory     bool // is this a mandatory option (relevant for suboptions primarily)
}

var defaultOptions []CliOption = []CliOption{
	{CliString: "h", UsageMessage: "Prints out usage", ExpectedValue: "",  FilterOption: false, ID: false, ParentPath: "", Mandatory: false},
	{CliString: "help", UsageMessage: "Prints out usage", ExpectedValue: "",  FilterOption: false, ID: false, ParentPath: "", Mandatory: false},
	{CliString: "host", UsageMessage: "The name/IP address of the hellfire VM.", ExpectedValue: "<host/IP Address>",  FilterOption: false, ID: false, ParentPath: "", Mandatory: false},
{CliString: "port", UsageMessage: "The port number where the HF VM is running, default is 8080", DefaultValue: "8080", ExpectedValue: "<port>", URLPathString: "", FilterOption: false, ID: false, ParentPath: "", Mandatory: false},
{CliString: "user", UsageMessage: "The username to be used", DefaultValue: " ", ExpectedValue: "<username>", URLPathString: "", FilterOption: false, ID: false, ParentPath: "", Mandatory: false},
	{CliString: "password", UsageMessage: "The password for the user", DefaultValue: " ", ExpectedValue: "<password>", URLPathString: "", FilterOption: false, ID: false, ParentPath: "", Mandatory: false},
	{CliString: "useHttps", UsageMessage: "If https is to be used, default is false", DefaultValue: "false", ExpectedValue: "<true/false>", URLPathString: "", FilterOption: false, ID: false, ParentPath: "", Mandatory: false}}



func Usage(basecommand string, options []CliOption) {

	fmt.Print(" Usage: ", basecommand)
	for _, clioption := range defaultOptions {
		fmt.Print(" -" + clioption.CliString + " " + clioption.ExpectedValue)
	}
	for _, clioption := range options {
		fmt.Print(" -" + clioption.CliString + " " + clioption.ExpectedValue)
	}
	fmt.Println()
	PrintUsageForOptions(defaultOptions)
	PrintUsageForOptions(options)
}

func PrintUsageForOptions(options []CliOption){
	for _, clioption := range options {
		fmt.Println(" -" + clioption.CliString + ": " + clioption.UsageMessage)
	}
}

func GetUserValues(basecommand string, args []string, options []CliOption) (map[string]*string, error){

	var mandatoryOptions []string
	var values map[string]*string
	values = make(map[string]*string)
	var optionmap map[string]CliOption
	optionmap = make(map[string]CliOption)

	mandatoryOptions = make([]string, len(options)+len(defaultOptions), len(options)+len(defaultOptions))
	var noOfMandatoryOptions int


	for _, arg := range args {
		if( arg == "-h" || arg == "-help"){
			Usage(basecommand, options)
			return values, nil
		}

	}

	command := flag.NewFlagSet(basecommand, flag.ExitOnError)

	for _, clioption := range options {
		optionmap[clioption.CliString] = clioption
		values[clioption.CliString] = command.String(clioption.CliString, "", clioption.UsageMessage)
		if clioption.Mandatory {
			mandatoryOptions[noOfMandatoryOptions] = clioption.CliString
			noOfMandatoryOptions++
		}
	}
	fmt.Println("no of mandatory: " , noOfMandatoryOptions)
	for _, clioption := range defaultOptions {
		optionmap[clioption.CliString] = clioption
		values[clioption.CliString] = command.String(clioption.CliString, "", clioption.UsageMessage)
		if clioption.Mandatory {
			mandatoryOptions[noOfMandatoryOptions] = clioption.CliString
			noOfMandatoryOptions++
		}
	}

	fmt.Println(mandatoryOptions)
	command.Parse(args)


	if command.Parsed() {
		for i, mandatoryoption := range mandatoryOptions {
			if(i < noOfMandatoryOptions ) {
				fmt.Println(" Mandatory: " + mandatoryoption, " value: " , values[mandatoryoption])
				if values[mandatoryoption] == nil || *values[mandatoryoption] == "" {
					fmt.Println(" Missing: " + mandatoryoption + ". Must be specified for this command ")
					return values, ErrMandatoryParameterMissing
				}
			}
		}
	}else{
		fmt.Println("Could not parse args: " , args)
		Usage(basecommand, options)
		return values, ErrGeneral
	}


	return values, nil

}
