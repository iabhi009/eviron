package validate

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

func ValidateEnviromentVariables(input interface{}) error {
	var err error
	inputJson, err := json.Marshal(input)
	if err != nil {
		fmt.Println("error -  ", err.Error())
	}
	tempVar := map[string]interface{}{}
	// x = temp
	err = json.Unmarshal(inputJson, &tempVar)
	if err != nil {
		fmt.Println("error - ", err.Error())
	}
	enviromentVaribles := getEnviromentVariables()
	for key, _ := range tempVar {
		flag := false
		for envKey, _ := range enviromentVaribles {
			if envKey == key {
				flag = true
			}
		}
		if flag != true {
			errStr := "Unable to find enviroment variables key - " + key
			return errors.New(errStr)
		}
	}
	return err
}

func getEnviromentVariables() map[string]string {
	output := map[string]string{}
	for _, val := range os.Environ() {
		split := strings.Split(val, "=")
		output[split[0]] = split[1]
	}
	return output
}
