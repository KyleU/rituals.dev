// Code generated by qtc from "Table.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vestimate/Table.html:2
package vestimate

//line views/vestimate/Table.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/estimate"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/views/components"
)

//line views/vestimate/Table.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vestimate/Table.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vestimate/Table.html:12
func StreamTable(qw422016 *qt422016.Writer, models estimate.Estimates, teams team.Teams, sprints sprint.Sprints, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vestimate/Table.html:12
	qw422016.N().S(`
`)
//line views/vestimate/Table.html:13
	prms := params.Get("estimate", nil, ps.Logger).Sanitize("estimate")

//line views/vestimate/Table.html:13
	qw422016.N().S(`  <table class="mt">
    <thead>
      <tr>
        `)
//line views/vestimate/Table.html:17
	components.StreamTableHeaderSimple(qw422016, "estimate", "id", "ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vestimate/Table.html:17
	qw422016.N().S(`
        `)
//line views/vestimate/Table.html:18
	components.StreamTableHeaderSimple(qw422016, "estimate", "slug", "Slug", "String text", prms, ps.URI, ps)
//line views/vestimate/Table.html:18
	qw422016.N().S(`
        `)
//line views/vestimate/Table.html:19
	components.StreamTableHeaderSimple(qw422016, "estimate", "title", "Title", "String text", prms, ps.URI, ps)
//line views/vestimate/Table.html:19
	qw422016.N().S(`
        `)
//line views/vestimate/Table.html:20
	components.StreamTableHeaderSimple(qw422016, "estimate", "icon", "Icon", "String text", prms, ps.URI, ps)
//line views/vestimate/Table.html:20
	qw422016.N().S(`
        `)
//line views/vestimate/Table.html:21
	components.StreamTableHeaderSimple(qw422016, "estimate", "status", "Status", "Available options: [new, active, complete]", prms, ps.URI, ps)
//line views/vestimate/Table.html:21
	qw422016.N().S(`
        `)
//line views/vestimate/Table.html:22
	components.StreamTableHeaderSimple(qw422016, "estimate", "team_id", "Team ID", "UUID in format (00000000-0000-0000-0000-000000000000) (optional)", prms, ps.URI, ps)
//line views/vestimate/Table.html:22
	qw422016.N().S(`
        `)
//line views/vestimate/Table.html:23
	components.StreamTableHeaderSimple(qw422016, "estimate", "sprint_id", "Sprint ID", "UUID in format (00000000-0000-0000-0000-000000000000) (optional)", prms, ps.URI, ps)
//line views/vestimate/Table.html:23
	qw422016.N().S(`
        `)
//line views/vestimate/Table.html:24
	components.StreamTableHeaderSimple(qw422016, "estimate", "choices", "Choices", "Comma-separated list of values", prms, ps.URI, ps)
//line views/vestimate/Table.html:24
	qw422016.N().S(`
        `)
//line views/vestimate/Table.html:25
	components.StreamTableHeaderSimple(qw422016, "estimate", "created", "Created", "Date and time, in almost any format", prms, ps.URI, ps)
//line views/vestimate/Table.html:25
	qw422016.N().S(`
        `)
//line views/vestimate/Table.html:26
	components.StreamTableHeaderSimple(qw422016, "estimate", "updated", "Updated", "Date and time, in almost any format (optional)", prms, ps.URI, ps)
//line views/vestimate/Table.html:26
	qw422016.N().S(`
      </tr>
    </thead>
    <tbody>
`)
//line views/vestimate/Table.html:30
	for _, model := range models {
//line views/vestimate/Table.html:30
		qw422016.N().S(`      <tr>
        <td><a href="/admin/db/estimate/`)
//line views/vestimate/Table.html:32
		components.StreamDisplayUUID(qw422016, &model.ID)
//line views/vestimate/Table.html:32
		qw422016.N().S(`">`)
//line views/vestimate/Table.html:32
		components.StreamDisplayUUID(qw422016, &model.ID)
//line views/vestimate/Table.html:32
		qw422016.N().S(`</a></td>
        <td>`)
//line views/vestimate/Table.html:33
		qw422016.E().S(model.Slug)
//line views/vestimate/Table.html:33
		qw422016.N().S(`</td>
        <td><strong>`)
//line views/vestimate/Table.html:34
		qw422016.E().S(model.Title)
//line views/vestimate/Table.html:34
		qw422016.N().S(`</strong></td>
        <td>`)
//line views/vestimate/Table.html:35
		qw422016.E().S(model.Icon)
//line views/vestimate/Table.html:35
		qw422016.N().S(`</td>
        <td>`)
//line views/vestimate/Table.html:36
		qw422016.E().V(model.Status)
//line views/vestimate/Table.html:36
		qw422016.N().S(`</td>
        <td class="nowrap">
          `)
//line views/vestimate/Table.html:38
		components.StreamDisplayUUID(qw422016, model.TeamID)
//line views/vestimate/Table.html:38
		if model.TeamID != nil {
//line views/vestimate/Table.html:38
			if x := teams.Get(*model.TeamID); x != nil {
//line views/vestimate/Table.html:38
				qw422016.N().S(` (`)
//line views/vestimate/Table.html:38
				qw422016.E().S(x.TitleString())
//line views/vestimate/Table.html:38
				qw422016.N().S(`)`)
//line views/vestimate/Table.html:38
			}
//line views/vestimate/Table.html:38
		}
//line views/vestimate/Table.html:38
		qw422016.N().S(`
          `)
//line views/vestimate/Table.html:39
		if model.TeamID != nil {
//line views/vestimate/Table.html:39
			qw422016.N().S(`<a title="Team" href="`)
//line views/vestimate/Table.html:39
			qw422016.E().S(`/team` + `/` + model.TeamID.String())
//line views/vestimate/Table.html:39
			qw422016.N().S(`">`)
//line views/vestimate/Table.html:39
			components.StreamSVGRef(qw422016, "team", 18, 18, "", ps)
//line views/vestimate/Table.html:39
			qw422016.N().S(`</a>`)
//line views/vestimate/Table.html:39
		}
//line views/vestimate/Table.html:39
		qw422016.N().S(`
        </td>
        <td class="nowrap">
          `)
//line views/vestimate/Table.html:42
		components.StreamDisplayUUID(qw422016, model.SprintID)
//line views/vestimate/Table.html:42
		if model.SprintID != nil {
//line views/vestimate/Table.html:42
			if x := sprints.Get(*model.SprintID); x != nil {
//line views/vestimate/Table.html:42
				qw422016.N().S(` (`)
//line views/vestimate/Table.html:42
				qw422016.E().S(x.TitleString())
//line views/vestimate/Table.html:42
				qw422016.N().S(`)`)
//line views/vestimate/Table.html:42
			}
//line views/vestimate/Table.html:42
		}
//line views/vestimate/Table.html:42
		qw422016.N().S(`
          `)
//line views/vestimate/Table.html:43
		if model.SprintID != nil {
//line views/vestimate/Table.html:43
			qw422016.N().S(`<a title="Sprint" href="`)
//line views/vestimate/Table.html:43
			qw422016.E().S(`/sprint` + `/` + model.SprintID.String())
//line views/vestimate/Table.html:43
			qw422016.N().S(`">`)
//line views/vestimate/Table.html:43
			components.StreamSVGRef(qw422016, "sprint", 18, 18, "", ps)
//line views/vestimate/Table.html:43
			qw422016.N().S(`</a>`)
//line views/vestimate/Table.html:43
		}
//line views/vestimate/Table.html:43
		qw422016.N().S(`
        </td>
        <td>`)
//line views/vestimate/Table.html:45
		components.StreamDisplayStringArray(qw422016, model.Choices)
//line views/vestimate/Table.html:45
		qw422016.N().S(`</td>
        <td>`)
//line views/vestimate/Table.html:46
		components.StreamDisplayTimestamp(qw422016, &model.Created)
//line views/vestimate/Table.html:46
		qw422016.N().S(`</td>
        <td>`)
//line views/vestimate/Table.html:47
		components.StreamDisplayTimestamp(qw422016, model.Updated)
//line views/vestimate/Table.html:47
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vestimate/Table.html:49
	}
//line views/vestimate/Table.html:50
	if prms.HasNextPage(len(models)+prms.Offset) || prms.HasPreviousPage() {
//line views/vestimate/Table.html:50
		qw422016.N().S(`      <tr>
        <td colspan="10">`)
//line views/vestimate/Table.html:52
		components.StreamPagination(qw422016, len(models)+prms.Offset, prms, ps.URI)
//line views/vestimate/Table.html:52
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vestimate/Table.html:54
	}
//line views/vestimate/Table.html:54
	qw422016.N().S(`    </tbody>
  </table>
`)
//line views/vestimate/Table.html:57
}

//line views/vestimate/Table.html:57
func WriteTable(qq422016 qtio422016.Writer, models estimate.Estimates, teams team.Teams, sprints sprint.Sprints, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vestimate/Table.html:57
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vestimate/Table.html:57
	StreamTable(qw422016, models, teams, sprints, params, as, ps)
//line views/vestimate/Table.html:57
	qt422016.ReleaseWriter(qw422016)
//line views/vestimate/Table.html:57
}

//line views/vestimate/Table.html:57
func Table(models estimate.Estimates, teams team.Teams, sprints sprint.Sprints, params filter.ParamSet, as *app.State, ps *cutil.PageState) string {
//line views/vestimate/Table.html:57
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vestimate/Table.html:57
	WriteTable(qb422016, models, teams, sprints, params, as, ps)
//line views/vestimate/Table.html:57
	qs422016 := string(qb422016.B)
//line views/vestimate/Table.html:57
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vestimate/Table.html:57
	return qs422016
//line views/vestimate/Table.html:57
}
