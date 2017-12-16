package validator

import (
	"strings"

	"github.com/asaskevich/govalidator"
)

// add by hzwy23
// IsWord check if the string contains letters and numbers. Empty string is invalid.
// params accept 2 arguments
// first argument is min len.
// second argument is max len.
func IsWord(str string, params ...int) bool {
	if IsEmpty(str) {
		return false
	}
	return rxWord.MatchString(str)
}

// add by hzwy23
// IsEmpty check if the string contains any character, but spacing is invalid.
func IsEmpty(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}

// add by hzwy23
// IsDate check if the string is date ,the formatter support 2006-01-02 or 2006/01/02.
func IsDate(str string, format ...string) bool {

	if len(format) == 0 {

		if len(strings.Split(str, "/")) > 1 {
			return govalidator.IsTime(str, "2006/01/02")
		}

		if len(strings.Split(str, "-")) > 1 {
			return govalidator.IsTime(str, "2006-01-02")
		}

		return false
	}

	return govalidator.IsTime(str, format[0])
}

// add by hzwy23
// IsURI check if the string is URI address, relation address and absolute address is valid.
func IsURI(str string) bool {
	relation := false
	for idx, val := range str {
		if val == '.' {
			relation = true
			continue
		}
		if val == '/' || val == '\\' {
			if idx < len(str)-1 {
				if str[idx+1] == '.' {
					relation = true
					continue
				}
			}
		}

		if relation && (val != '/' && val != '\\') {
			return false
		}
		return govalidator.IsRequestURI(str[idx:])
	}
	return false
}

// add by hzwy23
// IsMobilePhone check if the str is mobile phone number.
func IsMobilePhone(str string) bool {
	if IsEmpty(str) {
		return false
	}
	return rxMobolePhone.MatchString(str)
}

func IsAlnum(str string) bool {
	if IsEmpty(str) {
		return false
	}
	return rxAlnum.MatchString(str)
}

// check if string str is a member of the set of strings params
func IsIn(str string, params ...string) bool {
	return govalidator.IsIn(str, params...)
}

func IsEmail(str string) bool {
	return govalidator.IsEmail(str)
}

func IsNumeric(str string) bool {
	return govalidator.IsNumeric(str)
}

func IsNull(str string) bool {
	return govalidator.IsNull(str)
}

func IsFloat(str string) bool {
	return govalidator.IsFloat(str)
}
