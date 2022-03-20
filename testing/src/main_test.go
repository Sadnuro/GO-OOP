package main

import "testing"

func TestSum(t *testing.T) {
	total := Sum(6, 5)
	expected := 11

	if total != expected {
		t.Errorf("Sum was incorrect!, got %d expected %d", total, expected)
	}

	/*
		tables := []struct {
			a int
			b int
			n int
		}{
			{1, 2, 3},
			{4, 5, 9,
			{7, 8, 15},
		}

		fmt.Println(tables)

		for _, item := range tables{
			total := Sum(item.a, item.b)

			if total != item.n {
				t.Errorf("Sum was incorrect!, got %d expected %d", total, item.n)
			}
		}
	*/

}
