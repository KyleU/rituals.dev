// Code generated by qtc from "Edit.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vretro/vrmember/Edit.html:2
package vrmember

//line views/vretro/vrmember/Edit.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/retro/rmember"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vretro/vrmember/Edit.html:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vretro/vrmember/Edit.html:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vretro/vrmember/Edit.html:10
type Edit struct {
	layout.Basic
	Model *rmember.RetroMember
	IsNew bool
}

//line views/vretro/vrmember/Edit.html:16
func (p *Edit) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vretro/vrmember/Edit.html:16
	qw422016.N().S(`
  <div class="card">
`)
//line views/vretro/vrmember/Edit.html:18
	if p.IsNew {
//line views/vretro/vrmember/Edit.html:18
		qw422016.N().S(`    <div class="right"><a href="/admin/db/retro/member/random"><button>Random</button></a></div>
    <h3>`)
//line views/vretro/vrmember/Edit.html:20
		components.StreamSVGRefIcon(qw422016, `users`, ps)
//line views/vretro/vrmember/Edit.html:20
		qw422016.N().S(` New Member</h3>
    <form action="/admin/db/retro/member/new" class="mt" method="post">
`)
//line views/vretro/vrmember/Edit.html:22
	} else {
//line views/vretro/vrmember/Edit.html:22
		qw422016.N().S(`    <div class="right"><a href="`)
//line views/vretro/vrmember/Edit.html:23
		qw422016.E().S(p.Model.WebPath())
//line views/vretro/vrmember/Edit.html:23
		qw422016.N().S(`/delete" onclick="return confirm('Are you sure you wish to delete member [`)
//line views/vretro/vrmember/Edit.html:23
		qw422016.E().S(p.Model.String())
//line views/vretro/vrmember/Edit.html:23
		qw422016.N().S(`]?')"><button>Delete</button></a></div>
    <h3>`)
//line views/vretro/vrmember/Edit.html:24
		components.StreamSVGRefIcon(qw422016, `users`, ps)
//line views/vretro/vrmember/Edit.html:24
		qw422016.N().S(` Edit Member [`)
//line views/vretro/vrmember/Edit.html:24
		qw422016.E().S(p.Model.String())
//line views/vretro/vrmember/Edit.html:24
		qw422016.N().S(`]</h3>
`)
//line views/vretro/vrmember/Edit.html:25
	}
//line views/vretro/vrmember/Edit.html:25
	qw422016.N().S(`    <form action="" method="post">
      <table class="mt expanded">
        <tbody>
          `)
//line views/vretro/vrmember/Edit.html:29
	if p.IsNew {
//line views/vretro/vrmember/Edit.html:29
		components.StreamTableInputUUID(qw422016, "retroID", "Retro ID", &p.Model.RetroID, 5, "UUID in format (00000000-0000-0000-0000-000000000000)")
//line views/vretro/vrmember/Edit.html:29
	}
//line views/vretro/vrmember/Edit.html:29
	qw422016.N().S(`
          `)
//line views/vretro/vrmember/Edit.html:30
	if p.IsNew {
//line views/vretro/vrmember/Edit.html:30
		components.StreamTableInputUUID(qw422016, "userID", "User ID", &p.Model.UserID, 5, "UUID in format (00000000-0000-0000-0000-000000000000)")
//line views/vretro/vrmember/Edit.html:30
	}
//line views/vretro/vrmember/Edit.html:30
	qw422016.N().S(`
          `)
//line views/vretro/vrmember/Edit.html:31
	components.StreamTableInput(qw422016, "name", "Name", p.Model.Name, 5, "String text")
//line views/vretro/vrmember/Edit.html:31
	qw422016.N().S(`
          `)
//line views/vretro/vrmember/Edit.html:32
	components.StreamTableInput(qw422016, "picture", "Picture", p.Model.Picture, 5, "URL in string form")
//line views/vretro/vrmember/Edit.html:32
	qw422016.N().S(`
          `)
//line views/vretro/vrmember/Edit.html:33
	components.StreamTableInput(qw422016, "role", "Role", p.Model.Role, 5, "String text")
//line views/vretro/vrmember/Edit.html:33
	qw422016.N().S(`
          <tr><td colspan="2"><button type="submit">Save Changes</button></td></tr>
        </tbody>
      </table>
    </form>
  </div>
`)
//line views/vretro/vrmember/Edit.html:39
}

//line views/vretro/vrmember/Edit.html:39
func (p *Edit) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vretro/vrmember/Edit.html:39
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vretro/vrmember/Edit.html:39
	p.StreamBody(qw422016, as, ps)
//line views/vretro/vrmember/Edit.html:39
	qt422016.ReleaseWriter(qw422016)
//line views/vretro/vrmember/Edit.html:39
}

//line views/vretro/vrmember/Edit.html:39
func (p *Edit) Body(as *app.State, ps *cutil.PageState) string {
//line views/vretro/vrmember/Edit.html:39
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vretro/vrmember/Edit.html:39
	p.WriteBody(qb422016, as, ps)
//line views/vretro/vrmember/Edit.html:39
	qs422016 := string(qb422016.B)
//line views/vretro/vrmember/Edit.html:39
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vretro/vrmember/Edit.html:39
	return qs422016
//line views/vretro/vrmember/Edit.html:39
}
