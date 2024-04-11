// Package umember - Content managed by Project Forge, see [projectforge.md] for details.
package umember

import (
	"github.com/kyleu/rituals/app/lib/database"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/lib/svc"
)

var _ svc.Service[*StandupMember, StandupMembers] = (*Service)(nil)

type Service struct {
	db *database.Service
}

func NewService(db *database.Service) *Service {
	filter.AllowedColumns["umember"] = columns
	return &Service{db: db}
}

func filters(orig *filter.Params) *filter.Params {
	return orig.Sanitize("umember", &filter.Ordering{Column: "created"})
}
