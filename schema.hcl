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
  
  primary_key {
    columns = [column.id]
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
  
  primary_key {
    columns = [column.id]
  }
  
  foreign_key "author_fk" {
    columns = [column.author_id]
    ref_columns = [table.authors.column.id]
    on_delete = CASCADE
    on_update = CASCADE
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
  
  primary_key {
    columns = [column.id]
  }
  
  foreign_key "post_fk" {
    columns = [column.post_id]
    ref_columns = [table.posts.column.id]
    on_delete = CASCADE
    on_update = CASCADE
  }
  
  foreign_key "author_fk" {
    columns = [column.author_id]
    ref_columns = [table.authors.column.id]
    on_delete = CASCADE
    on_update = CASCADE
  }
}

schema "public" {
  comment = "schema for blog website"
}

schema "private" {}