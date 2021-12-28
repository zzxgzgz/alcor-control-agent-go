package main

import (
	"context"
	"fmt"
	"github.com/zzxgzgz/alcor-control-agent-go/api/schema"
	"github.com/zzxgzgz/alcor-control-agent-go/server"
	"google.golang.org/grpc"
	//"google.golang.org/grpc/connectivity"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"sync"
	system_time "time"
	"golang.org/x/time/rate"
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
//var global_client_connection *grpc.ClientConn
var global_received_goalstate_count int = 0
var global_sent_on_demand_request_limit_per_second int = 0
var global_number_of_clients int = 0
var global_number_of_connections int = 0

/*
Arguments:
1. aca_server_port, the port of the gRPC server
2. ncm_ip, IP address of the NCM, used by the gRPC client to connect to the NCM
3. ncm_gRPC_port, the gRPC port of NCM, used by the gRPC client to connect to the NCM
4. test_mode_latency_or_throughput, a flag used to indicate which mode to run for the gRPC client.
The throughput_mode, which sends as many requests as possible in a time interval, is 2;
and the latency_mode, which sends a fixed number of requests, and see how much time it takes, is 1.
5.1. client_call_length_in_seconds, used with the throughput_mode, indicates the time interval(in seconds) of the test.
5.2. number_of_calls, used with the latency_mode, indicates how many requests the gRPC client will send.
6. global_sent_on_demand_request_limit_per_second, used to limit how many request the program can send per second.
7. global_number_of_clients, indicates how many gRPC clients this program creates.
8. global_number_of_connections, indicates how many gRPC connections this program creates.

Examples:
1) Run the throughput test for 10 seconds, limits send requests to 100 per second, create 20 clients and 30 connections:
```
go build cmd/main.go
./main 50001 ${ncm_ip} ${ncm_gRPC_port} 2 10 100 20 30
```

2) Run the latency test with 1000 on-demand requests, limits send requests to 100 per second, create 20 clients and 30 connections:
```
go build cmd/main.go
./main 50001 ${ncm_ip} ${ncm_gRPC_port} 1 1000 100 20 30
```
*/
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

	global_sent_on_demand_request_limit_per_second, _  = strconv.Atoi(args_without_program_name[5])

	global_number_of_clients, _ = strconv.Atoi(args_without_program_name[6])

	global_number_of_connections, _ = strconv.Atoi(args_without_program_name[7])

	log.Printf("Running gRPC server at localhost:%s, gRPC client connecting to server at: %s:%s, the test will last %d seconds, the send limit is %d request/second\n", aca_server_port, ncm_ip, ncm_gRPC_port, client_call_length_in_seconds, global_sent_on_demand_request_limit_per_second)

	go runServer()

	runClient()
	//
	//select {
	//case <-time.After(60 * time.Second):
	//	log.Println("It has been 60 seconds, goodbye")
	//}
	log.Println("Goodbye")
}

// runs the gRPC server that receives GoalStateV2
func runServer(){

	lis, err := net.Listen("tcp", ":50001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption

	global_server = grpc.NewServer(opts ...)

	global_server_api_instance := server.Goalstate_receiving_server{
		Received_goalstatev2_count: &global_received_goalstate_count,
	}

	schema.RegisterGoalStateProvisionerServer(global_server, &global_server_api_instance)
	log.Printf("Now running a goalstate receiving server, current count: %d", global_received_goalstate_count)
	log.Println("Now running a goalstate receiving server")
	if err := global_server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}

// runs the gRPC client that sends out on-demand requests
func runClient(){
	limit := rate.NewLimiter(rate.Limit(global_sent_on_demand_request_limit_per_second), 1)
	//number_of_calls := 200
	var waitGroup = sync.WaitGroup{}
	log.Println("Running client and trying to connect to server at ", ncm_ip+":"+ncm_gRPC_port)
	//time.Sleep(10 * time.Second)

	client_connections := make([]*grpc.ClientConn, global_number_of_connections)

	for i := 0 ; i < global_number_of_connections ; i ++ {
		client_connection, err := grpc.Dial(ncm_ip + ":" + ncm_gRPC_port, grpc.WithInsecure())
		defer client_connection.Close()

		if err != nil {
			log.Fatalf("The %dth connection did not connect: %s\n", i, err)
		}
		client_connections = append(client_connections, client_connection)
	}

	clients := make([]schema.GoalStateProvisionerClient, global_number_of_clients)

	for i := 0 ; i < global_number_of_clients ; i ++ {
		client := schema.NewGoalStateProvisionerClient(client_connections[i%global_number_of_connections])
		clients = append(clients, client)
	}

	begin := system_time.Now()

	log.Println("Press the Enter Key to terminate the console screen!")
	fmt.Scanln() // wait for Enter Key

	request_id := 0
	if test_mode_latency_or_throughput == latency_mode{
		log.Printf("Client running in latency mode, it will send %d on-demand requests concurrently\n", number_of_calls)
		for i := 0 ; i < number_of_calls ; i ++ {
			waitGroup.Add(1)

			go func(id int) {
				defer waitGroup.Done()
				limit.Wait(context.Background())
				call_start := system_time.Now()
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

				host_request := schema.HostRequest{
					FormatVersion: rand.Uint32(),
					StateRequests: state_request_array,
				}
				//jobs <- &host_request
				send_request_time := system_time.Now()
				host_request_reply, err := clients[(global_number_of_clients % id)].RequestGoalStates(context.Background(), &host_request)
				received_reply_time := system_time.Now()
				if err != nil {
					log.Fatalf("Error when calling RequestGoalStates: %s", err)
				}
				log.Printf("For the %dth request, total time took %d ms,\ngRPC call time took %d ms\nResponse from server: %v\n",
					id, received_reply_time.Sub(call_start).Milliseconds(), received_reply_time.Sub(send_request_time).Milliseconds(),
					host_request_reply.OperationStatuses[0].RequestId)
			}(request_id)
			// update now and update request ID
			request_id ++
			//now = time.Now()
		}
		waitGroup.Wait()
		end := system_time.Now()
		diff := end.Sub(begin)
		log.Println("Finishing ", number_of_calls, " RequestGoalStates calls took ",diff.Milliseconds(), " ms")

	}else if test_mode_latency_or_throughput == throughput_mode{
		log.Printf("Client running in throughput mode, it will send on-demand requests concurrently for %d seconds", client_call_length_in_seconds)
		through_put_test_start_time := system_time.Now()
		through_put_test_end_time := through_put_test_start_time.Add(system_time.Duration(client_call_length_in_seconds) * system_time.Second )
		//test_stop_channel := make(chan struct{})
		request_id := 0
		count := 0
		ctx := context.Background()

		ctx, cancel := context.WithDeadline(ctx, through_put_test_end_time)
		defer cancel()

		for {
			if through_put_test_end_time.Sub(system_time.Now()).Milliseconds() <= 0 {
				log.Println("Time's up, stop the server")
				for _, connection := range client_connections {
					connection.Close()
				}
				global_server.Stop()
				break
			}else{
				go func(id int, ctx *context.Context) {
					//connection_to_use := client_connections[global_number_of_connections % id]
					call_start := system_time.Now()
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

					host_request := schema.HostRequest{
						FormatVersion: rand.Uint32(),
						StateRequests: state_request_array,
					}
					//if connection_to_use.GetState() == connectivity.Idle ||
					//	connection_to_use.GetState() == connectivity.Ready{
					limit.Wait(*ctx)
					defer waitGroup.Done()
					waitGroup.Add(1)
					send_request_time := system_time.Now()
					select {
					case <-(*ctx).Done():
						break
					default:
						host_request_reply, err := clients[id % global_number_of_clients].RequestGoalStates(*ctx, &host_request)
						received_reply_time := system_time.Now()
						if err != nil {
							log.Printf("Error when calling RequestGoalStates: %s\n", err)
							break
						}
						//time.Sleep(time.Millisecond * 30)
						log.Printf("For the %dth request, total time took %d ms,\ngRPC call time took %d ms\nResponse from server: %v\n",
							id, received_reply_time.Sub(call_start).Milliseconds(), received_reply_time.Sub(send_request_time).Milliseconds(),
							host_request_reply.OperationStatuses[0].RequestId)
						count ++}

					//}

				}(request_id, &ctx)
				// update now and update request ID
				request_id ++
			}
		}
		log.Println("Outside of the for loop, now wait a little bit")
		waitGroup.Wait()
		end := system_time.Now()
		diff := end.Sub(begin)
		log.Printf("Requests sent: %d, request finished: %d, received GoalStateV2 amount: %d\n", request_id +1, count, global_received_goalstate_count)
		log.Println("Sent ", request_id + 1, " RequestGoalStates calls in ", client_call_length_in_seconds ," seconds, took ",diff.Milliseconds(), " ms to receive the results")
	}else{
		log.Fatalf("Client got unknown test mode: %v, exiting...", test_mode_latency_or_throughput)
	}

	// commenting out this goalstate sending part in the client, as it should now be done by NCM.
/*
	log.Println("Time to call the same amount of PushGoalStates streaming calls")
	stream, err := c.PushGoalStatesStream(context.Background())
	waitc := make(chan struct{})
	go func(){
		for {
			in, err := stream.Recv()
			if err == io.EOF{
				close(waitc)
				log.Println("Returning because the stream received io.EOF")
				return
			}
			if err != nil {
				log.Printf("Failed to receive a goalstate programming result: %v\n", err)
			}
			log.Printf("Received a goalstate operation reply for the %vth goalstatev2: %v\n", (*in).FormatVersion,(*in).MessageTotalOperationTime)
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
			log.Printf("Sending GoalStateV2: %v gave this error: %v\n", v2.FormatVersion, err)
		}
		log.Printf("Sent the %vth gsv2\n", a)
	}
	log.Println("All gsv2 sent, closing the stream")
	time.Sleep(time.Second * 10)
	stream.CloseSend()
	<-waitc
 */
	log.Println("This is the end of runClient")
}