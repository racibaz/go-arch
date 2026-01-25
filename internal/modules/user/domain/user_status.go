package domain

import "fmt"

type UserStatus int

const (
	StatusDraft UserStatus = iota
	StatusPublished
	StatusArchived
)

var postStatusToString = map[UserStatus]string{
	StatusDraft:     "draft",
	StatusPublished: "published",
	StatusArchived:  "archived",
}

var stringToStatus = map[string]UserStatus{
	"draft":     StatusDraft,
	"published": StatusPublished,
	"archived":  StatusArchived,
}

func NewStatus(status UserStatus) UserStatus {
	isValidStatus := IsValidStatus(status)
	if !isValidStatus {
		panic(fmt.Sprintf("Invalid status: %s", status))
	}

	return status
}

func (p UserStatus) String() string {
	if val, ok := postStatusToString[p]; ok {
		return val
	}
	return "unknown"
}

func IsValidStatus(p UserStatus) bool {
	_, ok := postStatusToString[p]
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
