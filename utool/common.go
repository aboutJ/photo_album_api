package utool

import (
	"crypto/md5"
	"fmt"
)

func Md5Crypt(str string, salt string) string {
	var result string = salt
	if len(str) > 0 {
		result = fmt.Sprintf("%s%s", str, salt)
	}
	ret := md5.Sum([]byte(result))
	return fmt.Sprintf("%x", ret)
}
