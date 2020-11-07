package util

import (
	"fmt"
	"reflect"
	"strconv"
)

// ConvertMapToStruct
func ConvertMapToStruct(params map[string]interface{}, val interface{}) error {

	var err error

	//ValueOf返回接口指代的具体值,Elem返回v持有的接口保管的值的Value封装，或者v持有的指针指向的值的Value封装
	structVal := reflect.ValueOf(val).Elem()
	//value的Type类型
	structType := structVal.Type()
	//N通过NumField获取结构体有几个字段
	for i := 0; i < structVal.NumField(); i++ {
		//通过Field获取结构体字段
		fieldType := structType.Field(i)
		//通过FieldByName获取结构体字段
		field := structVal.FieldByName(fieldType.Name)
		//字段的tag名称
		name := fieldType.Tag.Get("tag")
		if name == "" {
			continue
		}
		ignore := fieldType.Tag.Get("ignore")
		def := fieldType.Tag.Get("default")
		if ignore == "true" {
			continue
		}
		data, ok := params[name]
		if !ok {
			if def == "" {
				return err
			} else {
				data = def
			}
		}
		////将data根据val的底层类型进行转换
		err := convertStringToVal(data, field)
		if err != nil {
			return err
		}
	}
	return nil
}

//将data根据val的底层类型进行转换
func convertStringToVal(data interface{}, val reflect.Value) error {

	var err error

	//获取底层类型
	kind := getKind(val)
	switch kind {
	case reflect.String:
		//设置val的值为TransToString(data)
		val.SetString(TransToString(data))
		return nil
	case reflect.Int:
		//返回字符串表示的整数值
		//base代表进制
		//bitSize代表int，int32，int64
		num, err := strconv.ParseInt(TransToString(data), 10, 64)
		if err != nil {
			return err
		}
		val.SetInt(num)
		return nil
	case reflect.Float64:
		num, err := strconv.ParseFloat(TransToString(data), 64)
		if err != nil {
			return err
		}
		val.SetFloat(num)
		return nil
	case reflect.Bool:
		b, err := strconv.ParseBool(TransToString(data))
		if err != nil {
			return err
		}
		val.SetBool(b)
		return nil
	default:
		return err
	}
}

// TransToString 强制类型转换为 string
func TransToString(data interface{}) (res string) {
	switch v := data.(type) {
	case bool:
		res = strconv.FormatBool(v)
	case float32:
		res = strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		res = strconv.FormatFloat(v, 'f', -1, 64)
	case int, int8, int16, int32, int64:
		val := reflect.ValueOf(data)
		res = strconv.FormatInt(val.Int(), 10)
	case uint, uint8, uint16, uint32, uint64:
		val := reflect.ValueOf(data)
		res = strconv.FormatUint(val.Uint(), 10)
	case string:
		res = v
	case []byte:
		res = string(v)
	default:
		res = fmt.Sprintf("%v", v)
	}
	return
}

//得到底层类型
func getKind(val reflect.Value) reflect.Kind {
	kind := val.Kind()

	switch {
	case kind >= reflect.Int && kind <= reflect.Int64:
		return reflect.Int
	case kind >= reflect.Uint && kind <= reflect.Uint64:
		return reflect.Uint
	case kind >= reflect.Float32 && kind <= reflect.Float64:
		return reflect.Float64
	default:
		return kind
	}
}
