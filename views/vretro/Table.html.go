// Code generated by qtc from "Table.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vretro/Table.html:2
package vretro

//line views/vretro/Table.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/retro"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/components/view"
)

//line views/vretro/Table.html:14
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vretro/Table.html:14
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vretro/Table.html:14
func StreamTable(qw422016 *qt422016.Writer, models retro.Retros, teamsByTeamID team.Teams, sprintsBySprintID sprint.Sprints, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vretro/Table.html:14
	qw422016.N().S(`
`)
//line views/vretro/Table.html:15
	prms := params.Get("retro", nil, ps.Logger).Sanitize("retro")

//line views/vretro/Table.html:15
	qw422016.N().S(`  <table>
    <thead>
      <tr>
        `)
//line views/vretro/Table.html:19
	components.StreamTableHeaderSimple(qw422016, "retro", "id", "ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vretro/Table.html:19
	qw422016.N().S(`
        `)
//line views/vretro/Table.html:20
	components.StreamTableHeaderSimple(qw422016, "retro", "slug", "Slug", "String text", prms, ps.URI, ps)
//line views/vretro/Table.html:20
	qw422016.N().S(`
        `)
//line views/vretro/Table.html:21
	components.StreamTableHeaderSimple(qw422016, "retro", "title", "Title", "String text", prms, ps.URI, ps)
//line views/vretro/Table.html:21
	qw422016.N().S(`
        `)
//line views/vretro/Table.html:22
	components.StreamTableHeaderSimple(qw422016, "retro", "icon", "Icon", "String text", prms, ps.URI, ps)
//line views/vretro/Table.html:22
	qw422016.N().S(`
        `)
//line views/vretro/Table.html:23
	components.StreamTableHeaderSimple(qw422016, "retro", "status", "Status", enum.AllSessionStatuses.Help(), prms, ps.URI, ps)
//line views/vretro/Table.html:23
	qw422016.N().S(`
        `)
//line views/vretro/Table.html:24
	components.StreamTableHeaderSimple(qw422016, "retro", "team_id", "Team ID", "UUID in format (00000000-0000-0000-0000-000000000000) (optional)", prms, ps.URI, ps)
//line views/vretro/Table.html:24
	qw422016.N().S(`
        `)
//line views/vretro/Table.html:25
	components.StreamTableHeaderSimple(qw422016, "retro", "sprint_id", "Sprint ID", "UUID in format (00000000-0000-0000-0000-000000000000) (optional)", prms, ps.URI, ps)
//line views/vretro/Table.html:25
	qw422016.N().S(`
        `)
//line views/vretro/Table.html:26
	components.StreamTableHeaderSimple(qw422016, "retro", "categories", "Categories", "Comma-separated list of values", prms, ps.URI, ps)
//line views/vretro/Table.html:26
	qw422016.N().S(`
        `)
//line views/vretro/Table.html:27
	components.StreamTableHeaderSimple(qw422016, "retro", "created", "Created", "Date and time, in almost any format", prms, ps.URI, ps)
//line views/vretro/Table.html:27
	qw422016.N().S(`
        `)
//line views/vretro/Table.html:28
	components.StreamTableHeaderSimple(qw422016, "retro", "updated", "Updated", "Date and time, in almost any format (optional)", prms, ps.URI, ps)
//line views/vretro/Table.html:28
	qw422016.N().S(`
      </tr>
    </thead>
    <tbody>
`)
//line views/vretro/Table.html:32
	for _, model := range models {
//line views/vretro/Table.html:32
		qw422016.N().S(`      <tr>
        <td><a href="/admin/db/retro/`)
//line views/vretro/Table.html:34
		view.StreamUUID(qw422016, &model.ID)
//line views/vretro/Table.html:34
		qw422016.N().S(`">`)
//line views/vretro/Table.html:34
		view.StreamUUID(qw422016, &model.ID)
//line views/vretro/Table.html:34
		qw422016.N().S(`</a></td>
        <td>`)
//line views/vretro/Table.html:35
		view.StreamString(qw422016, model.Slug)
//line views/vretro/Table.html:35
		qw422016.N().S(`</td>
        <td><strong>`)
//line views/vretro/Table.html:36
		view.StreamString(qw422016, model.Title)
//line views/vretro/Table.html:36
		qw422016.N().S(`</strong></td>
        <td>`)
//line views/vretro/Table.html:37
		view.StreamString(qw422016, model.Icon)
//line views/vretro/Table.html:37
		qw422016.N().S(`</td>
        <td>`)
//line views/vretro/Table.html:38
		qw422016.E().S(model.Status.String())
//line views/vretro/Table.html:38
		qw422016.N().S(`</td>
        <td class="nowrap">
          `)
//line views/vretro/Table.html:40
		view.StreamUUID(qw422016, model.TeamID)
//line views/vretro/Table.html:40
		if model.TeamID != nil {
//line views/vretro/Table.html:40
			if x := teamsByTeamID.Get(*model.TeamID); x != nil {
//line views/vretro/Table.html:40
				qw422016.N().S(` (`)
//line views/vretro/Table.html:40
				qw422016.E().S(x.TitleString())
//line views/vretro/Table.html:40
				qw422016.N().S(`)`)
//line views/vretro/Table.html:40
			}
//line views/vretro/Table.html:40
		}
//line views/vretro/Table.html:40
		qw422016.N().S(`
          `)
//line views/vretro/Table.html:41
		if model.TeamID != nil {
//line views/vretro/Table.html:41
			qw422016.N().S(`<a title="Team" href="`)
//line views/vretro/Table.html:41
			qw422016.E().S(`/admin/db/team` + `/` + model.TeamID.String())
//line views/vretro/Table.html:41
			qw422016.N().S(`">`)
//line views/vretro/Table.html:41
			components.StreamSVGRef(qw422016, "team", 18, 18, "", ps)
//line views/vretro/Table.html:41
			qw422016.N().S(`</a>`)
//line views/vretro/Table.html:41
		}
//line views/vretro/Table.html:41
		qw422016.N().S(`
        </td>
        <td class="nowrap">
          `)
//line views/vretro/Table.html:44
		view.StreamUUID(qw422016, model.SprintID)
//line views/vretro/Table.html:44
		if model.SprintID != nil {
//line views/vretro/Table.html:44
			if x := sprintsBySprintID.Get(*model.SprintID); x != nil {
//line views/vretro/Table.html:44
				qw422016.N().S(` (`)
//line views/vretro/Table.html:44
				qw422016.E().S(x.TitleString())
//line views/vretro/Table.html:44
				qw422016.N().S(`)`)
//line views/vretro/Table.html:44
			}
//line views/vretro/Table.html:44
		}
//line views/vretro/Table.html:44
		qw422016.N().S(`
          `)
//line views/vretro/Table.html:45
		if model.SprintID != nil {
//line views/vretro/Table.html:45
			qw422016.N().S(`<a title="Sprint" href="`)
//line views/vretro/Table.html:45
			qw422016.E().S(`/admin/db/sprint` + `/` + model.SprintID.String())
//line views/vretro/Table.html:45
			qw422016.N().S(`">`)
//line views/vretro/Table.html:45
			components.StreamSVGRef(qw422016, "sprint", 18, 18, "", ps)
//line views/vretro/Table.html:45
			qw422016.N().S(`</a>`)
//line views/vretro/Table.html:45
		}
//line views/vretro/Table.html:45
		qw422016.N().S(`
        </td>
        <td>`)
//line views/vretro/Table.html:47
		view.StreamStringArray(qw422016, model.Categories)
//line views/vretro/Table.html:47
		qw422016.N().S(`</td>
        <td>`)
//line views/vretro/Table.html:48
		view.StreamTimestamp(qw422016, &model.Created)
//line views/vretro/Table.html:48
		qw422016.N().S(`</td>
        <td>`)
//line views/vretro/Table.html:49
		view.StreamTimestamp(qw422016, model.Updated)
//line views/vretro/Table.html:49
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vretro/Table.html:51
	}
//line views/vretro/Table.html:52
	if prms.HasNextPage(len(models)+prms.Offset) || prms.HasPreviousPage() {
//line views/vretro/Table.html:52
		qw422016.N().S(`      <tr>
        <td colspan="10">`)
//line views/vretro/Table.html:54
		components.StreamPagination(qw422016, len(models)+prms.Offset, prms, ps.URI)
//line views/vretro/Table.html:54
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vretro/Table.html:56
	}
//line views/vretro/Table.html:56
	qw422016.N().S(`    </tbody>
  </table>
`)
//line views/vretro/Table.html:59
}

//line views/vretro/Table.html:59
func WriteTable(qq422016 qtio422016.Writer, models retro.Retros, teamsByTeamID team.Teams, sprintsBySprintID sprint.Sprints, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vretro/Table.html:59
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vretro/Table.html:59
	StreamTable(qw422016, models, teamsByTeamID, sprintsBySprintID, params, as, ps)
//line views/vretro/Table.html:59
	qt422016.ReleaseWriter(qw422016)
//line views/vretro/Table.html:59
}

//line views/vretro/Table.html:59
func Table(models retro.Retros, teamsByTeamID team.Teams, sprintsBySprintID sprint.Sprints, params filter.ParamSet, as *app.State, ps *cutil.PageState) string {
//line views/vretro/Table.html:59
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vretro/Table.html:59
	WriteTable(qb422016, models, teamsByTeamID, sprintsBySprintID, params, as, ps)
//line views/vretro/Table.html:59
	qs422016 := string(qb422016.B)
//line views/vretro/Table.html:59
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vretro/Table.html:59
	return qs422016
//line views/vretro/Table.html:59
}
