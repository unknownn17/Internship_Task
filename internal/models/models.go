package models

type Register_User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Verify_User struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

type LogIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Task struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Title       string `json:"title"`
	Status      string `json:"status"`
	Description string `json:"desciption"`
	Important   bool   `json:"important"`
}

type GetTaskRequest struct {
	ID     int `json:"id" bson:"id"`
	UserID int `json:"user_id" bson:"user_id"`
}

type GetTaskResponse struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Title       string `json:"title"`
	Status      string `json:"status"`
	Description string `json:"desciption"`
	Important   bool   `json:"important"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type MonogTask struct {
	ID          int    `json:"id" bson:"id"`
	UserID      int    `json:"user_id" bson:"user_id"`
	Status      string `json:"status" bson:"status"`
	Description string `json:"desciption" bson:"description"`
	Important   bool   `json:"important" bson:"important"`
}
