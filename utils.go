package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Obj struct {
	name string
	age  int
}

func sortElements() {
	// sorting numbers
	arr := []int{4, 5, 1, 2, 3, 6, 8, 9, 7}
	sort.Ints(arr)
	// slice must be sorted in ascending order
	fmt.Println(sort.SearchInts(arr, 2))
	fmt.Println(arr)

	// sorting strings
	str := []string{"bharath", "manoj"}
	sort.Slice(str, func(i int, j int) bool {
		return str[i] > str[j]
	})
	fmt.Println(str)

	// sorting array of objects
	objArr := []Obj{{name: "bahrath", age: 10}, {name: "manoj", age: 20}}
	sort.Slice(objArr, func(i int, j int) bool {
		return objArr[i].age > objArr[j].age
	})
	fmt.Println(objArr)

}

func strMethods() {
	str := "hi hello"
	fmt.Println(strings.Index(str, "hiii"))
	fmt.Println(strings.Split(str, " "))
	fmt.Println(strings.ReplaceAll(str, "hi", "hello"))
	fmt.Println(strings.Join([]string{"hi", "hello"}, ","))
	fmt.Println(str[2:3])
}

func typeConvert() {
	val := "true"
	newVal, _ := strconv.ParseBool(val)
	fmt.Println(newVal)
}

func arrMethods() {
	arr1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr1[0:2])
	arr1 = append(arr1, 5)
	fmt.Println(arr1)

	arr2 := []int{7, 8}
	arr3 := append(arr1, arr2...)
	fmt.Println(arr3)

	arr4 := make([]int, 5)
	copy(arr4, arr2)
	fmt.Println(arr4)

	arr5 := append([]int{}, arr1...)
	arr5[0] = 100
	fmt.Println(arr1)
}

// find element by index
func findEleByIndex() {
	list := []int{1, 2, 3, 4}
	fmt.Print(list[0])
}

// find index by value
func findIndexByVal() {
	list := []int{1, 2, 3, 4}
	ele := 3
	for i, v := range list {
		if v == ele {
			fmt.Println(i)
		}
	}
}

// add new element to the specified pos
func addItems() {
	list := []int{1, 2, 3, 4}
	pos := 2
	ele := []int{9, 8, 10}

	for _, v := range ele {
		list = append(list[:pos+1], list[pos:]...)
		list[pos] = v
		pos++
	}

	fmt.Println(list)
}

// remove an element from specified pos
func removeItems() {
	list := []int{1, 2, 3, 4}
	pos := 2
	list = append(list[:pos], list[pos+1:]...)

	fmt.Println(list)
}

func getTime() {
	fmt.Println(time.Now().Local().UnixMilli())
}

func genUUID() {
	id := uuid.New()
	fmt.Println(id.String())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	rand.Seed(time.Now().UnixNano())
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)

	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func utils() {
	fmt.Println(RandStringRunes(10))
}
