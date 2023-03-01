package userentity

type User struct {
	ID        int
	UID       string
	Name      string
	Thumbnail string
}

type UserSlice []*User
