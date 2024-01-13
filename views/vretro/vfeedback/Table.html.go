// Code generated by qtc from "Table.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vretro/vfeedback/Table.html:2
package vfeedback

//line views/vretro/vfeedback/Table.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/retro"
	"github.com/kyleu/rituals/app/retro/feedback"
	"github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/components/view"
)

//line views/vretro/vfeedback/Table.html:13
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vretro/vfeedback/Table.html:13
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vretro/vfeedback/Table.html:13
func StreamTable(qw422016 *qt422016.Writer, models feedback.Feedbacks, retrosByRetroID retro.Retros, usersByUserID user.Users, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vretro/vfeedback/Table.html:13
	qw422016.N().S(`
`)
//line views/vretro/vfeedback/Table.html:14
	prms := params.Get("feedback", nil, ps.Logger).Sanitize("feedback")

//line views/vretro/vfeedback/Table.html:14
	qw422016.N().S(`  <table>
    <thead>
      <tr>
        `)
//line views/vretro/vfeedback/Table.html:18
	components.StreamTableHeaderSimple(qw422016, "feedback", "id", "ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vretro/vfeedback/Table.html:18
	qw422016.N().S(`
        `)
//line views/vretro/vfeedback/Table.html:19
	components.StreamTableHeaderSimple(qw422016, "feedback", "retro_id", "Retro ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vretro/vfeedback/Table.html:19
	qw422016.N().S(`
        `)
//line views/vretro/vfeedback/Table.html:20
	components.StreamTableHeaderSimple(qw422016, "feedback", "idx", "Idx", "Integer", prms, ps.URI, ps)
//line views/vretro/vfeedback/Table.html:20
	qw422016.N().S(`
        `)
//line views/vretro/vfeedback/Table.html:21
	components.StreamTableHeaderSimple(qw422016, "feedback", "user_id", "User ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vretro/vfeedback/Table.html:21
	qw422016.N().S(`
        `)
//line views/vretro/vfeedback/Table.html:22
	components.StreamTableHeaderSimple(qw422016, "feedback", "category", "Category", "String text", prms, ps.URI, ps)
//line views/vretro/vfeedback/Table.html:22
	qw422016.N().S(`
        `)
//line views/vretro/vfeedback/Table.html:23
	components.StreamTableHeaderSimple(qw422016, "feedback", "content", "Content", "String text", prms, ps.URI, ps)
//line views/vretro/vfeedback/Table.html:23
	qw422016.N().S(`
        `)
//line views/vretro/vfeedback/Table.html:24
	components.StreamTableHeaderSimple(qw422016, "feedback", "created", "Created", "Date and time, in almost any format", prms, ps.URI, ps)
//line views/vretro/vfeedback/Table.html:24
	qw422016.N().S(`
        `)
//line views/vretro/vfeedback/Table.html:25
	components.StreamTableHeaderSimple(qw422016, "feedback", "updated", "Updated", "Date and time, in almost any format (optional)", prms, ps.URI, ps)
//line views/vretro/vfeedback/Table.html:25
	qw422016.N().S(`
      </tr>
    </thead>
    <tbody>
`)
//line views/vretro/vfeedback/Table.html:29
	for _, model := range models {
//line views/vretro/vfeedback/Table.html:29
		qw422016.N().S(`      <tr>
        <td><a href="/admin/db/retro/feedback/`)
//line views/vretro/vfeedback/Table.html:31
		view.StreamUUID(qw422016, &model.ID)
//line views/vretro/vfeedback/Table.html:31
		qw422016.N().S(`">`)
//line views/vretro/vfeedback/Table.html:31
		view.StreamUUID(qw422016, &model.ID)
//line views/vretro/vfeedback/Table.html:31
		qw422016.N().S(`</a></td>
        <td class="nowrap">
          `)
//line views/vretro/vfeedback/Table.html:33
		view.StreamUUID(qw422016, &model.RetroID)
//line views/vretro/vfeedback/Table.html:33
		if x := retrosByRetroID.Get(model.RetroID); x != nil {
//line views/vretro/vfeedback/Table.html:33
			qw422016.N().S(` (`)
//line views/vretro/vfeedback/Table.html:33
			qw422016.E().S(x.TitleString())
//line views/vretro/vfeedback/Table.html:33
			qw422016.N().S(`)`)
//line views/vretro/vfeedback/Table.html:33
		}
//line views/vretro/vfeedback/Table.html:33
		qw422016.N().S(`
          <a title="Retro" href="`)
//line views/vretro/vfeedback/Table.html:34
		qw422016.E().S(`/admin/db/retro` + `/` + model.RetroID.String())
//line views/vretro/vfeedback/Table.html:34
		qw422016.N().S(`">`)
//line views/vretro/vfeedback/Table.html:34
		components.StreamSVGRef(qw422016, "retro", 18, 18, "", ps)
//line views/vretro/vfeedback/Table.html:34
		qw422016.N().S(`</a>
        </td>
        <td>`)
//line views/vretro/vfeedback/Table.html:36
		qw422016.N().D(model.Idx)
//line views/vretro/vfeedback/Table.html:36
		qw422016.N().S(`</td>
        <td class="nowrap">
          `)
//line views/vretro/vfeedback/Table.html:38
		view.StreamUUID(qw422016, &model.UserID)
//line views/vretro/vfeedback/Table.html:38
		if x := usersByUserID.Get(model.UserID); x != nil {
//line views/vretro/vfeedback/Table.html:38
			qw422016.N().S(` (`)
//line views/vretro/vfeedback/Table.html:38
			qw422016.E().S(x.TitleString())
//line views/vretro/vfeedback/Table.html:38
			qw422016.N().S(`)`)
//line views/vretro/vfeedback/Table.html:38
		}
//line views/vretro/vfeedback/Table.html:38
		qw422016.N().S(`
          <a title="User" href="`)
//line views/vretro/vfeedback/Table.html:39
		qw422016.E().S(`/admin/db/user` + `/` + model.UserID.String())
//line views/vretro/vfeedback/Table.html:39
		qw422016.N().S(`">`)
//line views/vretro/vfeedback/Table.html:39
		components.StreamSVGRef(qw422016, "profile", 18, 18, "", ps)
//line views/vretro/vfeedback/Table.html:39
		qw422016.N().S(`</a>
        </td>
        <td>`)
//line views/vretro/vfeedback/Table.html:41
		view.StreamString(qw422016, model.Category)
//line views/vretro/vfeedback/Table.html:41
		qw422016.N().S(`</td>
        <td>`)
//line views/vretro/vfeedback/Table.html:42
		view.StreamString(qw422016, model.Content)
//line views/vretro/vfeedback/Table.html:42
		qw422016.N().S(`</td>
        <td>`)
//line views/vretro/vfeedback/Table.html:43
		view.StreamTimestamp(qw422016, &model.Created)
//line views/vretro/vfeedback/Table.html:43
		qw422016.N().S(`</td>
        <td>`)
//line views/vretro/vfeedback/Table.html:44
		view.StreamTimestamp(qw422016, model.Updated)
//line views/vretro/vfeedback/Table.html:44
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vretro/vfeedback/Table.html:46
	}
//line views/vretro/vfeedback/Table.html:47
	if prms.HasNextPage(len(models)+prms.Offset) || prms.HasPreviousPage() {
//line views/vretro/vfeedback/Table.html:47
		qw422016.N().S(`      <tr>
        <td colspan="8">`)
//line views/vretro/vfeedback/Table.html:49
		components.StreamPagination(qw422016, len(models)+prms.Offset, prms, ps.URI)
//line views/vretro/vfeedback/Table.html:49
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vretro/vfeedback/Table.html:51
	}
//line views/vretro/vfeedback/Table.html:51
	qw422016.N().S(`    </tbody>
  </table>
`)
//line views/vretro/vfeedback/Table.html:54
}

//line views/vretro/vfeedback/Table.html:54
func WriteTable(qq422016 qtio422016.Writer, models feedback.Feedbacks, retrosByRetroID retro.Retros, usersByUserID user.Users, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vretro/vfeedback/Table.html:54
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vretro/vfeedback/Table.html:54
	StreamTable(qw422016, models, retrosByRetroID, usersByUserID, params, as, ps)
//line views/vretro/vfeedback/Table.html:54
	qt422016.ReleaseWriter(qw422016)
//line views/vretro/vfeedback/Table.html:54
}

//line views/vretro/vfeedback/Table.html:54
func Table(models feedback.Feedbacks, retrosByRetroID retro.Retros, usersByUserID user.Users, params filter.ParamSet, as *app.State, ps *cutil.PageState) string {
//line views/vretro/vfeedback/Table.html:54
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vretro/vfeedback/Table.html:54
	WriteTable(qb422016, models, retrosByRetroID, usersByUserID, params, as, ps)
//line views/vretro/vfeedback/Table.html:54
	qs422016 := string(qb422016.B)
//line views/vretro/vfeedback/Table.html:54
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vretro/vfeedback/Table.html:54
	return qs422016
//line views/vretro/vfeedback/Table.html:54
}
