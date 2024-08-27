package entity

type User struct {
	ID           string
	Username     string
	Password     string
	NIM          string
	Email        string
	Faculty      string
	StudyProgram string
	Role         UserRole
}

type UserRole uint8

const (
	UserRoleUnknown    UserRole = 0
	UserRoleAdmin      UserRole = 1
	UserRoleDelegation UserRole = 2
	UserRoleStudent    UserRole = 3
)

var UserRolesMap = map[UserRole]string{
	UserRoleAdmin:      "Admin",
	UserRoleDelegation: "Delegation",
	UserRoleStudent:    "Student",
}

func (u UserRole) String() string {
	return UserRolesMap[u]
}

func (u UserRole) Int() int {
	return int(u)
}
