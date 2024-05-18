// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vestimate/vepermission/Detail.html:2
package vepermission

//line views/vestimate/vepermission/Detail.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/estimate"
	"github.com/kyleu/rituals/app/estimate/epermission"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/components/view"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vestimate/vepermission/Detail.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vestimate/vepermission/Detail.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vestimate/vepermission/Detail.html:12
type Detail struct {
	layout.Basic
	Model                *epermission.EstimatePermission
	EstimateByEstimateID *estimate.Estimate
}

//line views/vestimate/vepermission/Detail.html:18
func (p *Detail) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vestimate/vepermission/Detail.html:18
	qw422016.N().S(`
  <div class="card">
    <div class="right">
      <a href="#modal-estimatePermission"><button type="button">`)
//line views/vestimate/vepermission/Detail.html:21
	components.StreamSVGRef(qw422016, "file", 15, 15, "icon", ps)
//line views/vestimate/vepermission/Detail.html:21
	qw422016.N().S(`JSON</button></a>
      <a href="`)
//line views/vestimate/vepermission/Detail.html:22
	qw422016.E().S(p.Model.WebPath())
//line views/vestimate/vepermission/Detail.html:22
	qw422016.N().S(`/edit"><button>`)
//line views/vestimate/vepermission/Detail.html:22
	components.StreamSVGRef(qw422016, "edit", 15, 15, "icon", ps)
//line views/vestimate/vepermission/Detail.html:22
	qw422016.N().S(`Edit</button></a>
    </div>
    <h3>`)
//line views/vestimate/vepermission/Detail.html:24
	components.StreamSVGRefIcon(qw422016, `permission`, ps)
//line views/vestimate/vepermission/Detail.html:24
	qw422016.N().S(` `)
//line views/vestimate/vepermission/Detail.html:24
	qw422016.E().S(p.Model.TitleString())
//line views/vestimate/vepermission/Detail.html:24
	qw422016.N().S(`</h3>
    <div><a href="/admin/db/estimate/permission"><em>Permission</em></a></div>
    <div class="mt overflow full-width">
      <table>
        <tbody>
          <tr>
            <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">Estimate ID</th>
            <td class="nowrap">
              `)
//line views/vestimate/vepermission/Detail.html:32
	view.StreamUUID(qw422016, &p.Model.EstimateID)
//line views/vestimate/vepermission/Detail.html:32
	if p.EstimateByEstimateID != nil {
//line views/vestimate/vepermission/Detail.html:32
		qw422016.N().S(` (`)
//line views/vestimate/vepermission/Detail.html:32
		qw422016.E().S(p.EstimateByEstimateID.TitleString())
//line views/vestimate/vepermission/Detail.html:32
		qw422016.N().S(`)`)
//line views/vestimate/vepermission/Detail.html:32
	}
//line views/vestimate/vepermission/Detail.html:32
	qw422016.N().S(`
              <a title="Estimate" href="`)
//line views/vestimate/vepermission/Detail.html:33
	qw422016.E().S(`/admin/db/estimate` + `/` + p.Model.EstimateID.String())
//line views/vestimate/vepermission/Detail.html:33
	qw422016.N().S(`">`)
//line views/vestimate/vepermission/Detail.html:33
	components.StreamSVGRef(qw422016, "estimate", 18, 18, "", ps)
//line views/vestimate/vepermission/Detail.html:33
	qw422016.N().S(`</a>
            </td>
          </tr>
          <tr>
            <th class="shrink" title="String text">Key</th>
            <td>`)
//line views/vestimate/vepermission/Detail.html:38
	view.StreamString(qw422016, p.Model.Key)
//line views/vestimate/vepermission/Detail.html:38
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="String text">Value</th>
            <td>`)
//line views/vestimate/vepermission/Detail.html:42
	view.StreamString(qw422016, p.Model.Value)
//line views/vestimate/vepermission/Detail.html:42
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="String text">Access</th>
            <td>`)
//line views/vestimate/vepermission/Detail.html:46
	view.StreamString(qw422016, p.Model.Access)
//line views/vestimate/vepermission/Detail.html:46
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="Date and time, in almost any format">Created</th>
            <td>`)
//line views/vestimate/vepermission/Detail.html:50
	view.StreamTimestamp(qw422016, &p.Model.Created)
//line views/vestimate/vepermission/Detail.html:50
	qw422016.N().S(`</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
`)
//line views/vestimate/vepermission/Detail.html:57
	qw422016.N().S(`  `)
//line views/vestimate/vepermission/Detail.html:58
	components.StreamJSONModal(qw422016, "estimatePermission", "Permission JSON", p.Model, 1)
//line views/vestimate/vepermission/Detail.html:58
	qw422016.N().S(`
`)
//line views/vestimate/vepermission/Detail.html:59
}

//line views/vestimate/vepermission/Detail.html:59
func (p *Detail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vestimate/vepermission/Detail.html:59
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vestimate/vepermission/Detail.html:59
	p.StreamBody(qw422016, as, ps)
//line views/vestimate/vepermission/Detail.html:59
	qt422016.ReleaseWriter(qw422016)
//line views/vestimate/vepermission/Detail.html:59
}

//line views/vestimate/vepermission/Detail.html:59
func (p *Detail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vestimate/vepermission/Detail.html:59
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vestimate/vepermission/Detail.html:59
	p.WriteBody(qb422016, as, ps)
//line views/vestimate/vepermission/Detail.html:59
	qs422016 := string(qb422016.B)
//line views/vestimate/vepermission/Detail.html:59
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vestimate/vepermission/Detail.html:59
	return qs422016
//line views/vestimate/vepermission/Detail.html:59
}
