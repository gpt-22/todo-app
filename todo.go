package todo_app

type TodoList struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	IsDone bool `json:"is_done"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type TodoItem struct {
	Id int `json:"id"`
	Text string `json:"text"`
	IsDone bool `json:"is_done"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type TodoListItem struct {
	Id int `json:"id"`
	TodoItemId int `json:"todo_item_id"`
	UserTodoListId int `json:"user_todo_list_id"`
}

type UserTodoList struct {
	Id int `json:"id"`
	TodoListId int `json:"todo_list_id"`
	UserId int `json:"user_id"`
}
