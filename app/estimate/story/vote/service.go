// Package vote - Content managed by Project Forge, see [projectforge.md] for details.
package vote

import (
	"github.com/kyleu/rituals/app/lib/database"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/lib/svc"
)

var _ svc.Service[*Vote, Votes] = (*Service)(nil)

type Service struct {
	db *database.Service
}

func NewService(db *database.Service) *Service {
	filter.AllowedColumns["vote"] = columns
	return &Service{db: db}
}

func filters(orig *filter.Params) *filter.Params {
	return orig.Sanitize("vote", &filter.Ordering{Column: "created"})
}
