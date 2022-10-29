// Code generated by qtc from "Table.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vstandup/Table.html:2
package vstandup

//line views/vstandup/Table.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/standup"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/views/components"
)

//line views/vstandup/Table.html:13
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vstandup/Table.html:13
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vstandup/Table.html:13
func StreamTable(qw422016 *qt422016.Writer, models standup.Standups, users user.Users, teams team.Teams, sprints sprint.Sprints, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vstandup/Table.html:13
	qw422016.N().S(`
`)
//line views/vstandup/Table.html:14
	prms := params.Get("standup", nil, ps.Logger).Sanitize("standup")

//line views/vstandup/Table.html:14
	qw422016.N().S(`  <table class="mt">
    <thead>
      <tr>
        `)
//line views/vstandup/Table.html:18
	components.StreamTableHeaderSimple(qw422016, "standup", "id", "ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vstandup/Table.html:18
	qw422016.N().S(`
        `)
//line views/vstandup/Table.html:19
	components.StreamTableHeaderSimple(qw422016, "standup", "slug", "Slug", "String text", prms, ps.URI, ps)
//line views/vstandup/Table.html:19
	qw422016.N().S(`
        `)
//line views/vstandup/Table.html:20
	components.StreamTableHeaderSimple(qw422016, "standup", "title", "Title", "String text", prms, ps.URI, ps)
//line views/vstandup/Table.html:20
	qw422016.N().S(`
        `)
//line views/vstandup/Table.html:21
	components.StreamTableHeaderSimple(qw422016, "standup", "status", "Status", "Available options: [new, active, complete, deleted]", prms, ps.URI, ps)
//line views/vstandup/Table.html:21
	qw422016.N().S(`
        `)
//line views/vstandup/Table.html:22
	components.StreamTableHeaderSimple(qw422016, "standup", "team_id", "Team ID", "UUID in format (00000000-0000-0000-0000-000000000000) (optional)", prms, ps.URI, ps)
//line views/vstandup/Table.html:22
	qw422016.N().S(`
        `)
//line views/vstandup/Table.html:23
	components.StreamTableHeaderSimple(qw422016, "standup", "sprint_id", "Sprint ID", "UUID in format (00000000-0000-0000-0000-000000000000) (optional)", prms, ps.URI, ps)
//line views/vstandup/Table.html:23
	qw422016.N().S(`
        `)
//line views/vstandup/Table.html:24
	components.StreamTableHeaderSimple(qw422016, "standup", "owner", "Owner", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vstandup/Table.html:24
	qw422016.N().S(`
        `)
//line views/vstandup/Table.html:25
	components.StreamTableHeaderSimple(qw422016, "standup", "created", "Created", "Date and time, in almost any format", prms, ps.URI, ps)
//line views/vstandup/Table.html:25
	qw422016.N().S(`
        `)
//line views/vstandup/Table.html:26
	components.StreamTableHeaderSimple(qw422016, "standup", "updated", "Updated", "Date and time, in almost any format (optional)", prms, ps.URI, ps)
//line views/vstandup/Table.html:26
	qw422016.N().S(`
      </tr>
    </thead>
    <tbody>
`)
//line views/vstandup/Table.html:30
	for _, model := range models {
//line views/vstandup/Table.html:30
		qw422016.N().S(`      <tr>
        <td><a href="/admin/db/standup/`)
//line views/vstandup/Table.html:32
		components.StreamDisplayUUID(qw422016, &model.ID)
//line views/vstandup/Table.html:32
		qw422016.N().S(`">`)
//line views/vstandup/Table.html:32
		components.StreamDisplayUUID(qw422016, &model.ID)
//line views/vstandup/Table.html:32
		qw422016.N().S(`</a></td>
        <td>`)
//line views/vstandup/Table.html:33
		qw422016.E().S(model.Slug)
//line views/vstandup/Table.html:33
		qw422016.N().S(`</td>
        <td><strong>`)
//line views/vstandup/Table.html:34
		qw422016.E().S(model.Title)
//line views/vstandup/Table.html:34
		qw422016.N().S(`</strong></td>
        <td>`)
//line views/vstandup/Table.html:35
		qw422016.E().V(model.Status)
//line views/vstandup/Table.html:35
		qw422016.N().S(`</td>
        <td>
          <div class="icon">`)
//line views/vstandup/Table.html:37
		components.StreamDisplayUUID(qw422016, model.TeamID)
//line views/vstandup/Table.html:37
		if model.TeamID != nil {
//line views/vstandup/Table.html:37
			if x := teams.Get(*model.TeamID); x != nil {
//line views/vstandup/Table.html:37
				qw422016.N().S(` (`)
//line views/vstandup/Table.html:37
				qw422016.E().S(x.TitleString())
//line views/vstandup/Table.html:37
				qw422016.N().S(`)`)
//line views/vstandup/Table.html:37
			}
//line views/vstandup/Table.html:37
		}
//line views/vstandup/Table.html:37
		qw422016.N().S(`</div>
          <a title="Team" href="`)
//line views/vstandup/Table.html:38
		qw422016.E().S(`/team` + `/` + model.TeamID.String())
//line views/vstandup/Table.html:38
		qw422016.N().S(`">`)
//line views/vstandup/Table.html:38
		components.StreamSVGRefIcon(qw422016, "team", ps)
//line views/vstandup/Table.html:38
		qw422016.N().S(`</a>
        </td>
        <td>
          <div class="icon">`)
//line views/vstandup/Table.html:41
		components.StreamDisplayUUID(qw422016, model.SprintID)
//line views/vstandup/Table.html:41
		if model.SprintID != nil {
//line views/vstandup/Table.html:41
			if x := sprints.Get(*model.SprintID); x != nil {
//line views/vstandup/Table.html:41
				qw422016.N().S(` (`)
//line views/vstandup/Table.html:41
				qw422016.E().S(x.TitleString())
//line views/vstandup/Table.html:41
				qw422016.N().S(`)`)
//line views/vstandup/Table.html:41
			}
//line views/vstandup/Table.html:41
		}
//line views/vstandup/Table.html:41
		qw422016.N().S(`</div>
          <a title="Sprint" href="`)
//line views/vstandup/Table.html:42
		qw422016.E().S(`/sprint` + `/` + model.SprintID.String())
//line views/vstandup/Table.html:42
		qw422016.N().S(`">`)
//line views/vstandup/Table.html:42
		components.StreamSVGRefIcon(qw422016, "sprint", ps)
//line views/vstandup/Table.html:42
		qw422016.N().S(`</a>
        </td>
        <td>
          <div class="icon">`)
//line views/vstandup/Table.html:45
		components.StreamDisplayUUID(qw422016, &model.Owner)
//line views/vstandup/Table.html:45
		if x := users.Get(model.Owner); x != nil {
//line views/vstandup/Table.html:45
			qw422016.N().S(` (`)
//line views/vstandup/Table.html:45
			qw422016.E().S(x.TitleString())
//line views/vstandup/Table.html:45
			qw422016.N().S(`)`)
//line views/vstandup/Table.html:45
		}
//line views/vstandup/Table.html:45
		qw422016.N().S(`</div>
          <a title="User" href="`)
//line views/vstandup/Table.html:46
		qw422016.E().S(`/user` + `/` + model.Owner.String())
//line views/vstandup/Table.html:46
		qw422016.N().S(`">`)
//line views/vstandup/Table.html:46
		components.StreamSVGRefIcon(qw422016, "profile", ps)
//line views/vstandup/Table.html:46
		qw422016.N().S(`</a>
        </td>
        <td>`)
//line views/vstandup/Table.html:48
		components.StreamDisplayTimestamp(qw422016, &model.Created)
//line views/vstandup/Table.html:48
		qw422016.N().S(`</td>
        <td>`)
//line views/vstandup/Table.html:49
		components.StreamDisplayTimestamp(qw422016, model.Updated)
//line views/vstandup/Table.html:49
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vstandup/Table.html:51
	}
//line views/vstandup/Table.html:52
	if prms.HasNextPage(len(models)+prms.Offset) || prms.HasPreviousPage() {
//line views/vstandup/Table.html:52
		qw422016.N().S(`      <tr>
        <td colspan="9">`)
//line views/vstandup/Table.html:54
		components.StreamPagination(qw422016, len(models)+prms.Offset, prms, ps.URI)
//line views/vstandup/Table.html:54
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vstandup/Table.html:56
	}
//line views/vstandup/Table.html:56
	qw422016.N().S(`    </tbody>
  </table>
`)
//line views/vstandup/Table.html:59
}

//line views/vstandup/Table.html:59
func WriteTable(qq422016 qtio422016.Writer, models standup.Standups, users user.Users, teams team.Teams, sprints sprint.Sprints, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vstandup/Table.html:59
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vstandup/Table.html:59
	StreamTable(qw422016, models, users, teams, sprints, params, as, ps)
//line views/vstandup/Table.html:59
	qt422016.ReleaseWriter(qw422016)
//line views/vstandup/Table.html:59
}

//line views/vstandup/Table.html:59
func Table(models standup.Standups, users user.Users, teams team.Teams, sprints sprint.Sprints, params filter.ParamSet, as *app.State, ps *cutil.PageState) string {
//line views/vstandup/Table.html:59
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vstandup/Table.html:59
	WriteTable(qb422016, models, users, teams, sprints, params, as, ps)
//line views/vstandup/Table.html:59
	qs422016 := string(qb422016.B)
//line views/vstandup/Table.html:59
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vstandup/Table.html:59
	return qs422016
//line views/vstandup/Table.html:59
}
