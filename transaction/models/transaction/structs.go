package transaction

type UpdateModel struct {
	IdUser    int    `json:"id_users"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
