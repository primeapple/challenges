package maxMealPrepTasks

import "slices"

// Idea: Gready Algorithm
// Sort by endTime ascending
// Do the first task, ignore all tasks that are not possible anymore
// Do it until queue at end of queue
func MaxMealPrepTasks(tasks []Task) TasksOrder {
	chosenTasks := []string{}

	sortedTasks := slices.Clone(tasks)
	slices.SortFunc(sortedTasks, func(p1, p2 Task) int {
		if p1.end < p2.end {
			return -1
		} else {
			return 1
		}
	})

	currentTime := 0
	for _, task := range sortedTasks {
		if currentTime <= task.start {
			chosenTasks = append(chosenTasks, task.name)
			currentTime = task.end
		}
	}


	return TasksOrder{count: len(chosenTasks), chosen: chosenTasks}
}
