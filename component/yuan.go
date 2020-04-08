package component

// Yuan 元, 上元/中元/下元
type Yuan int

func (y Yuan) IsValid() bool {
	return y >= 0 && y < 3
}

var YuanEnum = struct {
	Upper  Yuan
	Middle Yuan
	Lower  Yuan
}{
	Upper:  0,
	Middle: 1,
	Lower:  2,
}
