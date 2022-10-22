package main

import "time"

type Question struct {
	Id         int64
	content    string
	CreatedAt  time.Time
	DeletedAt  time.Time
	UpdateAt   time.Time
	comments   string
	like       bool
	collection bool
	forward    bool
}

//不知道是不是这么做
