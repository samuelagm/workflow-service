package common

import "github.com/samuelagm/workflow-service/common/enum"

// TripSignalName ...
const TripSignalName = "taskflow_signal"

type taskState int

type (
	// TaskEvent ...
	TaskEvent struct {
		ID     string
		Signal enum.TaskSignal
	}
)

const (
	// TaskQueue ...
	TaskQueue = "projectflow"
)

type (
	// ProjectEvent ...
	ProjectEvent struct {
		ID     string
		Signal enum.ProjectSignal
	}

	// Note ..
	Note struct {
		ID   string
		Text string
	}

	// Note ..
	Milestone struct {
		Name   string
		Status int
	}

	// Task ...
	Task struct {
		ID           string
		Milestone    []Milestone
		Notes        []Note
		Dependencies []Task
	}

	// ProjectState ...
	ProjectState struct {
		Running   []Task
		Completed []Task
		Pending   []Task
	}
)
