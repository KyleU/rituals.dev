// Code generated by qtc from "Display.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/components/Display.html:2
package components

//line views/components/Display.html:2
import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/util"
)

//line views/components/Display.html:14
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/Display.html:14
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/Display.html:14
func StreamDisplayTimestamp(qw422016 *qt422016.Writer, value *time.Time) {
//line views/components/Display.html:14
	qw422016.N().S(`<span class="nowrap">`)
//line views/components/Display.html:15
	qw422016.E().S(util.TimeToFull(value))
//line views/components/Display.html:15
	qw422016.N().S(`</span>`)
//line views/components/Display.html:16
}

//line views/components/Display.html:16
func WriteDisplayTimestamp(qq422016 qtio422016.Writer, value *time.Time) {
//line views/components/Display.html:16
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Display.html:16
	StreamDisplayTimestamp(qw422016, value)
//line views/components/Display.html:16
	qt422016.ReleaseWriter(qw422016)
//line views/components/Display.html:16
}

//line views/components/Display.html:16
func DisplayTimestamp(value *time.Time) string {
//line views/components/Display.html:16
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Display.html:16
	WriteDisplayTimestamp(qb422016, value)
//line views/components/Display.html:16
	qs422016 := string(qb422016.B)
//line views/components/Display.html:16
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Display.html:16
	return qs422016
//line views/components/Display.html:16
}

//line views/components/Display.html:18
func StreamDisplayTimestampRelative(qw422016 *qt422016.Writer, value *time.Time) {
//line views/components/Display.html:18
	qw422016.N().S(`<span class="nowrap reltime" data-time="`)
//line views/components/Display.html:19
	qw422016.E().S(util.TimeToFull(value))
//line views/components/Display.html:19
	qw422016.N().S(`">`)
//line views/components/Display.html:19
	qw422016.E().S(util.TimeRelative(value))
//line views/components/Display.html:19
	qw422016.N().S(`</span>`)
//line views/components/Display.html:20
}

//line views/components/Display.html:20
func WriteDisplayTimestampRelative(qq422016 qtio422016.Writer, value *time.Time) {
//line views/components/Display.html:20
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Display.html:20
	StreamDisplayTimestampRelative(qw422016, value)
//line views/components/Display.html:20
	qt422016.ReleaseWriter(qw422016)
//line views/components/Display.html:20
}

//line views/components/Display.html:20
func DisplayTimestampRelative(value *time.Time) string {
//line views/components/Display.html:20
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Display.html:20
	WriteDisplayTimestampRelative(qb422016, value)
//line views/components/Display.html:20
	qs422016 := string(qb422016.B)
//line views/components/Display.html:20
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Display.html:20
	return qs422016
//line views/components/Display.html:20
}

//line views/components/Display.html:22
func StreamDisplayTimestampDay(qw422016 *qt422016.Writer, value *time.Time) {
//line views/components/Display.html:23
	qw422016.E().S(util.TimeToYMD(value))
//line views/components/Display.html:24
}

//line views/components/Display.html:24
func WriteDisplayTimestampDay(qq422016 qtio422016.Writer, value *time.Time) {
//line views/components/Display.html:24
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Display.html:24
	StreamDisplayTimestampDay(qw422016, value)
//line views/components/Display.html:24
	qt422016.ReleaseWriter(qw422016)
//line views/components/Display.html:24
}

//line views/components/Display.html:24
func DisplayTimestampDay(value *time.Time) string {
//line views/components/Display.html:24
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Display.html:24
	WriteDisplayTimestampDay(qb422016, value)
//line views/components/Display.html:24
	qs422016 := string(qb422016.B)
//line views/components/Display.html:24
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Display.html:24
	return qs422016
//line views/components/Display.html:24
}

//line views/components/Display.html:26
func StreamDisplayUUID(qw422016 *qt422016.Writer, value *uuid.UUID) {
//line views/components/Display.html:27
	if value != nil {
//line views/components/Display.html:28
		qw422016.E().S(value.String())
//line views/components/Display.html:29
	}
//line views/components/Display.html:30
}

//line views/components/Display.html:30
func WriteDisplayUUID(qq422016 qtio422016.Writer, value *uuid.UUID) {
//line views/components/Display.html:30
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Display.html:30
	StreamDisplayUUID(qw422016, value)
//line views/components/Display.html:30
	qt422016.ReleaseWriter(qw422016)
//line views/components/Display.html:30
}

//line views/components/Display.html:30
func DisplayUUID(value *uuid.UUID) string {
//line views/components/Display.html:30
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Display.html:30
	WriteDisplayUUID(qb422016, value)
//line views/components/Display.html:30
	qs422016 := string(qb422016.B)
//line views/components/Display.html:30
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Display.html:30
	return qs422016
//line views/components/Display.html:30
}

//line views/components/Display.html:32
func StreamDisplayStringArray(qw422016 *qt422016.Writer, value []string) {
//line views/components/Display.html:33
	if len(value) == 0 {
//line views/components/Display.html:33
		qw422016.N().S(`<em>empty</em>`)
//line views/components/Display.html:35
	}
//line views/components/Display.html:37
	maxCount := 5
	display := value
	var extra int
	if len(value) > maxCount {
		extra = len(value) - maxCount
		display = display[:maxCount]
	}

//line views/components/Display.html:45
	if extra > 0 {
//line views/components/Display.html:45
		qw422016.N().S(`<span title="`)
//line views/components/Display.html:45
		qw422016.E().S(strings.Join(value, `, `))
//line views/components/Display.html:45
		qw422016.N().S(`">`)
//line views/components/Display.html:45
	}
//line views/components/Display.html:46
	for idx, v := range display {
//line views/components/Display.html:47
		if idx > 0 {
//line views/components/Display.html:47
			qw422016.N().S(`,`)
//line views/components/Display.html:47
			qw422016.N().S(` `)
//line views/components/Display.html:47
		}
//line views/components/Display.html:48
		qw422016.E().S(v)
//line views/components/Display.html:49
	}
//line views/components/Display.html:50
	if extra > 0 {
//line views/components/Display.html:50
		qw422016.N().S(`, <em>and`)
//line views/components/Display.html:50
		qw422016.N().S(` `)
//line views/components/Display.html:50
		qw422016.N().D(extra)
//line views/components/Display.html:50
		qw422016.N().S(` `)
//line views/components/Display.html:50
		qw422016.N().S(`more...</em>`)
//line views/components/Display.html:50
	}
//line views/components/Display.html:51
	if extra > 0 {
//line views/components/Display.html:51
		qw422016.N().S(`</span>`)
//line views/components/Display.html:51
	}
//line views/components/Display.html:52
}

//line views/components/Display.html:52
func WriteDisplayStringArray(qq422016 qtio422016.Writer, value []string) {
//line views/components/Display.html:52
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Display.html:52
	StreamDisplayStringArray(qw422016, value)
//line views/components/Display.html:52
	qt422016.ReleaseWriter(qw422016)
//line views/components/Display.html:52
}

//line views/components/Display.html:52
func DisplayStringArray(value []string) string {
//line views/components/Display.html:52
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Display.html:52
	WriteDisplayStringArray(qb422016, value)
//line views/components/Display.html:52
	qs422016 := string(qb422016.B)
//line views/components/Display.html:52
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Display.html:52
	return qs422016
//line views/components/Display.html:52
}

//line views/components/Display.html:54
func StreamDisplayIntArray(qw422016 *qt422016.Writer, value []int) {
//line views/components/Display.html:55
	StreamDisplayStringArray(qw422016, util.ArrayToStringArray(value))
//line views/components/Display.html:56
}

//line views/components/Display.html:56
func WriteDisplayIntArray(qq422016 qtio422016.Writer, value []int) {
//line views/components/Display.html:56
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Display.html:56
	StreamDisplayIntArray(qw422016, value)
//line views/components/Display.html:56
	qt422016.ReleaseWriter(qw422016)
//line views/components/Display.html:56
}

//line views/components/Display.html:56
func DisplayIntArray(value []int) string {
//line views/components/Display.html:56
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Display.html:56
	WriteDisplayIntArray(qb422016, value)
//line views/components/Display.html:56
	qs422016 := string(qb422016.B)
//line views/components/Display.html:56
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Display.html:56
	return qs422016
//line views/components/Display.html:56
}

//line views/components/Display.html:58
func StreamDisplayFloatArray(qw422016 *qt422016.Writer, value []float64) {
//line views/components/Display.html:59
	StreamDisplayStringArray(qw422016, util.ArrayToStringArray(value))
//line views/components/Display.html:60
}

//line views/components/Display.html:60
func WriteDisplayFloatArray(qq422016 qtio422016.Writer, value []float64) {
//line views/components/Display.html:60
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Display.html:60
	StreamDisplayFloatArray(qw422016, value)
//line views/components/Display.html:60
	qt422016.ReleaseWriter(qw422016)
//line views/components/Display.html:60
}

//line views/components/Display.html:60
func DisplayFloatArray(value []float64) string {
//line views/components/Display.html:60
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Display.html:60
	WriteDisplayFloatArray(qb422016, value)
//line views/components/Display.html:60
	qs422016 := string(qb422016.B)
//line views/components/Display.html:60
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Display.html:60
	return qs422016
//line views/components/Display.html:60
}

//line views/components/Display.html:62
func StreamDisplayDiffs(qw422016 *qt422016.Writer, value util.Diffs) {
//line views/components/Display.html:63
	if len(value) == 0 {
//line views/components/Display.html:63
		qw422016.N().S(`<em>no changes</em>`)
//line views/components/Display.html:65
	} else {
//line views/components/Display.html:65
		qw422016.N().S(`<div class="overflow full-width"><table><tbody>`)
//line views/components/Display.html:69
		for _, d := range value {
//line views/components/Display.html:69
			qw422016.N().S(`<tr><td style="width: 30%;"><code>`)
//line views/components/Display.html:71
			qw422016.E().S(d.Path)
//line views/components/Display.html:71
			qw422016.N().S(`</code></td><td style="width: 30%;"><code class="error">`)
//line views/components/Display.html:72
			qw422016.E().S(d.Old)
//line views/components/Display.html:72
			qw422016.N().S(`</code></td><td style="width: 10%;">→</td><td style="width: 30%;"><code class="success">`)
//line views/components/Display.html:74
			qw422016.E().S(d.New)
//line views/components/Display.html:74
			qw422016.N().S(`</code></td></tr>`)
//line views/components/Display.html:76
		}
//line views/components/Display.html:76
		qw422016.N().S(`</tbody></table></div>`)
//line views/components/Display.html:80
	}
//line views/components/Display.html:81
}

//line views/components/Display.html:81
func WriteDisplayDiffs(qq422016 qtio422016.Writer, value util.Diffs) {
//line views/components/Display.html:81
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Display.html:81
	StreamDisplayDiffs(qw422016, value)
//line views/components/Display.html:81
	qt422016.ReleaseWriter(qw422016)
//line views/components/Display.html:81
}

//line views/components/Display.html:81
func DisplayDiffs(value util.Diffs) string {
//line views/components/Display.html:81
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Display.html:81
	WriteDisplayDiffs(qb422016, value)
//line views/components/Display.html:81
	qs422016 := string(qb422016.B)
//line views/components/Display.html:81
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Display.html:81
	return qs422016
//line views/components/Display.html:81
}

//line views/components/Display.html:83
func StreamDisplayMaps(qw422016 *qt422016.Writer, maps []util.ValueMap, params *filter.Params, preserveWhitespace bool, ps *cutil.PageState) {
//line views/components/Display.html:84
	if len(maps) == 0 {
//line views/components/Display.html:84
		qw422016.N().S(`<em>no results</em>`)
//line views/components/Display.html:86
	} else {
//line views/components/Display.html:86
		qw422016.N().S(`<div class="overflow full-width"><table><thead><tr>`)
//line views/components/Display.html:91
		for _, k := range maps[0].Keys() {
//line views/components/Display.html:92
			StreamTableHeaderSimple(qw422016, "map", k, k, "", params, nil, ps)
//line views/components/Display.html:93
		}
//line views/components/Display.html:93
		qw422016.N().S(`</tr></thead><tbody>`)
//line views/components/Display.html:97
		for _, m := range maps {
//line views/components/Display.html:97
			qw422016.N().S(`<tr>`)
//line views/components/Display.html:99
			for _, k := range m.Keys() {
//line views/components/Display.html:101
				res := ""
				switch t := m[k].(type) {
				case string:
					res = t
				case []byte:
					res = string(t)
				default:
					res = fmt.Sprint(m[k])
				}

//line views/components/Display.html:111
				if preserveWhitespace {
//line views/components/Display.html:111
					qw422016.N().S(`<td class="prews">`)
//line views/components/Display.html:112
					qw422016.E().S(res)
//line views/components/Display.html:112
					qw422016.N().S(`</td>`)
//line views/components/Display.html:113
				} else {
//line views/components/Display.html:113
					qw422016.N().S(`<td>`)
//line views/components/Display.html:114
					qw422016.E().S(res)
//line views/components/Display.html:114
					qw422016.N().S(`</td>`)
//line views/components/Display.html:115
				}
//line views/components/Display.html:116
			}
//line views/components/Display.html:116
			qw422016.N().S(`</tr>`)
//line views/components/Display.html:118
		}
//line views/components/Display.html:118
		qw422016.N().S(`</tbody></table></div>`)
//line views/components/Display.html:122
	}
//line views/components/Display.html:123
}

//line views/components/Display.html:123
func WriteDisplayMaps(qq422016 qtio422016.Writer, maps []util.ValueMap, params *filter.Params, preserveWhitespace bool, ps *cutil.PageState) {
//line views/components/Display.html:123
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Display.html:123
	StreamDisplayMaps(qw422016, maps, params, preserveWhitespace, ps)
//line views/components/Display.html:123
	qt422016.ReleaseWriter(qw422016)
//line views/components/Display.html:123
}

//line views/components/Display.html:123
func DisplayMaps(maps []util.ValueMap, params *filter.Params, preserveWhitespace bool, ps *cutil.PageState) string {
//line views/components/Display.html:123
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Display.html:123
	WriteDisplayMaps(qb422016, maps, params, preserveWhitespace, ps)
//line views/components/Display.html:123
	qs422016 := string(qb422016.B)
//line views/components/Display.html:123
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Display.html:123
	return qs422016
//line views/components/Display.html:123
}

//line views/components/Display.html:125
func StreamFormat(qw422016 *qt422016.Writer, v string, ext string) {
//line views/components/Display.html:126
	out, err := cutil.FormatLang(v, ext)

//line views/components/Display.html:127
	if err == nil {
//line views/components/Display.html:128
		qw422016.N().S(out)
//line views/components/Display.html:129
	} else {
//line views/components/Display.html:130
		qw422016.E().S(err.Error())
//line views/components/Display.html:131
	}
//line views/components/Display.html:132
}

//line views/components/Display.html:132
func WriteFormat(qq422016 qtio422016.Writer, v string, ext string) {
//line views/components/Display.html:132
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Display.html:132
	StreamFormat(qw422016, v, ext)
//line views/components/Display.html:132
	qt422016.ReleaseWriter(qw422016)
//line views/components/Display.html:132
}

//line views/components/Display.html:132
func Format(v string, ext string) string {
//line views/components/Display.html:132
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Display.html:132
	WriteFormat(qb422016, v, ext)
//line views/components/Display.html:132
	qs422016 := string(qb422016.B)
//line views/components/Display.html:132
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Display.html:132
	return qs422016
//line views/components/Display.html:132
}
