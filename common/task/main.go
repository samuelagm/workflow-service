package task

import (
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

// TripSignalName ...
const TripSignalName = "taskflow_signal"

type taskState int

// TaskSignal ...
type TaskSignal int

const (
	// ACCEPTANCE ...
	ACCEPTANCE TaskSignal = iota
	// REJECTION ...
	REJECTION
	// COMPLETION ...
	COMPLETION
)

type (
	// TaskEvent ...
	TaskEvent struct {
		ID     string
		Signal TaskSignal
	}
)

// Workflow ...
func Workflow(ctx workflow.Context, taskID string) error {

	//TODO: get task object
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute * 10,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Minute,
			BackoffCoefficient: 2.0,
			MaximumInterval:    30 * time.Minute,
			MaximumAttempts:    5,
		},
	}
	ctx1 := workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)

	//send taskId to daily email scheduler service
	err := workflow.ExecuteActivity(ctx1, ScheduleEmails, taskID).Get(ctx1, nil)
	if err != nil {
		logger.Error("Failed to create expense report", "Error", err)
		return err
	}

	//wait for command signals signals
	//// if done signal send task acceptance email to child tasks if any
	//// if child task send acceptance signal, start child task workflow
	//// if project accepted return

	tripCh := workflow.GetSignalChannel(ctx, TripSignalName)
	for {
		var ev TaskEvent
		tripCh.Receive(ctx, &ev)
		switch ev.Signal {
		case ACCEPTANCE:
			return nil
		case REJECTION:
			//TODO: ...
		case COMPLETION:
			workflow.ExecuteChildWorkflow(ctx, Workflow, "child-taskid") //TODO: get child task id
		}

	}
}
