package qi

var qiWords = []string{"乙", "丙", "丁"}

// Qi 奇
type Qi int

func (q Qi) String() string {
	if int(q) >= len(qiWords) || q < 0 {
		return ""
	}
	return qiWords[int(q)]
}

// 乙奇,丙奇,丁奇
var (
	Yi   Qi = 0
	Bing Qi = 1
	Ding Qi = 2
)
