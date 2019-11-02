package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

//时区问题
//格式问题
var locales map[string]map[string]string

func main() {
	locales = make(map[string]map[string]string, 2)
	en := make(map[string]string, 10)
	en["time_zone"] = "America/Chicago"
	en["date_format"] = "%Y-%m-%d %H:%M:%S"
	locales["en"] = en

	cn := make(map[string]string, 10)
	cn["time_zone"] = "Asia/Shanghai"
	cn["date_format"] = "%Y年%m月%d日 %H时%M分%S秒"
	locales["cn_ZH"] = cn

	lang := "en"
	loc, _ := time.LoadLocation(msg(lang, "time_zone"))
	t := time.Now()
	t = t.In(loc)
	fmt.Println(t.Format(time.RFC3339))

	fmt.Println(date(msg(lang, "date_format"), t))
}

func date(fomate string, t time.Time) string {
	year, month, day := t.Date()
	hour, min, sec := t.Clock()
	//解析相应的%Y %m %d %H %M %S然后返回信息
	//%Y 替换成2012
	//%m 替换成10
	//%d 替换成24
	s1 := strings.Replace(fomate, "%Y", strconv.Itoa(year), -1)
	s2 := strings.Replace(s1, "%m", month.String(), -1)
	s3 := strings.Replace(s2, "%d", strconv.Itoa(day), -1)
	s4 := strings.Replace(s3, "%H", strconv.Itoa(hour), -1)
	s5 := strings.Replace(s4, "%M", strconv.Itoa(min), -1)
	return strings.Replace(s5, "%S", strconv.Itoa(sec), -1)
}

func msg(locale, key string) string {
	if v, ok := locales[locale]; ok {
		if v2, ok := v[key]; ok {
			return v2
		}
	}
	return ""
}
