// Code generated by qtc from "Table.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vstandup/vumember/Table.html:2
package vumember

//line views/vstandup/vumember/Table.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/standup"
	"github.com/kyleu/rituals/app/standup/umember"
	"github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/views/components"
)

//line views/vstandup/vumember/Table.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vstandup/vumember/Table.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vstandup/vumember/Table.html:12
func StreamTable(qw422016 *qt422016.Writer, models umember.StandupMembers, standups standup.Standups, users user.Users, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vstandup/vumember/Table.html:12
	qw422016.N().S(`
`)
//line views/vstandup/vumember/Table.html:13
	prms := params.Get("umember", nil, ps.Logger).Sanitize("umember")

//line views/vstandup/vumember/Table.html:13
	qw422016.N().S(`  <table class="mt">
    <thead>
      <tr>
        `)
//line views/vstandup/vumember/Table.html:17
	components.StreamTableHeaderSimple(qw422016, "umember", "standup_id", "Standup ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vstandup/vumember/Table.html:17
	qw422016.N().S(`
        `)
//line views/vstandup/vumember/Table.html:18
	components.StreamTableHeaderSimple(qw422016, "umember", "user_id", "User ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vstandup/vumember/Table.html:18
	qw422016.N().S(`
        `)
//line views/vstandup/vumember/Table.html:19
	components.StreamTableHeaderSimple(qw422016, "umember", "name", "Name", "String text", prms, ps.URI, ps)
//line views/vstandup/vumember/Table.html:19
	qw422016.N().S(`
        `)
//line views/vstandup/vumember/Table.html:20
	components.StreamTableHeaderSimple(qw422016, "umember", "picture", "Picture", "URL in string form", prms, ps.URI, ps)
//line views/vstandup/vumember/Table.html:20
	qw422016.N().S(`
        `)
//line views/vstandup/vumember/Table.html:21
	components.StreamTableHeaderSimple(qw422016, "umember", "role", "Role", "String text", prms, ps.URI, ps)
//line views/vstandup/vumember/Table.html:21
	qw422016.N().S(`
        `)
//line views/vstandup/vumember/Table.html:22
	components.StreamTableHeaderSimple(qw422016, "umember", "created", "Created", "Date and time, in almost any format", prms, ps.URI, ps)
//line views/vstandup/vumember/Table.html:22
	qw422016.N().S(`
        `)
//line views/vstandup/vumember/Table.html:23
	components.StreamTableHeaderSimple(qw422016, "umember", "updated", "Updated", "Date and time, in almost any format (optional)", prms, ps.URI, ps)
//line views/vstandup/vumember/Table.html:23
	qw422016.N().S(`
      </tr>
    </thead>
    <tbody>
`)
//line views/vstandup/vumember/Table.html:27
	for _, model := range models {
//line views/vstandup/vumember/Table.html:27
		qw422016.N().S(`      <tr>
        <td>
          <div class="icon"><a href="/admin/db/standup/member/`)
//line views/vstandup/vumember/Table.html:30
		components.StreamDisplayUUID(qw422016, &model.StandupID)
//line views/vstandup/vumember/Table.html:30
		qw422016.N().S(`/`)
//line views/vstandup/vumember/Table.html:30
		components.StreamDisplayUUID(qw422016, &model.UserID)
//line views/vstandup/vumember/Table.html:30
		qw422016.N().S(`">`)
//line views/vstandup/vumember/Table.html:30
		components.StreamDisplayUUID(qw422016, &model.StandupID)
//line views/vstandup/vumember/Table.html:30
		if x := standups.Get(model.StandupID); x != nil {
//line views/vstandup/vumember/Table.html:30
			qw422016.N().S(` (`)
//line views/vstandup/vumember/Table.html:30
			qw422016.E().S(x.TitleString())
//line views/vstandup/vumember/Table.html:30
			qw422016.N().S(`)`)
//line views/vstandup/vumember/Table.html:30
		}
//line views/vstandup/vumember/Table.html:30
		qw422016.N().S(`</a></div>
          <a title="Standup" href="`)
//line views/vstandup/vumember/Table.html:31
		qw422016.E().S(`/standup` + `/` + model.StandupID.String())
//line views/vstandup/vumember/Table.html:31
		qw422016.N().S(`">`)
//line views/vstandup/vumember/Table.html:31
		components.StreamSVGRefIcon(qw422016, "star", ps)
//line views/vstandup/vumember/Table.html:31
		qw422016.N().S(`</a>
        </td>
        <td>
          <div class="icon"><a href="/admin/db/standup/member/`)
//line views/vstandup/vumember/Table.html:34
		components.StreamDisplayUUID(qw422016, &model.StandupID)
//line views/vstandup/vumember/Table.html:34
		qw422016.N().S(`/`)
//line views/vstandup/vumember/Table.html:34
		components.StreamDisplayUUID(qw422016, &model.UserID)
//line views/vstandup/vumember/Table.html:34
		qw422016.N().S(`">`)
//line views/vstandup/vumember/Table.html:34
		components.StreamDisplayUUID(qw422016, &model.UserID)
//line views/vstandup/vumember/Table.html:34
		if x := users.Get(model.UserID); x != nil {
//line views/vstandup/vumember/Table.html:34
			qw422016.N().S(` (`)
//line views/vstandup/vumember/Table.html:34
			qw422016.E().S(x.TitleString())
//line views/vstandup/vumember/Table.html:34
			qw422016.N().S(`)`)
//line views/vstandup/vumember/Table.html:34
		}
//line views/vstandup/vumember/Table.html:34
		qw422016.N().S(`</a></div>
          <a title="User" href="`)
//line views/vstandup/vumember/Table.html:35
		qw422016.E().S(`/user` + `/` + model.UserID.String())
//line views/vstandup/vumember/Table.html:35
		qw422016.N().S(`">`)
//line views/vstandup/vumember/Table.html:35
		components.StreamSVGRefIcon(qw422016, "profile", ps)
//line views/vstandup/vumember/Table.html:35
		qw422016.N().S(`</a>
        </td>
        <td><strong>`)
//line views/vstandup/vumember/Table.html:37
		qw422016.E().S(model.Name)
//line views/vstandup/vumember/Table.html:37
		qw422016.N().S(`</strong></td>
        <td><a href="`)
//line views/vstandup/vumember/Table.html:38
		qw422016.E().S(model.Picture)
//line views/vstandup/vumember/Table.html:38
		qw422016.N().S(`" target="_blank">`)
//line views/vstandup/vumember/Table.html:38
		qw422016.E().S(model.Picture)
//line views/vstandup/vumember/Table.html:38
		qw422016.N().S(`</a></td>
        <td>`)
//line views/vstandup/vumember/Table.html:39
		qw422016.E().S(model.Role)
//line views/vstandup/vumember/Table.html:39
		qw422016.N().S(`</td>
        <td>`)
//line views/vstandup/vumember/Table.html:40
		components.StreamDisplayTimestamp(qw422016, &model.Created)
//line views/vstandup/vumember/Table.html:40
		qw422016.N().S(`</td>
        <td>`)
//line views/vstandup/vumember/Table.html:41
		components.StreamDisplayTimestamp(qw422016, model.Updated)
//line views/vstandup/vumember/Table.html:41
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vstandup/vumember/Table.html:43
	}
//line views/vstandup/vumember/Table.html:44
	if prms.HasNextPage(len(models)+prms.Offset) || prms.HasPreviousPage() {
//line views/vstandup/vumember/Table.html:44
		qw422016.N().S(`      <tr>
        <td colspan="7">`)
//line views/vstandup/vumember/Table.html:46
		components.StreamPagination(qw422016, len(models)+prms.Offset, prms, ps.URI)
//line views/vstandup/vumember/Table.html:46
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vstandup/vumember/Table.html:48
	}
//line views/vstandup/vumember/Table.html:48
	qw422016.N().S(`    </tbody>
  </table>
`)
//line views/vstandup/vumember/Table.html:51
}

//line views/vstandup/vumember/Table.html:51
func WriteTable(qq422016 qtio422016.Writer, models umember.StandupMembers, standups standup.Standups, users user.Users, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vstandup/vumember/Table.html:51
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vstandup/vumember/Table.html:51
	StreamTable(qw422016, models, standups, users, params, as, ps)
//line views/vstandup/vumember/Table.html:51
	qt422016.ReleaseWriter(qw422016)
//line views/vstandup/vumember/Table.html:51
}

//line views/vstandup/vumember/Table.html:51
func Table(models umember.StandupMembers, standups standup.Standups, users user.Users, params filter.ParamSet, as *app.State, ps *cutil.PageState) string {
//line views/vstandup/vumember/Table.html:51
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vstandup/vumember/Table.html:51
	WriteTable(qb422016, models, standups, users, params, as, ps)
//line views/vstandup/vumember/Table.html:51
	qs422016 := string(qb422016.B)
//line views/vstandup/vumember/Table.html:51
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vstandup/vumember/Table.html:51
	return qs422016
//line views/vstandup/vumember/Table.html:51
}
