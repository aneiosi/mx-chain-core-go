package testscommon

import (
	"github.com/multiversx/mx-chain-core-go/webSocket"
)

// WebSocketTransceiverStub -
type WebSocketTransceiverStub struct {
	SendCalled              func(payload []byte, topic string, conn webSocket.WSConClient) error
	CloseCalled             func() error
	SetPayloadHandlerCalled func(handler webSocket.PayloadHandler) error
	ListenCalled            func(conn webSocket.WSConClient) (closed bool)
}

// Send -
func (w *WebSocketTransceiverStub) Send(payload []byte, topic string, conn webSocket.WSConClient) error {
	if w.SendCalled != nil {
		return w.SendCalled(payload, topic, conn)
	}
	return nil
}

// Close -
func (w *WebSocketTransceiverStub) Close() error {
	if w.CloseCalled != nil {
		return w.CloseCalled()
	}
	return nil
}

// SetPayloadHandler -
func (w *WebSocketTransceiverStub) SetPayloadHandler(handler webSocket.PayloadHandler) error {
	if w.SetPayloadHandlerCalled != nil {
		return w.SetPayloadHandlerCalled(handler)
	}

	return nil
}

// Listen -
func (w *WebSocketTransceiverStub) Listen(conn webSocket.WSConClient) (closed bool) {
	if w.ListenCalled != nil {
		return w.ListenCalled(conn)
	}
	return false
}
