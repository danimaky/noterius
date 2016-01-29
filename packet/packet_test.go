package packet

import (
	"testing"

	"github.com/Nyarum/noterius/core"
	"github.com/Nyarum/noterius/entitie"
)

func BenchmarkEncode940Packet(b *testing.B) {
	buffers := core.NewBuffers()
	player := entitie.NewPlayer(buffers)
	packet := NewPacket(player)

	for n := 0; n < b.N; n++ {
		packet.Encode(940)
	}
}

func BenchmarkEncode931Packet(b *testing.B) {
	buffers := core.NewBuffers()
	player := entitie.NewPlayer(buffers)
	packet := NewPacket(player)

	for n := 0; n < b.N; n++ {
		packet.Encode(931)
	}
}