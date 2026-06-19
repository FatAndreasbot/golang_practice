package main

import (
	"errors"
	"fmt"
)

var People map[string]int

var ErrNotFound error = errors.New("map not found")

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
	People = map[string]int{}

}
