package models

import (
	"time"
)

type Device struct {
	ID       int       `json:"id"`
	Device   string    `json:"device"`
	Unit     string    `json:"unit"`
	Value    float64   `json:"value"`
	Datetime time.Time `json:"datetime"`
}
