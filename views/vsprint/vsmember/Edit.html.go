// Code generated by qtc from "Edit.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vsprint/vsmember/Edit.html:2
package vsmember

//line views/vsprint/vsmember/Edit.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/sprint/smember"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/components/edit"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vsprint/vsmember/Edit.html:13
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vsprint/vsmember/Edit.html:13
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vsprint/vsmember/Edit.html:13
type Edit struct {
	layout.Basic
	Model *smember.SprintMember
	IsNew bool
}

//line views/vsprint/vsmember/Edit.html:19
func (p *Edit) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsprint/vsmember/Edit.html:19
	qw422016.N().S(`
  <div class="card">
`)
//line views/vsprint/vsmember/Edit.html:21
	if p.IsNew {
//line views/vsprint/vsmember/Edit.html:21
		qw422016.N().S(`    <div class="right"><a href="?prototype=random"><button>Random</button></a></div>
    <h3>`)
//line views/vsprint/vsmember/Edit.html:23
		components.StreamSVGIcon(qw422016, `users`, ps)
//line views/vsprint/vsmember/Edit.html:23
		qw422016.N().S(` New Member</h3>
`)
//line views/vsprint/vsmember/Edit.html:24
	} else {
//line views/vsprint/vsmember/Edit.html:24
		qw422016.N().S(`    <div class="right"><a class="link-confirm" href="`)
//line views/vsprint/vsmember/Edit.html:25
		qw422016.E().S(p.Model.WebPath())
//line views/vsprint/vsmember/Edit.html:25
		qw422016.N().S(`/delete" data-message="Are you sure you wish to delete member [`)
//line views/vsprint/vsmember/Edit.html:25
		qw422016.E().S(p.Model.String())
//line views/vsprint/vsmember/Edit.html:25
		qw422016.N().S(`]?"><button>`)
//line views/vsprint/vsmember/Edit.html:25
		components.StreamSVGButton(qw422016, "times", ps)
//line views/vsprint/vsmember/Edit.html:25
		qw422016.N().S(` Delete</button></a></div>
    <h3>`)
//line views/vsprint/vsmember/Edit.html:26
		components.StreamSVGIcon(qw422016, `users`, ps)
//line views/vsprint/vsmember/Edit.html:26
		qw422016.N().S(` Edit Member [`)
//line views/vsprint/vsmember/Edit.html:26
		qw422016.E().S(p.Model.String())
//line views/vsprint/vsmember/Edit.html:26
		qw422016.N().S(`]</h3>
`)
//line views/vsprint/vsmember/Edit.html:27
	}
//line views/vsprint/vsmember/Edit.html:27
	qw422016.N().S(`    <form action="`)
//line views/vsprint/vsmember/Edit.html:28
	qw422016.E().S(util.Choose(p.IsNew, `/admin/db/sprint/member/_new`, ``))
//line views/vsprint/vsmember/Edit.html:28
	qw422016.N().S(`" class="mt" method="post">
      <table class="mt expanded">
        <tbody>
          `)
//line views/vsprint/vsmember/Edit.html:31
	if p.IsNew {
//line views/vsprint/vsmember/Edit.html:31
		edit.StreamUUIDTable(qw422016, "sprintID", "input-sprintID", "Sprint ID", &p.Model.SprintID, 5, "UUID in format (00000000-0000-0000-0000-000000000000)")
//line views/vsprint/vsmember/Edit.html:31
	}
//line views/vsprint/vsmember/Edit.html:31
	qw422016.N().S(`
          `)
//line views/vsprint/vsmember/Edit.html:32
	if p.IsNew {
//line views/vsprint/vsmember/Edit.html:32
		edit.StreamUUIDTable(qw422016, "userID", "input-userID", "User ID", &p.Model.UserID, 5, "UUID in format (00000000-0000-0000-0000-000000000000)")
//line views/vsprint/vsmember/Edit.html:32
	}
//line views/vsprint/vsmember/Edit.html:32
	qw422016.N().S(`
          `)
//line views/vsprint/vsmember/Edit.html:33
	edit.StreamStringTable(qw422016, "name", "", "Name", p.Model.Name, 5, "String text")
//line views/vsprint/vsmember/Edit.html:33
	qw422016.N().S(`
          `)
//line views/vsprint/vsmember/Edit.html:34
	edit.StreamStringTable(qw422016, "picture", "", "Picture", p.Model.Picture, 5, "URL in string form")
//line views/vsprint/vsmember/Edit.html:34
	qw422016.N().S(`
          `)
//line views/vsprint/vsmember/Edit.html:35
	edit.StreamSelectTable(qw422016, "role", "", "Role", p.Model.Role.Key, enum.AllMemberStatuses.Keys(), enum.AllMemberStatuses.Strings(), 5, enum.AllMemberStatuses.Help())
//line views/vsprint/vsmember/Edit.html:35
	qw422016.N().S(`
          <tr><td colspan="2"><button type="submit">Save Changes</button></td></tr>
        </tbody>
      </table>
    </form>
  </div>
  <script>
    document.addEventListener("DOMContentLoaded", function() {
      rituals.autocomplete(document.getElementById("input-sprintID"), "/admin/db/sprint?sprint.l=10", "q", (o) => o["slug"] + " / " + o["title"] + " (" + o["id"] + ")", (o) => o["id"]);
      rituals.autocomplete(document.getElementById("input-userID"), "/admin/db/user?user.l=10", "q", (o) => o["name"] + " (" + o["id"] + ")", (o) => o["id"]);
    });
  </script>
`)
//line views/vsprint/vsmember/Edit.html:47
}

//line views/vsprint/vsmember/Edit.html:47
func (p *Edit) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsprint/vsmember/Edit.html:47
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vsprint/vsmember/Edit.html:47
	p.StreamBody(qw422016, as, ps)
//line views/vsprint/vsmember/Edit.html:47
	qt422016.ReleaseWriter(qw422016)
//line views/vsprint/vsmember/Edit.html:47
}

//line views/vsprint/vsmember/Edit.html:47
func (p *Edit) Body(as *app.State, ps *cutil.PageState) string {
//line views/vsprint/vsmember/Edit.html:47
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vsprint/vsmember/Edit.html:47
	p.WriteBody(qb422016, as, ps)
//line views/vsprint/vsmember/Edit.html:47
	qs422016 := string(qb422016.B)
//line views/vsprint/vsmember/Edit.html:47
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vsprint/vsmember/Edit.html:47
	return qs422016
//line views/vsprint/vsmember/Edit.html:47
}
