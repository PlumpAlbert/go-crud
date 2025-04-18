package main

import "database/sql"

func GetTasks(dbPointer *sql.DB) ([]Task, error) {
	query := "SELECT id, task, done FROM tasks;"

	rows, err := dbPointer.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tasks []Task

	for rows.Next() { // Loop over rows in result set
		var todo Task

		err = rows.Scan(&todo.Id, &todo.Task, &todo.Done)
		if err != nil {
			return nil, err
		}

		// Build the array of Tasks from returned rows
		tasks = append(tasks, todo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}
