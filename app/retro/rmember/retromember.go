// Package rmember - Content managed by Project Forge, see [projectforge.md] for details.
package rmember

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

type PK struct {
	RetroID uuid.UUID `json:"retroID,omitempty"`
	UserID  uuid.UUID `json:"userID,omitempty"`
}

func (p *PK) String() string {
	return fmt.Sprintf("%v::%v", p.RetroID, p.UserID)
}

type RetroMember struct {
	RetroID uuid.UUID         `json:"retroID,omitempty"`
	UserID  uuid.UUID         `json:"userID,omitempty"`
	Name    string            `json:"name,omitempty"`
	Picture string            `json:"picture,omitempty"`
	Role    enum.MemberStatus `json:"role,omitempty"`
	Created time.Time         `json:"created,omitempty"`
	Updated *time.Time        `json:"updated,omitempty"`
}

func New(retroID uuid.UUID, userID uuid.UUID) *RetroMember {
	return &RetroMember{RetroID: retroID, UserID: userID}
}

func (r *RetroMember) Clone() *RetroMember {
	return &RetroMember{r.RetroID, r.UserID, r.Name, r.Picture, r.Role, r.Created, r.Updated}
}

func (r *RetroMember) String() string {
	return fmt.Sprintf("%s::%s", r.RetroID.String(), r.UserID.String())
}

func (r *RetroMember) TitleString() string {
	return r.RetroID.String() + " / " + r.Name
}

func (r *RetroMember) ToPK() *PK {
	return &PK{
		RetroID: r.RetroID,
		UserID:  r.UserID,
	}
}

func Random() *RetroMember {
	return &RetroMember{
		RetroID: util.UUID(),
		UserID:  util.UUID(),
		Name:    util.RandomString(12),
		Picture: "https://" + util.RandomString(6) + ".com/" + util.RandomString(6),
		Role:    enum.AllMemberStatuses.Random(),
		Created: util.TimeCurrent(),
		Updated: util.TimeCurrentP(),
	}
}

func (r *RetroMember) WebPath() string {
	return "/admin/db/retro/member/" + r.RetroID.String() + "/" + r.UserID.String()
}

func (r *RetroMember) ToData() []any {
	return []any{r.RetroID, r.UserID, r.Name, r.Picture, r.Role, r.Created, r.Updated}
}
