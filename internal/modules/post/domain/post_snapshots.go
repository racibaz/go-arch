// Package domain defines the domain models and logic for the post module.

package domain

type PostV1 struct {
	UserID  string
	Title   string
	Content string
	Status  PostStatus
}

func (PostV1) SnapshotName() string { return "posts.PostV1" }
