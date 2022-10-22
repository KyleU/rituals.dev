// Code generated by qtc from "Edit.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vestimate/vstory/Edit.html:2
package vstory

//line views/vestimate/vstory/Edit.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/estimate/story"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vestimate/vstory/Edit.html:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vestimate/vstory/Edit.html:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vestimate/vstory/Edit.html:10
type Edit struct {
	layout.Basic
	Model *story.Story
	IsNew bool
}

//line views/vestimate/vstory/Edit.html:16
func (p *Edit) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vestimate/vstory/Edit.html:16
	qw422016.N().S(`
  <div class="card">
`)
//line views/vestimate/vstory/Edit.html:18
	if p.IsNew {
//line views/vestimate/vstory/Edit.html:18
		qw422016.N().S(`    <div class="right"><a href="/admin/db/estimate/story/random"><button>Random</button></a></div>
    <h3>`)
//line views/vestimate/vstory/Edit.html:20
		components.StreamSVGRefIcon(qw422016, `star`, ps)
//line views/vestimate/vstory/Edit.html:20
		qw422016.N().S(` New Story</h3>
    <form action="/admin/db/estimate/story/new" class="mt" method="post">
`)
//line views/vestimate/vstory/Edit.html:22
	} else {
//line views/vestimate/vstory/Edit.html:22
		qw422016.N().S(`    <div class="right"><a href="`)
//line views/vestimate/vstory/Edit.html:23
		qw422016.E().S(p.Model.WebPath())
//line views/vestimate/vstory/Edit.html:23
		qw422016.N().S(`/delete" onclick="return confirm('Are you sure you wish to delete story [`)
//line views/vestimate/vstory/Edit.html:23
		qw422016.E().S(p.Model.String())
//line views/vestimate/vstory/Edit.html:23
		qw422016.N().S(`]?')"><button>Delete</button></a></div>
    <h3>`)
//line views/vestimate/vstory/Edit.html:24
		components.StreamSVGRefIcon(qw422016, `star`, ps)
//line views/vestimate/vstory/Edit.html:24
		qw422016.N().S(` Edit Story [`)
//line views/vestimate/vstory/Edit.html:24
		qw422016.E().S(p.Model.String())
//line views/vestimate/vstory/Edit.html:24
		qw422016.N().S(`]</h3>
`)
//line views/vestimate/vstory/Edit.html:25
	}
//line views/vestimate/vstory/Edit.html:25
	qw422016.N().S(`    <form action="" method="post">
      <table class="mt expanded">
        <tbody>
          `)
//line views/vestimate/vstory/Edit.html:29
	if p.IsNew {
//line views/vestimate/vstory/Edit.html:29
		components.StreamTableInputUUID(qw422016, "id", "ID", &p.Model.ID, 5, "UUID in format (00000000-0000-0000-0000-000000000000)")
//line views/vestimate/vstory/Edit.html:29
	}
//line views/vestimate/vstory/Edit.html:29
	qw422016.N().S(`
          `)
//line views/vestimate/vstory/Edit.html:30
	components.StreamTableInputUUID(qw422016, "estimateID", "Estimate ID", &p.Model.EstimateID, 5, "UUID in format (00000000-0000-0000-0000-000000000000)")
//line views/vestimate/vstory/Edit.html:30
	qw422016.N().S(`
          `)
//line views/vestimate/vstory/Edit.html:31
	components.StreamTableInputNumber(qw422016, "idx", "Idx", p.Model.Idx, 5, "Integer")
//line views/vestimate/vstory/Edit.html:31
	qw422016.N().S(`
          `)
//line views/vestimate/vstory/Edit.html:32
	components.StreamTableInputUUID(qw422016, "userID", "User ID", &p.Model.UserID, 5, "UUID in format (00000000-0000-0000-0000-000000000000)")
//line views/vestimate/vstory/Edit.html:32
	qw422016.N().S(`
          `)
//line views/vestimate/vstory/Edit.html:33
	components.StreamTableInput(qw422016, "title", "Title", p.Model.Title, 5, "String text")
//line views/vestimate/vstory/Edit.html:33
	qw422016.N().S(`
          `)
//line views/vestimate/vstory/Edit.html:34
	components.StreamTableInput(qw422016, "status", "Status", string(p.Model.Status), 5, "Available options: [new, active, complete, deleted]")
//line views/vestimate/vstory/Edit.html:34
	qw422016.N().S(`
          `)
//line views/vestimate/vstory/Edit.html:35
	components.StreamTableInput(qw422016, "finalVote", "Final Vote", p.Model.FinalVote, 5, "String text")
//line views/vestimate/vstory/Edit.html:35
	qw422016.N().S(`
          <tr><td colspan="2"><button type="submit">Save Changes</button></td></tr>
        </tbody>
      </table>
    </form>
  </div>
`)
//line views/vestimate/vstory/Edit.html:41
}

//line views/vestimate/vstory/Edit.html:41
func (p *Edit) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vestimate/vstory/Edit.html:41
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vestimate/vstory/Edit.html:41
	p.StreamBody(qw422016, as, ps)
//line views/vestimate/vstory/Edit.html:41
	qt422016.ReleaseWriter(qw422016)
//line views/vestimate/vstory/Edit.html:41
}

//line views/vestimate/vstory/Edit.html:41
func (p *Edit) Body(as *app.State, ps *cutil.PageState) string {
//line views/vestimate/vstory/Edit.html:41
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vestimate/vstory/Edit.html:41
	p.WriteBody(qb422016, as, ps)
//line views/vestimate/vstory/Edit.html:41
	qs422016 := string(qb422016.B)
//line views/vestimate/vstory/Edit.html:41
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vestimate/vstory/Edit.html:41
	return qs422016
//line views/vestimate/vstory/Edit.html:41
}
