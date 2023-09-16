// Code generated by qtc from "StandupList.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vworkspace/vwstandup/StandupList.html:1
package vwstandup

//line views/vworkspace/vwstandup/StandupList.html:1
import (
	"fmt"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/comment"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/standup"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/layout"
	"github.com/kyleu/rituals/views/vworkspace/vwutil"
)

//line views/vworkspace/vwstandup/StandupList.html:19
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vworkspace/vwstandup/StandupList.html:19
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vworkspace/vwstandup/StandupList.html:19
type StandupList struct {
	layout.Basic
	Sprints  sprint.Sprints
	Standups standup.Standups
	Teams    team.Teams
}

//line views/vworkspace/vwstandup/StandupList.html:26
func (p *StandupList) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vworkspace/vwstandup/StandupList.html:26
	qw422016.N().S(`
  <div class="card">
    <h3>`)
//line views/vworkspace/vwstandup/StandupList.html:28
	components.StreamSVGRefIcon(qw422016, util.KeyStandup, ps)
//line views/vworkspace/vwstandup/StandupList.html:28
	qw422016.E().S(util.StringPlural(len(p.Standups), "Standup"))
//line views/vworkspace/vwstandup/StandupList.html:28
	qw422016.N().S(`</h3>
    <em>`)
//line views/vworkspace/vwstandup/StandupList.html:29
	qw422016.E().S(util.KeyStandupDesc)
//line views/vworkspace/vwstandup/StandupList.html:29
	qw422016.N().S(`</em>
    <table class="mt expanded">
      <tbody>
`)
//line views/vworkspace/vwstandup/StandupList.html:32
	for _, u := range p.Standups {
//line views/vworkspace/vwstandup/StandupList.html:32
		qw422016.N().S(`        <tr>
          <td><a href="`)
//line views/vworkspace/vwstandup/StandupList.html:34
		qw422016.E().S(u.PublicWebPath())
//line views/vworkspace/vwstandup/StandupList.html:34
		qw422016.N().S(`">`)
//line views/vworkspace/vwstandup/StandupList.html:34
		components.StreamSVGRef(qw422016, u.IconSafe(), 16, 16, "icon", ps)
//line views/vworkspace/vwstandup/StandupList.html:34
		qw422016.E().S(u.TitleString())
//line views/vworkspace/vwstandup/StandupList.html:34
		qw422016.N().S(`</a></td>
          <td class="text-align-right">`)
//line views/vworkspace/vwstandup/StandupList.html:35
		components.StreamDisplayTimestamp(qw422016, &u.Created)
//line views/vworkspace/vwstandup/StandupList.html:35
		qw422016.N().S(`</td>
        </tr>
`)
//line views/vworkspace/vwstandup/StandupList.html:37
	}
//line views/vworkspace/vwstandup/StandupList.html:37
	qw422016.N().S(`      </tbody>
    </table>
  </div>
  <div class="card">
    <h3>`)
//line views/vworkspace/vwstandup/StandupList.html:42
	components.StreamSVGRefIcon(qw422016, util.KeyStandup, ps)
//line views/vworkspace/vwstandup/StandupList.html:42
	qw422016.N().S(`New Standup</h3>
    `)
//line views/vworkspace/vwstandup/StandupList.html:43
	StreamStandupForm(qw422016, &standup.Standup{}, p.Teams, p.Sprints, as, ps)
//line views/vworkspace/vwstandup/StandupList.html:43
	qw422016.N().S(`
  </div>
`)
//line views/vworkspace/vwstandup/StandupList.html:45
}

//line views/vworkspace/vwstandup/StandupList.html:45
func (p *StandupList) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vworkspace/vwstandup/StandupList.html:45
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwstandup/StandupList.html:45
	p.StreamBody(qw422016, as, ps)
//line views/vworkspace/vwstandup/StandupList.html:45
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwstandup/StandupList.html:45
}

//line views/vworkspace/vwstandup/StandupList.html:45
func (p *StandupList) Body(as *app.State, ps *cutil.PageState) string {
//line views/vworkspace/vwstandup/StandupList.html:45
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwstandup/StandupList.html:45
	p.WriteBody(qb422016, as, ps)
//line views/vworkspace/vwstandup/StandupList.html:45
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwstandup/StandupList.html:45
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwstandup/StandupList.html:45
	return qs422016
//line views/vworkspace/vwstandup/StandupList.html:45
}

//line views/vworkspace/vwstandup/StandupList.html:47
func StreamStandupForm(qw422016 *qt422016.Writer, s *standup.Standup, teams team.Teams, sprints sprint.Sprints, as *app.State, ps *cutil.PageState) {
//line views/vworkspace/vwstandup/StandupList.html:47
	qw422016.N().S(`
  <form action="" method="post">
    <table class="mt expanded">
      <tbody>
        `)
//line views/vworkspace/vwstandup/StandupList.html:51
	components.StreamTableInput(qw422016, "title", "", "Standup Title", s.Title, 5, "The name of your standup")
//line views/vworkspace/vwstandup/StandupList.html:51
	qw422016.N().S(`
        `)
//line views/vworkspace/vwstandup/StandupList.html:52
	components.StreamTableInput(qw422016, "name", "", "Your Name", ps.Username(), 5, "Whatever you prefer to be called")
//line views/vworkspace/vwstandup/StandupList.html:52
	qw422016.N().S(`
        `)
//line views/vworkspace/vwstandup/StandupList.html:53
	components.StreamTableSelect(qw422016, util.KeyTeam, "", "Team", fmt.Sprint(s.TeamID), teams.IDStrings(true), teams.TitleStrings("- no team -"), 5, "The team associated to this standup")
//line views/vworkspace/vwstandup/StandupList.html:53
	qw422016.N().S(`
        `)
//line views/vworkspace/vwstandup/StandupList.html:54
	components.StreamTableSelect(qw422016, util.KeySprint, "", "Sprint", fmt.Sprint(s.SprintID), sprints.IDStrings(true), sprints.TitleStrings("- no sprint -"), 5, "The sprint associated to this standup")
//line views/vworkspace/vwstandup/StandupList.html:54
	qw422016.N().S(`
        <tr><td colspan="2"><button type="submit">Add Standup</button></td></tr>
      </tbody>
    </table>
  </form>
`)
//line views/vworkspace/vwstandup/StandupList.html:59
}

//line views/vworkspace/vwstandup/StandupList.html:59
func WriteStandupForm(qq422016 qtio422016.Writer, s *standup.Standup, teams team.Teams, sprints sprint.Sprints, as *app.State, ps *cutil.PageState) {
//line views/vworkspace/vwstandup/StandupList.html:59
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwstandup/StandupList.html:59
	StreamStandupForm(qw422016, s, teams, sprints, as, ps)
//line views/vworkspace/vwstandup/StandupList.html:59
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwstandup/StandupList.html:59
}

//line views/vworkspace/vwstandup/StandupList.html:59
func StandupForm(s *standup.Standup, teams team.Teams, sprints sprint.Sprints, as *app.State, ps *cutil.PageState) string {
//line views/vworkspace/vwstandup/StandupList.html:59
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwstandup/StandupList.html:59
	WriteStandupForm(qb422016, s, teams, sprints, as, ps)
//line views/vworkspace/vwstandup/StandupList.html:59
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwstandup/StandupList.html:59
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwstandup/StandupList.html:59
	return qs422016
//line views/vworkspace/vwstandup/StandupList.html:59
}

//line views/vworkspace/vwstandup/StandupList.html:61
func StreamStandupListTable(qw422016 *qt422016.Writer, standups standup.Standups, teamID *uuid.UUID, sprintID *uuid.UUID, showComments bool, comments comment.Comments, members util.Members, ps *cutil.PageState) {
//line views/vworkspace/vwstandup/StandupList.html:61
	qw422016.N().S(`
  <div class="card">
    <div class="right">`)
//line views/vworkspace/vwstandup/StandupList.html:63
	vwutil.StreamEditWorkspaceForm(qw422016, util.KeyStandup, teamID, sprintID, "New Standup")
//line views/vworkspace/vwstandup/StandupList.html:63
	qw422016.N().S(`</div>
    <h3 title="`)
//line views/vworkspace/vwstandup/StandupList.html:64
	qw422016.E().S(util.KeyStandupDesc)
//line views/vworkspace/vwstandup/StandupList.html:64
	qw422016.N().S(`">`)
//line views/vworkspace/vwstandup/StandupList.html:64
	components.StreamSVGRefIcon(qw422016, util.KeyStandup, ps)
//line views/vworkspace/vwstandup/StandupList.html:64
	qw422016.N().S(`Standups</h3>
    <table id="standup-list" class="mt expanded">
      <tbody>
`)
//line views/vworkspace/vwstandup/StandupList.html:67
	if len(standups) == 0 {
//line views/vworkspace/vwstandup/StandupList.html:67
		qw422016.N().S(`          <tr class="empty"><td><em>no standups</em></td></tr>
`)
//line views/vworkspace/vwstandup/StandupList.html:69
	} else {
//line views/vworkspace/vwstandup/StandupList.html:70
		for _, x := range standups {
//line views/vworkspace/vwstandup/StandupList.html:70
			qw422016.N().S(`          <tr id="standup-list-`)
//line views/vworkspace/vwstandup/StandupList.html:71
			qw422016.E().S(x.ID.String())
//line views/vworkspace/vwstandup/StandupList.html:71
			qw422016.N().S(`">
            <td>
`)
//line views/vworkspace/vwstandup/StandupList.html:73
			if showComments {
//line views/vworkspace/vwstandup/StandupList.html:73
				qw422016.N().S(`              <div class="right">
                `)
//line views/vworkspace/vwstandup/StandupList.html:75
				vwutil.StreamComments(qw422016, enum.ModelServiceStandup, x.ID, x.TitleString(), comments, nil, "member-icon", ps)
//line views/vworkspace/vwstandup/StandupList.html:75
				qw422016.N().S(`
              </div>
`)
//line views/vworkspace/vwstandup/StandupList.html:77
			}
//line views/vworkspace/vwstandup/StandupList.html:77
			qw422016.N().S(`              <a href="`)
//line views/vworkspace/vwstandup/StandupList.html:78
			qw422016.E().S(x.PublicWebPath())
//line views/vworkspace/vwstandup/StandupList.html:78
			qw422016.N().S(`"><div>
                <span>`)
//line views/vworkspace/vwstandup/StandupList.html:79
			components.StreamSVGRef(qw422016, x.IconSafe(), 16, 16, "icon", ps)
//line views/vworkspace/vwstandup/StandupList.html:79
			qw422016.N().S(`</span><span>`)
//line views/vworkspace/vwstandup/StandupList.html:79
			qw422016.E().S(x.TitleString())
//line views/vworkspace/vwstandup/StandupList.html:79
			qw422016.N().S(`</span>
              </div></a>
            </td>
          </tr>
`)
//line views/vworkspace/vwstandup/StandupList.html:83
		}
//line views/vworkspace/vwstandup/StandupList.html:84
	}
//line views/vworkspace/vwstandup/StandupList.html:84
	qw422016.N().S(`      </tbody>
    </table>
  </div>
`)
//line views/vworkspace/vwstandup/StandupList.html:88
}

//line views/vworkspace/vwstandup/StandupList.html:88
func WriteStandupListTable(qq422016 qtio422016.Writer, standups standup.Standups, teamID *uuid.UUID, sprintID *uuid.UUID, showComments bool, comments comment.Comments, members util.Members, ps *cutil.PageState) {
//line views/vworkspace/vwstandup/StandupList.html:88
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwstandup/StandupList.html:88
	StreamStandupListTable(qw422016, standups, teamID, sprintID, showComments, comments, members, ps)
//line views/vworkspace/vwstandup/StandupList.html:88
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwstandup/StandupList.html:88
}

//line views/vworkspace/vwstandup/StandupList.html:88
func StandupListTable(standups standup.Standups, teamID *uuid.UUID, sprintID *uuid.UUID, showComments bool, comments comment.Comments, members util.Members, ps *cutil.PageState) string {
//line views/vworkspace/vwstandup/StandupList.html:88
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwstandup/StandupList.html:88
	WriteStandupListTable(qb422016, standups, teamID, sprintID, showComments, comments, members, ps)
//line views/vworkspace/vwstandup/StandupList.html:88
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwstandup/StandupList.html:88
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwstandup/StandupList.html:88
	return qs422016
//line views/vworkspace/vwstandup/StandupList.html:88
}
