package chinese_calendar

import (
	"strings"
)

// SexagenaryTerm 干支
type SexagenaryTerm struct {
	CelestialStem
	TerrestrialBranch
}

// NewSexagenaryTermFromIndex 通过索引值创建干支
// 索引范围0-59, 分别表示 0:甲子, 1:乙丑, ... 59:癸亥
// 大于59,或小于0会按照模算数,转换到0-59,取对应干支
func NewSexagenaryTermFromIndex(idx int) SexagenaryTerm {
	if idx < 0 {
		idx = idx%60 + 60
	}
	return SexagenaryTermEnum.JiaZi.Add(idx)
}

// NewSexagenaryTermFromWord 通过干支中文返回该类型值
func NewSexagenaryTermFromWord(word string) (SexagenaryTerm, bool) {
	ct := strings.Split(word, "")
	if len(ct) != 2 {
		return SexagenaryTermEnum.JiaZi, false
	}
	c, cValid := NewCelestialStemFromWord(ct[0])
	t, tValid := NewTerrestrialBranchFromWord(ct[1])
	if !cValid || !tValid {
		return SexagenaryTermEnum.JiaZi, false
	}
	return SexagenaryTerm{
		CelestialStem:     c,
		TerrestrialBranch: t,
	}, true
}

// Next 获得该干支的下一项
// 例: x=甲子, x.Next() -> 乙丑
//     x=乙丑, x.Next() -> 丙寅
//	   ...
//     x=癸亥, x.Next() -> 甲子
func (s SexagenaryTerm) Next() SexagenaryTerm {
	return SexagenaryTerm{
		CelestialStem:     s.CelestialStem.Next(),
		TerrestrialBranch: s.TerrestrialBranch.Next(),
	}
}

// Prev 获得该干支的上一项, Next的逆向操作
func (s SexagenaryTerm) Prev() SexagenaryTerm {
	return SexagenaryTerm{
		CelestialStem:     s.CelestialStem.Prev(),
		TerrestrialBranch: s.TerrestrialBranch.Prev(),
	}
}

// Add 获得该干支之前/后的任一项, Next和Prev的推广
// n < 0 向前回朔, n > 0 向后推算
func (s SexagenaryTerm) Add(n int) SexagenaryTerm {
	return SexagenaryTerm{
		CelestialStem:     s.CelestialStem.Add(n),
		TerrestrialBranch: s.TerrestrialBranch.Add(n),
	}
}

// Index 返回该干支的索引值(0-59)
func (s SexagenaryTerm) Index() int {
	if s.CelestialStem >= 10 || s.CelestialStem < 0 || s.TerrestrialBranch >= 12 || s.TerrestrialBranch < 0 {
		return 0
	}
	c := int(s.CelestialStem)
	t := int(s.TerrestrialBranch)

	return (((c+12)-t)%12)/2*10 + c
}

// String 返回干支中文
func (s SexagenaryTerm) String() string {
	c := s.CelestialStem.String()
	t := s.TerrestrialBranch.String()
	if c == "" || t == "" {
		return ""
	}
	return c + t
}

// SexagenaryTermEnum 60干支枚举项
var SexagenaryTermEnum = struct {
	JiaZi   SexagenaryTerm // 甲子
	JiaXu   SexagenaryTerm // 甲戌
	JiaShen SexagenaryTerm // 甲申
	JiaWu   SexagenaryTerm // 甲午
	JiaChen SexagenaryTerm // 甲辰
	JiaYin  SexagenaryTerm // 甲寅

	YiChou SexagenaryTerm // 乙丑
	YiHai  SexagenaryTerm // 乙亥
	YiYou  SexagenaryTerm // 乙酉
	YiWei  SexagenaryTerm // 乙未
	YiSi   SexagenaryTerm // 乙巳
	YiMao  SexagenaryTerm // 乙卯

	BingYin  SexagenaryTerm // 丙寅
	BingZi   SexagenaryTerm // 丙子
	BingXu   SexagenaryTerm // 丙戌
	BingShen SexagenaryTerm // 丙申
	BingWu   SexagenaryTerm // 丙午
	BingChen SexagenaryTerm // 丙辰

	DingMao  SexagenaryTerm // 丁卯
	DingChou SexagenaryTerm // 丁丑
	DingHai  SexagenaryTerm // 丁亥
	DingYou  SexagenaryTerm // 丁酉
	DingWei  SexagenaryTerm // 丁未
	DingSi   SexagenaryTerm // 丁巳

	WuChen SexagenaryTerm // 戊辰
	WuYin  SexagenaryTerm // 戊寅
	WuZi   SexagenaryTerm // 戊子
	WuXu   SexagenaryTerm // 戊戌
	WuShen SexagenaryTerm // 戊申
	WuWu   SexagenaryTerm // 戊午

	JiSi   SexagenaryTerm // 己巳
	JiMao  SexagenaryTerm // 己卯
	JiChou SexagenaryTerm // 己丑
	JiHai  SexagenaryTerm // 己亥
	JiYou  SexagenaryTerm // 己酉
	JiWei  SexagenaryTerm // 己未

	GengWu   SexagenaryTerm // 庚午
	GengChen SexagenaryTerm // 庚辰
	GengYin  SexagenaryTerm // 庚寅
	GengZi   SexagenaryTerm // 庚子
	GengXu   SexagenaryTerm // 庚戌
	GengShen SexagenaryTerm // 庚申

	XinWei  SexagenaryTerm // 辛未
	XinSi   SexagenaryTerm // 辛巳
	XinMao  SexagenaryTerm // 辛卯
	XinChou SexagenaryTerm // 辛丑
	XinHai  SexagenaryTerm // 辛亥
	XinYou  SexagenaryTerm // 辛酉

	RenShen SexagenaryTerm // 壬申
	RenWu   SexagenaryTerm // 壬午
	RenChen SexagenaryTerm // 壬辰
	RenYin  SexagenaryTerm // 壬寅
	RenZi   SexagenaryTerm // 壬子
	RenXu   SexagenaryTerm // 壬戌

	GuiYou  SexagenaryTerm // 癸酉
	GuiWei  SexagenaryTerm // 癸未
	GuiSi   SexagenaryTerm // 癸巳
	GuiMao  SexagenaryTerm // 癸卯
	GuiChou SexagenaryTerm // 癸丑
	GuiHai  SexagenaryTerm // 癸亥
}{
	JiaZi:   SexagenaryTerm{CelestialStemEnum.Jia, TerrestrialBranchEnum.Zi},   // 甲子
	JiaXu:   SexagenaryTerm{CelestialStemEnum.Jia, TerrestrialBranchEnum.Xu},   // 甲戌
	JiaShen: SexagenaryTerm{CelestialStemEnum.Jia, TerrestrialBranchEnum.Shen}, // 甲申
	JiaWu:   SexagenaryTerm{CelestialStemEnum.Jia, TerrestrialBranchEnum.Wu},   // 甲午
	JiaChen: SexagenaryTerm{CelestialStemEnum.Jia, TerrestrialBranchEnum.Chen}, // 甲辰
	JiaYin:  SexagenaryTerm{CelestialStemEnum.Jia, TerrestrialBranchEnum.Yin},  // 甲寅

	YiChou: SexagenaryTerm{CelestialStemEnum.Yi, TerrestrialBranchEnum.Chou}, // 乙丑
	YiHai:  SexagenaryTerm{CelestialStemEnum.Yi, TerrestrialBranchEnum.Hai},  // 乙亥
	YiYou:  SexagenaryTerm{CelestialStemEnum.Yi, TerrestrialBranchEnum.You},  // 乙酉
	YiWei:  SexagenaryTerm{CelestialStemEnum.Yi, TerrestrialBranchEnum.Wei},  // 乙未
	YiSi:   SexagenaryTerm{CelestialStemEnum.Yi, TerrestrialBranchEnum.Si},   // 乙巳
	YiMao:  SexagenaryTerm{CelestialStemEnum.Yi, TerrestrialBranchEnum.Mao},  // 乙卯

	BingYin:  SexagenaryTerm{CelestialStemEnum.Bing, TerrestrialBranchEnum.Yin},  // 丙寅
	BingZi:   SexagenaryTerm{CelestialStemEnum.Bing, TerrestrialBranchEnum.Zi},   // 丙子
	BingXu:   SexagenaryTerm{CelestialStemEnum.Bing, TerrestrialBranchEnum.Xu},   // 丙戌
	BingShen: SexagenaryTerm{CelestialStemEnum.Bing, TerrestrialBranchEnum.Shen}, // 丙申
	BingWu:   SexagenaryTerm{CelestialStemEnum.Bing, TerrestrialBranchEnum.Wu},   // 丙午
	BingChen: SexagenaryTerm{CelestialStemEnum.Bing, TerrestrialBranchEnum.Chen}, // 丙辰

	DingMao:  SexagenaryTerm{CelestialStemEnum.Ding, TerrestrialBranchEnum.Mao},  // 丁卯
	DingChou: SexagenaryTerm{CelestialStemEnum.Ding, TerrestrialBranchEnum.Chou}, // 丁丑
	DingHai:  SexagenaryTerm{CelestialStemEnum.Ding, TerrestrialBranchEnum.Hai},  // 丁亥
	DingYou:  SexagenaryTerm{CelestialStemEnum.Ding, TerrestrialBranchEnum.You},  // 丁酉
	DingWei:  SexagenaryTerm{CelestialStemEnum.Ding, TerrestrialBranchEnum.Wei},  // 丁未
	DingSi:   SexagenaryTerm{CelestialStemEnum.Ding, TerrestrialBranchEnum.Si},   // 丁巳

	WuChen: SexagenaryTerm{CelestialStemEnum.Wu, TerrestrialBranchEnum.Chen}, // 戊辰
	WuYin:  SexagenaryTerm{CelestialStemEnum.Wu, TerrestrialBranchEnum.Yin},  // 戊寅
	WuZi:   SexagenaryTerm{CelestialStemEnum.Wu, TerrestrialBranchEnum.Zi},   // 戊子
	WuXu:   SexagenaryTerm{CelestialStemEnum.Wu, TerrestrialBranchEnum.Xu},   // 戊戌
	WuShen: SexagenaryTerm{CelestialStemEnum.Wu, TerrestrialBranchEnum.Shen}, // 戊申
	WuWu:   SexagenaryTerm{CelestialStemEnum.Wu, TerrestrialBranchEnum.Wu},   // 戊午

	JiSi:   SexagenaryTerm{CelestialStemEnum.Ji, TerrestrialBranchEnum.Si},   // 己巳
	JiMao:  SexagenaryTerm{CelestialStemEnum.Ji, TerrestrialBranchEnum.Mao},  // 己卯
	JiChou: SexagenaryTerm{CelestialStemEnum.Ji, TerrestrialBranchEnum.Chou}, // 己丑
	JiHai:  SexagenaryTerm{CelestialStemEnum.Ji, TerrestrialBranchEnum.Hai},  // 己亥
	JiYou:  SexagenaryTerm{CelestialStemEnum.Ji, TerrestrialBranchEnum.You},  // 己酉
	JiWei:  SexagenaryTerm{CelestialStemEnum.Ji, TerrestrialBranchEnum.Wei},  // 己未

	GengWu:   SexagenaryTerm{CelestialStemEnum.Geng, TerrestrialBranchEnum.Wu},   // 庚午
	GengChen: SexagenaryTerm{CelestialStemEnum.Geng, TerrestrialBranchEnum.Chen}, // 庚辰
	GengYin:  SexagenaryTerm{CelestialStemEnum.Geng, TerrestrialBranchEnum.Yin},  // 庚寅
	GengZi:   SexagenaryTerm{CelestialStemEnum.Geng, TerrestrialBranchEnum.Zi},   // 庚子
	GengXu:   SexagenaryTerm{CelestialStemEnum.Geng, TerrestrialBranchEnum.Xu},   // 庚戌
	GengShen: SexagenaryTerm{CelestialStemEnum.Geng, TerrestrialBranchEnum.Shen}, // 庚申

	XinWei:  SexagenaryTerm{CelestialStemEnum.Xin, TerrestrialBranchEnum.Wei},  // 辛未
	XinSi:   SexagenaryTerm{CelestialStemEnum.Xin, TerrestrialBranchEnum.Si},   // 辛巳
	XinMao:  SexagenaryTerm{CelestialStemEnum.Xin, TerrestrialBranchEnum.Mao},  // 辛卯
	XinChou: SexagenaryTerm{CelestialStemEnum.Xin, TerrestrialBranchEnum.Chou}, // 辛丑
	XinHai:  SexagenaryTerm{CelestialStemEnum.Xin, TerrestrialBranchEnum.Hai},  // 辛亥
	XinYou:  SexagenaryTerm{CelestialStemEnum.Xin, TerrestrialBranchEnum.You},  // 辛酉

	RenShen: SexagenaryTerm{CelestialStemEnum.Ren, TerrestrialBranchEnum.Shen}, // 壬申
	RenWu:   SexagenaryTerm{CelestialStemEnum.Ren, TerrestrialBranchEnum.Wu},   // 壬午
	RenChen: SexagenaryTerm{CelestialStemEnum.Ren, TerrestrialBranchEnum.Chen}, // 壬辰
	RenYin:  SexagenaryTerm{CelestialStemEnum.Ren, TerrestrialBranchEnum.Yin},  // 壬寅
	RenZi:   SexagenaryTerm{CelestialStemEnum.Ren, TerrestrialBranchEnum.Zi},   // 壬子
	RenXu:   SexagenaryTerm{CelestialStemEnum.Ren, TerrestrialBranchEnum.Xu},   // 壬戌

	GuiYou:  SexagenaryTerm{CelestialStemEnum.Gui, TerrestrialBranchEnum.You},  // 癸酉
	GuiWei:  SexagenaryTerm{CelestialStemEnum.Gui, TerrestrialBranchEnum.Wei},  // 癸未
	GuiSi:   SexagenaryTerm{CelestialStemEnum.Gui, TerrestrialBranchEnum.Si},   // 癸巳
	GuiMao:  SexagenaryTerm{CelestialStemEnum.Gui, TerrestrialBranchEnum.Mao},  // 癸卯
	GuiChou: SexagenaryTerm{CelestialStemEnum.Gui, TerrestrialBranchEnum.Chou}, // 癸丑
	GuiHai:  SexagenaryTerm{CelestialStemEnum.Gui, TerrestrialBranchEnum.Hai},  // 癸亥
}
