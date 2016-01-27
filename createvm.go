package main

import (
	"fmt"
	"github.com/swapneshb/flags/common"
	"os"
)

func main() {

	clioptions := []common.CliOption{{CliString: "name", UsageMessage: "The name of the virtual machine to be created.", ExpectedValue: "hostname", URLPathString: "/Operation/vm/create", FilterOption: false, ID: false, ParentPath: "", Mandatory: true},
		{CliString: "cpus", UsageMessage: "The number of CPUs required", DefaultValue: "2", ExpectedValue: "<numCpus>", URLPathString: "", FilterOption: false, ID: false, ParentPath: "", Mandatory: false},
		{CliString: "mem", UsageMessage: "The amount of memory required (GB)", DefaultValue: "2", ExpectedValue: "<memValue>", URLPathString: "", FilterOption: false, ID: false, ParentPath: "", Mandatory: false}}

	if len(os.Args) == 1 {
		common.Usage("createvm", clioptions)
		os.Exit(2)
	}

	var err error
	var values map[string]*string
	values, err = common.GetUserValues("createvm", os.Args[1:], clioptions )
	fmt.Println("Done, err: ", err)
	for option, value := range values {
		if value != nil {
			fmt.Println(option, ":", *value)
		}

	}
}
