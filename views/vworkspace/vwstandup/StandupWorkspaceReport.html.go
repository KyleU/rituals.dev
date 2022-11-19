// Code generated by qtc from "StandupWorkspaceReport.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:1
package vwstandup

//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:1
import (
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/standup/report"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/app/workspace"
	"github.com/kyleu/rituals/views/components"
)

//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:9
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:9
func StreamStandupWorkspaceReports(qw422016 *qt422016.Writer, w *workspace.FullStandup, ps *cutil.PageState) {
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:9
	qw422016.N().S(`
  <ul class="accordion">
`)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:11
	for _, g := range w.Reports.Grouped() {
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:11
		qw422016.N().S(`    <li>
      <input id="accordion-`)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:13
		qw422016.E().S(util.TimeToYMD(&g.Day))
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:13
		qw422016.N().S(`" type="checkbox" hidden checked="checked" />
      <label for="accordion-`)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:14
		qw422016.E().S(util.TimeToYMD(&g.Day))
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:14
		qw422016.N().S(`">`)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:14
		components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:14
		qw422016.N().S(` `)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:14
		qw422016.E().S(util.TimeToYMD(&g.Day))
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:14
		qw422016.N().S(`</label>
      <div class="bd">
`)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:16
		for _, r := range g.Reports {
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:16
			qw422016.N().S(`        <a class="report" href="#modal-report-`)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:17
			qw422016.E().S(r.ID.String())
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:17
			qw422016.N().S(`" style="text-decoration: none; color: var(--color-foreground);">
          <div>
`)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:20
			uname := "Former Guest"
			curr := w.Members.Get(w.Standup.ID, r.UserID)
			if curr != nil {
				uname = curr.Name
			}

//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:25
			qw422016.N().S(`            <h4>`)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:26
			qw422016.E().S(uname)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:26
			qw422016.N().S(`</h4>
            <div>`)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:27
			qw422016.N().S(r.HTML)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:27
			qw422016.N().S(`</div>
          </div>
        </a>
`)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:30
		}
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:30
		qw422016.N().S(`      </div>
    </li>
`)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:33
	}
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:33
	qw422016.N().S(`  </ul>
`)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:35
	for _, r := range w.Reports {
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:36
		if ps.Profile.ID == r.UserID {
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:36
			qw422016.N().S(`  `)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:37
			StreamStandupWorkspaceReportModalEdit(qw422016, r)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:37
			qw422016.N().S(`
`)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:38
		} else {
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:38
			qw422016.N().S(`  `)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:39
			StreamStandupWorkspaceReportModalView(qw422016, r)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:39
			qw422016.N().S(`
`)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:40
		}
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:41
	}
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:42
}

//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:42
func WriteStandupWorkspaceReports(qq422016 qtio422016.Writer, w *workspace.FullStandup, ps *cutil.PageState) {
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:42
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:42
	StreamStandupWorkspaceReports(qw422016, w, ps)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:42
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:42
}

//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:42
func StandupWorkspaceReports(w *workspace.FullStandup, ps *cutil.PageState) string {
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:42
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:42
	WriteStandupWorkspaceReports(qb422016, w, ps)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:42
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:42
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:42
	return qs422016
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:42
}

//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:44
func StreamStandupWorkspaceReportModalAdd(qw422016 *qt422016.Writer) {
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:44
	qw422016.N().S(`
  <div id="modal-report--add" class="modal" style="display: none;">
    <a class="backdrop" href="#"></a>
    <div class="modal-content">
      <div class="modal-header">
        <a href="#" class="modal-close">×</a>
        <h2>New Report</h2>
      </div>
      <div class="modal-body">
        <form action="" method="post">
          <input type="hidden" name="action" value="report-add" />
          <table class="mt expanded">
            <tbody>
            `)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:57
	components.StreamFormVerticalInputTimestampDay(qw422016, "day", "Day", util.TimeToday(), 5)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:57
	qw422016.N().S(`
            `)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:58
	components.StreamFormVerticalTextarea(qw422016, "content", "Content", 8, "", 5, "HTML and Markdown supported")
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:58
	qw422016.N().S(`
            <tr><td colspan="2"><button type="submit">Add Report</button></td></tr>
            </tbody>
          </table>
        </form>
      </div>
    </div>
  </div>
`)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:66
}

//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:66
func WriteStandupWorkspaceReportModalAdd(qq422016 qtio422016.Writer) {
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:66
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:66
	StreamStandupWorkspaceReportModalAdd(qw422016)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:66
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:66
}

//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:66
func StandupWorkspaceReportModalAdd() string {
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:66
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:66
	WriteStandupWorkspaceReportModalAdd(qb422016)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:66
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:66
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:66
	return qs422016
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:66
}

//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:68
func StreamStandupWorkspaceReportModalEdit(qw422016 *qt422016.Writer, r *report.Report) {
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:68
	qw422016.N().S(`
  <div id="modal-report-`)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:69
	qw422016.E().S(r.ID.String())
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:69
	qw422016.N().S(`" class="modal" style="display: none;">
    <a class="backdrop" href="#"></a>
    <div class="modal-content">
      <div class="modal-header">
        <a href="#" class="modal-close">×</a>
        <h2>`)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:74
	qw422016.E().S(r.TitleString())
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:74
	qw422016.N().S(`</h2>
      </div>
      <div class="modal-body">
        <form action="" method="post">
          <input type="hidden" name="action" value="report-edit" />
          <input type="hidden" name="reportID" value="`)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:79
	qw422016.E().S(r.ID.String())
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:79
	qw422016.N().S(`" />
          <table class="mt expanded">
            <tbody>
            `)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:82
	components.StreamFormVerticalInputTimestampDay(qw422016, "day", "Day", &r.Day, 5)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:82
	qw422016.N().S(`
            `)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:83
	components.StreamFormVerticalTextarea(qw422016, "content", "Content", 8, r.Content, 5, "HTML and Markdown supported")
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:83
	qw422016.N().S(`
            <tr><td colspan="2"><button type="submit">Save Changes</button></td></tr>
            </tbody>
          </table>
        </form>
      </div>
    </div>
  </div>
`)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:91
}

//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:91
func WriteStandupWorkspaceReportModalEdit(qq422016 qtio422016.Writer, r *report.Report) {
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:91
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:91
	StreamStandupWorkspaceReportModalEdit(qw422016, r)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:91
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:91
}

//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:91
func StandupWorkspaceReportModalEdit(r *report.Report) string {
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:91
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:91
	WriteStandupWorkspaceReportModalEdit(qb422016, r)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:91
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:91
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:91
	return qs422016
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:91
}

//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:93
func StreamStandupWorkspaceReportModalView(qw422016 *qt422016.Writer, r *report.Report) {
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:93
	qw422016.N().S(`
  <div id="modal-report-`)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:94
	qw422016.E().S(r.ID.String())
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:94
	qw422016.N().S(`" class="modal" style="display: none;">
    <a class="backdrop" href="#"></a>
    <div class="modal-content">
      <div class="modal-header">
        <a href="#" class="modal-close">×</a>
        <h2>`)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:99
	qw422016.E().S(r.TitleString())
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:99
	qw422016.N().S(`</h2>
      </div>
      <div class="modal-body">
        `)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:102
	qw422016.N().S(r.HTML)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:102
	qw422016.N().S(`
      </div>
    </div>
  </div>
`)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:106
}

//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:106
func WriteStandupWorkspaceReportModalView(qq422016 qtio422016.Writer, r *report.Report) {
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:106
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:106
	StreamStandupWorkspaceReportModalView(qw422016, r)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:106
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:106
}

//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:106
func StandupWorkspaceReportModalView(r *report.Report) string {
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:106
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:106
	WriteStandupWorkspaceReportModalView(qb422016, r)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:106
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:106
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:106
	return qs422016
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:106
}
