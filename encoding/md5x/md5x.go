package md5x

import (
	gmd5 "crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"strings"
)

// Result md5的结果
type Result []byte

// Bytes 获取md5的[]byte类型的结果
func (r Result) Bytes() []byte {
	return []byte(r)
}

// String 获取md5的string类型的结果
func (r Result) String() string {
	return hex.EncodeToString(r)
}

// UpperString 获取md5的大写的string类型的结果
func (r Result) UpperString() string {
	return strings.ToUpper(hex.EncodeToString(r))
}

// String16 获取md5的16位的string类型结果
func (r Result) String16() string {
	return hex.EncodeToString(r)[8:24]
}

// UpperString16 获取md5的16位的大写的string类型结果
func (r Result) UpperString16() string {
	return strings.ToUpper(hex.EncodeToString(r))[8:24]
}

// SumFile 计算文件的md5值
func SumFile(file *os.File) (Result, error) {
	hash := gmd5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return nil, err
	}
	return Result(hash.Sum(nil)), nil
}

// SumFileByPath 计算文件的md5值
func SumFileByPath(filePath string) (Result, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	hash := gmd5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return nil, err
	}
	return Result(hash.Sum(nil)), nil
}

// SumString 计算string的md5值
func SumString(data string) Result {
	h := gmd5.New()
	h.Write([]byte(data))
	return Result(h.Sum(nil))
}

// Sum 计算[]byte的md5值
func Sum(data []byte) Result {
	h := gmd5.New()
	h.Write(data)
	return Result(h.Sum(nil))
}
