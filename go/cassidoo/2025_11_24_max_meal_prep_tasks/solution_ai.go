package maxMealPrepTasks

import (
	"sort"
)

func MaxMealPrepTasksAI(tasks []Task) TasksOrder {
	if len(tasks) == 0 {
		return TasksOrder{count: 0, chosen: []string{}}
	}

	// Sort tasks by end time ascending
	sortedTasks := make([]Task, len(tasks))
	copy(sortedTasks, tasks)
	sort.Slice(sortedTasks, func(i, j int) bool {
		return sortedTasks[i].end < sortedTasks[j].end
	})

	var chosen []string
	currentEnd := 0

	for _, task := range sortedTasks {
		if task.start >= currentEnd {
			chosen = append(chosen, task.name)
			currentEnd = task.end
		}
	}

	return TasksOrder{
		count:  len(chosen),
		chosen: chosen,
	}
}
