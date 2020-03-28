package component

import (
	"github.com/hsldymq/go-chinese-calendar/sexagenary"
)

// Qi 三奇六仪
type QiYi sexagenary.CelestialStem

// CelestialStem 转换为对应的天支类型
func (q QiYi) CelestialStem() sexagenary.CelestialStem {
	return sexagenary.CelestialStem(q)
}

// Value 返回枚举值
func (q QiYi) Value() int {
	return int(q)
}

func (q QiYi) String() string {
	qcl := sexagenary.CelestialStem(q)
	if qcl == sexagenary.CelestialStemEnum.Jia {
		return ""
	}
	return qcl.String()
}

// QiYiEnum 三奇六仪枚举
var QiYiEnum = struct {
	Yi   QiYi
	Bing QiYi
	Ding QiYi
	Wu   QiYi
	Ji   QiYi
	Geng QiYi
	Xin  QiYi
	Ren  QiYi
	Gui  QiYi
}{
	Yi:   QiYi(sexagenary.CelestialStemEnum.Yi),
	Bing: QiYi(sexagenary.CelestialStemEnum.Bing),
	Ding: QiYi(sexagenary.CelestialStemEnum.Ding),
	Wu:   QiYi(sexagenary.CelestialStemEnum.Wu),
	Ji:   QiYi(sexagenary.CelestialStemEnum.Ji),
	Geng: QiYi(sexagenary.CelestialStemEnum.Geng),
	Xin:  QiYi(sexagenary.CelestialStemEnum.Xin),
	Ren:  QiYi(sexagenary.CelestialStemEnum.Ren),
	Gui:  QiYi(sexagenary.CelestialStemEnum.Gui),
}
