package gaz

import r "reflect"
import "strconv"
//import "dump"
/*
const mapping = map[string]string {
	"VARCHAR": "string",
	"TINYINT": "byte",
	"SMALLINT": "int16",
	"INT": "int",
	"FLOAT": "float32",
	"DOUBLE": "float64"
}
*/
func maptype(data map[string]interface{}) (result map[string]string) {
	result = make(map[string]string)
	v := r.ValueOf(data)
	for _, k := range v.MapKeys() {
		switch r.ValueOf(v.MapIndex(k).Interface()).Kind() {
		case r.String : 
			result[k.Interface().(string)] = "'"+v.MapIndex(k).Interface().(string)+"'"
		case r.Int :
			result[k.Interface().(string)] = strconv.Itoa(v.MapIndex(k).Interface().(int))
		case r.Float64 : 
			result[k.Interface().(string)] = strconv.Ftoa64(v.MapIndex(k).Interface().(float64), 'f', -1)
		}
	}
	return
}
