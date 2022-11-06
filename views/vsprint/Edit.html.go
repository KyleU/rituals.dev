// Code generated by qtc from "Edit.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vsprint/Edit.html:2
package vsprint

//line views/vsprint/Edit.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vsprint/Edit.html:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vsprint/Edit.html:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vsprint/Edit.html:10
type Edit struct {
	layout.Basic
	Model *sprint.Sprint
	IsNew bool
}

//line views/vsprint/Edit.html:16
func (p *Edit) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsprint/Edit.html:16
	qw422016.N().S(`
  <div class="card">
`)
//line views/vsprint/Edit.html:18
	if p.IsNew {
//line views/vsprint/Edit.html:18
		qw422016.N().S(`    <div class="right"><a href="/admin/db/sprint/random"><button>Random</button></a></div>
    <h3>`)
//line views/vsprint/Edit.html:20
		components.StreamSVGRefIcon(qw422016, `sprint`, ps)
//line views/vsprint/Edit.html:20
		qw422016.N().S(` New Sprint</h3>
    <form action="/admin/db/sprint/new" class="mt" method="post">
`)
//line views/vsprint/Edit.html:22
	} else {
//line views/vsprint/Edit.html:22
		qw422016.N().S(`    <div class="right"><a href="`)
//line views/vsprint/Edit.html:23
		qw422016.E().S(p.Model.WebPath())
//line views/vsprint/Edit.html:23
		qw422016.N().S(`/delete" onclick="return confirm('Are you sure you wish to delete sprint [`)
//line views/vsprint/Edit.html:23
		qw422016.E().S(p.Model.String())
//line views/vsprint/Edit.html:23
		qw422016.N().S(`]?')"><button>Delete</button></a></div>
    <h3>`)
//line views/vsprint/Edit.html:24
		components.StreamSVGRefIcon(qw422016, `sprint`, ps)
//line views/vsprint/Edit.html:24
		qw422016.N().S(` Edit Sprint [`)
//line views/vsprint/Edit.html:24
		qw422016.E().S(p.Model.String())
//line views/vsprint/Edit.html:24
		qw422016.N().S(`]</h3>
`)
//line views/vsprint/Edit.html:25
	}
//line views/vsprint/Edit.html:25
	qw422016.N().S(`    <form action="" method="post">
      <table class="mt expanded">
        <tbody>
          `)
//line views/vsprint/Edit.html:29
	if p.IsNew {
//line views/vsprint/Edit.html:29
		components.StreamTableInputUUID(qw422016, "id", "ID", &p.Model.ID, 5, "UUID in format (00000000-0000-0000-0000-000000000000)")
//line views/vsprint/Edit.html:29
	}
//line views/vsprint/Edit.html:29
	qw422016.N().S(`
          `)
//line views/vsprint/Edit.html:30
	components.StreamTableInput(qw422016, "slug", "Slug", p.Model.Slug, 5, "String text")
//line views/vsprint/Edit.html:30
	qw422016.N().S(`
          `)
//line views/vsprint/Edit.html:31
	components.StreamTableInput(qw422016, "title", "Title", p.Model.Title, 5, "String text")
//line views/vsprint/Edit.html:31
	qw422016.N().S(`
          `)
//line views/vsprint/Edit.html:32
	components.StreamTableSelect(qw422016, "status", "Status", string(p.Model.Status), []string{"new", "active", "complete", "deleted"}, []string{"new", "active", "complete", "deleted"}, 5, "Available options: [new, active, complete, deleted]")
//line views/vsprint/Edit.html:32
	qw422016.N().S(`
          `)
//line views/vsprint/Edit.html:33
	components.StreamTableInputUUID(qw422016, "teamID", "Team ID", p.Model.TeamID, 5, "UUID in format (00000000-0000-0000-0000-000000000000) (optional)")
//line views/vsprint/Edit.html:33
	qw422016.N().S(`
          `)
//line views/vsprint/Edit.html:34
	components.StreamTableInputUUID(qw422016, "owner", "Owner", &p.Model.Owner, 5, "UUID in format (00000000-0000-0000-0000-000000000000)")
//line views/vsprint/Edit.html:34
	qw422016.N().S(`
          `)
//line views/vsprint/Edit.html:35
	components.StreamTableInputTimestampDay(qw422016, "startDate", "Start Date", p.Model.StartDate, 5, "Calendar date (optional)")
//line views/vsprint/Edit.html:35
	qw422016.N().S(`
          `)
//line views/vsprint/Edit.html:36
	components.StreamTableInputTimestampDay(qw422016, "endDate", "End Date", p.Model.EndDate, 5, "Calendar date (optional)")
//line views/vsprint/Edit.html:36
	qw422016.N().S(`
          <tr><td colspan="2"><button type="submit">Save Changes</button></td></tr>
        </tbody>
      </table>
    </form>
  </div>
`)
//line views/vsprint/Edit.html:42
}

//line views/vsprint/Edit.html:42
func (p *Edit) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsprint/Edit.html:42
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vsprint/Edit.html:42
	p.StreamBody(qw422016, as, ps)
//line views/vsprint/Edit.html:42
	qt422016.ReleaseWriter(qw422016)
//line views/vsprint/Edit.html:42
}

//line views/vsprint/Edit.html:42
func (p *Edit) Body(as *app.State, ps *cutil.PageState) string {
//line views/vsprint/Edit.html:42
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vsprint/Edit.html:42
	p.WriteBody(qb422016, as, ps)
//line views/vsprint/Edit.html:42
	qs422016 := string(qb422016.B)
//line views/vsprint/Edit.html:42
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vsprint/Edit.html:42
	return qs422016
//line views/vsprint/Edit.html:42
}
