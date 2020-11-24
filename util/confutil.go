package util

import (
	"encoding/json"
	"io/ioutil"
	"regexp"
	"strings"
)

// LoadConfigJSON 解析json文件为一个对象 忽略#注释的行
func LoadConfigJSON(path string, obj interface{}) (err error) {
	var bytes []byte
	//ReadFile从filename指定的文件中读取数据并返回文件的内容。成功的调用返回的err为nil而非EOF
	bytes, err = ioutil.ReadFile(path)
	if err != nil {
		return
	}
	//Compile解析并返回一个正则表达式,MustCompile类似Compile但会在解析失败时panic，主要用于全局正则表达式变量的安全初始化
	re1 := regexp.MustCompile("^#.*$")
	//func Split(s, sep string) []string
	//用去掉s中出现的sep的方式进行分割，会分割到结尾，并返回生成的所有片段组成的切片
	lines := strings.Split(string(bytes), "\n")
	var loc []int
	var cnt string
	for _, line := range lines {
		//Find返回保管正则表达式re在b中的最左侧的一个匹配结果的起止位置的切片
		//re := regexp.MustCompile("ab?")
		//fmt.Println(re.FindStringIndex("tablett"))
		//输出:[1 3]
		loc = re1.FindStringIndex(line)
		if len(loc) == 0 {
			cnt += line
		}
	}
	//json类型解析为结构体
	//解析json编码的数据并将结果存入obj指向的值
	err = json.Unmarshal([]byte(cnt), obj)
	return
}
