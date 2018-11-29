package serv

import (
	"database/sql"
	"fmt"
	"log"

	"bz-client/api"
	"bz-lib/cfg"

	_ "github.com/lib/pq"

	nats "github.com/nats-io/go-nats"
)

func Start_ListClients_v1(config *cfg.Config) error {
	log.Println("Запускаю ", api.LIST_CLIENTS_V1, "...")

	nc, err := nats.Connect(config.Bus.Urls)
	if err != nil {
		log.Fatalf("NATS Connect error!")
	}
	c, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Fatalf("NATS Connect JSON Encoder error!")
	}

	dbinfo := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		config.Datastore.User,
		config.Datastore.Password,
		config.Datastore.Addr,
		config.Datastore.Database)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatalf("sql.Open: %v", err)
	}

	_, err = c.Subscribe(api.LIST_CLIENTS_V1,
		func(subj, reply string, req *api.ListClientsReq_v1) {
			//			log.Printf("Запрос: %+v", req)
			var resp api.ListClientsResp_v1
			listClients_v1(db, req, &resp)
			c.Publish(reply, resp)
			//			log.Printf("Ответ: %+v", resp)
		})
	if err != nil {
		log.Fatalf("Subscribe error!")
	}

	log.Println(api.LIST_CLIENTS_V1, " стартовал.")
	return nil
}

func listClients_v1(db *sql.DB, req *api.ListClientsReq_v1, resp *api.ListClientsResp_v1) {

	resp.Meta = api.Meta{
		Ref: req.Meta.Ref,
	}

	rows, err1 := db.Query("SELECT * FROM clients;")
	if err1 != nil {
		resp.Status = api.Status{
			Code: 1,
			Note: fmt.Sprint(err1),
		}
		return
	}
	defer rows.Close()

	for rows.Next() {
		cd := new(api.ClientData)
		err2 := rows.Scan(&cd.ID, &cd.Code, &cd.Name, &cd.Note, &cd.Status, &cd.CreateTime)
		if err2 != nil {
			resp.Status = api.Status{
				Code: 2,
				Note: fmt.Sprint(err2),
			}
			return
		}

		resp.Data = append(resp.Data, *cd)
	}
	if err3 := rows.Err(); err3 != nil {
		resp.Status = api.Status{
			Code: 3,
			Note: fmt.Sprint(err1),
		}
		return
	}

	resp.Status = api.Status{
		Code: 0,
	}
	return
}
