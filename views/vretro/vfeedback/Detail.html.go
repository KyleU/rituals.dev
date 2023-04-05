// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vretro/vfeedback/Detail.html:2
package vfeedback

//line views/vretro/vfeedback/Detail.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/retro"
	"github.com/kyleu/rituals/app/retro/feedback"
	"github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vretro/vfeedback/Detail.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vretro/vfeedback/Detail.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vretro/vfeedback/Detail.html:12
type Detail struct {
	layout.Basic
	Model          *feedback.Feedback
	RetroByRetroID *retro.Retro
	UserByUserID   *user.User
}

//line views/vretro/vfeedback/Detail.html:19
func (p *Detail) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vretro/vfeedback/Detail.html:19
	qw422016.N().S(`
  <div class="card">
    <div class="right">
      <a href="#modal-feedback"><button type="button">JSON</button></a>
      <a href="`)
//line views/vretro/vfeedback/Detail.html:23
	qw422016.E().S(p.Model.WebPath())
//line views/vretro/vfeedback/Detail.html:23
	qw422016.N().S(`/edit"><button>`)
//line views/vretro/vfeedback/Detail.html:23
	components.StreamSVGRef(qw422016, "edit", 15, 15, "icon", ps)
//line views/vretro/vfeedback/Detail.html:23
	qw422016.N().S(`Edit</button></a>
    </div>
    <h3>`)
//line views/vretro/vfeedback/Detail.html:25
	components.StreamSVGRefIcon(qw422016, `comment`, ps)
//line views/vretro/vfeedback/Detail.html:25
	qw422016.N().S(` `)
//line views/vretro/vfeedback/Detail.html:25
	qw422016.E().S(p.Model.TitleString())
//line views/vretro/vfeedback/Detail.html:25
	qw422016.N().S(`</h3>
    <div><a href="/admin/db/retro/feedback"><em>Feedback</em></a></div>
    <table class="mt">
      <tbody>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">ID</th>
          <td>`)
//line views/vretro/vfeedback/Detail.html:31
	components.StreamDisplayUUID(qw422016, &p.Model.ID)
//line views/vretro/vfeedback/Detail.html:31
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">Retro ID</th>
          <td class="nowrap">
            `)
//line views/vretro/vfeedback/Detail.html:36
	components.StreamDisplayUUID(qw422016, &p.Model.RetroID)
//line views/vretro/vfeedback/Detail.html:36
	if p.RetroByRetroID != nil {
//line views/vretro/vfeedback/Detail.html:36
		qw422016.N().S(` (`)
//line views/vretro/vfeedback/Detail.html:36
		qw422016.E().S(p.RetroByRetroID.TitleString())
//line views/vretro/vfeedback/Detail.html:36
		qw422016.N().S(`)`)
//line views/vretro/vfeedback/Detail.html:36
	}
//line views/vretro/vfeedback/Detail.html:36
	qw422016.N().S(`
            <a title="Retro" href="`)
//line views/vretro/vfeedback/Detail.html:37
	qw422016.E().S(`/admin/db/retro` + `/` + p.Model.RetroID.String())
//line views/vretro/vfeedback/Detail.html:37
	qw422016.N().S(`">`)
//line views/vretro/vfeedback/Detail.html:37
	components.StreamSVGRef(qw422016, "retro", 18, 18, "", ps)
//line views/vretro/vfeedback/Detail.html:37
	qw422016.N().S(`</a>
          </td>
        </tr>
        <tr>
          <th class="shrink" title="Integer">Idx</th>
          <td>`)
//line views/vretro/vfeedback/Detail.html:42
	qw422016.N().D(p.Model.Idx)
//line views/vretro/vfeedback/Detail.html:42
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">User ID</th>
          <td class="nowrap">
            `)
//line views/vretro/vfeedback/Detail.html:47
	components.StreamDisplayUUID(qw422016, &p.Model.UserID)
//line views/vretro/vfeedback/Detail.html:47
	if p.UserByUserID != nil {
//line views/vretro/vfeedback/Detail.html:47
		qw422016.N().S(` (`)
//line views/vretro/vfeedback/Detail.html:47
		qw422016.E().S(p.UserByUserID.TitleString())
//line views/vretro/vfeedback/Detail.html:47
		qw422016.N().S(`)`)
//line views/vretro/vfeedback/Detail.html:47
	}
//line views/vretro/vfeedback/Detail.html:47
	qw422016.N().S(`
            <a title="User" href="`)
//line views/vretro/vfeedback/Detail.html:48
	qw422016.E().S(`/admin/db/user` + `/` + p.Model.UserID.String())
//line views/vretro/vfeedback/Detail.html:48
	qw422016.N().S(`">`)
//line views/vretro/vfeedback/Detail.html:48
	components.StreamSVGRef(qw422016, "profile", 18, 18, "", ps)
//line views/vretro/vfeedback/Detail.html:48
	qw422016.N().S(`</a>
          </td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Category</th>
          <td>`)
//line views/vretro/vfeedback/Detail.html:53
	qw422016.E().S(p.Model.Category)
//line views/vretro/vfeedback/Detail.html:53
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Content</th>
          <td>`)
//line views/vretro/vfeedback/Detail.html:57
	qw422016.E().S(p.Model.Content)
//line views/vretro/vfeedback/Detail.html:57
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="String text">HTML</th>
          <td><pre>`)
//line views/vretro/vfeedback/Detail.html:61
	qw422016.E().S(p.Model.HTML)
//line views/vretro/vfeedback/Detail.html:61
	qw422016.N().S(`</pre></td>
        </tr>
        <tr>
          <th class="shrink" title="Date and time, in almost any format">Created</th>
          <td>`)
//line views/vretro/vfeedback/Detail.html:65
	components.StreamDisplayTimestamp(qw422016, &p.Model.Created)
//line views/vretro/vfeedback/Detail.html:65
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="Date and time, in almost any format (optional)">Updated</th>
          <td>`)
//line views/vretro/vfeedback/Detail.html:69
	components.StreamDisplayTimestamp(qw422016, p.Model.Updated)
//line views/vretro/vfeedback/Detail.html:69
	qw422016.N().S(`</td>
        </tr>
      </tbody>
    </table>
  </div>
`)
//line views/vretro/vfeedback/Detail.html:75
	qw422016.N().S(`  `)
//line views/vretro/vfeedback/Detail.html:76
	components.StreamJSONModal(qw422016, "feedback", "Feedback JSON", p.Model, 1)
//line views/vretro/vfeedback/Detail.html:76
	qw422016.N().S(`
`)
//line views/vretro/vfeedback/Detail.html:77
}

//line views/vretro/vfeedback/Detail.html:77
func (p *Detail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vretro/vfeedback/Detail.html:77
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vretro/vfeedback/Detail.html:77
	p.StreamBody(qw422016, as, ps)
//line views/vretro/vfeedback/Detail.html:77
	qt422016.ReleaseWriter(qw422016)
//line views/vretro/vfeedback/Detail.html:77
}

//line views/vretro/vfeedback/Detail.html:77
func (p *Detail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vretro/vfeedback/Detail.html:77
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vretro/vfeedback/Detail.html:77
	p.WriteBody(qb422016, as, ps)
//line views/vretro/vfeedback/Detail.html:77
	qs422016 := string(qb422016.B)
//line views/vretro/vfeedback/Detail.html:77
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vretro/vfeedback/Detail.html:77
	return qs422016
//line views/vretro/vfeedback/Detail.html:77
}
