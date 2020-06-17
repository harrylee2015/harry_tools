package redis

import "testing"

func TestExampleClient(t *testing.T) {
	ExampleClient()
	ExampleConn()
	ExampleClient_Incr()
	ExampleClient_Incr()
	ExampleClient_BLPop()
	ExampleClient_Scan()
	ExampleClient_Pipelined()
	ExampleClient_Pipelined()
	ExampleClient_Pipeline()


}