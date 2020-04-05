package qimen

import (
	"testing"

	"github.com/hsldymq/go-chinese-calendar/sexagenary"
	"github.com/hsldymq/qimen/component"
)

// TestRotateStarCelestialPlate 测试旋转九星天盘
func TestRotateStarCelestialPlate(t *testing.T) {
	qiYiPlate, _ := component.NewQiYiPlate(component.ThirdPalace, component.YangEscaping)
	hour := sexagenary.SexagenaryTermEnum.GengWu
	info := roundInfo{
		LeadingHour:      component.LeadingHourOfSexagenary(hour),
		Escaping:         component.YangEscaping,
		RoundPalaceIndex: component.ThirdPalace,
	}
	leadingHourPalaceIndex := qiYiPlate.FindPalaceIndex(info.LeadingHour.SixYi().Value())

	dutyStar, qiYiCelestialPlate, starCelestialPlate := rotateStarCelestialPlate(leadingHourPalaceIndex, hour, qiYiPlate)
	if dutyStar != component.StarEnum.TianChong {
		t.Fatalf("wrong duty star, expect %s, got %s", component.StarEnum.TianChong.String(), dutyStar.String())
	}

	pIndex := component.FirstPalace
	expect := [9]component.PalaceValue{
		component.FourthPalace: component.PalaceValue(component.QiYiEnum.Bing), component.NinthPalace: component.PalaceValue(component.QiYiEnum.Gui), component.SecondPalace: component.PalaceValue(component.QiYiEnum.Wu),
		component.ThirdPalace: component.PalaceValue(component.QiYiEnum.Xin), component.FifthPalace: component.PalaceValue(component.QiYiEnum.Geng), component.SeventhPalace: component.PalaceValue(component.QiYiEnum.Ji),
		component.EighthPalace: component.PalaceValue(component.QiYiEnum.Ren), component.FirstPalace: component.PalaceValue(component.QiYiEnum.Yi), component.SixthPalace: component.PalaceValue(component.QiYiEnum.Ding),
	}
	for i := 0; i < 9; i++ {
		actual := qiYiCelestialPlate.Value(pIndex)
		if expect[i] != actual {
			t.Fatalf("qi yi plate, palace index %d, expect %s, got %s", i, expect[i].QiYi().String(), actual.QiYi().String())
		}
		pIndex = pIndex.Next()
	}

	pIndex = component.FirstPalace
	expect = [9]component.PalaceValue{
		component.FourthPalace: component.PalaceValue(component.StarEnum.TianPeng), component.NinthPalace: component.PalaceValue(component.StarEnum.TianRen), component.SecondPalace: component.PalaceValue(component.StarEnum.TianChong),
		component.ThirdPalace: component.PalaceValue(component.StarEnum.TianXin), component.FifthPalace: component.PalaceValue(component.StarEnum.TianQin), component.SeventhPalace: component.PalaceValue(component.StarEnum.TianFu),
		component.EighthPalace: component.PalaceValue(component.StarEnum.TianZhu), component.FirstPalace: component.PalaceValue(component.StarEnum.TianRui), component.SixthPalace: component.PalaceValue(component.StarEnum.TianYing),
	}
	for i := 0; i < 9; i++ {
		actual := starCelestialPlate.Value(pIndex)
		if expect[i] != actual {
			t.Fatalf("star celestial plate, palace index %d, expect %s, got %s", i, expect[i].Star().String(), actual.Star().String())
		}
		pIndex = pIndex.Next()
	}
}
