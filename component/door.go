package component

// 八门, 代表人世间凶吉的八个行动方向布于人盘, 分别为:
//		休门, 死门, 伤门, 杜门, 开门, 惊门, 生门, 景门

// doorWords 门中文
var doorWords = [9]string{"休", "死", "伤", "杜", "开", "惊", "生", "景"}

// Door 门
type Door int

// Value 返回门对应的枚举值
func (m Door) Value() int {
	return int(m)
}

// String
func (m Door) String() string {
	if int(m) >= 9 || int(m) < 0 {
		return ""
	}
	return doorWords[int(m)]
}

// DoorEnum 门枚举
var DoorEnum = struct {
	Xiu   Door // 休
	Si    Door // 死
	Shang Door // 伤
	Du    Door // 杜
	Kai   Door // 开
	Jing  Door // 惊
	Sheng Door // 生
	PJing Door // 景
}{
	Xiu:   0,
	Si:    1,
	Shang: 2,
	Du:    3,
	Kai:   4,
	Jing:  5,
	Sheng: 6,
	PJing: 7,
}
