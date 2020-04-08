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
	Xiu   Door // 休门
	Si    Door // 死门
	Shang Door // 伤门
	Du    Door // 杜门
	Kai   Door // 开门
	Jing  Door // 惊门
	Sheng Door // 生门
	PJing Door // 景门
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
