package cli

import (
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/hortonworks/cb-cli/cli/utils"
	"github.com/hortonworks/cb-cli/client_cloudbreak/v3_workspace_id_stacks"
	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

type status string

func (c status) String() string {
	return string(c)
}

const (
	AVAILABLE        = status("AVAILABLE")
	STOPPED          = status("STOPPED")
	DELETE_COMPLETED = status("DELETE_COMPLETED")
	SKIP             = status("")
)

func (c *Cloudbreak) waitForOperationToFinish(workspaceID int64, name string, stackStatus, clusterStatus status) {
	defer utils.TimeTrack(time.Now(), "wait for operation to finish")

	log.Infof("[waitForOperationToFinish] start waiting")
	waitForOperationToFinishImpl(workspaceID, name, stackStatus, clusterStatus, c.Cloudbreak.V3WorkspaceIDStacks)
}

type getStackClient interface {
	GetStackInWorkspace(*v3_workspace_id_stacks.GetStackInWorkspaceParams) (*v3_workspace_id_stacks.GetStackInWorkspaceOK, error)
}

func waitForOperationToFinishImpl(workspaceID int64, name string, desiredStackStatus, desiredClusterStatus status, client getStackClient) {
	for {
		resp, err := client.GetStackInWorkspace(v3_workspace_id_stacks.NewGetStackInWorkspaceParams().WithWorkspaceID(workspaceID).WithName(name))
		if err != nil {
			utils.LogErrorAndExit(err)
		}

		stackStatus := resp.Payload.Status
		clusterStatus := resp.Payload.Cluster.Status
		log.Infof("[waitForClusterToFinishImpl] stack status: %s, cluster status: %s", stackStatus, clusterStatus)

		if (desiredStackStatus == SKIP || stackStatus == desiredStackStatus.String()) && (desiredClusterStatus == SKIP || clusterStatus == desiredClusterStatus.String()) {
			log.Infof("[waitForClusterToFinishImpl] cluster operation successfully finished")
			break
		}
		if strings.Contains(stackStatus, "FAILED") || strings.Contains(clusterStatus, "FAILED") {
			utils.LogErrorMessageAndExit("cluster operation failed")
		}

		log.Infof("[waitForClusterToFinishImpl] cluster operation is in progress, wait for 20 seconds")
		time.Sleep(20 * time.Second)
	}
}

func (c *Cloudbreak) getStackByName(workspaceID int64, name string) *models_cloudbreak.StackResponse {
	defer utils.TimeTrack(time.Now(), "get stack by name")

	log.Infof("[getStackByName] fetch stack, name: %s", name)
	stack, err := c.Cloudbreak.V3WorkspaceIDStacks.GetStackInWorkspace(v3_workspace_id_stacks.NewGetStackInWorkspaceParams().WithWorkspaceID(workspaceID).WithName(name))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	return stack.Payload
}

func (c *Cloudbreak) createStack(workspaceID int64, req *models_cloudbreak.StackV2Request) *models_cloudbreak.StackResponse {
	var stack *models_cloudbreak.StackResponse
	log.Infof("[createStack] sending create stack request")
	resp, err := c.Cloudbreak.V3WorkspaceIDStacks.CreateStackInWorkspace(v3_workspace_id_stacks.NewCreateStackInWorkspaceParams().WithWorkspaceID(workspaceID).WithBody(req))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	stack = resp.Payload

	log.Infof("[createStack] stack created: %s (id: %d)", *stack.Name, stack.ID)
	return stack
}

func (c *Cloudbreak) deleteStack(workspaceID int64, name string, forced bool) {
	defer utils.TimeTrack(time.Now(), "delete stack by name")

	log.Infof("[deleteStack] deleting stack, name: %s", name)
	err := c.Cloudbreak.V3WorkspaceIDStacks.DeleteStackInWorkspace(v3_workspace_id_stacks.NewDeleteStackInWorkspaceParams().WithWorkspaceID(workspaceID).WithName(name).WithForced(&forced))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
}
