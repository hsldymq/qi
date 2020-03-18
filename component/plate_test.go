package component

import "testing"

func TestPalace(t *testing.T) {
	t.Run("test RotateValue method", func(t *testing.T) {
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
				First: 6, Eighth: 7, Third: 8, Fourth: 1, Ninth: 2, Second: 3, Seventh: 4, Sixth: 5,
				Fifth: 9,
			},
			{
				First: 8, Eighth: 1, Third: 2, Fourth: 3, Ninth: 4, Second: 5, Seventh: 6, Sixth: 7,
				Fifth: 9,
			},
			{
				First: 3, Eighth: 4, Third: 5, Fourth: 6, Ninth: 7, Second: 8, Seventh: 1, Sixth: 2,
				Fifth: 9,
			},
			{
				First: 5, Eighth: 6, Third: 7, Fourth: 8, Ninth: 1, Second: 2, Seventh: 3, Sixth: 4,
				Fifth: 9,
			},
		}
		for idx, by := range inputs {
			np := original.RotateValue(by)
			if np != expect[idx] {
				t.Fatalf("RotateValue palace, expect %+v, got %+v", expect[idx], np)
			}
		}
	})

	t.Run("test OffsetValue Method", func(t *testing.T) {
		original := Palace{
			First:   1,
			Second:  2,
			Third:   3,
			Fourth:  4,
			Fifth:   5,
			Sixth:   6,
			Seventh: 7,
			Eighth:  8,
			Ninth:   9,
		}
		nonSkipInputs := []int{1, 5, 9, -1, -5, -10}
		nonSkipExpect := []Palace{
			{First: 9, Second: 1, Third: 2, Fourth: 3, Fifth: 4, Sixth: 5, Seventh: 6, Eighth: 7, Ninth: 8},
			{First: 5, Second: 6, Third: 7, Fourth: 8, Fifth: 9, Sixth: 1, Seventh: 2, Eighth: 3, Ninth: 4},
			{First: 1, Second: 2, Third: 3, Fourth: 4, Fifth: 5, Sixth: 6, Seventh: 7, Eighth: 8, Ninth: 9},
			{First: 2, Second: 3, Third: 4, Fourth: 5, Fifth: 6, Sixth: 7, Seventh: 8, Eighth: 9, Ninth: 1},
			{First: 6, Second: 7, Third: 8, Fourth: 9, Fifth: 1, Sixth: 2, Seventh: 3, Eighth: 4, Ninth: 5},
			{First: 2, Second: 3, Third: 4, Fourth: 5, Fifth: 6, Sixth: 7, Seventh: 8, Eighth: 9, Ninth: 1},
		}

		for idx, offset := range nonSkipInputs {
			actual := original.OffsetValue(offset, false)
			if actual != nonSkipExpect[idx] {
				t.Fatalf("offset by %d: expect %+v, got %+v", offset, nonSkipExpect[idx], actual)
			}
		}

		skipInputs := []int{1, 4, 10, -1, -5, -13}
		skipExpect := []Palace{
			{First: 9, Second: 1, Third: 2, Fourth: 3, Fifth: 5, Sixth: 4, Seventh: 6, Eighth: 7, Ninth: 8},
			{First: 6, Second: 7, Third: 8, Fourth: 9, Fifth: 5, Sixth: 1, Seventh: 2, Eighth: 3, Ninth: 4},
			{First: 8, Second: 9, Third: 1, Fourth: 2, Fifth: 5, Sixth: 3, Seventh: 4, Eighth: 6, Ninth: 7},
			{First: 2, Second: 3, Third: 4, Fourth: 6, Fifth: 5, Sixth: 7, Seventh: 8, Eighth: 9, Ninth: 1},
			{First: 7, Second: 8, Third: 9, Fourth: 1, Fifth: 5, Sixth: 2, Seventh: 3, Eighth: 4, Ninth: 6},
			{First: 7, Second: 8, Third: 9, Fourth: 1, Fifth: 5, Sixth: 2, Seventh: 3, Eighth: 4, Ninth: 6},
		}

		for idx, offset := range skipInputs {
			actual := original.OffsetValue(offset, true)
			if actual != skipExpect[idx] {
				t.Fatalf("skip 5th palace, offset by %d: expect %+v, got %+v", offset, skipExpect[idx], actual)
			}
		}
	})
}
