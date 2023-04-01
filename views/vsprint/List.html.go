// Code generated by qtc from "List.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vsprint/List.html:2
package vsprint

//line views/vsprint/List.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vsprint/List.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vsprint/List.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vsprint/List.html:12
type List struct {
	layout.Basic
	Models        sprint.Sprints
	TeamsByTeamID team.Teams
	Params        filter.ParamSet
	SearchQuery   string
}

//line views/vsprint/List.html:20
func (p *List) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsprint/List.html:20
	qw422016.N().S(`
  <div class="card">
    <div class="right">`)
//line views/vsprint/List.html:22
	components.StreamSearchForm(qw422016, "", "q", "Search sprints", p.SearchQuery, ps)
//line views/vsprint/List.html:22
	qw422016.N().S(`</div>
    <h3>`)
//line views/vsprint/List.html:23
	components.StreamSVGRefIcon(qw422016, `sprint`, ps)
//line views/vsprint/List.html:23
	qw422016.E().S(ps.Title)
//line views/vsprint/List.html:23
	qw422016.N().S(` <a href="/admin/db/sprint/new"><button>New</button></a></h3>
    <div class="clear"></div>
`)
//line views/vsprint/List.html:25
	if p.SearchQuery != "" {
//line views/vsprint/List.html:25
		qw422016.N().S(`    <em>Search results for [`)
//line views/vsprint/List.html:26
		qw422016.E().S(p.SearchQuery)
//line views/vsprint/List.html:26
		qw422016.N().S(`]</em>
`)
//line views/vsprint/List.html:27
	}
//line views/vsprint/List.html:28
	if len(p.Models) == 0 {
//line views/vsprint/List.html:28
		qw422016.N().S(`    <div class="mt"><em>No sprints available</em></div>
`)
//line views/vsprint/List.html:30
	} else {
//line views/vsprint/List.html:30
		qw422016.N().S(`    <div class="overflow clear">
      `)
//line views/vsprint/List.html:32
		StreamTable(qw422016, p.Models, p.TeamsByTeamID, p.Params, as, ps)
//line views/vsprint/List.html:32
		qw422016.N().S(`
    </div>
`)
//line views/vsprint/List.html:34
	}
//line views/vsprint/List.html:34
	qw422016.N().S(`  </div>
`)
//line views/vsprint/List.html:36
}

//line views/vsprint/List.html:36
func (p *List) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsprint/List.html:36
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vsprint/List.html:36
	p.StreamBody(qw422016, as, ps)
//line views/vsprint/List.html:36
	qt422016.ReleaseWriter(qw422016)
//line views/vsprint/List.html:36
}

//line views/vsprint/List.html:36
func (p *List) Body(as *app.State, ps *cutil.PageState) string {
//line views/vsprint/List.html:36
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vsprint/List.html:36
	p.WriteBody(qb422016, as, ps)
//line views/vsprint/List.html:36
	qs422016 := string(qb422016.B)
//line views/vsprint/List.html:36
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vsprint/List.html:36
	return qs422016
//line views/vsprint/List.html:36
}
