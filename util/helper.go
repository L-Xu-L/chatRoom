package util

import (
	"math/rand"
	"reflect"
	"time"
)

/**
	助手函数
 */

/**
	结构体转map
 */
func StructToMap(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		field := []byte(t.Field(i).Name)
		field[0] += 32
		data[string(field)] = v.Field(i).Interface()
	}
	return data
}

/**
	获取Number位时间戳
 */
func GetRandom(number int) string {
	rawStr := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLNMOPQRSTUVWXYZ1234567890"
	var randSlice []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i:=0; i<number; i++ {
		index := r.Intn(len(rawStr))
		randSlice = append(randSlice,rawStr[index])
	}
	return string(randSlice)
}