// Code generated by qtc from "SVG.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/components/SVG.html:2
package components

//line views/components/SVG.html:2
import (
	"fmt"
	"strings"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/util"
)

//line views/components/SVG.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/SVG.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/SVG.html:11
func StreamSVG(qw422016 *qt422016.Writer, k string) {
//line views/components/SVG.html:12
	if svg, ok := util.SVGLibrary[k]; ok {
//line views/components/SVG.html:13
		qw422016.N().S(svg)
//line views/components/SVG.html:14
	} else {
//line views/components/SVG.html:14
		qw422016.N().S(`<!-- missing icon definition for [`)
//line views/components/SVG.html:15
		qw422016.E().S(k)
//line views/components/SVG.html:15
		qw422016.N().S(`], using default icon -->`)
//line views/components/SVG.html:15
		qw422016.N().S(strings.ReplaceAll(util.SVGLibrary["question"], "-question", "-"+k))
//line views/components/SVG.html:16
	}
//line views/components/SVG.html:17
}

//line views/components/SVG.html:17
func WriteSVG(qq422016 qtio422016.Writer, k string) {
//line views/components/SVG.html:17
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/SVG.html:17
	StreamSVG(qw422016, k)
//line views/components/SVG.html:17
	qt422016.ReleaseWriter(qw422016)
//line views/components/SVG.html:17
}

//line views/components/SVG.html:17
func SVG(k string) string {
//line views/components/SVG.html:17
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/SVG.html:17
	WriteSVG(qb422016, k)
//line views/components/SVG.html:17
	qs422016 := string(qb422016.B)
//line views/components/SVG.html:17
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/SVG.html:17
	return qs422016
//line views/components/SVG.html:17
}

//line views/components/SVG.html:19
func StreamSVGRef(qw422016 *qt422016.Writer, k string, w int, h int, cls string, ps *cutil.PageState) {
//line views/components/SVG.html:20
	if k != "" {
//line views/components/SVG.html:22
		ps.AddIcon(k)
		if w == 0 {
			w = 20
		}
		if h == 0 {
			h = 20
		}
		style := fmt.Sprintf("width: %dpx; height: %dpx;", w, h)

//line views/components/SVG.html:27
		if cls == "" {
//line views/components/SVG.html:27
			qw422016.N().S(`<svg style="`)
//line views/components/SVG.html:28
			qw422016.E().S(style)
//line views/components/SVG.html:28
			qw422016.N().S(`"><use xlink:href="#svg-`)
//line views/components/SVG.html:28
			qw422016.E().S(k)
//line views/components/SVG.html:28
			qw422016.N().S(`" /></svg>`)
//line views/components/SVG.html:29
		} else {
//line views/components/SVG.html:29
			qw422016.N().S(`<svg class="`)
//line views/components/SVG.html:30
			qw422016.E().S(cls)
//line views/components/SVG.html:30
			qw422016.N().S(`" style="`)
//line views/components/SVG.html:30
			qw422016.E().S(style)
//line views/components/SVG.html:30
			qw422016.N().S(`"><use xlink:href="#svg-`)
//line views/components/SVG.html:30
			qw422016.E().S(k)
//line views/components/SVG.html:30
			qw422016.N().S(`" /></svg>`)
//line views/components/SVG.html:31
		}
//line views/components/SVG.html:32
	}
//line views/components/SVG.html:33
}

//line views/components/SVG.html:33
func WriteSVGRef(qq422016 qtio422016.Writer, k string, w int, h int, cls string, ps *cutil.PageState) {
//line views/components/SVG.html:33
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/SVG.html:33
	StreamSVGRef(qw422016, k, w, h, cls, ps)
//line views/components/SVG.html:33
	qt422016.ReleaseWriter(qw422016)
//line views/components/SVG.html:33
}

//line views/components/SVG.html:33
func SVGRef(k string, w int, h int, cls string, ps *cutil.PageState) string {
//line views/components/SVG.html:33
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/SVG.html:33
	WriteSVGRef(qb422016, k, w, h, cls, ps)
//line views/components/SVG.html:33
	qs422016 := string(qb422016.B)
//line views/components/SVG.html:33
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/SVG.html:33
	return qs422016
//line views/components/SVG.html:33
}

//line views/components/SVG.html:35
func StreamIcon(qw422016 *qt422016.Writer, k string, size int, cls string, ps *cutil.PageState) {
//line views/components/SVG.html:36
	if strings.Contains(k, "/") {
//line views/components/SVG.html:36
		qw422016.N().S(`<img alt="SVG icon [`)
//line views/components/SVG.html:37
		qw422016.E().S(k)
//line views/components/SVG.html:37
		qw422016.N().S(`]" src="`)
//line views/components/SVG.html:37
		qw422016.E().S(k)
//line views/components/SVG.html:37
		qw422016.N().S(`" style="width:`)
//line views/components/SVG.html:37
		qw422016.N().D(size)
//line views/components/SVG.html:37
		qw422016.N().S(`px; height:`)
//line views/components/SVG.html:37
		qw422016.N().D(size)
//line views/components/SVG.html:37
		qw422016.N().S(`px;" />`)
//line views/components/SVG.html:38
	} else {
//line views/components/SVG.html:39
		StreamSVGRef(qw422016, k, size, size, cls, ps)
//line views/components/SVG.html:40
	}
//line views/components/SVG.html:41
}

//line views/components/SVG.html:41
func WriteIcon(qq422016 qtio422016.Writer, k string, size int, cls string, ps *cutil.PageState) {
//line views/components/SVG.html:41
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/SVG.html:41
	StreamIcon(qw422016, k, size, cls, ps)
//line views/components/SVG.html:41
	qt422016.ReleaseWriter(qw422016)
//line views/components/SVG.html:41
}

//line views/components/SVG.html:41
func Icon(k string, size int, cls string, ps *cutil.PageState) string {
//line views/components/SVG.html:41
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/SVG.html:41
	WriteIcon(qb422016, k, size, cls, ps)
//line views/components/SVG.html:41
	qs422016 := string(qb422016.B)
//line views/components/SVG.html:41
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/SVG.html:41
	return qs422016
//line views/components/SVG.html:41
}

//line views/components/SVG.html:43
func StreamSVGSimple(qw422016 *qt422016.Writer, k string, size int, ps *cutil.PageState) {
//line views/components/SVG.html:44
	StreamSVGRef(qw422016, k, size, size, "", ps)
//line views/components/SVG.html:45
}

//line views/components/SVG.html:45
func WriteSVGSimple(qq422016 qtio422016.Writer, k string, size int, ps *cutil.PageState) {
//line views/components/SVG.html:45
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/SVG.html:45
	StreamSVGSimple(qw422016, k, size, ps)
//line views/components/SVG.html:45
	qt422016.ReleaseWriter(qw422016)
//line views/components/SVG.html:45
}

//line views/components/SVG.html:45
func SVGSimple(k string, size int, ps *cutil.PageState) string {
//line views/components/SVG.html:45
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/SVG.html:45
	WriteSVGSimple(qb422016, k, size, ps)
//line views/components/SVG.html:45
	qs422016 := string(qb422016.B)
//line views/components/SVG.html:45
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/SVG.html:45
	return qs422016
//line views/components/SVG.html:45
}

//line views/components/SVG.html:47
func StreamSVGButton(qw422016 *qt422016.Writer, k string, ps *cutil.PageState) {
//line views/components/SVG.html:48
	StreamSVGRef(qw422016, k, 14, 14, "icon", ps)
//line views/components/SVG.html:49
}

//line views/components/SVG.html:49
func WriteSVGButton(qq422016 qtio422016.Writer, k string, ps *cutil.PageState) {
//line views/components/SVG.html:49
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/SVG.html:49
	StreamSVGButton(qw422016, k, ps)
//line views/components/SVG.html:49
	qt422016.ReleaseWriter(qw422016)
//line views/components/SVG.html:49
}

//line views/components/SVG.html:49
func SVGButton(k string, ps *cutil.PageState) string {
//line views/components/SVG.html:49
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/SVG.html:49
	WriteSVGButton(qb422016, k, ps)
//line views/components/SVG.html:49
	qs422016 := string(qb422016.B)
//line views/components/SVG.html:49
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/SVG.html:49
	return qs422016
//line views/components/SVG.html:49
}

//line views/components/SVG.html:51
func StreamSVGIcon(qw422016 *qt422016.Writer, k string, ps *cutil.PageState) {
//line views/components/SVG.html:52
	StreamSVGRef(qw422016, k, 18, 18, "icon", ps)
//line views/components/SVG.html:53
}

//line views/components/SVG.html:53
func WriteSVGIcon(qq422016 qtio422016.Writer, k string, ps *cutil.PageState) {
//line views/components/SVG.html:53
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/SVG.html:53
	StreamSVGIcon(qw422016, k, ps)
//line views/components/SVG.html:53
	qt422016.ReleaseWriter(qw422016)
//line views/components/SVG.html:53
}

//line views/components/SVG.html:53
func SVGIcon(k string, ps *cutil.PageState) string {
//line views/components/SVG.html:53
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/SVG.html:53
	WriteSVGIcon(qb422016, k, ps)
//line views/components/SVG.html:53
	qs422016 := string(qb422016.B)
//line views/components/SVG.html:53
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/SVG.html:53
	return qs422016
//line views/components/SVG.html:53
}

//line views/components/SVG.html:55
func StreamSVGBreadcrumb(qw422016 *qt422016.Writer, k string, ps *cutil.PageState) {
//line views/components/SVG.html:56
	StreamSVGRef(qw422016, k, 18, 18, "breadcrumb-icon", ps)
//line views/components/SVG.html:57
}

//line views/components/SVG.html:57
func WriteSVGBreadcrumb(qq422016 qtio422016.Writer, k string, ps *cutil.PageState) {
//line views/components/SVG.html:57
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/SVG.html:57
	StreamSVGBreadcrumb(qw422016, k, ps)
//line views/components/SVG.html:57
	qt422016.ReleaseWriter(qw422016)
//line views/components/SVG.html:57
}

//line views/components/SVG.html:57
func SVGBreadcrumb(k string, ps *cutil.PageState) string {
//line views/components/SVG.html:57
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/SVG.html:57
	WriteSVGBreadcrumb(qb422016, k, ps)
//line views/components/SVG.html:57
	qs422016 := string(qb422016.B)
//line views/components/SVG.html:57
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/SVG.html:57
	return qs422016
//line views/components/SVG.html:57
}

//line views/components/SVG.html:59
func StreamSVGInline(qw422016 *qt422016.Writer, k string, size int, ps *cutil.PageState) {
//line views/components/SVG.html:60
	StreamSVGRef(qw422016, k, size, size, "inline", ps)
//line views/components/SVG.html:61
}

//line views/components/SVG.html:61
func WriteSVGInline(qq422016 qtio422016.Writer, k string, size int, ps *cutil.PageState) {
//line views/components/SVG.html:61
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/SVG.html:61
	StreamSVGInline(qw422016, k, size, ps)
//line views/components/SVG.html:61
	qt422016.ReleaseWriter(qw422016)
//line views/components/SVG.html:61
}

//line views/components/SVG.html:61
func SVGInline(k string, size int, ps *cutil.PageState) string {
//line views/components/SVG.html:61
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/SVG.html:61
	WriteSVGInline(qb422016, k, size, ps)
//line views/components/SVG.html:61
	qs422016 := string(qb422016.B)
//line views/components/SVG.html:61
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/SVG.html:61
	return qs422016
//line views/components/SVG.html:61
}

//line views/components/SVG.html:63
func StreamSVGLink(qw422016 *qt422016.Writer, k string, ps *cutil.PageState) {
//line views/components/SVG.html:64
	StreamSVGRef(qw422016, k, 18, 18, "link", ps)
//line views/components/SVG.html:65
}

//line views/components/SVG.html:65
func WriteSVGLink(qq422016 qtio422016.Writer, k string, ps *cutil.PageState) {
//line views/components/SVG.html:65
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/SVG.html:65
	StreamSVGLink(qw422016, k, ps)
//line views/components/SVG.html:65
	qt422016.ReleaseWriter(qw422016)
//line views/components/SVG.html:65
}

//line views/components/SVG.html:65
func SVGLink(k string, ps *cutil.PageState) string {
//line views/components/SVG.html:65
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/SVG.html:65
	WriteSVGLink(qb422016, k, ps)
//line views/components/SVG.html:65
	qs422016 := string(qb422016.B)
//line views/components/SVG.html:65
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/SVG.html:65
	return qs422016
//line views/components/SVG.html:65
}

//line views/components/SVG.html:67
func StreamIconGallery(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/components/SVG.html:67
	qw422016.N().S(`  <div class="flex-wrap mt">
`)
//line views/components/SVG.html:69
	for _, k := range util.SVGIconKeys {
//line views/components/SVG.html:69
		qw422016.N().S(`    <div class="icon-gallery-icon">
      <div class="gallery-svg">`)
//line views/components/SVG.html:71
		StreamSVGRef(qw422016, k, 64, 64, "icon", ps)
//line views/components/SVG.html:71
		qw422016.N().S(`</div>
      <div class="gallery-title">`)
//line views/components/SVG.html:72
		qw422016.E().S(k)
//line views/components/SVG.html:72
		qw422016.N().S(`</div>
    </div>
`)
//line views/components/SVG.html:74
	}
//line views/components/SVG.html:74
	qw422016.N().S(`  </div>
`)
//line views/components/SVG.html:76
}

//line views/components/SVG.html:76
func WriteIconGallery(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/components/SVG.html:76
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/SVG.html:76
	StreamIconGallery(qw422016, as, ps)
//line views/components/SVG.html:76
	qt422016.ReleaseWriter(qw422016)
//line views/components/SVG.html:76
}

//line views/components/SVG.html:76
func IconGallery(as *app.State, ps *cutil.PageState) string {
//line views/components/SVG.html:76
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/SVG.html:76
	WriteIconGallery(qb422016, as, ps)
//line views/components/SVG.html:76
	qs422016 := string(qb422016.B)
//line views/components/SVG.html:76
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/SVG.html:76
	return qs422016
//line views/components/SVG.html:76
}
