package server

import (
	"context"
	"fmt"
	"github.com/zzxgzgz/alcor-control-agent-go/api/schema"
	"time"
)

type Goalstate_receiving_server struct {
}

func (s *Goalstate_receiving_server) PushNetworkResourceStates(ctx context.Context, goalState *schema.GoalState) (*schema.GoalStateOperationReply, error){
	fmt.Println("Called PushNetworkResourceStates")
	return nil, nil
}

func (s *Goalstate_receiving_server) PushGoalStatesStream(stream_server schema.GoalStateProvisioner_PushGoalStatesStreamServer) error{

	fmt.Println("Called PushGoalStatesStream")

	return  nil
}

func (s *Goalstate_receiving_server) RequestGoalStates(ctx context.Context, host_request *schema.HostRequest) (*schema.HostRequestReply, error){
	fmt.Println("Called RequestGoalStates")
	response := schema.HostRequestReply{
		FormatVersion:      123,
		OperationStatuses:  []*schema.HostRequestReply_HostRequestOperationStatus{
			&schema.HostRequestReply_HostRequestOperationStatus{
				RequestId:       host_request.StateRequests[0].RequestId,
				OperationStatus: schema.OperationStatus_SUCCESS,
			},
		},
		TotalOperationTime: 321,
	}
	// sleep 20 ms before returning, as the avg time for NCM to return is 20 ms.
	time.Sleep(time.Millisecond * 20)
	return &response, nil
}