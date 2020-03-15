package component

// starWords 九星中文
var starWords = [9]string{"天蓬", "天芮", "天冲", "天辅", "天禽", "天心", "天柱", "天任", "天英"}

// Star 九星
type Star int

func (s Star) String() string {
	if s >= 9 || s < 0 {
		return ""
	}
	return starWords[s]
}

// StarEnum 九星枚举
var StarEnum = struct {
	TianPeng  Star
	TianRui   Star
	TianChong Star
	TianFu    Star
	TianQin   Star
	TianXin   Star
	TianZhu   Star
	TianRen   Star
	TianYing  Star
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
