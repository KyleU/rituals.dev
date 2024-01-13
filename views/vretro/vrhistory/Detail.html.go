// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vretro/vrhistory/Detail.html:2
package vrhistory

//line views/vretro/vrhistory/Detail.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/retro"
	"github.com/kyleu/rituals/app/retro/rhistory"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/components/view"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vretro/vrhistory/Detail.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vretro/vrhistory/Detail.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vretro/vrhistory/Detail.html:12
type Detail struct {
	layout.Basic
	Model          *rhistory.RetroHistory
	RetroByRetroID *retro.Retro
}

//line views/vretro/vrhistory/Detail.html:18
func (p *Detail) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vretro/vrhistory/Detail.html:18
	qw422016.N().S(`
  <div class="card">
    <div class="right">
      <a href="#modal-retroHistory"><button type="button">JSON</button></a>
      <a href="`)
//line views/vretro/vrhistory/Detail.html:22
	qw422016.E().S(p.Model.WebPath())
//line views/vretro/vrhistory/Detail.html:22
	qw422016.N().S(`/edit"><button>`)
//line views/vretro/vrhistory/Detail.html:22
	components.StreamSVGRef(qw422016, "edit", 15, 15, "icon", ps)
//line views/vretro/vrhistory/Detail.html:22
	qw422016.N().S(`Edit</button></a>
    </div>
    <h3>`)
//line views/vretro/vrhistory/Detail.html:24
	components.StreamSVGRefIcon(qw422016, `history`, ps)
//line views/vretro/vrhistory/Detail.html:24
	qw422016.N().S(` `)
//line views/vretro/vrhistory/Detail.html:24
	qw422016.E().S(p.Model.TitleString())
//line views/vretro/vrhistory/Detail.html:24
	qw422016.N().S(`</h3>
    <div><a href="/admin/db/retro/history"><em>History</em></a></div>
    <table class="mt">
      <tbody>
        <tr>
          <th class="shrink" title="String text">Slug</th>
          <td>`)
//line views/vretro/vrhistory/Detail.html:30
	view.StreamString(qw422016, p.Model.Slug)
//line views/vretro/vrhistory/Detail.html:30
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">Retro ID</th>
          <td class="nowrap">
            `)
//line views/vretro/vrhistory/Detail.html:35
	view.StreamUUID(qw422016, &p.Model.RetroID)
//line views/vretro/vrhistory/Detail.html:35
	if p.RetroByRetroID != nil {
//line views/vretro/vrhistory/Detail.html:35
		qw422016.N().S(` (`)
//line views/vretro/vrhistory/Detail.html:35
		qw422016.E().S(p.RetroByRetroID.TitleString())
//line views/vretro/vrhistory/Detail.html:35
		qw422016.N().S(`)`)
//line views/vretro/vrhistory/Detail.html:35
	}
//line views/vretro/vrhistory/Detail.html:35
	qw422016.N().S(`
            <a title="Retro" href="`)
//line views/vretro/vrhistory/Detail.html:36
	qw422016.E().S(`/admin/db/retro` + `/` + p.Model.RetroID.String())
//line views/vretro/vrhistory/Detail.html:36
	qw422016.N().S(`">`)
//line views/vretro/vrhistory/Detail.html:36
	components.StreamSVGRef(qw422016, "retro", 18, 18, "", ps)
//line views/vretro/vrhistory/Detail.html:36
	qw422016.N().S(`</a>
          </td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Retro Name</th>
          <td>`)
//line views/vretro/vrhistory/Detail.html:41
	view.StreamString(qw422016, p.Model.RetroName)
//line views/vretro/vrhistory/Detail.html:41
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="Date and time, in almost any format">Created</th>
          <td>`)
//line views/vretro/vrhistory/Detail.html:45
	view.StreamTimestamp(qw422016, &p.Model.Created)
//line views/vretro/vrhistory/Detail.html:45
	qw422016.N().S(`</td>
        </tr>
      </tbody>
    </table>
  </div>
`)
//line views/vretro/vrhistory/Detail.html:51
	qw422016.N().S(`  `)
//line views/vretro/vrhistory/Detail.html:52
	components.StreamJSONModal(qw422016, "retroHistory", "History JSON", p.Model, 1)
//line views/vretro/vrhistory/Detail.html:52
	qw422016.N().S(`
`)
//line views/vretro/vrhistory/Detail.html:53
}

//line views/vretro/vrhistory/Detail.html:53
func (p *Detail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vretro/vrhistory/Detail.html:53
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vretro/vrhistory/Detail.html:53
	p.StreamBody(qw422016, as, ps)
//line views/vretro/vrhistory/Detail.html:53
	qt422016.ReleaseWriter(qw422016)
//line views/vretro/vrhistory/Detail.html:53
}

//line views/vretro/vrhistory/Detail.html:53
func (p *Detail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vretro/vrhistory/Detail.html:53
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vretro/vrhistory/Detail.html:53
	p.WriteBody(qb422016, as, ps)
//line views/vretro/vrhistory/Detail.html:53
	qs422016 := string(qb422016.B)
//line views/vretro/vrhistory/Detail.html:53
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vretro/vrhistory/Detail.html:53
	return qs422016
//line views/vretro/vrhistory/Detail.html:53
}
