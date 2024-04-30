package test

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var mRootDir string
var mSeparator string
var mJsonData map[string]any

const mJsonFileName = "dir.json"

func loadJson() {
	mSeparator = string(filepath.Separator)
	workDir, _ := os.Getwd()
	mRootDir = workDir[0:strings.LastIndex(workDir, mSeparator)]

	jsonBytes, _ := os.ReadFile(workDir + mSeparator + mJsonFileName)

	err := json.Unmarshal(jsonBytes, &mJsonData)
	if err != nil {
		panic("加载json错误")
	}

}

func parseMap(mapData map[string]any, parentDir string) {
	for k, v := range mapData {
		switch v.(type) {
		case string:
			{
				path := v.(string)
				if path == "" {
					continue
				}

				if parentDir != "" {
					path = parentDir + mSeparator + path
					if k == "text" {
						parentDir = path
					}
				} else {
					parentDir = path
				}
				creteDir(path)
			}
		case []any:
			parseArray(v.([]any), parentDir)
		}
	}
}

func parseArray(jsonData []any, parentDir string) {
	for _, v := range jsonData {
		mapV, _ := v.(map[string]any)
		parseMap(mapV, parentDir)
	}
}

func creteDir(path string) {
	if path == "" {
		return
	}
	fmt.Println(path)
	err := os.MkdirAll(mRootDir+mSeparator+path, fs.ModePerm)
	if err != nil {
		panic("文件夹创建失败")
	}
}
func TestGenerateDir001(t *testing.T) {
	loadJson()
	parseMap(mJsonData, "")
}
