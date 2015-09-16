package chipmunkdb

import (
	"time"
)

type DB struct{}

func (db *DB) AddTimeValue(key string, timestamp time.Time, value float32) {
}

type TimeValue struct {
	Timestamp time.Time
	Value     float32
}

func (db *DB) AddPack(key string, deltaTime uint8, values []TimeValue) {
}
