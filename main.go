package main

import (
	"fmt"
	"strconv"
)

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

func MunculSekali(num string) []int {
	counts := make([]int, 10) // Membuat slice dengan 10 elemen, satu untuk setiap digit (0-9)

	for j := 0; j < len(num); j++ {
		digit, err := strconv.Atoi(string(num[j])) // Mengubah karakter ke dalam bilangan
		if err == nil {
			counts[digit]++
		}
	}

	uniqueDigits := []int{}
	for i := 0; i < 10; i++ {
		if counts[i] == 1 {
			uniqueDigits = append(uniqueDigits, i)
		}
	}

	return uniqueDigits
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

	fmt.Println(MunculSekali("1234123"))
	fmt.Println(MunculSekali("67523752"))
	fmt.Println(MunculSekali("12345"))
	fmt.Println(MunculSekali("1122334455"))
	fmt.Println(MunculSekali("0872504"))
}