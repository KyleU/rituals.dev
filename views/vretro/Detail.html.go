// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vretro/Detail.html:2
package vretro

//line views/vretro/Detail.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/retro"
	"github.com/kyleu/rituals/app/retro/feedback"
	"github.com/kyleu/rituals/app/retro/rhistory"
	"github.com/kyleu/rituals/app/retro/rmember"
	"github.com/kyleu/rituals/app/retro/rpermission"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/layout"
	"github.com/kyleu/rituals/views/vretro/vfeedback"
	"github.com/kyleu/rituals/views/vretro/vrhistory"
	"github.com/kyleu/rituals/views/vretro/vrmember"
	"github.com/kyleu/rituals/views/vretro/vrpermission"
)

//line views/vretro/Detail.html:22
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vretro/Detail.html:22
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vretro/Detail.html:22
type Detail struct {
	layout.Basic
	Model                     *retro.Retro
	Users                     user.Users
	Teams                     team.Teams
	Sprints                   sprint.Sprints
	Params                    filter.ParamSet
	FeedbacksByRetroID        feedback.Feedbacks
	RetroHistoriesByRetroID   rhistory.RetroHistories
	RetroMembersByRetroID     rmember.RetroMembers
	RetroPermissionsByRetroID rpermission.RetroPermissions
}

//line views/vretro/Detail.html:35
func (p *Detail) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vretro/Detail.html:35
	qw422016.N().S(`
  <div class="card">
    <div class="right">
      <a href="#modal-retro"><button type="button">JSON</button></a>
      <a href="`)
//line views/vretro/Detail.html:39
	qw422016.E().S(p.Model.WebPath())
//line views/vretro/Detail.html:39
	qw422016.N().S(`/edit"><button>`)
//line views/vretro/Detail.html:39
	components.StreamSVGRef(qw422016, "edit", 15, 15, "icon", ps)
//line views/vretro/Detail.html:39
	qw422016.N().S(`Edit</button></a>
    </div>
    <h3>`)
//line views/vretro/Detail.html:41
	components.StreamSVGRefIcon(qw422016, `star`, ps)
//line views/vretro/Detail.html:41
	qw422016.N().S(` `)
//line views/vretro/Detail.html:41
	qw422016.E().S(p.Model.TitleString())
//line views/vretro/Detail.html:41
	qw422016.N().S(`</h3>
    <div><a href="/admin/db/retro"><em>Retro</em></a></div>
    <table class="mt">
      <tbody>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">ID</th>
          <td>`)
//line views/vretro/Detail.html:47
	components.StreamDisplayUUID(qw422016, &p.Model.ID)
//line views/vretro/Detail.html:47
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Slug</th>
          <td>`)
//line views/vretro/Detail.html:51
	qw422016.E().S(p.Model.Slug)
//line views/vretro/Detail.html:51
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Title</th>
          <td><strong>`)
//line views/vretro/Detail.html:55
	qw422016.E().S(p.Model.Title)
//line views/vretro/Detail.html:55
	qw422016.N().S(`</strong></td>
        </tr>
        <tr>
          <th class="shrink" title="Available options: [new, active, complete, deleted]">Status</th>
          <td>`)
//line views/vretro/Detail.html:59
	qw422016.E().V(p.Model.Status)
//line views/vretro/Detail.html:59
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000) (optional)">Team ID</th>
          <td>
            <div class="icon">`)
//line views/vretro/Detail.html:64
	components.StreamDisplayUUID(qw422016, p.Model.TeamID)
//line views/vretro/Detail.html:64
	if p.Model.TeamID != nil {
//line views/vretro/Detail.html:64
		if x := p.Teams.Get(*p.Model.TeamID); x != nil {
//line views/vretro/Detail.html:64
			qw422016.N().S(` (`)
//line views/vretro/Detail.html:64
			qw422016.E().S(x.TitleString())
//line views/vretro/Detail.html:64
			qw422016.N().S(`)`)
//line views/vretro/Detail.html:64
		}
//line views/vretro/Detail.html:64
	}
//line views/vretro/Detail.html:64
	qw422016.N().S(`</div>
            <a title="Team" href="`)
//line views/vretro/Detail.html:65
	qw422016.E().S(`/team` + `/` + p.Model.TeamID.String())
//line views/vretro/Detail.html:65
	qw422016.N().S(`">`)
//line views/vretro/Detail.html:65
	components.StreamSVGRefIcon(qw422016, "users", ps)
//line views/vretro/Detail.html:65
	qw422016.N().S(`</a>
          </td>
        </tr>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000) (optional)">Sprint ID</th>
          <td>
            <div class="icon">`)
//line views/vretro/Detail.html:71
	components.StreamDisplayUUID(qw422016, p.Model.SprintID)
//line views/vretro/Detail.html:71
	if p.Model.SprintID != nil {
//line views/vretro/Detail.html:71
		if x := p.Sprints.Get(*p.Model.SprintID); x != nil {
//line views/vretro/Detail.html:71
			qw422016.N().S(` (`)
//line views/vretro/Detail.html:71
			qw422016.E().S(x.TitleString())
//line views/vretro/Detail.html:71
			qw422016.N().S(`)`)
//line views/vretro/Detail.html:71
		}
//line views/vretro/Detail.html:71
	}
//line views/vretro/Detail.html:71
	qw422016.N().S(`</div>
            <a title="Sprint" href="`)
//line views/vretro/Detail.html:72
	qw422016.E().S(`/sprint` + `/` + p.Model.SprintID.String())
//line views/vretro/Detail.html:72
	qw422016.N().S(`">`)
//line views/vretro/Detail.html:72
	components.StreamSVGRefIcon(qw422016, "star", ps)
//line views/vretro/Detail.html:72
	qw422016.N().S(`</a>
          </td>
        </tr>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">Owner</th>
          <td>
            <div class="icon">`)
//line views/vretro/Detail.html:78
	components.StreamDisplayUUID(qw422016, &p.Model.Owner)
//line views/vretro/Detail.html:78
	if x := p.Users.Get(p.Model.Owner); x != nil {
//line views/vretro/Detail.html:78
		qw422016.N().S(` (`)
//line views/vretro/Detail.html:78
		qw422016.E().S(x.TitleString())
//line views/vretro/Detail.html:78
		qw422016.N().S(`)`)
//line views/vretro/Detail.html:78
	}
//line views/vretro/Detail.html:78
	qw422016.N().S(`</div>
            <a title="User" href="`)
//line views/vretro/Detail.html:79
	qw422016.E().S(`/user` + `/` + p.Model.Owner.String())
//line views/vretro/Detail.html:79
	qw422016.N().S(`">`)
//line views/vretro/Detail.html:79
	components.StreamSVGRefIcon(qw422016, "profile", ps)
//line views/vretro/Detail.html:79
	qw422016.N().S(`</a>
          </td>
        </tr>
        <tr>
          <th class="shrink" title="Comma-separated list of values">Categories</th>
          <td>`)
//line views/vretro/Detail.html:84
	components.StreamDisplayStringArray(qw422016, p.Model.Categories)
//line views/vretro/Detail.html:84
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="Date and time, in almost any format">Created</th>
          <td>`)
//line views/vretro/Detail.html:88
	components.StreamDisplayTimestamp(qw422016, &p.Model.Created)
//line views/vretro/Detail.html:88
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="Date and time, in almost any format (optional)">Updated</th>
          <td>`)
//line views/vretro/Detail.html:92
	components.StreamDisplayTimestamp(qw422016, p.Model.Updated)
//line views/vretro/Detail.html:92
	qw422016.N().S(`</td>
        </tr>
      </tbody>
    </table>
  </div>
`)
//line views/vretro/Detail.html:99
	if len(p.FeedbacksByRetroID) > 0 {
//line views/vretro/Detail.html:99
		qw422016.N().S(`  <div class="card">
    <h3>`)
//line views/vretro/Detail.html:101
		components.StreamSVGRefIcon(qw422016, `star`, ps)
//line views/vretro/Detail.html:101
		qw422016.N().S(` Related feedbacks by [retro id]</h3>
    <div class="overflow clear">
      `)
//line views/vretro/Detail.html:103
		vfeedback.StreamTable(qw422016, p.FeedbacksByRetroID, nil, nil, p.Params, as, ps)
//line views/vretro/Detail.html:103
		qw422016.N().S(`
    </div>
  </div>
`)
//line views/vretro/Detail.html:106
	}
//line views/vretro/Detail.html:107
	if len(p.RetroHistoriesByRetroID) > 0 {
//line views/vretro/Detail.html:107
		qw422016.N().S(`  <div class="card">
    <h3>`)
//line views/vretro/Detail.html:109
		components.StreamSVGRefIcon(qw422016, `clock`, ps)
//line views/vretro/Detail.html:109
		qw422016.N().S(` Related histories by [retro id]</h3>
    <div class="overflow clear">
      `)
//line views/vretro/Detail.html:111
		vrhistory.StreamTable(qw422016, p.RetroHistoriesByRetroID, nil, p.Params, as, ps)
//line views/vretro/Detail.html:111
		qw422016.N().S(`
    </div>
  </div>
`)
//line views/vretro/Detail.html:114
	}
//line views/vretro/Detail.html:115
	if len(p.RetroMembersByRetroID) > 0 {
//line views/vretro/Detail.html:115
		qw422016.N().S(`  <div class="card">
    <h3>`)
//line views/vretro/Detail.html:117
		components.StreamSVGRefIcon(qw422016, `users`, ps)
//line views/vretro/Detail.html:117
		qw422016.N().S(` Related members by [retro id]</h3>
    <div class="overflow clear">
      `)
//line views/vretro/Detail.html:119
		vrmember.StreamTable(qw422016, p.RetroMembersByRetroID, nil, nil, p.Params, as, ps)
//line views/vretro/Detail.html:119
		qw422016.N().S(`
    </div>
  </div>
`)
//line views/vretro/Detail.html:122
	}
//line views/vretro/Detail.html:123
	if len(p.RetroPermissionsByRetroID) > 0 {
//line views/vretro/Detail.html:123
		qw422016.N().S(`  <div class="card">
    <h3>`)
//line views/vretro/Detail.html:125
		components.StreamSVGRefIcon(qw422016, `users`, ps)
//line views/vretro/Detail.html:125
		qw422016.N().S(` Related permissions by [retro id]</h3>
    <div class="overflow clear">
      `)
//line views/vretro/Detail.html:127
		vrpermission.StreamTable(qw422016, p.RetroPermissionsByRetroID, nil, p.Params, as, ps)
//line views/vretro/Detail.html:127
		qw422016.N().S(`
    </div>
  </div>
`)
//line views/vretro/Detail.html:130
	}
//line views/vretro/Detail.html:130
	qw422016.N().S(`  `)
//line views/vretro/Detail.html:131
	components.StreamJSONModal(qw422016, "retro", "Retro JSON", p.Model, 1)
//line views/vretro/Detail.html:131
	qw422016.N().S(`
`)
//line views/vretro/Detail.html:132
}

//line views/vretro/Detail.html:132
func (p *Detail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vretro/Detail.html:132
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vretro/Detail.html:132
	p.StreamBody(qw422016, as, ps)
//line views/vretro/Detail.html:132
	qt422016.ReleaseWriter(qw422016)
//line views/vretro/Detail.html:132
}

//line views/vretro/Detail.html:132
func (p *Detail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vretro/Detail.html:132
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vretro/Detail.html:132
	p.WriteBody(qb422016, as, ps)
//line views/vretro/Detail.html:132
	qs422016 := string(qb422016.B)
//line views/vretro/Detail.html:132
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vretro/Detail.html:132
	return qs422016
//line views/vretro/Detail.html:132
}
