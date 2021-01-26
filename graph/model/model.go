package model

//NewVideo ...
type NewVideo struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

//Video ...
type Video struct {
	ID    string `json:"id" bson:"_id"`
	Title string `json:"title"`
	URL   string `json:"url"`
}
