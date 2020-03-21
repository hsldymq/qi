package component

import (
	"github.com/hsldymq/go-chinese-calendar/solar"
)

// Escaping 遁
type Escaping int

// 阳遁阴遁
const (
	YangEscaping Escaping = 0
	YinEscaping  Escaping = 1
)

// solarTermEscapingMap 节气与阴阳遁的对应关系
var solarTermEscapingMap = map[solar.SolarTerm]Escaping{
	solar.SolarTermEnum.TheBeginningOfSummer:  YinEscaping,
	solar.SolarTermEnum.LesserFullnessOfGrain: YinEscaping,
	solar.SolarTermEnum.GrainInBeard:          YinEscaping,
	solar.SolarTermEnum.TheSummerSolstice:     YinEscaping,
	solar.SolarTermEnum.LesserHeat:            YinEscaping,
	solar.SolarTermEnum.GreaterHeat:           YinEscaping,
	solar.SolarTermEnum.TheBeginningOfAutumn:  YinEscaping,
	solar.SolarTermEnum.TheEndOfHeat:          YinEscaping,
	solar.SolarTermEnum.WhiteDew:              YinEscaping,
	solar.SolarTermEnum.TheAutumnEquinox:      YinEscaping,
	solar.SolarTermEnum.ColdDew:               YinEscaping,
	solar.SolarTermEnum.FrostsDescent:         YinEscaping,
	solar.SolarTermEnum.TheBeginningOfWinter:  YangEscaping,
	solar.SolarTermEnum.LesserSnow:            YangEscaping,
	solar.SolarTermEnum.GreaterSnow:           YangEscaping,
	solar.SolarTermEnum.TheWinterSolstice:     YangEscaping,
	solar.SolarTermEnum.LesserCold:            YangEscaping,
	solar.SolarTermEnum.GreaterCold:           YangEscaping,
	solar.SolarTermEnum.TheBeginningOfSpring:  YangEscaping,
	solar.SolarTermEnum.RainWater:             YangEscaping,
	solar.SolarTermEnum.TheWakingOfInsects:    YangEscaping,
	solar.SolarTermEnum.TheSpringEquinox:      YangEscaping,
	solar.SolarTermEnum.PureBrightness:        YangEscaping,
	solar.SolarTermEnum.GrainRain:             YangEscaping,
}

// SolarTermEscaping 根据节气返回对应的遁
func SolarTermEscaping(term solar.SolarTerm) Escaping {
	return solarTermEscapingMap[term]
}
