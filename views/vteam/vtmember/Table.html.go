// Code generated by qtc from "Table.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vteam/vtmember/Table.html:1
package vtmember

//line views/vteam/vtmember/Table.html:1
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/team/tmember"
	"github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/components/view"
)

//line views/vteam/vtmember/Table.html:13
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vteam/vtmember/Table.html:13
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vteam/vtmember/Table.html:13
func StreamTable(qw422016 *qt422016.Writer, models tmember.TeamMembers, teamsByTeamID team.Teams, usersByUserID user.Users, params filter.ParamSet, as *app.State, ps *cutil.PageState, paths ...string) {
//line views/vteam/vtmember/Table.html:13
	qw422016.N().S(`
`)
//line views/vteam/vtmember/Table.html:14
	prms := params.Sanitized("tmember", ps.Logger)

//line views/vteam/vtmember/Table.html:14
	qw422016.N().S(`  <div class="overflow clear">
    <table>
      <thead>
        <tr>
          `)
//line views/vteam/vtmember/Table.html:19
	components.StreamTableHeaderSimple(qw422016, "tmember", "team_id", "Team ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vteam/vtmember/Table.html:19
	qw422016.N().S(`
          `)
//line views/vteam/vtmember/Table.html:20
	components.StreamTableHeaderSimple(qw422016, "tmember", "user_id", "User ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vteam/vtmember/Table.html:20
	qw422016.N().S(`
          `)
//line views/vteam/vtmember/Table.html:21
	components.StreamTableHeaderSimple(qw422016, "tmember", "name", "Name", "String text", prms, ps.URI, ps)
//line views/vteam/vtmember/Table.html:21
	qw422016.N().S(`
          `)
//line views/vteam/vtmember/Table.html:22
	components.StreamTableHeaderSimple(qw422016, "tmember", "picture", "Picture", "URL in string form", prms, ps.URI, ps)
//line views/vteam/vtmember/Table.html:22
	qw422016.N().S(`
          `)
//line views/vteam/vtmember/Table.html:23
	components.StreamTableHeaderSimple(qw422016, "tmember", "role", "Role", enum.AllMemberStatuses.Help(), prms, ps.URI, ps)
//line views/vteam/vtmember/Table.html:23
	qw422016.N().S(`
          `)
//line views/vteam/vtmember/Table.html:24
	components.StreamTableHeaderSimple(qw422016, "tmember", "created", "Created", "Date and time, in almost any format", prms, ps.URI, ps)
//line views/vteam/vtmember/Table.html:24
	qw422016.N().S(`
          `)
//line views/vteam/vtmember/Table.html:25
	components.StreamTableHeaderSimple(qw422016, "tmember", "updated", "Updated", "Date and time, in almost any format (optional)", prms, ps.URI, ps)
//line views/vteam/vtmember/Table.html:25
	qw422016.N().S(`
        </tr>
      </thead>
      <tbody>
`)
//line views/vteam/vtmember/Table.html:29
	for _, model := range models {
//line views/vteam/vtmember/Table.html:29
		qw422016.N().S(`        <tr>
          <td class="nowrap">
            <a href="`)
//line views/vteam/vtmember/Table.html:32
		qw422016.E().S(model.WebPath(paths...))
//line views/vteam/vtmember/Table.html:32
		qw422016.N().S(`">`)
//line views/vteam/vtmember/Table.html:32
		view.StreamUUID(qw422016, &model.TeamID)
//line views/vteam/vtmember/Table.html:32
		if x := teamsByTeamID.Get(model.TeamID); x != nil {
//line views/vteam/vtmember/Table.html:32
			qw422016.N().S(` (`)
//line views/vteam/vtmember/Table.html:32
			qw422016.E().S(x.TitleString())
//line views/vteam/vtmember/Table.html:32
			qw422016.N().S(`)`)
//line views/vteam/vtmember/Table.html:32
		}
//line views/vteam/vtmember/Table.html:32
		qw422016.N().S(`</a>
            <a title="Team" href="`)
//line views/vteam/vtmember/Table.html:33
		qw422016.E().S(model.WebPath(paths...))
//line views/vteam/vtmember/Table.html:33
		qw422016.N().S(`">`)
//line views/vteam/vtmember/Table.html:33
		components.StreamSVGLink(qw422016, `team`, ps)
//line views/vteam/vtmember/Table.html:33
		qw422016.N().S(`</a>
          </td>
          <td class="nowrap">
            <a href="`)
//line views/vteam/vtmember/Table.html:36
		qw422016.E().S(model.WebPath(paths...))
//line views/vteam/vtmember/Table.html:36
		qw422016.N().S(`">`)
//line views/vteam/vtmember/Table.html:36
		view.StreamUUID(qw422016, &model.UserID)
//line views/vteam/vtmember/Table.html:36
		if x := usersByUserID.Get(model.UserID); x != nil {
//line views/vteam/vtmember/Table.html:36
			qw422016.N().S(` (`)
//line views/vteam/vtmember/Table.html:36
			qw422016.E().S(x.TitleString())
//line views/vteam/vtmember/Table.html:36
			qw422016.N().S(`)`)
//line views/vteam/vtmember/Table.html:36
		}
//line views/vteam/vtmember/Table.html:36
		qw422016.N().S(`</a>
            <a title="User" href="`)
//line views/vteam/vtmember/Table.html:37
		qw422016.E().S(model.WebPath(paths...))
//line views/vteam/vtmember/Table.html:37
		qw422016.N().S(`">`)
//line views/vteam/vtmember/Table.html:37
		components.StreamSVGLink(qw422016, `profile`, ps)
//line views/vteam/vtmember/Table.html:37
		qw422016.N().S(`</a>
          </td>
          <td><strong>`)
//line views/vteam/vtmember/Table.html:39
		view.StreamString(qw422016, model.Name)
//line views/vteam/vtmember/Table.html:39
		qw422016.N().S(`</strong></td>
          <td><a href="`)
//line views/vteam/vtmember/Table.html:40
		qw422016.E().S(model.Picture)
//line views/vteam/vtmember/Table.html:40
		qw422016.N().S(`" target="_blank" rel="noopener noreferrer">`)
//line views/vteam/vtmember/Table.html:40
		qw422016.E().S(model.Picture)
//line views/vteam/vtmember/Table.html:40
		qw422016.N().S(`</a></td>
          <td>`)
//line views/vteam/vtmember/Table.html:41
		qw422016.E().S(model.Role.String())
//line views/vteam/vtmember/Table.html:41
		qw422016.N().S(`</td>
          <td>`)
//line views/vteam/vtmember/Table.html:42
		view.StreamTimestamp(qw422016, &model.Created)
//line views/vteam/vtmember/Table.html:42
		qw422016.N().S(`</td>
          <td>`)
//line views/vteam/vtmember/Table.html:43
		view.StreamTimestamp(qw422016, model.Updated)
//line views/vteam/vtmember/Table.html:43
		qw422016.N().S(`</td>
        </tr>
`)
//line views/vteam/vtmember/Table.html:45
	}
//line views/vteam/vtmember/Table.html:45
	qw422016.N().S(`      </tbody>
    </table>
  </div>
`)
//line views/vteam/vtmember/Table.html:49
	if prms.HasNextPage(len(models)+prms.Offset) || prms.HasPreviousPage() {
//line views/vteam/vtmember/Table.html:49
		qw422016.N().S(`  <hr />
  `)
//line views/vteam/vtmember/Table.html:51
		components.StreamPagination(qw422016, len(models)+prms.Offset, prms, ps.URI)
//line views/vteam/vtmember/Table.html:51
		qw422016.N().S(`
  <div class="clear"></div>
`)
//line views/vteam/vtmember/Table.html:53
	}
//line views/vteam/vtmember/Table.html:54
}

//line views/vteam/vtmember/Table.html:54
func WriteTable(qq422016 qtio422016.Writer, models tmember.TeamMembers, teamsByTeamID team.Teams, usersByUserID user.Users, params filter.ParamSet, as *app.State, ps *cutil.PageState, paths ...string) {
//line views/vteam/vtmember/Table.html:54
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vteam/vtmember/Table.html:54
	StreamTable(qw422016, models, teamsByTeamID, usersByUserID, params, as, ps, paths...)
//line views/vteam/vtmember/Table.html:54
	qt422016.ReleaseWriter(qw422016)
//line views/vteam/vtmember/Table.html:54
}

//line views/vteam/vtmember/Table.html:54
func Table(models tmember.TeamMembers, teamsByTeamID team.Teams, usersByUserID user.Users, params filter.ParamSet, as *app.State, ps *cutil.PageState, paths ...string) string {
//line views/vteam/vtmember/Table.html:54
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vteam/vtmember/Table.html:54
	WriteTable(qb422016, models, teamsByTeamID, usersByUserID, params, as, ps, paths...)
//line views/vteam/vtmember/Table.html:54
	qs422016 := string(qb422016.B)
//line views/vteam/vtmember/Table.html:54
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vteam/vtmember/Table.html:54
	return qs422016
//line views/vteam/vtmember/Table.html:54
}
