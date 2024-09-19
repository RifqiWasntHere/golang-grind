package model

import "time"

type UserComplex struct {
	userId     string
	userName   string
	email      string
	marrie     bool
	occupation string
	balance    int32
	score      float32
	birthdate  time.Time
	createdAt  time.Time
}
