package main

import (
	"fmt"
	"reflect"
)

type Kisi struct {
	ad    string "Kisi adı"
	soyad string
	yas   int "Kisi yas"
}

func main() {
	yapi3 := Kisi{"Ayşe", "Yıldız", 52}
	if reflect.ValueOf(yapi3).Kind() == reflect.Struct {
		v := reflect.ValueOf(yapi3)
		fmt.Println("Alan sayısı", v.NumField())
		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("Alan:%d  Değer:%v Etiket: %v\n", i, v.Field(i), reflect.TypeOf(yapi3).Field(i).Tag)
		}
	}
}
