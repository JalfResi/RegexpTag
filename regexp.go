package regexptag

import (
	"reflect"
	"regexp"
)

const tagName = "regexp"

func Parse(i interface{}, s string) {
	ptyp := reflect.TypeOf(i)  // a reflect.Type
	pval := reflect.ValueOf(i) // a reflect.Value

	var typ reflect.Type
	var val reflect.Value

	if ptyp.Kind() == reflect.Ptr {
		typ = ptyp.Elem()
		val = pval.Elem()
	} else {
		typ = ptyp
		val = pval
	}

	if typ.Kind() != reflect.Struct {
		return
	}

	for n := 0; n < typ.NumField(); n++ {
		sfld := typ.Field(n)
		tfld := sfld.Type   // The Type of the StructField of the struct
		kind := tfld.Kind() // The Kind of the Type of the StructField
		vfld := val.Field(n)

		tagVal := sfld.Tag.Get(tagName)
		re := regexp.MustCompile(tagVal)
		m := re.FindStringSubmatch(s)

		if kind == reflect.String && vfld.CanSet() {
			vfld.SetString(m[1])
		}

		if kind == reflect.Struct && vfld.CanAddr() && vfld.CanSet() {
			Parse(vfld.Addr().Interface(), s)
		}
	}
}
