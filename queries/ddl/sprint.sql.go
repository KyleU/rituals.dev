// Code generated by qtc from "sprint.sql". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// -- Content managed by Project Forge, see [projectforge.md] for details.
// --

//line queries/ddl/sprint.sql:2
package ddl

//line queries/ddl/sprint.sql:2
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line queries/ddl/sprint.sql:2
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line queries/ddl/sprint.sql:2
func StreamSprintDrop(qw422016 *qt422016.Writer) {
//line queries/ddl/sprint.sql:2
	qw422016.N().S(`
drop table if exists "sprint";
-- `)
//line queries/ddl/sprint.sql:4
}

//line queries/ddl/sprint.sql:4
func WriteSprintDrop(qq422016 qtio422016.Writer) {
//line queries/ddl/sprint.sql:4
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/ddl/sprint.sql:4
	StreamSprintDrop(qw422016)
//line queries/ddl/sprint.sql:4
	qt422016.ReleaseWriter(qw422016)
//line queries/ddl/sprint.sql:4
}

//line queries/ddl/sprint.sql:4
func SprintDrop() string {
//line queries/ddl/sprint.sql:4
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/ddl/sprint.sql:4
	WriteSprintDrop(qb422016)
//line queries/ddl/sprint.sql:4
	qs422016 := string(qb422016.B)
//line queries/ddl/sprint.sql:4
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/ddl/sprint.sql:4
	return qs422016
//line queries/ddl/sprint.sql:4
}

// --

//line queries/ddl/sprint.sql:6
func StreamSprintCreate(qw422016 *qt422016.Writer) {
//line queries/ddl/sprint.sql:6
	qw422016.N().S(`
create table if not exists "sprint" (
  "id" uuid not null,
  "slug" text not null,
  "title" text not null,
  "status" session_status not null,
  "team_id" uuid,
  "owner" uuid not null,
  "start_date" timestamp,
  "end_date" timestamp,
  "created" timestamp not null default now(),
  "updated" timestamp default now(),
  foreign key ("owner") references "user" ("id"),
  foreign key ("team_id") references "team" ("id"),
  unique ("slug"),
  primary key ("id")
);

create index if not exists "sprint__slug_idx" on "sprint" ("slug");

create index if not exists "sprint__status_idx" on "sprint" ("status");

create index if not exists "sprint__owner_idx" on "sprint" ("owner");

create index if not exists "sprint__team_id_idx" on "sprint" ("team_id");
-- `)
//line queries/ddl/sprint.sql:31
}

//line queries/ddl/sprint.sql:31
func WriteSprintCreate(qq422016 qtio422016.Writer) {
//line queries/ddl/sprint.sql:31
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/ddl/sprint.sql:31
	StreamSprintCreate(qw422016)
//line queries/ddl/sprint.sql:31
	qt422016.ReleaseWriter(qw422016)
//line queries/ddl/sprint.sql:31
}

//line queries/ddl/sprint.sql:31
func SprintCreate() string {
//line queries/ddl/sprint.sql:31
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/ddl/sprint.sql:31
	WriteSprintCreate(qb422016)
//line queries/ddl/sprint.sql:31
	qs422016 := string(qb422016.B)
//line queries/ddl/sprint.sql:31
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/ddl/sprint.sql:31
	return qs422016
//line queries/ddl/sprint.sql:31
}
