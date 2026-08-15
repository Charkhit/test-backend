package tables

import "time"

type Staff struct {
	ID        uint       `json:"id"`
	Username  string     `json:"username"`
	Password  string     `json:"-"`
	Hospital  string     `json:"hospital"`
	LastLogin *time.Time `json:"last_login"`
}
