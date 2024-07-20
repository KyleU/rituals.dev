// Code generated by qtc from "List.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vestimate/vstory/vvote/List.html:1
package vvote

//line views/vestimate/vstory/vvote/List.html:1
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/estimate/story"
	"github.com/kyleu/rituals/app/estimate/story/vote"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vestimate/vstory/vvote/List.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vestimate/vstory/vvote/List.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vestimate/vstory/vvote/List.html:12
type List struct {
	layout.Basic
	Models           vote.Votes
	StoriesByStoryID story.Stories
	UsersByUserID    user.Users
	Params           filter.ParamSet
	Paths            []string
}

//line views/vestimate/vstory/vvote/List.html:21
func (p *List) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vestimate/vstory/vvote/List.html:21
	qw422016.N().S(`
  <div class="card">
    <div class="right mrs large-buttons">
`)
//line views/vestimate/vstory/vvote/List.html:24
	if len(p.Models) > 1 {
//line views/vestimate/vstory/vvote/List.html:24
		qw422016.N().S(`<a href="/admin/db/estimate/story/vote/_random"><button>`)
//line views/vestimate/vstory/vvote/List.html:24
		components.StreamSVGButton(qw422016, "gift", ps)
//line views/vestimate/vstory/vvote/List.html:24
		qw422016.N().S(` Random</button></a>`)
//line views/vestimate/vstory/vvote/List.html:24
	}
//line views/vestimate/vstory/vvote/List.html:24
	qw422016.N().S(`      <a href="`)
//line views/vestimate/vstory/vvote/List.html:25
	qw422016.E().S(vote.Route(p.Paths...))
//line views/vestimate/vstory/vvote/List.html:25
	qw422016.N().S(`/_new"><button>`)
//line views/vestimate/vstory/vvote/List.html:25
	components.StreamSVGButton(qw422016, "plus", ps)
//line views/vestimate/vstory/vvote/List.html:25
	qw422016.N().S(` New</button></a>
    </div>
    <h3>`)
//line views/vestimate/vstory/vvote/List.html:27
	components.StreamSVGIcon(qw422016, `vote-yea`, ps)
//line views/vestimate/vstory/vvote/List.html:27
	qw422016.N().S(` `)
//line views/vestimate/vstory/vvote/List.html:27
	qw422016.E().S(ps.Title)
//line views/vestimate/vstory/vvote/List.html:27
	qw422016.N().S(`</h3>
`)
//line views/vestimate/vstory/vvote/List.html:28
	if len(p.Models) == 0 {
//line views/vestimate/vstory/vvote/List.html:28
		qw422016.N().S(`    <div class="mt"><em>No votes available</em></div>
`)
//line views/vestimate/vstory/vvote/List.html:30
	} else {
//line views/vestimate/vstory/vvote/List.html:30
		qw422016.N().S(`    <div class="mt">
      `)
//line views/vestimate/vstory/vvote/List.html:32
		StreamTable(qw422016, p.Models, p.StoriesByStoryID, p.UsersByUserID, p.Params, as, ps, p.Paths...)
//line views/vestimate/vstory/vvote/List.html:32
		qw422016.N().S(`
    </div>
`)
//line views/vestimate/vstory/vvote/List.html:34
	}
//line views/vestimate/vstory/vvote/List.html:34
	qw422016.N().S(`  </div>
`)
//line views/vestimate/vstory/vvote/List.html:36
}

//line views/vestimate/vstory/vvote/List.html:36
func (p *List) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vestimate/vstory/vvote/List.html:36
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vestimate/vstory/vvote/List.html:36
	p.StreamBody(qw422016, as, ps)
//line views/vestimate/vstory/vvote/List.html:36
	qt422016.ReleaseWriter(qw422016)
//line views/vestimate/vstory/vvote/List.html:36
}

//line views/vestimate/vstory/vvote/List.html:36
func (p *List) Body(as *app.State, ps *cutil.PageState) string {
//line views/vestimate/vstory/vvote/List.html:36
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vestimate/vstory/vvote/List.html:36
	p.WriteBody(qb422016, as, ps)
//line views/vestimate/vstory/vvote/List.html:36
	qs422016 := string(qb422016.B)
//line views/vestimate/vstory/vvote/List.html:36
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vestimate/vstory/vvote/List.html:36
	return qs422016
//line views/vestimate/vstory/vvote/List.html:36
}
