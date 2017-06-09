package main

import "time"

type Bucket struct {
	Key        string          `json:"key"`
	Items      map[string]Item `json:"items"`
	CreatedAt  time.Time       `json:"created"`
	DefaultTTL int             `json:"default_ttl"`
}
