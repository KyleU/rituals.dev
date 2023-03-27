// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vstandup/Detail.html:2
package vstandup

//line views/vstandup/Detail.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/standup"
	"github.com/kyleu/rituals/app/standup/report"
	"github.com/kyleu/rituals/app/standup/uhistory"
	"github.com/kyleu/rituals/app/standup/umember"
	"github.com/kyleu/rituals/app/standup/upermission"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/layout"
	"github.com/kyleu/rituals/views/vstandup/vreport"
	"github.com/kyleu/rituals/views/vstandup/vuhistory"
	"github.com/kyleu/rituals/views/vstandup/vumember"
	"github.com/kyleu/rituals/views/vstandup/vupermission"
)

//line views/vstandup/Detail.html:21
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vstandup/Detail.html:21
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vstandup/Detail.html:21
type Detail struct {
	layout.Basic
	Model                         *standup.Standup
	Teams                         team.Teams
	Sprints                       sprint.Sprints
	Params                        filter.ParamSet
	ReportsByStandupID            report.Reports
	StandupHistoriesByStandupID   uhistory.StandupHistories
	StandupMembersByStandupID     umember.StandupMembers
	StandupPermissionsByStandupID upermission.StandupPermissions
}

//line views/vstandup/Detail.html:33
func (p *Detail) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vstandup/Detail.html:33
	qw422016.N().S(`
  <div class="card">
    <div class="right">
      <a href="#modal-standup"><button type="button">JSON</button></a>
      <a href="`)
//line views/vstandup/Detail.html:37
	qw422016.E().S(p.Model.WebPath())
//line views/vstandup/Detail.html:37
	qw422016.N().S(`/edit"><button>`)
//line views/vstandup/Detail.html:37
	components.StreamSVGRef(qw422016, "edit", 15, 15, "icon", ps)
//line views/vstandup/Detail.html:37
	qw422016.N().S(`Edit</button></a>
    </div>
    <h3>`)
//line views/vstandup/Detail.html:39
	components.StreamSVGRefIcon(qw422016, `standup`, ps)
//line views/vstandup/Detail.html:39
	qw422016.N().S(` `)
//line views/vstandup/Detail.html:39
	qw422016.E().S(p.Model.TitleString())
//line views/vstandup/Detail.html:39
	qw422016.N().S(`</h3>
    <div><a href="/admin/db/standup"><em>Standup</em></a></div>
    <table class="mt">
      <tbody>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">ID</th>
          <td>`)
//line views/vstandup/Detail.html:45
	components.StreamDisplayUUID(qw422016, &p.Model.ID)
//line views/vstandup/Detail.html:45
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Slug</th>
          <td>`)
//line views/vstandup/Detail.html:49
	qw422016.E().S(p.Model.Slug)
//line views/vstandup/Detail.html:49
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Title</th>
          <td><strong>`)
//line views/vstandup/Detail.html:53
	qw422016.E().S(p.Model.Title)
//line views/vstandup/Detail.html:53
	qw422016.N().S(`</strong></td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Icon</th>
          <td>`)
//line views/vstandup/Detail.html:57
	qw422016.E().S(p.Model.Icon)
//line views/vstandup/Detail.html:57
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="Available options: [new, active, complete]">Status</th>
          <td>`)
//line views/vstandup/Detail.html:61
	qw422016.E().V(p.Model.Status)
//line views/vstandup/Detail.html:61
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000) (optional)">Team ID</th>
          <td class="nowrap">
            `)
//line views/vstandup/Detail.html:66
	components.StreamDisplayUUID(qw422016, p.Model.TeamID)
//line views/vstandup/Detail.html:66
	if p.Model.TeamID != nil {
//line views/vstandup/Detail.html:66
		if x := p.Teams.Get(*p.Model.TeamID); x != nil {
//line views/vstandup/Detail.html:66
			qw422016.N().S(` (`)
//line views/vstandup/Detail.html:66
			qw422016.E().S(x.TitleString())
//line views/vstandup/Detail.html:66
			qw422016.N().S(`)`)
//line views/vstandup/Detail.html:66
		}
//line views/vstandup/Detail.html:66
	}
//line views/vstandup/Detail.html:66
	qw422016.N().S(`
            `)
//line views/vstandup/Detail.html:67
	if p.Model.TeamID != nil {
//line views/vstandup/Detail.html:67
		qw422016.N().S(`<a title="Team" href="`)
//line views/vstandup/Detail.html:67
		qw422016.E().S(`/team` + `/` + p.Model.TeamID.String())
//line views/vstandup/Detail.html:67
		qw422016.N().S(`">`)
//line views/vstandup/Detail.html:67
		components.StreamSVGRef(qw422016, "team", 18, 18, "", ps)
//line views/vstandup/Detail.html:67
		qw422016.N().S(`</a>`)
//line views/vstandup/Detail.html:67
	}
//line views/vstandup/Detail.html:67
	qw422016.N().S(`
          </td>
        </tr>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000) (optional)">Sprint ID</th>
          <td class="nowrap">
            `)
//line views/vstandup/Detail.html:73
	components.StreamDisplayUUID(qw422016, p.Model.SprintID)
//line views/vstandup/Detail.html:73
	if p.Model.SprintID != nil {
//line views/vstandup/Detail.html:73
		if x := p.Sprints.Get(*p.Model.SprintID); x != nil {
//line views/vstandup/Detail.html:73
			qw422016.N().S(` (`)
//line views/vstandup/Detail.html:73
			qw422016.E().S(x.TitleString())
//line views/vstandup/Detail.html:73
			qw422016.N().S(`)`)
//line views/vstandup/Detail.html:73
		}
//line views/vstandup/Detail.html:73
	}
//line views/vstandup/Detail.html:73
	qw422016.N().S(`
            `)
//line views/vstandup/Detail.html:74
	if p.Model.SprintID != nil {
//line views/vstandup/Detail.html:74
		qw422016.N().S(`<a title="Sprint" href="`)
//line views/vstandup/Detail.html:74
		qw422016.E().S(`/sprint` + `/` + p.Model.SprintID.String())
//line views/vstandup/Detail.html:74
		qw422016.N().S(`">`)
//line views/vstandup/Detail.html:74
		components.StreamSVGRef(qw422016, "sprint", 18, 18, "", ps)
//line views/vstandup/Detail.html:74
		qw422016.N().S(`</a>`)
//line views/vstandup/Detail.html:74
	}
//line views/vstandup/Detail.html:74
	qw422016.N().S(`
          </td>
        </tr>
        <tr>
          <th class="shrink" title="Date and time, in almost any format">Created</th>
          <td>`)
//line views/vstandup/Detail.html:79
	components.StreamDisplayTimestamp(qw422016, &p.Model.Created)
//line views/vstandup/Detail.html:79
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="Date and time, in almost any format (optional)">Updated</th>
          <td>`)
//line views/vstandup/Detail.html:83
	components.StreamDisplayTimestamp(qw422016, p.Model.Updated)
//line views/vstandup/Detail.html:83
	qw422016.N().S(`</td>
        </tr>
      </tbody>
    </table>
  </div>
`)
//line views/vstandup/Detail.html:90
	if len(p.ReportsByStandupID) > 0 {
//line views/vstandup/Detail.html:90
		qw422016.N().S(`  <div class="card">
    <h3>`)
//line views/vstandup/Detail.html:92
		components.StreamSVGRefIcon(qw422016, `file-alt`, ps)
//line views/vstandup/Detail.html:92
		qw422016.N().S(` Related reports by [standup id]</h3>
    <div class="overflow clear">
      `)
//line views/vstandup/Detail.html:94
		vreport.StreamTable(qw422016, p.ReportsByStandupID, nil, nil, p.Params, as, ps)
//line views/vstandup/Detail.html:94
		qw422016.N().S(`
    </div>
  </div>
`)
//line views/vstandup/Detail.html:97
	}
//line views/vstandup/Detail.html:98
	if len(p.StandupHistoriesByStandupID) > 0 {
//line views/vstandup/Detail.html:98
		qw422016.N().S(`  <div class="card">
    <h3>`)
//line views/vstandup/Detail.html:100
		components.StreamSVGRefIcon(qw422016, `history`, ps)
//line views/vstandup/Detail.html:100
		qw422016.N().S(` Related histories by [standup id]</h3>
    <div class="overflow clear">
      `)
//line views/vstandup/Detail.html:102
		vuhistory.StreamTable(qw422016, p.StandupHistoriesByStandupID, nil, p.Params, as, ps)
//line views/vstandup/Detail.html:102
		qw422016.N().S(`
    </div>
  </div>
`)
//line views/vstandup/Detail.html:105
	}
//line views/vstandup/Detail.html:106
	if len(p.StandupMembersByStandupID) > 0 {
//line views/vstandup/Detail.html:106
		qw422016.N().S(`  <div class="card">
    <h3>`)
//line views/vstandup/Detail.html:108
		components.StreamSVGRefIcon(qw422016, `users`, ps)
//line views/vstandup/Detail.html:108
		qw422016.N().S(` Related members by [standup id]</h3>
    <div class="overflow clear">
      `)
//line views/vstandup/Detail.html:110
		vumember.StreamTable(qw422016, p.StandupMembersByStandupID, nil, nil, p.Params, as, ps)
//line views/vstandup/Detail.html:110
		qw422016.N().S(`
    </div>
  </div>
`)
//line views/vstandup/Detail.html:113
	}
//line views/vstandup/Detail.html:114
	if len(p.StandupPermissionsByStandupID) > 0 {
//line views/vstandup/Detail.html:114
		qw422016.N().S(`  <div class="card">
    <h3>`)
//line views/vstandup/Detail.html:116
		components.StreamSVGRefIcon(qw422016, `permission`, ps)
//line views/vstandup/Detail.html:116
		qw422016.N().S(` Related permissions by [standup id]</h3>
    <div class="overflow clear">
      `)
//line views/vstandup/Detail.html:118
		vupermission.StreamTable(qw422016, p.StandupPermissionsByStandupID, nil, p.Params, as, ps)
//line views/vstandup/Detail.html:118
		qw422016.N().S(`
    </div>
  </div>
`)
//line views/vstandup/Detail.html:121
	}
//line views/vstandup/Detail.html:121
	qw422016.N().S(`  `)
//line views/vstandup/Detail.html:122
	components.StreamJSONModal(qw422016, "standup", "Standup JSON", p.Model, 1)
//line views/vstandup/Detail.html:122
	qw422016.N().S(`
`)
//line views/vstandup/Detail.html:123
}

//line views/vstandup/Detail.html:123
func (p *Detail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vstandup/Detail.html:123
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vstandup/Detail.html:123
	p.StreamBody(qw422016, as, ps)
//line views/vstandup/Detail.html:123
	qt422016.ReleaseWriter(qw422016)
//line views/vstandup/Detail.html:123
}

//line views/vstandup/Detail.html:123
func (p *Detail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vstandup/Detail.html:123
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vstandup/Detail.html:123
	p.WriteBody(qb422016, as, ps)
//line views/vstandup/Detail.html:123
	qs422016 := string(qb422016.B)
//line views/vstandup/Detail.html:123
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vstandup/Detail.html:123
	return qs422016
//line views/vstandup/Detail.html:123
}
