package domain

type PostCreated struct {
	Post *Post
}

func (PostCreated) EventName() string { return "post.PostCreated" }

type PostDeleted struct {
	Post *Post
}

func (PostDeleted) EventName() string { return "post.PostDeleted" }
