package common

//for debug purpose
import (
	"fmt"
	"reflect"
)

func Beautify(data interface{}, level int) {
	for i := 0; i < level; i++ {
		fmt.Print(" ")
	}
	f := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	fmt.Println("Calling field")
	for i := 0; i < f.NumField(); i++ {
		kind := f.Field(i).Type.Kind()
		fmt.Println(kind)
		exported := f.Field(i).IsExported()
		if !exported {
			continue
		}
		if kind == reflect.Slice {
			fmt.Printf("%v: \n", f.Field(i).Name)
			for j := 0; j < v.Field(i).Len(); j++ {
				// fmt.Printf("%+v\n", v.Field(i).Index(j))
				// fmt.Println(j)
				fmt.Println("Calling slice")
				Beautify(v.Field(i).Index(j).Convert(v.Field(i).Type().Elem()), level+1)
			}
			continue
		}
		if kind == reflect.Pointer {
			fmt.Println(v.Elem())
			continue
		}
		if kind == reflect.Struct {
			// fmt.Println("sdfsdfsfaewe")
			fmt.Printf("%v: \n", f.Field(i).Name)
			Beautify(v.Field(i).Convert(f.Field(i).Type), level+1)
			continue
		}
		if kind == reflect.Map {
			continue
		}
		fmt.Printf("%v: %v\n", f.Field(i).Name, v.Field(i).Interface())
	}
}
func BeautifyReflection(data reflect.Value, level int) {
	// for i := 0; i < level; i++ {
	// 	fmt.Print(" ")
	// }
	// type_ := data.Type()
	// type_:=data.Func
	for i := 0; i < data.NumField(); i++ {
		kind := data.Field(i).Type().Kind()
		exported := data.Type().Field(i).IsExported()
		if !exported {
			continue
		}
		switch kind {
		case reflect.Slice:
			for i := 0; i < level; i++ {
				fmt.Print(" ")
			}
			fmt.Printf("%v: \n", data.Type().Field(i).Name)
			for j := 0; j < data.Field(i).Len(); j++ {
				fmt.Println("-")
				elem := data.Field(i).Index(j)
				BeautifyReflection(elem, level+1)
				fmt.Println("-")
			}
			// SliceReflection(data.Field(i))
			continue
		case reflect.Pointer:
			continue
		case reflect.Struct:
			if res, ok := data.Field(i).Interface().(fmt.Stringer); ok {
				for i := 0; i < level; i++ {
					fmt.Print(" ")
				}
				fmt.Printf("%v: %v \n", data.Type().Field(i).Name, res)
				continue
			}
			fmt.Printf("%v: \n", data.Type().Field(i).Name)
			BeautifyReflection(data.Field(i).Convert(data.Type().Field(i).Type), level+1)
			continue
		case reflect.Map:
			continue
		default:
			for i := 0; i < level; i++ {
				fmt.Print(" ")
			}
			fmt.Printf("%v: %v\n", data.Type().Field(i).Name, data.Field(i).Interface())
			continue
		}
	}
}
func SliceReflection(data reflect.Value) {
	if data.Type().Kind() != reflect.Slice {
		return
	}
	for i := 0; i < data.Len(); i++ {
		elem := data.Index(i)
		BeautifyReflection(elem, 0)
	}
}
