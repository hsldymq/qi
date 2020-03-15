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

// rotate 按圆盘方式旋转他们的值
func (p Palace) rotate(by int) Palace {
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
