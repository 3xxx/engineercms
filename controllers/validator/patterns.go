package validator

import "regexp"

const (
	word        string = "^[a-zA-Z0-9_]+$"
	alnum       string = "^[a-zA-Z0-9]+$"
	mobilePhone string = "^1(3[0-9]|4[57]|5[0-35-9]|7[0135678]|8[0-9])\\d{8}$"
)

var (
	rxWord        = regexp.MustCompile(word)
	rxAlnum       = regexp.MustCompile(alnum)
	rxMobolePhone = regexp.MustCompile(mobilePhone)
)
