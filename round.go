package qimen

import (
	"fmt"

	"github.com/hsldymq/go-chinese-calendar/sexagenary"
	"github.com/hsldymq/go-chinese-calendar/solar"
	"github.com/hsldymq/qimen/component"
)

// solarTermPalaceIndex 每个节气坐落的宫索引,以及对应的上元局起宫索引
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
	HourLeader           component.HourLeader // 旬首
	DutyStar             component.Star       // 值符
	DutyDoor             component.Door       // 值使
	QiYiTerrestrialPlate component.Palaces    // 地盘三奇六仪
	QiYiCelestialPlate   component.Palaces    // 天盘三奇六仪
	StarCelestialPlate   component.Palaces    // 九星天盘
	DoorCelestialPlate   component.Palaces    // 八门天盘
	GodCelestialPlate    component.Palaces    // 八神天盘
}

// RoundParams 起局参数
type RoundParams struct {
	Escaping             component.Escaping        // 阳/阴遁
	SolarTermPalaceIndex component.PalaceIndex     // 节气所在宫索引
	RoundPalaceIndex     component.PalaceIndex     // 此局所在宫索引
	SexagenaryHour       sexagenary.SexagenaryTerm // 干支时辰
}

func (ri RoundParams) Validate() error {
	if !ri.RoundPalaceIndex.IsValid() {
		return fmt.Errorf("invalid round palace index: %d", ri.RoundPalaceIndex)
	}
	if !ri.SexagenaryHour.IsValid() {
		return fmt.Errorf("invalid sexagenary hour: %d", ri.SexagenaryHour.Index())
	}
	if !ri.Escaping.IsValid() {
		return fmt.Errorf("invalid escaping: %d", ri.Escaping)
	}

	return nil
}

// NewRoundParams 创建起局参数
// term: 节气
// dayNo: 该节气的第几日(从1计数)
// sexagenaryHour: 干支时
func NewRoundParams(term solar.SolarTerm, dayNo int, sexagenaryHour sexagenary.SexagenaryTerm) (RoundParams, error) {
	palaceIndexes, ok := solarTermPalaceIndex[term]
	if !ok {
		return RoundParams{}, fmt.Errorf("invalid solar term: %d", term)
	}
	if dayNo < 0 {
		return RoundParams{}, fmt.Errorf("invalid solar term day number: %d", dayNo)
	}

	yuan := dayNo / 5
	if yuan > 2 {
		yuan = 2
	}
	roundPalaceIndex := palaceIndexes[1].OffsetBy(yuan * 6)
	return RoundParams{
		Escaping:             component.SolarTermEscaping(term),
		SolarTermPalaceIndex: palaceIndexes[0],
		RoundPalaceIndex:     roundPalaceIndex,
		SexagenaryHour:       sexagenaryHour,
	}, nil
}

// GenerateRoundV1 起局
func GenerateRoundV1(roundParams RoundParams) (Round, error) {
	if err := roundParams.Validate(); err != nil {
		return Round{}, err
	}

	// 地盘三奇六仪
	qiYiTerrestrialPlate, err := component.NewQiYiPlate(roundParams.RoundPalaceIndex, roundParams.Escaping)
	if err != nil {
		return Round{}, err
	}

	// 天盘三奇六仪
	qiYiCelestialPlate, err := component.NewQiYiPlate(roundParams.RoundPalaceIndex, roundParams.Escaping)
	if err != nil {
		return Round{}, err
	}

	// 旬首, 旬首所在的宫位索引
	hourLeader, hourLeaderPalaceIndex := hourLeader(roundParams.SexagenaryHour, qiYiCelestialPlate)

	// 获得值符(值星), 旋转后的天盘三奇六仪, 旋转后的天盘九星
	dutyStar, qiYiCelestialPlate, starCelestialPlate := rotateStarCelestialPlateV1(hourLeaderPalaceIndex, roundParams.SexagenaryHour, qiYiCelestialPlate)

	// 获得值使(值门), 八门天盘
	dutyDoor, doorCelestialPlate := makeHumanCelestialPlateV1(roundParams, hourLeader, hourLeaderPalaceIndex)

	// 生成八神天盘
	godCelestialPlate := makeGodCelestialPlate(dutyStar, starCelestialPlate, roundParams.Escaping)

	return Round{
		HourLeader:           hourLeader,
		DutyStar:             dutyStar,
		DutyDoor:             dutyDoor,
		QiYiTerrestrialPlate: qiYiTerrestrialPlate,
		QiYiCelestialPlate:   qiYiCelestialPlate,
		StarCelestialPlate:   starCelestialPlate,
		DoorCelestialPlate:   doorCelestialPlate,
		GodCelestialPlate:    godCelestialPlate,
	}, nil
}

func GenerateRoundV2(roundParams RoundParams) (Round, error) {
	if err := roundParams.Validate(); err != nil {
		return Round{}, err
	}

	// 地盘三奇六仪
	qiYiTerrestrialPlate, err := component.NewQiYiPlate(roundParams.RoundPalaceIndex, roundParams.Escaping)
	if err != nil {
		return Round{}, err
	}

	// 天盘三奇六仪
	qiYiCelestialPlate, err := component.NewQiYiPlate(roundParams.RoundPalaceIndex, roundParams.Escaping)
	if err != nil {
		return Round{}, err
	}

	// 旬首, 旬首所在的宫位索引
	hourLeader, hourLeaderPalaceIndex := hourLeader(roundParams.SexagenaryHour, qiYiCelestialPlate)

	// 获得值符(值星), 旋转后的天盘三奇六仪, 旋转后的天盘九星
	dutyStar, qiYiCelestialPlate, starCelestialPlate := rotateStarCelestialPlateV2(hourLeaderPalaceIndex, roundParams, qiYiCelestialPlate)

	// 获得值使(值门), 八门天盘
	dutyDoor, humanCelestialPlate := makeHumanCelestialPlateV2(roundParams, roundParams.SolarTermPalaceIndex, hourLeader, hourLeaderPalaceIndex)

	// 生成八神天盘
	godCelestialPlate := makeGodCelestialPlate(dutyStar, starCelestialPlate, roundParams.Escaping)

	return Round{
		HourLeader:           hourLeader,
		DutyStar:             dutyStar,
		DutyDoor:             dutyDoor,
		QiYiTerrestrialPlate: qiYiTerrestrialPlate,
		QiYiCelestialPlate:   qiYiCelestialPlate,
		StarCelestialPlate:   starCelestialPlate,
		DoorCelestialPlate:   humanCelestialPlate,
		GodCelestialPlate:    godCelestialPlate,
	}, nil
}

// rotateStarCelestialPlateV1 转动九星天盘
// 当时干位于五宫时,需寄于二宫
func rotateStarCelestialPlateV1(
	hourLeaderPalaceIndex component.PalaceIndex,
	hour sexagenary.SexagenaryTerm,
	qiYiCelestialPlate component.Palaces,
) (component.Star, component.Palaces, component.Palaces) {
	// 从旬首所在宫位要转到的目标宫位
	rotatedToCelestialPalaceIndex := hourLeaderPalaceIndex
	if hour.CelestialStem != sexagenary.CelestialStemEnum.Jia {
		rotatedToCelestialPalaceIndex = qiYiCelestialPlate.FindPalaceIndex(component.QiYi(hour.CelestialStem).Value())
	}
	// 指定实际的天盘旋转宫位, 由于五宫寄于二宫, 所以这里需要进行调整
	from := hourLeaderPalaceIndex
	if from == component.FifthPalace {
		from = component.SecondPalace
	}
	to := rotatedToCelestialPalaceIndex
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

// rotateStarCelestialPlateV2 转动九星天盘
// 当时干位于五宫时,需寄于节气所在宫
func rotateStarCelestialPlateV2(
	hourLeaderPalaceIndex component.PalaceIndex,
	roundParams RoundParams,
	qiYiCelestialPlate component.Palaces,
) (component.Star, component.Palaces, component.Palaces) {
	// 从旬首所在宫位要转到的目标宫位
	rotatedToCelestialPalaceIndex := hourLeaderPalaceIndex
	if roundParams.SexagenaryHour.CelestialStem != sexagenary.CelestialStemEnum.Jia {
		rotatedToCelestialPalaceIndex = qiYiCelestialPlate.FindPalaceIndex(component.QiYi(roundParams.SexagenaryHour.CelestialStem).Value())
	}
	// 指定实际的天盘旋转宫位, 由于五宫寄于节气所在的宫, 所以这里需要进行调整
	from := hourLeaderPalaceIndex
	if from == component.FifthPalace {
		from = roundParams.SolarTermPalaceIndex
	}
	to := rotatedToCelestialPalaceIndex
	if to == component.FifthPalace {
		to = roundParams.SolarTermPalaceIndex
	}

	rotatedDistance := from.RoundDistance(to)
	starCelestialPlate := component.NewOriginStarPlate()
	rotatedQiYiCelestialPlate := qiYiCelestialPlate.RotateValues(rotatedDistance)
	rotatedStarCelestialPlate := starCelestialPlate.RotateValues(rotatedDistance)
	dutyStar := hourLeaderPalaceIndex.OriginalStar()

	return dutyStar, rotatedQiYiCelestialPlate, rotatedStarCelestialPlate
}

// makeHumanCelestialPlateV1 生成八门天盘
// 当时飞临五宫时,需寄于二宫
func makeHumanCelestialPlateV1(
	info RoundParams,
	hourLeader component.HourLeader,
	hourLeaderPalaceIndex component.PalaceIndex,
) (component.Door, component.Palaces) {
	fromPalaceIndex := hourLeaderPalaceIndex
	toPalaceIndex := hourLeaderPalaceIndex
	// 八门按旬首到干支时, 根据阳遁或阴遁进行顺飞或逆飞
	for st := hourLeader.SexagenaryTerm(); st.Index() != info.SexagenaryHour.Index(); st = st.Next() {
		if info.Escaping == component.YangEscaping {
			toPalaceIndex = toPalaceIndex.Next()
		} else {
			toPalaceIndex = toPalaceIndex.Prev()
		}
	}

	if fromPalaceIndex == component.FifthPalace {
		fromPalaceIndex = component.SecondPalace
	}

	if toPalaceIndex == component.FifthPalace {
		toPalaceIndex = component.SecondPalace
	}

	distance := fromPalaceIndex.RoundDistance(toPalaceIndex)
	dutyDoor := fromPalaceIndex.OriginalDoor()
	humanPlate := component.NewOriginalHumanPlate()

	return dutyDoor, humanPlate.RotateValues(distance)
}

// makeHumanCelestialPlateV2 生成八门天盘
// 当时飞临五宫时,需寄于节气所在宫
func makeHumanCelestialPlateV2(
	info RoundParams,
	solarTermPalaceIndex component.PalaceIndex,
	hourLeader component.HourLeader,
	hourLeaderPalaceIndex component.PalaceIndex,
) (component.Door, component.Palaces) {
	fromPalaceIndex := hourLeaderPalaceIndex
	toPalaceIndex := hourLeaderPalaceIndex
	// 八门按旬首到干支时, 根据阳遁或阴遁进行顺飞或逆飞
	for st := hourLeader.SexagenaryTerm(); st.Index() != info.SexagenaryHour.Index(); st = st.Next() {
		if info.Escaping == component.YangEscaping {
			toPalaceIndex = toPalaceIndex.Next()
		} else {
			toPalaceIndex = toPalaceIndex.Prev()
		}
	}

	if fromPalaceIndex == component.FifthPalace {
		fromPalaceIndex = solarTermPalaceIndex
	}

	if toPalaceIndex == component.FifthPalace {
		toPalaceIndex = solarTermPalaceIndex
	}

	distance := fromPalaceIndex.RoundDistance(toPalaceIndex)
	dutyDoor := fromPalaceIndex.OriginalDoor()
	humanPlate := component.NewOriginalHumanPlate()

	return dutyDoor, humanPlate.RotateValues(distance)
}

// makeGodCelestialPlate 生成八神天盘
func makeGodCelestialPlate(
	dutyStar component.Star,
	starCelestialPlate component.Palaces,
	escaping component.Escaping,
) component.Palaces {
	// 通过旋转将值符对准九星天盘值符
	// 原始神盘值符位于一宫
	godCelestialPlate := component.NewOriginalGodPlate(escaping)
	ds := dutyStar
	if ds == component.StarEnum.TianQin {
		ds = component.StarEnum.TianRui
	}

	from := godCelestialPlate.FindPalaceIndex(int(component.GodEnum.ZhiFu))
	rotateTo := starCelestialPlate.FindPalaceIndex(ds.Value())

	return godCelestialPlate.RotateValues(from.RoundDistance(rotateTo))
}

func hourLeader(hour sexagenary.SexagenaryTerm, qiYiCelestialPlate component.Palaces) (component.HourLeader, component.PalaceIndex) {
	lh := component.NewHourLeader(hour)

	return lh, qiYiCelestialPlate.FindPalaceIndex(lh.SixYi().Value())
}
