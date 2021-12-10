package main

import (
	"context"
	"fmt"
	"github.com/zzxgzgz/alcor-control-agent-go/api/schema"
	"github.com/zzxgzgz/alcor-control-agent-go/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

var aca_server_port string
var ncm_ip string
var ncm_gRPC_port string
const latency_mode = 1
const throughput_mode = 2
var test_mode_latency_or_throughput int = latency_mode
var number_of_calls = 0
var client_call_length_in_seconds int = 1
var global_server *grpc.Server
var global_server_api_instance *server.Goalstate_receiving_server
var global_client_connection *grpc.ClientConn

func main() {
	log.Println("hello world")

	args_without_program_name := os.Args[1:]

	aca_server_port = args_without_program_name[0]

	ncm_ip = args_without_program_name[1]

	ncm_gRPC_port = args_without_program_name[2]

	test_mode_latency_or_throughput,_ = strconv.Atoi(args_without_program_name[3])

	if test_mode_latency_or_throughput == throughput_mode{
		client_call_length_in_seconds,_ = strconv.Atoi(args_without_program_name[4])
	}else if test_mode_latency_or_throughput == latency_mode {
		number_of_calls,_ = strconv.Atoi(args_without_program_name[4])
	}else{
		log.Fatalf("Client got unknown test mode: %v, exiting...", test_mode_latency_or_throughput)
	}

	log.Printf("Running gRPC server at localhost:%s, gRPC client connecting to server at: %s:%s, the test will last %d seconds\n", aca_server_port, ncm_ip, ncm_gRPC_port, client_call_length_in_seconds)

	go runServer()

	runClient()

	select {
	case <-time.After(60 * time.Second):
		log.Println("It has been 60 seconds, goodbye")
	}
	fmt.Println("Goodbye")
}


func runServer(){

	lis, err := net.Listen("tcp", ":50001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption

	global_server = grpc.NewServer(opts ...)

	global_server_api_instance := server.Goalstate_receiving_server{
		Received_goalstatev2_count: 0,
	}

	schema.RegisterGoalStateProvisionerServer(global_server, &global_server_api_instance)
	fmt.Println("Now running a goalstate receiving server")
	if err := global_server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}

func runClient(){
	//number_of_calls := 200
	var waitGroup = sync.WaitGroup{}
	fmt.Println("Running client and trying to connect to server at ", ncm_ip+":"+ncm_gRPC_port)
	//time.Sleep(10 * time.Second)
	var err error
	global_client_connection, err = grpc.Dial(ncm_ip + ":" + ncm_gRPC_port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer global_client_connection.Close()
	c := schema.NewGoalStateProvisionerClient(global_client_connection)
	begin := time.Now()
	// try to use the same amount of workers (thread pool size) as the current ACA in test environment
	//number_of_workers := 16
	//jobs := make(chan *schema.HostRequest, number_of_calls)
	//results := make(chan *schema.HostRequestReply, number_of_calls)

	//for w:=0 ; w < number_of_workers ; w ++{
	//	go worker(&waitGroup, &c, w, jobs, results)
	//}

	fmt.Println("Press the Enter Key to terminate the console screen!")
	fmt.Scanln() // wait for Enter Key

	request_id := 0
	if test_mode_latency_or_throughput == latency_mode{
		log.Printf("Client running in latency mode, it will send %d on-demand requests concurrently\n", number_of_calls)
		for i := 0 ; i < number_of_calls ; i ++ {
			waitGroup.Add(1)

			go func(id int) {
				defer waitGroup.Done()
				//fmt.Println(fmt.Sprintf("Preparing the %v th request", id))
				call_start := time.Now()
				state_request := schema.HostRequest_ResourceStateRequest{
					RequestType:     schema.RequestType_ON_DEMAND,
					RequestId:       strconv.Itoa(id),
					TunnelId:        21,				// this should be changed when testing with more than 1 VPCs
					SourceIp:        "10.0.0.3",		// this should be changed when so that we can test different src IPs
					SourcePort:      1,
					DestinationIp:   "10.0.2.2",		// this should be changed when so that we can test different dest IPs
					DestinationPort: 1,
					Ethertype:       schema.EtherType_IPV4,
					Protocol:        schema.Protocol_ARP,
				}
				state_request_array := []*schema.HostRequest_ResourceStateRequest{&state_request}
				//fmt.Println(fmt.Sprintf("Sending the %v th request", id))

				host_request := schema.HostRequest{
					FormatVersion: rand.Uint32(),
					StateRequests: state_request_array,
				}
				//jobs <- &host_request
				send_request_time := time.Now()
				host_request_reply, err := c.RequestGoalStates(context.Background(), &host_request)
				received_reply_time := time.Now()
				if err != nil {
					log.Fatalf("Error when calling RequestGoalStates: %s", err)
				}
				time.Sleep(time.Millisecond * 30)
				log.Printf("For the %dth request, total time took %d ms,\ngRPC call time took %d ms\nResponse from server: %v\n",
					id, received_reply_time.Sub(call_start).Milliseconds(), received_reply_time.Sub(send_request_time).Milliseconds(),
					host_request_reply.OperationStatuses[0].RequestId)
			}(request_id)
			// update now and update request ID
			request_id ++
			//now = time.Now()
		}
		waitGroup.Wait()
		end := time.Now()
		diff := end.Sub(begin)
		fmt.Println("Finishing ", number_of_calls, " RequestGoalStates calls took ",diff.Milliseconds(), " ms")

	}else if test_mode_latency_or_throughput == throughput_mode{
		log.Printf("Client running in throughput mode, it will send on-demand requests concurrently for %d seconds", client_call_length_in_seconds)
		through_put_test_start_time := time.Now()
		through_put_test_end_time := through_put_test_start_time.Add(time.Duration(client_call_length_in_seconds) * time.Second )
		//test_stop_channel := make(chan struct{})
		request_id := 0
		count := 0
		ctx := context.Background()

		ctx, cancel := context.WithDeadline(ctx, through_put_test_end_time)
		defer cancel()

		for {
			if through_put_test_end_time.Sub(time.Now()).Milliseconds() <= 0 {
				//test_stop_channel <- struct{}{}
				ctx.Done()
				(*global_server_api_instance).Mu.Lock()
				global_server.GracefulStop()
				(*global_server_api_instance).Mu.Unlock()
				global_client_connection.Close()
				log.Println("Time's up, stop the server")
				log.Printf("Time to stop sending requests, requests sent: %d, request finished: %d, received GoalStateV2 amount: %d\n", request_id +1, count, &global_server_api_instance.Received_goalstatev2_count)
				break
			}

			go func(id int, ctx *context.Context) {
				defer waitGroup.Done()
				//fmt.Println(fmt.Sprintf("Preparing the %v th request", id))
				call_start := time.Now()
				state_request := schema.HostRequest_ResourceStateRequest{
					RequestType:     schema.RequestType_ON_DEMAND,
					RequestId:       strconv.Itoa(id),
					TunnelId:        21,				// this should be changed when testing with more than 1 VPCs
					SourceIp:        "10.0.0.3",		// this should be changed when so that we can test different src IPs
					SourcePort:      1,
					DestinationIp:   "10.0.2.2",		// this should be changed when so that we can test different dest IPs
					DestinationPort: 1,
					Ethertype:       schema.EtherType_IPV4,
					Protocol:        schema.Protocol_ARP,
				}
				state_request_array := []*schema.HostRequest_ResourceStateRequest{&state_request}
				//fmt.Println(fmt.Sprintf("Sending the %v th request", id))

				host_request := schema.HostRequest{
					FormatVersion: rand.Uint32(),
					StateRequests: state_request_array,
				}
				//jobs <- &host_request
				if global_client_connection.GetState() == connectivity.Idle ||
					global_client_connection.GetState() == connectivity.Ready{
					waitGroup.Add(1)
					send_request_time := time.Now()
					select {
						case <-(*ctx).Done():
							return
					default:
						host_request_reply, err := c.RequestGoalStates(*ctx, &host_request)
						received_reply_time := time.Now()
						if err != nil {
							log.Printf("Error when calling RequestGoalStates: %s\n", err)
							return
						}
						//time.Sleep(time.Millisecond * 30)
						log.Printf("For the %dth request, total time took %d ms,\ngRPC call time took %d ms\nResponse from server: %v\n",
							id, received_reply_time.Sub(call_start).Milliseconds(), received_reply_time.Sub(send_request_time).Milliseconds(),
							host_request_reply.OperationStatuses[0].RequestId)
						count ++}

				}

			}(request_id, &ctx)
			// update now and update request ID
			request_id ++
			//now = time.Now()
		}
		log.Println("Outside of the for loop, now wait a little bit")
		waitGroup.Wait()
		end := time.Now()
		diff := end.Sub(begin)
		fmt.Println("Sent ", count, " RequestGoalStates calls in ", client_call_length_in_seconds ," seconds, took ",diff.Milliseconds(), " ms to receive the results")
	}else{
		log.Fatalf("Client got unknown test mode: %v, exiting...", test_mode_latency_or_throughput)
	}

	// commenting out this goalstate sending part in the client, as it should now be done by NCM.
/*
	fmt.Println("Time to call the same amount of PushGoalStates streaming calls")
	stream, err := c.PushGoalStatesStream(context.Background())
	waitc := make(chan struct{})
	go func(){
		for {
			in, err := stream.Recv()
			if err == io.EOF{
				close(waitc)
				fmt.Println("Returning because the stream received io.EOF")
				return
			}
			if err != nil {
				fmt.Printf("Failed to receive a goalstate programming result: %v\n", err)
			}
			fmt.Printf("Received a goalstate operation reply for the %vth goalstatev2: %v\n", (*in).FormatVersion,(*in).MessageTotalOperationTime)
		}
	}()
	for a:= 0 ; a < client_call_length_in_seconds ; a ++{
		v2 := schema.GoalStateV2{
			FormatVersion:       uint32(a),
			HostResources:       nil,
			VpcStates:           nil,
			SubnetStates:        nil,
			PortStates:          nil,
			DhcpStates:          nil,
			NeighborStates:      nil,
			SecurityGroupStates: nil,
			RouterStates:        nil,
			GatewayStates:       nil,
		}
		if err:= stream.Send(&v2); err != nil{
			fmt.Printf("Sending GoalStateV2: %v gave this error: %v\n", v2.FormatVersion, err)
		}
		fmt.Printf("Sent the %vth gsv2\n", a)
	}
	fmt.Println("All gsv2 sent, closing the stream")
	time.Sleep(time.Second * 10)
	stream.CloseSend()
	<-waitc
 */
	log.Println("This is the end of runClient")
}