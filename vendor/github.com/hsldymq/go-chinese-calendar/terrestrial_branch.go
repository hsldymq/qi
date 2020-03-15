package chinese_calendar

import "time"

// terrestrialBranch 地支中文列表
var terrestrialBranchWords = [12]string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
var terrestrialBranchWordMap = map[string]TerrestrialBranch{
	"子": TerrestrialBranchEnum.Zi,
	"丑": TerrestrialBranchEnum.Chou,
	"寅": TerrestrialBranchEnum.Yin,
	"卯": TerrestrialBranchEnum.Mao,
	"辰": TerrestrialBranchEnum.Chen,
	"巳": TerrestrialBranchEnum.Si,
	"午": TerrestrialBranchEnum.Wu,
	"未": TerrestrialBranchEnum.Wei,
	"申": TerrestrialBranchEnum.Shen,
	"酉": TerrestrialBranchEnum.You,
	"戌": TerrestrialBranchEnum.Xu,
	"亥": TerrestrialBranchEnum.Hai,
}

// TerrestrialBranch 地支
type TerrestrialBranch int

// NewTerrestrialBranchFromTime 根据时间返回地支类型
func NewTerrestrialBranchFromTime(t time.Time) TerrestrialBranch {
	h := t.Hour()
	if h == 23 {
		h = 0
	}
	if h%2 == 1 {
		h += 1
	}
	return TerrestrialBranch(h / 2)
}

// NewTerrestrialBranchFromWord 从地支中文返回其类型
func NewTerrestrialBranchFromWord(word string) (TerrestrialBranch, bool) {
	cs, valid := terrestrialBranchWordMap[word]
	return cs, valid
}

// Next 获得该地支的下一项
// 例:   x=子, x.Next() -> 丑
//		 x=丑, x.Next() -> 寅
// 		 ...
//       x=亥, x.Next() -> 子
func (tb TerrestrialBranch) Next() TerrestrialBranch {
	return tb.Add(1)
}

// Prev 获得该地支的上一项, Next的逆操作
func (tb TerrestrialBranch) Prev() TerrestrialBranch {
	return tb.Add(-1)
}

// Add 获得该地支之前/后的任意项, Next和Prev的推广
// n < 0 向前回朔, n > 0 向后推算
func (tb TerrestrialBranch) Add(n int) TerrestrialBranch {
	ntb := int(tb) + n
	if ntb < 0 {
		ntb = ntb%12 + 12
	}
	ntb = ntb % 12
	return TerrestrialBranch(ntb)
}

// Month 返回对应的地支纪月
func (tb TerrestrialBranch) Month() int {
	if tb > 11 || tb < 0 {
		return 0
	}
	if tb == TerrestrialBranchEnum.Zi {
		return 11
	} else if tb == TerrestrialBranchEnum.Chou {
		return 12
	}
	return int(tb - 1)
}

// String 返回地支中文
func (tb TerrestrialBranch) String() string {
	if tb >= 12 || tb < 0 {
		return ""
	}
	return terrestrialBranchWords[tb]
}

// ZodiacSign 返回对应的生肖
func (tb TerrestrialBranch) ZodiacSign() ZodiacSign {
	return ZodiacSign(tb)
}

// TerrestrialBranchEnum 地支枚举项
var TerrestrialBranchEnum = struct {
	Zi   TerrestrialBranch
	Chou TerrestrialBranch
	Yin  TerrestrialBranch
	Mao  TerrestrialBranch
	Chen TerrestrialBranch
	Si   TerrestrialBranch
	Wu   TerrestrialBranch
	Wei  TerrestrialBranch
	Shen TerrestrialBranch
	You  TerrestrialBranch
	Xu   TerrestrialBranch
	Hai  TerrestrialBranch
}{
	Zi:   0,
	Chou: 1,
	Yin:  2,
	Mao:  3,
	Chen: 4,
	Si:   5,
	Wu:   6,
	Wei:  7,
	Shen: 8,
	You:  9,
	Xu:   10,
	Hai:  11,
}

// zodiac 生肖简体
var zodiacSignWords = [12]string{"鼠", "牛", "虎", "兔", "龙", "蛇", "马", "羊", "猴", "鸡", "狗", "猪"}

// zodiaz 生肖繁体
var zodiacSignWordsTraditional = [12]string{"鼠", "牛", "虎", "兔", "龍", "蛇", "馬", "羊", "猴", "雞", "狗", "豬"}

// ZodiacSign 生肖
type ZodiacSign int

// TerrestrialBranch 返回对应的地支
func (zs ZodiacSign) TerrestrialBranch() TerrestrialBranch {
	return TerrestrialBranch(zs)
}

func (zs ZodiacSign) String(simplified bool) string {
	if zs >= 12 || zs < 0 {
		return ""
	}

	if simplified {
		return zodiacSignWords[zs]
	}
	return zodiacSignWordsTraditional[zs]
}

// ZodiacSignEnum 生肖枚举项
var ZodiacSignEnum = struct {
	Rat     ZodiacSign
	Ox      ZodiacSign
	Tiger   ZodiacSign
	Rabbit  ZodiacSign
	Dragoon ZodiacSign
	Snake   ZodiacSign
	Horse   ZodiacSign
	Goat    ZodiacSign
	Monkey  ZodiacSign
	Rooster ZodiacSign
	Dog     ZodiacSign
	Pig     ZodiacSign
}{
	Rat:     0,
	Ox:      1,
	Tiger:   2,
	Rabbit:  3,
	Dragoon: 4,
	Snake:   5,
	Horse:   6,
	Goat:    7,
	Monkey:  8,
	Rooster: 9,
	Dog:     10,
	Pig:     11,
}
