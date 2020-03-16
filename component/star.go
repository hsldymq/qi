package component

// starWords 九星中文
var starWords = [9]string{"天蓬", "天芮", "天冲", "天辅", "天禽", "天心", "天柱", "天任", "天英"}

// godWords 八神中文
var godWords = [10]string{"直符", "朱雀", "玄武", "太阴", "六合", "九天", "九地", "腾蛇", "勾陈", "白虎"}

// Star 九星
type Star int

// String 返回九星对应的中文
func (s Star) String() string {
	if s >= 9 || s < 0 {
		return ""
	}
	return starWords[s]
}

// String 返回九星所对应的8神
// 阳遁对应关系:
// 		直符 => 天蓬
// 		天芮(天禽) => 朱雀
// 		天冲 => 太阴
// 		天辅 => 六合
// 		天心 => 九天
// 		天柱 => 九地
// 		天任 => 腾蛇
// 		天英 => 勾陈
// 阴遁对应关系:
// 		直符 => 天蓬
// 		天芮(天禽) => 玄武
// 		天冲 => 太阴
// 		天辅 => 六合
// 		天心 => 九天
// 		天柱 => 九地
// 		天任 => 腾蛇
// 		天英 => 白虎
func (s Star) God(escaping Escaping) God {
	yangEscapingMap := map[Star]God{
		StarEnum.TianPeng:  GodEnum.Zhifu,
		StarEnum.TianRui:   GodEnum.ZhuQue,
		StarEnum.TianChong: GodEnum.TaiYin,
		StarEnum.TianFu:    GodEnum.LiuHe,
		StarEnum.TianXin:   GodEnum.JiuTian,
		StarEnum.TianZhu:   GodEnum.JiuDi,
		StarEnum.TianRen:   GodEnum.TengShe,
		StarEnum.TianYing:  GodEnum.GouChen,
	}
	YinEscapingMap := map[Star]God{
		StarEnum.TianPeng:  GodEnum.Zhifu,
		StarEnum.TianRui:   GodEnum.XuanWu,
		StarEnum.TianChong: GodEnum.TaiYin,
		StarEnum.TianFu:    GodEnum.LiuHe,
		StarEnum.TianXin:   GodEnum.JiuTian,
		StarEnum.TianZhu:   GodEnum.JiuDi,
		StarEnum.TianRen:   GodEnum.TengShe,
		StarEnum.TianYing:  GodEnum.BaiHu,
	}
	if escaping == YangEscaping {
		return yangEscapingMap[s]
	}
	return YinEscapingMap[s]
}

// God 八神
// 阳遁八神: 直符, 朱雀, 太阴, 六合, 九天, 九地, 腾蛇, 勾陈
// 阴遁八神: 直符, 玄武, 太阴, 六合, 九天, 九地, 腾蛇, 白虎
type God int

// String 返回八神中文
func (g God) String() string {
	if g >= 10 || g < 0 {
		return ""
	}
	return godWords[g]
}

// Star 返回神所在的星
// 对应关系:
// 		直符 => 天蓬
// 		朱雀(玄武) => 天芮(天禽)
// 		太阴 => 天冲
// 		六合 => 天辅
// 		九天 => 天心
// 		九地 => 天柱
// 		腾蛇 => 天任
// 		勾陈(白虎) => 天英
func (g God) Star() Star {
	starMap := map[God]Star{
		GodEnum.Zhifu:   StarEnum.TianPeng,
		GodEnum.ZhuQue:  StarEnum.TianRui,
		GodEnum.XuanWu:  StarEnum.TianRui,
		GodEnum.TaiYin:  StarEnum.TianChong,
		GodEnum.LiuHe:   StarEnum.TianFu,
		GodEnum.JiuTian: StarEnum.TianXin,
		GodEnum.JiuDi:   StarEnum.TianZhu,
		GodEnum.TengShe: StarEnum.TianRen,
		GodEnum.GouChen: StarEnum.TianYing,
		GodEnum.BaiHu:   StarEnum.TianYing,
	}
	return starMap[g]
}

// StarEnum 九星枚举
// 九星: 天蓬, 天芮, 天冲, 天辅, 天禽, 天心, 天柱, 天任, 天英
var StarEnum = struct {
	TianPeng  Star // 天蓬
	TianRui   Star // 天芮
	TianChong Star // 天冲
	TianFu    Star // 天辅
	TianQin   Star // 天禽
	TianXin   Star // 天心
	TianZhu   Star // 天柱
	TianRen   Star // 天任
	TianYing  Star // 天英
}{
	TianPeng:  0,
	TianRui:   1,
	TianChong: 2,
	TianFu:    3,
	TianQin:   4,
	TianXin:   5,
	TianZhu:   6,
	TianRen:   7,
	TianYing:  8,
}

// GodEnum 八神枚举
var GodEnum = struct {
	Zhifu   God // 直符
	ZhuQue  God // 朱雀
	XuanWu  God // 玄武
	TaiYin  God // 太阴
	LiuHe   God // 六合
	JiuTian God // 九天
	JiuDi   God // 九地
	TengShe God // 腾蛇
	GouChen God // 勾陈
	BaiHu   God // 白虎
}{
	Zhifu:   0,
	ZhuQue:  1,
	XuanWu:  2,
	TaiYin:  3,
	LiuHe:   4,
	JiuTian: 5,
	JiuDi:   6,
	TengShe: 7,
	GouChen: 8,
	BaiHu:   9,
}
