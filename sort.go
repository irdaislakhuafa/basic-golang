package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

// create slice of user
type UserSortByName []User
type UserSortByAge []User

func main() {
	users := []User{
		{
			Name: "irda", Age: 20,
		},
		{
			Name: "ana", Age: 21,
		},
		{
			Name: "bna", Age: 21,
		},
	}

	fmt.Println("Before sort :", users)
	sort.Sort(UserSortByName(users))
	fmt.Println("After sort (Name) :", users)
	sort.Sort(UserSortByAge(users))
	fmt.Println("After sort (Age) :", users)
}

// by name
func (us UserSortByName) Len() int {
	return len(us)
}
func (us UserSortByName) Less(i, j int) bool {
	return us[i].Name < us[j].Name
}
func (us UserSortByName) Swap(i, j int) {
	// reverse in golang
	us[i], us[j] = us[j], us[i]
}

// by age
func (us UserSortByAge) Len() int {
	return len(us)
}
func (us UserSortByAge) Less(i, j int) bool {
	return us[i].Age < us[j].Age
}
func (us UserSortByAge) Swap(i, j int) {
	// reverse in golang
	us[i], us[j] = us[j], us[i]
}

/*
PACKAGE SORT :
- adalah package yang berisikan utilitas untuk pengurutan
- agar data bisa diurutkan maka kita harus mengimplementasikan kontrak interface sort.Interface
- docs : https://golang.org/pkg/sort
*/
