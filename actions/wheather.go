package actions

import "persey/payloads"

/*WheatherAction - тип для работа с погодой*/
type WheatherAction struct {
}

/*Action - выполнить действия интеграционного модуля*/
func (actor *WheatherAction) Action(requestPayload *payloads.NewRequestPayload) error {
	// city := ""
	// time := ""

	// for _, value := range requestPayload.Entities {
	// 	switch value {
	// 	case "cities":
	// 		city = value.Value
	// 		break
	// 	case "time":
	// 		time = value.Value
	// 		break
	// 	default:
	// 		break
	// 	}
	// }

	return nil

}

func (actor *WheatherAction) getWheater(city, time string) (string, error) {
	return "not implemented", nil
}
