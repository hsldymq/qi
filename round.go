package qimen

import (
	"fmt"

	calendar "github.com/hsldymq/go-chinese-calendar"
	"github.com/hsldymq/go-chinese-calendar/solar"
	"github.com/hsldymq/qimen/component"
)

var solarTermPalaceIndex = map[solar.SolarTerm][2]component.PalaceIndex{
	solar.SolarTermEnum.TheWinterSolstice:     {component.FirstPalace, component.FirstPalace},     // 冬至1宫, 上元局1宫
	solar.SolarTermEnum.LesserCold:            {component.FirstPalace, component.SecondPalace},    // 小寒1宫, 上元局2宫
	solar.SolarTermEnum.GreaterCold:           {component.FirstPalace, component.ThirdPalace},     // 大寒1宫, 上元局3宫
	solar.SolarTermEnum.TheBeginningOfSpring:  {component.EighthPalace, component.EighthPalace},   // 立春8宫, 上元局8宫
	solar.SolarTermEnum.RainWater:             {component.EighthPalace, component.NinthPalace},    // 雨水8宫, 上元局9宫
	solar.SolarTermEnum.TheWakingOfInsects:    {component.EighthPalace, component.FirstPalace},    // 惊蛰8宫, 上元局1宫
	solar.SolarTermEnum.TheSpringEquinox:      {component.ThirdPalace, component.ThirdPalace},     // 春分3宫, 上元局3宫
	solar.SolarTermEnum.PureBrightness:        {component.ThirdPalace, component.FourthPalace},    // 清明3宫, 上元局4宫
	solar.SolarTermEnum.GrainRain:             {component.ThirdPalace, component.FifthPalace},     // 谷雨3宫, 上元局5宫
	solar.SolarTermEnum.TheBeginningOfSummer:  {component.FourthPalace, component.FourthPalace},   // 立夏4宫, 上元局4宫
	solar.SolarTermEnum.LesserFullnessOfGrain: {component.FourthPalace, component.FifthPalace},    // 小满4宫, 上元局5宫
	solar.SolarTermEnum.GrainInBeard:          {component.FourthPalace, component.SixthPalace},    // 芒种4宫, 上元局6宫
	solar.SolarTermEnum.TheSummerSolstice:     {component.NinthPalace, component.NinthPalace},     // 夏至9宫, 上元局9宫
	solar.SolarTermEnum.LesserHeat:            {component.NinthPalace, component.EighthPalace},    // 小暑9宫, 上元局8宫
	solar.SolarTermEnum.GreaterHeat:           {component.NinthPalace, component.SeventhPalace},   // 大暑9宫, 上元局7宫
	solar.SolarTermEnum.TheBeginningOfAutumn:  {component.SecondPalace, component.SecondPalace},   // 立秋2宫, 上元局2宫
	solar.SolarTermEnum.TheEndOfHeat:          {component.SecondPalace, component.FirstPalace},    // 处暑2宫, 上元局1宫
	solar.SolarTermEnum.WhiteDew:              {component.SecondPalace, component.NinthPalace},    // 白露2宫, 上元局9宫
	solar.SolarTermEnum.TheAutumnEquinox:      {component.SeventhPalace, component.SeventhPalace}, // 秋分7宫, 上元局7宫
	solar.SolarTermEnum.ColdDew:               {component.SeventhPalace, component.SixthPalace},   // 寒露7宫, 上元局6宫
	solar.SolarTermEnum.FrostsDescent:         {component.SeventhPalace, component.FifthPalace},   // 霜降7宫, 上元局5宫
	solar.SolarTermEnum.TheBeginningOfWinter:  {component.SixthPalace, component.SixthPalace},     // 立冬6宫, 上元局6宫
	solar.SolarTermEnum.LesserSnow:            {component.SixthPalace, component.FifthPalace},     // 小雪6宫, 上元局5宫
	solar.SolarTermEnum.GreaterSnow:           {component.SixthPalace, component.FourthPalace},    // 大雪6宫, 上元局4宫
}

func GenerateRound(year, month, day, hour calendar.SexagenaryTerm, term solar.SolarTerm, pentad solar.Pentad) error {
	escaping := component.SolarTermEscaping(term)
	palaceIndexes, ok := solarTermPalaceIndex[term]
	if !ok {
		return fmt.Errorf("invalid solar term %d", term)
	}
	solarPalaceIndex, roundPalaceIndex := palaceIndexes[0], palaceIndexes[1].OffsetBy(int(pentad)*6)

	// TODO delete
	fmt.Println(solarPalaceIndex)

	// 地盘
	terrestrialPlate, err := makeTerrestrialPlate(roundPalaceIndex, escaping)
	fmt.Println(terrestrialPlate)
	if err != nil {
		return err
	}

	return nil
}

func makeTerrestrialPlate(pIdx component.PalaceIndex, escaping component.Escaping) (component.Palaces, error) {
	return component.NewPalace(
		[9]int{
			int(calendar.CelestialStemEnum.Wu), int(calendar.CelestialStemEnum.Ji), int(calendar.CelestialStemEnum.Geng),
			int(calendar.CelestialStemEnum.Xin), int(calendar.CelestialStemEnum.Ren), int(calendar.CelestialStemEnum.Gui),
			int(calendar.CelestialStemEnum.Ding), int(calendar.CelestialStemEnum.Bing), int(calendar.CelestialStemEnum.Yi),
		},
		pIdx,
		escaping == component.YangEscaping,
	)
}
