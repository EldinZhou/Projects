package parser

import (
	"Projects/crawler/model"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseProfile(contents, "猖狂猫儿.")

	if len(result.Items) != 1 {
		t.Errorf("Result should contain 1 element; but was %v", result.Items)
	}

	profile := result.Items[0].(model.Profile)

	expectedProfile := model.Profile {
		Name : "猖狂猫儿.",
		Gender : "女",
		Age : 99,
		Height : 141,
		Weight : 202,
		Income : "财务自由",
		Marriage : "离异",
		Education :	"初中",
		Occupation : "产品经理",
		Hukou :		"郑州市",
		Xingzuo :	"天蝎座",
		House :		"租房",
		Car :		"有车",
	}

	if profile != expectedProfile {
		t.Errorf("expected %v; but was %v", expectedProfile, profile)
	}
}