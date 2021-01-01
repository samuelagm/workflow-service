package project

import (
	"context"
	"log"

	"github.com/samuelagm/workflow-service/common"
	"github.com/samuelagm/workflow-service/common/enum"
	"github.com/samuelagm/workflow-service/common/task"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/workflow"
)

// Workflow workflow definition
func Flow(ctx workflow.Context, state common.ProjectState) error {
	//logger := workflow.GetLogger(ctx)

	for _, t := range state.Pending {
		workflow.ExecuteChildWorkflow(ctx, task.Workflow, t) //TODO: get child task id
	}
	//wait for timer
	//run workflow for active task
	//wait for command signal
	//// if command is done, move task to completed
	//// if command is task update, update related task
	//// if command is acceptance, signal child workflows and return
	tripCh := workflow.GetSignalChannel(ctx, common.TaskQueue)
	for {
		var ev common.ProjectEvent
		tripCh.Receive(ctx, &ev)
		switch ev.Signal {
		case enum.REPORTUPDATE:
			//TODO: ...
		case enum.COMMENCEMENT:
			//TODO: ...
		case enum.PROJECTCOMPLETION:
			return nil
		}
	}
}

// StartWorkflow ...
func StartWorkflow(workflowID string, state common.ProjectState) error {
	c, err := client.NewClient(client.Options{
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		return err
	}
	defer c.Close()

	workflowOptions := client.StartWorkflowOptions{
		ID:        workflowID,
		TaskQueue: common.TaskQueue,
	}

	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, Flow, state)
	if err != nil {
		return err
	}
	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())
	return nil
}
