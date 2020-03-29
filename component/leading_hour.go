package component

import (
	"github.com/hsldymq/go-chinese-calendar/sexagenary"
)

// LeadingHour 旬首
type LeadingHour int

// LeadingHourOfSexagenary 根据干支创建对应的旬首
func LeadingHourOfSexagenary(term sexagenary.SexagenaryTerm) LeadingHour {
	return LeadingHour(term.Index() / 10)
}

// SixYi 返回旬首对应的六仪
func (lh LeadingHour) SixYi() QiYi {
	yMap := map[LeadingHour]QiYi{
		LeadingHourEnum.JiaZi:   QiYiEnum.Wu,
		LeadingHourEnum.JiaXu:   QiYiEnum.Ji,
		LeadingHourEnum.JiaShen: QiYiEnum.Geng,
		LeadingHourEnum.JiaWu:   QiYiEnum.Xin,
		LeadingHourEnum.JiaChen: QiYiEnum.Ren,
		LeadingHourEnum.JiaYin:  QiYiEnum.Gui,
	}

	return yMap[lh]
}

func (lh LeadingHour) SexagenaryTerm() sexagenary.SexagenaryTerm {
	sMap := map[LeadingHour]sexagenary.SexagenaryTerm{
		LeadingHourEnum.JiaZi:   sexagenary.SexagenaryTermEnum.JiaZi,
		LeadingHourEnum.JiaXu:   sexagenary.SexagenaryTermEnum.JiaXu,
		LeadingHourEnum.JiaShen: sexagenary.SexagenaryTermEnum.JiaShen,
		LeadingHourEnum.JiaWu:   sexagenary.SexagenaryTermEnum.JiaWu,
		LeadingHourEnum.JiaChen: sexagenary.SexagenaryTermEnum.JiaChen,
		LeadingHourEnum.JiaYin:  sexagenary.SexagenaryTermEnum.JiaYin,
	}

	return sMap[lh]
}

func (lh LeadingHour) IsValid() bool {
	return lh >= 0 && lh < 6
}

var LeadingHourEnum = struct {
	JiaZi   LeadingHour
	JiaXu   LeadingHour
	JiaShen LeadingHour
	JiaWu   LeadingHour
	JiaChen LeadingHour
	JiaYin  LeadingHour
}{
	JiaZi:   0,
	JiaXu:   1,
	JiaShen: 2,
	JiaWu:   3,
	JiaChen: 4,
	JiaYin:  5,
}
