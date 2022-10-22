// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vteam/vthistory/Detail.html:2
package vthistory

//line views/vteam/vthistory/Detail.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/team/thistory"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vteam/vthistory/Detail.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vteam/vthistory/Detail.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vteam/vthistory/Detail.html:11
type Detail struct {
	layout.Basic
	Model *thistory.TeamHistory
	Teams team.Teams
}

//line views/vteam/vthistory/Detail.html:17
func (p *Detail) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vteam/vthistory/Detail.html:17
	qw422016.N().S(`
  <div class="card">
    <div class="right">
      <a href="#modal-teamHistory"><button type="button">JSON</button></a>
      <a href="`)
//line views/vteam/vthistory/Detail.html:21
	qw422016.E().S(p.Model.WebPath())
//line views/vteam/vthistory/Detail.html:21
	qw422016.N().S(`/edit"><button>`)
//line views/vteam/vthistory/Detail.html:21
	components.StreamSVGRef(qw422016, "edit", 15, 15, "icon", ps)
//line views/vteam/vthistory/Detail.html:21
	qw422016.N().S(`Edit</button></a>
    </div>
    <h3>`)
//line views/vteam/vthistory/Detail.html:23
	components.StreamSVGRefIcon(qw422016, `clock`, ps)
//line views/vteam/vthistory/Detail.html:23
	qw422016.N().S(` `)
//line views/vteam/vthistory/Detail.html:23
	qw422016.E().S(p.Model.TitleString())
//line views/vteam/vthistory/Detail.html:23
	qw422016.N().S(`</h3>
    <div><a href="/admin/db/team/history"><em>History</em></a></div>
    <table class="mt">
      <tbody>
        <tr>
          <th class="shrink" title="String text">Slug</th>
          <td>`)
//line views/vteam/vthistory/Detail.html:29
	qw422016.E().S(p.Model.Slug)
//line views/vteam/vthistory/Detail.html:29
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">Team ID</th>
          <td>
            <div class="icon">`)
//line views/vteam/vthistory/Detail.html:34
	components.StreamDisplayUUID(qw422016, &p.Model.TeamID)
//line views/vteam/vthistory/Detail.html:34
	if x := p.Teams.Get(p.Model.TeamID); x != nil {
//line views/vteam/vthistory/Detail.html:34
		qw422016.N().S(` (`)
//line views/vteam/vthistory/Detail.html:34
		qw422016.E().S(x.TitleString())
//line views/vteam/vthistory/Detail.html:34
		qw422016.N().S(`)`)
//line views/vteam/vthistory/Detail.html:34
	}
//line views/vteam/vthistory/Detail.html:34
	qw422016.N().S(`</div>
            <a title="Team" href="`)
//line views/vteam/vthistory/Detail.html:35
	qw422016.E().S(`/team` + `/` + p.Model.TeamID.String())
//line views/vteam/vthistory/Detail.html:35
	qw422016.N().S(`">`)
//line views/vteam/vthistory/Detail.html:35
	components.StreamSVGRefIcon(qw422016, "users", ps)
//line views/vteam/vthistory/Detail.html:35
	qw422016.N().S(`</a>
          </td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Team Name</th>
          <td>`)
//line views/vteam/vthistory/Detail.html:40
	qw422016.E().S(p.Model.TeamName)
//line views/vteam/vthistory/Detail.html:40
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="Date and time, in almost any format">Created</th>
          <td>`)
//line views/vteam/vthistory/Detail.html:44
	components.StreamDisplayTimestamp(qw422016, &p.Model.Created)
//line views/vteam/vthistory/Detail.html:44
	qw422016.N().S(`</td>
        </tr>
      </tbody>
    </table>
  </div>
`)
//line views/vteam/vthistory/Detail.html:50
	qw422016.N().S(`  `)
//line views/vteam/vthistory/Detail.html:51
	components.StreamJSONModal(qw422016, "teamHistory", "History JSON", p.Model, 1)
//line views/vteam/vthistory/Detail.html:51
	qw422016.N().S(`
`)
//line views/vteam/vthistory/Detail.html:52
}

//line views/vteam/vthistory/Detail.html:52
func (p *Detail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vteam/vthistory/Detail.html:52
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vteam/vthistory/Detail.html:52
	p.StreamBody(qw422016, as, ps)
//line views/vteam/vthistory/Detail.html:52
	qt422016.ReleaseWriter(qw422016)
//line views/vteam/vthistory/Detail.html:52
}

//line views/vteam/vthistory/Detail.html:52
func (p *Detail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vteam/vthistory/Detail.html:52
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vteam/vthistory/Detail.html:52
	p.WriteBody(qb422016, as, ps)
//line views/vteam/vthistory/Detail.html:52
	qs422016 := string(qb422016.B)
//line views/vteam/vthistory/Detail.html:52
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vteam/vthistory/Detail.html:52
	return qs422016
//line views/vteam/vthistory/Detail.html:52
}
