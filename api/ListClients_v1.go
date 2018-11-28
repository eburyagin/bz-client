package api

const LIST_CLIENTS_V1 = "bz.ListClients_v1"

type ListClientsReq_v1 struct {
	Meta Meta
}

type ListClientsResp_v1 struct {
	Meta   Meta
	Data   []ClientData
	Status Status
}
