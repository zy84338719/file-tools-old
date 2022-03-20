package utils

import (
	"bufio"
	"io"
	"os"
)

type File struct {
	file *os.File
}

func NewFile(path string) (*File, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return &File{file: f}, err
}

func (f *File) ReadLine(hookfn func([]byte)) error {
	bfRd := bufio.NewReader(f.file)
	for {
		line, err := bfRd.ReadBytes('\n')
		hookfn(line)    //放在错误处理前面，即使发生错误，也会处理已经读取到的数据。
		if err != nil { //遇到任何错误立即返回，并忽略 EOF 错误信息
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
	f.file.Close()
	return nil
}
