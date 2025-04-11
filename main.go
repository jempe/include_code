package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/jempe/include_code/utils"
)

const version = "0.0.2"

// What's new
// Ability to insert code between HTML and JS/CSS C style comments

func main() {
	if len(os.Args) < 2 {
		fmt.Println("include_code v:", version)
		fmt.Println("Search the comments <!--include:file_to_include:--> and <!--includeend--> or /*--include:file_to_include:--*/ and /*--includeend--*/ and insert the code of the file_to_include between them")
		fmt.Println("Usage of include_code:")
		fmt.Println("include_code <file>")
	} else {
		mainFile := os.Args[1]

		if utils.Exists(mainFile) {
			if utils.IsDirectory(mainFile) {
				fmt.Println(mainFile, "is a directory")
			} else {
				dir := filepath.Dir(mainFile)

				content, err := ioutil.ReadFile(mainFile)
				if err != nil {
					fmt.Println(err)
				}

				mainContent := string(content)

				rHTML := regexp.MustCompile("<!--include:[^:]+:-->")
				rCStyle := regexp.MustCompile("\\/\\*--include:[^:]+:--\\*\\/")

				includesHTML := rHTML.FindAll(content, -1)
				includesCStyle := rCStyle.FindAll(content, -1)

				includes := append(includesHTML, includesCStyle...)

				saveFile := true

				for _, include := range includes {
					includeFileName := strings.Replace(strings.Replace(strings.Replace(strings.Replace(string(include), ":-->", "", 1), "<!--include:", "", 1), ":--*/", "", 1), "/*--include:", "", 1)
					includeFile := dir + "/" + includeFileName

					if utils.Exists(includeFile) && !utils.IsDirectory(includeFile) {
						fmt.Println("include:", includeFile)

						fileContent, err := ioutil.ReadFile(includeFile)
						if err != nil {
							fmt.Println(err)

							saveFile = false
						} else {

							var startTag, endTag string
							if strings.Contains(string(include), "<!--include:") {
								startTag = "<!--include:" + includeFileName + ":-->"
								endTag = "<!--includeend-->"
							} else {
								startTag = "/*--include:" + includeFileName + ":--*/"
								endTag = "/*--includeend--*/"
							}

							mainContent, err = utils.InsertBeetweenMatches(mainContent, startTag, endTag, "\n"+string(fileContent))
							if err != nil {
								fmt.Println(err)

								saveFile = false
							}
						}
					} else {
						fmt.Println("file:", includeFile, "doesn't exist or is a directory")
					}
				}

				if saveFile {
					fmt.Println("saving", mainFile)

					err = ioutil.WriteFile(mainFile, []byte(mainContent), 0644)
					if err != nil {
						fmt.Println(err)
					}

				}
			}
		} else {
			fmt.Println("file:", mainFile, "doesn't exists")
		}
	}
}
