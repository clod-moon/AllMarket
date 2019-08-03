package webclient

import (
	"bytes"
	//"compress/gzip"
	"io/ioutil"
	"math/rand"
	"time"
	"compress/flate"
)

var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// getRandomString 返回随机字符串
func getRandomString(n uint) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// getUinxMillisecond 取毫秒时间戳
func getUinxMillisecond() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// unGzipData 解压gzip的数据
func unGzipData(buf []byte) ([]byte, error) {
	r := flate.NewReader(bytes.NewReader(buf))
	defer r.Close()
	return ioutil.ReadAll(r)
}
