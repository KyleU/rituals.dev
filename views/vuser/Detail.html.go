// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vuser/Detail.html:1
package vuser

//line views/vuser/Detail.html:1
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/comment"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/email"
	"github.com/kyleu/rituals/app/estimate/emember"
	"github.com/kyleu/rituals/app/estimate/story"
	"github.com/kyleu/rituals/app/estimate/story/vote"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/retro/feedback"
	"github.com/kyleu/rituals/app/retro/rmember"
	"github.com/kyleu/rituals/app/sprint/smember"
	"github.com/kyleu/rituals/app/standup/report"
	"github.com/kyleu/rituals/app/standup/umember"
	"github.com/kyleu/rituals/app/team/tmember"
	"github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/components/view"
	"github.com/kyleu/rituals/views/layout"
	"github.com/kyleu/rituals/views/vaction"
	"github.com/kyleu/rituals/views/vcomment"
	"github.com/kyleu/rituals/views/vemail"
	"github.com/kyleu/rituals/views/vestimate/vemember"
	"github.com/kyleu/rituals/views/vestimate/vstory"
	"github.com/kyleu/rituals/views/vestimate/vstory/vvote"
	"github.com/kyleu/rituals/views/vretro/vfeedback"
	"github.com/kyleu/rituals/views/vretro/vrmember"
	"github.com/kyleu/rituals/views/vsprint/vsmember"
	"github.com/kyleu/rituals/views/vstandup/vreport"
	"github.com/kyleu/rituals/views/vstandup/vumember"
	"github.com/kyleu/rituals/views/vteam/vtmember"
)

//line views/vuser/Detail.html:36
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vuser/Detail.html:36
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vuser/Detail.html:36
type Detail struct {
	layout.Basic
	Model                      *user.User
	Params                     filter.ParamSet
	RelActionsByUserID         action.Actions
	RelCommentsByUserID        comment.Comments
	RelEmailsByUserID          email.Emails
	RelEstimateMembersByUserID emember.EstimateMembers
	RelFeedbacksByUserID       feedback.Feedbacks
	RelReportsByUserID         report.Reports
	RelRetroMembersByUserID    rmember.RetroMembers
	RelSprintMembersByUserID   smember.SprintMembers
	RelStandupMembersByUserID  umember.StandupMembers
	RelStoriesByUserID         story.Stories
	RelTeamMembersByUserID     tmember.TeamMembers
	RelVotesByUserID           vote.Votes
	Paths                      []string
}

//line views/vuser/Detail.html:55
func (p *Detail) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vuser/Detail.html:55
	qw422016.N().S(`
  <div class="card">
    <div class="right">
      <a href="#modal-user"><button type="button">`)
//line views/vuser/Detail.html:58
	components.StreamSVGButton(qw422016, "file", ps)
//line views/vuser/Detail.html:58
	qw422016.N().S(` JSON</button></a>
      <a href="`)
//line views/vuser/Detail.html:59
	qw422016.E().S(p.Model.WebPath(p.Paths...))
//line views/vuser/Detail.html:59
	qw422016.N().S(`/edit"><button>`)
//line views/vuser/Detail.html:59
	components.StreamSVGButton(qw422016, "edit", ps)
//line views/vuser/Detail.html:59
	qw422016.N().S(` Edit</button></a>
    </div>
    <h3>`)
//line views/vuser/Detail.html:61
	components.StreamSVGIcon(qw422016, `profile`, ps)
//line views/vuser/Detail.html:61
	qw422016.N().S(` `)
//line views/vuser/Detail.html:61
	qw422016.E().S(p.Model.TitleString())
//line views/vuser/Detail.html:61
	qw422016.N().S(`</h3>
    <div><a href="`)
//line views/vuser/Detail.html:62
	qw422016.E().S(user.Route(p.Paths...))
//line views/vuser/Detail.html:62
	qw422016.N().S(`"><em>User</em></a></div>
    <div class="mt overflow full-width">
      <table>
        <tbody>
          <tr>
            <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">ID</th>
            <td>`)
//line views/vuser/Detail.html:68
	view.StreamUUID(qw422016, &p.Model.ID)
//line views/vuser/Detail.html:68
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="String text">Name</th>
            <td><strong>`)
//line views/vuser/Detail.html:72
	view.StreamString(qw422016, p.Model.Name)
//line views/vuser/Detail.html:72
	qw422016.N().S(`</strong></td>
          </tr>
          <tr>
            <th class="shrink" title="URL in string form">Picture</th>
            <td><a href="`)
//line views/vuser/Detail.html:76
	qw422016.E().S(p.Model.Picture)
//line views/vuser/Detail.html:76
	qw422016.N().S(`" target="_blank" rel="noopener noreferrer">`)
//line views/vuser/Detail.html:76
	qw422016.E().S(p.Model.Picture)
//line views/vuser/Detail.html:76
	qw422016.N().S(`</a></td>
          </tr>
          <tr>
            <th class="shrink" title="Date and time, in almost any format">Created</th>
            <td>`)
//line views/vuser/Detail.html:80
	view.StreamTimestamp(qw422016, &p.Model.Created)
//line views/vuser/Detail.html:80
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="Date and time, in almost any format (optional)">Updated</th>
            <td>`)
//line views/vuser/Detail.html:84
	view.StreamTimestamp(qw422016, p.Model.Updated)
//line views/vuser/Detail.html:84
	qw422016.N().S(`</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
`)
//line views/vuser/Detail.html:92
	relationHelper := user.Users{p.Model}

//line views/vuser/Detail.html:92
	qw422016.N().S(`  <div class="card">
    <h3 class="mb">Relations</h3>
    <ul class="accordion">
      <li>
        <input id="accordion-ActionsByUserID" type="checkbox" hidden="hidden"`)
//line views/vuser/Detail.html:97
	if p.Params.Specifies(`action`) {
//line views/vuser/Detail.html:97
		qw422016.N().S(` checked="checked"`)
//line views/vuser/Detail.html:97
	}
//line views/vuser/Detail.html:97
	qw422016.N().S(` />
        <label for="accordion-ActionsByUserID">
          `)
//line views/vuser/Detail.html:99
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vuser/Detail.html:99
	qw422016.N().S(`
          `)
//line views/vuser/Detail.html:100
	components.StreamSVGInline(qw422016, `action`, 16, ps)
//line views/vuser/Detail.html:100
	qw422016.N().S(`
          `)
//line views/vuser/Detail.html:101
	qw422016.E().S(util.StringPlural(len(p.RelActionsByUserID), "Action"))
//line views/vuser/Detail.html:101
	qw422016.N().S(` by [User ID]
        </label>
        <div class="bd"><div><div>
`)
//line views/vuser/Detail.html:104
	if len(p.RelActionsByUserID) == 0 {
//line views/vuser/Detail.html:104
		qw422016.N().S(`          <em>no related Actions</em>
`)
//line views/vuser/Detail.html:106
	} else {
//line views/vuser/Detail.html:106
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vuser/Detail.html:108
		vaction.StreamTable(qw422016, p.RelActionsByUserID, relationHelper, p.Params, as, ps)
//line views/vuser/Detail.html:108
		qw422016.N().S(`
          </div>
`)
//line views/vuser/Detail.html:110
	}
//line views/vuser/Detail.html:110
	qw422016.N().S(`        </div></div></div>
      </li>
      <li>
        <input id="accordion-CommentsByUserID" type="checkbox" hidden="hidden"`)
//line views/vuser/Detail.html:114
	if p.Params.Specifies(`comment`) {
//line views/vuser/Detail.html:114
		qw422016.N().S(` checked="checked"`)
//line views/vuser/Detail.html:114
	}
//line views/vuser/Detail.html:114
	qw422016.N().S(` />
        <label for="accordion-CommentsByUserID">
          `)
//line views/vuser/Detail.html:116
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vuser/Detail.html:116
	qw422016.N().S(`
          `)
//line views/vuser/Detail.html:117
	components.StreamSVGInline(qw422016, `comments`, 16, ps)
//line views/vuser/Detail.html:117
	qw422016.N().S(`
          `)
//line views/vuser/Detail.html:118
	qw422016.E().S(util.StringPlural(len(p.RelCommentsByUserID), "Comment"))
//line views/vuser/Detail.html:118
	qw422016.N().S(` by [User ID]
        </label>
        <div class="bd"><div><div>
`)
//line views/vuser/Detail.html:121
	if len(p.RelCommentsByUserID) == 0 {
//line views/vuser/Detail.html:121
		qw422016.N().S(`          <em>no related Comments</em>
`)
//line views/vuser/Detail.html:123
	} else {
//line views/vuser/Detail.html:123
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vuser/Detail.html:125
		vcomment.StreamTable(qw422016, p.RelCommentsByUserID, relationHelper, p.Params, as, ps)
//line views/vuser/Detail.html:125
		qw422016.N().S(`
          </div>
`)
//line views/vuser/Detail.html:127
	}
//line views/vuser/Detail.html:127
	qw422016.N().S(`        </div></div></div>
      </li>
      <li>
        <input id="accordion-EmailsByUserID" type="checkbox" hidden="hidden"`)
//line views/vuser/Detail.html:131
	if p.Params.Specifies(`email`) {
//line views/vuser/Detail.html:131
		qw422016.N().S(` checked="checked"`)
//line views/vuser/Detail.html:131
	}
//line views/vuser/Detail.html:131
	qw422016.N().S(` />
        <label for="accordion-EmailsByUserID">
          `)
//line views/vuser/Detail.html:133
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vuser/Detail.html:133
	qw422016.N().S(`
          `)
//line views/vuser/Detail.html:134
	components.StreamSVGInline(qw422016, `email`, 16, ps)
//line views/vuser/Detail.html:134
	qw422016.N().S(`
          `)
//line views/vuser/Detail.html:135
	qw422016.E().S(util.StringPlural(len(p.RelEmailsByUserID), "Email"))
//line views/vuser/Detail.html:135
	qw422016.N().S(` by [User ID]
        </label>
        <div class="bd"><div><div>
`)
//line views/vuser/Detail.html:138
	if len(p.RelEmailsByUserID) == 0 {
//line views/vuser/Detail.html:138
		qw422016.N().S(`          <em>no related Emails</em>
`)
//line views/vuser/Detail.html:140
	} else {
//line views/vuser/Detail.html:140
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vuser/Detail.html:142
		vemail.StreamTable(qw422016, p.RelEmailsByUserID, relationHelper, p.Params, as, ps)
//line views/vuser/Detail.html:142
		qw422016.N().S(`
          </div>
`)
//line views/vuser/Detail.html:144
	}
//line views/vuser/Detail.html:144
	qw422016.N().S(`        </div></div></div>
      </li>
      <li>
        <input id="accordion-EstimateMembersByUserID" type="checkbox" hidden="hidden"`)
//line views/vuser/Detail.html:148
	if p.Params.Specifies(`emember`) {
//line views/vuser/Detail.html:148
		qw422016.N().S(` checked="checked"`)
//line views/vuser/Detail.html:148
	}
//line views/vuser/Detail.html:148
	qw422016.N().S(` />
        <label for="accordion-EstimateMembersByUserID">
          `)
//line views/vuser/Detail.html:150
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vuser/Detail.html:150
	qw422016.N().S(`
          `)
//line views/vuser/Detail.html:151
	components.StreamSVGInline(qw422016, `users`, 16, ps)
//line views/vuser/Detail.html:151
	qw422016.N().S(`
          `)
//line views/vuser/Detail.html:152
	qw422016.E().S(util.StringPlural(len(p.RelEstimateMembersByUserID), "Member"))
//line views/vuser/Detail.html:152
	qw422016.N().S(` by [User ID]
        </label>
        <div class="bd"><div><div>
`)
//line views/vuser/Detail.html:155
	if len(p.RelEstimateMembersByUserID) == 0 {
//line views/vuser/Detail.html:155
		qw422016.N().S(`          <em>no related Members</em>
`)
//line views/vuser/Detail.html:157
	} else {
//line views/vuser/Detail.html:157
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vuser/Detail.html:159
		vemember.StreamTable(qw422016, p.RelEstimateMembersByUserID, nil, relationHelper, p.Params, as, ps)
//line views/vuser/Detail.html:159
		qw422016.N().S(`
          </div>
`)
//line views/vuser/Detail.html:161
	}
//line views/vuser/Detail.html:161
	qw422016.N().S(`        </div></div></div>
      </li>
      <li>
        <input id="accordion-FeedbacksByUserID" type="checkbox" hidden="hidden"`)
//line views/vuser/Detail.html:165
	if p.Params.Specifies(`feedback`) {
//line views/vuser/Detail.html:165
		qw422016.N().S(` checked="checked"`)
//line views/vuser/Detail.html:165
	}
//line views/vuser/Detail.html:165
	qw422016.N().S(` />
        <label for="accordion-FeedbacksByUserID">
          `)
//line views/vuser/Detail.html:167
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vuser/Detail.html:167
	qw422016.N().S(`
          `)
//line views/vuser/Detail.html:168
	components.StreamSVGInline(qw422016, `comment`, 16, ps)
//line views/vuser/Detail.html:168
	qw422016.N().S(`
          `)
//line views/vuser/Detail.html:169
	qw422016.E().S(util.StringPlural(len(p.RelFeedbacksByUserID), "Feedback"))
//line views/vuser/Detail.html:169
	qw422016.N().S(` by [User ID]
        </label>
        <div class="bd"><div><div>
`)
//line views/vuser/Detail.html:172
	if len(p.RelFeedbacksByUserID) == 0 {
//line views/vuser/Detail.html:172
		qw422016.N().S(`          <em>no related Feedbacks</em>
`)
//line views/vuser/Detail.html:174
	} else {
//line views/vuser/Detail.html:174
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vuser/Detail.html:176
		vfeedback.StreamTable(qw422016, p.RelFeedbacksByUserID, nil, relationHelper, p.Params, as, ps)
//line views/vuser/Detail.html:176
		qw422016.N().S(`
          </div>
`)
//line views/vuser/Detail.html:178
	}
//line views/vuser/Detail.html:178
	qw422016.N().S(`        </div></div></div>
      </li>
      <li>
        <input id="accordion-ReportsByUserID" type="checkbox" hidden="hidden"`)
//line views/vuser/Detail.html:182
	if p.Params.Specifies(`report`) {
//line views/vuser/Detail.html:182
		qw422016.N().S(` checked="checked"`)
//line views/vuser/Detail.html:182
	}
//line views/vuser/Detail.html:182
	qw422016.N().S(` />
        <label for="accordion-ReportsByUserID">
          `)
//line views/vuser/Detail.html:184
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vuser/Detail.html:184
	qw422016.N().S(`
          `)
//line views/vuser/Detail.html:185
	components.StreamSVGInline(qw422016, `file-alt`, 16, ps)
//line views/vuser/Detail.html:185
	qw422016.N().S(`
          `)
//line views/vuser/Detail.html:186
	qw422016.E().S(util.StringPlural(len(p.RelReportsByUserID), "Report"))
//line views/vuser/Detail.html:186
	qw422016.N().S(` by [User ID]
        </label>
        <div class="bd"><div><div>
`)
//line views/vuser/Detail.html:189
	if len(p.RelReportsByUserID) == 0 {
//line views/vuser/Detail.html:189
		qw422016.N().S(`          <em>no related Reports</em>
`)
//line views/vuser/Detail.html:191
	} else {
//line views/vuser/Detail.html:191
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vuser/Detail.html:193
		vreport.StreamTable(qw422016, p.RelReportsByUserID, nil, relationHelper, p.Params, as, ps)
//line views/vuser/Detail.html:193
		qw422016.N().S(`
          </div>
`)
//line views/vuser/Detail.html:195
	}
//line views/vuser/Detail.html:195
	qw422016.N().S(`        </div></div></div>
      </li>
      <li>
        <input id="accordion-RetroMembersByUserID" type="checkbox" hidden="hidden"`)
//line views/vuser/Detail.html:199
	if p.Params.Specifies(`rmember`) {
//line views/vuser/Detail.html:199
		qw422016.N().S(` checked="checked"`)
//line views/vuser/Detail.html:199
	}
//line views/vuser/Detail.html:199
	qw422016.N().S(` />
        <label for="accordion-RetroMembersByUserID">
          `)
//line views/vuser/Detail.html:201
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vuser/Detail.html:201
	qw422016.N().S(`
          `)
//line views/vuser/Detail.html:202
	components.StreamSVGInline(qw422016, `users`, 16, ps)
//line views/vuser/Detail.html:202
	qw422016.N().S(`
          `)
//line views/vuser/Detail.html:203
	qw422016.E().S(util.StringPlural(len(p.RelRetroMembersByUserID), "Member"))
//line views/vuser/Detail.html:203
	qw422016.N().S(` by [User ID]
        </label>
        <div class="bd"><div><div>
`)
//line views/vuser/Detail.html:206
	if len(p.RelRetroMembersByUserID) == 0 {
//line views/vuser/Detail.html:206
		qw422016.N().S(`          <em>no related Members</em>
`)
//line views/vuser/Detail.html:208
	} else {
//line views/vuser/Detail.html:208
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vuser/Detail.html:210
		vrmember.StreamTable(qw422016, p.RelRetroMembersByUserID, nil, relationHelper, p.Params, as, ps)
//line views/vuser/Detail.html:210
		qw422016.N().S(`
          </div>
`)
//line views/vuser/Detail.html:212
	}
//line views/vuser/Detail.html:212
	qw422016.N().S(`        </div></div></div>
      </li>
      <li>
        <input id="accordion-SprintMembersByUserID" type="checkbox" hidden="hidden"`)
//line views/vuser/Detail.html:216
	if p.Params.Specifies(`smember`) {
//line views/vuser/Detail.html:216
		qw422016.N().S(` checked="checked"`)
//line views/vuser/Detail.html:216
	}
//line views/vuser/Detail.html:216
	qw422016.N().S(` />
        <label for="accordion-SprintMembersByUserID">
          `)
//line views/vuser/Detail.html:218
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vuser/Detail.html:218
	qw422016.N().S(`
          `)
//line views/vuser/Detail.html:219
	components.StreamSVGInline(qw422016, `users`, 16, ps)
//line views/vuser/Detail.html:219
	qw422016.N().S(`
          `)
//line views/vuser/Detail.html:220
	qw422016.E().S(util.StringPlural(len(p.RelSprintMembersByUserID), "Member"))
//line views/vuser/Detail.html:220
	qw422016.N().S(` by [User ID]
        </label>
        <div class="bd"><div><div>
`)
//line views/vuser/Detail.html:223
	if len(p.RelSprintMembersByUserID) == 0 {
//line views/vuser/Detail.html:223
		qw422016.N().S(`          <em>no related Members</em>
`)
//line views/vuser/Detail.html:225
	} else {
//line views/vuser/Detail.html:225
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vuser/Detail.html:227
		vsmember.StreamTable(qw422016, p.RelSprintMembersByUserID, nil, relationHelper, p.Params, as, ps)
//line views/vuser/Detail.html:227
		qw422016.N().S(`
          </div>
`)
//line views/vuser/Detail.html:229
	}
//line views/vuser/Detail.html:229
	qw422016.N().S(`        </div></div></div>
      </li>
      <li>
        <input id="accordion-StandupMembersByUserID" type="checkbox" hidden="hidden"`)
//line views/vuser/Detail.html:233
	if p.Params.Specifies(`umember`) {
//line views/vuser/Detail.html:233
		qw422016.N().S(` checked="checked"`)
//line views/vuser/Detail.html:233
	}
//line views/vuser/Detail.html:233
	qw422016.N().S(` />
        <label for="accordion-StandupMembersByUserID">
          `)
//line views/vuser/Detail.html:235
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vuser/Detail.html:235
	qw422016.N().S(`
          `)
//line views/vuser/Detail.html:236
	components.StreamSVGInline(qw422016, `users`, 16, ps)
//line views/vuser/Detail.html:236
	qw422016.N().S(`
          `)
//line views/vuser/Detail.html:237
	qw422016.E().S(util.StringPlural(len(p.RelStandupMembersByUserID), "Member"))
//line views/vuser/Detail.html:237
	qw422016.N().S(` by [User ID]
        </label>
        <div class="bd"><div><div>
`)
//line views/vuser/Detail.html:240
	if len(p.RelStandupMembersByUserID) == 0 {
//line views/vuser/Detail.html:240
		qw422016.N().S(`          <em>no related Members</em>
`)
//line views/vuser/Detail.html:242
	} else {
//line views/vuser/Detail.html:242
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vuser/Detail.html:244
		vumember.StreamTable(qw422016, p.RelStandupMembersByUserID, nil, relationHelper, p.Params, as, ps)
//line views/vuser/Detail.html:244
		qw422016.N().S(`
          </div>
`)
//line views/vuser/Detail.html:246
	}
//line views/vuser/Detail.html:246
	qw422016.N().S(`        </div></div></div>
      </li>
      <li>
        <input id="accordion-StoriesByUserID" type="checkbox" hidden="hidden"`)
//line views/vuser/Detail.html:250
	if p.Params.Specifies(`story`) {
//line views/vuser/Detail.html:250
		qw422016.N().S(` checked="checked"`)
//line views/vuser/Detail.html:250
	}
//line views/vuser/Detail.html:250
	qw422016.N().S(` />
        <label for="accordion-StoriesByUserID">
          `)
//line views/vuser/Detail.html:252
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vuser/Detail.html:252
	qw422016.N().S(`
          `)
//line views/vuser/Detail.html:253
	components.StreamSVGInline(qw422016, `story`, 16, ps)
//line views/vuser/Detail.html:253
	qw422016.N().S(`
          `)
//line views/vuser/Detail.html:254
	qw422016.E().S(util.StringPlural(len(p.RelStoriesByUserID), "Story"))
//line views/vuser/Detail.html:254
	qw422016.N().S(` by [User ID]
        </label>
        <div class="bd"><div><div>
`)
//line views/vuser/Detail.html:257
	if len(p.RelStoriesByUserID) == 0 {
//line views/vuser/Detail.html:257
		qw422016.N().S(`          <em>no related Stories</em>
`)
//line views/vuser/Detail.html:259
	} else {
//line views/vuser/Detail.html:259
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vuser/Detail.html:261
		vstory.StreamTable(qw422016, p.RelStoriesByUserID, nil, relationHelper, p.Params, as, ps)
//line views/vuser/Detail.html:261
		qw422016.N().S(`
          </div>
`)
//line views/vuser/Detail.html:263
	}
//line views/vuser/Detail.html:263
	qw422016.N().S(`        </div></div></div>
      </li>
      <li>
        <input id="accordion-TeamMembersByUserID" type="checkbox" hidden="hidden"`)
//line views/vuser/Detail.html:267
	if p.Params.Specifies(`tmember`) {
//line views/vuser/Detail.html:267
		qw422016.N().S(` checked="checked"`)
//line views/vuser/Detail.html:267
	}
//line views/vuser/Detail.html:267
	qw422016.N().S(` />
        <label for="accordion-TeamMembersByUserID">
          `)
//line views/vuser/Detail.html:269
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vuser/Detail.html:269
	qw422016.N().S(`
          `)
//line views/vuser/Detail.html:270
	components.StreamSVGInline(qw422016, `users`, 16, ps)
//line views/vuser/Detail.html:270
	qw422016.N().S(`
          `)
//line views/vuser/Detail.html:271
	qw422016.E().S(util.StringPlural(len(p.RelTeamMembersByUserID), "Member"))
//line views/vuser/Detail.html:271
	qw422016.N().S(` by [User ID]
        </label>
        <div class="bd"><div><div>
`)
//line views/vuser/Detail.html:274
	if len(p.RelTeamMembersByUserID) == 0 {
//line views/vuser/Detail.html:274
		qw422016.N().S(`          <em>no related Members</em>
`)
//line views/vuser/Detail.html:276
	} else {
//line views/vuser/Detail.html:276
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vuser/Detail.html:278
		vtmember.StreamTable(qw422016, p.RelTeamMembersByUserID, nil, relationHelper, p.Params, as, ps)
//line views/vuser/Detail.html:278
		qw422016.N().S(`
          </div>
`)
//line views/vuser/Detail.html:280
	}
//line views/vuser/Detail.html:280
	qw422016.N().S(`        </div></div></div>
      </li>
      <li>
        <input id="accordion-VotesByUserID" type="checkbox" hidden="hidden"`)
//line views/vuser/Detail.html:284
	if p.Params.Specifies(`vote`) {
//line views/vuser/Detail.html:284
		qw422016.N().S(` checked="checked"`)
//line views/vuser/Detail.html:284
	}
//line views/vuser/Detail.html:284
	qw422016.N().S(` />
        <label for="accordion-VotesByUserID">
          `)
//line views/vuser/Detail.html:286
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vuser/Detail.html:286
	qw422016.N().S(`
          `)
//line views/vuser/Detail.html:287
	components.StreamSVGInline(qw422016, `vote-yea`, 16, ps)
//line views/vuser/Detail.html:287
	qw422016.N().S(`
          `)
//line views/vuser/Detail.html:288
	qw422016.E().S(util.StringPlural(len(p.RelVotesByUserID), "Vote"))
//line views/vuser/Detail.html:288
	qw422016.N().S(` by [User ID]
        </label>
        <div class="bd"><div><div>
`)
//line views/vuser/Detail.html:291
	if len(p.RelVotesByUserID) == 0 {
//line views/vuser/Detail.html:291
		qw422016.N().S(`          <em>no related Votes</em>
`)
//line views/vuser/Detail.html:293
	} else {
//line views/vuser/Detail.html:293
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vuser/Detail.html:295
		vvote.StreamTable(qw422016, p.RelVotesByUserID, nil, relationHelper, p.Params, as, ps)
//line views/vuser/Detail.html:295
		qw422016.N().S(`
          </div>
`)
//line views/vuser/Detail.html:297
	}
//line views/vuser/Detail.html:297
	qw422016.N().S(`        </div></div></div>
      </li>
    </ul>
  </div>
  `)
//line views/vuser/Detail.html:302
	components.StreamJSONModal(qw422016, "user", "User JSON", p.Model, 1)
//line views/vuser/Detail.html:302
	qw422016.N().S(`
`)
//line views/vuser/Detail.html:303
}

//line views/vuser/Detail.html:303
func (p *Detail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vuser/Detail.html:303
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vuser/Detail.html:303
	p.StreamBody(qw422016, as, ps)
//line views/vuser/Detail.html:303
	qt422016.ReleaseWriter(qw422016)
//line views/vuser/Detail.html:303
}

//line views/vuser/Detail.html:303
func (p *Detail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vuser/Detail.html:303
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vuser/Detail.html:303
	p.WriteBody(qb422016, as, ps)
//line views/vuser/Detail.html:303
	qs422016 := string(qb422016.B)
//line views/vuser/Detail.html:303
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vuser/Detail.html:303
	return qs422016
//line views/vuser/Detail.html:303
}
