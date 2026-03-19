package scanner

import (
	"fmt"
	"sort"
	"todoapp/todo"
)

func printResult(result string) {
	fmt.Println("Результат вывода команды:", result)
	fmt.Println("")
}

func printPromt() {
	fmt.Println("Введите команду: ")
}

func printExit() {
	fmt.Println("Завершение работы. До скорого!")
}

func printAdd(title string) {
	fmt.Println("Задача '" + title + "' успешно добавлена")
	fmt.Println("")
}

func printTasks(tasks map[string]todo.Task) {
	if len(tasks) == 0 {
		fmt.Println("Список дел пуст.")
		fmt.Println("")
		return
	}

	titles := make([]string, 0, len(tasks))
	for title := range tasks {
		titles = append(titles, title)
	}
	sort.Strings(titles)

	fmt.Println("=== Список текущих задач ===")
	for i, title := range titles {
		task := tasks[title]

		status := " "
		if task.IsDone {
			status = "X"
		}

		timeStr := task.CreatedAt.Format("2006-01-02 15:04")

		fmt.Printf("%d. [%s] %-25s | Добавлено: %s\n", i+1, status, task.Title, timeStr)
	}
	fmt.Println("")
}

func printDone(title string) {
	fmt.Println("Задача '" + title + "' помечена как выполненная")
	fmt.Println("")
}

func printDel(title string) {
	fmt.Println("Задача '" + title + "' успешно удалена")
	fmt.Println("")

}

func printHelp() {
	fmt.Println("Список команд:")
	fmt.Println("help — эта команда позволяет узнать доступные команды и их формат")
	fmt.Println("")
	fmt.Println("add {заголовок задачи из одного слова} {текст задачи из одного или нескольких слов} — эта команда позволяет добавлять новые задачи в список задач")
	fmt.Println("")
	fmt.Println("list — эта команда позволяет получить полный список всех задач")
	fmt.Println("")
	fmt.Println("del {заголовок существующей задачи} — эта команда позволяет удалить задачу по её заголовку")
	fmt.Println("")
	fmt.Println("done {заголовок существующей задачи} — эта команда позволяет отменить задачу как выполненную")
	fmt.Println("")
	fmt.Println("events — эта команда позволяет получить список всех событий")
	fmt.Println("")
	fmt.Println("exit — эта команда позволяет завершить выполнение программы")
	fmt.Println("")
}

func printEvents(events []Event) {
	if len(events) == 0 {
		fmt.Println("Список событий пуст.")
		fmt.Println("")
		return
	}

	fmt.Println("Список событий")
	for i, event := range events {
		timeStr := event.dateAt.Format("2006-01-02 15:04:05")

		fmt.Printf("[%d] Время:  %s\n", i+1, timeStr)
		fmt.Printf("    Ввод:   %s\n", event.userInput)
		fmt.Printf("    Статус: %s\n", event.description)
	}
	fmt.Println("")
}
