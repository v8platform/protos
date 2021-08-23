package rasbinary

import (
	messagesv1 "github.com/v8platform/protos/gen/ras/messages/v1"
	"google.golang.org/protobuf/proto"
	"reflect"
	"testing"
)

func TestGetMessageType(t *testing.T) {

	tests := []struct {
		name  string
		m     proto.Message
		want  messagesv1.MessageType
		want1 bool
	}{
		{
			"message",
			&messagesv1.GetClustersRequest{},
			messagesv1.MessageType_GET_CLUSTERS_REQUEST,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetMessageType(tt.m)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMessageType() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetMessageType() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
