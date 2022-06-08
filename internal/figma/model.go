package figma

type Comment struct {
	Id string `json:"id"`
	User User `json:"user"`
}

type User struct {
	Email string `json:"email"`
	Handle string `json:"handle"`
}
