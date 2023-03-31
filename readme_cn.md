# gpt-cli

一个基于chatgpt的终端工具

## 特性

1. 聊天
2. 文件翻译器(markdown,text)
3. 待挖掘

## 安装

```shell
  go install github.com/trateou/gpt-cli@latest
```

## 使用

参数配置

### 基于config.yml(不推荐,还在优化)

```yaml
mode: c   		#c 聊天, t 翻译
trans_mode: 	#翻译模式生效.0 直接终端输出 todo, 1 文本翻译
src_file: 		#翻译模式生效. 源文件
dest_file:	 	#翻译模式生效. 目标文件.如果不指定,默认为${dest_lang}_src_file
src_lang: "" 	# 翻译模式生效.可以不指定,由chatgpt自动识别
dest_lang:   	#翻译模式生效.目标语言,Not Null.example ‘cn,us,jp’
openai_secret: #翻译模式生效. openai token
ss: 8096       #翻译模式生效.每次翻译的文件段落大小.默认8k
```

example

```yaml
mode: t   		#c 聊天, t 翻译
trans_mode: 1 	#翻译模式生效.0 直接终端输出 todo, 1 文本翻译
src_file: /path/to/src_file 		#翻译模式生效. 要翻译的文件, 如果该选项不为空,trans_mode=1
dest_file:/path/to/dest_file	 	#翻译模式生效. 目标文件.如果不指定,默认为${dest_lang}_src_file
src_lang: "" 	# 翻译模式生效.可以不指定,由chatgpt自动识别
dest_lang: "cn"   	#翻译模式生效.目标语言,Not Null.example ‘cn,us,jp’
openai_secret:  your_openai_token #翻译模式生效. openai token
ss: 8096       #翻译模式生效.每次翻译的文件段落大小.默认8k
```

### 命令行

聊天

```shell
gpt-cli --mode c
```

翻译

```shell
gpt-cli --mode t --dest_lang cn --openai_secret your_openai_token --src_file /path/to/src_file --dest_file /path/to/dest_file

```



## TODO

- [ ] ​	添加help命令