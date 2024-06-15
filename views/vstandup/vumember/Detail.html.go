// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vstandup/vumember/Detail.html:2
package vumember

//line views/vstandup/vumember/Detail.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/standup"
	"github.com/kyleu/rituals/app/standup/umember"
	"github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/components/view"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vstandup/vumember/Detail.html:14
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vstandup/vumember/Detail.html:14
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vstandup/vumember/Detail.html:14
type Detail struct {
	layout.Basic
	Model              *umember.StandupMember
	StandupByStandupID *standup.Standup
	UserByUserID       *user.User
}

//line views/vstandup/vumember/Detail.html:21
func (p *Detail) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vstandup/vumember/Detail.html:21
	qw422016.N().S(`
  <div class="card">
    <div class="right">
      <a href="#modal-standupMember"><button type="button">`)
//line views/vstandup/vumember/Detail.html:24
	components.StreamSVGButton(qw422016, "file", ps)
//line views/vstandup/vumember/Detail.html:24
	qw422016.N().S(` JSON</button></a>
      <a href="`)
//line views/vstandup/vumember/Detail.html:25
	qw422016.E().S(p.Model.WebPath())
//line views/vstandup/vumember/Detail.html:25
	qw422016.N().S(`/edit"><button>`)
//line views/vstandup/vumember/Detail.html:25
	components.StreamSVGButton(qw422016, "edit", ps)
//line views/vstandup/vumember/Detail.html:25
	qw422016.N().S(` Edit</button></a>
    </div>
    <h3>`)
//line views/vstandup/vumember/Detail.html:27
	components.StreamSVGIcon(qw422016, `users`, ps)
//line views/vstandup/vumember/Detail.html:27
	qw422016.N().S(` `)
//line views/vstandup/vumember/Detail.html:27
	qw422016.E().S(p.Model.TitleString())
//line views/vstandup/vumember/Detail.html:27
	qw422016.N().S(`</h3>
    <div><a href="/admin/db/standup/member"><em>Member</em></a></div>
    <div class="mt overflow full-width">
      <table>
        <tbody>
          <tr>
            <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">Standup ID</th>
            <td class="nowrap">
              `)
//line views/vstandup/vumember/Detail.html:35
	view.StreamUUID(qw422016, &p.Model.StandupID)
//line views/vstandup/vumember/Detail.html:35
	if p.StandupByStandupID != nil {
//line views/vstandup/vumember/Detail.html:35
		qw422016.N().S(` (`)
//line views/vstandup/vumember/Detail.html:35
		qw422016.E().S(p.StandupByStandupID.TitleString())
//line views/vstandup/vumember/Detail.html:35
		qw422016.N().S(`)`)
//line views/vstandup/vumember/Detail.html:35
	}
//line views/vstandup/vumember/Detail.html:35
	qw422016.N().S(`
              <a title="Standup" href="`)
//line views/vstandup/vumember/Detail.html:36
	qw422016.E().S(`/admin/db/standup` + `/` + p.Model.StandupID.String())
//line views/vstandup/vumember/Detail.html:36
	qw422016.N().S(`">`)
//line views/vstandup/vumember/Detail.html:36
	components.StreamSVGLink(qw422016, `standup`, ps)
//line views/vstandup/vumember/Detail.html:36
	qw422016.N().S(`</a>
            </td>
          </tr>
          <tr>
            <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">User ID</th>
            <td class="nowrap">
              `)
//line views/vstandup/vumember/Detail.html:42
	view.StreamUUID(qw422016, &p.Model.UserID)
//line views/vstandup/vumember/Detail.html:42
	if p.UserByUserID != nil {
//line views/vstandup/vumember/Detail.html:42
		qw422016.N().S(` (`)
//line views/vstandup/vumember/Detail.html:42
		qw422016.E().S(p.UserByUserID.TitleString())
//line views/vstandup/vumember/Detail.html:42
		qw422016.N().S(`)`)
//line views/vstandup/vumember/Detail.html:42
	}
//line views/vstandup/vumember/Detail.html:42
	qw422016.N().S(`
              <a title="User" href="`)
//line views/vstandup/vumember/Detail.html:43
	qw422016.E().S(`/admin/db/user` + `/` + p.Model.UserID.String())
//line views/vstandup/vumember/Detail.html:43
	qw422016.N().S(`">`)
//line views/vstandup/vumember/Detail.html:43
	components.StreamSVGLink(qw422016, `profile`, ps)
//line views/vstandup/vumember/Detail.html:43
	qw422016.N().S(`</a>
            </td>
          </tr>
          <tr>
            <th class="shrink" title="String text">Name</th>
            <td><strong>`)
//line views/vstandup/vumember/Detail.html:48
	view.StreamString(qw422016, p.Model.Name)
//line views/vstandup/vumember/Detail.html:48
	qw422016.N().S(`</strong></td>
          </tr>
          <tr>
            <th class="shrink" title="URL in string form">Picture</th>
            <td><a href="`)
//line views/vstandup/vumember/Detail.html:52
	qw422016.E().S(p.Model.Picture)
//line views/vstandup/vumember/Detail.html:52
	qw422016.N().S(`" target="_blank" rel="noopener noreferrer">`)
//line views/vstandup/vumember/Detail.html:52
	qw422016.E().S(p.Model.Picture)
//line views/vstandup/vumember/Detail.html:52
	qw422016.N().S(`</a></td>
          </tr>
          <tr>
            <th class="shrink" title="`)
//line views/vstandup/vumember/Detail.html:55
	qw422016.E().S(enum.AllMemberStatuses.Help())
//line views/vstandup/vumember/Detail.html:55
	qw422016.N().S(`">Role</th>
            <td>`)
//line views/vstandup/vumember/Detail.html:56
	qw422016.E().S(p.Model.Role.String())
//line views/vstandup/vumember/Detail.html:56
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="Date and time, in almost any format">Created</th>
            <td>`)
//line views/vstandup/vumember/Detail.html:60
	view.StreamTimestamp(qw422016, &p.Model.Created)
//line views/vstandup/vumember/Detail.html:60
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="Date and time, in almost any format (optional)">Updated</th>
            <td>`)
//line views/vstandup/vumember/Detail.html:64
	view.StreamTimestamp(qw422016, p.Model.Updated)
//line views/vstandup/vumember/Detail.html:64
	qw422016.N().S(`</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
`)
//line views/vstandup/vumember/Detail.html:71
	qw422016.N().S(`  `)
//line views/vstandup/vumember/Detail.html:72
	components.StreamJSONModal(qw422016, "standupMember", "Member JSON", p.Model, 1)
//line views/vstandup/vumember/Detail.html:72
	qw422016.N().S(`
`)
//line views/vstandup/vumember/Detail.html:73
}

//line views/vstandup/vumember/Detail.html:73
func (p *Detail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vstandup/vumember/Detail.html:73
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vstandup/vumember/Detail.html:73
	p.StreamBody(qw422016, as, ps)
//line views/vstandup/vumember/Detail.html:73
	qt422016.ReleaseWriter(qw422016)
//line views/vstandup/vumember/Detail.html:73
}

//line views/vstandup/vumember/Detail.html:73
func (p *Detail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vstandup/vumember/Detail.html:73
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vstandup/vumember/Detail.html:73
	p.WriteBody(qb422016, as, ps)
//line views/vstandup/vumember/Detail.html:73
	qs422016 := string(qb422016.B)
//line views/vstandup/vumember/Detail.html:73
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vstandup/vumember/Detail.html:73
	return qs422016
//line views/vstandup/vumember/Detail.html:73
}
