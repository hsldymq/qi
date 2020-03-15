package chinese_calendar

import "math"

// solarTerms 24节气
// 通常春分的黄经定义为0°, 所以我们以春分作为开始
var solarTerms = []string{
	"春分", "清明", "谷雨", "立夏", "小满", "芒种",
	"夏至", "小暑", "大暑", "立秋", "处暑", "白露",
	"秋分", "寒露", "霜降", "立冬", "小雪", "大雪",
	"冬至", "小寒", "大寒", "立春", "雨水", "惊蛰",
}

// solarTermsTraditional 24节气繁体
var solarTermsTraditional = []string{
	"春分", "清明", "穀雨", "立夏", "小滿", "芒種",
	"夏至", "小暑", "大暑", "立秋", "處暑", "白露",
	"秋分", "寒露", "霜降", "立冬", "小雪", "大雪",
	"冬至", "小寒", "大寒", "立春", "雨水", "驚蟄",
}

// SolarTerm 24节气类型
type SolarTerm int

func (st SolarTerm) String(simplified bool) string {
	if int(st) >= len(solarTerms) || int(st) < 0 {
		return ""
	}

	if simplified {
		return solarTerms[int(st)]
	}
	return solarTermsTraditional[int(st)]
}

// EclipticLongitude 黄道经度
type EclipticLongitude float64

// SolarTerm 根据黄道经度得到对应节气
func (l EclipticLongitude) SolarTerm() SolarTerm {
	longitude := float64(l)
	if longitude >= 360 || longitude < 0 {
		longitude -= math.Floor(longitude/360) * 360
	}
	idx := int(math.Floor(longitude / 15))
	return SolarTerm(idx)
}

// SolarTermEnum 24节气枚举
// 英文命名参照: http://www.cma.gov.cn/2011xzt/essjqzt/jqhz/jqhz02/201312/t20131213_233952.html
var SolarTermEnum = struct {
	TheSpringEquinox      SolarTerm // 春分
	PureBrightness        SolarTerm // 清明
	GrainRain             SolarTerm // 谷雨
	TheBeginningOfSummer  SolarTerm // 立夏
	LesserFullnessOfGrain SolarTerm // 小满
	GrainInBeard          SolarTerm // 芒种
	TheSummerSolstice     SolarTerm // 夏至
	LesserHeat            SolarTerm // 小暑
	GreaterHeat           SolarTerm // 大暑
	TheBeginningOfAutumn  SolarTerm // 立秋
	TheEndOfHeat          SolarTerm // 处暑
	WhiteDew              SolarTerm // 白露
	TheAutumnEquinox      SolarTerm // 秋分
	ColdDew               SolarTerm // 寒露
	FrostsDescent         SolarTerm // 霜降
	TheBeginningOfWinter  SolarTerm // 立冬
	LesserSnow            SolarTerm // 小雪
	GreaterSnow           SolarTerm // 大雪
	TheWinterSolstice     SolarTerm // 冬至
	LesserCold            SolarTerm // 小寒
	GreaterCold           SolarTerm // 大寒
	TheBeginningOfSpring  SolarTerm // 立春
	RainWater             SolarTerm // 雨水
	TheWakingOfInsects    SolarTerm // 惊蛰
}{
	TheSpringEquinox:      0,
	PureBrightness:        1,
	GrainRain:             2,
	TheBeginningOfSummer:  3,
	LesserFullnessOfGrain: 4,
	GrainInBeard:          5,
	TheSummerSolstice:     6,
	LesserHeat:            7,
	GreaterHeat:           8,
	TheBeginningOfAutumn:  9,
	TheEndOfHeat:          10,
	WhiteDew:              11,
	TheAutumnEquinox:      12,
	ColdDew:               13,
	FrostsDescent:         14,
	TheBeginningOfWinter:  15,
	LesserSnow:            16,
	GreaterSnow:           17,
	TheWinterSolstice:     18,
	LesserCold:            19,
	GreaterCold:           20,
	TheBeginningOfSpring:  21,
	RainWater:             22,
	TheWakingOfInsects:    23,
}
