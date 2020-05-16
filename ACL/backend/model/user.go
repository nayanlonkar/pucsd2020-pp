package model

type User struct {
	Id        int64  `json:"user_id,omitempty" key:"primary" autoincr:"1" column:"id"`
	FirstName string `json:"first_name" column:"first_name"`
	LastName  string `json:"last_name" column:"last_name"`
	UserName  string `json:"username" column:"username"`
	Password  string `json:"password" column:"password"`
}

func (user *User) Table() string {
	return "users"
}

func (user *User) String() string {
	return Stringify(user)
}
