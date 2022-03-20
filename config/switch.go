package config

import "green-hat/model"

func (r *Result) SwitchWith(types string, files []string) {
	for _, p := range files {
		model.WG.Add(1)
		switch types {
		case "push_type":
			go r.PushChannel(p)
		}
	}
	model.WG.Wait()
}
