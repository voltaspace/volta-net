package voltanet

import "errors"

var componentInstances map[string]interface{} = make(map[string]interface{})

func bind(name string, service interface{}) {
	componentInstances[name] = service
}

func making(name string) (service interface{}, err error) {
	if _, ok := componentInstances[name]; !ok {
		err = errors.New("nil")
		return
	}
	service = componentInstances[name]
	return
}
