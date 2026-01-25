package domain

const (
	UserRegisteredEvent = "user.UserRegistered"
	UserDeletedEvent    = "user.UserDeleted"
)

type UserRegistered struct {
	User *User
}

func (UserRegistered) EventName() string { return UserRegisteredEvent }

type UserDeleted struct {
	User *User
}

func (UserDeleted) EventName() string { return UserDeletedEvent }
