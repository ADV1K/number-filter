package main

import (
	"flag"
	"fmt"
	"strconv"
)

const HelpText = `filter - A CLI tool for filtering numbers

Usage:
	filter [options] [numbers...]

Options:`

// Help prints the help message
func Help() {
	fmt.Println(HelpText)
	flag.VisitAll(func(flag *flag.Flag) {
		fmt.Printf("\t--%s\t\t%s\n", flag.Name, flag.Usage)
	})
}

// IsEven returns true if the number is even, false otherwise
func IsEven(num int) bool {
	return num%2 == 0
}

// IsOdd returns true if the number is odd, false otherwise
func IsOdd(num int) bool {
	return num%2 != 0
}

// IsPrime returns true if the number is prime, false otherwise
func IsPrime(num int) bool {
	if num == 2 || num == 3 {
		return true
	}

	if num < 2 || num%2 == 0 || num%3 == 0 {
		return false
	}

	i := 5
	for i*i <= num {
		if num%i == 0 || num%(i+2) == 0 {
			return false
		}
		i += 4
	}
	return true
}

// IsMultipleOf returns true if the number is a multiple of the given number, false otherwise
func IsMultipleOf(num, multiple int) bool {
	return num%multiple == 0
}

// IsGreater returns true if the number is greater than the given number, false otherwise
func IsGreater(num, greater int) bool {
	return num > greater
}

// IsLess returns true if the number is less than the given number, false otherwise
func IsLess(num, less int) bool {
	return num < less
}

// IsEqual returns true if the number is equal to the given number, false otherwise
func IsEqual(num, equal int) bool {
	return num == equal
}

// IsGreaterOrEqual returns true if the number is greater than or equal to the given number, false otherwise
func IsGreaterOrEqual(num, greaterOrEqual int) bool {
	return num >= greaterOrEqual
}

// IsLessOrEqual returns true if the number is less than or equal to the given number, false otherwise
func IsLessOrEqual(num, lessOrEqual int) bool {
	return num <= lessOrEqual
}

// PartialFilter is a curried function that returns a filter function
func PartialFilter(filter func(int, int) bool, other int) func(int) bool {
	return func(num int) bool {
		return filter(num, other)
	}
}

// FilterAll returns a slice of numbers that pass all the filters
func FilterAll(nums []int, filters ...func(int) bool) []int {
	var result []int
	for _, num := range nums {
		passesAllFilters := true
		for _, filter := range filters {
			if !filter(num) {
				passesAllFilters = false
				break
			}
		}
		if passesAllFilters {
			result = append(result, num)
		}
	}
	return result
}

// FilterAny returns a slice of numbers that pass any of the filters
func FilterAny(nums []int, filters ...func(int) bool) []int {
	var result []int
	for _, num := range nums {
		passesAnyFilter := false
		for _, filter := range filters {
			if filter(num) {
				passesAnyFilter = true
				break
			}
		}
		if passesAnyFilter {
			result = append(result, num)
		}
	}
	return result
}

// StringsToNumbers converts a slice of strings to a slice of numbers
func StringsToNumbers(args []string) []int {
	var nums []int
	for _, arg := range args {
		num, err := strconv.Atoi(arg)
		if err != nil {
			panic(err)
		}
		nums = append(nums, num)
	}
	return nums
}

// IsFilterPresent returns true if the given filter is present, false otherwise
func IsFilterPresent(filter string) (present bool) {
	flag.Visit(func(flag *flag.Flag) {
		if flag.Name == filter {
			present = true
		}
	})
	return present
}

func main() {
	// Help flag
	var help bool
	flag.BoolVar(&help, "help", false, "Print help message")

	// Filter Flags
	var even, odd, prime bool
	flag.BoolVar(&even, "even", false, "Filter even numbers")
	flag.BoolVar(&odd, "odd", false, "Filter odd numbers")
	flag.BoolVar(&prime, "prime", false, "Filter prime numbers")

	var mult, greater, less, equal, greaterOrEqual, lessOrEqual int
	flag.IntVar(&mult, "mult", 0, "Filter multiples of the given number")
	flag.IntVar(&greater, "gt", 0, "Filter numbers greater than the given number")
	flag.IntVar(&less, "lt", 0, "Filter numbers less than the given number")
	flag.IntVar(&equal, "eq", 0, "Filter numbers equal to the given number")
	flag.IntVar(&greaterOrEqual, "ge", 0, "Filter numbers greater than or equal to the given number")
	flag.IntVar(&lessOrEqual, "le", 0, "Filter numbers less than or equal to the given number")

	// At least one filter must be true
	var anyFilter bool
	flag.BoolVar(&anyFilter, "any", false, "At least one of the filters must be true")

	// Parse flags
	flag.Parse()

	// Print help message if no arguments are given or if help flag is set
	if help || flag.NArg() == 0 {
		Help()
		return
	}

	// Parse filters
	var filters []func(int) bool
	if even {
		filters = append(filters, IsEven)
	}
	if odd {
		filters = append(filters, IsOdd)
	}
	if prime {
		filters = append(filters, IsPrime)
	}
	if IsFilterPresent("mult") {
		filters = append(filters, PartialFilter(IsMultipleOf, mult))
	}
	if IsFilterPresent("gt") {
		filters = append(filters, PartialFilter(IsGreater, greater))
	}
	if IsFilterPresent("lt") {
		filters = append(filters, PartialFilter(IsLess, less))
	}
	if IsFilterPresent("eq") {
		filters = append(filters, PartialFilter(IsEqual, equal))
	}
	if IsFilterPresent("ge") {
		filters = append(filters, PartialFilter(IsGreaterOrEqual, greaterOrEqual))
	}
	if IsFilterPresent("le") {
		filters = append(filters, PartialFilter(IsLessOrEqual, lessOrEqual))
	}

	// Filter numbers
	var result []int
	nums := StringsToNumbers(flag.Args())
	if anyFilter {
		result = FilterAny(nums, filters...)
	} else {
		result = FilterAll(nums, filters...)
	}

	// Print result
	for _, num := range result {
		fmt.Printf("%d ", num)
	}
	fmt.Println()
}
