package todo

var todos []Todo
var nextID = 1

func GetTodos() []Todo {
	return todos
}

func GetTodoByID(id int) *Todo {
	for _, t := range todos {
		if t.ID == id {
			return &t
		}
	}
	return nil
}

func CreateTodo(title string) Todo {
	t := Todo{
		ID:        nextID,
		Title:     title,
		Completed: false,
	}
	todos = append(todos, t)
	nextID++
	return t
}

func UpdateTodoByID(id int, title string, completed bool) *Todo {
	for i, t := range todos {
		if t.ID == id {
			todos[i].Title = title
			todos[i].Completed = completed
			return &todos[i]
		}
	}
	return nil
}

func DeleteTodoByID(id int) bool {
	for i, t := range todos {
		if t.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return true
		}
	}
	return false
}
