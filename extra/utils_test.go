package extra

import (
	protocolv1 "github.com/v8platform/protos/gen/ras/protocol/v1"
	"google.golang.org/protobuf/proto"
	"reflect"
	"testing"
)

func TestUnpackPacketDataNew(t *testing.T) {
	type args struct {
		p *protocolv1.Packet
	}
	tests := []struct {
		name    string
		p       *protocolv1.Packet
		want    proto.Message
		wantErr bool
	}{
		{"name",
			&protocolv1.Packet{
				Type: protocolv1.PacketType_PACKET_TYPE_ENDPOINT_OPEN_ACK,
			},
			&protocolv1.EndpointOpenAck{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnpackPacketDataNew(tt.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnpackPacketDataNew() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnpackPacketDataNew() got = %v, want %v", got, tt.want)
			}
		})
	}
}
