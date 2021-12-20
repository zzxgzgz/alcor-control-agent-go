package server

import (
	"context"
	"github.com/zzxgzgz/alcor-control-agent-go/api/schema"
	"log"
	"sync"
	"time"
)

type Goalstate_receiving_server struct {
	Received_goalstatev2_count *int
	Mu sync.Mutex
}

func (s *Goalstate_receiving_server) PushNetworkResourceStates(ctx context.Context, goalState *schema.GoalState) (*schema.GoalStateOperationReply, error){
	log.Println("Called PushNetworkResourceStates")
	return nil, nil
}

func (s *Goalstate_receiving_server) PushGoalStatesStream(stream_server schema.GoalStateProvisioner_PushGoalStatesStreamServer) error{
	select {
		case <-stream_server.Context().Done():
			log.Println("Canceled because context is Done")
			return stream_server.Context().Err()
	default:
		//for{
		gsv2_ptr, err := stream_server.Recv()
		//if err == io.EOF{
		//	fmt.Println("End of stream, getting out")
		//	break
		//}
		if err != nil {
			log.Printf("Got this error when reading from stream: %v\n", err)
		}
		if gsv2_ptr != nil {
			s.Mu.Lock()
			*(s.Received_goalstatev2_count) = *(s.Received_goalstatev2_count) + 1
			s.Mu.Unlock()
			log.Println("Called PushGoalStatesStream for the ", *(s.Received_goalstatev2_count), " time")
			log.Println("Read a gsv2 from the stream")
			// use this following go routine to simulate using another thread to program goalstatev2, and reply
			//go func() {
			reply := schema.GoalStateOperationReply{
				FormatVersion:             (*gsv2_ptr).FormatVersion,
				OperationStatuses:         nil,
				MessageTotalOperationTime: 30,
			}
			//time.Sleep(time.Millisecond * 30)
			stream_server.SendMsg(&reply)
			//stream_server.Send(&reply)
			//}()
			//}
		}

	}

	return  nil
}

func (s *Goalstate_receiving_server) RequestGoalStates(ctx context.Context, host_request *schema.HostRequest) (*schema.HostRequestReply, error){
	log.Println("Called RequestGoalStates")
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