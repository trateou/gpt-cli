package main

import (
	"bufio"
	"fmt"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
	"os"
	"strconv"
	"strings"
)

func addLangPrefix(path string, lang string) string {
	return lang + "_" + path
}

// ParseFileAsSegmentWithFileSize Reads the file by the specified segment size
func ParseFileAsSegmentWithFileSize(path string, bufSize int) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	var bufferSize = bufSize
	var segment string
	var dataList []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// 读取一行数据
		line := scanner.Text()
		if len(line) >= bufferSize {
			if len(segment) != 0 {
				dataList = append(dataList, segment)
				segment = ""
			}
			dataList = append(dataList, string(line))
			continue
		}

		// todo if one line's size bigger than chatgpt token limitation
		if len(segment)+len(line) > bufferSize {
			dataList = append(dataList, segment)
			segment = ""
		}
		segment += string(line)
	}
	if len(segment) != 0 {
		dataList = append(dataList, segment)
	}

	return dataList
}

func main() {
	// add driver for support yaml content
	config.AddDriver(yaml.Driver)

	err := config.LoadFiles("config.yml")
	if err != nil {
		panic(err)
	}
	var sys_config SysConfig

	err = config.LoadFlags(os.Args)

	if err != nil {
		panic(err)
	}
	config.BindStruct("", &sys_config)
	//config check
	if len(sys_config.Mode) == 0 {
		panic("Config.Mode Is Empty")
	}
	if sys_config.Mode != "c" && sys_config.Mode != "t" {
		panic("Config.Mode is invalid value:" + sys_config.Mode)
	}
	if sys_config.Mode == "t" {
		if sys_config.TransMode != 0 && sys_config.TransMode != 1 {
			panic("Config.TransMode is invalid value:" + strconv.Itoa(sys_config.TransMode))
		}
		if sys_config.TransMode == 1 {
			if len(sys_config.SrcFile) == 0 {
				panic("Config.SrcFile Is Empty")
			} else {
				sys_config.TransMode = 1
				if len(sys_config.DestLang) == 0 {
					panic("Config.Destlang Is Empty")
				}
				if len(strings.TrimSpace(sys_config.DestFile)) == 0 {
					sys_config.DestFile = addLangPrefix(sys_config.SrcFile, sys_config.DestLang)
				}
				if sys_config.SegementSize == 0 {
					sys_config.SegementSize = 8096
				}
			}
		}
	}
	if len(sys_config.OpenaiSecret) == 0 {
		panic("Config.openaiSecret Is Empty")
	}

	if sys_config.Mode == "c" {
		//chat
		model := NewGptTrans(&sys_config)
		model.Chat()
	} else if sys_config.Mode == "t" {
		fmt.Println("translating file:" + sys_config.SrcFile + " to " + sys_config.DestFile)
		// crate 7K buffer
		bufferSize := sys_config.SegementSize
		model := NewGptTrans(&sys_config)
		var content string
		var lines = ParseFileAsSegmentWithFileSize(sys_config.SrcFile, bufferSize)
		var lineCount = len(lines)
		for index, segment := range lines {
			fmt.Println("translating index:" + strconv.Itoa(index) + "/" + strconv.Itoa(lineCount))
			d := model.TransContent("en", "cn", segment)
			fmt.Println("response content size:" + strconv.Itoa(len(d)))
			content += d
		}
		f, err := os.Create(sys_config.DestFile)
		if err != nil {
			panic(err)
		}
		writer := bufio.NewWriter(f)
		_, err = writer.WriteString(content)
		if err != nil {
			panic(err)
		}

		err = writer.Flush()
		if err != nil {
			panic(err)
		}
		fmt.Println("task finished. from:" + sys_config.SrcLang + " to " + sys_config.DestLang)
	}

}
