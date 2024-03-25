package metartools

const KNOT_CONVERSION float32 = 1.15078

type WindDirectional string

const (
	WIND_N   = "N"
	WIND_S   = "S"
	WIND_E   = "E"
	WIND_W   = "W"
	WIND_NE  = "NE"
	WIND_SE  = "SE"
	WIND_SW  = "SW"
	WIND_NW  = "NW"
	WIND_NNE = "NNE"
	WIND_ENE = "ENE"
	WIND_SSE = "SSE"
	WIND_ESE = "ESE"
	WIND_WSW = "WSW"
	WIND_SSW = "SSW"
	WIND_NNW = "NNW"
	WIND_WNW = "WNW"
)

func WindAngle(direction string) int {
	switch direction {
	case WIND_N:
		return 0
	case WIND_E:
		return 90
	case WIND_S:
		return 180
	case WIND_W:
		return 270
	case WIND_NE:
		return 45
	case WIND_SE:
		return 135
	case WIND_SW:
		return 225
	case WIND_NW:
		return 315
	case WIND_NNE:
		return 22
	case WIND_ENE:
		return 67
	case WIND_ESE:
		return 112
	case WIND_SSE:
		return 157
	case WIND_SSW:
		return 202
	case WIND_WSW:
		return 247
	case WIND_WNW:
		return 292
	case WIND_NNW:
		return 337
	}
	return 0
}

func WindBarbImageIndex(s float32) int {
	switch speed := s; {
	case speed > 0 && speed < 5:
		return 26
	case speed > 125:
		return 25
	case speed > 120:
		return 24
	case speed > 115:
		return 23
	case speed > 110:
		return 22
	case speed > 105:
		return 21
	case speed > 100:
		return 20
	case speed > 95:
		return 19
	case speed > 90:
		return 18
	case speed > 85:
		return 17
	case speed > 80:
		return 16
	case speed > 75:
		return 15
	case speed > 70:
		return 14
	case speed > 65:
		return 13
	case speed > 60:
		return 12
	case speed > 55:
		return 11
	case speed > 50:
		return 10
	case speed > 45:
		return 9
	case speed > 40:
		return 8
	case speed > 35:
		return 7
	case speed > 30:
		return 6
	case speed > 25:
		return 5
	case speed > 20:
		return 4
	case speed > 15:
		return 3
	case speed > 10:
		return 2
	case speed > 5:
		return 1
	}
	return 26
}
