package qimen

import (
	"fmt"

	"github.com/hsldymq/go-chinese-calendar/sexagenary"
	"github.com/hsldymq/go-chinese-calendar/solar"
	"github.com/hsldymq/qimen/component"
)

var solarTermPalaceIndex = map[solar.SolarTerm][2]component.PalaceIndex{
	solar.SolarTermEnum.TheWinterSolstice:     {component.FirstPalace, component.FirstPalace},     // 冬至1宫, 上元局起宫1宫
	solar.SolarTermEnum.LesserCold:            {component.FirstPalace, component.SecondPalace},    // 小寒1宫, 上元局起宫2宫
	solar.SolarTermEnum.GreaterCold:           {component.FirstPalace, component.ThirdPalace},     // 大寒1宫, 上元局起宫3宫
	solar.SolarTermEnum.TheBeginningOfSpring:  {component.EighthPalace, component.EighthPalace},   // 立春8宫, 上元局起宫8宫
	solar.SolarTermEnum.RainWater:             {component.EighthPalace, component.NinthPalace},    // 雨水8宫, 上元局起宫9宫
	solar.SolarTermEnum.TheWakingOfInsects:    {component.EighthPalace, component.FirstPalace},    // 惊蛰8宫, 上元局起宫1宫
	solar.SolarTermEnum.TheSpringEquinox:      {component.ThirdPalace, component.ThirdPalace},     // 春分3宫, 上元局起宫3宫
	solar.SolarTermEnum.PureBrightness:        {component.ThirdPalace, component.FourthPalace},    // 清明3宫, 上元局起宫4宫
	solar.SolarTermEnum.GrainRain:             {component.ThirdPalace, component.FifthPalace},     // 谷雨3宫, 上元局起宫5宫
	solar.SolarTermEnum.TheBeginningOfSummer:  {component.FourthPalace, component.FourthPalace},   // 立夏4宫, 上元局起宫4宫
	solar.SolarTermEnum.LesserFullnessOfGrain: {component.FourthPalace, component.FifthPalace},    // 小满4宫, 上元局起宫5宫
	solar.SolarTermEnum.GrainInBeard:          {component.FourthPalace, component.SixthPalace},    // 芒种4宫, 上元局起宫6宫
	solar.SolarTermEnum.TheSummerSolstice:     {component.NinthPalace, component.NinthPalace},     // 夏至9宫, 上元局起宫9宫
	solar.SolarTermEnum.LesserHeat:            {component.NinthPalace, component.EighthPalace},    // 小暑9宫, 上元局起宫8宫
	solar.SolarTermEnum.GreaterHeat:           {component.NinthPalace, component.SeventhPalace},   // 大暑9宫, 上元局起宫7宫
	solar.SolarTermEnum.TheBeginningOfAutumn:  {component.SecondPalace, component.SecondPalace},   // 立秋2宫, 上元局起宫2宫
	solar.SolarTermEnum.TheEndOfHeat:          {component.SecondPalace, component.FirstPalace},    // 处暑2宫, 上元局起宫1宫
	solar.SolarTermEnum.WhiteDew:              {component.SecondPalace, component.NinthPalace},    // 白露2宫, 上元局起宫9宫
	solar.SolarTermEnum.TheAutumnEquinox:      {component.SeventhPalace, component.SeventhPalace}, // 秋分7宫, 上元局起宫7宫
	solar.SolarTermEnum.ColdDew:               {component.SeventhPalace, component.SixthPalace},   // 寒露7宫, 上元局起宫6宫
	solar.SolarTermEnum.FrostsDescent:         {component.SeventhPalace, component.FifthPalace},   // 霜降7宫, 上元局起宫5宫
	solar.SolarTermEnum.TheBeginningOfWinter:  {component.SixthPalace, component.SixthPalace},     // 立冬6宫, 上元局起宫6宫
	solar.SolarTermEnum.LesserSnow:            {component.SixthPalace, component.FifthPalace},     // 小雪6宫, 上元局起宫5宫
	solar.SolarTermEnum.GreaterSnow:           {component.SixthPalace, component.FourthPalace},    // 大雪6宫, 上元局起宫4宫
}

func GenerateRound(year, month, day, hour sexagenary.SexagenaryTerm, solarTerm solar.SolarTerm, pentad solar.Pentad) error {
	if !solarTerm.IsValid() {
		return fmt.Errorf("invalid solar solarTerm %d", solarTerm)
	}
	if !pentad.IsValid() {
		return fmt.Errorf("invalid pentad %d", pentad)
	}
	if !year.IsValid() {
		return fmt.Errorf("invalid sexagenary year %d:%d", year.CelestialStem, year.TerrestrialBranch)
	}
	if !month.IsValid() {
		return fmt.Errorf("invalid sexagenary month %d:%d", month.CelestialStem, month.TerrestrialBranch)
	}
	if !month.IsValid() {
		return fmt.Errorf("invalid sexagenary day %d:%d", day.CelestialStem, day.TerrestrialBranch)
	}
	if !month.IsValid() {
		return fmt.Errorf("invalid sexagenary hour %d:%d", hour.CelestialStem, hour.TerrestrialBranch)
	}

	// 旬首, 阴/阳遁, 节气所在的宫索引, 局的宫索引
	leadingHour, escaping, solarTermPalaceIndex, roundPalaceIndex := parseTime(year, month, day, hour, solarTerm, pentad)

	// 天盘
	celestialPlate, err := makeBasePlate(roundPalaceIndex, escaping)
	if err != nil {
		return err
	}

	// 地盘
	terrestrialPlate, err := makeBasePlate(roundPalaceIndex, escaping)
	if err != nil {
		return err
	}
	leadHourPalaceIndex := terrestrialPlate.FindPalaceIndex(leadingHour.SixYi().Value())

	// 值使. 又名值门. 使:八门之意
	dutyDoor := leadHourPalaceIndex.OriginalDoor()
	// 值符. 又名值星
	dutyStar := leadHourPalaceIndex.OriginalStar()
	fmt.Println(solarTermPalaceIndex, celestialPlate, terrestrialPlate, dutyDoor, dutyStar)

	return nil
}

func parseTime(year, month, day, hour sexagenary.SexagenaryTerm,
	solarTerm solar.SolarTerm,
	pentad solar.Pentad) (component.LeadingHour, component.Escaping, component.PalaceIndex, component.PalaceIndex) {
	// 旬首
	leadingHour := component.LeadingHourOfSexagenary(hour)
	// 阴/阳遁
	escaping := component.SolarTermEscaping(solarTerm)
	palaceIndexes := solarTermPalaceIndex[solarTerm]
	// 节气所在的宫索引, 局的宫索引
	solarTermPalaceIndex, roundPalaceIndex := palaceIndexes[0], palaceIndexes[1].OffsetBy(int(pentad)*6)

	return leadingHour, escaping, solarTermPalaceIndex, roundPalaceIndex
}

// makeBasePlate 创建天盘地盘的原始位
func makeBasePlate(pIdx component.PalaceIndex, escaping component.Escaping) (component.Palaces, error) {
	return component.NewPalace(
		[9]int{
			int(component.QiYiEnum.Wu), int(component.QiYiEnum.Ji), int(component.QiYiEnum.Geng),
			int(component.QiYiEnum.Xin), int(component.QiYiEnum.Ren), int(component.QiYiEnum.Gui),
			int(component.QiYiEnum.Ding), int(component.QiYiEnum.Bing), int(component.QiYiEnum.Yi),
		},
		pIdx,
		escaping == component.YangEscaping,
	)
}
