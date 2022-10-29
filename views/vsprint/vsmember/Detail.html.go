// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vsprint/vsmember/Detail.html:2
package vsmember

//line views/vsprint/vsmember/Detail.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/sprint/smember"
	"github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vsprint/vsmember/Detail.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vsprint/vsmember/Detail.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vsprint/vsmember/Detail.html:12
type Detail struct {
	layout.Basic
	Model   *smember.SprintMember
	Sprints sprint.Sprints
	Users   user.Users
}

//line views/vsprint/vsmember/Detail.html:19
func (p *Detail) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsprint/vsmember/Detail.html:19
	qw422016.N().S(`
  <div class="card">
    <div class="right">
      <a href="#modal-sprintMember"><button type="button">JSON</button></a>
      <a href="`)
//line views/vsprint/vsmember/Detail.html:23
	qw422016.E().S(p.Model.WebPath())
//line views/vsprint/vsmember/Detail.html:23
	qw422016.N().S(`/edit"><button>`)
//line views/vsprint/vsmember/Detail.html:23
	components.StreamSVGRef(qw422016, "edit", 15, 15, "icon", ps)
//line views/vsprint/vsmember/Detail.html:23
	qw422016.N().S(`Edit</button></a>
    </div>
    <h3>`)
//line views/vsprint/vsmember/Detail.html:25
	components.StreamSVGRefIcon(qw422016, `users`, ps)
//line views/vsprint/vsmember/Detail.html:25
	qw422016.N().S(` `)
//line views/vsprint/vsmember/Detail.html:25
	qw422016.E().S(p.Model.TitleString())
//line views/vsprint/vsmember/Detail.html:25
	qw422016.N().S(`</h3>
    <div><a href="/admin/db/sprint/member"><em>Member</em></a></div>
    <table class="mt">
      <tbody>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">Sprint ID</th>
          <td>
            <div class="icon">`)
//line views/vsprint/vsmember/Detail.html:32
	components.StreamDisplayUUID(qw422016, &p.Model.SprintID)
//line views/vsprint/vsmember/Detail.html:32
	if x := p.Sprints.Get(p.Model.SprintID); x != nil {
//line views/vsprint/vsmember/Detail.html:32
		qw422016.N().S(` (`)
//line views/vsprint/vsmember/Detail.html:32
		qw422016.E().S(x.TitleString())
//line views/vsprint/vsmember/Detail.html:32
		qw422016.N().S(`)`)
//line views/vsprint/vsmember/Detail.html:32
	}
//line views/vsprint/vsmember/Detail.html:32
	qw422016.N().S(`</div>
            <a title="Sprint" href="`)
//line views/vsprint/vsmember/Detail.html:33
	qw422016.E().S(`/sprint` + `/` + p.Model.SprintID.String())
//line views/vsprint/vsmember/Detail.html:33
	qw422016.N().S(`">`)
//line views/vsprint/vsmember/Detail.html:33
	components.StreamSVGRefIcon(qw422016, "sprint", ps)
//line views/vsprint/vsmember/Detail.html:33
	qw422016.N().S(`</a>
          </td>
        </tr>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">User ID</th>
          <td>
            <div class="icon">`)
//line views/vsprint/vsmember/Detail.html:39
	components.StreamDisplayUUID(qw422016, &p.Model.UserID)
//line views/vsprint/vsmember/Detail.html:39
	if x := p.Users.Get(p.Model.UserID); x != nil {
//line views/vsprint/vsmember/Detail.html:39
		qw422016.N().S(` (`)
//line views/vsprint/vsmember/Detail.html:39
		qw422016.E().S(x.TitleString())
//line views/vsprint/vsmember/Detail.html:39
		qw422016.N().S(`)`)
//line views/vsprint/vsmember/Detail.html:39
	}
//line views/vsprint/vsmember/Detail.html:39
	qw422016.N().S(`</div>
            <a title="User" href="`)
//line views/vsprint/vsmember/Detail.html:40
	qw422016.E().S(`/user` + `/` + p.Model.UserID.String())
//line views/vsprint/vsmember/Detail.html:40
	qw422016.N().S(`">`)
//line views/vsprint/vsmember/Detail.html:40
	components.StreamSVGRefIcon(qw422016, "profile", ps)
//line views/vsprint/vsmember/Detail.html:40
	qw422016.N().S(`</a>
          </td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Name</th>
          <td><strong>`)
//line views/vsprint/vsmember/Detail.html:45
	qw422016.E().S(p.Model.Name)
//line views/vsprint/vsmember/Detail.html:45
	qw422016.N().S(`</strong></td>
        </tr>
        <tr>
          <th class="shrink" title="URL in string form">Picture</th>
          <td><a href="`)
//line views/vsprint/vsmember/Detail.html:49
	qw422016.E().S(p.Model.Picture)
//line views/vsprint/vsmember/Detail.html:49
	qw422016.N().S(`" target="_blank">`)
//line views/vsprint/vsmember/Detail.html:49
	qw422016.E().S(p.Model.Picture)
//line views/vsprint/vsmember/Detail.html:49
	qw422016.N().S(`</a></td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Role</th>
          <td>`)
//line views/vsprint/vsmember/Detail.html:53
	qw422016.E().S(p.Model.Role)
//line views/vsprint/vsmember/Detail.html:53
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="Date and time, in almost any format">Created</th>
          <td>`)
//line views/vsprint/vsmember/Detail.html:57
	components.StreamDisplayTimestamp(qw422016, &p.Model.Created)
//line views/vsprint/vsmember/Detail.html:57
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="Date and time, in almost any format (optional)">Updated</th>
          <td>`)
//line views/vsprint/vsmember/Detail.html:61
	components.StreamDisplayTimestamp(qw422016, p.Model.Updated)
//line views/vsprint/vsmember/Detail.html:61
	qw422016.N().S(`</td>
        </tr>
      </tbody>
    </table>
  </div>
`)
//line views/vsprint/vsmember/Detail.html:67
	qw422016.N().S(`  `)
//line views/vsprint/vsmember/Detail.html:68
	components.StreamJSONModal(qw422016, "sprintMember", "Member JSON", p.Model, 1)
//line views/vsprint/vsmember/Detail.html:68
	qw422016.N().S(`
`)
//line views/vsprint/vsmember/Detail.html:69
}

//line views/vsprint/vsmember/Detail.html:69
func (p *Detail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsprint/vsmember/Detail.html:69
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vsprint/vsmember/Detail.html:69
	p.StreamBody(qw422016, as, ps)
//line views/vsprint/vsmember/Detail.html:69
	qt422016.ReleaseWriter(qw422016)
//line views/vsprint/vsmember/Detail.html:69
}

//line views/vsprint/vsmember/Detail.html:69
func (p *Detail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vsprint/vsmember/Detail.html:69
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vsprint/vsmember/Detail.html:69
	p.WriteBody(qb422016, as, ps)
//line views/vsprint/vsmember/Detail.html:69
	qs422016 := string(qb422016.B)
//line views/vsprint/vsmember/Detail.html:69
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vsprint/vsmember/Detail.html:69
	return qs422016
//line views/vsprint/vsmember/Detail.html:69
}
