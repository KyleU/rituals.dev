// Code generated by qtc from "Table.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/components/Table.html:2
package components

//line views/components/Table.html:2
import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/views/vutil"
)

//line views/components/Table.html:13
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/Table.html:13
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/Table.html:13
func StreamTableInput(qw422016 *qt422016.Writer, key string, title string, value string, indent int, help ...string) {
//line views/components/Table.html:13
	qw422016.N().S(`<tr>`)
//line views/components/Table.html:15
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:15
	qw422016.N().S(`<th class="shrink"><label for="input-`)
//line views/components/Table.html:16
	qw422016.E().S(key)
//line views/components/Table.html:16
	qw422016.N().S(`"`)
//line views/components/Table.html:16
	streamtitleFor(qw422016, help)
//line views/components/Table.html:16
	qw422016.N().S(`>`)
//line views/components/Table.html:16
	qw422016.E().S(title)
//line views/components/Table.html:16
	qw422016.N().S(`</label></th>`)
//line views/components/Table.html:17
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:17
	qw422016.N().S(`<td>`)
//line views/components/Table.html:18
	StreamFormInput(qw422016, key, "input-"+key, value, help...)
//line views/components/Table.html:18
	qw422016.N().S(`</td>`)
//line views/components/Table.html:19
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/Table.html:19
	qw422016.N().S(`</tr>`)
//line views/components/Table.html:21
}

//line views/components/Table.html:21
func WriteTableInput(qq422016 qtio422016.Writer, key string, title string, value string, indent int, help ...string) {
//line views/components/Table.html:21
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Table.html:21
	StreamTableInput(qw422016, key, title, value, indent, help...)
//line views/components/Table.html:21
	qt422016.ReleaseWriter(qw422016)
//line views/components/Table.html:21
}

//line views/components/Table.html:21
func TableInput(key string, title string, value string, indent int, help ...string) string {
//line views/components/Table.html:21
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Table.html:21
	WriteTableInput(qb422016, key, title, value, indent, help...)
//line views/components/Table.html:21
	qs422016 := string(qb422016.B)
//line views/components/Table.html:21
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Table.html:21
	return qs422016
//line views/components/Table.html:21
}

//line views/components/Table.html:23
func StreamTableInputPassword(qw422016 *qt422016.Writer, key string, title string, value string, indent int, help ...string) {
//line views/components/Table.html:23
	qw422016.N().S(`<tr>`)
//line views/components/Table.html:25
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:25
	qw422016.N().S(`<th class="shrink"><label for="input-`)
//line views/components/Table.html:26
	qw422016.E().S(key)
//line views/components/Table.html:26
	qw422016.N().S(`"`)
//line views/components/Table.html:26
	streamtitleFor(qw422016, help)
//line views/components/Table.html:26
	qw422016.N().S(`>`)
//line views/components/Table.html:26
	qw422016.E().S(title)
//line views/components/Table.html:26
	qw422016.N().S(`</label></th>`)
//line views/components/Table.html:27
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:27
	qw422016.N().S(`<td>`)
//line views/components/Table.html:28
	StreamFormInputPassword(qw422016, key, "input-"+key, value, help...)
//line views/components/Table.html:28
	qw422016.N().S(`</td>`)
//line views/components/Table.html:29
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/Table.html:29
	qw422016.N().S(`</tr>`)
//line views/components/Table.html:31
}

//line views/components/Table.html:31
func WriteTableInputPassword(qq422016 qtio422016.Writer, key string, title string, value string, indent int, help ...string) {
//line views/components/Table.html:31
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Table.html:31
	StreamTableInputPassword(qw422016, key, title, value, indent, help...)
//line views/components/Table.html:31
	qt422016.ReleaseWriter(qw422016)
//line views/components/Table.html:31
}

//line views/components/Table.html:31
func TableInputPassword(key string, title string, value string, indent int, help ...string) string {
//line views/components/Table.html:31
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Table.html:31
	WriteTableInputPassword(qb422016, key, title, value, indent, help...)
//line views/components/Table.html:31
	qs422016 := string(qb422016.B)
//line views/components/Table.html:31
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Table.html:31
	return qs422016
//line views/components/Table.html:31
}

//line views/components/Table.html:33
func StreamTableInputNumber(qw422016 *qt422016.Writer, key string, title string, value int, indent int, help ...string) {
//line views/components/Table.html:33
	qw422016.N().S(`<tr>`)
//line views/components/Table.html:35
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:35
	qw422016.N().S(`<th class="shrink"><label for="input-`)
//line views/components/Table.html:36
	qw422016.E().S(key)
//line views/components/Table.html:36
	qw422016.N().S(`"`)
//line views/components/Table.html:36
	streamtitleFor(qw422016, help)
//line views/components/Table.html:36
	qw422016.N().S(`>`)
//line views/components/Table.html:36
	qw422016.E().S(title)
//line views/components/Table.html:36
	qw422016.N().S(`</label></th>`)
//line views/components/Table.html:37
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:37
	qw422016.N().S(`<td>`)
//line views/components/Table.html:38
	StreamFormInputNumber(qw422016, key, "input-"+key, value, help...)
//line views/components/Table.html:38
	qw422016.N().S(`</td>`)
//line views/components/Table.html:39
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/Table.html:39
	qw422016.N().S(`</tr>`)
//line views/components/Table.html:41
}

//line views/components/Table.html:41
func WriteTableInputNumber(qq422016 qtio422016.Writer, key string, title string, value int, indent int, help ...string) {
//line views/components/Table.html:41
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Table.html:41
	StreamTableInputNumber(qw422016, key, title, value, indent, help...)
//line views/components/Table.html:41
	qt422016.ReleaseWriter(qw422016)
//line views/components/Table.html:41
}

//line views/components/Table.html:41
func TableInputNumber(key string, title string, value int, indent int, help ...string) string {
//line views/components/Table.html:41
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Table.html:41
	WriteTableInputNumber(qb422016, key, title, value, indent, help...)
//line views/components/Table.html:41
	qs422016 := string(qb422016.B)
//line views/components/Table.html:41
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Table.html:41
	return qs422016
//line views/components/Table.html:41
}

//line views/components/Table.html:43
func StreamTableInputFloat(qw422016 *qt422016.Writer, key string, title string, value float64, indent int, help ...string) {
//line views/components/Table.html:43
	qw422016.N().S(`<tr>`)
//line views/components/Table.html:45
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:45
	qw422016.N().S(`<th class="shrink"><label for="input-`)
//line views/components/Table.html:46
	qw422016.E().S(key)
//line views/components/Table.html:46
	qw422016.N().S(`"`)
//line views/components/Table.html:46
	streamtitleFor(qw422016, help)
//line views/components/Table.html:46
	qw422016.N().S(`>`)
//line views/components/Table.html:46
	qw422016.E().S(title)
//line views/components/Table.html:46
	qw422016.N().S(`</label></th>`)
//line views/components/Table.html:47
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:47
	qw422016.N().S(`<td>`)
//line views/components/Table.html:48
	StreamFormInputFloat(qw422016, key, "input-"+key, value, help...)
//line views/components/Table.html:48
	qw422016.N().S(`</td>`)
//line views/components/Table.html:49
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/Table.html:49
	qw422016.N().S(`</tr>`)
//line views/components/Table.html:51
}

//line views/components/Table.html:51
func WriteTableInputFloat(qq422016 qtio422016.Writer, key string, title string, value float64, indent int, help ...string) {
//line views/components/Table.html:51
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Table.html:51
	StreamTableInputFloat(qw422016, key, title, value, indent, help...)
//line views/components/Table.html:51
	qt422016.ReleaseWriter(qw422016)
//line views/components/Table.html:51
}

//line views/components/Table.html:51
func TableInputFloat(key string, title string, value float64, indent int, help ...string) string {
//line views/components/Table.html:51
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Table.html:51
	WriteTableInputFloat(qb422016, key, title, value, indent, help...)
//line views/components/Table.html:51
	qs422016 := string(qb422016.B)
//line views/components/Table.html:51
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Table.html:51
	return qs422016
//line views/components/Table.html:51
}

//line views/components/Table.html:53
func StreamTableInputTimestamp(qw422016 *qt422016.Writer, key string, title string, value *time.Time, indent int, help ...string) {
//line views/components/Table.html:53
	qw422016.N().S(`<tr>`)
//line views/components/Table.html:55
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:55
	qw422016.N().S(`<th class="shrink"><label for="input-`)
//line views/components/Table.html:56
	qw422016.E().S(key)
//line views/components/Table.html:56
	qw422016.N().S(`"`)
//line views/components/Table.html:56
	streamtitleFor(qw422016, help)
//line views/components/Table.html:56
	qw422016.N().S(`>`)
//line views/components/Table.html:56
	qw422016.E().S(title)
//line views/components/Table.html:56
	qw422016.N().S(`</label></th>`)
//line views/components/Table.html:57
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:57
	qw422016.N().S(`<td>`)
//line views/components/Table.html:58
	StreamFormInputTimestamp(qw422016, key, "input-"+key, value, help...)
//line views/components/Table.html:58
	qw422016.N().S(`</td>`)
//line views/components/Table.html:59
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/Table.html:59
	qw422016.N().S(`</tr>`)
//line views/components/Table.html:61
}

//line views/components/Table.html:61
func WriteTableInputTimestamp(qq422016 qtio422016.Writer, key string, title string, value *time.Time, indent int, help ...string) {
//line views/components/Table.html:61
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Table.html:61
	StreamTableInputTimestamp(qw422016, key, title, value, indent, help...)
//line views/components/Table.html:61
	qt422016.ReleaseWriter(qw422016)
//line views/components/Table.html:61
}

//line views/components/Table.html:61
func TableInputTimestamp(key string, title string, value *time.Time, indent int, help ...string) string {
//line views/components/Table.html:61
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Table.html:61
	WriteTableInputTimestamp(qb422016, key, title, value, indent, help...)
//line views/components/Table.html:61
	qs422016 := string(qb422016.B)
//line views/components/Table.html:61
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Table.html:61
	return qs422016
//line views/components/Table.html:61
}

//line views/components/Table.html:63
func StreamTableInputTimestampDay(qw422016 *qt422016.Writer, key string, title string, value *time.Time, indent int, help ...string) {
//line views/components/Table.html:63
	qw422016.N().S(`<tr>`)
//line views/components/Table.html:65
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:65
	qw422016.N().S(`<th class="shrink"><label for="input-`)
//line views/components/Table.html:66
	qw422016.E().S(key)
//line views/components/Table.html:66
	qw422016.N().S(`"`)
//line views/components/Table.html:66
	streamtitleFor(qw422016, help)
//line views/components/Table.html:66
	qw422016.N().S(`>`)
//line views/components/Table.html:66
	qw422016.E().S(title)
//line views/components/Table.html:66
	qw422016.N().S(`</label></th>`)
//line views/components/Table.html:67
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:67
	qw422016.N().S(`<td>`)
//line views/components/Table.html:68
	StreamFormInputTimestampDay(qw422016, key, "input-"+key, value, help...)
//line views/components/Table.html:68
	qw422016.N().S(`</td>`)
//line views/components/Table.html:69
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/Table.html:69
	qw422016.N().S(`</tr>`)
//line views/components/Table.html:71
}

//line views/components/Table.html:71
func WriteTableInputTimestampDay(qq422016 qtio422016.Writer, key string, title string, value *time.Time, indent int, help ...string) {
//line views/components/Table.html:71
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Table.html:71
	StreamTableInputTimestampDay(qw422016, key, title, value, indent, help...)
//line views/components/Table.html:71
	qt422016.ReleaseWriter(qw422016)
//line views/components/Table.html:71
}

//line views/components/Table.html:71
func TableInputTimestampDay(key string, title string, value *time.Time, indent int, help ...string) string {
//line views/components/Table.html:71
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Table.html:71
	WriteTableInputTimestampDay(qb422016, key, title, value, indent, help...)
//line views/components/Table.html:71
	qs422016 := string(qb422016.B)
//line views/components/Table.html:71
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Table.html:71
	return qs422016
//line views/components/Table.html:71
}

//line views/components/Table.html:73
func StreamTableInputUUID(qw422016 *qt422016.Writer, key string, title string, value *uuid.UUID, indent int, help ...string) {
//line views/components/Table.html:73
	qw422016.N().S(`<tr>`)
//line views/components/Table.html:75
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:75
	qw422016.N().S(`<th class="shrink"><label for="input-`)
//line views/components/Table.html:76
	qw422016.E().S(key)
//line views/components/Table.html:76
	qw422016.N().S(`"`)
//line views/components/Table.html:76
	streamtitleFor(qw422016, help)
//line views/components/Table.html:76
	qw422016.N().S(`>`)
//line views/components/Table.html:76
	qw422016.E().S(title)
//line views/components/Table.html:76
	qw422016.N().S(`</label></th>`)
//line views/components/Table.html:77
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:77
	qw422016.N().S(`<td>`)
//line views/components/Table.html:78
	StreamFormInputUUID(qw422016, key, "input-"+key, value, help...)
//line views/components/Table.html:78
	qw422016.N().S(`</td>`)
//line views/components/Table.html:79
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/Table.html:79
	qw422016.N().S(`</tr>`)
//line views/components/Table.html:81
}

//line views/components/Table.html:81
func WriteTableInputUUID(qq422016 qtio422016.Writer, key string, title string, value *uuid.UUID, indent int, help ...string) {
//line views/components/Table.html:81
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Table.html:81
	StreamTableInputUUID(qw422016, key, title, value, indent, help...)
//line views/components/Table.html:81
	qt422016.ReleaseWriter(qw422016)
//line views/components/Table.html:81
}

//line views/components/Table.html:81
func TableInputUUID(key string, title string, value *uuid.UUID, indent int, help ...string) string {
//line views/components/Table.html:81
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Table.html:81
	WriteTableInputUUID(qb422016, key, title, value, indent, help...)
//line views/components/Table.html:81
	qs422016 := string(qb422016.B)
//line views/components/Table.html:81
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Table.html:81
	return qs422016
//line views/components/Table.html:81
}

//line views/components/Table.html:83
func StreamTableTextarea(qw422016 *qt422016.Writer, key string, title string, rows int, value string, indent int, help ...string) {
//line views/components/Table.html:83
	qw422016.N().S(`<tr>`)
//line views/components/Table.html:85
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:85
	qw422016.N().S(`<th class="shrink"><label for="input-`)
//line views/components/Table.html:86
	qw422016.E().S(key)
//line views/components/Table.html:86
	qw422016.N().S(`"`)
//line views/components/Table.html:86
	streamtitleFor(qw422016, help)
//line views/components/Table.html:86
	qw422016.N().S(`>`)
//line views/components/Table.html:86
	qw422016.E().S(title)
//line views/components/Table.html:86
	qw422016.N().S(`</label></th>`)
//line views/components/Table.html:87
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:87
	qw422016.N().S(`<td>`)
//line views/components/Table.html:88
	StreamFormTextarea(qw422016, key, "input-"+key, rows, value, help...)
//line views/components/Table.html:88
	qw422016.N().S(`</td>`)
//line views/components/Table.html:89
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/Table.html:89
	qw422016.N().S(`</tr>`)
//line views/components/Table.html:91
}

//line views/components/Table.html:91
func WriteTableTextarea(qq422016 qtio422016.Writer, key string, title string, rows int, value string, indent int, help ...string) {
//line views/components/Table.html:91
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Table.html:91
	StreamTableTextarea(qw422016, key, title, rows, value, indent, help...)
//line views/components/Table.html:91
	qt422016.ReleaseWriter(qw422016)
//line views/components/Table.html:91
}

//line views/components/Table.html:91
func TableTextarea(key string, title string, rows int, value string, indent int, help ...string) string {
//line views/components/Table.html:91
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Table.html:91
	WriteTableTextarea(qb422016, key, title, rows, value, indent, help...)
//line views/components/Table.html:91
	qs422016 := string(qb422016.B)
//line views/components/Table.html:91
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Table.html:91
	return qs422016
//line views/components/Table.html:91
}

//line views/components/Table.html:93
func StreamTableSelect(qw422016 *qt422016.Writer, key string, title string, value string, opts []string, titles []string, indent int, help ...string) {
//line views/components/Table.html:93
	qw422016.N().S(`<tr>`)
//line views/components/Table.html:95
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:95
	qw422016.N().S(`<th class="shrink"><label for="input-`)
//line views/components/Table.html:96
	qw422016.E().S(key)
//line views/components/Table.html:96
	qw422016.N().S(`"`)
//line views/components/Table.html:96
	streamtitleFor(qw422016, help)
//line views/components/Table.html:96
	qw422016.N().S(`>`)
//line views/components/Table.html:96
	qw422016.E().S(title)
//line views/components/Table.html:96
	qw422016.N().S(`</label></th>`)
//line views/components/Table.html:97
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:97
	qw422016.N().S(`<td>`)
//line views/components/Table.html:98
	StreamFormSelect(qw422016, key, "input-"+key, value, opts, titles, indent)
//line views/components/Table.html:98
	qw422016.N().S(`</td>`)
//line views/components/Table.html:99
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/Table.html:99
	qw422016.N().S(`</tr>`)
//line views/components/Table.html:101
}

//line views/components/Table.html:101
func WriteTableSelect(qq422016 qtio422016.Writer, key string, title string, value string, opts []string, titles []string, indent int, help ...string) {
//line views/components/Table.html:101
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Table.html:101
	StreamTableSelect(qw422016, key, title, value, opts, titles, indent, help...)
//line views/components/Table.html:101
	qt422016.ReleaseWriter(qw422016)
//line views/components/Table.html:101
}

//line views/components/Table.html:101
func TableSelect(key string, title string, value string, opts []string, titles []string, indent int, help ...string) string {
//line views/components/Table.html:101
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Table.html:101
	WriteTableSelect(qb422016, key, title, value, opts, titles, indent, help...)
//line views/components/Table.html:101
	qs422016 := string(qb422016.B)
//line views/components/Table.html:101
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Table.html:101
	return qs422016
//line views/components/Table.html:101
}

//line views/components/Table.html:103
func StreamTableDatalist(qw422016 *qt422016.Writer, key string, title string, value string, opts []string, titles []string, indent int, help ...string) {
//line views/components/Table.html:103
	qw422016.N().S(`<tr>`)
//line views/components/Table.html:105
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:105
	qw422016.N().S(`<th class="shrink"><label for="input-`)
//line views/components/Table.html:106
	qw422016.E().S(key)
//line views/components/Table.html:106
	qw422016.N().S(`"`)
//line views/components/Table.html:106
	streamtitleFor(qw422016, help)
//line views/components/Table.html:106
	qw422016.N().S(`>`)
//line views/components/Table.html:106
	qw422016.E().S(title)
//line views/components/Table.html:106
	qw422016.N().S(`</label></th>`)
//line views/components/Table.html:107
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:107
	qw422016.N().S(`<td>`)
//line views/components/Table.html:108
	StreamFormDatalist(qw422016, key, "input-"+key, value, opts, titles, indent)
//line views/components/Table.html:108
	qw422016.N().S(`</td>`)
//line views/components/Table.html:109
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/Table.html:109
	qw422016.N().S(`</tr>`)
//line views/components/Table.html:111
}

//line views/components/Table.html:111
func WriteTableDatalist(qq422016 qtio422016.Writer, key string, title string, value string, opts []string, titles []string, indent int, help ...string) {
//line views/components/Table.html:111
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Table.html:111
	StreamTableDatalist(qw422016, key, title, value, opts, titles, indent, help...)
//line views/components/Table.html:111
	qt422016.ReleaseWriter(qw422016)
//line views/components/Table.html:111
}

//line views/components/Table.html:111
func TableDatalist(key string, title string, value string, opts []string, titles []string, indent int, help ...string) string {
//line views/components/Table.html:111
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Table.html:111
	WriteTableDatalist(qb422016, key, title, value, opts, titles, indent, help...)
//line views/components/Table.html:111
	qs422016 := string(qb422016.B)
//line views/components/Table.html:111
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Table.html:111
	return qs422016
//line views/components/Table.html:111
}

//line views/components/Table.html:113
func StreamTableRadio(qw422016 *qt422016.Writer, key string, title string, value string, opts []string, titles []string, indent int, help ...string) {
//line views/components/Table.html:113
	qw422016.N().S(`<tr>`)
//line views/components/Table.html:115
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:115
	qw422016.N().S(`<th class="shrink"><label`)
//line views/components/Table.html:116
	streamtitleFor(qw422016, help)
//line views/components/Table.html:116
	qw422016.N().S(`>`)
//line views/components/Table.html:116
	qw422016.E().S(title)
//line views/components/Table.html:116
	qw422016.N().S(`</label></th>`)
//line views/components/Table.html:117
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:117
	qw422016.N().S(`<td>`)
//line views/components/Table.html:119
	StreamFormRadio(qw422016, key, value, opts, titles, indent+2)
//line views/components/Table.html:120
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:120
	qw422016.N().S(`</td>`)
//line views/components/Table.html:122
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/Table.html:122
	qw422016.N().S(`</tr>`)
//line views/components/Table.html:124
}

//line views/components/Table.html:124
func WriteTableRadio(qq422016 qtio422016.Writer, key string, title string, value string, opts []string, titles []string, indent int, help ...string) {
//line views/components/Table.html:124
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Table.html:124
	StreamTableRadio(qw422016, key, title, value, opts, titles, indent, help...)
//line views/components/Table.html:124
	qt422016.ReleaseWriter(qw422016)
//line views/components/Table.html:124
}

//line views/components/Table.html:124
func TableRadio(key string, title string, value string, opts []string, titles []string, indent int, help ...string) string {
//line views/components/Table.html:124
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Table.html:124
	WriteTableRadio(qb422016, key, title, value, opts, titles, indent, help...)
//line views/components/Table.html:124
	qs422016 := string(qb422016.B)
//line views/components/Table.html:124
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Table.html:124
	return qs422016
//line views/components/Table.html:124
}

//line views/components/Table.html:126
func StreamTableBoolean(qw422016 *qt422016.Writer, key string, title string, value bool, indent int, help ...string) {
//line views/components/Table.html:127
	StreamTableRadio(qw422016, key, title, fmt.Sprint(value), []string{"true", "false"}, []string{"True", "False"}, indent)
//line views/components/Table.html:128
}

//line views/components/Table.html:128
func WriteTableBoolean(qq422016 qtio422016.Writer, key string, title string, value bool, indent int, help ...string) {
//line views/components/Table.html:128
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Table.html:128
	StreamTableBoolean(qw422016, key, title, value, indent, help...)
//line views/components/Table.html:128
	qt422016.ReleaseWriter(qw422016)
//line views/components/Table.html:128
}

//line views/components/Table.html:128
func TableBoolean(key string, title string, value bool, indent int, help ...string) string {
//line views/components/Table.html:128
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Table.html:128
	WriteTableBoolean(qb422016, key, title, value, indent, help...)
//line views/components/Table.html:128
	qs422016 := string(qb422016.B)
//line views/components/Table.html:128
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Table.html:128
	return qs422016
//line views/components/Table.html:128
}

//line views/components/Table.html:130
func StreamTableCheckbox(qw422016 *qt422016.Writer, key string, title string, values []string, opts []string, titles []string, linebreaks bool, indent int, help ...string) {
//line views/components/Table.html:130
	qw422016.N().S(`<tr>`)
//line views/components/Table.html:132
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:132
	qw422016.N().S(`<th class="shrink"><label>`)
//line views/components/Table.html:133
	qw422016.E().S(title)
//line views/components/Table.html:133
	qw422016.N().S(`</label></th>`)
//line views/components/Table.html:134
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:134
	qw422016.N().S(`<td>`)
//line views/components/Table.html:136
	StreamFormCheckbox(qw422016, key, values, opts, titles, linebreaks, indent+2)
//line views/components/Table.html:137
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:137
	qw422016.N().S(`</td>`)
//line views/components/Table.html:139
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/Table.html:139
	qw422016.N().S(`</tr>`)
//line views/components/Table.html:141
}

//line views/components/Table.html:141
func WriteTableCheckbox(qq422016 qtio422016.Writer, key string, title string, values []string, opts []string, titles []string, linebreaks bool, indent int, help ...string) {
//line views/components/Table.html:141
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Table.html:141
	StreamTableCheckbox(qw422016, key, title, values, opts, titles, linebreaks, indent, help...)
//line views/components/Table.html:141
	qt422016.ReleaseWriter(qw422016)
//line views/components/Table.html:141
}

//line views/components/Table.html:141
func TableCheckbox(key string, title string, values []string, opts []string, titles []string, linebreaks bool, indent int, help ...string) string {
//line views/components/Table.html:141
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Table.html:141
	WriteTableCheckbox(qb422016, key, title, values, opts, titles, linebreaks, indent, help...)
//line views/components/Table.html:141
	qs422016 := string(qb422016.B)
//line views/components/Table.html:141
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Table.html:141
	return qs422016
//line views/components/Table.html:141
}

//line views/components/Table.html:143
func StreamTableIcons(qw422016 *qt422016.Writer, key string, title string, value string, ps *cutil.PageState, indent int, help ...string) {
//line views/components/Table.html:143
	qw422016.N().S(`<tr>`)
//line views/components/Table.html:145
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:145
	qw422016.N().S(`<th class="shrink"><label>`)
//line views/components/Table.html:146
	qw422016.E().S(title)
//line views/components/Table.html:146
	qw422016.N().S(`</label></th>`)
//line views/components/Table.html:147
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:147
	qw422016.N().S(`<td>`)
//line views/components/Table.html:148
	StreamIconPicker(qw422016, value, key, ps, indent+2)
//line views/components/Table.html:148
	qw422016.N().S(`</td>`)
//line views/components/Table.html:150
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/Table.html:150
	qw422016.N().S(`</tr>`)
//line views/components/Table.html:152
}

//line views/components/Table.html:152
func WriteTableIcons(qq422016 qtio422016.Writer, key string, title string, value string, ps *cutil.PageState, indent int, help ...string) {
//line views/components/Table.html:152
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Table.html:152
	StreamTableIcons(qw422016, key, title, value, ps, indent, help...)
//line views/components/Table.html:152
	qt422016.ReleaseWriter(qw422016)
//line views/components/Table.html:152
}

//line views/components/Table.html:152
func TableIcons(key string, title string, value string, ps *cutil.PageState, indent int, help ...string) string {
//line views/components/Table.html:152
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Table.html:152
	WriteTableIcons(qb422016, key, title, value, ps, indent, help...)
//line views/components/Table.html:152
	qs422016 := string(qb422016.B)
//line views/components/Table.html:152
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Table.html:152
	return qs422016
//line views/components/Table.html:152
}

//line views/components/Table.html:154
func streamtitleFor(qw422016 *qt422016.Writer, help []string) {
//line views/components/Table.html:155
	if len(help) > 0 {
//line views/components/Table.html:155
		qw422016.N().S(` `)
//line views/components/Table.html:155
		qw422016.N().S(`title="`)
//line views/components/Table.html:155
		qw422016.E().S(strings.Join(help, "; "))
//line views/components/Table.html:155
		qw422016.N().S(`"`)
//line views/components/Table.html:155
	}
//line views/components/Table.html:156
}

//line views/components/Table.html:156
func writetitleFor(qq422016 qtio422016.Writer, help []string) {
//line views/components/Table.html:156
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Table.html:156
	streamtitleFor(qw422016, help)
//line views/components/Table.html:156
	qt422016.ReleaseWriter(qw422016)
//line views/components/Table.html:156
}

//line views/components/Table.html:156
func titleFor(help []string) string {
//line views/components/Table.html:156
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Table.html:156
	writetitleFor(qb422016, help)
//line views/components/Table.html:156
	qs422016 := string(qb422016.B)
//line views/components/Table.html:156
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Table.html:156
	return qs422016
//line views/components/Table.html:156
}
