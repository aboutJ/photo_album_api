package test

import (
	"fmt"
	"photo_album/config"
	"photo_album/utool"
	"testing"
)

func TestMd5(t *testing.T) {
	str := utool.Md5Crypt("123456", config.SaltPassword)
	fmt.Println(str == "787cdf241f8d06bdc71a3472ba16afa5")
}
