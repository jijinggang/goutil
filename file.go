/*
golang utils package

doc from: ` godocdown . > README.md `

##Install

    go get -u github.com/jijinggang/goutil

*/
package goutil

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"
	"strings"
)

//复制文件，把srcName复制到dstName
func CopyFile(srcName, dstName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	//确保创建目标目录
	CreateDir(dstName)

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}

//写文本文件，如果文件存在，则覆盖
func WriteStringFile(file string, str string) (written int, err error) {
	//确保创建目标目录
	CreateDir(file)
	dst, err := os.Create(file)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.WriteString(dst, str)
}

//递归创建文件所在的目录
func CreateDir(file string) {
	path := GetPath(file)
	os.MkdirAll(path, os.ModeDir)
}

//格式化路径，统一用/替代\，并且去掉最后面的/
func FormatPath(path string) string {
	path = strings.Replace(path, "\\", "/", -1)
	path = strings.TrimRight(path, "/")
	return path
}

//获取file或目录相对其一个父目录dir的相对路径名
func GetRelativePath(file string, dir string) string {
	rfile := file[len(dir):] //相对路径
	rfile = strings.TrimLeft(rfile, "/")
	return rfile
}

//获取一个文件的路径，不带最后的/
func GetPath(file string) string {
	file = FormatPath(file)
	pos := strings.LastIndex(file, "/")
	return file[0:pos]
}

//获取文件的Md码
func GetFileMd5(file string) string {
	f, err := os.Open(file)
	if err != nil {
		panic(err.Error())
	}
	h := md5.New()
	const BUFSIZE = 10240
	buf := make([]byte, BUFSIZE)
	for {
		rlen, err := f.Read(buf)
		if err != nil {
			break
		}
		h.Write(buf[0:rlen])
	}
	return hex.EncodeToString(h.Sum(nil))
}

//获取一个目录的所有子文件
func GetAllFiles(dir string) []string {
	files := []string{}
	filepath.Walk(dir, func(file string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if !f.IsDir() {
			file = FormatPath(file)
			files = append(files, file)
		}
		return nil
	})
	return files
}

//判断文件或路径是否存在
func Exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}
