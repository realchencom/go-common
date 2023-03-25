package realsmile

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
	"reflect"
	"strconv"
)

// 定义普通类型，非结构体

var baseTypes = map[reflect.Kind]bool{
	reflect.Bool:    true,
	reflect.Int:     true,
	reflect.Int8:    true,
	reflect.Int16:   true,
	reflect.Int32:   true,
	reflect.Int64:   true,
	reflect.Uint:    true,
	reflect.Uint8:   true,
	reflect.Uint16:  true,
	reflect.Uint32:  true,
	reflect.Uint64:  true,
	reflect.Uintptr: true,
	reflect.Float32: true,
	reflect.Float64: true,
	reflect.Map:     true,
	reflect.String:  true,
}

func Clone(target interface{}, source interface{}) error {
	targetValue := reflect.ValueOf(target)
	targetType := reflect.TypeOf(target)
	if targetType.Kind() != reflect.Ptr {
		return errors.New("转换失败，目标结构体不是指针")
	}
	sourceValue := reflect.ValueOf(source)
	sourceType := reflect.TypeOf(source)
	for targetType.Kind() == reflect.Ptr {
		targetType = targetType.Elem()
	}
	for targetValue.Kind() == reflect.Ptr {
		targetValue = targetValue.Elem()
	}

	for sourceValue.Kind() == reflect.Ptr {
		sourceValue = sourceValue.Elem()
	}
	for sourceType.Kind() == reflect.Ptr {
		sourceType = sourceType.Elem()
	}
	if sourceType.Kind() != reflect.Struct || targetType.Kind() != reflect.Struct {
		return errors.New("转换失败，只支持转结构体")
	}
	clone(targetValue, sourceValue, targetType, sourceType)
	return nil
}

func clone(targetValue, sourceValue reflect.Value, targetType, sourceType reflect.Type) {

	for i := 0; i < targetType.NumField(); i++ {
		targetField := targetType.Field(i)
		if baseTypes[targetField.Type.Kind()] {
			if v, _, b := sourceValueType(sourceValue, sourceType, targetField.Name); b {
				setValue(targetValue.Field(i), *v, targetField.Name)
			}
			continue
		} else if targetField.Type.Kind() == reflect.Struct {
			clone(targetValue.Field(i), sourceValue, targetField.Type, sourceType)
		}
	}
}
func sourceValueType(sourceValue reflect.Value, sourceType reflect.Type, field string) (*reflect.Value, reflect.Type, bool) {
	for i := 0; i < sourceType.NumField(); i++ {
		structField := sourceType.Field(i)
		if baseTypes[structField.Type.Kind()] && field == structField.Name {
			fieldValue := sourceValue.Field(i)
			return &fieldValue, structField.Type, true
		} else if structField.Type.Kind() == reflect.Struct {
			if val, t, b := sourceValueType(sourceValue.Field(i), structField.Type, field); b {
				return val, t, b
			}
		}
	}
	return nil, nil, false
}
func setValue(targetValue reflect.Value, sourceValue reflect.Value, name string) {

	if targetValue.CanSet() {
		if sourceValue.Type().Kind() == targetValue.Type().Kind() {
			targetValue.Set(sourceValue)
		} else {
			switch sourceValue.Type().Kind() {
			case reflect.String:
				{
					if val, err := strconv.ParseInt(sourceValue.String(), 10, 64); err != nil {
						Log.Errorf("Error parsing oValue for field. %v: %v", name, err)
						break
					} else {
						targetValue.SetInt(val)
					}
					break
				}
			case reflect.Int64:
				{
					targetValue.SetString(strconv.FormatInt(sourceValue.Int(), 10))
					break
				}
			}
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
