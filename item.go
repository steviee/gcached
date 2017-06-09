package main

import "time"

type Item struct {
	Key         string    `json:"key"`
	Value       string    `json:"value"`
	CreatedAt   time.Time `json:"created"`
	ProlongedAt time.Time `json:"created"`
	TTL         int       `json:"ttl"`
}
