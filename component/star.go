package component

// starWords 九星中文
var starWords = [9]string{"天蓬", "天芮", "天冲", "天辅", "天禽", "天心", "天柱", "天任", "天英"}

// Star 九星
type Star int

// String 返回九星对应的中文
func (s Star) String() string {
	if !s.IsValid() {
		return ""
	}
	return starWords[s]
}

func (s Star) IsValid() bool {
	return s >= 0 && s < 9
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
