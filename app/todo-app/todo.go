package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// Todo 構造体は個々のタスクを表します
type Todo struct {
	ID        int       `json:"id"`
	Task      string    `json:"task"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

// todosはTodoのリストを保持します
var todos []Todo
var nextID int = 1 // 新しいTodoに割り当てるID

const dataFile = "data.json"

// loadTodosはJSONファイルからTodoを読み込みます
func loadTodos() error {
	data, err := os.ReadFile(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			// ファイルが存在しない場合は、空のTodoリストを初期化します
			todos = []Todo{}
			nextID = 1
			return nil
		}
		return fmt.Errorf("ファイルの読み込み中にエラーが発生しました: %w", err)
	}

	if len(data) == 0 {
		todos = []Todo{}
		nextID = 1
		return nil
	}

	err = json.Unmarshal(data, &todos)
	if err != nil {
		return fmt.Errorf("JSONのデコード中にエラーが発生しました: %w", err)
	}

	// nextIDを設定する
	if len(todos) > 0 {
		maxID := 0
		for _, todo := range todos {
			if todo.ID > maxID {
				maxID = todo.ID
			}
		}
		nextID = maxID + 1
	} else {
		nextID = 1
	}

	return nil
}

// saveTodosはTodoをJSONファイルに保存します
func saveTodos() error {
	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return fmt.Errorf("JSONのエンコード中にエラーが発生しました: %w", err)
	}
	err = os.WriteFile(dataFile, data, 0644)
	if err != nil {
		return fmt.Errorf("ファイルの書き込み中にエラーが発生しました: %w", err)
	}
	return nil
}

// addTodoは新しいTodoをリストに追加します
func addTodo(task string) {
	todo := Todo{
		ID:        nextID,
		Task:      task,
		Completed: false,
		CreatedAt: time.Now(),
	}
	todos = append(todos, todo)
	nextID++
	fmt.Printf("Todo \"%s\" (ID: %d) が追加されました。\n", task, todo.ID)
}

// listTodosはすべてのTodoを表示します
func listTodos() {
	if len(todos) == 0 {
		fmt.Println("Todoがありません。")
		return
	}
	fmt.Println("\n--- Todoリスト ---")
	for _, todo := range todos {
		status := "未完了"
		if todo.Completed {
			status = "完了"
		}
		fmt.Printf("ID: %d | タスク: %s | 状態: %s | 作成日時: %s\n",
			todo.ID, todo.Task, status, todo.CreatedAt.Format("2006-01-02 15:04:05"))
	}
	fmt.Println("------------------")
}

// completeTodoは指定されたIDのTodoを完了済みにします
func completeTodo(id int) {
	found := false
	for i := range todos {
		if todos[i].ID == id {
			todos[i].Completed = true
			found = true
			fmt.Printf("Todo (ID: %d) が完了しました。\n", id)
			break
		}
	}
	if !found {
		fmt.Printf("Todo (ID: %d) が見つかりませんでした。\n", id)
	}
}

// deleteTodoは指定されたIDのTodoを削除します
func deleteTodo(id int) {
	found := false
	newTodos := []Todo{}
	for _, todo := range todos {
		if todo.ID == id {
			found = true
			continue
		}
		newTodos = append(newTodos, todo)
	}
	if found {
		todos = newTodos
		fmt.Printf("Todo (ID: %d) が削除されました。\n", id)
	} else {
		fmt.Printf("Todo (ID: %d) が見つかりませんでした。\n", id)
	}
}