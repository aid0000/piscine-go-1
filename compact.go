package piscine

import ()

func Compact(ptr *[]string) int {

	add := 0
	var arr []string
	for _, v := range *ptr {

		if v != "" {
			arr=append(arr,v)
			add++	
		}
	}
	*ptr=arr
	return add
}
