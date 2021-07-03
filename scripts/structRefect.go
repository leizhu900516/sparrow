package main

import (
	"fmt"
	"log"
	"reflect"
)

type Person struct {
	Name string
	Age int
}
func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		log.Println("Check type error not Struct")
		return nil
	}
	var data = make(map[string]interface{})
	fmt.Println(">>>",t.NumField())
	for i := 0; i < t.NumField(); i++ {
		fmt.Println("name=",t.Field(i).Name)
		fmt.Println("interface=",v.Field(i))
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}
func main()  {
	tom := Person{Name:"chen",Age:14}
	fmt.Println(tom)
	data:=Struct2Map(tom)
	fmt.Println(data)
}
