package amqp

import (
	"sync"

	"github.com/up9inc/mizu/tap/api"
)

type tcpStream struct {
	isClosed        bool
	protoIdentifier *api.ProtoIdentifier
	isTapTarget     bool
	origin          api.Capture
	reqResMatchers  []api.RequestResponseMatcher
	sync.Mutex
}

func NewTcpStream(capture api.Capture) api.TcpStream {
	return &tcpStream{
		origin:          capture,
		protoIdentifier: &api.ProtoIdentifier{},
	}
}

func (t *tcpStream) SetProtocol(protocol *api.Protocol) {}

func (t *tcpStream) GetOrigin() api.Capture {
	return t.origin
}

func (t *tcpStream) GetProtoIdentifier() *api.ProtoIdentifier {
	return t.protoIdentifier
}

func (t *tcpStream) GetReqResMatchers() []api.RequestResponseMatcher {
	return t.reqResMatchers
}

func (t *tcpStream) GetIsTapTarget() bool {
	return t.isTapTarget
}

func (t *tcpStream) GetIsClosed() bool {
	return t.isClosed
}
