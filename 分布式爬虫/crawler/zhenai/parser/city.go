package parser

import (
	"Projects/crawler/engine"
	"regexp"
)

//mockCityRe := `<a href="(http://localhost:8080/mock/album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`
const cityRe = `<a href="(http://localhost:8080/mock/album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`//`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1) //-1表示要所有的匹配
	result := engine.ParseResult{}
	for _, m := range matches {
		name := string(m[2])
		result.Items = append(result.Items, "User " + name)//人名
		result.Requests = append(
			result.Requests, engine.Request{
				Url: string(m[1]),
				//巧妙利用函数式编程解决name的问题->将参数name固定进去:函数不是在循环进行的时候运行
				//等退出循环的时候才会运行m
				ParserFunc: func(contents []byte) engine.ParseResult {
					return ParseProfile(contents, name)
				},
			})
	}
	return result
}
