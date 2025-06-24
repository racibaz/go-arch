package domain

import (
	"fmt"
)

type PostStatus int

const (
	// PostStatusDraft @Description PostStatusDraft represents the status of a post in the system.
	PostStatusDraft PostStatus = iota
	// PostStatusPublished @Description PostStatusPublished represents the status of a post in the system.
	PostStatusPublished
	// PostStatusArchived @Description PostStatusArchived represents the status of a post in the system.
	PostStatusArchived
)

var postStatusToString = map[PostStatus]string{
	PostStatusDraft:     "draft",
	PostStatusPublished: "published",
	PostStatusArchived:  "archived",
}

var stringToPostStatus = map[string]PostStatus{
	"draft":     PostStatusDraft,
	"published": PostStatusPublished,
	"archived":  PostStatusArchived,
}

func NewPostStatus(status PostStatus) PostStatus {

	isValidPostStatus := IsValidPostStatus(status)
	if !isValidPostStatus {
		panic(fmt.Sprintf("Invalid post status: %s", status))
	}

	return status
}

func (p PostStatus) String() string {
	if val, ok := postStatusToString[p]; ok {
		return val
	}
	return "unknown"
}

func IsValidPostStatus(p PostStatus) bool {
	_, ok := postStatusToString[p]
	return ok
}

func (p PostStatus) EqualTo(other PostStatus) bool {
	return p == other
}

func (p PostStatus) IsPublished() bool {
	return p == PostStatusPublished
}

func (p PostStatus) IsDraft() bool {
	return p == PostStatusDraft
}

func (p PostStatus) IsArchived() bool {
	return p == PostStatusArchived
}

func (p PostStatus) ToInt() (PostStatus, error) {
	if val, ok := stringToPostStatus[p.String()]; ok {
		return val, nil
	}
	return -1, fmt.Errorf("invalid post status: %s", p.String())
}
