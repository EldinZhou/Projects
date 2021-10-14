package parser

import (
	"Projects/crawler/engine"
	"Projects/crawler/model"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(`(\d+)岁`) //\d:表示是数字
var heightRe = regexp.MustCompile(`<td><span class="label">身高：</span><span field="">(\d+)CM</span></td>`)
var incomeRe = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
var weightRe = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">(\d+)KG</span></td>`)
var genderRe = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var xingzuoRe = regexp.MustCompile(`<td><span class="label">星座：</span><span field="">([^<]+)</span></td>`)
var marriageRe = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
var educationRe = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
var occupationRe = regexp.MustCompile(`<td><span class="label">职业：</span><span field="">([^<]+)</span></td>`)
var hukouRe = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var houseRe = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var carRe = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)
//var guessRe = regexp.MustCompile(`<a class="exp-user-name"[^>]*href="(.*album\.zhenai\.com/u/[\d]+)">([^<]+)</a>`)
//var idUrlRe = regexp.MustCompile(`.*album\.zhenai\.com/u/([\d]+)`)


func ParseProfile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name
	//fmt.Println(string(contents))
	age, _ := strconv.Atoi(extractString(contents, ageRe))
	weight, _ := strconv.Atoi(extractString(contents, weightRe))
	height, _ := strconv.Atoi(extractString(contents, heightRe))
	profile.Age = age
	profile.Weight = weight
	profile.Height = height
	profile.Marriage = extractString(contents, marriageRe)
	profile.Income = extractString(contents, incomeRe)
	profile.Gender = extractString(contents, genderRe)
	profile.Car = extractString(contents, carRe)
	profile.Education = extractString(contents, educationRe)
	profile.Hukou = extractString(contents, hukouRe)
	profile.House = extractString(contents, houseRe)
	profile.Occupation = extractString(contents, occupationRe)
	profile.Xingzuo = extractString(contents, xingzuoRe)

	result := engine.ParseResult{
		Items : []interface{}{profile},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents) //只找一个就可以

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
