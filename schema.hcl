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
    default = "now"
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
    default = "now"
  }
}

schema "public" {
  comment = "schema for blog website"
}

schema "private" {}
