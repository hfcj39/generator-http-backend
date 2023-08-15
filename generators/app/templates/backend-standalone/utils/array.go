package utils

import (
	"math/rand"
	"reflect"
	"sort"
	"time"
)

// ItemInStructArray 判断元素是否在结构体数组中
func ItemInStructArray(item interface{}, array interface{}) bool {
	sVal := reflect.ValueOf(array)
	kind := sVal.Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		for i := 0; i < sVal.Len(); i++ {
			if reflect.DeepEqual(item, sVal.Index(i).Interface()) {
				return true
			}
		}
	}
	return false
}

func AttrInStructArray(attr_value interface{}, array interface{}, attr_name string) bool {
	sVal := reflect.ValueOf(array)
	kind := sVal.Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		for i := 0; i < sVal.Len(); i++ {
			var val interface{}
			switch sVal.Index(i).Kind() {
			case reflect.Ptr:
				val = sVal.Index(i).Elem().FieldByName(attr_name).Interface()
			case reflect.Struct:
				val = sVal.Index(i).FieldByName(attr_name).Interface()
			default:
				val = sVal.Index(i).FieldByName(attr_name).Interface()
			}
			if attr_value == val {
				return true
			}
		}
	}

	return false
}

func RemoveDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func FindItemInSlice(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func RandRangeNumbers(min, max, count int) []int {
	// 检查参数
	if max < min || (max-min+1) < count {
		return nil
	}
	nums := make([]int, max-min+1)
	position := -1
	if min <= 0 && max >= 0 {
		position = -min
		nums[position] = max + 1
	}
	// 随机数生成器
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < count; i++ {
		num := r.Intn(max - min + 1)
		if nums[i] == 0 {
			nums[i] = min + i
		}
		if nums[num] == 0 {
			nums[num] = min + num
		}
		if position != -1 {
			if i == position {
				position = num
			} else if num == position {
				position = i
			}
		}
		nums[i], nums[num] = nums[num], nums[i]
	}

	if position != -1 {
		nums[position] = 0
	}
	rst := nums[0:count]
	sort.Ints(rst)
	return rst
}
