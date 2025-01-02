## example graphql queries

### fetch all authors

```gql
query {
  authors {
    id
    name
    email
    bio
    posts {
      id
      title
      created_at
    }
    comments {
      id
      content
      created_at
      post {
        id
        title
      }
    }
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
    author {
      id
      name
      email
    }
    comments {
      id
      content
      created_at
      author {
        id
        name
      }
    }
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
    author {
      id
      name
    }
    post {
      id
      title
    }
  }
}
```

### create author 
```gql
mutation {
  insert_authors(objects: {
    name: "john doe",
    email: "johndoe@example.com",
    bio: "an author who loves writing about technology."
  }) {
    returning {
      id
      name
      email
    }
  }
}
```

### create post

```gql
mutation {
  insert_posts(objects: {
    title: "graphql and go",
    content: "an introduction to using graphql with go using graphjin.",
    author_id: 1
  }) {
    returning {
      id
      title
      created_at
    }
  }
}
```

### create comment 

```gql
mutation {
  insert_comments(objects: {
    content: "this post was super helpful!",
    post_id: 1,
    author_id: 2
  }) {
    returning {
      id
      content
      created_at
    }
  }
}
```

### multiple posts at once 

```gql
mutation {
  insert_posts(objects: [
    {title: "introduction to go", content: "go is a statically typed, compiled language.", author_id: 1},
    {title: "building rest apis", content: "learn how to build restful apis with go.", author_id: 1}
  ]) {
    returning {
      id
      title
    }
  }
}
```

### multiple comments at once

```gql
mutation {
  insert_comments(objects: [
    {content: "great post!", post_id: 1, author_id: 2},
    {content: "i learned a lot from this.", post_id: 2, author_id: 3}
  ]) {
    returning {
      id
      content
    }
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