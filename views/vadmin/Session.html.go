// Code generated by qtc from "Session.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vadmin/Session.html:2
package vadmin

//line views/vadmin/Session.html:2
import (
	"fmt"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vadmin/Session.html:14
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vadmin/Session.html:14
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vadmin/Session.html:14
type Session struct{ layout.Basic }

//line views/vadmin/Session.html:16
func (p *Session) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Session.html:16
	qw422016.N().S(`
  <div class="card">
    <h3>Session</h3>
    <em>`)
//line views/vadmin/Session.html:19
	qw422016.N().D(len(ps.Session))
//line views/vadmin/Session.html:19
	qw422016.N().S(` values</em>
  </div>
`)
//line views/vadmin/Session.html:21
	if len(ps.Session) > 0 {
//line views/vadmin/Session.html:21
		qw422016.N().S(`  <div class="card">
    <h3>Values</h3>
    <table class="mt expanded">
      <tbody>
`)
//line views/vadmin/Session.html:27
		keys := maps.Keys(ps.Session)
		slices.Sort(keys)

//line views/vadmin/Session.html:30
		for _, k := range keys {
//line views/vadmin/Session.html:31
			v := ps.Session[k]

//line views/vadmin/Session.html:31
			qw422016.N().S(`          <tr>
            <th class="shrink">`)
//line views/vadmin/Session.html:33
			qw422016.E().S(k)
//line views/vadmin/Session.html:33
			qw422016.N().S(`</th>
            <td>`)
//line views/vadmin/Session.html:34
			qw422016.E().S(fmt.Sprint(v))
//line views/vadmin/Session.html:34
			qw422016.N().S(`</td>
          </tr>
`)
//line views/vadmin/Session.html:36
		}
//line views/vadmin/Session.html:36
		qw422016.N().S(`      </tbody>
    </table>
  </div>
`)
//line views/vadmin/Session.html:40
	} else {
//line views/vadmin/Session.html:40
		qw422016.N().S(`  <div class="card">
    <em>Empty session</em>
  </div>
`)
//line views/vadmin/Session.html:44
	}
//line views/vadmin/Session.html:44
	qw422016.N().S(`  <div class="card">
    <h3>Profile</h3>
    <div class="mt">`)
//line views/vadmin/Session.html:47
	components.StreamJSON(qw422016, ps.Profile)
//line views/vadmin/Session.html:47
	qw422016.N().S(`</div>
  </div>
`)
//line views/vadmin/Session.html:49
	if len(ps.Accounts) > 0 {
//line views/vadmin/Session.html:49
		qw422016.N().S(`  <div class="card">
    <h3>Accounts</h3>
    <table class="mt">
      <thead>
        <tr>
          <th>Provider</th>
          <th>Email</th>
          <th>Token</th>
        </tr>
      </thead>
      <tbody>
`)
//line views/vadmin/Session.html:61
		for _, acct := range ps.Accounts {
//line views/vadmin/Session.html:61
			qw422016.N().S(`      <tr>
        <td>`)
//line views/vadmin/Session.html:63
			qw422016.E().S(acct.Provider)
//line views/vadmin/Session.html:63
			qw422016.N().S(`</td>
        <td>`)
//line views/vadmin/Session.html:64
			qw422016.E().S(acct.Email)
//line views/vadmin/Session.html:64
			qw422016.N().S(`</td>
        <td><div class="break-word">`)
//line views/vadmin/Session.html:65
			qw422016.E().S(acct.Token)
//line views/vadmin/Session.html:65
			qw422016.N().S(`</div></td>
      </tr>
`)
//line views/vadmin/Session.html:67
		}
//line views/vadmin/Session.html:67
		qw422016.N().S(`      </tbody>
    </table>
  </div>
`)
//line views/vadmin/Session.html:71
	}
//line views/vadmin/Session.html:72
}

//line views/vadmin/Session.html:72
func (p *Session) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Session.html:72
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vadmin/Session.html:72
	p.StreamBody(qw422016, as, ps)
//line views/vadmin/Session.html:72
	qt422016.ReleaseWriter(qw422016)
//line views/vadmin/Session.html:72
}

//line views/vadmin/Session.html:72
func (p *Session) Body(as *app.State, ps *cutil.PageState) string {
//line views/vadmin/Session.html:72
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vadmin/Session.html:72
	p.WriteBody(qb422016, as, ps)
//line views/vadmin/Session.html:72
	qs422016 := string(qb422016.B)
//line views/vadmin/Session.html:72
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vadmin/Session.html:72
	return qs422016
//line views/vadmin/Session.html:72
}
