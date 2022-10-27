package realsmile

import (
	"fmt"
	"os"
	"reflect"
)

func Clone[T any](new *T, old *interface{}) {
	oValue := reflect.ValueOf(old)
	nValue := reflect.ValueOf(new)
	for i := 0; i < nValue.NumField(); i++ {
		nType := nValue.Type()
		name := nType.Field(i).Name
		_, isExists := oValue.Type().FieldByName(name)
		if isExists {
			nValue.Field(i).Set(oValue.FieldByName(name))
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
