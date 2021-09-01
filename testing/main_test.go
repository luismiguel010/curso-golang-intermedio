package main

import "testing"

func TestSum(t *testing.T) {
	/* total := Sum(5, 5)

	if total != 10 {
		t.Errorf("Sum was incorrect, got %d exprected %d", total, 10)
	} */
	tables := []struct {
		a int
		b int
		n int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{25, 26, 51},
	}

	for _, item := range tables {
		total := Sum(item.a, item.b)
		if total != item.n {
			t.Errorf("Sum was incorrect, got %d exprected %d", total, item.n)
		}
	}
}

func TestMax(t *testing.T) {
	tables:= []struct{
		a int
		b int
		n int
	} {
		{4, 2, 4},
		{3, 2, 3},
		{1, 2, 2},
	}

	for _, item := range tables {
		max := GetMax(item.a, item.b)
		if max != item.n {
			t.Errorf("GetMax was incorrect, got %d, expected %d", max, item.n)
		}
	}
}