package common

import "reflect"

const tagName = "gf_test"

type pathMap struct {
	ma reflect.Value
	key string
	value reflect.Value
	path string
}

func NewDecoder() {

}