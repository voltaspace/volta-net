package container

import (
	"errors"
	"github.com/voltaspace/volta-net/show"
	"reflect"
	"runtime"
	"strconv"
)

type Container struct {
	instances map[string]InstancesBind
}

type InstancesBind struct {
	Method string
	ConValue reflect.Value
	ConType reflect.Type
}

var AppKeys []string = make([]string,0)

var App = Container{
	make(map[string]InstancesBind),
}

func (container *Container) Bind(appNm interface{},relay interface{})  {
	var name string
	refV := reflect.ValueOf(relay)
	refT := reflect.Indirect(refV).Type()
	relayNm := refT.Name()
	switch v := appNm.(type) {
		case int:
			name = strconv.FormatInt(int64(v),10)
		case string:
			name = appNm.(string)
	}
	if name == "" {
		return
	}
	show.VoltaPrint(name,runtime.FuncForPC(refV.Pointer()).Name())
	container.instances[name] = InstancesBind{
		relayNm,
		refV,
		refT,
	}
	AppKeys = append(AppKeys,name)
}

func (container *Container) Make(method string) (app InstancesBind,err error){
	if _,ok := container.instances[method]; !ok{
		err = errors.New("reflect nil")
		return InstancesBind{},err
	}
	app = container.instances[method]
	return
}
func (container *Container) GetInstances() map[string]InstancesBind {
	return container.instances
}

