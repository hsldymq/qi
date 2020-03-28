package component

import (
	"fmt"
)

const (
	FirstPalace   PalaceIndex = 0
	SecondPalace  PalaceIndex = 1
	ThirdPalace   PalaceIndex = 2
	FourthPalace  PalaceIndex = 3
	FifthPalace   PalaceIndex = 4
	SixthPalace   PalaceIndex = 5
	SeventhPalace PalaceIndex = 6
	EighthPalace  PalaceIndex = 7
	NinthPalace   PalaceIndex = 8
)

type PalaceIndex int

func (pi PalaceIndex) Next() PalaceIndex {
	return pi.OffsetBy(1)
}

func (pi PalaceIndex) Prev() PalaceIndex {
	return pi.OffsetBy(-1)
}

func (pi PalaceIndex) OffsetBy(num int) PalaceIndex {
	i := (pi + PalaceIndex(num)) % 9
	if i < 0 {
		i += 9
	}
	return i
}

func (pi PalaceIndex) IsValid() bool {
	return pi >= 0 && pi < 9
}

// Palaces 用于保存9宫的值
// 按照索引, 0代表1宫, 1代表2宫,一次类推
type Palaces struct {
	values [9]int
}

func NewPalace(values [9]int, startPalaceIndex PalaceIndex, asc bool) (Palaces, error) {
	p := Palaces{}
	if !startPalaceIndex.IsValid() {
		return p, fmt.Errorf("invalid start index for creating palace")
	}

	pIndex := startPalaceIndex
	for i := 0; i < 9; i++ {
		p.values[pIndex] = values[i]
		if asc {
			pIndex = pIndex.Next()
		} else {
			pIndex = pIndex.Prev()
		}
	}

	return p, nil
}

// Value 根据宫索引值获得其值
// 0表示1宫, 1表示2宫, 以此类推
func (p Palaces) Value(palaceIndex PalaceIndex) int {
	if !palaceIndex.IsValid() {
		return -1
	}
	return p.values[palaceIndex]
}

// FindPalaceIndex 根据值找到它所在的宫索引
// 返回0代表1宫, 返回1代表2宫, 以此类推
func (p Palaces) FindPalaceIndex(value int) int {
	for idx, v := range p.values {
		if value == v {
			return idx
		}
	}
	return -1
}

// RotateValues 按圆盘方式旋转他们的值
// 例如,在圆盘上有如下值
//		  4 9
//	    3 	  2
//		8     7
//		  1 6
//	如果RotateValue(1), 则顺时旋转1变为
//		  3 4
//	    8 	  9
//		1     2
//		  6 7
//	如果RotateValue(-1), 则逆时针旋转变为
//		  9 2
//	    4 	  7
//		3     6
//		  8 1
func (p Palaces) RotateValues(by int) Palaces {
	original := [8]int{
		p.values[FirstPalace], p.values[EighthPalace], p.values[ThirdPalace], p.values[FourthPalace],
		p.values[NinthPalace], p.values[SecondPalace], p.values[SeventhPalace], p.values[SixthPalace],
	}
	rotated := [8]int{}
	by = by % 8
	if by < 0 {
		by += 8
	}
	for idx, v := range original {
		newIdx := (idx + by) % 8
		rotated[newIdx] = v
	}
	return Palaces{
		values: [9]int{
			FourthPalace: rotated[3], NinthPalace: rotated[4], SecondPalace: rotated[5],
			ThirdPalace: rotated[2], FifthPalace: p.values[FifthPalace], SeventhPalace: rotated[6],
			EighthPalace: rotated[1], FirstPalace: rotated[0], SixthPalace: rotated[7],
		},
	}
}

// OffsetValue 按照宫序进行位移其值
// 例:
//		offsetBy == 1时,
//		1宫的值移到2宫, 2宫的值移到3宫, 以此类推, 9宫的值移到1宫
//		offsetBy == -1时,
//		9宫的值移到8宫, 8宫的值移到7宫, 以此类推, 1宫的值移到9宫
// 当skipFifthPalace为true时, 5宫的值固定,不参与位移, 那么当offsetBy == 1, 4宫的值会移到6宫
func (p Palaces) OffsetValue(offsetBy int, skipFifthPalace bool) Palaces {
	newPalaces := Palaces{}
	palace := FirstPalace
	for i := 0; i < 9; i++ {
		offsetPalace := palace.OffsetBy(offsetBy)
		//if skipFifthPalace && palace == FifthPalace {
		//	newPalaces.values[FifthPalace] = original[FifthPalace]
		//	goto next
		//}
		//if skipFifthPalace {
		//
		//	offsetPalace = offsetPalace.OffsetBy(moreOffset)
		//	if offsetBy > 0 && offsetPalace >= FifthPalace && (palace < FifthPalace || palace > offsetPalace) {
		//		offsetPalace = offsetPalace.Next()
		//	} else if offsetBy < 0 && offsetPalace <= FifthPalace && (palace > FifthPalace || palace < offsetPalace) {
		//		offsetPalace = offsetPalace.Prev()
		//	}
		//}
		newPalaces.values[offsetPalace] = p.values[palace]
		//next:
		palace = palace.Next()
	}

	return newPalaces
}
