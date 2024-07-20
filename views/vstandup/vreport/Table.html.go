// Code generated by qtc from "Table.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vstandup/vreport/Table.html:1
package vreport

//line views/vstandup/vreport/Table.html:1
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/standup"
	"github.com/kyleu/rituals/app/standup/report"
	"github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/components/view"
)

//line views/vstandup/vreport/Table.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vstandup/vreport/Table.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vstandup/vreport/Table.html:12
func StreamTable(qw422016 *qt422016.Writer, models report.Reports, standupsByStandupID standup.Standups, usersByUserID user.Users, params filter.ParamSet, as *app.State, ps *cutil.PageState, paths ...string) {
//line views/vstandup/vreport/Table.html:12
	qw422016.N().S(`
`)
//line views/vstandup/vreport/Table.html:13
	prms := params.Sanitized("report", ps.Logger)

//line views/vstandup/vreport/Table.html:13
	qw422016.N().S(`  <div class="overflow clear">
    <table>
      <thead>
        <tr>
          `)
//line views/vstandup/vreport/Table.html:18
	components.StreamTableHeaderSimple(qw422016, "report", "id", "ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vstandup/vreport/Table.html:18
	qw422016.N().S(`
          `)
//line views/vstandup/vreport/Table.html:19
	components.StreamTableHeaderSimple(qw422016, "report", "standup_id", "Standup ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vstandup/vreport/Table.html:19
	qw422016.N().S(`
          `)
//line views/vstandup/vreport/Table.html:20
	components.StreamTableHeaderSimple(qw422016, "report", "day", "Day", "Calendar date", prms, ps.URI, ps)
//line views/vstandup/vreport/Table.html:20
	qw422016.N().S(`
          `)
//line views/vstandup/vreport/Table.html:21
	components.StreamTableHeaderSimple(qw422016, "report", "user_id", "User ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vstandup/vreport/Table.html:21
	qw422016.N().S(`
          `)
//line views/vstandup/vreport/Table.html:22
	components.StreamTableHeaderSimple(qw422016, "report", "content", "Content", "String text", prms, ps.URI, ps)
//line views/vstandup/vreport/Table.html:22
	qw422016.N().S(`
          `)
//line views/vstandup/vreport/Table.html:23
	components.StreamTableHeaderSimple(qw422016, "report", "created", "Created", "Date and time, in almost any format", prms, ps.URI, ps)
//line views/vstandup/vreport/Table.html:23
	qw422016.N().S(`
          `)
//line views/vstandup/vreport/Table.html:24
	components.StreamTableHeaderSimple(qw422016, "report", "updated", "Updated", "Date and time, in almost any format (optional)", prms, ps.URI, ps)
//line views/vstandup/vreport/Table.html:24
	qw422016.N().S(`
        </tr>
      </thead>
      <tbody>
`)
//line views/vstandup/vreport/Table.html:28
	for _, model := range models {
//line views/vstandup/vreport/Table.html:28
		qw422016.N().S(`        <tr>
          <td><a href="`)
//line views/vstandup/vreport/Table.html:30
		qw422016.E().S(model.WebPath(paths...))
//line views/vstandup/vreport/Table.html:30
		qw422016.N().S(`">`)
//line views/vstandup/vreport/Table.html:30
		view.StreamUUID(qw422016, &model.ID)
//line views/vstandup/vreport/Table.html:30
		qw422016.N().S(`</a></td>
          <td class="nowrap">
            `)
//line views/vstandup/vreport/Table.html:32
		view.StreamUUID(qw422016, &model.StandupID)
//line views/vstandup/vreport/Table.html:32
		if x := standupsByStandupID.Get(model.StandupID); x != nil {
//line views/vstandup/vreport/Table.html:32
			qw422016.N().S(` (`)
//line views/vstandup/vreport/Table.html:32
			qw422016.E().S(x.TitleString())
//line views/vstandup/vreport/Table.html:32
			qw422016.N().S(`)`)
//line views/vstandup/vreport/Table.html:32
		}
//line views/vstandup/vreport/Table.html:32
		qw422016.N().S(`
            <a title="Standup" href="`)
//line views/vstandup/vreport/Table.html:33
		qw422016.E().S(model.WebPath(paths...))
//line views/vstandup/vreport/Table.html:33
		qw422016.N().S(`">`)
//line views/vstandup/vreport/Table.html:33
		components.StreamSVGLink(qw422016, `standup`, ps)
//line views/vstandup/vreport/Table.html:33
		qw422016.N().S(`</a>
          </td>
          <td>`)
//line views/vstandup/vreport/Table.html:35
		view.StreamTimestampDay(qw422016, &model.Day)
//line views/vstandup/vreport/Table.html:35
		qw422016.N().S(`</td>
          <td class="nowrap">
            `)
//line views/vstandup/vreport/Table.html:37
		view.StreamUUID(qw422016, &model.UserID)
//line views/vstandup/vreport/Table.html:37
		if x := usersByUserID.Get(model.UserID); x != nil {
//line views/vstandup/vreport/Table.html:37
			qw422016.N().S(` (`)
//line views/vstandup/vreport/Table.html:37
			qw422016.E().S(x.TitleString())
//line views/vstandup/vreport/Table.html:37
			qw422016.N().S(`)`)
//line views/vstandup/vreport/Table.html:37
		}
//line views/vstandup/vreport/Table.html:37
		qw422016.N().S(`
            <a title="User" href="`)
//line views/vstandup/vreport/Table.html:38
		qw422016.E().S(model.WebPath(paths...))
//line views/vstandup/vreport/Table.html:38
		qw422016.N().S(`">`)
//line views/vstandup/vreport/Table.html:38
		components.StreamSVGLink(qw422016, `profile`, ps)
//line views/vstandup/vreport/Table.html:38
		qw422016.N().S(`</a>
          </td>
          <td>`)
//line views/vstandup/vreport/Table.html:40
		view.StreamString(qw422016, model.Content)
//line views/vstandup/vreport/Table.html:40
		qw422016.N().S(`</td>
          <td>`)
//line views/vstandup/vreport/Table.html:41
		view.StreamTimestamp(qw422016, &model.Created)
//line views/vstandup/vreport/Table.html:41
		qw422016.N().S(`</td>
          <td>`)
//line views/vstandup/vreport/Table.html:42
		view.StreamTimestamp(qw422016, model.Updated)
//line views/vstandup/vreport/Table.html:42
		qw422016.N().S(`</td>
        </tr>
`)
//line views/vstandup/vreport/Table.html:44
	}
//line views/vstandup/vreport/Table.html:44
	qw422016.N().S(`      </tbody>
    </table>
  </div>
`)
//line views/vstandup/vreport/Table.html:48
	if prms.HasNextPage(len(models)+prms.Offset) || prms.HasPreviousPage() {
//line views/vstandup/vreport/Table.html:48
		qw422016.N().S(`  <hr />
  `)
//line views/vstandup/vreport/Table.html:50
		components.StreamPagination(qw422016, len(models)+prms.Offset, prms, ps.URI)
//line views/vstandup/vreport/Table.html:50
		qw422016.N().S(`
  <div class="clear"></div>
`)
//line views/vstandup/vreport/Table.html:52
	}
//line views/vstandup/vreport/Table.html:53
}

//line views/vstandup/vreport/Table.html:53
func WriteTable(qq422016 qtio422016.Writer, models report.Reports, standupsByStandupID standup.Standups, usersByUserID user.Users, params filter.ParamSet, as *app.State, ps *cutil.PageState, paths ...string) {
//line views/vstandup/vreport/Table.html:53
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vstandup/vreport/Table.html:53
	StreamTable(qw422016, models, standupsByStandupID, usersByUserID, params, as, ps, paths...)
//line views/vstandup/vreport/Table.html:53
	qt422016.ReleaseWriter(qw422016)
//line views/vstandup/vreport/Table.html:53
}

//line views/vstandup/vreport/Table.html:53
func Table(models report.Reports, standupsByStandupID standup.Standups, usersByUserID user.Users, params filter.ParamSet, as *app.State, ps *cutil.PageState, paths ...string) string {
//line views/vstandup/vreport/Table.html:53
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vstandup/vreport/Table.html:53
	WriteTable(qb422016, models, standupsByStandupID, usersByUserID, params, as, ps, paths...)
//line views/vstandup/vreport/Table.html:53
	qs422016 := string(qb422016.B)
//line views/vstandup/vreport/Table.html:53
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vstandup/vreport/Table.html:53
	return qs422016
//line views/vstandup/vreport/Table.html:53
}
