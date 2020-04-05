package component

import (
	"testing"
)

func TestPalace(t *testing.T) {
	t.Run("test RotateValues method", func(t *testing.T) {
		original := Palaces{
			values: [9]PalaceValue{
				FourthPalace: 4, NinthPalace: 9, SecondPalace: 2,
				ThirdPalace: 3, FifthPalace: 5, SeventhPalace: 7,
				EighthPalace: 8, FirstPalace: 1, SixthPalace: 6,
			},
		}
		inputs := []int{3, 9, -2, -28}
		expect := []Palaces{
			{
				values: [9]PalaceValue{
					FourthPalace: 1, NinthPalace: 8, SecondPalace: 3,
					ThirdPalace: 6, FifthPalace: 5, SeventhPalace: 4,
					EighthPalace: 7, FirstPalace: 2, SixthPalace: 9,
				},
			},
			{
				values: [9]PalaceValue{
					FourthPalace: 3, NinthPalace: 4, SecondPalace: 9,
					ThirdPalace: 8, FifthPalace: 5, SeventhPalace: 2,
					EighthPalace: 1, FirstPalace: 6, SixthPalace: 7,
				},
			},
			{
				values: [9]PalaceValue{
					FourthPalace: 2, NinthPalace: 7, SecondPalace: 6,
					ThirdPalace: 9, FifthPalace: 5, SeventhPalace: 1,
					EighthPalace: 4, FirstPalace: 3, SixthPalace: 8,
				},
			},
			{
				values: [9]PalaceValue{
					FourthPalace: 6, NinthPalace: 1, SecondPalace: 8,
					ThirdPalace: 7, FifthPalace: 5, SeventhPalace: 3,
					EighthPalace: 2, FirstPalace: 9, SixthPalace: 4,
				},
			},
		}
		for idx, by := range inputs {
			np := original.RotateValues(by)
			if np != expect[idx] {
				t.Fatalf("RotateValues palace, expect %+v, got %+v", expect[idx], np)
			}
		}
	})

	t.Run("test OffsetValue Method", func(t *testing.T) {
		original := Palaces{
			values: [9]PalaceValue{1, 2, 3, 4, 5, 6, 7, 8, 9},
		}
		nonSkipInputs := []int{1, 5, 9, 67, -1, -5, -10, -31}
		nonSkipExpect := []Palaces{
			{values: [9]PalaceValue{9, 1, 2, 3, 4, 5, 6, 7, 8}},
			{values: [9]PalaceValue{5, 6, 7, 8, 9, 1, 2, 3, 4}},
			{values: [9]PalaceValue{1, 2, 3, 4, 5, 6, 7, 8, 9}},
			{values: [9]PalaceValue{6, 7, 8, 9, 1, 2, 3, 4, 5}},
			{values: [9]PalaceValue{2, 3, 4, 5, 6, 7, 8, 9, 1}},
			{values: [9]PalaceValue{6, 7, 8, 9, 1, 2, 3, 4, 5}},
			{values: [9]PalaceValue{2, 3, 4, 5, 6, 7, 8, 9, 1}},
			{values: [9]PalaceValue{5, 6, 7, 8, 9, 1, 2, 3, 4}},
		}

		for idx, offset := range nonSkipInputs {
			actual := original.OffsetValue(offset, false)
			if actual != nonSkipExpect[idx] {
				t.Fatalf("offset by %d: expect %+v, got %+v", offset, nonSkipExpect[idx], actual)
			}
		}

		//skipInputs := []int{1, 4, 10, 65, -1, -5, -13, -39}
		//skipExpect := []Palaces{
		//	{values: [9]int{9, 1, 2, 3, 5, 4, 6, 7, 8}},
		//	{values: [9]int{6, 7, 8, 9, 5, 1, 2, 3, 4}},
		//	{values: [9]int{8, 9, 1, 2, 5, 3, 4, 6, 7}},
		//	{values: [9]int{9, 1, 2, 3, 5, 4, 6, 7, 8}},
		//	{values: [9]int{2, 3, 4, 6, 5, 7, 8, 9, 1}},
		//	{values: [9]int{7, 8, 9, 1, 5, 2, 3, 4, 6}},
		//	{values: [9]int{9, 1, 2, 3, 5, 4, 6, 7, 8}},
		//}
		//
		//for idx, offset := range skipInputs {
		//	actual := original.OffsetValue(offset, true)
		//	if actual != skipExpect[idx] {
		//		t.Fatalf("skip 5th palace, offset by %d: expect %+v, got %+v", offset, skipExpect[idx], actual)
		//	}
		//}
	})
}

func TestNewQiYiPlate(t *testing.T) {
	p, err := NewQiYiPlate(ThirdPalace, YangEscaping)
	if err != nil {
		t.Fatalf("expect a yang escaping terrestrial plate, got error: %s", err)
	}
	expect := [9]PalaceValue{
		PalaceValue(QiYiEnum.Bing), PalaceValue(QiYiEnum.Yi), PalaceValue(QiYiEnum.Wu),
		PalaceValue(QiYiEnum.Ji), PalaceValue(QiYiEnum.Geng), PalaceValue(QiYiEnum.Xin),
		PalaceValue(QiYiEnum.Ren), PalaceValue(QiYiEnum.Gui), PalaceValue(QiYiEnum.Ding),
	}

	for i := 0; i < 9; i++ {
		pi := PalaceIndex(i)
		if p.Value(pi) != expect[i] {
			t.Fatalf("yang escaping, expect value of palace index %d to be %d, got %d", i, expect[i], p.Value(pi))
		}
	}

	p, err = NewQiYiPlate(SeventhPalace, YinEscaping)
	if err != nil {
		t.Fatalf("expect a yin escaping terrestrial plate, got error: %s", err)
	}
	expect = [9]PalaceValue{
		PalaceValue(QiYiEnum.Ding), PalaceValue(QiYiEnum.Gui), PalaceValue(QiYiEnum.Ren),
		PalaceValue(QiYiEnum.Xin), PalaceValue(QiYiEnum.Geng), PalaceValue(QiYiEnum.Ji),
		PalaceValue(QiYiEnum.Wu), PalaceValue(QiYiEnum.Yi), PalaceValue(QiYiEnum.Bing),
	}

	for i := 0; i < 9; i++ {
		pi := PalaceIndex(i)
		if p.Value(pi) != expect[i] {
			t.Fatalf("yin escaping, expect value of palace index %d to be %d, got %d", i, expect[i], p.Value(pi))
		}
	}
}
