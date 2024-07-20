// Code generated by qtc from "Table.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vestimate/Table.html:1
package vestimate

//line views/vestimate/Table.html:1
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/estimate"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/components/view"
)

//line views/vestimate/Table.html:13
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vestimate/Table.html:13
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vestimate/Table.html:13
func StreamTable(qw422016 *qt422016.Writer, models estimate.Estimates, teamsByTeamID team.Teams, sprintsBySprintID sprint.Sprints, params filter.ParamSet, as *app.State, ps *cutil.PageState, paths ...string) {
//line views/vestimate/Table.html:13
	qw422016.N().S(`
`)
//line views/vestimate/Table.html:14
	prms := params.Sanitized("estimate", ps.Logger)

//line views/vestimate/Table.html:14
	qw422016.N().S(`  <div class="overflow clear">
    <table>
      <thead>
        <tr>
          `)
//line views/vestimate/Table.html:19
	components.StreamTableHeaderSimple(qw422016, "estimate", "id", "ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vestimate/Table.html:19
	qw422016.N().S(`
          `)
//line views/vestimate/Table.html:20
	components.StreamTableHeaderSimple(qw422016, "estimate", "slug", "Slug", "String text", prms, ps.URI, ps)
//line views/vestimate/Table.html:20
	qw422016.N().S(`
          `)
//line views/vestimate/Table.html:21
	components.StreamTableHeaderSimple(qw422016, "estimate", "title", "Title", "String text", prms, ps.URI, ps)
//line views/vestimate/Table.html:21
	qw422016.N().S(`
          `)
//line views/vestimate/Table.html:22
	components.StreamTableHeaderSimple(qw422016, "estimate", "icon", "Icon", "String text", prms, ps.URI, ps)
//line views/vestimate/Table.html:22
	qw422016.N().S(`
          `)
//line views/vestimate/Table.html:23
	components.StreamTableHeaderSimple(qw422016, "estimate", "status", "Status", enum.AllSessionStatuses.Help(), prms, ps.URI, ps)
//line views/vestimate/Table.html:23
	qw422016.N().S(`
          `)
//line views/vestimate/Table.html:24
	components.StreamTableHeaderSimple(qw422016, "estimate", "team_id", "Team ID", "UUID in format (00000000-0000-0000-0000-000000000000) (optional)", prms, ps.URI, ps)
//line views/vestimate/Table.html:24
	qw422016.N().S(`
          `)
//line views/vestimate/Table.html:25
	components.StreamTableHeaderSimple(qw422016, "estimate", "sprint_id", "Sprint ID", "UUID in format (00000000-0000-0000-0000-000000000000) (optional)", prms, ps.URI, ps)
//line views/vestimate/Table.html:25
	qw422016.N().S(`
          `)
//line views/vestimate/Table.html:26
	components.StreamTableHeaderSimple(qw422016, "estimate", "choices", "Choices", "Comma-separated list of values", prms, ps.URI, ps)
//line views/vestimate/Table.html:26
	qw422016.N().S(`
          `)
//line views/vestimate/Table.html:27
	components.StreamTableHeaderSimple(qw422016, "estimate", "created", "Created", "Date and time, in almost any format", prms, ps.URI, ps)
//line views/vestimate/Table.html:27
	qw422016.N().S(`
          `)
//line views/vestimate/Table.html:28
	components.StreamTableHeaderSimple(qw422016, "estimate", "updated", "Updated", "Date and time, in almost any format (optional)", prms, ps.URI, ps)
//line views/vestimate/Table.html:28
	qw422016.N().S(`
        </tr>
      </thead>
      <tbody>
`)
//line views/vestimate/Table.html:32
	for _, model := range models {
//line views/vestimate/Table.html:32
		qw422016.N().S(`        <tr>
          <td><a href="`)
//line views/vestimate/Table.html:34
		qw422016.E().S(model.WebPath(paths...))
//line views/vestimate/Table.html:34
		qw422016.N().S(`">`)
//line views/vestimate/Table.html:34
		view.StreamUUID(qw422016, &model.ID)
//line views/vestimate/Table.html:34
		qw422016.N().S(`</a></td>
          <td>`)
//line views/vestimate/Table.html:35
		view.StreamString(qw422016, model.Slug)
//line views/vestimate/Table.html:35
		qw422016.N().S(`</td>
          <td><strong>`)
//line views/vestimate/Table.html:36
		view.StreamString(qw422016, model.Title)
//line views/vestimate/Table.html:36
		qw422016.N().S(`</strong></td>
          <td>`)
//line views/vestimate/Table.html:37
		view.StreamString(qw422016, model.Icon)
//line views/vestimate/Table.html:37
		qw422016.N().S(`</td>
          <td>`)
//line views/vestimate/Table.html:38
		qw422016.E().S(model.Status.String())
//line views/vestimate/Table.html:38
		qw422016.N().S(`</td>
          <td class="nowrap">
            `)
//line views/vestimate/Table.html:40
		view.StreamUUID(qw422016, model.TeamID)
//line views/vestimate/Table.html:40
		if model.TeamID != nil {
//line views/vestimate/Table.html:40
			if x := teamsByTeamID.Get(*model.TeamID); x != nil {
//line views/vestimate/Table.html:40
				qw422016.N().S(` (`)
//line views/vestimate/Table.html:40
				qw422016.E().S(x.TitleString())
//line views/vestimate/Table.html:40
				qw422016.N().S(`)`)
//line views/vestimate/Table.html:40
			}
//line views/vestimate/Table.html:40
		}
//line views/vestimate/Table.html:40
		qw422016.N().S(`
            `)
//line views/vestimate/Table.html:41
		if model.TeamID != nil {
//line views/vestimate/Table.html:41
			qw422016.N().S(`<a title="Team" href="`)
//line views/vestimate/Table.html:41
			qw422016.E().S(model.WebPath(paths...))
//line views/vestimate/Table.html:41
			qw422016.N().S(`">`)
//line views/vestimate/Table.html:41
			components.StreamSVGLink(qw422016, `team`, ps)
//line views/vestimate/Table.html:41
			qw422016.N().S(`</a>`)
//line views/vestimate/Table.html:41
		}
//line views/vestimate/Table.html:41
		qw422016.N().S(`
          </td>
          <td class="nowrap">
            `)
//line views/vestimate/Table.html:44
		view.StreamUUID(qw422016, model.SprintID)
//line views/vestimate/Table.html:44
		if model.SprintID != nil {
//line views/vestimate/Table.html:44
			if x := sprintsBySprintID.Get(*model.SprintID); x != nil {
//line views/vestimate/Table.html:44
				qw422016.N().S(` (`)
//line views/vestimate/Table.html:44
				qw422016.E().S(x.TitleString())
//line views/vestimate/Table.html:44
				qw422016.N().S(`)`)
//line views/vestimate/Table.html:44
			}
//line views/vestimate/Table.html:44
		}
//line views/vestimate/Table.html:44
		qw422016.N().S(`
            `)
//line views/vestimate/Table.html:45
		if model.SprintID != nil {
//line views/vestimate/Table.html:45
			qw422016.N().S(`<a title="Sprint" href="`)
//line views/vestimate/Table.html:45
			qw422016.E().S(model.WebPath(paths...))
//line views/vestimate/Table.html:45
			qw422016.N().S(`">`)
//line views/vestimate/Table.html:45
			components.StreamSVGLink(qw422016, `sprint`, ps)
//line views/vestimate/Table.html:45
			qw422016.N().S(`</a>`)
//line views/vestimate/Table.html:45
		}
//line views/vestimate/Table.html:45
		qw422016.N().S(`
          </td>
          <td>`)
//line views/vestimate/Table.html:47
		view.StreamStringArray(qw422016, model.Choices)
//line views/vestimate/Table.html:47
		qw422016.N().S(`</td>
          <td>`)
//line views/vestimate/Table.html:48
		view.StreamTimestamp(qw422016, &model.Created)
//line views/vestimate/Table.html:48
		qw422016.N().S(`</td>
          <td>`)
//line views/vestimate/Table.html:49
		view.StreamTimestamp(qw422016, model.Updated)
//line views/vestimate/Table.html:49
		qw422016.N().S(`</td>
        </tr>
`)
//line views/vestimate/Table.html:51
	}
//line views/vestimate/Table.html:51
	qw422016.N().S(`      </tbody>
    </table>
  </div>
`)
//line views/vestimate/Table.html:55
	if prms.HasNextPage(len(models)+prms.Offset) || prms.HasPreviousPage() {
//line views/vestimate/Table.html:55
		qw422016.N().S(`  <hr />
  `)
//line views/vestimate/Table.html:57
		components.StreamPagination(qw422016, len(models)+prms.Offset, prms, ps.URI)
//line views/vestimate/Table.html:57
		qw422016.N().S(`
  <div class="clear"></div>
`)
//line views/vestimate/Table.html:59
	}
//line views/vestimate/Table.html:60
}

//line views/vestimate/Table.html:60
func WriteTable(qq422016 qtio422016.Writer, models estimate.Estimates, teamsByTeamID team.Teams, sprintsBySprintID sprint.Sprints, params filter.ParamSet, as *app.State, ps *cutil.PageState, paths ...string) {
//line views/vestimate/Table.html:60
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vestimate/Table.html:60
	StreamTable(qw422016, models, teamsByTeamID, sprintsBySprintID, params, as, ps, paths...)
//line views/vestimate/Table.html:60
	qt422016.ReleaseWriter(qw422016)
//line views/vestimate/Table.html:60
}

//line views/vestimate/Table.html:60
func Table(models estimate.Estimates, teamsByTeamID team.Teams, sprintsBySprintID sprint.Sprints, params filter.ParamSet, as *app.State, ps *cutil.PageState, paths ...string) string {
//line views/vestimate/Table.html:60
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vestimate/Table.html:60
	WriteTable(qb422016, models, teamsByTeamID, sprintsBySprintID, params, as, ps, paths...)
//line views/vestimate/Table.html:60
	qs422016 := string(qb422016.B)
//line views/vestimate/Table.html:60
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vestimate/Table.html:60
	return qs422016
//line views/vestimate/Table.html:60
}
