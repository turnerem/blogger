package blogger_test

import (
  blogger "github.com/turnerem/blogger"
  "testing"
  "testing/fstest"
  "io/fs"
  "errors"
  "reflect"
)

func TestNewBlogPosts(t *testing.T) {
  t.Run("returns all blog posts", func(t *testing.T) {
    fs := fstest.MapFS{
      "hellowrld.md":	{Data: []byte("Title: Post 1")},
      "hey_world2.md":	{Data: []byte("Title: Post 2")},
    }

    posts, err := blogger.NewPostsFromFS(fs)

    if err != nil {
      t.Fatal(err)
    }

    if len(posts) != len(fs) {
      t.Errorf("got %d posts but wanted %d", len(posts), len(fs))
    }

    got := posts[0]
    want := blogger.Post{Title: "Post 1"}

    if !reflect.DeepEqual(got, want) {
      t.Errorf("got %+v, want %+v", got, want)
    }
  })

  t.Run("returns error when can't open file", func(t *testing.T) {
    _, err := blogger.NewPostsFromFS(StubFailingFS{})

    if err == nil {
      t.Errorf("expected an error")
    }
  })
}


type StubFailingFS struct {
}

func (s StubFailingFS) Open(name string) (fs.File, error) {
  return nil, errors.New("uh oh")
}
