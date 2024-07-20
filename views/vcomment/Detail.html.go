// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vcomment/Detail.html:1
package vcomment

//line views/vcomment/Detail.html:1
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/comment"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/components/view"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vcomment/Detail.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vcomment/Detail.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vcomment/Detail.html:12
type Detail struct {
	layout.Basic
	Model        *comment.Comment
	UserByUserID *user.User
	Paths        []string
}

//line views/vcomment/Detail.html:19
func (p *Detail) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vcomment/Detail.html:19
	qw422016.N().S(`
  <div class="card">
    <div class="right">
      <a href="#modal-comment"><button type="button">`)
//line views/vcomment/Detail.html:22
	components.StreamSVGButton(qw422016, "file", ps)
//line views/vcomment/Detail.html:22
	qw422016.N().S(` JSON</button></a>
      <a href="`)
//line views/vcomment/Detail.html:23
	qw422016.E().S(p.Model.WebPath(p.Paths...))
//line views/vcomment/Detail.html:23
	qw422016.N().S(`/edit"><button>`)
//line views/vcomment/Detail.html:23
	components.StreamSVGButton(qw422016, "edit", ps)
//line views/vcomment/Detail.html:23
	qw422016.N().S(` Edit</button></a>
    </div>
    <h3>`)
//line views/vcomment/Detail.html:25
	components.StreamSVGIcon(qw422016, `comments`, ps)
//line views/vcomment/Detail.html:25
	qw422016.N().S(` `)
//line views/vcomment/Detail.html:25
	qw422016.E().S(p.Model.TitleString())
//line views/vcomment/Detail.html:25
	qw422016.N().S(`</h3>
    <div><a href="`)
//line views/vcomment/Detail.html:26
	qw422016.E().S(comment.Route(p.Paths...))
//line views/vcomment/Detail.html:26
	qw422016.N().S(`"><em>Comment</em></a></div>
    <div class="mt overflow full-width">
      <table>
        <tbody>
          <tr>
            <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">ID</th>
            <td>`)
//line views/vcomment/Detail.html:32
	view.StreamUUID(qw422016, &p.Model.ID)
//line views/vcomment/Detail.html:32
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="`)
//line views/vcomment/Detail.html:35
	qw422016.E().S(enum.AllModelServices.Help())
//line views/vcomment/Detail.html:35
	qw422016.N().S(`">Svc</th>
            <td>`)
//line views/vcomment/Detail.html:36
	qw422016.E().S(p.Model.Svc.String())
//line views/vcomment/Detail.html:36
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">Model ID</th>
            <td>`)
//line views/vcomment/Detail.html:40
	view.StreamUUID(qw422016, &p.Model.ModelID)
//line views/vcomment/Detail.html:40
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">User ID</th>
            <td class="nowrap">
              `)
//line views/vcomment/Detail.html:45
	view.StreamUUID(qw422016, &p.Model.UserID)
//line views/vcomment/Detail.html:45
	if p.UserByUserID != nil {
//line views/vcomment/Detail.html:45
		qw422016.N().S(` (`)
//line views/vcomment/Detail.html:45
		qw422016.E().S(p.UserByUserID.TitleString())
//line views/vcomment/Detail.html:45
		qw422016.N().S(`)`)
//line views/vcomment/Detail.html:45
	}
//line views/vcomment/Detail.html:45
	qw422016.N().S(`
              <a title="User" href="`)
//line views/vcomment/Detail.html:46
	qw422016.E().S(p.Model.WebPath(p.Paths...))
//line views/vcomment/Detail.html:46
	qw422016.N().S(`">`)
//line views/vcomment/Detail.html:46
	components.StreamSVGLink(qw422016, `profile`, ps)
//line views/vcomment/Detail.html:46
	qw422016.N().S(`</a>
            </td>
          </tr>
          <tr>
            <th class="shrink" title="String text">Content</th>
            <td>`)
//line views/vcomment/Detail.html:51
	view.StreamString(qw422016, p.Model.Content)
//line views/vcomment/Detail.html:51
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="HTML code, in string form">HTML</th>
            <td>`)
//line views/vcomment/Detail.html:55
	view.StreamFormat(qw422016, p.Model.HTML, "html")
//line views/vcomment/Detail.html:55
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="Date and time, in almost any format">Created</th>
            <td>`)
//line views/vcomment/Detail.html:59
	view.StreamTimestamp(qw422016, &p.Model.Created)
//line views/vcomment/Detail.html:59
	qw422016.N().S(`</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
`)
//line views/vcomment/Detail.html:66
	qw422016.N().S(`  `)
//line views/vcomment/Detail.html:67
	components.StreamJSONModal(qw422016, "comment", "Comment JSON", p.Model, 1)
//line views/vcomment/Detail.html:67
	qw422016.N().S(`
`)
//line views/vcomment/Detail.html:68
}

//line views/vcomment/Detail.html:68
func (p *Detail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vcomment/Detail.html:68
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vcomment/Detail.html:68
	p.StreamBody(qw422016, as, ps)
//line views/vcomment/Detail.html:68
	qt422016.ReleaseWriter(qw422016)
//line views/vcomment/Detail.html:68
}

//line views/vcomment/Detail.html:68
func (p *Detail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vcomment/Detail.html:68
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vcomment/Detail.html:68
	p.WriteBody(qb422016, as, ps)
//line views/vcomment/Detail.html:68
	qs422016 := string(qb422016.B)
//line views/vcomment/Detail.html:68
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vcomment/Detail.html:68
	return qs422016
//line views/vcomment/Detail.html:68
}
