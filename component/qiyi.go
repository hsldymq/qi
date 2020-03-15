package component

// 三奇 六仪

import chinese_calendar "github.com/hsldymq/go-chinese-calendar"

// Qi 三奇类型
type Qi chinese_calendar.CelestialStem

// CelestialStem 转换为对应的天支类型
func (q Qi) CelestialStem() chinese_calendar.CelestialStem {
	return chinese_calendar.CelestialStem(q)
}

// Value 返回枚举值
func (q Qi) Value() int {
	return int(q)
}

func (q Qi) String() string {
	qcl := chinese_calendar.CelestialStem(q)
	if qcl != chinese_calendar.CelestialStemEnum.Yi &&
		qcl != chinese_calendar.CelestialStemEnum.Bing &&
		qcl != chinese_calendar.CelestialStemEnum.Ding {
		return ""
	}
	return qcl.String()
}

// QiEnum 三奇枚举
var QiEnum = struct {
	Yi   chinese_calendar.CelestialStem
	Bing chinese_calendar.CelestialStem
	Ding chinese_calendar.CelestialStem
}{
	Yi:   chinese_calendar.CelestialStemEnum.Yi,
	Bing: chinese_calendar.CelestialStemEnum.Bing,
	Ding: chinese_calendar.CelestialStemEnum.Ding,
}

// Yi 六仪类型
type Yi chinese_calendar.CelestialStem

// CelestialStem 转换为对应的天支类型
func (y Yi) CelestialStem() chinese_calendar.CelestialStem {
	return chinese_calendar.CelestialStem(y)
}

// Value 返回枚举值
func (y Yi) Value() int {
	return int(y)
}

func (y Yi) String() string {
	qcl := chinese_calendar.CelestialStem(y)
	if qcl == chinese_calendar.CelestialStemEnum.Wu ||
		qcl == chinese_calendar.CelestialStemEnum.Ji ||
		qcl == chinese_calendar.CelestialStemEnum.Geng ||
		qcl == chinese_calendar.CelestialStemEnum.Jia {
		return ""
	}

	return qcl.String()
}

// YiEnum 三奇枚举
var YiEnum = struct {
	Wu   chinese_calendar.CelestialStem
	Ji   chinese_calendar.CelestialStem
	Geng chinese_calendar.CelestialStem
	Xin  chinese_calendar.CelestialStem
	Ren  chinese_calendar.CelestialStem
	Gui  chinese_calendar.CelestialStem
}{
	Wu:   chinese_calendar.CelestialStemEnum.Wu,
	Ji:   chinese_calendar.CelestialStemEnum.Ji,
	Geng: chinese_calendar.CelestialStemEnum.Geng,
	Xin:  chinese_calendar.CelestialStemEnum.Xin,
	Ren:  chinese_calendar.CelestialStemEnum.Ren,
	Gui:  chinese_calendar.CelestialStemEnum.Gui,
}
