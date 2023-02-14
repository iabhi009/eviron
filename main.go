package main

import (
	"encoding/json"
	"environ/initialize"
	"environ/validate"
	"fmt"
	"os"
)

func init() {
	os.Setenv("FargateCommandBucket", "a")
	os.Setenv("Region", "b")
	os.Setenv("InvokeLambdaFunctionName", "c")
	os.Setenv("test4", "d")
}

type Env struct {
	FargateCommandBucket     string `json:"FargateCommandBucket"`
	Region                   string `json:"Region"`
	InvokeLambdaFunctionName string `json:"InvokeLambdaFunctionName,omitempty"`
}

func main() {
	x := Env{}
	err := SetupEnviroment(&x)
	if err != nil {
		fmt.Println("Error: Unable to setup enviroment variables ", err.Error())
		return
	}
	fmt.Println("final - ", x)
	str, err := json.Marshal(x)
	if err != nil {
		fmt.Println("error ", err.Error())
	}
	fmt.Println(string(str))
}
func SetupEnviroment(input interface{}) error {
	var err error
	if err = validate.ValidateEnviromentVariables(input); err != nil {
		fmt.Println("Error: Unable to validate enviroment variables ", err.Error())
		return err
	}

	if err = initialize.InitializeEnviromentVariables(&input); err != nil {
		fmt.Println("Error: Unable to initialize enviroment variables ", err.Error())
		return err
	}
	return err
}
