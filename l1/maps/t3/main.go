package main

import (
	"fmt"
)

// FilterByValue возвращает новую map, содержащую только элементы,
// значения которых присутствуют в allowedValues.
func FilterByValue(m map[int]string, allowedValues []string) map[int]string {
	// Преобразовать allowedValues в set для быстрой проверки
	// Создать новую map и заполнить её подходящими элементами

	filteredMap := make(map[int]string)
	allowedValuesSet := make(map[string]bool)

	for _, v := range allowedValues {
		allowedValuesSet[v] = true
	}

	for k, v := range m {
		_, isAllowed := allowedValuesSet[v]
		if isAllowed {
			filteredMap[k] = v
		}
	}

	return filteredMap
}

// InvertMap меняет ключи и значения местами.
// Если значения исходной map не уникальны, возвращает ошибку.
func InvertMap(m map[string]int) (map[int]string, error) {
	// Проверять уникальность значений
	// При обнаружении дубликата вернуть ошибку с описанием конфликта

	invertedMap := make(map[int]string)
	for k, v := range m {
		_, double := invertedMap[v]
		if double {
			return nil, fmt.Errorf("value %d was repeaded", v)
		}
		invertedMap[v] = k
	}

	return invertedMap, nil
}

func main() {
	testMap := make(map[string]int)

	testMap["word"] = 4
	testMap["que"] = 3
	testMap["slice"] = 5
	testMap["food"] = 4

	_, err := InvertMap(testMap)
	if err == nil {
		panic("actually i expected here an error")
	}
	delete(testMap, "food")

	inverted, err := InvertMap(testMap)
	if err != nil {
		panic("here should be no errors, though")
	}

	filtered := FilterByValue(inverted, []string{"word", "slice"})

	for k, v := range filtered {
		fmt.Printf("%d, %s\n", k, v)
	}
}
