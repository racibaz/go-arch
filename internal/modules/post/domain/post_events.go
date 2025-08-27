package domain

const (
	PostCreatedEvent = "posts.PostCreated"
)

type PostCreated struct {
	Post *Post
}

func (PostCreated) EventName() string { return "post.PostCreated" }

type PostDeleted struct {
	Post *Post
}

func (PostDeleted) EventName() string { return "post.PostDeleted" }
