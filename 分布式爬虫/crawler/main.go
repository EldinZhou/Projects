package main

import (
	"Projects/crawler/engine"
	"Projects/crawler/zhenai/parser"
)

//const text = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

//这个就算是城市列表的解析器->各个城市名字+对应的url: printCityList
//城市解析器: ->解析出每个人的名字+对应的url
//用户解析器: ->获取当前人的信息->现阶段所要的结果: 单任务版爬虫的目标

//抽象一下解析器: input:  utf-8的文本
//				 output: Request列表{url, 对应的Parser}, DataItem列表(需要的有价值的数据)
//每个url对应的解析器可能是不一样的，所以需要将其封装在Request里面
//比如某个城市的用户页面，第一页每个用户的url对应的解析器就是用户解析器，但是下一页对应的是城市解析器
//只有通过城市解析器才可以链接到下一页的用户
//单任务版是获取所有城市第一页用户的数据

/*
func printCityList(contents []byte) {
	re := regexp.MustCompile(text)
	matches := re.FindAllSubmatch(contents, -1) //-1表示要所有的匹配
	//matches: [][]byte->[]string
	//subMatches: [][][]byte->[][]string
	for _, m := range matches {
		fmt.Printf("City: %s, URL: %s\n", m[2], m[1])
	}
	fmt.Printf("Matches found : %d\n", len(matches))
}*/

//正则表达式 or CSS选择器找到想要的字符串
//mockUrl := http://localhost:8080/mock/www.zhenai.com/zhenghun
func main() {
	engine.Run(engine.Request{
		Url: 		"http://localhost:8080/mock/www.zhenai.com/zhenghun",//"http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}

