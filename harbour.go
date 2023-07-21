package harbour

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type structInterface interface{}

func ImportConfig(fileName string, structVar structInterface) {
	//open file to variable
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	defer file.Close()

	//read file data
	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}

	err = json.Unmarshal(data, structVar)
	if err != nil {
		fmt.Println("Error unmarshaling JSON: ", err)
		return
	}

}
