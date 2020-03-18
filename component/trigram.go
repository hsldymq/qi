package component

var trigramWords = [8]string{"坤", "艮", "坎", "巽", "震", "離", "兌", "乾"}

// 八卦
type Trigram int

// NewTrigramFromPalaceIndex 从宫位索引创建八卦
// 0(1宫): 坎
// 1(2宫): 坤
// 2(3宫): 震
// 3(4宫): 巽
// 5(6宫): 乾
// 6(7宫): 兌
// 7(8宫): 艮
// 8(9宫): 离
func NewTrigramFromPalaceIndex(idx int) (Trigram, bool) {
	tMap := map[int]Trigram{
		0: TrigramEnum.Kan,
		1: TrigramEnum.Kun,
		2: TrigramEnum.Zhen,
		3: TrigramEnum.Xun,
		5: TrigramEnum.Qian,
		6: TrigramEnum.Dui,
		7: TrigramEnum.Gen,
		8: TrigramEnum.Li,
	}
	e, valid := tMap[idx]
	return e, valid
}

// Element 返回所对应的五行元素
func (t Trigram) Element() (Element, bool) {
	eMap := map[Trigram]Element{
		TrigramEnum.Kun:  ElementEnum.Earth,
		TrigramEnum.Gen:  ElementEnum.Earth,
		TrigramEnum.Kan:  ElementEnum.Water,
		TrigramEnum.Xun:  ElementEnum.Wood,
		TrigramEnum.Zhen: ElementEnum.Wood,
		TrigramEnum.Li:   ElementEnum.Fire,
		TrigramEnum.Dui:  ElementEnum.Metal,
		TrigramEnum.Qian: ElementEnum.Metal,
	}
	e, valid := eMap[t]
	return e, valid
}

// PalaceIndex 返回八卦对应宫位索引
// 坎: 0(1宫)
// 坤: 1(2宫)
// 震: 2(3宫)
// 巽: 3(4宫)
// 乾: 5(6宫)
// 兌: 6(7宫)
// 艮: 7(8宫)
// 离: 8(9宫)
func (t Trigram) PalaceIndex() (int, bool) {
	iMap := map[Trigram]int{
		TrigramEnum.Kan:  0,
		TrigramEnum.Kun:  1,
		TrigramEnum.Zhen: 2,
		TrigramEnum.Xun:  3,
		TrigramEnum.Qian: 5,
		TrigramEnum.Dui:  6,
		TrigramEnum.Gen:  7,
		TrigramEnum.Li:   8,
	}
	e, valid := iMap[t]
	return e, valid
}

func (t Trigram) String() string {
	if t >= 8 || t < 0 {
		return ""
	}
	return trigramWords[t]
}

// TrigramEnum 八卦枚举
var TrigramEnum = struct {
	Kun  Trigram
	Gen  Trigram
	Kan  Trigram
	Xun  Trigram
	Zhen Trigram
	Li   Trigram
	Dui  Trigram
	Qian Trigram
}{
	Kun:  0, // 坤
	Gen:  1, // 艮
	Kan:  2, // 坎
	Xun:  3, // 巽
	Zhen: 4, // 震
	Li:   5, // 離
	Dui:  6, // 兌
	Qian: 7, // 乾
}
