package men

// 门有 休/生/伤/杜/景/死/惊/开 八门, 分布于人盘, 其中
// 开/休/生 为吉门
// 死/惊/伤 为凶门
// 杜/景    为中平

// menWorkds 门列表
var menWords = []string{"开", "休", "生", "死", "惊", "伤", "杜", "景"}

// Men 门
type Men int

// String
func (m Men) String() string {
	if int(m) >= len(menWords) || int(m) < 0 {
		return ""
	}
	return menWords[int(m)]
}

// IsJiMen 是否为吉门
func (m Men) IsJiMen() bool {
	return m == 0 || m == 1 || m == 2
}

// IsXiongMen 是否为凶门
func (m Men) IsXiongMen() bool {
	return m == 3 || m == 4 || m == 5
}

// IsZhongPing 是否为中平
func (m Men) IsZhongPing() bool {
	return m == 6 || m == 7
}

// JiMen 吉门
var JiMen = struct {
	Kai   Men
	Xiu   Men
	Sheng Men
}{
	Kai:   0,
	Xiu:   1,
	Sheng: 2,
}

// XiongMen 凶门
var XiongMen = struct {
	Si    Men
	Jing  Men
	Shang Men
}{
	Si:    3,
	Jing:  4,
	Shang: 5,
}

// ZhongPing 平门
var ZhongPing = struct {
	Du   Men
	Jing Men
}{
	Du:   6,
	Jing: 7,
}
