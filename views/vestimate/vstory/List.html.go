// Code generated by qtc from "List.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vestimate/vstory/List.html:2
package vstory

//line views/vestimate/vstory/List.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/estimate"
	"github.com/kyleu/rituals/app/estimate/story"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vestimate/vstory/List.html:13
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vestimate/vstory/List.html:13
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vestimate/vstory/List.html:13
type List struct {
	layout.Basic
	Models                story.Stories
	EstimatesByEstimateID estimate.Estimates
	UsersByUserID         user.Users
	Params                filter.ParamSet
	SearchQuery           string
}

//line views/vestimate/vstory/List.html:22
func (p *List) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vestimate/vstory/List.html:22
	qw422016.N().S(`
  <div class="card">
    <div class="right">`)
//line views/vestimate/vstory/List.html:24
	components.StreamSearchForm(qw422016, "", "q", "Search stories", p.SearchQuery, ps)
//line views/vestimate/vstory/List.html:24
	qw422016.N().S(`</div>
    <h3>`)
//line views/vestimate/vstory/List.html:25
	components.StreamSVGRefIcon(qw422016, `story`, ps)
//line views/vestimate/vstory/List.html:25
	qw422016.E().S(ps.Title)
//line views/vestimate/vstory/List.html:25
	qw422016.N().S(` <a href="/admin/db/estimate/story/new"><button>New</button></a></h3>
    <div class="clear"></div>
`)
//line views/vestimate/vstory/List.html:27
	if p.SearchQuery != "" {
//line views/vestimate/vstory/List.html:27
		qw422016.N().S(`    <em>Search results for [`)
//line views/vestimate/vstory/List.html:28
		qw422016.E().S(p.SearchQuery)
//line views/vestimate/vstory/List.html:28
		qw422016.N().S(`]</em>
`)
//line views/vestimate/vstory/List.html:29
	}
//line views/vestimate/vstory/List.html:30
	if len(p.Models) == 0 {
//line views/vestimate/vstory/List.html:30
		qw422016.N().S(`    <div class="mt"><em>No stories available</em></div>
`)
//line views/vestimate/vstory/List.html:32
	} else {
//line views/vestimate/vstory/List.html:32
		qw422016.N().S(`    <div class="overflow clear">
      `)
//line views/vestimate/vstory/List.html:34
		StreamTable(qw422016, p.Models, p.EstimatesByEstimateID, p.UsersByUserID, p.Params, as, ps)
//line views/vestimate/vstory/List.html:34
		qw422016.N().S(`
    </div>
`)
//line views/vestimate/vstory/List.html:36
	}
//line views/vestimate/vstory/List.html:36
	qw422016.N().S(`  </div>
`)
//line views/vestimate/vstory/List.html:38
}

//line views/vestimate/vstory/List.html:38
func (p *List) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vestimate/vstory/List.html:38
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vestimate/vstory/List.html:38
	p.StreamBody(qw422016, as, ps)
//line views/vestimate/vstory/List.html:38
	qt422016.ReleaseWriter(qw422016)
//line views/vestimate/vstory/List.html:38
}

//line views/vestimate/vstory/List.html:38
func (p *List) Body(as *app.State, ps *cutil.PageState) string {
//line views/vestimate/vstory/List.html:38
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vestimate/vstory/List.html:38
	p.WriteBody(qb422016, as, ps)
//line views/vestimate/vstory/List.html:38
	qs422016 := string(qb422016.B)
//line views/vestimate/vstory/List.html:38
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vestimate/vstory/List.html:38
	return qs422016
//line views/vestimate/vstory/List.html:38
}
