package LogicFunc

import (
	"fmt"
	"reflect"
)

/*
	实现了传入一个变量指针(该变量及其下层字段都不是interface{})，则将其中为切片类型的字段设置为空切片
	Array
	Chan
	Func
	Interface	×（不支持）
	Map			√
	Ptr			√
	Slice		√
	Struct		√
	UnsafePointer
*/
func SetEmptySlice(Value reflect.Value) {

	if Value.Kind() != reflect.Ptr {
		return
	}
	Value = Value.Elem()
	fmt.Println(Value.Interface())
	//fmt.Println(Value.Kind())

	// 遍历结构体的字段
	if Value.Kind() == reflect.Struct {
		for i := 0; i < Value.NumField(); i++ {
			if Value.Field(i).CanAddr(){
				SetEmptySlice(Value.Field(i).Addr())
			}
		}
	}

	// 遍历切片的值
	if Value.Kind() == reflect.Slice {
		if Value.IsNil() {
			sliceValue := reflect.MakeSlice(Value.Type(), 0, 0)
			if Value.CanSet() {
				Value.Set(sliceValue)
			}
		} else {
			for i := 0; i < Value.Len(); i++ {
				SetEmptySlice(Value.Index(i).Addr())
			}
		}
	}

	// 遍历map
	if Value.Kind() == reflect.Map {
		for _, key := range Value.MapKeys() {
			value := Value.MapIndex(key)
			// 创建指针类型的反射值
			valuePtr := reflect.New(value.Type())
			if valuePtr.Elem().CanSet() {
				valuePtr.Elem().Set(value)
				SetEmptySlice(valuePtr)
				if Value.CanSet() {
					Value.SetMapIndex(key, valuePtr.Elem())
				}
			}

		}
	}

	if Value.Kind() == reflect.Ptr {
		SetEmptySlice(Value.Elem().Addr())
	}
}
