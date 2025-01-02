## example graphql queries

### fetch all authors

```gql
query {
  authors {
    id
    name
    email
    bio
  }
}
```

### fetch all posts with their authors and comments

```gql
query {
  posts {
    id
    title
    content
    created_at
    author_id
  }
}
```

### fetch comments for a specific post

```gql
query {
  comments(where: {post_id: {eq: 1}}) {
    id
    content
    created_at
    author_id
    post_id
  }
}
```

### create author 
```gql
mutation {
  authors(insert: {
    name: "john doe",
    email: "johndoe@example.com",
    bio: "an author who loves writing about technology."
  }) {
      id
      name
      email
    }
}
```

### create post

```gql
mutation {
  posts(insert: {
    title: "graphql and go",
    content: "an introduction to using graphql with go using graphjin.",
    author_id: 1
  }) {
    id
    title
    created_at
  }
}
```

### create comment 

```gql
mutation {
  comments(insert: {
    content: "this post was super helpful!",
    post_id: 1,
    author_id: 2
  }) {
    id
    content
    created_at
  }
}
```

### multiple posts at once 

```gql
mutation {
  posts(insert: [
    {title: "introduction to go", content: "go is a statically typed, compiled language.", author_id: 1},
    {title: "building rest apis", content: "learn how to build restful apis with go.", author_id: 1}
  ]) {
    id
    title
  }
}
```

### multiple comments at once

```gql
mutation {
  comments(insert: [
    {content: "great post!", post_id: 1, author_id: 2},
    {content: "i learned a lot from this.", post_id: 2, author_id: 3}
  ]) {
    id
    content
  }
}

```

### create author with posts

```gql
mutation {
  insert_authors(objects: {
    name: "jane doe",
    email: "janedoe@example.com",
    bio: "a passionate author.",
    posts: {
      data: [
        {title: "getting started with go", content: "basics of go programming."},
        {title: "advanced go concepts", content: "diving deep into go."}
      ]
    }
  }) {
    returning {
      id
      name
      posts {
        id
        title
      }
    }
  }
}
```