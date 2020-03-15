package component

var elementWords = [6]string{"", "水", "火", "木", "金", "土"}

// Element 五行元素
type Element int

// Generate 返回该元素所生的元素
// 木生火, 火生土, 土生金, 金生水, 水生木
func (e Element) Generate() Element {
	gMap := map[Element]Element{
		ElementEnum.Wood:  ElementEnum.Fire,
		ElementEnum.Fire:  ElementEnum.Earth,
		ElementEnum.Earth: ElementEnum.Metal,
		ElementEnum.Metal: ElementEnum.Water,
		ElementEnum.Water: ElementEnum.Wood,
	}

	g, ok := gMap[e]
	if !ok {
		return e
	}
	return g
}

// GeneratedBy 返回生成该元素的元素, Generate方法的逆操作
func (e Element) GeneratedBy() Element {
	gMap := map[Element]Element{
		ElementEnum.Fire:  ElementEnum.Wood,
		ElementEnum.Earth: ElementEnum.Fire,
		ElementEnum.Metal: ElementEnum.Earth,
		ElementEnum.Water: ElementEnum.Metal,
		ElementEnum.Wood:  ElementEnum.Water,
	}

	g, ok := gMap[e]
	if !ok {
		return e
	}
	return g
}

func (e Element) IsGeneratedBy(ele Element) bool {
	return e.GeneratedBy() == ele
}

// Overcome 返回该元素所克的元素
// 木克土, 土克水, 水克火, 火克金, 金克木
func (e Element) Overcome() Element {
	gMap := map[Element]Element{
		ElementEnum.Wood:  ElementEnum.Earth,
		ElementEnum.Earth: ElementEnum.Water,
		ElementEnum.Water: ElementEnum.Fire,
		ElementEnum.Fire:  ElementEnum.Metal,
		ElementEnum.Metal: ElementEnum.Wood,
	}
	g, ok := gMap[e]
	if !ok {
		return e
	}
	return g
}

// OvercameBy 返回克该元素的元素, Overcome方法的逆操作
func (e Element) OvercameBy() Element {
	gMap := map[Element]Element{
		ElementEnum.Earth: ElementEnum.Wood,
		ElementEnum.Water: ElementEnum.Earth,
		ElementEnum.Fire:  ElementEnum.Water,
		ElementEnum.Metal: ElementEnum.Fire,
		ElementEnum.Wood:  ElementEnum.Metal,
	}
	g, ok := gMap[e]
	if !ok {
		return e
	}
	return g
}

func (e Element) IsOvercameBy(ele Element) bool {
	return e.OvercameBy() == ele
}

func (e Element) String() string {
	if e >= 6 || e < 0 {
		return ""
	}
	return elementWords[e]
}

// ElementEnum 元素枚举
var ElementEnum = struct {
	Water Element
	Fire  Element
	Wood  Element
	Metal Element
	Earth Element
}{
	Water: 1,
	Fire:  2,
	Wood:  3,
	Metal: 4,
	Earth: 5,
}
