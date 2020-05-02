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
		HourLeaderEnum.JiaZi:   QiYiEnum.Wu,
		HourLeaderEnum.JiaXu:   QiYiEnum.Ji,
		HourLeaderEnum.JiaShen: QiYiEnum.Geng,
		HourLeaderEnum.JiaWu:   QiYiEnum.Xin,
		HourLeaderEnum.JiaChen: QiYiEnum.Ren,
		HourLeaderEnum.JiaYin:  QiYiEnum.Gui,
	}

	return yMap[lh]
}

func (lh HourLeader) SexagenaryTerm() sexagenary.SexagenaryTerm {
	sMap := map[HourLeader]sexagenary.SexagenaryTerm{
		HourLeaderEnum.JiaZi:   sexagenary.SexagenaryTermEnum.JiaZi,
		HourLeaderEnum.JiaXu:   sexagenary.SexagenaryTermEnum.JiaXu,
		HourLeaderEnum.JiaShen: sexagenary.SexagenaryTermEnum.JiaShen,
		HourLeaderEnum.JiaWu:   sexagenary.SexagenaryTermEnum.JiaWu,
		HourLeaderEnum.JiaChen: sexagenary.SexagenaryTermEnum.JiaChen,
		HourLeaderEnum.JiaYin:  sexagenary.SexagenaryTermEnum.JiaYin,
	}

	return sMap[lh]
}

// HourVoid 返回旬首对应的旬空
// 将60干支按顺序排列成6排,每排10个,如下:
//	"甲子", "乙丑", "丙寅", "丁卯", "戊辰", "己巳", "庚午", "辛未", "壬申", "癸酉"
//  	"甲戌", "乙亥", "丙子", "丁丑", "戊寅", "己卯", "庚辰", "辛巳", "壬午", "癸未"
//  	"甲申", "乙酉", "丙戌", "丁亥", "戊子", "己丑", "庚寅", "辛卯", "壬辰", "癸巳"
//  	"甲午", "乙未", "丙申", "丁酉", "戊戌", "己亥", "庚子", "辛丑", "壬寅", "癸卯"
//  	"甲辰", "乙巳", "丙午", "丁未", "戊申", "己酉", "庚戌", "辛亥", "壬子", "癸丑"
//  	"甲寅", "乙卯", "丙辰", "丁巳", "戊午", "己未", "庚申", "辛酉", "壬戌", "癸亥"
// 每排缺少的那两个地支就是旬空,也就是下一排的前两个地支
func (lh HourLeader) HourVoid() [2]sexagenary.TerrestrialBranch {
	return map[HourLeader][2]sexagenary.TerrestrialBranch{
		HourLeaderEnum.JiaZi:   {sexagenary.TerrestrialBranchEnum.Xu, sexagenary.TerrestrialBranchEnum.Hai},
		HourLeaderEnum.JiaXu:   {sexagenary.TerrestrialBranchEnum.Shen, sexagenary.TerrestrialBranchEnum.You},
		HourLeaderEnum.JiaShen: {sexagenary.TerrestrialBranchEnum.Wu, sexagenary.TerrestrialBranchEnum.Wei},
		HourLeaderEnum.JiaWu:   {sexagenary.TerrestrialBranchEnum.Chen, sexagenary.TerrestrialBranchEnum.Si},
		HourLeaderEnum.JiaChen: {sexagenary.TerrestrialBranchEnum.Yin, sexagenary.TerrestrialBranchEnum.Mao},
		HourLeaderEnum.JiaYin:  {sexagenary.TerrestrialBranchEnum.Zi, sexagenary.TerrestrialBranchEnum.Chou},
	}[lh]
}

func (lh HourLeader) IsValid() bool {
	return lh >= 0 && lh < 6
}

var HourLeaderEnum = struct {
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
