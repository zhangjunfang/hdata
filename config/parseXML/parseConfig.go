package parseXML

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/zhangjunfang/hdata/config"
)

func ParseXML() {

	workPath := os.Args[0]
	workDir := filepath.Dir(workPath)
	workDir, err := filepath.Abs(workDir)
	if err != nil {
		fmt.Println(err)
	}
	workDir = strings.Replace(workDir, "\\", "/", -1)
	content, err := ioutil.ReadFile(workDir)
	if err != nil {
		workDir, err = os.Getwd()
		workDir = strings.Replace(workDir, "\\", "/", -1)
		workDir = filepath.Clean(workDir + "/../hdata.xml")
		content, err = ioutil.ReadFile(workDir)
		if err != nil {
			fmt.Println(err)
		}
	}
	hdata := &config.Hdata{}
	xml.Unmarshal(content, hdata)
	for _, v := range hdata.Configuration.Properties {
		fmt.Println(v.Name)
		fmt.Println(v.Value)
		fmt.Println(v.Description)
	}
}
