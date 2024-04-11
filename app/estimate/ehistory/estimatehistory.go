// Package ehistory - Content managed by Project Forge, see [projectforge.md] for details.
package ehistory

import (
	"net/url"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/lib/svc"
	"github.com/kyleu/rituals/app/util"
)

var _ svc.Model = (*EstimateHistory)(nil)

type EstimateHistory struct {
	Slug         string    `json:"slug,omitempty"`
	EstimateID   uuid.UUID `json:"estimateID,omitempty"`
	EstimateName string    `json:"estimateName,omitempty"`
	Created      time.Time `json:"created,omitempty"`
}

func New(slug string) *EstimateHistory {
	return &EstimateHistory{Slug: slug}
}

func (e *EstimateHistory) Clone() *EstimateHistory {
	return &EstimateHistory{e.Slug, e.EstimateID, e.EstimateName, e.Created}
}

func (e *EstimateHistory) String() string {
	return e.Slug
}

func (e *EstimateHistory) TitleString() string {
	return e.String()
}

func Random() *EstimateHistory {
	return &EstimateHistory{
		Slug:         util.RandomString(12),
		EstimateID:   util.UUID(),
		EstimateName: util.RandomString(12),
		Created:      util.TimeCurrent(),
	}
}

func (e *EstimateHistory) Strings() []string {
	return []string{e.Slug, e.EstimateID.String(), e.EstimateName, util.TimeToFull(&e.Created)}
}

func (e *EstimateHistory) ToCSV() ([]string, [][]string) {
	return FieldDescs.Keys(), [][]string{e.Strings()}
}

func (e *EstimateHistory) WebPath() string {
	return "/admin/db/estimate/history/" + url.QueryEscape(e.Slug)
}

func (e *EstimateHistory) ToData() []any {
	return []any{e.Slug, e.EstimateID, e.EstimateName, e.Created}
}

var FieldDescs = util.FieldDescs{
	{Key: "slug", Title: "Slug", Description: "", Type: "string"},
	{Key: "estimateID", Title: "Estimate ID", Description: "", Type: "uuid"},
	{Key: "estimateName", Title: "Estimate Name", Description: "", Type: "string"},
	{Key: "created", Title: "Created", Description: "", Type: "timestamp"},
}
