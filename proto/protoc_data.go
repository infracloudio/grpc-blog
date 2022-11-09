package main

import "github.com/infracloudio/grpc-blog/proto/protoc"

var (
	protocSmall = protoc.BenchSmall{
		Action: "benchmark",
		Key:    []byte("data to be sent"),
	}

	protocMedium = protoc.BenchMedium{
		Name:   "Tester",
		Age:    20,
		Height: 5.8,
		Weight: 180.7,
		Alive:  true,
		Desc: []byte(`If you’ve ever heard of ProtoBuf you may be thinking 
		that the results of this benchmarking experiment will be obvious,
		JSON < ProtoBuf.`),
	}

	protocLarge = protoc.BenchLarge{
		Name:     "Tester",
		Age:      20,
		Height:   5.8,
		Weight:   180.7,
		Alive:    true,
		Desc:     []byte("Lets benchmark some json and protobuf"),
		Nickname: "Another name",
		Num:      2314,
		Flt:      123451231.1234,
		Data: []byte(`If you’ve ever heard of ProtoBuf you may be thinking that
		the results of this benchmarking experiment will be obvious, JSON < ProtoBuf.
		My interest was in how much they actually differ in practice.
		How do they compare on a couple of different metrics, specifically serialization
		and de-serialization speeds, and the memory footprint of encoding the data.
		I was also curious about how the different serialization methods would
		behave under small, medium, and large chunks of data.`),
	}
)
