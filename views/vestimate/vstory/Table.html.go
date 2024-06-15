// Code generated by qtc from "Table.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vestimate/vstory/Table.html:2
package vstory

//line views/vestimate/vstory/Table.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/estimate"
	"github.com/kyleu/rituals/app/estimate/story"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/components/view"
)

//line views/vestimate/vstory/Table.html:14
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vestimate/vstory/Table.html:14
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vestimate/vstory/Table.html:14
func StreamTable(qw422016 *qt422016.Writer, models story.Stories, estimatesByEstimateID estimate.Estimates, usersByUserID user.Users, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vestimate/vstory/Table.html:14
	qw422016.N().S(`
`)
//line views/vestimate/vstory/Table.html:15
	prms := params.Sanitized("story", ps.Logger)

//line views/vestimate/vstory/Table.html:15
	qw422016.N().S(`  <div class="overflow clear">
    <table>
      <thead>
        <tr>
          `)
//line views/vestimate/vstory/Table.html:20
	components.StreamTableHeaderSimple(qw422016, "story", "id", "ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vestimate/vstory/Table.html:20
	qw422016.N().S(`
          `)
//line views/vestimate/vstory/Table.html:21
	components.StreamTableHeaderSimple(qw422016, "story", "estimate_id", "Estimate ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vestimate/vstory/Table.html:21
	qw422016.N().S(`
          `)
//line views/vestimate/vstory/Table.html:22
	components.StreamTableHeaderSimple(qw422016, "story", "idx", "Idx", "Integer", prms, ps.URI, ps)
//line views/vestimate/vstory/Table.html:22
	qw422016.N().S(`
          `)
//line views/vestimate/vstory/Table.html:23
	components.StreamTableHeaderSimple(qw422016, "story", "user_id", "User ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vestimate/vstory/Table.html:23
	qw422016.N().S(`
          `)
//line views/vestimate/vstory/Table.html:24
	components.StreamTableHeaderSimple(qw422016, "story", "title", "Title", "String text", prms, ps.URI, ps)
//line views/vestimate/vstory/Table.html:24
	qw422016.N().S(`
          `)
//line views/vestimate/vstory/Table.html:25
	components.StreamTableHeaderSimple(qw422016, "story", "status", "Status", enum.AllSessionStatuses.Help(), prms, ps.URI, ps)
//line views/vestimate/vstory/Table.html:25
	qw422016.N().S(`
          `)
//line views/vestimate/vstory/Table.html:26
	components.StreamTableHeaderSimple(qw422016, "story", "final_vote", "Final Vote", "String text", prms, ps.URI, ps)
//line views/vestimate/vstory/Table.html:26
	qw422016.N().S(`
          `)
//line views/vestimate/vstory/Table.html:27
	components.StreamTableHeaderSimple(qw422016, "story", "created", "Created", "Date and time, in almost any format", prms, ps.URI, ps)
//line views/vestimate/vstory/Table.html:27
	qw422016.N().S(`
          `)
//line views/vestimate/vstory/Table.html:28
	components.StreamTableHeaderSimple(qw422016, "story", "updated", "Updated", "Date and time, in almost any format (optional)", prms, ps.URI, ps)
//line views/vestimate/vstory/Table.html:28
	qw422016.N().S(`
        </tr>
      </thead>
      <tbody>
`)
//line views/vestimate/vstory/Table.html:32
	for _, model := range models {
//line views/vestimate/vstory/Table.html:32
		qw422016.N().S(`        <tr>
          <td><a href="/admin/db/estimate/story/`)
//line views/vestimate/vstory/Table.html:34
		view.StreamUUID(qw422016, &model.ID)
//line views/vestimate/vstory/Table.html:34
		qw422016.N().S(`">`)
//line views/vestimate/vstory/Table.html:34
		view.StreamUUID(qw422016, &model.ID)
//line views/vestimate/vstory/Table.html:34
		qw422016.N().S(`</a></td>
          <td class="nowrap">
            `)
//line views/vestimate/vstory/Table.html:36
		view.StreamUUID(qw422016, &model.EstimateID)
//line views/vestimate/vstory/Table.html:36
		if x := estimatesByEstimateID.Get(model.EstimateID); x != nil {
//line views/vestimate/vstory/Table.html:36
			qw422016.N().S(` (`)
//line views/vestimate/vstory/Table.html:36
			qw422016.E().S(x.TitleString())
//line views/vestimate/vstory/Table.html:36
			qw422016.N().S(`)`)
//line views/vestimate/vstory/Table.html:36
		}
//line views/vestimate/vstory/Table.html:36
		qw422016.N().S(`
            <a title="Estimate" href="`)
//line views/vestimate/vstory/Table.html:37
		qw422016.E().S(`/admin/db/estimate` + `/` + model.EstimateID.String())
//line views/vestimate/vstory/Table.html:37
		qw422016.N().S(`">`)
//line views/vestimate/vstory/Table.html:37
		components.StreamSVGLink(qw422016, `estimate`, ps)
//line views/vestimate/vstory/Table.html:37
		qw422016.N().S(`</a>
          </td>
          <td>`)
//line views/vestimate/vstory/Table.html:39
		qw422016.N().D(model.Idx)
//line views/vestimate/vstory/Table.html:39
		qw422016.N().S(`</td>
          <td class="nowrap">
            `)
//line views/vestimate/vstory/Table.html:41
		view.StreamUUID(qw422016, &model.UserID)
//line views/vestimate/vstory/Table.html:41
		if x := usersByUserID.Get(model.UserID); x != nil {
//line views/vestimate/vstory/Table.html:41
			qw422016.N().S(` (`)
//line views/vestimate/vstory/Table.html:41
			qw422016.E().S(x.TitleString())
//line views/vestimate/vstory/Table.html:41
			qw422016.N().S(`)`)
//line views/vestimate/vstory/Table.html:41
		}
//line views/vestimate/vstory/Table.html:41
		qw422016.N().S(`
            <a title="User" href="`)
//line views/vestimate/vstory/Table.html:42
		qw422016.E().S(`/admin/db/user` + `/` + model.UserID.String())
//line views/vestimate/vstory/Table.html:42
		qw422016.N().S(`">`)
//line views/vestimate/vstory/Table.html:42
		components.StreamSVGLink(qw422016, `profile`, ps)
//line views/vestimate/vstory/Table.html:42
		qw422016.N().S(`</a>
          </td>
          <td><strong>`)
//line views/vestimate/vstory/Table.html:44
		view.StreamString(qw422016, model.Title)
//line views/vestimate/vstory/Table.html:44
		qw422016.N().S(`</strong></td>
          <td>`)
//line views/vestimate/vstory/Table.html:45
		qw422016.E().S(model.Status.String())
//line views/vestimate/vstory/Table.html:45
		qw422016.N().S(`</td>
          <td>`)
//line views/vestimate/vstory/Table.html:46
		view.StreamString(qw422016, model.FinalVote)
//line views/vestimate/vstory/Table.html:46
		qw422016.N().S(`</td>
          <td>`)
//line views/vestimate/vstory/Table.html:47
		view.StreamTimestamp(qw422016, &model.Created)
//line views/vestimate/vstory/Table.html:47
		qw422016.N().S(`</td>
          <td>`)
//line views/vestimate/vstory/Table.html:48
		view.StreamTimestamp(qw422016, model.Updated)
//line views/vestimate/vstory/Table.html:48
		qw422016.N().S(`</td>
        </tr>
`)
//line views/vestimate/vstory/Table.html:50
	}
//line views/vestimate/vstory/Table.html:50
	qw422016.N().S(`      </tbody>
    </table>
  </div>
`)
//line views/vestimate/vstory/Table.html:54
	if prms.HasNextPage(len(models)+prms.Offset) || prms.HasPreviousPage() {
//line views/vestimate/vstory/Table.html:54
		qw422016.N().S(`  <hr />
  `)
//line views/vestimate/vstory/Table.html:56
		components.StreamPagination(qw422016, len(models)+prms.Offset, prms, ps.URI)
//line views/vestimate/vstory/Table.html:56
		qw422016.N().S(`
  <div class="clear"></div>
`)
//line views/vestimate/vstory/Table.html:58
	}
//line views/vestimate/vstory/Table.html:59
}

//line views/vestimate/vstory/Table.html:59
func WriteTable(qq422016 qtio422016.Writer, models story.Stories, estimatesByEstimateID estimate.Estimates, usersByUserID user.Users, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vestimate/vstory/Table.html:59
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vestimate/vstory/Table.html:59
	StreamTable(qw422016, models, estimatesByEstimateID, usersByUserID, params, as, ps)
//line views/vestimate/vstory/Table.html:59
	qt422016.ReleaseWriter(qw422016)
//line views/vestimate/vstory/Table.html:59
}

//line views/vestimate/vstory/Table.html:59
func Table(models story.Stories, estimatesByEstimateID estimate.Estimates, usersByUserID user.Users, params filter.ParamSet, as *app.State, ps *cutil.PageState) string {
//line views/vestimate/vstory/Table.html:59
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vestimate/vstory/Table.html:59
	WriteTable(qb422016, models, estimatesByEstimateID, usersByUserID, params, as, ps)
//line views/vestimate/vstory/Table.html:59
	qs422016 := string(qb422016.B)
//line views/vestimate/vstory/Table.html:59
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vestimate/vstory/Table.html:59
	return qs422016
//line views/vestimate/vstory/Table.html:59
}
