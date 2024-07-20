// Code generated by qtc from "Edit.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vteam/vtpermission/Edit.html:1
package vtpermission

//line views/vteam/vtpermission/Edit.html:1
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/team/tpermission"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/components/edit"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vteam/vtpermission/Edit.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vteam/vtpermission/Edit.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vteam/vtpermission/Edit.html:11
type Edit struct {
	layout.Basic
	Model *tpermission.TeamPermission
	Paths []string
	IsNew bool
}

//line views/vteam/vtpermission/Edit.html:18
func (p *Edit) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vteam/vtpermission/Edit.html:18
	qw422016.N().S(`
  <div class="card">
`)
//line views/vteam/vtpermission/Edit.html:20
	if p.IsNew {
//line views/vteam/vtpermission/Edit.html:20
		qw422016.N().S(`    <div class="right"><a href="?prototype=random"><button>Random</button></a></div>
    <h3>`)
//line views/vteam/vtpermission/Edit.html:22
		components.StreamSVGIcon(qw422016, `permission`, ps)
//line views/vteam/vtpermission/Edit.html:22
		qw422016.N().S(` New Permission</h3>
`)
//line views/vteam/vtpermission/Edit.html:23
	} else {
//line views/vteam/vtpermission/Edit.html:23
		qw422016.N().S(`    <div class="right"><a class="link-confirm" href="`)
//line views/vteam/vtpermission/Edit.html:24
		qw422016.E().S(p.Model.WebPath(p.Paths...))
//line views/vteam/vtpermission/Edit.html:24
		qw422016.N().S(`/delete" data-message="Are you sure you wish to delete permission [`)
//line views/vteam/vtpermission/Edit.html:24
		qw422016.E().S(p.Model.String())
//line views/vteam/vtpermission/Edit.html:24
		qw422016.N().S(`]?"><button>`)
//line views/vteam/vtpermission/Edit.html:24
		components.StreamSVGButton(qw422016, "times", ps)
//line views/vteam/vtpermission/Edit.html:24
		qw422016.N().S(` Delete</button></a></div>
    <h3>`)
//line views/vteam/vtpermission/Edit.html:25
		components.StreamSVGIcon(qw422016, `permission`, ps)
//line views/vteam/vtpermission/Edit.html:25
		qw422016.N().S(` Edit Permission [`)
//line views/vteam/vtpermission/Edit.html:25
		qw422016.E().S(p.Model.String())
//line views/vteam/vtpermission/Edit.html:25
		qw422016.N().S(`]</h3>
`)
//line views/vteam/vtpermission/Edit.html:26
	}
//line views/vteam/vtpermission/Edit.html:26
	qw422016.N().S(`    <form action="`)
//line views/vteam/vtpermission/Edit.html:27
	qw422016.E().S(util.Choose(p.IsNew, tpermission.Route(p.Paths...)+`/_new`, p.Model.WebPath(p.Paths...)+`/edit`))
//line views/vteam/vtpermission/Edit.html:27
	qw422016.N().S(`" class="mt" method="post">
      <table class="mt expanded">
        <tbody>
          `)
//line views/vteam/vtpermission/Edit.html:30
	if p.IsNew {
//line views/vteam/vtpermission/Edit.html:30
		edit.StreamUUIDTable(qw422016, "teamID", "input-teamID", "Team ID", &p.Model.TeamID, 5, "UUID in format (00000000-0000-0000-0000-000000000000)")
//line views/vteam/vtpermission/Edit.html:30
	}
//line views/vteam/vtpermission/Edit.html:30
	qw422016.N().S(`
          `)
//line views/vteam/vtpermission/Edit.html:31
	if p.IsNew {
//line views/vteam/vtpermission/Edit.html:31
		edit.StreamStringTable(qw422016, "key", "", "Key", p.Model.Key, 5, "String text")
//line views/vteam/vtpermission/Edit.html:31
	}
//line views/vteam/vtpermission/Edit.html:31
	qw422016.N().S(`
          `)
//line views/vteam/vtpermission/Edit.html:32
	if p.IsNew {
//line views/vteam/vtpermission/Edit.html:32
		edit.StreamStringTable(qw422016, "value", "", "Value", p.Model.Value, 5, "String text")
//line views/vteam/vtpermission/Edit.html:32
	}
//line views/vteam/vtpermission/Edit.html:32
	qw422016.N().S(`
          `)
//line views/vteam/vtpermission/Edit.html:33
	edit.StreamStringTable(qw422016, "access", "", "Access", p.Model.Access, 5, "String text")
//line views/vteam/vtpermission/Edit.html:33
	qw422016.N().S(`
          <tr><td colspan="2"><button type="submit">Save Changes</button></td></tr>
        </tbody>
      </table>
    </form>
  </div>
  <script>
    document.addEventListener("DOMContentLoaded", function() {
      rituals.autocomplete(document.getElementById("input-teamID"), "/admin/db/team?team.l=10", "q", (o) => o["slug"] + " / " + o["title"] + " (" + o["id"] + ")", (o) => o["id"]);
    });
  </script>
`)
//line views/vteam/vtpermission/Edit.html:44
}

//line views/vteam/vtpermission/Edit.html:44
func (p *Edit) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vteam/vtpermission/Edit.html:44
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vteam/vtpermission/Edit.html:44
	p.StreamBody(qw422016, as, ps)
//line views/vteam/vtpermission/Edit.html:44
	qt422016.ReleaseWriter(qw422016)
//line views/vteam/vtpermission/Edit.html:44
}

//line views/vteam/vtpermission/Edit.html:44
func (p *Edit) Body(as *app.State, ps *cutil.PageState) string {
//line views/vteam/vtpermission/Edit.html:44
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vteam/vtpermission/Edit.html:44
	p.WriteBody(qb422016, as, ps)
//line views/vteam/vtpermission/Edit.html:44
	qs422016 := string(qb422016.B)
//line views/vteam/vtpermission/Edit.html:44
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vteam/vtpermission/Edit.html:44
	return qs422016
//line views/vteam/vtpermission/Edit.html:44
}
