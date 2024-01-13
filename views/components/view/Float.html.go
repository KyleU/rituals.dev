// Code generated by qtc from "Float.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/components/view/Float.html:2
package view

//line views/components/view/Float.html:2
import "github.com/kyleu/rituals/app/util"

//line views/components/view/Float.html:4
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/view/Float.html:4
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/view/Float.html:4
func StreamFloat(qw422016 *qt422016.Writer, f float64) {
//line views/components/view/Float.html:5
	qw422016.E().V(f)
//line views/components/view/Float.html:6
}

//line views/components/view/Float.html:6
func WriteFloat(qq422016 qtio422016.Writer, f float64) {
//line views/components/view/Float.html:6
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/view/Float.html:6
	StreamFloat(qw422016, f)
//line views/components/view/Float.html:6
	qt422016.ReleaseWriter(qw422016)
//line views/components/view/Float.html:6
}

//line views/components/view/Float.html:6
func Float(f float64) string {
//line views/components/view/Float.html:6
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/view/Float.html:6
	WriteFloat(qb422016, f)
//line views/components/view/Float.html:6
	qs422016 := string(qb422016.B)
//line views/components/view/Float.html:6
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/view/Float.html:6
	return qs422016
//line views/components/view/Float.html:6
}

//line views/components/view/Float.html:8
func StreamFloatArray(qw422016 *qt422016.Writer, value []float64) {
//line views/components/view/Float.html:9
	StreamStringArray(qw422016, util.ArrayToStringArray(value))
//line views/components/view/Float.html:10
}

//line views/components/view/Float.html:10
func WriteFloatArray(qq422016 qtio422016.Writer, value []float64) {
//line views/components/view/Float.html:10
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/view/Float.html:10
	StreamFloatArray(qw422016, value)
//line views/components/view/Float.html:10
	qt422016.ReleaseWriter(qw422016)
//line views/components/view/Float.html:10
}

//line views/components/view/Float.html:10
func FloatArray(value []float64) string {
//line views/components/view/Float.html:10
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/view/Float.html:10
	WriteFloatArray(qb422016, value)
//line views/components/view/Float.html:10
	qs422016 := string(qb422016.B)
//line views/components/view/Float.html:10
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/view/Float.html:10
	return qs422016
//line views/components/view/Float.html:10
}
