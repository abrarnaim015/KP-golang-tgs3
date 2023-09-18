package main

import "fmt"

func ArrayMerge(arrA, ArrB []string) []string {
	arrA = append(arrA, ArrB...)
	return arrA
}

func Mapping(slice []string) map[string]int {
	counts := make(map[string]int)

	for _, s := range slice {
			counts[s]++
	}

	return counts
}

func main()  {
	fmt.Println(ArrayMerge([]string{
		"king",
		"devil jin",
		"akuma",
	}, []string{
		"eddie",
		"steve",
		"geese",
	}))

	fmt.Println(ArrayMerge([]string{
		"sergei",
		"jim",
	}, []string{
		"jin",
		"steve",
		"bryan",
	}))

	fmt.Println(ArrayMerge([]string{
		"alisa",
		"yoshimitsu",
	}, []string{
		"devil jin",
		"yoshimitsu",
		"alisa",
		"law",
	}))

	fmt.Println(ArrayMerge([]string{},
		[]string{
		"devil jin",
		"sergei",
	}))

	fmt.Println(ArrayMerge([]string{
		"hwoarang",
	},
		[]string{}))

	fmt.Println(ArrayMerge([]string{}, []string{}))

	fmt.Println(Mapping([]string{
		"asd",
		"qwe",
		"asd",
		"adi",
		"qwe",
		"qwe",
	}))
	fmt.Println(Mapping([]string{}))
}