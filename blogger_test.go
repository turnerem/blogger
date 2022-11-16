package blogger_test

import (
  "testing"
  "testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
  fs := fstest.MapFS{
    "hellowrld.md":	{Data: []byte("hi")},
    "hey_worldie.md":	{Data: []byte("hiya")},
  }

  posts := blogger.NewPostsFromFS(fs)

  if len(posts) != len(fs) {
    t.Errorf("got %d posts but wanted %d", len(posts), len(fs))
  }
}
