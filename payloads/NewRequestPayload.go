package payloads

/*NewRequestPayload - новый запрос на интеграцию с другими сервисами*/
type NewRequestPayload struct {
	ActionType int             `json:"action_type"`
	Entities   []EntityPayload `json:"entites"`
}
