package main

import (
	"reflect"
	"fmt"
)

func StructToMap(s interface{}, tagName string) map[string]interface{} {
	val := reflect.ValueOf(s)
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return nil
		}
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return nil
	}

	typ := val.Type()
	r := make(map[string]interface{})
	for i := 0; i < typ.NumField(); i++ {
		vf := val.Field(i)
		// 判断是否是导出的字段
		if vf.CanInterface() {
			key := typ.Field(i).Name
			if tagName != "" {
				tag := typ.Field(i).Tag.Get(tagName)
				if tag != "" {
					key = tag
				}
			}
			r[key] = vf.Interface()
		}
	}
	return r
}

func main() {
	type ts struct {
		A string
		B int
		C interface{}
		D interface{}
	}
	sm := &ts{
		A: "a", B: 1, C: nil,
	}
	fmt.Println(StructToMap(sm, "json"))
	//t.Log(StructSliceToMaps(sm, "", false))
	sm2 := ts{
		A: "a", B: 1, C: *sm,
	}
	fmt.Println(StructToMap(sm2, "json"))
	fmt.Println(sm2.C)
	tableNum :=50
	format := fmt.Sprintf("%%s_%%0%dd", IntNumDigit(tableNum-1))
	i:=uint64(4200152524)
	baseTable:= "inews_aa"
	a:=fmt.Sprintf(format, baseTable, i%uint64(tableNum))
	fmt.Println(a)
	//t.Log(StructSliceToMaps(sm2, "", false))
}

// IntNumDigit 获取输入数字是几位数
func IntNumDigit(n int) int {
	return Int64NumDigit(int64(n))
}
// Int64NumDigit 获取输入数字是几位数
func Int64NumDigit(n int64) int {
	for i := 0; ; {
		i++
		t := n / 10
		if t == 0 {
			return i
		}
		n = t
	}
}