package serv

import (
	"fmt"
	"testing"
	"time"

	"bz-client/api"
	"bz-client/internal/cfg"

	nats "github.com/nats-io/go-nats"
)

func Test_ListClients_v1(t *testing.T) {
	cf := "../../config_test.json"
	config, _ := cfg.Load(cf)
	Start_ListClients_v1(config)

	nc, _ := nats.Connect(config.Bus.Urls)
	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer c.Close()

	req := &api.ListClientsReq_v1{
		api.Meta{Ref: "3842793847239478"},
	}
	var resp api.ListClientsResp_v1
	err := c.Request(api.LIST_CLIENTS_V1, req, &resp, 100*time.Millisecond)
	if err != nil {
		t.Error("Request", err)
	}
	if len(resp.Data) != 3 {
		t.Errorf("Response %+v", resp)
	}

}

func Benchmark_ListClients_v1(b *testing.B) {
	cf := "../../config_test.json"
	config, _ := cfg.Load(cf)
	Start_ListClients_v1(config)

	nc, _ := nats.Connect(config.Bus.Urls)
	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer c.Close()

	for i := 0; i < b.N; i++ {
		req := &api.ListClientsReq_v1{
			api.Meta{Ref: fmt.Sprint(i)},
		}
		var resp api.ListClientsResp_v1
		c.Request(api.LIST_CLIENTS_V1, req, &resp, 100*time.Millisecond)
	}
}
