// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vstandup/vupermission/Detail.html:2
package vupermission

//line views/vstandup/vupermission/Detail.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/standup"
	"github.com/kyleu/rituals/app/standup/upermission"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vstandup/vupermission/Detail.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vstandup/vupermission/Detail.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vstandup/vupermission/Detail.html:11
type Detail struct {
	layout.Basic
	Model               *upermission.StandupPermission
	StandupsByStandupID standup.Standups
}

//line views/vstandup/vupermission/Detail.html:17
func (p *Detail) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vstandup/vupermission/Detail.html:17
	qw422016.N().S(`
  <div class="card">
    <div class="right">
      <a href="#modal-standupPermission"><button type="button">JSON</button></a>
      <a href="`)
//line views/vstandup/vupermission/Detail.html:21
	qw422016.E().S(p.Model.WebPath())
//line views/vstandup/vupermission/Detail.html:21
	qw422016.N().S(`/edit"><button>`)
//line views/vstandup/vupermission/Detail.html:21
	components.StreamSVGRef(qw422016, "edit", 15, 15, "icon", ps)
//line views/vstandup/vupermission/Detail.html:21
	qw422016.N().S(`Edit</button></a>
    </div>
    <h3>`)
//line views/vstandup/vupermission/Detail.html:23
	components.StreamSVGRefIcon(qw422016, `permission`, ps)
//line views/vstandup/vupermission/Detail.html:23
	qw422016.N().S(` `)
//line views/vstandup/vupermission/Detail.html:23
	qw422016.E().S(p.Model.TitleString())
//line views/vstandup/vupermission/Detail.html:23
	qw422016.N().S(`</h3>
    <div><a href="/admin/db/standup/permission"><em>Permission</em></a></div>
    <table class="mt">
      <tbody>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">Standup ID</th>
          <td class="nowrap">
            `)
//line views/vstandup/vupermission/Detail.html:30
	components.StreamDisplayUUID(qw422016, &p.Model.StandupID)
//line views/vstandup/vupermission/Detail.html:30
	if x := p.StandupsByStandupID.Get(p.Model.StandupID); x != nil {
//line views/vstandup/vupermission/Detail.html:30
		qw422016.N().S(` (`)
//line views/vstandup/vupermission/Detail.html:30
		qw422016.E().S(x.TitleString())
//line views/vstandup/vupermission/Detail.html:30
		qw422016.N().S(`)`)
//line views/vstandup/vupermission/Detail.html:30
	}
//line views/vstandup/vupermission/Detail.html:30
	qw422016.N().S(`
            <a title="Standup" href="`)
//line views/vstandup/vupermission/Detail.html:31
	qw422016.E().S(`/standup` + `/` + p.Model.StandupID.String())
//line views/vstandup/vupermission/Detail.html:31
	qw422016.N().S(`">`)
//line views/vstandup/vupermission/Detail.html:31
	components.StreamSVGRef(qw422016, "standup", 18, 18, "", ps)
//line views/vstandup/vupermission/Detail.html:31
	qw422016.N().S(`</a>
          </td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Key</th>
          <td>`)
//line views/vstandup/vupermission/Detail.html:36
	qw422016.E().S(p.Model.Key)
//line views/vstandup/vupermission/Detail.html:36
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Value</th>
          <td>`)
//line views/vstandup/vupermission/Detail.html:40
	qw422016.E().S(p.Model.Value)
//line views/vstandup/vupermission/Detail.html:40
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Access</th>
          <td>`)
//line views/vstandup/vupermission/Detail.html:44
	qw422016.E().S(p.Model.Access)
//line views/vstandup/vupermission/Detail.html:44
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="Date and time, in almost any format">Created</th>
          <td>`)
//line views/vstandup/vupermission/Detail.html:48
	components.StreamDisplayTimestamp(qw422016, &p.Model.Created)
//line views/vstandup/vupermission/Detail.html:48
	qw422016.N().S(`</td>
        </tr>
      </tbody>
    </table>
  </div>
`)
//line views/vstandup/vupermission/Detail.html:54
	qw422016.N().S(`  `)
//line views/vstandup/vupermission/Detail.html:55
	components.StreamJSONModal(qw422016, "standupPermission", "Permission JSON", p.Model, 1)
//line views/vstandup/vupermission/Detail.html:55
	qw422016.N().S(`
`)
//line views/vstandup/vupermission/Detail.html:56
}

//line views/vstandup/vupermission/Detail.html:56
func (p *Detail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vstandup/vupermission/Detail.html:56
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vstandup/vupermission/Detail.html:56
	p.StreamBody(qw422016, as, ps)
//line views/vstandup/vupermission/Detail.html:56
	qt422016.ReleaseWriter(qw422016)
//line views/vstandup/vupermission/Detail.html:56
}

//line views/vstandup/vupermission/Detail.html:56
func (p *Detail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vstandup/vupermission/Detail.html:56
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vstandup/vupermission/Detail.html:56
	p.WriteBody(qb422016, as, ps)
//line views/vstandup/vupermission/Detail.html:56
	qs422016 := string(qb422016.B)
//line views/vstandup/vupermission/Detail.html:56
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vstandup/vupermission/Detail.html:56
	return qs422016
//line views/vstandup/vupermission/Detail.html:56
}
