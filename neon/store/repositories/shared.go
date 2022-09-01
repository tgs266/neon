package repositories

import "github.com/uptrace/bun"

type Query struct {
	Query string
	Arg   any
}

func (q Query) Apply(call *bun.SelectQuery) *bun.SelectQuery {
	return call.Where(q.Query, q.Arg)
}
