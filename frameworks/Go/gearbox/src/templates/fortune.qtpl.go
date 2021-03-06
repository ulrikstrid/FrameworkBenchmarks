// Code generated by qtc from "fortune.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line fortune.qtpl:1
package templates

//line fortune.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line fortune.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line fortune.qtpl:1
func StreamFortunePage(qw422016 *qt422016.Writer, rows []Fortune) {
//line fortune.qtpl:1
	qw422016.N().S(`<!DOCTYPE html>
<html>
<head>
<title>Fortunes</title>
</head>
<body>
<table>
<tr><th>id</th><th>message</th></tr>
`)
//line fortune.qtpl:9
	for _, r := range rows {
//line fortune.qtpl:9
		qw422016.N().S(`
<tr><td>`)
//line fortune.qtpl:10
		qw422016.N().D(int(r.ID))
//line fortune.qtpl:10
		qw422016.N().S(`</td><td>`)
//line fortune.qtpl:10
		qw422016.E().S(r.Message)
//line fortune.qtpl:10
		qw422016.N().S(`</td></tr>
`)
//line fortune.qtpl:11
	}
//line fortune.qtpl:11
	qw422016.N().S(`
</table>
</body>
</html>
`)
//line fortune.qtpl:15
}

//line fortune.qtpl:15
func WriteFortunePage(qq422016 qtio422016.Writer, rows []Fortune) {
//line fortune.qtpl:15
	qw422016 := qt422016.AcquireWriter(qq422016)
//line fortune.qtpl:15
	StreamFortunePage(qw422016, rows)
//line fortune.qtpl:15
	qt422016.ReleaseWriter(qw422016)
//line fortune.qtpl:15
}

//line fortune.qtpl:15
func FortunePage(rows []Fortune) string {
//line fortune.qtpl:15
	qb422016 := qt422016.AcquireByteBuffer()
//line fortune.qtpl:15
	WriteFortunePage(qb422016, rows)
//line fortune.qtpl:15
	qs422016 := string(qb422016.B)
//line fortune.qtpl:15
	qt422016.ReleaseByteBuffer(qb422016)
//line fortune.qtpl:15
	return qs422016
//line fortune.qtpl:15
}
