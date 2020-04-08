package component

// godWords 八神中文
var godWords = [10]string{"直符", "朱雀", "玄武", "太阴", "六合", "九天", "九地", "螣蛇", "勾陈", "白虎"}

// God 八神
// 阳遁八神: 直符, 朱雀, 太阴, 六合, 九天, 九地, 腾蛇, 勾陈
// 阴遁八神: 直符, 玄武, 太阴, 六合, 九天, 九地, 腾蛇, 白虎
type God int

// String 返回八神中文
func (g God) String() string {
	if !g.IsValid() {
		return ""
	}
	return godWords[g]
}

func (g God) IsValid() bool {
	return g >= 0 && g < 10
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
		GodEnum.ZhiFu:   StarEnum.TianPeng,
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

// GodEnum 八神枚举
var GodEnum = struct {
	ZhiFu   God // 直符
	ZhuQue  God // 朱雀
	XuanWu  God // 玄武
	TaiYin  God // 太阴
	LiuHe   God // 六合
	JiuTian God // 九天
	JiuDi   God // 九地
	TengShe God // 螣蛇
	GouChen God // 勾陈
	BaiHu   God // 白虎
}{
	ZhiFu:   0,
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
