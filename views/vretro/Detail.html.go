// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vretro/Detail.html:1
package vretro

//line views/vretro/Detail.html:1
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/retro"
	"github.com/kyleu/rituals/app/retro/feedback"
	"github.com/kyleu/rituals/app/retro/rhistory"
	"github.com/kyleu/rituals/app/retro/rmember"
	"github.com/kyleu/rituals/app/retro/rpermission"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/components/view"
	"github.com/kyleu/rituals/views/layout"
	"github.com/kyleu/rituals/views/vretro/vfeedback"
	"github.com/kyleu/rituals/views/vretro/vrhistory"
	"github.com/kyleu/rituals/views/vretro/vrmember"
	"github.com/kyleu/rituals/views/vretro/vrpermission"
)

//line views/vretro/Detail.html:23
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vretro/Detail.html:23
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vretro/Detail.html:23
type Detail struct {
	layout.Basic
	Model                        *retro.Retro
	TeamByTeamID                 *team.Team
	SprintBySprintID             *sprint.Sprint
	Params                       filter.ParamSet
	RelFeedbacksByRetroID        feedback.Feedbacks
	RelRetroHistoriesByRetroID   rhistory.RetroHistories
	RelRetroMembersByRetroID     rmember.RetroMembers
	RelRetroPermissionsByRetroID rpermission.RetroPermissions
	Paths                        []string
}

//line views/vretro/Detail.html:36
func (p *Detail) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vretro/Detail.html:36
	qw422016.N().S(`
  <div class="card">
    <div class="right">
      <a href="#modal-retro"><button type="button">`)
//line views/vretro/Detail.html:39
	components.StreamSVGButton(qw422016, "file", ps)
//line views/vretro/Detail.html:39
	qw422016.N().S(` JSON</button></a>
      <a href="`)
//line views/vretro/Detail.html:40
	qw422016.E().S(p.Model.WebPath(p.Paths...))
//line views/vretro/Detail.html:40
	qw422016.N().S(`/edit"><button>`)
//line views/vretro/Detail.html:40
	components.StreamSVGButton(qw422016, "edit", ps)
//line views/vretro/Detail.html:40
	qw422016.N().S(` Edit</button></a>
    </div>
    <h3>`)
//line views/vretro/Detail.html:42
	components.StreamSVGIcon(qw422016, `retro`, ps)
//line views/vretro/Detail.html:42
	qw422016.N().S(` `)
//line views/vretro/Detail.html:42
	qw422016.E().S(p.Model.TitleString())
//line views/vretro/Detail.html:42
	qw422016.N().S(`</h3>
    <div><a href="`)
//line views/vretro/Detail.html:43
	qw422016.E().S(retro.Route(p.Paths...))
//line views/vretro/Detail.html:43
	qw422016.N().S(`"><em>Retro</em></a></div>
    <div class="mt overflow full-width">
      <table>
        <tbody>
          <tr>
            <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">ID</th>
            <td>`)
//line views/vretro/Detail.html:49
	view.StreamUUID(qw422016, &p.Model.ID)
//line views/vretro/Detail.html:49
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="String text">Slug</th>
            <td>`)
//line views/vretro/Detail.html:53
	view.StreamString(qw422016, p.Model.Slug)
//line views/vretro/Detail.html:53
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="String text">Title</th>
            <td><strong>`)
//line views/vretro/Detail.html:57
	view.StreamString(qw422016, p.Model.Title)
//line views/vretro/Detail.html:57
	qw422016.N().S(`</strong></td>
          </tr>
          <tr>
            <th class="shrink" title="String text">Icon</th>
            <td>`)
//line views/vretro/Detail.html:61
	view.StreamString(qw422016, p.Model.Icon)
//line views/vretro/Detail.html:61
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="`)
//line views/vretro/Detail.html:64
	qw422016.E().S(enum.AllSessionStatuses.Help())
//line views/vretro/Detail.html:64
	qw422016.N().S(`">Status</th>
            <td>`)
//line views/vretro/Detail.html:65
	qw422016.E().S(p.Model.Status.String())
//line views/vretro/Detail.html:65
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000) (optional)">Team ID</th>
            <td class="nowrap">
              `)
//line views/vretro/Detail.html:70
	view.StreamUUID(qw422016, p.Model.TeamID)
//line views/vretro/Detail.html:70
	if p.TeamByTeamID != nil {
//line views/vretro/Detail.html:70
		qw422016.N().S(` (`)
//line views/vretro/Detail.html:70
		qw422016.E().S(p.TeamByTeamID.TitleString())
//line views/vretro/Detail.html:70
		qw422016.N().S(`)`)
//line views/vretro/Detail.html:70
	}
//line views/vretro/Detail.html:70
	qw422016.N().S(`
              `)
//line views/vretro/Detail.html:71
	if p.Model.TeamID != nil {
//line views/vretro/Detail.html:71
		qw422016.N().S(`<a title="Team" href="`)
//line views/vretro/Detail.html:71
		qw422016.E().S(p.Model.WebPath(p.Paths...))
//line views/vretro/Detail.html:71
		qw422016.N().S(`">`)
//line views/vretro/Detail.html:71
		components.StreamSVGLink(qw422016, `team`, ps)
//line views/vretro/Detail.html:71
		qw422016.N().S(`</a>`)
//line views/vretro/Detail.html:71
	}
//line views/vretro/Detail.html:71
	qw422016.N().S(`
            </td>
          </tr>
          <tr>
            <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000) (optional)">Sprint ID</th>
            <td class="nowrap">
              `)
//line views/vretro/Detail.html:77
	view.StreamUUID(qw422016, p.Model.SprintID)
//line views/vretro/Detail.html:77
	if p.SprintBySprintID != nil {
//line views/vretro/Detail.html:77
		qw422016.N().S(` (`)
//line views/vretro/Detail.html:77
		qw422016.E().S(p.SprintBySprintID.TitleString())
//line views/vretro/Detail.html:77
		qw422016.N().S(`)`)
//line views/vretro/Detail.html:77
	}
//line views/vretro/Detail.html:77
	qw422016.N().S(`
              `)
//line views/vretro/Detail.html:78
	if p.Model.SprintID != nil {
//line views/vretro/Detail.html:78
		qw422016.N().S(`<a title="Sprint" href="`)
//line views/vretro/Detail.html:78
		qw422016.E().S(p.Model.WebPath(p.Paths...))
//line views/vretro/Detail.html:78
		qw422016.N().S(`">`)
//line views/vretro/Detail.html:78
		components.StreamSVGLink(qw422016, `sprint`, ps)
//line views/vretro/Detail.html:78
		qw422016.N().S(`</a>`)
//line views/vretro/Detail.html:78
	}
//line views/vretro/Detail.html:78
	qw422016.N().S(`
            </td>
          </tr>
          <tr>
            <th class="shrink" title="Comma-separated list of values">Categories</th>
            <td>`)
//line views/vretro/Detail.html:83
	view.StreamStringArray(qw422016, p.Model.Categories)
//line views/vretro/Detail.html:83
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="Date and time, in almost any format">Created</th>
            <td>`)
//line views/vretro/Detail.html:87
	view.StreamTimestamp(qw422016, &p.Model.Created)
//line views/vretro/Detail.html:87
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="Date and time, in almost any format (optional)">Updated</th>
            <td>`)
//line views/vretro/Detail.html:91
	view.StreamTimestamp(qw422016, p.Model.Updated)
//line views/vretro/Detail.html:91
	qw422016.N().S(`</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
`)
//line views/vretro/Detail.html:99
	relationHelper := retro.Retros{p.Model}

//line views/vretro/Detail.html:99
	qw422016.N().S(`  <div class="card">
    <h3 class="mb">Relations</h3>
    <ul class="accordion">
      <li>
        <input id="accordion-FeedbacksByRetroID" type="checkbox" hidden="hidden"`)
//line views/vretro/Detail.html:104
	if p.Params.Specifies(`feedback`) {
//line views/vretro/Detail.html:104
		qw422016.N().S(` checked="checked"`)
//line views/vretro/Detail.html:104
	}
//line views/vretro/Detail.html:104
	qw422016.N().S(` />
        <label for="accordion-FeedbacksByRetroID">
          `)
//line views/vretro/Detail.html:106
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vretro/Detail.html:106
	qw422016.N().S(`
          `)
//line views/vretro/Detail.html:107
	components.StreamSVGInline(qw422016, `comment`, 16, ps)
//line views/vretro/Detail.html:107
	qw422016.N().S(`
          `)
//line views/vretro/Detail.html:108
	qw422016.E().S(util.StringPlural(len(p.RelFeedbacksByRetroID), "Feedback"))
//line views/vretro/Detail.html:108
	qw422016.N().S(` by [Retro ID]
        </label>
        <div class="bd"><div><div>
`)
//line views/vretro/Detail.html:111
	if len(p.RelFeedbacksByRetroID) == 0 {
//line views/vretro/Detail.html:111
		qw422016.N().S(`          <em>no related Feedbacks</em>
`)
//line views/vretro/Detail.html:113
	} else {
//line views/vretro/Detail.html:113
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vretro/Detail.html:115
		vfeedback.StreamTable(qw422016, p.RelFeedbacksByRetroID, relationHelper, nil, p.Params, as, ps)
//line views/vretro/Detail.html:115
		qw422016.N().S(`
          </div>
`)
//line views/vretro/Detail.html:117
	}
//line views/vretro/Detail.html:117
	qw422016.N().S(`        </div></div></div>
      </li>
      <li>
        <input id="accordion-RetroHistoriesByRetroID" type="checkbox" hidden="hidden"`)
//line views/vretro/Detail.html:121
	if p.Params.Specifies(`rhistory`) {
//line views/vretro/Detail.html:121
		qw422016.N().S(` checked="checked"`)
//line views/vretro/Detail.html:121
	}
//line views/vretro/Detail.html:121
	qw422016.N().S(` />
        <label for="accordion-RetroHistoriesByRetroID">
          `)
//line views/vretro/Detail.html:123
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vretro/Detail.html:123
	qw422016.N().S(`
          `)
//line views/vretro/Detail.html:124
	components.StreamSVGInline(qw422016, `history`, 16, ps)
//line views/vretro/Detail.html:124
	qw422016.N().S(`
          `)
//line views/vretro/Detail.html:125
	qw422016.E().S(util.StringPlural(len(p.RelRetroHistoriesByRetroID), "History"))
//line views/vretro/Detail.html:125
	qw422016.N().S(` by [Retro ID]
        </label>
        <div class="bd"><div><div>
`)
//line views/vretro/Detail.html:128
	if len(p.RelRetroHistoriesByRetroID) == 0 {
//line views/vretro/Detail.html:128
		qw422016.N().S(`          <em>no related Histories</em>
`)
//line views/vretro/Detail.html:130
	} else {
//line views/vretro/Detail.html:130
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vretro/Detail.html:132
		vrhistory.StreamTable(qw422016, p.RelRetroHistoriesByRetroID, relationHelper, p.Params, as, ps)
//line views/vretro/Detail.html:132
		qw422016.N().S(`
          </div>
`)
//line views/vretro/Detail.html:134
	}
//line views/vretro/Detail.html:134
	qw422016.N().S(`        </div></div></div>
      </li>
      <li>
        <input id="accordion-RetroMembersByRetroID" type="checkbox" hidden="hidden"`)
//line views/vretro/Detail.html:138
	if p.Params.Specifies(`rmember`) {
//line views/vretro/Detail.html:138
		qw422016.N().S(` checked="checked"`)
//line views/vretro/Detail.html:138
	}
//line views/vretro/Detail.html:138
	qw422016.N().S(` />
        <label for="accordion-RetroMembersByRetroID">
          `)
//line views/vretro/Detail.html:140
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vretro/Detail.html:140
	qw422016.N().S(`
          `)
//line views/vretro/Detail.html:141
	components.StreamSVGInline(qw422016, `users`, 16, ps)
//line views/vretro/Detail.html:141
	qw422016.N().S(`
          `)
//line views/vretro/Detail.html:142
	qw422016.E().S(util.StringPlural(len(p.RelRetroMembersByRetroID), "Member"))
//line views/vretro/Detail.html:142
	qw422016.N().S(` by [Retro ID]
        </label>
        <div class="bd"><div><div>
`)
//line views/vretro/Detail.html:145
	if len(p.RelRetroMembersByRetroID) == 0 {
//line views/vretro/Detail.html:145
		qw422016.N().S(`          <em>no related Members</em>
`)
//line views/vretro/Detail.html:147
	} else {
//line views/vretro/Detail.html:147
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vretro/Detail.html:149
		vrmember.StreamTable(qw422016, p.RelRetroMembersByRetroID, relationHelper, nil, p.Params, as, ps)
//line views/vretro/Detail.html:149
		qw422016.N().S(`
          </div>
`)
//line views/vretro/Detail.html:151
	}
//line views/vretro/Detail.html:151
	qw422016.N().S(`        </div></div></div>
      </li>
      <li>
        <input id="accordion-RetroPermissionsByRetroID" type="checkbox" hidden="hidden"`)
//line views/vretro/Detail.html:155
	if p.Params.Specifies(`rpermission`) {
//line views/vretro/Detail.html:155
		qw422016.N().S(` checked="checked"`)
//line views/vretro/Detail.html:155
	}
//line views/vretro/Detail.html:155
	qw422016.N().S(` />
        <label for="accordion-RetroPermissionsByRetroID">
          `)
//line views/vretro/Detail.html:157
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vretro/Detail.html:157
	qw422016.N().S(`
          `)
//line views/vretro/Detail.html:158
	components.StreamSVGInline(qw422016, `permission`, 16, ps)
//line views/vretro/Detail.html:158
	qw422016.N().S(`
          `)
//line views/vretro/Detail.html:159
	qw422016.E().S(util.StringPlural(len(p.RelRetroPermissionsByRetroID), "Permission"))
//line views/vretro/Detail.html:159
	qw422016.N().S(` by [Retro ID]
        </label>
        <div class="bd"><div><div>
`)
//line views/vretro/Detail.html:162
	if len(p.RelRetroPermissionsByRetroID) == 0 {
//line views/vretro/Detail.html:162
		qw422016.N().S(`          <em>no related Permissions</em>
`)
//line views/vretro/Detail.html:164
	} else {
//line views/vretro/Detail.html:164
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vretro/Detail.html:166
		vrpermission.StreamTable(qw422016, p.RelRetroPermissionsByRetroID, relationHelper, p.Params, as, ps)
//line views/vretro/Detail.html:166
		qw422016.N().S(`
          </div>
`)
//line views/vretro/Detail.html:168
	}
//line views/vretro/Detail.html:168
	qw422016.N().S(`        </div></div></div>
      </li>
    </ul>
  </div>
  `)
//line views/vretro/Detail.html:173
	components.StreamJSONModal(qw422016, "retro", "Retro JSON", p.Model, 1)
//line views/vretro/Detail.html:173
	qw422016.N().S(`
`)
//line views/vretro/Detail.html:174
}

//line views/vretro/Detail.html:174
func (p *Detail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vretro/Detail.html:174
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vretro/Detail.html:174
	p.StreamBody(qw422016, as, ps)
//line views/vretro/Detail.html:174
	qt422016.ReleaseWriter(qw422016)
//line views/vretro/Detail.html:174
}

//line views/vretro/Detail.html:174
func (p *Detail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vretro/Detail.html:174
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vretro/Detail.html:174
	p.WriteBody(qb422016, as, ps)
//line views/vretro/Detail.html:174
	qs422016 := string(qb422016.B)
//line views/vretro/Detail.html:174
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vretro/Detail.html:174
	return qs422016
//line views/vretro/Detail.html:174
}
