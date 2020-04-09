package component

import (
	"github.com/hsldymq/go-chinese-calendar/sexagenary"
)

// HourLeader 旬首
type HourLeader int

// NewHourLeader 根据干支时创建对应的旬首
func NewHourLeader(term sexagenary.SexagenaryTerm) HourLeader {
	return HourLeader(term.Index() / 10)
}

// SixYi 返回旬首对应的六仪
func (lh HourLeader) SixYi() QiYi {
	yMap := map[HourLeader]QiYi{
		LeadingHourEnum.JiaZi:   QiYiEnum.Wu,
		LeadingHourEnum.JiaXu:   QiYiEnum.Ji,
		LeadingHourEnum.JiaShen: QiYiEnum.Geng,
		LeadingHourEnum.JiaWu:   QiYiEnum.Xin,
		LeadingHourEnum.JiaChen: QiYiEnum.Ren,
		LeadingHourEnum.JiaYin:  QiYiEnum.Gui,
	}

	return yMap[lh]
}

func (lh HourLeader) SexagenaryTerm() sexagenary.SexagenaryTerm {
	sMap := map[HourLeader]sexagenary.SexagenaryTerm{
		LeadingHourEnum.JiaZi:   sexagenary.SexagenaryTermEnum.JiaZi,
		LeadingHourEnum.JiaXu:   sexagenary.SexagenaryTermEnum.JiaXu,
		LeadingHourEnum.JiaShen: sexagenary.SexagenaryTermEnum.JiaShen,
		LeadingHourEnum.JiaWu:   sexagenary.SexagenaryTermEnum.JiaWu,
		LeadingHourEnum.JiaChen: sexagenary.SexagenaryTermEnum.JiaChen,
		LeadingHourEnum.JiaYin:  sexagenary.SexagenaryTermEnum.JiaYin,
	}

	return sMap[lh]
}

func (lh HourLeader) IsValid() bool {
	return lh >= 0 && lh < 6
}

var LeadingHourEnum = struct {
	JiaZi   HourLeader
	JiaXu   HourLeader
	JiaShen HourLeader
	JiaWu   HourLeader
	JiaChen HourLeader
	JiaYin  HourLeader
}{
	JiaZi:   0,
	JiaXu:   1,
	JiaShen: 2,
	JiaWu:   3,
	JiaChen: 4,
	JiaYin:  5,
}
