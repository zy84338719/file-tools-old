package manager

import (
	"green-hat/model"
	"strings"
	"sync"
)

type FilterLog struct {
	filter        []string
	logfilterChan chan string
	res           sync.Map
}

func NewFilterLog(filter []string) *FilterLog {
	logfilterChan := make(chan string)
	return &FilterLog{
		filter:        filter,
		logfilterChan: logfilterChan}
}

func (f *FilterLog) FilterSourceData(b []byte) {
	s := string(b)
	all := strings.ReplaceAll(s, "\\", "")
	all = strings.ReplaceAll(all, "\"", "")
	f.logfilterChan <- all
}

func (f *FilterLog) Filter(fun func(d string, filter ...string)) {
	for s := range f.logfilterChan {
		fun(s, f.filter...)
	}
	model.WG.Done()
}

func (f *FilterLog) Close() {
	close(f.logfilterChan)
}
