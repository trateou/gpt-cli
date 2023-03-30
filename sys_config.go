package main

//lang list
//en
//cn

type SysConfig struct {
	Mode         string `mapstructure:"mode" comment:"c chat, t translate"`
	TransMode    int    `mapstructure:"trans_mode" comment:"0 terminal input, 1 file if SysConfig.Src is configed, TransMode will be set 1"`
	SrcFile      string `mapstructure:"src_file" comment:"filepath,not support folder now"`
	DestFile     string `mapstructure:"dest_file" comment:"filepath,not support folder now"`
	SrcLang      string `mapstructure:"src_lang" comment:"not sure if not specify the src lang, how chatgpt detects"`
	DestLang     string `mapstructure:"dest_lang"`
	OpenaiSecret string `mapstructure:"openai_secret"`
}
