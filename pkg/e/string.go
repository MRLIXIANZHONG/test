package e

import (
	"crypto/md5"
	"fmt"
	"regexp"
)

func MD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return fmt.Sprintf("%x", h.Sum(nil))//格式化输入小写16进制
}

func IsMobile(str string) bool {
	reg := regexp.MustCompile(`^(1\d{10})$`)
	return reg.MatchString(str)
}

func IsNumeric(str string) bool {
	reg := regexp.MustCompile(`^(\d+)$`)
	return reg.MatchString(str)
}
