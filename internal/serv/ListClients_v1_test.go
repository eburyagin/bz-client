package serv

import (
	"fmt"
	"testing"
	"time"

	"bz-client/api"
	"bz-lib/app"
	"bz-lib/cfg"

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
		Meta: api.Meta{Ref: "3842793847239478"},
	}
	var resp api.ListClientsResp_v1
	err := c.Request(api.LIST_CLIENTS_V1, req, &resp, 1000*time.Millisecond)
	if err != nil {
		t.Error("Request", err)
	}
	if len(resp.Data) != 3 {
		t.Errorf("Response data len not eq 3! %+v", resp)
	}

}

func Api_ListClients_v1_test(t *testing.T) {
	app := new(app.AppContext)
	cf := "../config_test.json"
	app.Conf, _ = cfg.Load(cf)
	Start_ListClients_v1(app.Conf)

	nc, _ := nats.Connect(app.Conf.Bus.Urls)
	app.Nats, _ = nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer app.Nats.Close()

	resp, err := api.ListClients_v1(app)
	if err != nil {
		t.Error(err)
	}
	if len(resp.Data) != 3 {
		t.Errorf("Response data len not eq 3! %+v", resp)
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
			Meta: api.Meta{Ref: fmt.Sprint(i)},
		}
		var resp api.ListClientsResp_v1
		c.Request(api.LIST_CLIENTS_V1, req, &resp, 100*time.Millisecond)
	}
}
