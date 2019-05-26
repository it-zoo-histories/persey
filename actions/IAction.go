package actions

import "persey/payloads"

/*IAction - интерфейс интеграционного сервиса*/
type IAction interface {
	Action(*payloads.NewRequestPayload) error
}
