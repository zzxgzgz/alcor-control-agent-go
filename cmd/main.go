package main

import (
	"context"
	"fmt"
	"github.com/zzxgzgz/alcor-control-agent-go/api/schema"
	"github.com/zzxgzgz/alcor-control-agent-go/server"
	"google.golang.org/grpc"
	"io"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	fmt.Println("hello world")

	args_without_program_name := os.Args[1:]

	if args_without_program_name[0] == "s" {
		 runServer()
	}else if args_without_program_name[0] == "c" {
		server_ip := "0.0.0.0"
		number_of_calls := 200
		if len(args_without_program_name )> 1{
			server_ip = args_without_program_name[1]
		}
		if len(args_without_program_name) > 2{
			number_of_calls, _ = strconv.Atoi(args_without_program_name[2])
		}
		runClient(server_ip, number_of_calls)
	}

	fmt.Println("Goodbye")
}


func runServer(){

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts ...)

	goalstate_receiving_server := server.Goalstate_receiving_server{}

	schema.RegisterGoalStateProvisionerServer(grpcServer, &goalstate_receiving_server)
	fmt.Println("Now running a goalstate receiving server")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}


//func worker(wg *sync.WaitGroup,c *schema.GoalStateProvisionerClient,id int, jobs <-chan *schema.HostRequest, results chan<-*schema.HostRequestReply){
//	for job := range jobs {
//		fmt.Println("worker", id, "started  job")
//		time.Sleep(time.Second)
//		fmt.Println("worker", id, "finished job")
//
//		host_request_reply, err := (*c).RequestGoalStates(context.Background(), job)
//		if err != nil {
//			log.Fatalf("Error when calling RequestGoalStates: %s", err)
//		}
//		log.Printf("Response from server: %v\n", host_request_reply.FormatVersion)
//		time.Sleep(time.Millisecond * 30)
//		results <- host_request_reply
//		(*wg).Done()
//	}
//}

func runClient(server_ip string, number_of_calls int){
	//number_of_calls := 200
	var waitGroup = sync.WaitGroup{}
	fmt.Println("Running client and trying to connect to server at ", server_ip+":9000")
	time.Sleep(10 * time.Second)
	conn, err := grpc.Dial(server_ip + ":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := schema.NewGoalStateProvisionerClient(conn)
	begin := time.Now()
	// try to use the same amount of workers (thread pool size) as the current ACA in test environment
	//number_of_workers := 16
	//jobs := make(chan *schema.HostRequest, number_of_calls)
	//results := make(chan *schema.HostRequestReply, number_of_calls)

	//for w:=0 ; w < number_of_workers ; w ++{
	//	go worker(&waitGroup, &c, w, jobs, results)
	//}
	for i:=0 ; i < number_of_calls ; i ++{
		waitGroup.Add(1)
		go func(id int) {
			defer waitGroup.Done()
			//fmt.Println(fmt.Sprintf("Preparing the %v th request", id))
			call_start := time.Now()
			state_request := schema.HostRequest_ResourceStateRequest{
				RequestType:     schema.RequestType_ON_DEMAND,
				RequestId:       strconv.Itoa(id),
				TunnelId:        1,
				SourceIp:        "123.123.321.321",
				SourcePort:      999,
				DestinationIp:   "333.222.111.000",
				DestinationPort: 888,
				Ethertype:       schema.EtherType_IPV4,
				Protocol:        schema.Protocol_ARP,
			}
			state_request_array := []*schema.HostRequest_ResourceStateRequest{&state_request}
			fmt.Println(fmt.Sprintf("Sending the %v th request", id))

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
		}(i)
	}
	waitGroup.Wait()
	end := time.Now()
	diff := end.Sub(begin)
	fmt.Println("Finishing ", number_of_calls, " RequestGoalStates calls took ",diff.Milliseconds(), " ms")


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
			fmt.Printf("Received a goalstate operation reply for the %vth goalstatev2: %v", (*in).FormatVersion,(*in).MessageTotalOperationTime)
		}
	}()
	for a:= 0 ; a < number_of_calls ; a ++{
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
	}
	fmt.Println("All gsv2 sent, closing the stream")
	stream.CloseSend()
	<-waitc
}