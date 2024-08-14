package di

import "reflect"

var instances = map[reflect.Type]any{}

func Bind(newInterface any, instance any) {
	instances[reflect.TypeOf(newInterface)] = instance
}

func Fetch[T any]() T {
	if instances == nil {
		instances = map[reflect.Type]any{}
	}

	newInterface := new(T)
	return instances[reflect.TypeOf(newInterface)].(T)
}
