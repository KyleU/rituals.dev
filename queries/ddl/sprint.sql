-- Content managed by Project Forge, see [projectforge.md] for details.
-- {% func SprintDrop() %}
drop table if exists "sprint";
-- {% endfunc %}

-- {% func SprintCreate() %}
create table if not exists "sprint" (
  "id" uuid not null,
  "slug" text not null,
  "title" text not null,
  "status" session_status not null,
  "team_id" uuid,
  "owner" uuid not null,
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
-- {% endfunc %}
