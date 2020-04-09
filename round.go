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

// Round 局
type Round struct {
	DutyStar             component.Star    // 值符
	DutyDoor             component.Door    // 值使
	QiYiTerrestrialPlate component.Palaces // 地盘三奇六仪
	QiYiCelestialPlate   component.Palaces // 天盘三奇六仪
	StarCelestialPlate   component.Palaces // 九星天盘
	HumanPlate           component.Palaces // 人盘
	GodPlate             component.Palaces // 神盘
}

// RoundInfo 起局需要的相关信息
type RoundInfo struct {
	Escaping         component.Escaping        // 阳/阴遁
	RoundPalaceIndex component.PalaceIndex     // 此局所在宫索引
	Yuan             component.Yuan            // 上/中/下元
	SexagenaryHour   sexagenary.SexagenaryTerm // 干支时辰
}

func (ri RoundInfo) Validate() error {
	if ri.Yuan.IsValid() {
		return fmt.Errorf("invalid yuan: %d", ri.Yuan)
	}
	if ri.RoundPalaceIndex.IsValid() {
		return fmt.Errorf("invalid round palace index: %d", ri.RoundPalaceIndex)
	}
	if ri.SexagenaryHour.IsValid() {
		return fmt.Errorf("invalid sexagenary hour: %d", ri.SexagenaryHour.Index())
	}
	if ri.Escaping.IsValid() {
		return fmt.Errorf("invalid escaping: %d", ri.Escaping)
	}

	return nil
}

func GenerateRound(roundInfo RoundInfo) (Round, error) {
	if err := roundInfo.Validate(); err != nil {
		return Round{}, err
	}

	// 旬首
	hourLeader := component.NewHourLeader(roundInfo.SexagenaryHour)

	// 天盘三奇六仪
	qiYiCelestialPlate, err := component.NewQiYiPlate(roundInfo.RoundPalaceIndex, roundInfo.Escaping)
	if err != nil {
		return Round{}, err
	}

	// 地盘三奇六仪
	qiYiTerrestrialPlate, err := component.NewQiYiPlate(roundInfo.RoundPalaceIndex, roundInfo.Escaping)
	if err != nil {
		return Round{}, err
	}

	// 旬首所在的宫位索引
	hourLeaderPalaceIndex := qiYiCelestialPlate.FindPalaceIndex(hourLeader.SixYi().Value())

	// 获得值符(值星), 旋转后的天盘三奇六仪, 旋转后的天盘九星
	dutyStar, qiYiCelestialPlate, starCelestialPlate := rotateStarCelestialPlate(hourLeaderPalaceIndex, roundInfo.SexagenaryHour, qiYiCelestialPlate)

	// 获得值使(值门), 旋转后的人盘
	dutyDoor, humanPlate := rotateHumanPlate(roundInfo, hourLeader, hourLeaderPalaceIndex)

	godPlate := rotateGodPlate(dutyStar, starCelestialPlate, roundInfo.Escaping)

	return Round{
		DutyStar:             dutyStar,
		DutyDoor:             dutyDoor,
		QiYiTerrestrialPlate: qiYiTerrestrialPlate,
		QiYiCelestialPlate:   qiYiCelestialPlate,
		StarCelestialPlate:   starCelestialPlate,
		HumanPlate:           humanPlate,
		GodPlate:             godPlate,
	}, nil
}

// rotateStarCelestialPlate 转动九星天盘
func rotateStarCelestialPlate(
	hourLeaderPalaceIndex component.PalaceIndex,
	hour sexagenary.SexagenaryTerm,
	qiYiCelestialPlate component.Palaces,
) (component.Star, component.Palaces, component.Palaces) {

	// 从旬首所在宫位要转到的目标宫位
	rotatedCelestialPalaceIndex := hourLeaderPalaceIndex
	if hour.CelestialStem != sexagenary.CelestialStemEnum.Jia {
		rotatedCelestialPalaceIndex = qiYiCelestialPlate.FindPalaceIndex(component.QiYi(hour.CelestialStem).Value())
	}
	fmt.Println(hourLeaderPalaceIndex)
	// 指定实际的天盘旋转宫位, 由于五宫寄于二宫, 所以这里需要进行调整
	from := hourLeaderPalaceIndex
	if from == component.FifthPalace {
		from = component.SecondPalace
	}
	to := rotatedCelestialPalaceIndex
	if to == component.FifthPalace {
		to = component.SecondPalace
	}

	rotatedDistance := from.RoundDistance(to)
	starCelestialPlate := component.NewOriginStarPlate()
	rotatedQiYiCelestialPlate := qiYiCelestialPlate.RotateValues(rotatedDistance)
	rotatedStarCelestialPlate := starCelestialPlate.RotateValues(rotatedDistance)
	dutyStar := hourLeaderPalaceIndex.OriginalStar()

	return dutyStar, rotatedQiYiCelestialPlate, rotatedStarCelestialPlate
}

// rotateHumanPlate 转动人盘
func rotateHumanPlate(
	info RoundInfo,
	hourLeader component.HourLeader,
	hourLeaderPalaceIndex component.PalaceIndex,
) (component.Door, component.Palaces) {

	endPalaceIndex := hourLeaderPalaceIndex
	// 八门按旬首到干支时, 根据阳遁或阴遁进行顺飞或逆飞
	for st := hourLeader.SexagenaryTerm(); st.Index() != info.SexagenaryHour.Index(); st = st.Next() {
		if info.Escaping == component.YangEscaping {
			endPalaceIndex = endPalaceIndex.Next()
		} else {
			endPalaceIndex = endPalaceIndex.Prev()
		}
	}
	// TODO 如果飞入五宫怎么办?
	if endPalaceIndex == component.FifthPalace {
		// endPalaceIndex = component.SecondPalace
	}

	distance := hourLeaderPalaceIndex.RoundDistance(endPalaceIndex)
	dutyDoor := hourLeaderPalaceIndex.OriginalDoor()
	humanPlate := component.NewOriginalHumanPlate()

	return dutyDoor, humanPlate.RotateValues(distance)
}

// rotateGodPlate 旋转神盘
// 通过旋转将直符对准九星天盘值符
// 原始神盘直符位于一宫
func rotateGodPlate(dutyStar component.Star, starCelestialPlate component.Palaces, escaping component.Escaping) component.Palaces {
	godPlate := component.NewOriginalGodPlate(escaping)
	rotateTo := starCelestialPlate.FindPalaceIndex(dutyStar.Value())

	return godPlate.RotateValues(component.FirstPalace.RoundDistance(rotateTo))
}
