package domain

import "fmt"

type UserStatus int

const (
	StatusDraft UserStatus = iota
	StatusPublished
	StatusArchived
	StatusPendingReview
	StatusBanned
)

var StatusToString = map[UserStatus]string{
	StatusDraft:         "draft",
	StatusPublished:     "published",
	StatusArchived:      "archived",
	StatusPendingReview: "pending_review",
	StatusBanned:        "banned",
}

var stringToStatus = map[string]UserStatus{
	"draft":          StatusDraft,
	"published":      StatusPublished,
	"archived":       StatusArchived,
	"pending_review": StatusPendingReview,
	"banned":         StatusBanned,
}

func NewStatus(status UserStatus) UserStatus {
	isValidStatus := IsValidStatus(status)
	if !isValidStatus {
		panic(fmt.Sprintf("Invalid status: %s", status))
	}

	return status
}

func (p UserStatus) String() string {
	if val, ok := StatusToString[p]; ok {
		return val
	}
	return "unknown"
}

func (p UserStatus) ToInt() int {
	return int(p)
}

func IsValidStatus(p UserStatus) bool {
	_, ok := StatusToString[p]
	return ok
}

func (p UserStatus) EqualTo(other UserStatus) bool {
	return p == other
}

func (p UserStatus) IsPublished() bool {
	return p == StatusPublished
}

func (p UserStatus) IsDraft() bool {
	return p == StatusDraft
}

func (p UserStatus) IsArchived() bool {
	return p == StatusArchived
}

func (p UserStatus) IsPendingReview() bool {
	return p == StatusPendingReview
}

func (p UserStatus) IsBanned() bool {
	return p == StatusBanned
}
