package qimen

import (
	"testing"

	"github.com/hsldymq/qimen/component"
)

func TestMakeTerrestrialPlate(t *testing.T) {
	p, err := makeBasePlate(component.ThirdPalace, component.YangEscaping)
	if err != nil {
		t.Fatalf("expect a yang escaping terrestrial plate, got error: %s", err)
	}
	expect := [9]int{
		int(component.QiYiEnum.Bing), int(component.QiYiEnum.Yi), int(component.QiYiEnum.Wu),
		int(component.QiYiEnum.Ji), int(component.QiYiEnum.Geng), int(component.QiYiEnum.Xin),
		int(component.QiYiEnum.Ren), int(component.QiYiEnum.Gui), int(component.QiYiEnum.Ding),
	}

	for i := 0; i < 9; i++ {
		pi := component.PalaceIndex(i)
		if p.Value(pi) != expect[i] {
			t.Fatalf("yang escaping, expect value of palace index %d to be %d, got %d", i, expect[i], p.Value(pi))
		}
	}

	p, err = makeBasePlate(component.SeventhPalace, component.YinEscaping)
	if err != nil {
		t.Fatalf("expect a yin escaping terrestrial plate, got error: %s", err)
	}
	expect = [9]int{
		int(component.QiYiEnum.Ding), int(component.QiYiEnum.Gui), int(component.QiYiEnum.Ren),
		int(component.QiYiEnum.Xin), int(component.QiYiEnum.Geng), int(component.QiYiEnum.Ji),
		int(component.QiYiEnum.Wu), int(component.QiYiEnum.Yi), int(component.QiYiEnum.Bing),
	}

	for i := 0; i < 9; i++ {
		pi := component.PalaceIndex(i)
		if p.Value(pi) != expect[i] {
			t.Fatalf("yin escaping, expect value of palace index %d to be %d, got %d", i, expect[i], p.Value(pi))
		}
	}
}
