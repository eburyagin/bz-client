package api

import "time"

type ClientData struct {
	ID         int
	Code       string
	Name       string
	Note       string
	Status     string
	CreateTime time.Time
}
