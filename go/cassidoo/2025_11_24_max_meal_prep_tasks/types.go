package maxMealPrepTasks

type Task struct {
	name  string
	start int
	end   int
}

type TasksOrder struct {
	count  int
	chosen []string
}
