package query

type RefreshTokenQueryV1 struct {
	RefreshToken string
	Platform     string
}

type RefreshTokenQueryResponseV1 struct {
	AccessToken  string
	RefreshToken string
	UserID       string
}
