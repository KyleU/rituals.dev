// Code generated by qtc from "standup.sql". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// -- Content managed by Project Forge, see [projectforge.md] for details.
// --

//line queries/ddl/standup.sql:2
package ddl

//line queries/ddl/standup.sql:2
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line queries/ddl/standup.sql:2
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line queries/ddl/standup.sql:2
func StreamStandupDrop(qw422016 *qt422016.Writer) {
//line queries/ddl/standup.sql:2
	qw422016.N().S(`
drop table if exists "standup";
-- `)
//line queries/ddl/standup.sql:4
}

//line queries/ddl/standup.sql:4
func WriteStandupDrop(qq422016 qtio422016.Writer) {
//line queries/ddl/standup.sql:4
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/ddl/standup.sql:4
	StreamStandupDrop(qw422016)
//line queries/ddl/standup.sql:4
	qt422016.ReleaseWriter(qw422016)
//line queries/ddl/standup.sql:4
}

//line queries/ddl/standup.sql:4
func StandupDrop() string {
//line queries/ddl/standup.sql:4
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/ddl/standup.sql:4
	WriteStandupDrop(qb422016)
//line queries/ddl/standup.sql:4
	qs422016 := string(qb422016.B)
//line queries/ddl/standup.sql:4
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/ddl/standup.sql:4
	return qs422016
//line queries/ddl/standup.sql:4
}

// --

//line queries/ddl/standup.sql:6
func StreamStandupCreate(qw422016 *qt422016.Writer) {
//line queries/ddl/standup.sql:6
	qw422016.N().S(`
create table if not exists "standup" (
  "id" uuid not null,
  "slug" text not null,
  "title" text not null,
  "status" session_status not null,
  "team_id" uuid,
  "sprint_id" uuid,
  "owner" uuid not null,
  "created" timestamp not null default now(),
  "updated" timestamp default now(),
  foreign key ("owner") references "user" ("id"),
  foreign key ("team_id") references "team" ("id"),
  foreign key ("sprint_id") references "sprint" ("id"),
  unique ("slug"),
  primary key ("id")
);

create index if not exists "standup__slug_idx" on "standup" ("slug");

create index if not exists "standup__status_idx" on "standup" ("status");

create index if not exists "standup__owner_idx" on "standup" ("owner");

create index if not exists "standup__team_id_idx" on "standup" ("team_id");

create index if not exists "standup__sprint_id_idx" on "standup" ("sprint_id");
-- `)
//line queries/ddl/standup.sql:33
}

//line queries/ddl/standup.sql:33
func WriteStandupCreate(qq422016 qtio422016.Writer) {
//line queries/ddl/standup.sql:33
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/ddl/standup.sql:33
	StreamStandupCreate(qw422016)
//line queries/ddl/standup.sql:33
	qt422016.ReleaseWriter(qw422016)
//line queries/ddl/standup.sql:33
}

//line queries/ddl/standup.sql:33
func StandupCreate() string {
//line queries/ddl/standup.sql:33
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/ddl/standup.sql:33
	WriteStandupCreate(qb422016)
//line queries/ddl/standup.sql:33
	qs422016 := string(qb422016.B)
//line queries/ddl/standup.sql:33
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/ddl/standup.sql:33
	return qs422016
//line queries/ddl/standup.sql:33
}
