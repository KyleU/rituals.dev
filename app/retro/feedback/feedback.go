package feedback

import (
	"fmt"
	"net/url"
	"path"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/lib/svc"
	"github.com/kyleu/rituals/app/util"
)

const DefaultRoute = "/admin/db/retro/feedback"

func Route(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return path.Join(paths...)
}

var _ svc.Model = (*Feedback)(nil)

type Feedback struct {
	ID       uuid.UUID  `json:"id,omitempty"`
	RetroID  uuid.UUID  `json:"retroID,omitempty"`
	Idx      int        `json:"idx,omitempty"`
	UserID   uuid.UUID  `json:"userID,omitempty"`
	Category string     `json:"category,omitempty"`
	Content  string     `json:"content,omitempty"`
	HTML     string     `json:"html,omitempty"`
	Created  time.Time  `json:"created,omitempty"`
	Updated  *time.Time `json:"updated,omitempty"`
}

func New(id uuid.UUID) *Feedback {
	return &Feedback{ID: id}
}

func (f *Feedback) Clone() *Feedback {
	return &Feedback{f.ID, f.RetroID, f.Idx, f.UserID, f.Category, f.Content, f.HTML, f.Created, f.Updated}
}

func (f *Feedback) String() string {
	return f.ID.String()
}

func (f *Feedback) TitleString() string {
	return f.String()
}

func Random() *Feedback {
	return &Feedback{
		ID:       util.UUID(),
		RetroID:  util.UUID(),
		Idx:      util.RandomInt(10000),
		UserID:   util.UUID(),
		Category: util.RandomString(12),
		Content:  util.RandomString(12),
		HTML:     "<h3>" + util.RandomString(6) + "</h3>",
		Created:  util.TimeCurrent(),
		Updated:  util.TimeCurrentP(),
	}
}

//nolint:lll
func (f *Feedback) Strings() []string {
	return []string{f.ID.String(), f.RetroID.String(), fmt.Sprint(f.Idx), f.UserID.String(), f.Category, f.Content, f.HTML, util.TimeToFull(&f.Created), util.TimeToFull(f.Updated)}
}

func (f *Feedback) ToCSV() ([]string, [][]string) {
	return FieldDescs.Keys(), [][]string{f.Strings()}
}

func (f *Feedback) WebPath(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return path.Join(append(paths, url.QueryEscape(f.ID.String()))...)
}

func (f *Feedback) ToData() []any {
	return []any{f.ID, f.RetroID, f.Idx, f.UserID, f.Category, f.Content, f.HTML, f.Created, f.Updated}
}

var FieldDescs = util.FieldDescs{
	{Key: "id", Title: "ID", Description: "", Type: "uuid"},
	{Key: "retroID", Title: "Retro ID", Description: "", Type: "uuid"},
	{Key: "idx", Title: "Idx", Description: "", Type: "int"},
	{Key: "userID", Title: "User ID", Description: "", Type: "uuid"},
	{Key: "category", Title: "Category", Description: "", Type: "string"},
	{Key: "content", Title: "Content", Description: "", Type: "string"},
	{Key: "html", Title: "HTML", Description: "", Type: "string"},
	{Key: "created", Title: "Created", Description: "", Type: "timestamp"},
	{Key: "updated", Title: "Updated", Description: "", Type: "timestamp"},
}
