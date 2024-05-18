// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vstandup/vuhistory/Detail.html:2
package vuhistory

//line views/vstandup/vuhistory/Detail.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/standup"
	"github.com/kyleu/rituals/app/standup/uhistory"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/components/view"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vstandup/vuhistory/Detail.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vstandup/vuhistory/Detail.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vstandup/vuhistory/Detail.html:12
type Detail struct {
	layout.Basic
	Model              *uhistory.StandupHistory
	StandupByStandupID *standup.Standup
}

//line views/vstandup/vuhistory/Detail.html:18
func (p *Detail) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vstandup/vuhistory/Detail.html:18
	qw422016.N().S(`
  <div class="card">
    <div class="right">
      <a href="#modal-standupHistory"><button type="button">`)
//line views/vstandup/vuhistory/Detail.html:21
	components.StreamSVGRef(qw422016, "file", 15, 15, "icon", ps)
//line views/vstandup/vuhistory/Detail.html:21
	qw422016.N().S(`JSON</button></a>
      <a href="`)
//line views/vstandup/vuhistory/Detail.html:22
	qw422016.E().S(p.Model.WebPath())
//line views/vstandup/vuhistory/Detail.html:22
	qw422016.N().S(`/edit"><button>`)
//line views/vstandup/vuhistory/Detail.html:22
	components.StreamSVGRef(qw422016, "edit", 15, 15, "icon", ps)
//line views/vstandup/vuhistory/Detail.html:22
	qw422016.N().S(`Edit</button></a>
    </div>
    <h3>`)
//line views/vstandup/vuhistory/Detail.html:24
	components.StreamSVGRefIcon(qw422016, `history`, ps)
//line views/vstandup/vuhistory/Detail.html:24
	qw422016.N().S(` `)
//line views/vstandup/vuhistory/Detail.html:24
	qw422016.E().S(p.Model.TitleString())
//line views/vstandup/vuhistory/Detail.html:24
	qw422016.N().S(`</h3>
    <div><a href="/admin/db/standup/history"><em>History</em></a></div>
    <div class="mt overflow full-width">
      <table>
        <tbody>
          <tr>
            <th class="shrink" title="String text">Slug</th>
            <td>`)
//line views/vstandup/vuhistory/Detail.html:31
	view.StreamString(qw422016, p.Model.Slug)
//line views/vstandup/vuhistory/Detail.html:31
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">Standup ID</th>
            <td class="nowrap">
              `)
//line views/vstandup/vuhistory/Detail.html:36
	view.StreamUUID(qw422016, &p.Model.StandupID)
//line views/vstandup/vuhistory/Detail.html:36
	if p.StandupByStandupID != nil {
//line views/vstandup/vuhistory/Detail.html:36
		qw422016.N().S(` (`)
//line views/vstandup/vuhistory/Detail.html:36
		qw422016.E().S(p.StandupByStandupID.TitleString())
//line views/vstandup/vuhistory/Detail.html:36
		qw422016.N().S(`)`)
//line views/vstandup/vuhistory/Detail.html:36
	}
//line views/vstandup/vuhistory/Detail.html:36
	qw422016.N().S(`
              <a title="Standup" href="`)
//line views/vstandup/vuhistory/Detail.html:37
	qw422016.E().S(`/admin/db/standup` + `/` + p.Model.StandupID.String())
//line views/vstandup/vuhistory/Detail.html:37
	qw422016.N().S(`">`)
//line views/vstandup/vuhistory/Detail.html:37
	components.StreamSVGRef(qw422016, "standup", 18, 18, "", ps)
//line views/vstandup/vuhistory/Detail.html:37
	qw422016.N().S(`</a>
            </td>
          </tr>
          <tr>
            <th class="shrink" title="String text">Standup Name</th>
            <td>`)
//line views/vstandup/vuhistory/Detail.html:42
	view.StreamString(qw422016, p.Model.StandupName)
//line views/vstandup/vuhistory/Detail.html:42
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="Date and time, in almost any format">Created</th>
            <td>`)
//line views/vstandup/vuhistory/Detail.html:46
	view.StreamTimestamp(qw422016, &p.Model.Created)
//line views/vstandup/vuhistory/Detail.html:46
	qw422016.N().S(`</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
`)
//line views/vstandup/vuhistory/Detail.html:53
	qw422016.N().S(`  `)
//line views/vstandup/vuhistory/Detail.html:54
	components.StreamJSONModal(qw422016, "standupHistory", "History JSON", p.Model, 1)
//line views/vstandup/vuhistory/Detail.html:54
	qw422016.N().S(`
`)
//line views/vstandup/vuhistory/Detail.html:55
}

//line views/vstandup/vuhistory/Detail.html:55
func (p *Detail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vstandup/vuhistory/Detail.html:55
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vstandup/vuhistory/Detail.html:55
	p.StreamBody(qw422016, as, ps)
//line views/vstandup/vuhistory/Detail.html:55
	qt422016.ReleaseWriter(qw422016)
//line views/vstandup/vuhistory/Detail.html:55
}

//line views/vstandup/vuhistory/Detail.html:55
func (p *Detail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vstandup/vuhistory/Detail.html:55
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vstandup/vuhistory/Detail.html:55
	p.WriteBody(qb422016, as, ps)
//line views/vstandup/vuhistory/Detail.html:55
	qs422016 := string(qb422016.B)
//line views/vstandup/vuhistory/Detail.html:55
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vstandup/vuhistory/Detail.html:55
	return qs422016
//line views/vstandup/vuhistory/Detail.html:55
}
