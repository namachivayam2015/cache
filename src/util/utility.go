package util

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

var cacheMap sync.Map
var m  = make(map[string]time.Time)
var Size int
var internalSize int = 0
var ReadThresholdInSec int;

func AddElement(key string, value string) error {
	if Size == internalSize || internalSize > Size {
		return &CustomException{ "You have reached the max size of the cache, please remove unwanted keys & try again" }
	}
	key = strings.Trim(key, " ")
	cacheMap.Store(key, value)
	fmt.Println("Key ", key, " Saved into Cache")
	internalSize++
	m[key] = time.Now()
	return nil
}

func GetElement(key string) string {
	key = strings.Trim(key, " ")
	result, ok  := cacheMap.Load(key)
	val := ""
	if ok {
		fmt.Println("Key ", key, " has ", result.(string), " Value in Cache")
		val = result.(string)
		m[key] = time.Now()
	}
	return val
}

func GetAllElement() map[string]string {
	var tempMap = make(map[string]string)
	cacheMap.Range(func(k, v interface{}) bool {
		tempMap[k.(string)] = v.(string)
		return true
	})
	return tempMap
}

func UpdateElement(key string, value string) {
	key = strings.Trim(key, " ")
	cacheMap.Store(key, value)
	fmt.Println("Key ", key, " update with ", value, " in Cache")
	m[key] = time.Now()
}

func DeleteElement(key string) {
	key = strings.Trim(key, " ")
	cacheMap.Delete(key)
	fmt.Println("Key ", key, " removed from Cache")
	internalSize--
	delete(m, key)
}

func GarbageCollector(){
	fmt.Println("start "+time.Now().Format("2006-01-02 15:04:05"))
	for k, v := range m {
		diff := time.Now().Sub(v)
		if int(diff.Milliseconds()) > ReadThresholdInSec*1000 {
			cacheMap.Delete(k)
			delete(m, k)
			fmt.Println(k, "removed from cache, due to inactivity")
		}
	}
	fmt.Println("end "+time.Now().Format("2006-01-02 15:04:05"))
}