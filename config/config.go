package config

import (
	"sync"
)

type Result struct {
	data sync.Map
}

func NewResult() *Result {
	return &Result{}
}

func (r *Result) GetMap() map[interface{}]interface{} {
	m := map[interface{}]interface{}{}
	r.data.Range(func(k, v interface{}) bool {
		m[k] = v
		return true
	})
	return m
}
