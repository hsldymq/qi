package chinese_calendar

// celestialStem 天干中文列表
var celestialStemWords = [10]string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}
var celestialStemWordMap = map[string]CelestialStem{
	"甲": CelestialStemEnum.Jia,
	"乙": CelestialStemEnum.Yi,
	"丙": CelestialStemEnum.Bing,
	"丁": CelestialStemEnum.Ding,
	"戊": CelestialStemEnum.Wu,
	"己": CelestialStemEnum.Ji,
	"庚": CelestialStemEnum.Geng,
	"辛": CelestialStemEnum.Xin,
	"壬": CelestialStemEnum.Ren,
	"癸": CelestialStemEnum.Gui,
}

// CelestialStem 天干
type CelestialStem int

// NewCelestialStemFromWord 从天干中文还原成CelestialStem类型
// 例: cs := NewCelestialStemFromWord("甲")
//	   assert(cs == CelestialStemEnum.Jia)
func NewCelestialStemFromWord(word string) (CelestialStem, bool) {
	cs, valid := celestialStemWordMap[word]
	return cs, valid
}

// Next 获得该天干的下一项
// 例:   x=甲, x.Next() -> 乙
//		 x=乙, x.Next() -> 丙
// 		 ...
//       x=癸, x.Next() -> 甲
func (cs CelestialStem) Next() CelestialStem {
	return cs.Add(1)
}

// Prev 获得该天干的上一项, Next的逆操作
func (cs CelestialStem) Prev() CelestialStem {
	return cs.Add(-1)
}

// Add 获得该天干向前/后的任意项, Next和Prev的推广
// n < 0 向前回朔, n > 0 向后推算
func (cs CelestialStem) Add(n int) CelestialStem {
	ncs := int(cs) + n
	if ncs < 0 {
		ncs = ncs%10 + 10
	}
	ncs = ncs % 10
	return CelestialStem(ncs)
}

// String 返回天干中文
func (cs CelestialStem) String() string {
	if cs >= 10 || cs < 0 {
		return ""
	}
	return celestialStemWords[cs]
}

// CelestialStemEnum 天干枚举项
var CelestialStemEnum = struct {
	Jia  CelestialStem
	Yi   CelestialStem
	Bing CelestialStem
	Ding CelestialStem
	Wu   CelestialStem
	Ji   CelestialStem
	Geng CelestialStem
	Xin  CelestialStem
	Ren  CelestialStem
	Gui  CelestialStem
}{
	Jia:  0,
	Yi:   1,
	Bing: 2,
	Ding: 3,
	Wu:   4,
	Ji:   5,
	Geng: 6,
	Xin:  7,
	Ren:  8,
	Gui:  9,
}
