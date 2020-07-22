package models

import (
	"crypto/md5"
	"fmt"
	"time"

	"github.com/astaxie/beego"
)

func Hello(in string) (out string) {
	out = in + "hello world!"
	return out
}

func UnixToDate(timestmap int) string {

	t := time.Unix(int64(timestmap), 0)
	return t.Format("2006-01-02 15:04:05")
}

func DateToUnix(str string) int64 {
	template := "2006-01-02 15:04:05"
	t, err := time.ParseInLocation(template, str, time.Local)

	if err != nil {
		beego.Info(err)
		return 0
	}
	return t.Unix()
}

func GetUnix() int64 {
	return time.Now().Unix()
}

func GetDate() string {
	template := "2006-01-02 15:04:05"
	return time.Now().Format(template)
}

func Md5(str string) string {
	data := []byte("str")
	return fmt.Sprintf("%X\n", md5.Sum(data))
}
