package payloads

/*EntityPayload - сущность, которую распарсил NLU логики бота*/
type EntityPayload struct {
	EntityName  string `json:"entity_name"`
	EntityValue string `json:"entity_value"`
}
