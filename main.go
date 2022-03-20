package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"green-hat/config"
	"green-hat/manager"
	"green-hat/model"
	"green-hat/utils"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	feature := flag.String("tools", "file", "功能 file,image404")
	image404 := flag.String("imgfile", "zy.txt", "")
	filterFile := flag.String("filePrx", "error", "文件前缀")
	changeTool := flag.String("fileType", "push_type", "功能")
	flag.Parse()

	switch *feature {
	case "image404":
		read := utils.CsvNew().Read(*image404)
		for i := 0; i < len(read); i++ {
			xhttp := manager.HttpXNew()
			if i%1000 == 0 {
				println(fmt.Sprintln("%.4f", float64(i)*float64(100)/float64(len(read))) + "%")
			}
			if len(read[i]) == 0 {
				continue
			}
			h, err := xhttp.GetHttp("http://172.31.5.58:11007/api_server/v1/user/cover?user_ids=" + read[i][0])
			if err != nil {
				println("第一步请求问题", read[i][0])
				continue
			}
			info := model.UserInfo{}
			if json.Unmarshal(h, &info) != nil {

				println("转换错", read[i][0])
			}
			if !xhttp.GetUrlStatus(info.Data.UserList[0].Covers[0].Url) {
				println("url：=", info.Data.UserList[0].Covers[0].Url, " uid=", read[i][0])
			}
		}
	case "file":
		files := getFiles(*filterFile)
		result := config.NewResult()
		result.SwitchWith(*changeTool, files)
		for k, v := range result.GetMap() {
			fmt.Println(k, v)
		}
	}

}

func getFiles(filterFile string) []string {
	var files []string
	root, _ := os.Getwd()
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, filterFile) {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fmt.Println(file)
	}
	return files
}
