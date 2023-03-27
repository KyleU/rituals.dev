// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vestimate/Detail.html:2
package vestimate

//line views/vestimate/Detail.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/estimate"
	"github.com/kyleu/rituals/app/estimate/ehistory"
	"github.com/kyleu/rituals/app/estimate/emember"
	"github.com/kyleu/rituals/app/estimate/epermission"
	"github.com/kyleu/rituals/app/estimate/story"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/layout"
	"github.com/kyleu/rituals/views/vestimate/vehistory"
	"github.com/kyleu/rituals/views/vestimate/vemember"
	"github.com/kyleu/rituals/views/vestimate/vepermission"
	"github.com/kyleu/rituals/views/vestimate/vstory"
)

//line views/vestimate/Detail.html:21
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vestimate/Detail.html:21
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vestimate/Detail.html:21
type Detail struct {
	layout.Basic
	Model                           *estimate.Estimate
	Teams                           team.Teams
	Sprints                         sprint.Sprints
	Params                          filter.ParamSet
	EstimateHistoriesByEstimateID   ehistory.EstimateHistories
	EstimateMembersByEstimateID     emember.EstimateMembers
	EstimatePermissionsByEstimateID epermission.EstimatePermissions
	StoriesByEstimateID             story.Stories
}

//line views/vestimate/Detail.html:33
func (p *Detail) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vestimate/Detail.html:33
	qw422016.N().S(`
  <div class="card">
    <div class="right">
      <a href="#modal-estimate"><button type="button">JSON</button></a>
      <a href="`)
//line views/vestimate/Detail.html:37
	qw422016.E().S(p.Model.WebPath())
//line views/vestimate/Detail.html:37
	qw422016.N().S(`/edit"><button>`)
//line views/vestimate/Detail.html:37
	components.StreamSVGRef(qw422016, "edit", 15, 15, "icon", ps)
//line views/vestimate/Detail.html:37
	qw422016.N().S(`Edit</button></a>
    </div>
    <h3>`)
//line views/vestimate/Detail.html:39
	components.StreamSVGRefIcon(qw422016, `estimate`, ps)
//line views/vestimate/Detail.html:39
	qw422016.N().S(` `)
//line views/vestimate/Detail.html:39
	qw422016.E().S(p.Model.TitleString())
//line views/vestimate/Detail.html:39
	qw422016.N().S(`</h3>
    <div><a href="/admin/db/estimate"><em>Estimate</em></a></div>
    <table class="mt">
      <tbody>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">ID</th>
          <td>`)
//line views/vestimate/Detail.html:45
	components.StreamDisplayUUID(qw422016, &p.Model.ID)
//line views/vestimate/Detail.html:45
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Slug</th>
          <td>`)
//line views/vestimate/Detail.html:49
	qw422016.E().S(p.Model.Slug)
//line views/vestimate/Detail.html:49
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Title</th>
          <td><strong>`)
//line views/vestimate/Detail.html:53
	qw422016.E().S(p.Model.Title)
//line views/vestimate/Detail.html:53
	qw422016.N().S(`</strong></td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Icon</th>
          <td>`)
//line views/vestimate/Detail.html:57
	qw422016.E().S(p.Model.Icon)
//line views/vestimate/Detail.html:57
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="Available options: [new, active, complete]">Status</th>
          <td>`)
//line views/vestimate/Detail.html:61
	qw422016.E().V(p.Model.Status)
//line views/vestimate/Detail.html:61
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000) (optional)">Team ID</th>
          <td class="nowrap">
            `)
//line views/vestimate/Detail.html:66
	components.StreamDisplayUUID(qw422016, p.Model.TeamID)
//line views/vestimate/Detail.html:66
	if p.Model.TeamID != nil {
//line views/vestimate/Detail.html:66
		if x := p.Teams.Get(*p.Model.TeamID); x != nil {
//line views/vestimate/Detail.html:66
			qw422016.N().S(` (`)
//line views/vestimate/Detail.html:66
			qw422016.E().S(x.TitleString())
//line views/vestimate/Detail.html:66
			qw422016.N().S(`)`)
//line views/vestimate/Detail.html:66
		}
//line views/vestimate/Detail.html:66
	}
//line views/vestimate/Detail.html:66
	qw422016.N().S(`
            `)
//line views/vestimate/Detail.html:67
	if p.Model.TeamID != nil {
//line views/vestimate/Detail.html:67
		qw422016.N().S(`<a title="Team" href="`)
//line views/vestimate/Detail.html:67
		qw422016.E().S(`/team` + `/` + p.Model.TeamID.String())
//line views/vestimate/Detail.html:67
		qw422016.N().S(`">`)
//line views/vestimate/Detail.html:67
		components.StreamSVGRef(qw422016, "team", 18, 18, "", ps)
//line views/vestimate/Detail.html:67
		qw422016.N().S(`</a>`)
//line views/vestimate/Detail.html:67
	}
//line views/vestimate/Detail.html:67
	qw422016.N().S(`
          </td>
        </tr>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000) (optional)">Sprint ID</th>
          <td class="nowrap">
            `)
//line views/vestimate/Detail.html:73
	components.StreamDisplayUUID(qw422016, p.Model.SprintID)
//line views/vestimate/Detail.html:73
	if p.Model.SprintID != nil {
//line views/vestimate/Detail.html:73
		if x := p.Sprints.Get(*p.Model.SprintID); x != nil {
//line views/vestimate/Detail.html:73
			qw422016.N().S(` (`)
//line views/vestimate/Detail.html:73
			qw422016.E().S(x.TitleString())
//line views/vestimate/Detail.html:73
			qw422016.N().S(`)`)
//line views/vestimate/Detail.html:73
		}
//line views/vestimate/Detail.html:73
	}
//line views/vestimate/Detail.html:73
	qw422016.N().S(`
            `)
//line views/vestimate/Detail.html:74
	if p.Model.SprintID != nil {
//line views/vestimate/Detail.html:74
		qw422016.N().S(`<a title="Sprint" href="`)
//line views/vestimate/Detail.html:74
		qw422016.E().S(`/sprint` + `/` + p.Model.SprintID.String())
//line views/vestimate/Detail.html:74
		qw422016.N().S(`">`)
//line views/vestimate/Detail.html:74
		components.StreamSVGRef(qw422016, "sprint", 18, 18, "", ps)
//line views/vestimate/Detail.html:74
		qw422016.N().S(`</a>`)
//line views/vestimate/Detail.html:74
	}
//line views/vestimate/Detail.html:74
	qw422016.N().S(`
          </td>
        </tr>
        <tr>
          <th class="shrink" title="Comma-separated list of values">Choices</th>
          <td>`)
//line views/vestimate/Detail.html:79
	components.StreamDisplayStringArray(qw422016, p.Model.Choices)
//line views/vestimate/Detail.html:79
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="Date and time, in almost any format">Created</th>
          <td>`)
//line views/vestimate/Detail.html:83
	components.StreamDisplayTimestamp(qw422016, &p.Model.Created)
//line views/vestimate/Detail.html:83
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="Date and time, in almost any format (optional)">Updated</th>
          <td>`)
//line views/vestimate/Detail.html:87
	components.StreamDisplayTimestamp(qw422016, p.Model.Updated)
//line views/vestimate/Detail.html:87
	qw422016.N().S(`</td>
        </tr>
      </tbody>
    </table>
  </div>
`)
//line views/vestimate/Detail.html:94
	if len(p.EstimateHistoriesByEstimateID) > 0 {
//line views/vestimate/Detail.html:94
		qw422016.N().S(`  <div class="card">
    <h3>`)
//line views/vestimate/Detail.html:96
		components.StreamSVGRefIcon(qw422016, `history`, ps)
//line views/vestimate/Detail.html:96
		qw422016.N().S(` Related histories by [estimate id]</h3>
    <div class="overflow clear">
      `)
//line views/vestimate/Detail.html:98
		vehistory.StreamTable(qw422016, p.EstimateHistoriesByEstimateID, nil, p.Params, as, ps)
//line views/vestimate/Detail.html:98
		qw422016.N().S(`
    </div>
  </div>
`)
//line views/vestimate/Detail.html:101
	}
//line views/vestimate/Detail.html:102
	if len(p.EstimateMembersByEstimateID) > 0 {
//line views/vestimate/Detail.html:102
		qw422016.N().S(`  <div class="card">
    <h3>`)
//line views/vestimate/Detail.html:104
		components.StreamSVGRefIcon(qw422016, `users`, ps)
//line views/vestimate/Detail.html:104
		qw422016.N().S(` Related members by [estimate id]</h3>
    <div class="overflow clear">
      `)
//line views/vestimate/Detail.html:106
		vemember.StreamTable(qw422016, p.EstimateMembersByEstimateID, nil, nil, p.Params, as, ps)
//line views/vestimate/Detail.html:106
		qw422016.N().S(`
    </div>
  </div>
`)
//line views/vestimate/Detail.html:109
	}
//line views/vestimate/Detail.html:110
	if len(p.EstimatePermissionsByEstimateID) > 0 {
//line views/vestimate/Detail.html:110
		qw422016.N().S(`  <div class="card">
    <h3>`)
//line views/vestimate/Detail.html:112
		components.StreamSVGRefIcon(qw422016, `permission`, ps)
//line views/vestimate/Detail.html:112
		qw422016.N().S(` Related permissions by [estimate id]</h3>
    <div class="overflow clear">
      `)
//line views/vestimate/Detail.html:114
		vepermission.StreamTable(qw422016, p.EstimatePermissionsByEstimateID, nil, p.Params, as, ps)
//line views/vestimate/Detail.html:114
		qw422016.N().S(`
    </div>
  </div>
`)
//line views/vestimate/Detail.html:117
	}
//line views/vestimate/Detail.html:118
	if len(p.StoriesByEstimateID) > 0 {
//line views/vestimate/Detail.html:118
		qw422016.N().S(`  <div class="card">
    <h3>`)
//line views/vestimate/Detail.html:120
		components.StreamSVGRefIcon(qw422016, `story`, ps)
//line views/vestimate/Detail.html:120
		qw422016.N().S(` Related stories by [estimate id]</h3>
    <div class="overflow clear">
      `)
//line views/vestimate/Detail.html:122
		vstory.StreamTable(qw422016, p.StoriesByEstimateID, nil, nil, p.Params, as, ps)
//line views/vestimate/Detail.html:122
		qw422016.N().S(`
    </div>
  </div>
`)
//line views/vestimate/Detail.html:125
	}
//line views/vestimate/Detail.html:125
	qw422016.N().S(`  `)
//line views/vestimate/Detail.html:126
	components.StreamJSONModal(qw422016, "estimate", "Estimate JSON", p.Model, 1)
//line views/vestimate/Detail.html:126
	qw422016.N().S(`
`)
//line views/vestimate/Detail.html:127
}

//line views/vestimate/Detail.html:127
func (p *Detail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vestimate/Detail.html:127
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vestimate/Detail.html:127
	p.StreamBody(qw422016, as, ps)
//line views/vestimate/Detail.html:127
	qt422016.ReleaseWriter(qw422016)
//line views/vestimate/Detail.html:127
}

//line views/vestimate/Detail.html:127
func (p *Detail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vestimate/Detail.html:127
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vestimate/Detail.html:127
	p.WriteBody(qb422016, as, ps)
//line views/vestimate/Detail.html:127
	qs422016 := string(qb422016.B)
//line views/vestimate/Detail.html:127
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vestimate/Detail.html:127
	return qs422016
//line views/vestimate/Detail.html:127
}
