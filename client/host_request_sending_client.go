package client

import (
	"context"
	"fmt"
	"github.com/zzxgzgz/alcor-control-agent-go/api/schema"
	"google.golang.org/grpc"
	"math/rand"
)

type Host_request_sending_client struct {
	
}


// Push a group of network resource states
//
// Input: a GoalState object consists of a list of operation requests, and each request contains an operation type and a resource configuration
// Results consist of a list of operation statuses, and each status is a response to one operation request in the input
//
// Note: It is a NoOps for Control Agents when the operation type is INFO or GET.
//       Use RetrieveNetworkResourceStates for state query.
// this is for DPM->ACA
func (c *Host_request_sending_client)PushNetworkResourceStates(ctx context.Context, in *schema.GoalState, opts ...grpc.CallOption) (*schema.GoalStateOperationReply, error){
	fmt.Println("PushNetworkResourceStates called in server")

	operation_status_array := []*schema.GoalStateOperationReply_GoalStateOperationStatus{
		&schema.GoalStateOperationReply_GoalStateOperationStatus{
			ResourceId:               "123",
			ResourceType:             schema.ResourceType_PORT,
			OperationType:            schema.OperationType_INFO,
			OperationStatus:          schema.OperationStatus_SUCCESS,
			DataplaneProgrammingTime: 1,
			NetworkConfigurationTime: 1,
			StateElapseTime:          1,
		},
	}


	reply := schema.GoalStateOperationReply{
		FormatVersion:             rand.Uint32(),
		OperationStatuses:         operation_status_array,
		MessageTotalOperationTime: 0,
	}
	return &reply, nil
}
// similar to PushGoalStatesStream with streaming GoalStateV2 and streaming GoalStateOperationReply
// this is for DPM->NCM, and NCM->ACA
func (c *Host_request_sending_client)PushGoalStatesStream(ctx context.Context, opts ...grpc.CallOption) (schema.GoalStateProvisioner_PushGoalStatesStreamClient, error){
	fmt.Println("PushGoalStatesStream called in server")


	return nil, nil
}
// Request goal states for ACA on-demand processing and also agent restart handling
//
// Input: a HostRequest object consists of a list of ResourceStateRequest, and each request contains a RequestType
// Results consist of a list of HostRequestOperationStatus, and each status is a reply to each request in the input
// this is for ACA->NCM
func (c *Host_request_sending_client)RequestGoalStates(ctx context.Context, in *schema.HostRequest, opts ...grpc.CallOption) (*schema.HostRequestReply, error){
	fmt.Println("Received request in server!")

	response := schema.HostRequestReply{
		FormatVersion:      rand.Uint32(),
		OperationStatuses:  []*schema.HostRequestReply_HostRequestOperationStatus{&schema.HostRequestReply_HostRequestOperationStatus{
			RequestId: in.StateRequests[0].RequestId,	//set the same request ID as the request's
			OperationStatus: schema.OperationStatus_SUCCESS,
		}},
		TotalOperationTime: rand.Uint32(),
	}

	return &response, nil
}