// Code generated by qtc from "About.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/About.html:2
package views

//line views/About.html:2
import (
	"time"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/components/view"
	"github.com/kyleu/rituals/views/layout"
)

//line views/About.html:13
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/About.html:13
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/About.html:13
type About struct {
	layout.Basic
	Version string
	Started time.Time
}

//line views/About.html:19
func (p *About) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/About.html:19
	qw422016.N().S(`
  <div class="card">
    <h3>`)
//line views/About.html:21
	components.StreamSVGIcon(qw422016, `app`, ps)
//line views/About.html:21
	qw422016.E().S(util.AppName)
//line views/About.html:21
	qw422016.N().S(`</h3>
    <em>v`)
//line views/About.html:22
	qw422016.E().S(p.Version)
//line views/About.html:22
	qw422016.N().S(`, started `)
//line views/About.html:22
	view.StreamTimestampRelative(qw422016, &p.Started, false)
//line views/About.html:22
	qw422016.N().S(`</em>
  </div>
  <div class="card">
    <h3>About</h3>
`)
//line views/About.html:26
	qw422016.N().S(`    <p>
      This app, <a href="https://rituals.dev">rituals.dev</a>, allows you to collaborate with your team to manage your work.
      It provides teams, sprints, estimate sessions, standup meetings, and retrospectives.
    </p>
    <p>
      Collaborative editing is supported, using a websocket to provide lightning-fast updates.
      If preferred, the app is fully functional without JavaScript.
    </p>
`)
//line views/About.html:35
	qw422016.N().S(`  </div>
  `)
//line views/About.html:37
	StreamSourceCode(qw422016)
//line views/About.html:37
	qw422016.N().S(`
  `)
//line views/About.html:38
	StreamFeedback(qw422016)
//line views/About.html:38
	qw422016.N().S(`
`)
//line views/About.html:39
}

//line views/About.html:39
func (p *About) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/About.html:39
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/About.html:39
	p.StreamBody(qw422016, as, ps)
//line views/About.html:39
	qt422016.ReleaseWriter(qw422016)
//line views/About.html:39
}

//line views/About.html:39
func (p *About) Body(as *app.State, ps *cutil.PageState) string {
//line views/About.html:39
	qb422016 := qt422016.AcquireByteBuffer()
//line views/About.html:39
	p.WriteBody(qb422016, as, ps)
//line views/About.html:39
	qs422016 := string(qb422016.B)
//line views/About.html:39
	qt422016.ReleaseByteBuffer(qb422016)
//line views/About.html:39
	return qs422016
//line views/About.html:39
}

//line views/About.html:41
func StreamSourceCode(qw422016 *qt422016.Writer) {
//line views/About.html:41
	qw422016.N().S(`
  <div class="card">
    <h3>Source Code</h3>
    <p>The project is available on <a href="https://github.com/kyleu/rituals" target="_blank" rel="noopener noreferrer">GitHub</a></p>
  </div>
`)
//line views/About.html:46
}

//line views/About.html:46
func WriteSourceCode(qq422016 qtio422016.Writer) {
//line views/About.html:46
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/About.html:46
	StreamSourceCode(qw422016)
//line views/About.html:46
	qt422016.ReleaseWriter(qw422016)
//line views/About.html:46
}

//line views/About.html:46
func SourceCode() string {
//line views/About.html:46
	qb422016 := qt422016.AcquireByteBuffer()
//line views/About.html:46
	WriteSourceCode(qb422016)
//line views/About.html:46
	qs422016 := string(qb422016.B)
//line views/About.html:46
	qt422016.ReleaseByteBuffer(qb422016)
//line views/About.html:46
	return qs422016
//line views/About.html:46
}

//line views/About.html:48
func StreamFeedback(qw422016 *qt422016.Writer) {
//line views/About.html:48
	qw422016.N().S(`
  <div class="card">
    <h3>Feedback</h3>
    <p>For now, email <a href="mailto:rituals.dev@kyleu.com">Kyle U</a></p>
  </div>
`)
//line views/About.html:53
}

//line views/About.html:53
func WriteFeedback(qq422016 qtio422016.Writer) {
//line views/About.html:53
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/About.html:53
	StreamFeedback(qw422016)
//line views/About.html:53
	qt422016.ReleaseWriter(qw422016)
//line views/About.html:53
}

//line views/About.html:53
func Feedback() string {
//line views/About.html:53
	qb422016 := qt422016.AcquireByteBuffer()
//line views/About.html:53
	WriteFeedback(qb422016)
//line views/About.html:53
	qs422016 := string(qb422016.B)
//line views/About.html:53
	qt422016.ReleaseByteBuffer(qb422016)
//line views/About.html:53
	return qs422016
//line views/About.html:53
}
