package api

import (
	"bz-lib/app"
	"time"
)

const LIST_CLIENTS_V1 = "bz-client.ListClients_v1"

type ListClientsReq_v1 struct {
	Meta Meta
}

type ListClientsResp_v1 struct {
	Meta   Meta
	Data   []ClientData
	Status Status
}

var timeout time.Duration

func ListClients_v1(app *app.AppContext) (ListClientsResp_v1, error) {
	if timeout == 0 {
		timeout, _ = time.ParseDuration(app.Conf.Bus.Timeout)
	}
	req := &ListClientsReq_v1{}
	var resp ListClientsResp_v1
	err := app.Nats.Request(LIST_CLIENTS_V1, req, &resp, timeout)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
