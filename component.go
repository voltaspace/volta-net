package voltanet

import (
	"errors"
	"reflect"
)


func init(){
	bind("wsEvents",&Events{})
}

func Register() {
	if err := serviceComponent(); err != nil {
		panic(err)
	}
	return
}


func serviceComponent() (err error){
	rfType := reflect.TypeOf(&AutoWaired).Elem()
	proType := reflect.ValueOf(&AutoWaired).Elem()
	kd := proType.Kind()
	if kd != reflect.Struct {
		err = errors.New("expect volta struct")
		return err
	}
	num := proType.NumField()
	//.快捷注册
	for i := 0; i < num; i++ {
		tag := rfType.Field(i).Tag.Get("volta")
		service, err := making(tag)
		if err != nil {
			err = errors.New("Not Fount volta Instance " + tag + ";error in gateway/Componet struct tag")
			return err
		}
		proType.Field(i).Set(reflect.ValueOf(service))
	}
	return nil
}
