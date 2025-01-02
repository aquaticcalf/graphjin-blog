## graphjin blog

### what i did so far

- installed graphjin

```sh
go install github.com/dosco/graphjin@latest
```

- had to alias graphjin, because it aint working 😠

```sh
alias graphjin='$(go env GOPATH)/bin/graphjin'
```

- created a new "blog" app

```sh
graphjin new blog
```

- installed postgresql

```sh
sudo pacman -S postgresql
```

- installed atlas 

```sh
curl -sSf https://atlasgo.sh | sh
```

- initialise a database cluster

```sh
sudo -u postgres initdb --locale en_US.UTF-8 -D /var/lib/postgres/data
```

- start and enable the postgres service

```sh
sudo systemctl start postgresql
sudo systemctl enable postgresql
```

- create database 

```sh
sudo -u postgres createdb blog_development
```

- create a new postgres user

```sh
sudo -u postgres createuser --interactive
```

- grant privileges to the new user

```sh
 sudo -u postgres psql
psql (16.6)
Type "help" for help.

postgres=# GRANT ALL PRIVILEGES ON DATABASE blog_development TO sam;
GRANT
postgres=# \q
```
> GRANT ALL PRIVILEGES ON DATABASE blog_development TO sam;
> if successful, it should display GRANT

( `sam` is the `username` of the user i just created )

- create `schema.hcl` file

```sh
touch schema.hcl
```

- fill it with the following
```hcl
schema "public" {
  comment = "test: schema comment"
}

schema "private" {}
```

- migrate that

```sh
atlas schema apply --url "postgres://sam:password@localhost:5432/blog_development?sslmode=disable" --to "file://schema.hcl"
```

- it shows error 

```sh
Planning migration statements (2 in total):

  -- set comment to schema: "public":
    -> COMMENT ON SCHEMA "public" IS 'test: schema comment';
  -- add new schema named "private":
    -> CREATE SCHEMA "private";

-------------------------------------------

Applying approved migration (2 statements in total):

  -- set comment to schema: "public"
    -> COMMENT ON SCHEMA "public" IS 'test: schema comment';
    pq: must be owner of schema public

  -------------------------
  -- 1.105405ms
  -- 1 migration with errors
  -- 1 sql statement with errors
Error: executing statement "COMMENT ON SCHEMA \"public\" IS 'test: schema comment';": pq: must be owner of schema public
```

- this is because `sam` isnt the owner of public schema *yet*

```sh
sudo -u postgres psql -d blog_development

psql (16.6)
Type "help" for help.

blog_development=# ALTER SCHEMA public OWNER TO sam;
ALTER SCHEMA
blog_development=# \q
```
- try migrating again 

```sh
 atlas schema apply --url "postgres://sam:password@localhost:5432/blog_development?sslmode=disable" --to "file://schema.hcl"

Planning migration statements (2 in total):

  -- set comment to schema: "public":
    -> COMMENT ON SCHEMA "public" IS 'test: schema comment';
  -- add new schema named "private":
    -> CREATE SCHEMA "private";

-------------------------------------------
Applying approved migration (2 statements in total):

  -- set comment to schema: "public"
    -> COMMENT ON SCHEMA "public" IS 'test: schema comment';
  -- ok (540.662µs)
  -- add new schema named "private"
    -> CREATE SCHEMA "private";
  -- ok (581.833µs)

  -------------------------                                                                          -- 1.210299ms
  -- 1 migration
  -- 2 sql statements
``` 

> it works yayy

- change this in `config/scripts/dev.yml`

```yml
database:
    type: postgres
    host: localhost
    port: 5432
    dbname: blog_development
    user: sam
    password: password
    schema: "public"
```

- write basic blog schema in schema.hcl

```hcl
table "authors" {
  schema = schema.public
  column "id" {
    null = false
    type = sql("serial")
  }
  column "name" {
    null = false
    type = sql("character varying(255)")
  }
  column "email" {
    null = false
    type = sql("character varying(255)")
  }
  column "bio" {
    null = true
    type = sql("text")
  }
}

table "posts" {
  schema = schema.public
  column "id" {
    null = false
    type = sql("serial")
  }
  column "title" {
    null = false
    type = sql("text")
  }
  column "content" {
    null = false
    type = sql("text")
  }
  column "author_id" {
    null = false
    type = sql("integer")
  }
  column "created_at" {
    null = false
    type = sql("timestamp")
  }
}

table "comments" {
  schema = schema.public
  column "id" {
    null = false
    type = sql("serial")
  }
  column "post_id" {
    null = false
    type = sql("integer")
  }
  column "author_id" {
    null = false
    type = sql("integer")
  }
  column "content" {
    null = false
    type = sql("text")
  }
  column "created_at" {
    null = false
    type = sql("timestamp")
  }
}

schema "public" {
  comment = "schema for blog website"
}

schema "private" {}
```

- migrations

```sh
atlas schema apply --url "postgres://sam:password@localhost:5432/blog_development?sslmode=disable" --to "file://schema.hcl"
```

```sh
Planning migration statements (4 in total):

  -- set comment to schema: "public":
    -> COMMENT ON SCHEMA "public" IS 'schema for blog website';
  -- create "authors" table:
    -> CREATE TABLE "public"."authors" (
         "id" serial NOT NULL,
         "name" character varying(255) NOT NULL,
         "email" character varying(255) NOT NULL,
         "bio" text NULL
       );
  -- create "posts" table:
    -> CREATE TABLE "public"."posts" (
         "id" serial NOT NULL,
         "title" text NOT NULL,
         "content" text NOT NULL,
         "author_id" integer NOT NULL,
         "created_at" timestamp NOT NULL
       );
  -- create "comments" table:
    -> CREATE TABLE "public"."comments" (
         "id" serial NOT NULL,
         "post_id" integer NOT NULL,
         "author_id" integer NOT NULL,
         "content" text NOT NULL,
         "created_at" timestamp NOT NULL
       );

-------------------------------------------

Applying approved migration (4 statements in total):

  -- set comment to schema: "public"
    -> COMMENT ON SCHEMA "public" IS 'schema for blog website';
  -- ok (498.846µs)

  -- create "authors" table
    -> CREATE TABLE "public"."authors" (
         "id" serial NOT NULL,
         "name" character varying(255) NOT NULL,
         "email" character varying(255) NOT NULL,
         "bio" text NULL
       );
  -- ok (5.388393ms)

  -- create "posts" table
    -> CREATE TABLE "public"."posts" (
         "id" serial NOT NULL,
         "title" text NOT NULL,
         "content" text NOT NULL,
         "author_id" integer NOT NULL,
         "created_at" timestamp NOT NULL
       );
  -- ok (3.874735ms)

  -- create "comments" table
    -> CREATE TABLE "public"."comments" (
         "id" serial NOT NULL,
         "post_id" integer NOT NULL,
         "author_id" integer NOT NULL,
         "content" text NOT NULL,
         "created_at" timestamp NOT NULL
       );
  -- ok (3.114995ms)

  -------------------------
  -- 13.391806ms
  -- 1 migration
  -- 4 sql statements
```

