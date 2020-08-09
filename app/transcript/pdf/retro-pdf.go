package pdf

import (
	"strings"

	"github.com/kyleu/npn/npncore"

	pdfgen "github.com/johnfercher/maroto/pkg/pdf"
	"github.com/kyleu/rituals.dev/app/member"
	"github.com/kyleu/rituals.dev/app/retro"
	"github.com/kyleu/rituals.dev/app/transcript"
	"github.com/kyleu/rituals.dev/app/util"
)

func renderRetro(rsp transcript.RetroResponse, m pdfgen.Maroto) string {
	hr(m)
	caption(rsp.Session.Title, m)
	detailRow(npncore.Title(npncore.KeyOwner), rsp.Members.GetName(rsp.Session.Owner), m)
	detailRow(npncore.PluralTitle(npncore.KeyCategory), strings.Join(rsp.Session.Categories, ", "), m)
	if rsp.Team != nil {
		detailRow(util.SvcTeam.Title, rsp.Team.Title, m)
	}
	if rsp.Sprint != nil {
		detailRow(util.SvcSprint.Title, rsp.Sprint.Title, m)
	}
	detailRow(npncore.Title(npncore.KeyCreated), npncore.ToDateString(&rsp.Session.Created), m)

	renderPermissionList(rsp.Permissions, m)
	renderMemberList(rsp.Members, m)
	renderFeedbackLists(rsp.Session.Categories, rsp.Feedback, rsp.Members, m)
	renderCommentList(rsp.Comments, rsp.Members, m, true)

	return rsp.Session.Slug
}

func renderRetroList(sessions retro.Sessions, members member.Entries, m pdfgen.Maroto) {
	if len(sessions) > 0 {
		hr(m)
		caption(util.SvcRetro.PluralTitle, m)
		cols := []string{npncore.Title(npncore.KeyOwner), npncore.Title(npncore.KeyTitle), npncore.Title(npncore.KeyCreated)}
		var data [][]string
		for _, s := range sessions {
			data = append(data, []string{members.GetName(s.Owner), s.Title, npncore.ToDateString(&s.Created)})
		}
		table(cols, data, []uint{3, 6, 3}, m)
	}
}

func renderFeedbackLists(categories []string, feedbacks retro.Feedbacks, members member.Entries, m pdfgen.Maroto) {
	for _, c := range categories {
		var fs retro.Feedbacks
		for _, f := range feedbacks {
			if f.Category == c {
				fs = append(fs, f)
			}
		}
		if len(fs) > 0 {
			hr(m)
			caption(c, m)
			cols := []string{npncore.Title(npncore.KeyUser), npncore.Title(npncore.KeyContent), npncore.Title(npncore.KeyCreated)}
			var data [][]string
			for _, s := range fs {
				data = append(data, []string{members.GetName(s.UserID), s.Content, npncore.ToDateString(&s.Created)})
			}
			table(cols, data, []uint{3, 6, 3}, m)
		}
	}
}
