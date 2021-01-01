package prodder

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/samuelagm/workflow-service/common"
	"github.com/samuelagm/workflow-service/common/project"
	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.NewClient(client.Options{
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	// This workflow ID can be user business logic identifier as well.
	workflowID := fmt.Sprintf("wk_%v", uuid.New())
	workflowOptions := client.StartWorkflowOptions{
		ID:        workflowID,
		TaskQueue: common.TaskQueue,
	}

	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, project.Flow, "9-")
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}
	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())
}
