package todo

type TodoList struct {
	Id          int    `json:"id"`
	Title       string `json:"title" binding="required"`
	Description string `json:"description"`
}

type UserList struct {
	Id     int
	UserId int
	Listid int
}

type TodoItem struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type ListItem struct {
	Id     int
	Listid int
	ItemId int
}
