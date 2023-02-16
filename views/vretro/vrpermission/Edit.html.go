// Code generated by qtc from "Edit.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vretro/vrpermission/Edit.html:2
package vrpermission

//line views/vretro/vrpermission/Edit.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/retro/rpermission"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vretro/vrpermission/Edit.html:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vretro/vrpermission/Edit.html:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vretro/vrpermission/Edit.html:10
type Edit struct {
	layout.Basic
	Model *rpermission.RetroPermission
	IsNew bool
}

//line views/vretro/vrpermission/Edit.html:16
func (p *Edit) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vretro/vrpermission/Edit.html:16
	qw422016.N().S(`
  <div class="card">
`)
//line views/vretro/vrpermission/Edit.html:18
	if p.IsNew {
//line views/vretro/vrpermission/Edit.html:18
		qw422016.N().S(`    <div class="right"><a href="/admin/db/retro/permission/random"><button>Random</button></a></div>
    <h3>`)
//line views/vretro/vrpermission/Edit.html:20
		components.StreamSVGRefIcon(qw422016, `permission`, ps)
//line views/vretro/vrpermission/Edit.html:20
		qw422016.N().S(` New Permission</h3>
    <form action="/admin/db/retro/permission/new" class="mt" method="post">
`)
//line views/vretro/vrpermission/Edit.html:22
	} else {
//line views/vretro/vrpermission/Edit.html:22
		qw422016.N().S(`    <div class="right"><a href="`)
//line views/vretro/vrpermission/Edit.html:23
		qw422016.E().S(p.Model.WebPath())
//line views/vretro/vrpermission/Edit.html:23
		qw422016.N().S(`/delete" onclick="return confirm('Are you sure you wish to delete permission [`)
//line views/vretro/vrpermission/Edit.html:23
		qw422016.E().S(p.Model.String())
//line views/vretro/vrpermission/Edit.html:23
		qw422016.N().S(`]?')"><button>Delete</button></a></div>
    <h3>`)
//line views/vretro/vrpermission/Edit.html:24
		components.StreamSVGRefIcon(qw422016, `permission`, ps)
//line views/vretro/vrpermission/Edit.html:24
		qw422016.N().S(` Edit Permission [`)
//line views/vretro/vrpermission/Edit.html:24
		qw422016.E().S(p.Model.String())
//line views/vretro/vrpermission/Edit.html:24
		qw422016.N().S(`]</h3>
    <form action="" method="post">
`)
//line views/vretro/vrpermission/Edit.html:26
	}
//line views/vretro/vrpermission/Edit.html:26
	qw422016.N().S(`      <table class="mt expanded">
        <tbody>
          `)
//line views/vretro/vrpermission/Edit.html:29
	if p.IsNew {
//line views/vretro/vrpermission/Edit.html:29
		components.StreamTableInputUUID(qw422016, "retroID", "input-retroID", "Retro ID", &p.Model.RetroID, 5, "UUID in format (00000000-0000-0000-0000-000000000000)")
//line views/vretro/vrpermission/Edit.html:29
	}
//line views/vretro/vrpermission/Edit.html:29
	qw422016.N().S(`
          `)
//line views/vretro/vrpermission/Edit.html:30
	if p.IsNew {
//line views/vretro/vrpermission/Edit.html:30
		components.StreamTableInput(qw422016, "key", "", "Key", p.Model.Key, 5, "String text")
//line views/vretro/vrpermission/Edit.html:30
	}
//line views/vretro/vrpermission/Edit.html:30
	qw422016.N().S(`
          `)
//line views/vretro/vrpermission/Edit.html:31
	if p.IsNew {
//line views/vretro/vrpermission/Edit.html:31
		components.StreamTableInput(qw422016, "value", "", "Value", p.Model.Value, 5, "String text")
//line views/vretro/vrpermission/Edit.html:31
	}
//line views/vretro/vrpermission/Edit.html:31
	qw422016.N().S(`
          `)
//line views/vretro/vrpermission/Edit.html:32
	components.StreamTableInput(qw422016, "access", "", "Access", p.Model.Access, 5, "String text")
//line views/vretro/vrpermission/Edit.html:32
	qw422016.N().S(`
          <tr><td colspan="2"><button type="submit">Save Changes</button></td></tr>
        </tbody>
      </table>
    </form>
  </div>
  <script>
    document.addEventListener("DOMContentLoaded", function() {
      rituals.autocomplete(document.getElementById("input-retroID"), "/admin/db/retro?retro.l=10", "q", (o) => o["slug"] + " / " + o["title"] + " / " + o["categories"] + " (" + o["id"] + ")", (o) => o["id"]);
    });
  </script>
`)
//line views/vretro/vrpermission/Edit.html:43
}

//line views/vretro/vrpermission/Edit.html:43
func (p *Edit) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vretro/vrpermission/Edit.html:43
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vretro/vrpermission/Edit.html:43
	p.StreamBody(qw422016, as, ps)
//line views/vretro/vrpermission/Edit.html:43
	qt422016.ReleaseWriter(qw422016)
//line views/vretro/vrpermission/Edit.html:43
}

//line views/vretro/vrpermission/Edit.html:43
func (p *Edit) Body(as *app.State, ps *cutil.PageState) string {
//line views/vretro/vrpermission/Edit.html:43
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vretro/vrpermission/Edit.html:43
	p.WriteBody(qb422016, as, ps)
//line views/vretro/vrpermission/Edit.html:43
	qs422016 := string(qb422016.B)
//line views/vretro/vrpermission/Edit.html:43
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vretro/vrpermission/Edit.html:43
	return qs422016
//line views/vretro/vrpermission/Edit.html:43
}
