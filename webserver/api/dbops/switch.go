package dbops

func SwitchKind(kind string)string{
	switch  kind {
	case "collect":
		return  `collect`
	case "favour":
		return  `favour`
	case "transmit":
		return  `transmit`
	}
	return  ""
}