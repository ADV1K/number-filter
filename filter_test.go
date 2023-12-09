package main

import "testing"

func TestIsEven(t *testing.T) {
	cases := []struct {
		Num  int
		Want bool
	}{
		{1, false},
		{2, true},
		{3, false},
		{4, true},
	}

	for _, c := range cases {
		got := IsEven(c.Num)
		if got != c.Want {
			t.Errorf("IsEven(%d) == %t, want %t", c.Num, got, c.Want)
		}
	}
}

func TestIsOdd(t *testing.T) {
	cases := []struct {
		Num  int
		Want bool
	}{
		{1, true},
		{2, false},
		{3, true},
		{4, false},
	}

	for _, c := range cases {
		got := IsOdd(c.Num)
		if got != c.Want {
			t.Errorf("IsOdd(%d) == %t, want %t", c.Num, got, c.Want)
		}
	}
}

func TestIsPrime(t *testing.T) {
	cases := []struct {
		Num  int
		Want bool
	}{
		{1, false},
		{2, true},
		{3, true},
		{5, true},
		{6, false},
		{17, true},
	}

	for _, c := range cases {
		got := IsPrime(c.Num)
		if got != c.Want {
			t.Errorf("IsPrime(%d) == %t, want %t", c.Num, got, c.Want)
		}
	}
}

func TestAll(t *testing.T) {
	cases := []struct {
		Description string
		Nums        []int
		Filters     []func(int) bool
		Want        []int
	}{
		{"Empty", []int{}, []func(int) bool{IsEven}, []int{}},
		{"Test with one filter", []int{1, 2, 3, 4, 5}, []func(int) bool{IsEven}, []int{2, 4}},
		{"Test with two filters", []int{1, 2, 3, 4, 5}, []func(int) bool{IsEven, IsPrime}, []int{2}},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			got := FilterAll(c.Nums, c.Filters...)
			if len(got) != len(c.Want) {
				t.Errorf("All(%v, %v) == %v, want %v", c.Nums, c.Filters, got, c.Want)
			}
			for i := range got {
				if got[i] != c.Want[i] {
					t.Errorf("All(%v, %v) == %v, want %v", c.Nums, c.Filters, got, c.Want)
				}
			}
		})
	}
}

func TestAny(t *testing.T) {
	cases := []struct {
		Description string
		Nums        []int
		Filters     []func(int) bool
		Want        []int
	}{
		{"Empty", []int{}, []func(int) bool{IsEven}, []int{}},
		{"Test with one filter", []int{1, 2, 3, 4, 5}, []func(int) bool{IsEven}, []int{2, 4}},
		{"Test with two filters", []int{1, 2, 3, 4, 5}, []func(int) bool{IsEven, IsPrime}, []int{2, 3, 4, 5}},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			got := FilterAny(c.Nums, c.Filters...)
			if len(got) != len(c.Want) {
				t.Errorf("Any(%v, %v) == %v, want %v", c.Nums, c.Filters, got, c.Want)
			}
			for i := range got {
				if got[i] != c.Want[i] {
					t.Errorf("Any(%v, %v) == %v, want %v", c.Nums, c.Filters, got, c.Want)
				}
			}
		})
	}
}
