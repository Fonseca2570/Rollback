package transaction

type UpdateModel struct {
	IdUser    int    `json:"id_users"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type UpdateResponseServiceOne struct {
	Error     string
	TableName string
	RowId     int64
}

type UpdateResponseServiceTwo struct {
	Error     string
	TableName string
	RowId     int64
}

type UpdateResponseServiceThree struct {
	Error     string
	TableName string
	RowId     int64
}