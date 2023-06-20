package Models

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string
}

func (u *User) SetEmail(email string) {
	u.Email = email
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) SetFirstName(firstName string) {
	u.FirstName = firstName
}

func (u *User) SetLastName(lastName string) {
	u.LastName = lastName
}

func (u *User) getFullName() string {
	return u.FirstName + " " + u.LastName
}

func (u *User) SetPassword(password string) {
	u.Password = password
}
