# gpt-cli
- a terminal tool based on chatgpt

[中文](/readme_cn.md)

## Features
1. Chatting
2. File Translator (markdown, text)
3. Yet to be explored

## Installation
```shell
  go get https://github.com/trateou/gpt-cli
```
## Usage 

Parameter Configuration

### Based on config.yml (not recommended, still being optimized)
```yaml
mode: c   		#c for chat, t for translation
trans_mode: 	#Translation mode enabled. 0 for direct terminal output todo, 1 for text translation
src_file: 		#Translation mode enabled. Source file
dest_file:	 	#Translation mode enabled. Target file. If not specified, defaults to ${dest_lang}_src_file
src_lang: "" 	# Translation mode enabled. Optional. Will be automatically recognized by chatgpt if not specified
dest_lang:   	# Translation mode enabled. Target language. Not Null. Example: 'cn, us, jp'
openai_secret:  # Translation mode enabled. Openai token
ss: 8096        # Translation mode enabled. Paragraph size for each translation. Default 8k
```
example
```yaml
mode: t   		#c for chat, t for translationtrans_mode: 1 	# Translation mode enabled. 0 for direct terminal output todo, 1 for text translation
src_file: /path/to/src_file 		# Translation mode enabled. File to be translated. If this option is not empty, trans_mode=1
dest_file:/path/to/dest_file	 	# Translation mode enabled. Target file. If not specified, defaults to ${dest_lang}_src_file
src_lang: "" 	# Translation mode enabled. Optional. Will be automatically recognized by chatgpt if not specified
dest_lang: "cn"   	# Translation mode enabled. Target language. Not Null. Example: 'cn, us, jp'
openai_secret: your_openai_token # Translation mode enabled. Openai token
ss: 8096       # Translation mode enabled. Paragraph size for each translation. Default 8k
```
### Chatting on the Command Line
```shell
gpt-cli --mode c
```
Translation
```
shellgpt-cli --mode t --dest_lang cn --openai_secret your_openai_token --src_file /path/to/src_file --dest_file /path/to/dest_file
```
## TODO
- [ ] Add help command