package datacenter

import (
	"github.com/centurylinkcloud/clc-go-cli/config"
	"reflect"
)

type SetDefault struct {
	DataCenter string `valid:"required"`
}

func ApplyDefault(model interface{}, conf *config.Config) {
	if model == nil {
		return
	}

	if _, ok := model.(*SetDefault); ok {
		return
	}

	applyDefault(model, conf)
}

func applyDefault(model interface{}, conf *config.Config) {
	yes, set := dependsOnDC(model)
	if yes && !set {
		if conf.DefaultDataCenter != "" {
			setDC(model, conf.DefaultDataCenter)
		}
	}
}

func dependsOnDC(model interface{}) (yes, set bool) {
	if model == nil {
		return
	}

	meta := reflect.ValueOf(model)
	if meta.Kind() == reflect.Ptr {
		meta = meta.Elem()
	}
	if meta.Kind() != reflect.Struct {
		panic("model must be nil or a struct or a pointer to a struct.")
	}

	dc := meta.FieldByName("DataCenter")
	if dc.IsValid() {
		yes = true
		set = dc.String() != ""
	}
	return
}

func setDC(model interface{}, dc string) {
	meta := reflect.ValueOf(model)
	if meta.Kind() == reflect.Ptr {
		meta = meta.Elem()
	}

	meta.FieldByName("DataCenter").SetString(dc)
}
