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
	hourLeader := component.NewHourLeader(hour)
	leadingHourPalaceIndex := qiYiPlate.FindPalaceIndex(hourLeader.SixYi().Value())

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

func TestGenerateRound(t *testing.T) {
	cases := []RoundInfo{
		{Escaping: component.YangEscaping, RoundPalaceIndex: component.FourthPalace, Yuan: component.YuanEnum.Lower, SexagenaryHour: sexagenary.SexagenaryTermEnum.WuZi},
	}
	expect := []Round{
		{DutyStar: component.StarEnum.TianXin, DutyDoor: component.DoorEnum.Kai, QiYiTerrestrialPlate: component.NewPalaces([9]component.PalaceValue{3, 2, 1, 4, 5, 6, 7, 8, 9}), QiYiCelestialPlate: component.NewPalaces([9]component.PalaceValue{9, 8, 7, 6, 5, 4, 1, 2, 3}), StarCelestialPlate: component.NewPalaces([9]component.PalaceValue{8, 7, 6, 5, 4, 3, 2, 1, 0}), HumanPlate: component.NewPalaces([9]component.PalaceValue{4, 7, 6, 2, -1, 5, 1, 0, 3}), GodPlate: component.NewPalaces([9]component.PalaceValue{1, 3, 5, 0, -1, 8, 4, 6, 7})},
	}
	for idx, each := range cases {
		actual, err := GenerateRound(each)
		exp := expect[idx]
		if err != nil {
			t.Fatalf("error occurred during generating round(%d): %s", idx, err)
		}
		if exp.DutyDoor != actual.DutyDoor {
			t.Fatalf("%d: expect %s, got %s", idx, exp.DutyDoor, actual.DutyDoor)
		}
		if exp.DutyStar != actual.DutyStar {
			t.Fatalf("case %d: expect %s, got %s", idx, exp.DutyStar, actual.DutyStar)
		}
		for i := 0; i < 9; i++ {
			if i == 4 {
				continue
			}
			expVal := exp.QiYiTerrestrialPlate.Value(component.PalaceIndex(i))
			actVal := actual.QiYiTerrestrialPlate.Value(component.PalaceIndex(i))
			if expVal != actVal {
				t.Fatalf("case %d: terrestrial plate, palace %d, expect %s, got %s", idx, i+1, expVal.QiYi(), actVal.QiYi())
			}
		}

		for i := 0; i < 9; i++ {
			if i == 4 {
				continue
			}
			expVal := exp.QiYiCelestialPlate.Value(component.PalaceIndex(i))
			actVal := actual.QiYiCelestialPlate.Value(component.PalaceIndex(i))
			if expVal != actVal {
				t.Fatalf("case %d: qi yi celestial plate, palace %d, expect %s, got %s", idx, i+1, expVal.QiYi(), actVal.QiYi())
			}
		}

		for i := 0; i < 9; i++ {
			if i == 4 {
				continue
			}
			expVal := exp.HumanPlate.Value(component.PalaceIndex(i))
			actVal := actual.HumanPlate.Value(component.PalaceIndex(i))
			if expVal != actVal {
				t.Fatalf("case %d: human plate, palace %d, expect %s, got %s", idx, i+1, expVal.Door(), actVal.Door())
			}
		}

		for i := 0; i < 9; i++ {
			if i == 4 {
				continue
			}
			expVal := exp.StarCelestialPlate.Value(component.PalaceIndex(i))
			actVal := actual.StarCelestialPlate.Value(component.PalaceIndex(i))
			if expVal != actVal {
				t.Fatalf("case %d: star celestial plate, palace %d, expect %s, got %s", idx, i+1, expVal.Star(), actVal.Star())
			}
		}

		for i := 0; i < 9; i++ {
			if i == 4 {
				continue
			}
			expVal := exp.GodPlate.Value(component.PalaceIndex(i))
			actVal := actual.GodPlate.Value(component.PalaceIndex(i))
			if expVal != actVal {
				t.Fatalf("case %d: god plate, palace %d, expect %s, got %s", idx, i+1, expVal.God(), actVal.God())
			}
		}
	}
}
