package entities

import (
	"time"

	"github.com/uptrace/bun"
)

type Credentials struct {
	bun.BaseModel `bun:"table:credentials,alias:creds"`
	Name          string `bun:",pk" json:"name"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	Token         string `json:"token"`

	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"createdAt"`
}

func (a Credentials) UsingBasic() bool {
	if a.Username != "" && a.Password != "" {
		return true
	}
	return false
}
