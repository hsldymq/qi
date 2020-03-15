package component

import "testing"

func TestPalace(t *testing.T) {
	t.Run("test rotate method", func(t *testing.T) {
		original := Palace{
			First:   1,
			Eighth:  2,
			Third:   3,
			Fourth:  4,
			Ninth:   5,
			Second:  6,
			Seventh: 7,
			Sixth:   8,
			Fifth:   9,
		}
		inputs := []int{3, 9, -2, -28}
		expect := []Palace{
			{
				First: 6, Eighth: 7, Third: 8, Fourth: 1,
				Ninth: 2, Second: 3, Seventh: 4, Sixth: 5,
				Fifth: 9,
			},
			{
				First: 8, Eighth: 1, Third: 2, Fourth: 3,
				Ninth: 4, Second: 5, Seventh: 6, Sixth: 7,
				Fifth: 9,
			},
			{
				First: 3, Eighth: 4, Third: 5, Fourth: 6,
				Ninth: 7, Second: 8, Seventh: 1, Sixth: 2,
				Fifth: 9,
			},
			{
				First: 5, Eighth: 6, Third: 7, Fourth: 8,
				Ninth: 1, Second: 2, Seventh: 3, Sixth: 4,
				Fifth: 9,
			},
		}
		for idx, by := range inputs {
			np := original.rotate(by)
			if np != expect[idx] {
				t.Fatalf("rotate palace, expect %+v, got %+v", expect[idx], np)
			}
		}
	})
}
