package component

// RoundPlate 圆盘
type RoundPlate struct {
	palace Palace
}

func NewRoundPlate(palace Palace) RoundPlate {
	return RoundPlate{
		palace: palace,
	}
}

// Palace 用于保存宫的值
type Palace struct {
	First   int // 1宫
	Second  int // 2宫
	Third   int // 3宫
	Fourth  int // 4宫
	Fifth   int // 5宫
	Sixth   int // 6宫
	Seventh int // 7宫
	Eighth  int // 8宫
	Ninth   int // 9宫
}

// Value 根据宫索引值获得其值
// 0表示1宫, 1表示2宫, 以此类推
func (p Palace) Value(palaceIndex int) int {
	switch palaceIndex {
	case 0:
		return p.First
	case 1:
		return p.Second
	case 2:
		return p.Third
	case 3:
		return p.Fourth
	case 4:
		return p.Fifth
	case 5:
		return p.Sixth
	case 6:
		return p.Seventh
	case 7:
		return p.Eighth
	case 8:
		return p.Ninth
	default:
		return -1
	}
}

// FindPalaceIndex 根据值找到它所在的宫索引
// 返回0代表1宫, 返回1代表2宫, 以此类推
func (p Palace) FindPalaceIndex(value int) int {
	switch value {
	case p.First:
		return 0
	case p.Second:
		return 1
	case p.Third:
		return 2
	case p.Fourth:
		return 3
	case p.Fifth:
		return 4
	case p.Sixth:
		return 5
	case p.Seventh:
		return 6
	case p.Eighth:
		return 7
	case p.Ninth:
		return 8
	default:
		return -1
	}
}

// RotateValue 按圆盘方式旋转他们的值
func (p Palace) RotateValue(by int) Palace {
	original := [8]int{
		p.First, p.Eighth, p.Third, p.Fourth,
		p.Ninth, p.Second, p.Seventh, p.Sixth,
	}
	rotated := [8]int{}
	if by < 0 {
		by = by%8 + 8
	}
	by = by % 8
	for idx, v := range original {
		newIdx := (idx + by) % 8
		rotated[newIdx] = v
	}
	return Palace{
		First:   rotated[0],
		Eighth:  rotated[1],
		Third:   rotated[2],
		Fourth:  rotated[3],
		Ninth:   rotated[4],
		Second:  rotated[5],
		Seventh: rotated[6],
		Sixth:   rotated[7],
		Fifth:   p.Fifth,
	}
}

// OffsetValue 按照宫序进行位移其值
// 例:
//		offsetBy == 1时,
//		1宫的值移到2宫, 2宫的值移到3宫, 以此类推, 9宫的值移到1宫
//		offsetBy == -1时,
//		9宫的值移到8宫, 8宫的值移到7宫, 以此类推, 1宫的值移到9宫
// 当skipFifthPalace为true时, 5宫的值固定,不参与位移, 那么当offsetBy == 1, 4宫的值会移到6宫
func (p Palace) OffsetValue(offsetBy int, skipFifthPalace bool) Palace {
	var original [9]int
	if skipFifthPalace {
		original = [9]int{
			p.First, p.Second, p.Third, p.Fourth, p.Sixth, p.Seventh, p.Eighth, p.Ninth, p.Fifth,
		}
	} else {
		original = [9]int{
			p.First, p.Second, p.Third, p.Fourth, p.Fifth, p.Sixth, p.Seventh, p.Eighth, p.Ninth,
		}
	}

	var rotated [9]int
	rotateNum := 9
	if skipFifthPalace {
		rotateNum = 8
	}
	if offsetBy < 0 {
		offsetBy = offsetBy%rotateNum + rotateNum
	}
	for i := 0; i < rotateNum; i++ {
		rotated[(i+offsetBy)%rotateNum] = original[i]
	}

	if skipFifthPalace {
		return Palace{
			First:   rotated[0],
			Second:  rotated[1],
			Third:   rotated[2],
			Fourth:  rotated[3],
			Fifth:   original[8],
			Sixth:   rotated[4],
			Seventh: rotated[5],
			Eighth:  rotated[6],
			Ninth:   rotated[7],
		}
	} else {
		return Palace{
			First:   rotated[0],
			Second:  rotated[1],
			Third:   rotated[2],
			Fourth:  rotated[3],
			Fifth:   rotated[4],
			Sixth:   rotated[5],
			Seventh: rotated[6],
			Eighth:  rotated[7],
			Ninth:   rotated[8],
		}
	}

}
