package main

import (
	"errors"
	"fmt"
)

var People map[string]int

var ErrNotFound error = errors.New("name not found")

func MapInit() {
	People = map[string]int{}
	People["Andy"] = 24
	People["Boone"] = 32
	People["James"] = 44
}

func GetAge(name string) (int, error) {
	age, exists := People[name]
	if exists {
		return age, nil
	}
	return 0, ErrNotFound
}

func DeletePerson(name string) error {
	_, exists := People[name]
	if exists {
		delete(People, name)
	}
	return ErrNotFound
}

func PrintAll() {
	for key, val := range People {
		fmt.Printf("%s, is %d years old\n", key, val)
	}
}

func main() {
	MapInit()

	age, _ := GetAge("Andy")
	fmt.Println(age)
	DeletePerson("Andy")
	age, err := GetAge("Andy")
	if err != nil {
		fmt.Println(err.Error())
	}

	PrintAll()
}
