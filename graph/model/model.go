package model

//Video ...
type Video struct {
	ID     string `json:"id" bson:"_id"`
	Title  string `json:"title"`
	URL    string `json:"url"`
	Author *User  `json:"author"`
}

//NewVideo ...
type NewVideo struct {
	Title  string `json:"title"`
	URL    string `json:"url"`
	UserID string `json:"user_id"`
}

//Audio ...
type Audio struct {
	ID     string `json:"id" bson:"_id"`
	Title  string `json:"title"`
	URL    string `json:"url"`
	Author *User  `json:"author"`
}

//NewAudio ...
type NewAudio struct {
	Title  string `json:"title"`
	URL    string `json:"url"`
	UserID string `json:"user_id"`
}

//User ...
type User struct {
	ID   string `json:"id" bson:"_id"`
	Name string `json:"name"`
}
