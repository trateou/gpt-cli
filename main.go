package main

import (
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
	"os"
	"strconv"
)

func addLangPrefix(path string, lang string) string {
	return lang + "_" + path
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
				if len(sys_config.DestFile) == 0 {
					sys_config.DestFile = addLangPrefix(sys_config.SrcFile, sys_config.DestLang)
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
		f, _ := os.ReadFile(sys_config.SrcFile)
		model := NewGptTrans(&sys_config)
		content := model.Trans("en", "cn", string(f))
		err = os.WriteFile(sys_config.DestFile, []byte(content), 0644)
		if err != nil {
			panic(err)
		}
	}

}
