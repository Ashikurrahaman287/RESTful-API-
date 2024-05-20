package todo

func GetAllTodos() []Todo {
	return GetTodos()
}

func GetSingleTodoByID(id int) *Todo {
	return GetTodoByID(id)
}

func CreateNewTodo(title string) Todo {
	return CreateTodo(title)
}

func UpdateExistingTodoByID(id int, title string, completed bool) *Todo {
	return UpdateTodoByID(id, title, completed)
}

func DeleteTodoByID(id int) bool {
	return DeleteTodoByID(id)
}
