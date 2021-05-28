package tools

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

// check the existence of the file
func FileExists(path string) bool {

	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println("File not found: ", path)
		return false
	}
	return true
}

//IsValidXML Check if the XML structure is valid
func IsValidXML(data []byte) error {
	err := xml.Unmarshal(data, new(interface{}))
	if err != nil {
		log.Println("xml-chk:IsValidXML: Error trying Unmarshal data error: ", err.Error())
		return err
	}
	return nil
}
