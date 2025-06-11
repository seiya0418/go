package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	err := loadTodos()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Todoの読み込み中にエラーが発生しました: %v\n", err)
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- Todoアプリ ---")
		fmt.Println("1. Todoを追加")
		fmt.Println("2. Todoを一覧表示")
		fmt.Println("3. Todoを完了する")
		fmt.Println("4. Todoを削除")
		fmt.Println("5. 終了")
		fmt.Print("選択してください: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			fmt.Print("新しいTodoのタスクを入力してください: ")
			task, _ := reader.ReadString('\n')
			addTodo(strings.TrimSpace(task))
			saveTodos()
		case "2":
			listTodos()
		case "3":
			fmt.Print("完了するTodoのIDを入力してください: ")
			idStr, _ := reader.ReadString('\n')
			id, err := strconv.Atoi(strings.TrimSpace(idStr))
			if err != nil {
				fmt.Println("無効なIDです。数値を入力してください。")
				continue
			}
			completeTodo(id)
			saveTodos()
		case "4":
			fmt.Print("削除するTodoのIDを入力してください: ")
			idStr, _ := reader.ReadString('\n')
			id, err := strconv.Atoi(strings.TrimSpace(idStr))
			if err != nil {
				fmt.Println("無効なIDです。数値を入力してください。")
				continue
			}
			deleteTodo(id)
			saveTodos()
		case "5":
			fmt.Println("Todoアプリを終了します。")
			return
		default:
			fmt.Println("無効な選択です。もう一度お試しください。")
		}
	}
}