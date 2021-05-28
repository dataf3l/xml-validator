/** *

The purpose of this program is to verify from a certain path, all the files that
are in that folder, to determine if these files are a valid .gz file. To determine
this, the first 2 bytes of each file, if equal to [31  139 ... ], then it is a valid .gz file.

Returns a string with all paths to valid files repaired by commas.

*/

package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

var pathOFValidGZFilesList []string

//checkFileType Check if the provided path corresponds to a valid .gz file or not
//If it is not valid, it is ignored
func checkFileType(path string) {

	//if path is empty
	if len(path) < 1 {
		return
	}

	//Open file
	content, err := ioutil.ReadFile(path)
	if err != nil {
		//if path are not a file then return
		//log.Printf("checkFileType: error trying open file %s. Error: %v", path, err.Error())
		return
	}

	//byteHeader corresponds to the first byte of any valid .gz file
	byteHeader := []byte{31, 139}

	//Compare the first byte of the file with byteHeader
	//if they are the same (rest == 0), the path corresponds to a valid .gz file
	res := bytes.Compare(byteHeader, content[0:2])
	if res == 0 {
		//fmt.Println("Zip file: ", path)
		//fmt.Println("len: ", len(content))
		pathOFValidGZFilesList = append(pathOFValidGZFilesList, path)
	}
}

type Color string

func run(pathArgument string) {

	//pathArgument := "/home/oswaldo/github.com/xml-validator"

	cmd := exec.Command("find", pathArgument, "-maxdepth", "1", "-mindepth", "1")
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Printf("xml-validator: Error trying run command find in path: %s. Error: %v", pathArgument, err.Error())
	}

	pathList := strings.Split(out.String(), "\n")

	for _, path := range pathList {

		checkFileType(path)
	}

	//Returns a string with all paths to valid files repaired by commas
	fmt.Println(strings.Join(pathOFValidGZFilesList[:], ","))

}

func main() {

	flag.Parse()
	pathArgument := flag.Arg(0)
	//fmt.Println("filename: ", filename)

	//check if argument is empty
	if len(pathArgument) <= 0 {
		log.Println("xml-validator:main: invalid path: ", pathArgument)
		return
	}

	run(pathArgument)

}
