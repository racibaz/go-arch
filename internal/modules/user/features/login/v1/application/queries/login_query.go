package queries

type LoginQueryV1 struct {
	Email    string
	Password string
	Platform string
}

type LoginQueryResponse struct {
	AccessToken  string
	RefreshToken string
	UserID       string
}
