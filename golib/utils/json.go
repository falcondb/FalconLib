package utils

import (
	"bytes"
	"encoding/json"
)

//func main() {
//	var content []byte
//	var err error
//	if len(os.Args) == 1 {
//		content, err = ioutil.ReadAll(os.Stdin)
//	} else {
//		content, err = ioutil.ReadFile(os.Args[1])
//	}
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	jsons, err := JsonPretty(content)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Print(string(jsons))
//}

func JsonPretty(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	return out.Bytes(), err
}
