package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type data struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Place string `json:"place,omitempty"`
}

// func main() {
// 	content, err := os.ReadFile("test.json")
// 	checkErr(err)

// 	var x data
// 	err = json.Unmarshal(content, &x)
// 	checkErr(err)
// 	fmt.Printf("%+v\n", x)

// 	y := data{Id: "4", Name: "Dibek"}
// 	data, err := json.Marshal(y)
// 	checkErr(err)
// 	fmt.Println(string(data))

// 	hello.Hello()
// }

func main() {
	parseJsonLogs()
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func parseJsonLogs() {
	f, err := os.Open("json.log")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		fmt.Printf("line: %s\n", scanner.Text())
		var temp map[string]interface{}
		err := json.Unmarshal([]byte(scanner.Text()), &temp)
		if err != nil {
			log.Println("[ERROR] unmarshal json data", err)
			continue
		}

		//create map[string]string for further processing
		result := make(map[string]string)

		//convert map[string]interface{} to map[string]string
		for k, v := range temp {
			result[k] = fmt.Sprint(v)
		}

		fmt.Println("result: ", result)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
