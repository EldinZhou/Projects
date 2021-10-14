package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	//test去网络外部拿数据这样不可取
	//test应该用本地去测试
	contents, err := ioutil.ReadFile("zhenai.html")//fetcher.Fetch("http://www.zhenai.com/zhenghun")

	if err != nil {
		panic(err)
	}

	result := ParseCityList(contents)
	//fmt.Printf("%s\n", contents)
	//verify result
	const resultSize = 470
	expectedUrls := []string {//const无法定义切片
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}

	expectedCities := []string {//const无法定义切片
		"City 阿坝", "City 阿克苏", "City 阿拉善盟",
	}

	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s; nut was %s", i, url, result.Requests[i].Url)
		}
	}

	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d"+
			     "requests but had %d", resultSize, len(result.Requests))
	}

	for i, city := range expectedCities {
		if result.Items[i].(string) != city {
			t.Errorf("expected url #%d: %s; nut was %s", i, city, result.Items[i].(string))
		}
	}

	if len(result.Items) != resultSize {
		t.Errorf("result should have %d"+
				 "items but had %d", resultSize, len(result.Requests))
	}

}