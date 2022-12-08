package main

import (
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/infracloudio/grpc-blog/proto/gogo_protobuf"
	"github.com/infracloudio/grpc-blog/proto/protoc"
	"testing"
)

func BenchmarkJSONMarshal(b *testing.B) {
	s := jsonSmall
	m := jsonMedium
	l := jsonLarge

	b.ResetTimer()

	b.Run("SmallData", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			d, _ := json.Marshal(&s)
			_ = d
		}
		b.ReportAllocs()
	})
	b.Run("MediumData", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			d, _ := json.Marshal(&m)
			_ = d
		}
		b.ReportAllocs()
	})
	b.Run("LargeData", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			d, _ := json.Marshal(&l)
			_ = d
		}
		b.ReportAllocs()
	})
	fmt.Printf("\n")
}

func BenchmarkJSONUnmarshal(b *testing.B) {
	s := jsonSmall
	m := jsonMedium
	l := jsonLarge

	sd, _ := json.Marshal(&s)
	md, _ := json.Marshal(&m)
	ld, _ := json.Marshal(&l)

	var sf []byte
	var mf []byte
	var lf []byte

	b.ResetTimer()

	b.Run("SmallData", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = json.Unmarshal(sd, &sf)
		}
		b.ReportAllocs()
	})
	b.Run("MediumData", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = json.Unmarshal(md, &mf)
		}
		b.ReportAllocs()
	})
	b.Run("LargeData", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = json.Unmarshal(ld, &lf)
		}
		b.ReportAllocs()
	})
	fmt.Printf("\n")
}

func BenchmarkProtocMarshal(b *testing.B) {
	s := protocSmall
	m := protocMedium
	l := protocLarge

	b.ResetTimer()

	b.Run("SmallData", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			d, _ := proto.Marshal(&s)
			_ = d
		}
		b.ReportAllocs()
	})
	b.Run("MediumData", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			d, _ := proto.Marshal(&m)
			_ = d
		}
		b.ReportAllocs()
	})
	b.Run("LargeData", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			d, _ := proto.Marshal(&l)
			_ = d
		}
		b.ReportAllocs()
	})
	fmt.Printf("\n")
}

func BenchmarkGogoPbMarshal(b *testing.B) {
	s := gogoPbSmall
	m := gogoPbMedium
	l := gogoPbLarge

	b.ResetTimer()

	b.Run("SmallData", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			d, _ := proto.Marshal(&s)
			_ = d
		}
		b.ReportAllocs()
	})
	b.Run("MediumData", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			d, _ := proto.Marshal(&m)
			_ = d
		}
		b.ReportAllocs()
	})
	b.Run("LargeData", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			d, _ := proto.Marshal(&l)
			_ = d
		}
		b.ReportAllocs()
	})
	fmt.Printf("\n")
}

func BenchmarkProtocUnMarshal(b *testing.B) {
	s := protocSmall
	m := protocMedium
	l := protocLarge

	sd, _ := proto.Marshal(&s)
	md, _ := proto.Marshal(&m)
	ld, _ := proto.Marshal(&l)

	var sf protoc.BenchSmall
	var mf protoc.BenchMedium
	var lf protoc.BenchLarge

	b.ResetTimer()

	b.Run("SmallData", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = proto.Unmarshal(sd, &sf)
		}
		b.ReportAllocs()
	})
	b.Run("MediumData", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = proto.Unmarshal(md, &mf)
		}
		b.ReportAllocs()
	})
	b.Run("LargeData", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = proto.Unmarshal(ld, &lf)
		}
		b.ReportAllocs()
	})
	fmt.Printf("\n")
}

func BenchmarkGogoPbUnMarshal(b *testing.B) {
	s := gogoPbSmall
	m := gogoPbMedium
	l := gogoPbLarge

	sd, _ := proto.Marshal(&s)
	md, _ := proto.Marshal(&m)
	ld, _ := proto.Marshal(&l)

	var sf gogo_protobuf.BenchSmall
	var mf gogo_protobuf.BenchMedium
	var lf gogo_protobuf.BenchLarge

	b.ResetTimer()

	b.Run("SmallData", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = proto.Unmarshal(sd, &sf)
		}
		b.ReportAllocs()
	})
	b.Run("MediumData", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = proto.Unmarshal(md, &mf)
		}
		b.ReportAllocs()
	})
	b.Run("LargeData", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = proto.Unmarshal(ld, &lf)
		}
		b.ReportAllocs()
	})
	fmt.Printf("\n")
}
