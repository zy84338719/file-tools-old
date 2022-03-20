package config

import (
	"fmt"
	"green-hat/manager"
	"green-hat/utils"
	"strings"
)

func (r *Result) PushChannel(filePath string) {
	file, err := utils.NewFile(filePath)
	if err != nil {
		fmt.Errorf("错误 %+v", err)
	}
	log := manager.NewFilterLog([]string{"存储数据为"})
	go log.Filter(r.PushChannelFilter)
	file.ReadLine(log.FilterSourceData)
	log.Close()

}

func (r *Result) PushChannelFilter(s string, filter ...string) {
	for _, f := range filter {
		if !strings.Contains(s, f) {
			return
		}
	}
	dataList := strings.Split(s, " ")
	if len(dataList) > 6 {
		r.data.Store(dataList[5], 1)
	}
}
