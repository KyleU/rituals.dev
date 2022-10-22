// Code generated by qtc from "Table.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vuser/Table.html:2
package vuser

//line views/vuser/Table.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/views/components"
)

//line views/vuser/Table.html:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vuser/Table.html:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vuser/Table.html:10
func StreamTable(qw422016 *qt422016.Writer, models user.Users, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vuser/Table.html:10
	qw422016.N().S(`
`)
//line views/vuser/Table.html:11
	prms := params.Get("user", nil, ps.Logger).Sanitize("user")

//line views/vuser/Table.html:11
	qw422016.N().S(`  <table class="mt">
    <thead>
      <tr>
        `)
//line views/vuser/Table.html:15
	components.StreamTableHeaderSimple(qw422016, "user", "id", "ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vuser/Table.html:15
	qw422016.N().S(`
        `)
//line views/vuser/Table.html:16
	components.StreamTableHeaderSimple(qw422016, "user", "name", "Name", "String text", prms, ps.URI, ps)
//line views/vuser/Table.html:16
	qw422016.N().S(`
        `)
//line views/vuser/Table.html:17
	components.StreamTableHeaderSimple(qw422016, "user", "role", "Role", "String text", prms, ps.URI, ps)
//line views/vuser/Table.html:17
	qw422016.N().S(`
        `)
//line views/vuser/Table.html:18
	components.StreamTableHeaderSimple(qw422016, "user", "picture", "Picture", "URL in string form", prms, ps.URI, ps)
//line views/vuser/Table.html:18
	qw422016.N().S(`
        `)
//line views/vuser/Table.html:19
	components.StreamTableHeaderSimple(qw422016, "user", "created", "Created", "Date and time, in almost any format", prms, ps.URI, ps)
//line views/vuser/Table.html:19
	qw422016.N().S(`
        `)
//line views/vuser/Table.html:20
	components.StreamTableHeaderSimple(qw422016, "user", "updated", "Updated", "Date and time, in almost any format (optional)", prms, ps.URI, ps)
//line views/vuser/Table.html:20
	qw422016.N().S(`
      </tr>
    </thead>
    <tbody>
`)
//line views/vuser/Table.html:24
	for _, model := range models {
//line views/vuser/Table.html:24
		qw422016.N().S(`      <tr>
        <td><a href="/admin/db/user/`)
//line views/vuser/Table.html:26
		components.StreamDisplayUUID(qw422016, &model.ID)
//line views/vuser/Table.html:26
		qw422016.N().S(`">`)
//line views/vuser/Table.html:26
		components.StreamDisplayUUID(qw422016, &model.ID)
//line views/vuser/Table.html:26
		qw422016.N().S(`</a></td>
        <td><strong>`)
//line views/vuser/Table.html:27
		qw422016.E().S(model.Name)
//line views/vuser/Table.html:27
		qw422016.N().S(`</strong></td>
        <td>`)
//line views/vuser/Table.html:28
		qw422016.E().S(model.Role)
//line views/vuser/Table.html:28
		qw422016.N().S(`</td>
        <td><a href="`)
//line views/vuser/Table.html:29
		qw422016.E().S(model.Picture)
//line views/vuser/Table.html:29
		qw422016.N().S(`" target="_blank">`)
//line views/vuser/Table.html:29
		qw422016.E().S(model.Picture)
//line views/vuser/Table.html:29
		qw422016.N().S(`</a></td>
        <td>`)
//line views/vuser/Table.html:30
		components.StreamDisplayTimestamp(qw422016, &model.Created)
//line views/vuser/Table.html:30
		qw422016.N().S(`</td>
        <td>`)
//line views/vuser/Table.html:31
		components.StreamDisplayTimestamp(qw422016, model.Updated)
//line views/vuser/Table.html:31
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vuser/Table.html:33
	}
//line views/vuser/Table.html:34
	if prms.HasNextPage(len(models)+prms.Offset) || prms.HasPreviousPage() {
//line views/vuser/Table.html:34
		qw422016.N().S(`      <tr>
        <td colspan="6">`)
//line views/vuser/Table.html:36
		components.StreamPagination(qw422016, len(models)+prms.Offset, prms, ps.URI)
//line views/vuser/Table.html:36
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vuser/Table.html:38
	}
//line views/vuser/Table.html:38
	qw422016.N().S(`    </tbody>
  </table>
`)
//line views/vuser/Table.html:41
}

//line views/vuser/Table.html:41
func WriteTable(qq422016 qtio422016.Writer, models user.Users, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vuser/Table.html:41
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vuser/Table.html:41
	StreamTable(qw422016, models, params, as, ps)
//line views/vuser/Table.html:41
	qt422016.ReleaseWriter(qw422016)
//line views/vuser/Table.html:41
}

//line views/vuser/Table.html:41
func Table(models user.Users, params filter.ParamSet, as *app.State, ps *cutil.PageState) string {
//line views/vuser/Table.html:41
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vuser/Table.html:41
	WriteTable(qb422016, models, params, as, ps)
//line views/vuser/Table.html:41
	qs422016 := string(qb422016.B)
//line views/vuser/Table.html:41
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vuser/Table.html:41
	return qs422016
//line views/vuser/Table.html:41
}
