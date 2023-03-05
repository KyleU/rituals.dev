// Code generated by qtc from "Self.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vworkspace/vwutil/Self.html:1
package vwutil

//line views/vworkspace/vwutil/Self.html:1
import (
	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/components"
)

//line views/vworkspace/vwutil/Self.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vworkspace/vwutil/Self.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vworkspace/vwutil/Self.html:11
func StreamSelfLink(qw422016 *qt422016.Writer, self *util.Member, ps *cutil.PageState) {
//line views/vworkspace/vwutil/Self.html:11
	qw422016.N().S(`
  <span id="self-id" style="display: none;">`)
//line views/vworkspace/vwutil/Self.html:12
	qw422016.E().S(self.UserID.String())
//line views/vworkspace/vwutil/Self.html:12
	qw422016.N().S(`</span>
  <a id="self-edit-link" href="#modal-member-`)
//line views/vworkspace/vwutil/Self.html:13
	qw422016.E().S(self.UserID.String())
//line views/vworkspace/vwutil/Self.html:13
	qw422016.N().S(`"><h3>
    <span id="self-picture">
`)
//line views/vworkspace/vwutil/Self.html:15
	if self.Picture == "" {
//line views/vworkspace/vwutil/Self.html:15
		qw422016.N().S(`      `)
//line views/vworkspace/vwutil/Self.html:16
		components.StreamSVGRefIcon(qw422016, `profile`, ps)
//line views/vworkspace/vwutil/Self.html:16
		qw422016.N().S(`
`)
//line views/vworkspace/vwutil/Self.html:17
	} else {
//line views/vworkspace/vwutil/Self.html:17
		qw422016.N().S(`      <img class="icon" style="width: 20px; height: 20px;" src="`)
//line views/vworkspace/vwutil/Self.html:18
		qw422016.E().S(self.Picture)
//line views/vworkspace/vwutil/Self.html:18
		qw422016.N().S(`" />
`)
//line views/vworkspace/vwutil/Self.html:19
	}
//line views/vworkspace/vwutil/Self.html:19
	qw422016.N().S(`    </span>
    <span id="self-name" class="member-`)
//line views/vworkspace/vwutil/Self.html:21
	qw422016.E().S(self.UserID.String())
//line views/vworkspace/vwutil/Self.html:21
	qw422016.N().S(`-name">`)
//line views/vworkspace/vwutil/Self.html:21
	qw422016.E().S(self.Name)
//line views/vworkspace/vwutil/Self.html:21
	qw422016.N().S(`</span>
  </h3></a>
  <em id="self-role">`)
//line views/vworkspace/vwutil/Self.html:23
	qw422016.E().S(string(self.Role))
//line views/vworkspace/vwutil/Self.html:23
	qw422016.N().S(`</em>
`)
//line views/vworkspace/vwutil/Self.html:24
}

//line views/vworkspace/vwutil/Self.html:24
func WriteSelfLink(qq422016 qtio422016.Writer, self *util.Member, ps *cutil.PageState) {
//line views/vworkspace/vwutil/Self.html:24
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwutil/Self.html:24
	StreamSelfLink(qw422016, self, ps)
//line views/vworkspace/vwutil/Self.html:24
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwutil/Self.html:24
}

//line views/vworkspace/vwutil/Self.html:24
func SelfLink(self *util.Member, ps *cutil.PageState) string {
//line views/vworkspace/vwutil/Self.html:24
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwutil/Self.html:24
	WriteSelfLink(qb422016, self, ps)
//line views/vworkspace/vwutil/Self.html:24
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwutil/Self.html:24
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwutil/Self.html:24
	return qs422016
//line views/vworkspace/vwutil/Self.html:24
}

//line views/vworkspace/vwutil/Self.html:26
func StreamSelfModal(qw422016 *qt422016.Writer, id uuid.UUID, name string, picture string, role enum.MemberStatus, formAction string, ps *cutil.PageState) {
//line views/vworkspace/vwutil/Self.html:26
	qw422016.N().S(`
  <div id="modal-member-`)
//line views/vworkspace/vwutil/Self.html:27
	qw422016.E().S(id.String())
//line views/vworkspace/vwutil/Self.html:27
	qw422016.N().S(`" class="modal self" style="display: none;">
    <a class="backdrop" href="#"></a>
    <div class="modal-content">
      <div class="modal-header">
        <a href="#" class="modal-close">×</a>
        <h2>Edit Self</h2>
      </div>
      <div class="modal-body">
        <form action="`)
//line views/vworkspace/vwutil/Self.html:35
	qw422016.E().S(formAction)
//line views/vworkspace/vwutil/Self.html:35
	qw422016.N().S(`" method="post" class="expanded">
          <input type="hidden" name="action" value="`)
//line views/vworkspace/vwutil/Self.html:36
	qw422016.E().S(string(action.ActMemberSelf))
//line views/vworkspace/vwutil/Self.html:36
	qw422016.N().S(`" />
          <em>Name</em><br />
          `)
//line views/vworkspace/vwutil/Self.html:38
	components.StreamFormInput(qw422016, "name", "self-name-input", name, "The name you wish to be called")
//line views/vworkspace/vwutil/Self.html:38
	qw422016.N().S(`
          <div class="mt"><label><input type="radio" name="choice" value="local" checked="checked"> Change for this session only</label></div>
          <div><label><input type="radio" name="choice" value="global"> Change global default</label></div>
          <hr />
          <em>Picture</em>
`)
//line views/vworkspace/vwutil/Self.html:43
	if images := ps.Accounts.Images(); len(images) > 0 {
//line views/vworkspace/vwutil/Self.html:43
		qw422016.N().S(`          <div class="choice">
            <label title="no picture">
`)
//line views/vworkspace/vwutil/Self.html:46
		if picture == "" {
//line views/vworkspace/vwutil/Self.html:46
			qw422016.N().S(`              <input type="radio" name="picture" value="" checked="checked">
`)
//line views/vworkspace/vwutil/Self.html:48
		} else {
//line views/vworkspace/vwutil/Self.html:48
			qw422016.N().S(`              <input type="radio" name="picture" value="">
`)
//line views/vworkspace/vwutil/Self.html:50
		}
//line views/vworkspace/vwutil/Self.html:50
		qw422016.N().S(`              `)
//line views/vworkspace/vwutil/Self.html:51
		components.StreamSVGRef(qw422016, "times", 32, 32, "", ps)
//line views/vworkspace/vwutil/Self.html:51
		qw422016.N().S(`
            </label>
`)
//line views/vworkspace/vwutil/Self.html:53
		for _, i := range images {
//line views/vworkspace/vwutil/Self.html:53
			qw422016.N().S(`            <label title="`)
//line views/vworkspace/vwutil/Self.html:54
			qw422016.E().S(i)
//line views/vworkspace/vwutil/Self.html:54
			qw422016.N().S(`">
`)
//line views/vworkspace/vwutil/Self.html:55
			if i == picture {
//line views/vworkspace/vwutil/Self.html:55
				qw422016.N().S(`              <input type="radio" name="picture" value="`)
//line views/vworkspace/vwutil/Self.html:56
				qw422016.E().S(i)
//line views/vworkspace/vwutil/Self.html:56
				qw422016.N().S(`" checked="checked">
`)
//line views/vworkspace/vwutil/Self.html:57
			} else {
//line views/vworkspace/vwutil/Self.html:57
				qw422016.N().S(`              <input type="radio" name="picture" value="`)
//line views/vworkspace/vwutil/Self.html:58
				qw422016.E().S(i)
//line views/vworkspace/vwutil/Self.html:58
				qw422016.N().S(`">
`)
//line views/vworkspace/vwutil/Self.html:59
			}
//line views/vworkspace/vwutil/Self.html:59
			qw422016.N().S(`              <img style="width: 32px; height: 32px;" src="`)
//line views/vworkspace/vwutil/Self.html:60
			qw422016.E().S(i)
//line views/vworkspace/vwutil/Self.html:60
			qw422016.N().S(`" />
            </label>
`)
//line views/vworkspace/vwutil/Self.html:62
		}
//line views/vworkspace/vwutil/Self.html:62
		qw422016.N().S(`            <div class="clear"></div>
          </div>
`)
//line views/vworkspace/vwutil/Self.html:65
	} else {
//line views/vworkspace/vwutil/Self.html:65
		qw422016.N().S(`          <div>To set a profile picture, <a href="/profile">sign in</a></div>
`)
//line views/vworkspace/vwutil/Self.html:67
	}
//line views/vworkspace/vwutil/Self.html:67
	qw422016.N().S(`          <hr />
          <div class="right"><button type="submit">Save</button></div>
          <a href="#"><button type="button">Leave</button></a>
        </form>
      </div>
    </div>
  </div>
`)
//line views/vworkspace/vwutil/Self.html:75
}

//line views/vworkspace/vwutil/Self.html:75
func WriteSelfModal(qq422016 qtio422016.Writer, id uuid.UUID, name string, picture string, role enum.MemberStatus, formAction string, ps *cutil.PageState) {
//line views/vworkspace/vwutil/Self.html:75
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwutil/Self.html:75
	StreamSelfModal(qw422016, id, name, picture, role, formAction, ps)
//line views/vworkspace/vwutil/Self.html:75
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwutil/Self.html:75
}

//line views/vworkspace/vwutil/Self.html:75
func SelfModal(id uuid.UUID, name string, picture string, role enum.MemberStatus, formAction string, ps *cutil.PageState) string {
//line views/vworkspace/vwutil/Self.html:75
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwutil/Self.html:75
	WriteSelfModal(qb422016, id, name, picture, role, formAction, ps)
//line views/vworkspace/vwutil/Self.html:75
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwutil/Self.html:75
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwutil/Self.html:75
	return qs422016
//line views/vworkspace/vwutil/Self.html:75
}
