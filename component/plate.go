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

// RoundDistance 计算从当前宫顺时针旋转到指定宫的距离
func (pi PalaceIndex) RoundDistance(to PalaceIndex) int {
	if pi == to || !pi.IsValid() || pi == FifthPalace || !to.IsValid() || to == FifthPalace {
		return 0
	}

	p := pi
	distance := 0
	for ; p != to; distance++ {
		p = p.RoundNext()
	}
	return distance
}

// RoundNext 返回以顺时针旋转方式的下一宫的索引
// clockwise == true以顺时针方式旋转
// 0 -> 7, 7 -> 2, 2 -> 3, ..., 5 -> 0
// 4 -> 4, 第五宫始终固定不变
func (pi PalaceIndex) RoundNext() PalaceIndex {
	pMap := map[PalaceIndex]PalaceIndex{
		FirstPalace:   EighthPalace,
		SecondPalace:  SeventhPalace,
		ThirdPalace:   FourthPalace,
		FourthPalace:  NinthPalace,
		FifthPalace:   FifthPalace,
		SixthPalace:   FirstPalace,
		SeventhPalace: SixthPalace,
		EighthPalace:  ThirdPalace,
		NinthPalace:   SecondPalace,
	}
	return pMap[pi]
}

// Next 返回下一宫的索引号
// 0 -> 1, 1 -> 2, ..., 8 -> 0
func (pi PalaceIndex) Next() PalaceIndex {
	return pi.OffsetBy(1)
}

// Next 返回上一宫的索引号
// 8 -> 7, 7 -> 6, ..., 0 -> 8
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

// OriginalDoor 返回作为原始宫位所对应的门
func (pi PalaceIndex) OriginalDoor() Door {
	dMap := map[PalaceIndex]Door{
		FirstPalace:   DoorEnum.Xiu,
		SecondPalace:  DoorEnum.Si,
		ThirdPalace:   DoorEnum.Shang,
		FourthPalace:  DoorEnum.Du,
		FifthPalace:   DoorEnum.Si,
		SixthPalace:   DoorEnum.Kai,
		SeventhPalace: DoorEnum.Jing,
		EighthPalace:  DoorEnum.Sheng,
		NinthPalace:   DoorEnum.PJing,
	}
	return dMap[pi]
}

// OriginalDoor 返回作为原始宫位所对应的星
func (pi PalaceIndex) OriginalStar() Star {
	sMap := map[PalaceIndex]Star{
		FirstPalace:   StarEnum.TianPeng,
		SecondPalace:  StarEnum.TianRui,
		ThirdPalace:   StarEnum.TianChong,
		FourthPalace:  StarEnum.TianFu,
		FifthPalace:   StarEnum.TianQin,
		SixthPalace:   StarEnum.TianXin,
		SeventhPalace: StarEnum.TianZhu,
		EighthPalace:  StarEnum.TianRen,
		NinthPalace:   StarEnum.TianYing,
	}
	return sMap[pi]
}

func (pi PalaceIndex) IsValid() bool {
	return pi >= 0 && pi < 9
}

type PalaceValue int

func (pv PalaceValue) QiYi() QiYi {
	return QiYi(pv)
}

func (pv PalaceValue) Door() Door {
	return Door(pv)
}

func (pv PalaceValue) Star() Star {
	return Star(pv)
}

func (pv PalaceValue) God() God {
	return God(pv)
}

// Palaces 用于保存9宫的值
// 按照索引, 0代表1宫, 1代表2宫,一次类推
type Palaces struct {
	values [9]PalaceValue
}

func NewPalaces(values [9]PalaceValue) Palaces {
	return Palaces{values: values}
}

func NewQiYiPlate(roundPalaceIndex PalaceIndex, escaping Escaping) (Palaces, error) {
	p := Palaces{}
	if !roundPalaceIndex.IsValid() {
		return p, fmt.Errorf("invalid start index for creating palace")
	}

	values := [9]PalaceValue{
		PalaceValue(QiYiEnum.Wu), PalaceValue(QiYiEnum.Ji), PalaceValue(QiYiEnum.Geng),
		PalaceValue(QiYiEnum.Xin), PalaceValue(QiYiEnum.Ren), PalaceValue(QiYiEnum.Gui),
		PalaceValue(QiYiEnum.Ding), PalaceValue(QiYiEnum.Bing), PalaceValue(QiYiEnum.Yi),
	}

	pIndex := roundPalaceIndex
	for i := 0; i < 9; i++ {
		p.values[pIndex] = values[i]
		if escaping == YangEscaping {
			pIndex = pIndex.Next()
		} else {
			pIndex = pIndex.Prev()
		}
	}

	return p, nil
}

// NewOriginalGodPlate 创建原始神盘
// 阳遁局
//	|4: 六合|9: 勾陈|2: 朱雀|
//  |3: 太阴|5  　　|7: 九地|
//  |8: 腾蛇|1: 值符|6: 九天|
// 阴遁局
//	|4: 腾蛇|9: 值符|2: 九天|
//  |3: 太阴|5  　　|7: 九地|
//  |8: 六合|1: 白虎|6: 玄武|
func NewOriginalGodPlate(escaping Escaping) Palaces {
	var p Palaces
	if escaping == YangEscaping {
		p.values = [9]PalaceValue{
			FourthPalace: PalaceValue(GodEnum.LiuHe), NinthPalace: PalaceValue(GodEnum.GouChen), SecondPalace: PalaceValue(GodEnum.ZhuQue),
			ThirdPalace: PalaceValue(GodEnum.TaiYin), FifthPalace: -1, SeventhPalace: PalaceValue(GodEnum.JiuDi),
			EighthPalace: PalaceValue(GodEnum.TengShe), FirstPalace: PalaceValue(GodEnum.ZhiFu), SixthPalace: PalaceValue(GodEnum.JiuTian),
		}
	} else {
		p.values = [9]PalaceValue{
			FourthPalace: PalaceValue(GodEnum.TengShe), NinthPalace: PalaceValue(GodEnum.ZhiFu), SecondPalace: PalaceValue(GodEnum.JiuTian),
			ThirdPalace: PalaceValue(GodEnum.TaiYin), FifthPalace: -1, SeventhPalace: PalaceValue(GodEnum.JiuDi),
			EighthPalace: PalaceValue(GodEnum.LiuHe), FirstPalace: PalaceValue(GodEnum.BaiHu), SixthPalace: PalaceValue(GodEnum.XuanWu),
		}
	}
	return p
}

// NewOriginalHumanPlate 创建原始人盘
//	|4: 杜门|9: 景门|2: 死门|
//  |3: 伤门|5  　　|7: 惊门|
//  |8: 生门|1: 休门|6: 开门|
func NewOriginalHumanPlate() Palaces {
	return Palaces{
		values: [9]PalaceValue{
			FourthPalace: PalaceValue(DoorEnum.Du), NinthPalace: PalaceValue(DoorEnum.PJing), SecondPalace: PalaceValue(DoorEnum.Si),
			ThirdPalace: PalaceValue(DoorEnum.Shang), FifthPalace: -1, SeventhPalace: PalaceValue(DoorEnum.Jing),
			EighthPalace: PalaceValue(DoorEnum.Sheng), FirstPalace: PalaceValue(DoorEnum.Xiu), SixthPalace: PalaceValue(DoorEnum.Kai),
		},
	}
}

// NewOriginalStarPlate 创建原始九星天盘
//	|4: 天辅|9: 天英|2: 天芮|
//  |3: 天冲|5  天禽|7: 天柱|
//  |8: 天任|1: 天蓬|6: 天心|
func NewOriginStarPlate() Palaces {
	return Palaces{
		values: [9]PalaceValue{
			FourthPalace: PalaceValue(StarEnum.TianFu), NinthPalace: PalaceValue(StarEnum.TianYing), SecondPalace: PalaceValue(StarEnum.TianRui),
			ThirdPalace: PalaceValue(StarEnum.TianChong), FifthPalace: PalaceValue(StarEnum.TianQin), SeventhPalace: PalaceValue(StarEnum.TianZhu),
			EighthPalace: PalaceValue(StarEnum.TianRen), FirstPalace: PalaceValue(StarEnum.TianPeng), SixthPalace: PalaceValue(StarEnum.TianXin),
		},
	}
}

// Value 根据宫索引值获得其值
// 0表示1宫, 1表示2宫, 以此类推
func (p Palaces) Value(palaceIndex PalaceIndex) PalaceValue {
	if !palaceIndex.IsValid() {
		return -1
	}
	return p.values[palaceIndex]
}

// FindPalaceIndex 根据值找到它所在的宫索引
// 返回0代表1宫, 返回1代表2宫, 以此类推
func (p Palaces) FindPalaceIndex(value int) PalaceIndex {
	for idx, v := range p.values {
		if PalaceValue(value) == v {
			return PalaceIndex(idx)
		}
	}
	return PalaceIndex(-1)
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
	original := [8]PalaceValue{
		p.values[FirstPalace], p.values[EighthPalace], p.values[ThirdPalace], p.values[FourthPalace],
		p.values[NinthPalace], p.values[SecondPalace], p.values[SeventhPalace], p.values[SixthPalace],
	}
	rotated := [8]PalaceValue{}
	by = by % 8
	if by < 0 {
		by += 8
	}
	for idx, v := range original {
		newIdx := (idx + by) % 8
		rotated[newIdx] = v
	}
	return Palaces{
		values: [9]PalaceValue{
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
