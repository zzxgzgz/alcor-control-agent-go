# alcor-control-agent-go
Golang version of the Alcor-Control-Agent

Arguments:
1. aca_server_port, the port of the gRPC server
2. ncm_ip, IP address of the NCM, used by the gRPC client to connect to the NCM
3. ncm_gRPC_port, the gRPC port of NCM, used by the gRPC client to connect to the NCM
4. test_mode_latency_or_throughput, a flag used to indicate which mode to run for the gRPC client.
   The throughput_mode, which sends as many requests as possible in a time interval, is 2;
   and the latency_mode, which sends a fixed number of requests, and see how much time it takes, is 1.
   5.1. client_call_length_in_seconds, used with the throughput_mode, indicates the time interval(in seconds) of the test.
   5.2. number_of_calls, used with the latency_mode, indicates how many requests the gRPC client will send.

Examples:
1) Run the throughput test for 10 seconds:
```
go build cmd/main.go
./main 50001 ${ncm_ip} ${ncm_gRPC_port} 2 10
```

2) Run the latency test with 1000 on-demand requests:
```
go build cmd/main.go
./main 50001 ${ncm_ip} ${ncm_gRPC_port} 1 1000
```
