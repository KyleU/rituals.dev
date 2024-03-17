// Code generated by qtc from "Edit.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vstandup/vuhistory/Edit.html:2
package vuhistory

//line views/vstandup/vuhistory/Edit.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/standup/uhistory"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/components/edit"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vstandup/vuhistory/Edit.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vstandup/vuhistory/Edit.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vstandup/vuhistory/Edit.html:11
type Edit struct {
	layout.Basic
	Model *uhistory.StandupHistory
	IsNew bool
}

//line views/vstandup/vuhistory/Edit.html:17
func (p *Edit) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vstandup/vuhistory/Edit.html:17
	qw422016.N().S(`
  <div class="card">
`)
//line views/vstandup/vuhistory/Edit.html:19
	if p.IsNew {
//line views/vstandup/vuhistory/Edit.html:19
		qw422016.N().S(`    <div class="right"><a href="?prototype=random"><button>Random</button></a></div>
    <h3>`)
//line views/vstandup/vuhistory/Edit.html:21
		components.StreamSVGRefIcon(qw422016, `history`, ps)
//line views/vstandup/vuhistory/Edit.html:21
		qw422016.N().S(` New History</h3>
    <form action="/admin/db/standup/history/_new" class="mt" method="post">
`)
//line views/vstandup/vuhistory/Edit.html:23
	} else {
//line views/vstandup/vuhistory/Edit.html:23
		qw422016.N().S(`    <div class="right"><a class="link-confirm" href="`)
//line views/vstandup/vuhistory/Edit.html:24
		qw422016.E().S(p.Model.WebPath())
//line views/vstandup/vuhistory/Edit.html:24
		qw422016.N().S(`/delete" data-message="Are you sure you wish to delete history [`)
//line views/vstandup/vuhistory/Edit.html:24
		qw422016.E().S(p.Model.String())
//line views/vstandup/vuhistory/Edit.html:24
		qw422016.N().S(`]?"><button>Delete</button></a></div>
    <h3>`)
//line views/vstandup/vuhistory/Edit.html:25
		components.StreamSVGRefIcon(qw422016, `history`, ps)
//line views/vstandup/vuhistory/Edit.html:25
		qw422016.N().S(` Edit History [`)
//line views/vstandup/vuhistory/Edit.html:25
		qw422016.E().S(p.Model.String())
//line views/vstandup/vuhistory/Edit.html:25
		qw422016.N().S(`]</h3>
    <form action="" method="post">
`)
//line views/vstandup/vuhistory/Edit.html:27
	}
//line views/vstandup/vuhistory/Edit.html:27
	qw422016.N().S(`      <table class="mt expanded">
        <tbody>
          `)
//line views/vstandup/vuhistory/Edit.html:30
	if p.IsNew {
//line views/vstandup/vuhistory/Edit.html:30
		edit.StreamStringTable(qw422016, "slug", "", "Slug", p.Model.Slug, 5, "String text")
//line views/vstandup/vuhistory/Edit.html:30
	}
//line views/vstandup/vuhistory/Edit.html:30
	qw422016.N().S(`
          `)
//line views/vstandup/vuhistory/Edit.html:31
	edit.StreamUUIDTable(qw422016, "standupID", "input-standupID", "Standup ID", &p.Model.StandupID, 5, "UUID in format (00000000-0000-0000-0000-000000000000)")
//line views/vstandup/vuhistory/Edit.html:31
	qw422016.N().S(`
          `)
//line views/vstandup/vuhistory/Edit.html:32
	edit.StreamStringTable(qw422016, "standupName", "", "Standup Name", p.Model.StandupName, 5, "String text")
//line views/vstandup/vuhistory/Edit.html:32
	qw422016.N().S(`
          <tr><td colspan="2"><button type="submit">Save Changes</button></td></tr>
        </tbody>
      </table>
    </form>
  </div>
  <script>
    document.addEventListener("DOMContentLoaded", function() {
      rituals.autocomplete(document.getElementById("input-standupID"), "/admin/db/standup?standup.l=10", "q", (o) => o["slug"] + " / " + o["title"] + " (" + o["id"] + ")", (o) => o["id"]);
    });
  </script>
`)
//line views/vstandup/vuhistory/Edit.html:43
}

//line views/vstandup/vuhistory/Edit.html:43
func (p *Edit) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vstandup/vuhistory/Edit.html:43
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vstandup/vuhistory/Edit.html:43
	p.StreamBody(qw422016, as, ps)
//line views/vstandup/vuhistory/Edit.html:43
	qt422016.ReleaseWriter(qw422016)
//line views/vstandup/vuhistory/Edit.html:43
}

//line views/vstandup/vuhistory/Edit.html:43
func (p *Edit) Body(as *app.State, ps *cutil.PageState) string {
//line views/vstandup/vuhistory/Edit.html:43
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vstandup/vuhistory/Edit.html:43
	p.WriteBody(qb422016, as, ps)
//line views/vstandup/vuhistory/Edit.html:43
	qs422016 := string(qb422016.B)
//line views/vstandup/vuhistory/Edit.html:43
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vstandup/vuhistory/Edit.html:43
	return qs422016
//line views/vstandup/vuhistory/Edit.html:43
}
