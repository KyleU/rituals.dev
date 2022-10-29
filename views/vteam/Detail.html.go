// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vteam/Detail.html:2
package vteam

//line views/vteam/Detail.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/estimate"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/retro"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/standup"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/team/thistory"
	"github.com/kyleu/rituals/app/team/tmember"
	"github.com/kyleu/rituals/app/team/tpermission"
	"github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/layout"
	"github.com/kyleu/rituals/views/vestimate"
	"github.com/kyleu/rituals/views/vretro"
	"github.com/kyleu/rituals/views/vsprint"
	"github.com/kyleu/rituals/views/vstandup"
	"github.com/kyleu/rituals/views/vteam/vthistory"
	"github.com/kyleu/rituals/views/vteam/vtmember"
	"github.com/kyleu/rituals/views/vteam/vtpermission"
)

//line views/vteam/Detail.html:26
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vteam/Detail.html:26
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vteam/Detail.html:26
type Detail struct {
	layout.Basic
	Model                   *team.Team
	Users                   user.Users
	Params                  filter.ParamSet
	EstimatesByTeamID       estimate.Estimates
	RetrosByTeamID          retro.Retros
	SprintsByTeamID         sprint.Sprints
	StandupsByTeamID        standup.Standups
	TeamHistoriesByTeamID   thistory.TeamHistories
	TeamMembersByTeamID     tmember.TeamMembers
	TeamPermissionsByTeamID tpermission.TeamPermissions
}

//line views/vteam/Detail.html:40
func (p *Detail) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vteam/Detail.html:40
	qw422016.N().S(`
  <div class="card">
    <div class="right">
      <a href="#modal-team"><button type="button">JSON</button></a>
      <a href="`)
//line views/vteam/Detail.html:44
	qw422016.E().S(p.Model.WebPath())
//line views/vteam/Detail.html:44
	qw422016.N().S(`/edit"><button>`)
//line views/vteam/Detail.html:44
	components.StreamSVGRef(qw422016, "edit", 15, 15, "icon", ps)
//line views/vteam/Detail.html:44
	qw422016.N().S(`Edit</button></a>
    </div>
    <h3>`)
//line views/vteam/Detail.html:46
	components.StreamSVGRefIcon(qw422016, `team`, ps)
//line views/vteam/Detail.html:46
	qw422016.N().S(` `)
//line views/vteam/Detail.html:46
	qw422016.E().S(p.Model.TitleString())
//line views/vteam/Detail.html:46
	qw422016.N().S(`</h3>
    <div><a href="/admin/db/team"><em>Team</em></a></div>
    <table class="mt">
      <tbody>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">ID</th>
          <td>`)
//line views/vteam/Detail.html:52
	components.StreamDisplayUUID(qw422016, &p.Model.ID)
//line views/vteam/Detail.html:52
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Slug</th>
          <td>`)
//line views/vteam/Detail.html:56
	qw422016.E().S(p.Model.Slug)
//line views/vteam/Detail.html:56
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Title</th>
          <td><strong>`)
//line views/vteam/Detail.html:60
	qw422016.E().S(p.Model.Title)
//line views/vteam/Detail.html:60
	qw422016.N().S(`</strong></td>
        </tr>
        <tr>
          <th class="shrink" title="Available options: [new, active, complete, deleted]">Status</th>
          <td>`)
//line views/vteam/Detail.html:64
	qw422016.E().V(p.Model.Status)
//line views/vteam/Detail.html:64
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">Owner</th>
          <td>
            <div class="icon">`)
//line views/vteam/Detail.html:69
	components.StreamDisplayUUID(qw422016, &p.Model.Owner)
//line views/vteam/Detail.html:69
	if x := p.Users.Get(p.Model.Owner); x != nil {
//line views/vteam/Detail.html:69
		qw422016.N().S(` (`)
//line views/vteam/Detail.html:69
		qw422016.E().S(x.TitleString())
//line views/vteam/Detail.html:69
		qw422016.N().S(`)`)
//line views/vteam/Detail.html:69
	}
//line views/vteam/Detail.html:69
	qw422016.N().S(`</div>
            <a title="User" href="`)
//line views/vteam/Detail.html:70
	qw422016.E().S(`/user` + `/` + p.Model.Owner.String())
//line views/vteam/Detail.html:70
	qw422016.N().S(`">`)
//line views/vteam/Detail.html:70
	components.StreamSVGRefIcon(qw422016, "profile", ps)
//line views/vteam/Detail.html:70
	qw422016.N().S(`</a>
          </td>
        </tr>
        <tr>
          <th class="shrink" title="Date and time, in almost any format">Created</th>
          <td>`)
//line views/vteam/Detail.html:75
	components.StreamDisplayTimestamp(qw422016, &p.Model.Created)
//line views/vteam/Detail.html:75
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="Date and time, in almost any format (optional)">Updated</th>
          <td>`)
//line views/vteam/Detail.html:79
	components.StreamDisplayTimestamp(qw422016, p.Model.Updated)
//line views/vteam/Detail.html:79
	qw422016.N().S(`</td>
        </tr>
      </tbody>
    </table>
  </div>
`)
//line views/vteam/Detail.html:86
	if len(p.EstimatesByTeamID) > 0 {
//line views/vteam/Detail.html:86
		qw422016.N().S(`  <div class="card">
    <h3>`)
//line views/vteam/Detail.html:88
		components.StreamSVGRefIcon(qw422016, `estimate`, ps)
//line views/vteam/Detail.html:88
		qw422016.N().S(` Related estimates by [team id]</h3>
    <div class="overflow clear">
      `)
//line views/vteam/Detail.html:90
		vestimate.StreamTable(qw422016, p.EstimatesByTeamID, nil, nil, nil, p.Params, as, ps)
//line views/vteam/Detail.html:90
		qw422016.N().S(`
    </div>
  </div>
`)
//line views/vteam/Detail.html:93
	}
//line views/vteam/Detail.html:94
	if len(p.RetrosByTeamID) > 0 {
//line views/vteam/Detail.html:94
		qw422016.N().S(`  <div class="card">
    <h3>`)
//line views/vteam/Detail.html:96
		components.StreamSVGRefIcon(qw422016, `retro`, ps)
//line views/vteam/Detail.html:96
		qw422016.N().S(` Related retros by [team id]</h3>
    <div class="overflow clear">
      `)
//line views/vteam/Detail.html:98
		vretro.StreamTable(qw422016, p.RetrosByTeamID, nil, nil, nil, p.Params, as, ps)
//line views/vteam/Detail.html:98
		qw422016.N().S(`
    </div>
  </div>
`)
//line views/vteam/Detail.html:101
	}
//line views/vteam/Detail.html:102
	if len(p.SprintsByTeamID) > 0 {
//line views/vteam/Detail.html:102
		qw422016.N().S(`  <div class="card">
    <h3>`)
//line views/vteam/Detail.html:104
		components.StreamSVGRefIcon(qw422016, `sprint`, ps)
//line views/vteam/Detail.html:104
		qw422016.N().S(` Related sprints by [team id]</h3>
    <div class="overflow clear">
      `)
//line views/vteam/Detail.html:106
		vsprint.StreamTable(qw422016, p.SprintsByTeamID, nil, nil, p.Params, as, ps)
//line views/vteam/Detail.html:106
		qw422016.N().S(`
    </div>
  </div>
`)
//line views/vteam/Detail.html:109
	}
//line views/vteam/Detail.html:110
	if len(p.StandupsByTeamID) > 0 {
//line views/vteam/Detail.html:110
		qw422016.N().S(`  <div class="card">
    <h3>`)
//line views/vteam/Detail.html:112
		components.StreamSVGRefIcon(qw422016, `standup`, ps)
//line views/vteam/Detail.html:112
		qw422016.N().S(` Related standups by [team id]</h3>
    <div class="overflow clear">
      `)
//line views/vteam/Detail.html:114
		vstandup.StreamTable(qw422016, p.StandupsByTeamID, nil, nil, nil, p.Params, as, ps)
//line views/vteam/Detail.html:114
		qw422016.N().S(`
    </div>
  </div>
`)
//line views/vteam/Detail.html:117
	}
//line views/vteam/Detail.html:118
	if len(p.TeamHistoriesByTeamID) > 0 {
//line views/vteam/Detail.html:118
		qw422016.N().S(`  <div class="card">
    <h3>`)
//line views/vteam/Detail.html:120
		components.StreamSVGRefIcon(qw422016, `clock`, ps)
//line views/vteam/Detail.html:120
		qw422016.N().S(` Related histories by [team id]</h3>
    <div class="overflow clear">
      `)
//line views/vteam/Detail.html:122
		vthistory.StreamTable(qw422016, p.TeamHistoriesByTeamID, nil, p.Params, as, ps)
//line views/vteam/Detail.html:122
		qw422016.N().S(`
    </div>
  </div>
`)
//line views/vteam/Detail.html:125
	}
//line views/vteam/Detail.html:126
	if len(p.TeamMembersByTeamID) > 0 {
//line views/vteam/Detail.html:126
		qw422016.N().S(`  <div class="card">
    <h3>`)
//line views/vteam/Detail.html:128
		components.StreamSVGRefIcon(qw422016, `users`, ps)
//line views/vteam/Detail.html:128
		qw422016.N().S(` Related members by [team id]</h3>
    <div class="overflow clear">
      `)
//line views/vteam/Detail.html:130
		vtmember.StreamTable(qw422016, p.TeamMembersByTeamID, nil, nil, p.Params, as, ps)
//line views/vteam/Detail.html:130
		qw422016.N().S(`
    </div>
  </div>
`)
//line views/vteam/Detail.html:133
	}
//line views/vteam/Detail.html:134
	if len(p.TeamPermissionsByTeamID) > 0 {
//line views/vteam/Detail.html:134
		qw422016.N().S(`  <div class="card">
    <h3>`)
//line views/vteam/Detail.html:136
		components.StreamSVGRefIcon(qw422016, `lock`, ps)
//line views/vteam/Detail.html:136
		qw422016.N().S(` Related permissions by [team id]</h3>
    <div class="overflow clear">
      `)
//line views/vteam/Detail.html:138
		vtpermission.StreamTable(qw422016, p.TeamPermissionsByTeamID, nil, p.Params, as, ps)
//line views/vteam/Detail.html:138
		qw422016.N().S(`
    </div>
  </div>
`)
//line views/vteam/Detail.html:141
	}
//line views/vteam/Detail.html:141
	qw422016.N().S(`  `)
//line views/vteam/Detail.html:142
	components.StreamJSONModal(qw422016, "team", "Team JSON", p.Model, 1)
//line views/vteam/Detail.html:142
	qw422016.N().S(`
`)
//line views/vteam/Detail.html:143
}

//line views/vteam/Detail.html:143
func (p *Detail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vteam/Detail.html:143
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vteam/Detail.html:143
	p.StreamBody(qw422016, as, ps)
//line views/vteam/Detail.html:143
	qt422016.ReleaseWriter(qw422016)
//line views/vteam/Detail.html:143
}

//line views/vteam/Detail.html:143
func (p *Detail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vteam/Detail.html:143
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vteam/Detail.html:143
	p.WriteBody(qb422016, as, ps)
//line views/vteam/Detail.html:143
	qs422016 := string(qb422016.B)
//line views/vteam/Detail.html:143
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vteam/Detail.html:143
	return qs422016
//line views/vteam/Detail.html:143
}
