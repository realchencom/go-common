package realsmile

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func Clone[N any](new *N, old interface{}) {
	if *new == nil {
		panic("new is nil")
		return
	}
	if old == nil {
		panic("old is nil")
		return
	}
	nType := reflect.TypeOf(*new)
	nValue := reflect.ValueOf(new)
	oValue := reflect.ValueOf(old)
	for i := 0; i < nType.NumField(); i++ {
		field := nType.Field(i)
		name := field.Name

		oField, isExists := oValue.Type().FieldByName(name)
		if isExists {
			if nValue.Elem().FieldByName(name).CanSet() {
				if oField.Type.Kind().String() == field.Type.Kind().String() {
					nValue.Elem().FieldByName(name).Set(oValue.FieldByName(name))
				} else {
					switch oField.Type.Kind().String() {
					case "string":
						{
							if val, err := strconv.ParseInt(oValue.FieldByName(name).String(), 10, 64); err != nil {
								Log.Errorf("Error parsing oValue for field. %v: %v", name, err)
								continue
							} else {
								nValue.Elem().FieldByName(name).SetInt(val)
							}
							break
						}
					case "int64":
						{
							nValue.Elem().FieldByName(name).SetString(strconv.FormatInt(oValue.FieldByName(name).Int(), 10))
							break
						}
					}
				}

			}
		} else {
			continue
		}
	}
}

func WritePidFile(Pid string) {
	//创建文件，返回两个值，一是创建的文件，二是错误信息
	var err error
	var file *os.File
	file, err = os.Create("./pid")
	if err != nil { // 如果有错误，打印错误，同时返回
		fmt.Println("创建pid文件错误。err = ", err)
	}
	_, err = file.WriteString(Pid)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("关闭pid文件错误。err = ", err)
		}
	}(file) // 在退出整个函数时，关闭文件
}
func DeletePidFile() {
	err := os.Remove("./pid")
	if err != nil {
		fmt.Println("删除pid文件错误。err = ", err)
	}
}
