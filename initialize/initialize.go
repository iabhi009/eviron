package initialize

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func InitializeEnviromentVariables(input interface{}) error {
	var err error
	enviromentMap := map[string]string{}
	for _, val := range os.Environ() {
		split := strings.Split(val, "=")
		enviromentMap[split[0]] = split[1]
	}

	enviromentStr, err := json.Marshal(enviromentMap)
	if err != nil {
		fmt.Println("error - ", err.Error())
		return err
	}
	err = json.Unmarshal(enviromentStr, &input)
	if err != nil {
		fmt.Println("error - ", err.Error())
		return err
	}
	return err
}
