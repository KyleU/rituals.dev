// Code generated by qtc from "List.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vstandup/List.html:2
package vstandup

//line views/vstandup/List.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/standup"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/components/edit"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vstandup/List.html:14
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vstandup/List.html:14
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vstandup/List.html:14
type List struct {
	layout.Basic
	Models            standup.Standups
	TeamsByTeamID     team.Teams
	SprintsBySprintID sprint.Sprints
	Params            filter.ParamSet
	SearchQuery       string
}

//line views/vstandup/List.html:23
func (p *List) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vstandup/List.html:23
	qw422016.N().S(`
  <div class="card">
    <div class="right">`)
//line views/vstandup/List.html:25
	edit.StreamSearchForm(qw422016, "", "q", "Search Standups", p.SearchQuery, ps)
//line views/vstandup/List.html:25
	qw422016.N().S(`</div>
    <div class="right mrs large-buttons">
`)
//line views/vstandup/List.html:27
	if len(p.Models) > 0 {
//line views/vstandup/List.html:27
		qw422016.N().S(`<a href="/admin/db/standup/_random"><button>`)
//line views/vstandup/List.html:27
		components.StreamSVGButton(qw422016, "gift", ps)
//line views/vstandup/List.html:27
		qw422016.N().S(` Random</button></a>`)
//line views/vstandup/List.html:27
	}
//line views/vstandup/List.html:27
	qw422016.N().S(`      <a href="/admin/db/standup/_new"><button>`)
//line views/vstandup/List.html:28
	components.StreamSVGButton(qw422016, "plus", ps)
//line views/vstandup/List.html:28
	qw422016.N().S(` New</button></a>
    </div>
    <h3>`)
//line views/vstandup/List.html:30
	components.StreamSVGIcon(qw422016, `standup`, ps)
//line views/vstandup/List.html:30
	qw422016.N().S(` `)
//line views/vstandup/List.html:30
	qw422016.E().S(ps.Title)
//line views/vstandup/List.html:30
	qw422016.N().S(`</h3>
    <div class="clear"></div>
`)
//line views/vstandup/List.html:32
	if p.SearchQuery != "" {
//line views/vstandup/List.html:32
		qw422016.N().S(`    <hr />
    <em>Search results for [`)
//line views/vstandup/List.html:34
		qw422016.E().S(p.SearchQuery)
//line views/vstandup/List.html:34
		qw422016.N().S(`]</em> (<a href="?">clear</a>)
`)
//line views/vstandup/List.html:35
	}
//line views/vstandup/List.html:36
	if len(p.Models) == 0 {
//line views/vstandup/List.html:36
		qw422016.N().S(`    <div class="mt"><em>No standups available</em></div>
`)
//line views/vstandup/List.html:38
	} else {
//line views/vstandup/List.html:38
		qw422016.N().S(`    <div class="mt">
      `)
//line views/vstandup/List.html:40
		StreamTable(qw422016, p.Models, p.TeamsByTeamID, p.SprintsBySprintID, p.Params, as, ps)
//line views/vstandup/List.html:40
		qw422016.N().S(`
    </div>
`)
//line views/vstandup/List.html:42
	}
//line views/vstandup/List.html:42
	qw422016.N().S(`  </div>
`)
//line views/vstandup/List.html:44
}

//line views/vstandup/List.html:44
func (p *List) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vstandup/List.html:44
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vstandup/List.html:44
	p.StreamBody(qw422016, as, ps)
//line views/vstandup/List.html:44
	qt422016.ReleaseWriter(qw422016)
//line views/vstandup/List.html:44
}

//line views/vstandup/List.html:44
func (p *List) Body(as *app.State, ps *cutil.PageState) string {
//line views/vstandup/List.html:44
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vstandup/List.html:44
	p.WriteBody(qb422016, as, ps)
//line views/vstandup/List.html:44
	qs422016 := string(qb422016.B)
//line views/vstandup/List.html:44
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vstandup/List.html:44
	return qs422016
//line views/vstandup/List.html:44
}
