package routes

import (
	"encoding/json"
	"net/http"
	"persey/dao"
	"persey/enhancer"
	"persey/payloads"

	"github.com/gorilla/mux"
)

/*EntryRoute - роутер для обработки входящих запросов*/
type EntryRoute struct {
	// ActionAvailables
	StatisticDAO *dao.StatisticDao
	EResponser   *enhancer.Responser
}

const (
	nameServer           = "[PERSEY: ROUTES]: "
	handleCheckIntegrate = "/action_check"
)

func (route *EntryRoute) hanldeNewRequest(w http.ResponseWriter, r *http.Request) {
	var payload payloads.NewRequestPayload

	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		route.EResponser.ResponseWithError(w, r, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"error":   err.Error(),
			"context": "Persey.EntryRoute",
		})
		return
	}

}

/*SettingRoute - конфигурирование маршрута*/
func (route *EntryRoute) SettingRoute(da *dao.StatisticDao, router *mux.Router) *mux.Router {
	router.HandleFunc(handleCheckIntegrate, route.hanldeNewRequest).Methods("POST")
	route.StatisticDAO = da
	route.EResponser = &enhancer.Responser{}

	return router
}
