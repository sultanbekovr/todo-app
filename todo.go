package todo

type TodoList struct {
	Id    int    `json:"-"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
}
type UserList struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id    int    `json:"-"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
	Done  bool   `json:"done"`
}

type ListItem struct {
	Id     int
	ListId int
	ItemId int
}
