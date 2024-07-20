// Code generated by qtc from "Edit.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vretro/vrmember/Edit.html:1
package vrmember

//line views/vretro/vrmember/Edit.html:1
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/retro/rmember"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/components/edit"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vretro/vrmember/Edit.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vretro/vrmember/Edit.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vretro/vrmember/Edit.html:12
type Edit struct {
	layout.Basic
	Model *rmember.RetroMember
	Paths []string
	IsNew bool
}

//line views/vretro/vrmember/Edit.html:19
func (p *Edit) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vretro/vrmember/Edit.html:19
	qw422016.N().S(`
  <div class="card">
`)
//line views/vretro/vrmember/Edit.html:21
	if p.IsNew {
//line views/vretro/vrmember/Edit.html:21
		qw422016.N().S(`    <div class="right"><a href="?prototype=random"><button>Random</button></a></div>
    <h3>`)
//line views/vretro/vrmember/Edit.html:23
		components.StreamSVGIcon(qw422016, `users`, ps)
//line views/vretro/vrmember/Edit.html:23
		qw422016.N().S(` New Member</h3>
`)
//line views/vretro/vrmember/Edit.html:24
	} else {
//line views/vretro/vrmember/Edit.html:24
		qw422016.N().S(`    <div class="right"><a class="link-confirm" href="`)
//line views/vretro/vrmember/Edit.html:25
		qw422016.E().S(p.Model.WebPath(p.Paths...))
//line views/vretro/vrmember/Edit.html:25
		qw422016.N().S(`/delete" data-message="Are you sure you wish to delete member [`)
//line views/vretro/vrmember/Edit.html:25
		qw422016.E().S(p.Model.String())
//line views/vretro/vrmember/Edit.html:25
		qw422016.N().S(`]?"><button>`)
//line views/vretro/vrmember/Edit.html:25
		components.StreamSVGButton(qw422016, "times", ps)
//line views/vretro/vrmember/Edit.html:25
		qw422016.N().S(` Delete</button></a></div>
    <h3>`)
//line views/vretro/vrmember/Edit.html:26
		components.StreamSVGIcon(qw422016, `users`, ps)
//line views/vretro/vrmember/Edit.html:26
		qw422016.N().S(` Edit Member [`)
//line views/vretro/vrmember/Edit.html:26
		qw422016.E().S(p.Model.String())
//line views/vretro/vrmember/Edit.html:26
		qw422016.N().S(`]</h3>
`)
//line views/vretro/vrmember/Edit.html:27
	}
//line views/vretro/vrmember/Edit.html:27
	qw422016.N().S(`    <form action="`)
//line views/vretro/vrmember/Edit.html:28
	qw422016.E().S(util.Choose(p.IsNew, rmember.Route(p.Paths...)+`/_new`, p.Model.WebPath(p.Paths...)+`/edit`))
//line views/vretro/vrmember/Edit.html:28
	qw422016.N().S(`" class="mt" method="post">
      <table class="mt expanded">
        <tbody>
          `)
//line views/vretro/vrmember/Edit.html:31
	if p.IsNew {
//line views/vretro/vrmember/Edit.html:31
		edit.StreamUUIDTable(qw422016, "retroID", "input-retroID", "Retro ID", &p.Model.RetroID, 5, "UUID in format (00000000-0000-0000-0000-000000000000)")
//line views/vretro/vrmember/Edit.html:31
	}
//line views/vretro/vrmember/Edit.html:31
	qw422016.N().S(`
          `)
//line views/vretro/vrmember/Edit.html:32
	if p.IsNew {
//line views/vretro/vrmember/Edit.html:32
		edit.StreamUUIDTable(qw422016, "userID", "input-userID", "User ID", &p.Model.UserID, 5, "UUID in format (00000000-0000-0000-0000-000000000000)")
//line views/vretro/vrmember/Edit.html:32
	}
//line views/vretro/vrmember/Edit.html:32
	qw422016.N().S(`
          `)
//line views/vretro/vrmember/Edit.html:33
	edit.StreamStringTable(qw422016, "name", "", "Name", p.Model.Name, 5, "String text")
//line views/vretro/vrmember/Edit.html:33
	qw422016.N().S(`
          `)
//line views/vretro/vrmember/Edit.html:34
	edit.StreamStringTable(qw422016, "picture", "", "Picture", p.Model.Picture, 5, "URL in string form")
//line views/vretro/vrmember/Edit.html:34
	qw422016.N().S(`
          `)
//line views/vretro/vrmember/Edit.html:35
	edit.StreamSelectTable(qw422016, "role", "", "Role", p.Model.Role.Key, enum.AllMemberStatuses.Keys(), enum.AllMemberStatuses.Strings(), 5, enum.AllMemberStatuses.Help())
//line views/vretro/vrmember/Edit.html:35
	qw422016.N().S(`
          <tr><td colspan="2"><button type="submit">Save Changes</button></td></tr>
        </tbody>
      </table>
    </form>
  </div>
  <script>
    document.addEventListener("DOMContentLoaded", function() {
      rituals.autocomplete(document.getElementById("input-retroID"), "/admin/db/retro?retro.l=10", "q", (o) => o["slug"] + " / " + o["title"] + " / " + o["categories"] + " (" + o["id"] + ")", (o) => o["id"]);
      rituals.autocomplete(document.getElementById("input-userID"), "/admin/db/user?user.l=10", "q", (o) => o["name"] + " (" + o["id"] + ")", (o) => o["id"]);
    });
  </script>
`)
//line views/vretro/vrmember/Edit.html:47
}

//line views/vretro/vrmember/Edit.html:47
func (p *Edit) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vretro/vrmember/Edit.html:47
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vretro/vrmember/Edit.html:47
	p.StreamBody(qw422016, as, ps)
//line views/vretro/vrmember/Edit.html:47
	qt422016.ReleaseWriter(qw422016)
//line views/vretro/vrmember/Edit.html:47
}

//line views/vretro/vrmember/Edit.html:47
func (p *Edit) Body(as *app.State, ps *cutil.PageState) string {
//line views/vretro/vrmember/Edit.html:47
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vretro/vrmember/Edit.html:47
	p.WriteBody(qb422016, as, ps)
//line views/vretro/vrmember/Edit.html:47
	qs422016 := string(qb422016.B)
//line views/vretro/vrmember/Edit.html:47
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vretro/vrmember/Edit.html:47
	return qs422016
//line views/vretro/vrmember/Edit.html:47
}
