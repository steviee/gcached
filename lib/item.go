package lib

import "time"

type Item struct {
	Key         string    `json:"key"`
	Value       string    `json:"value"`
	CreatedAt   time.Time `json:"created"`
	ProlongedAt time.Time `json:"prolongedAt"`
	TTL         int       `json:"ttl"`
}
