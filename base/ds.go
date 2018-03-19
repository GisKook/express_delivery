package base

type User struct {
	ID    string `json:"id"`
	Alias string `json:"alias"`
	Role  string `json:"role"`
}

type DBResult struct {
	Err   error
	Extra interface{}
}
