package easyws

import (
	"context"
	"easynet"
	_interface "easynet/interface"

	//_interface "github.com/gomystery/easynet/interface"
	//"github.com/gomystery/easynet"
	"fmt"
	"io"
)


type Handler struct {
}

func (h Handler) OnStart(conn interface{}) error {
	fmt.Println("test OnStart")
	return nil
}

func (h Handler) OnConnect(conn interface{}) error {
	return nil

}


func (h Handler) OnReceive(conn interface{}, stream _interface.IInputStream) ([]byte, error) {

	// todo handover

	// todo frame

	return []byte(""), nil

}

func (h Handler) OnShutdown(conn interface{}) error {
	return nil
}

func (h Handler) OnClose(conn interface{}, err error) error {
	return nil
}


func NewEasyWs() *EasyWs {
	config := easynet.NewDefaultNetConfig("tcp", "127.0.0.1", 9011)
	handler := &Handler{}
	net := easynet.NewEasyNet(context.Background(), "NetPoll", config, handler)
	ws :=&EasyWs{
		Handler: handler,
		EasyNet:net,
	}

	return ws
}

type EasyWs struct {

	Handler  *Handler

	EasyNet *easynet.EasyNet


}




// HandshakeHeader is the interface that writes both upgrade request or
// response headers into a given io.Writer.
type HandshakeHeader interface {
	io.WriterTo
}

// Upgrader contains options for upgrading connection to websocket.
type Upgrader struct {
	// ReadBufferSize and WriteBufferSize is an I/O buffer sizes.
	// They used to read and write http data while upgrading to WebSocket.
	// Allocated buffers are pooled with sync.Pool to avoid extra allocations.
	//
	// If a size is zero then default value is used.
	//
	// Usually it is useful to set read buffer size bigger than write buffer
	// size because incoming request could contain long header values, such as
	// Cookie. Response, in other way, could be big only if user write multiple
	// custom headers. Usually response takes less than 256 bytes.
	ReadBufferSize, WriteBufferSize int

	// Protocol is a select function that is used to select subprotocol
	// from list requested by client. If this field is set, then the first matched
	// protocol is sent to a client as negotiated.
	//
	// The argument is only valid until the callback returns.
	Protocol func([]byte) bool

	// ProtocolCustrom allow user to parse Sec-WebSocket-Protocol header manually.
	// Note that returned bytes must be valid until Upgrade returns.
	// If ProtocolCustom is set, it used instead of Protocol function.
	ProtocolCustom func([]byte) (string, bool)

	// Extension is a select function that is used to select extensions
	// from list requested by client. If this field is set, then the all matched
	// extensions are sent to a client as negotiated.
	//
	// Note that Extension may be called multiple times and implementations
	// must track uniqueness of accepted extensions manually.
	//
	// The argument is only valid until the callback returns.
	//
	// According to the RFC6455 order of extensions passed by a client is
	// significant. That is, returning true from this function means that no
	// other extension with the same name should be checked because server
	// accepted the most preferable extension right now:
	// "Note that the order of extensions is significant.  Any interactions between
	// multiple extensions MAY be defined in the documents defining the extensions.
	// In the absence of such definitions, the interpretation is that the header
	// fields listed by the client in its request represent a preference of the
	// header fields it wishes to use, with the first options listed being most
	// preferable."
	//
	// Deprecated: use Negotiate instead.
	//Extension func(httphead.Option) bool

	// ExtensionCustom allow user to parse Sec-WebSocket-Extensions header
	// manually.
	//
	// If ExtensionCustom() decides to accept received extension, it must
	// append appropriate option to the given slice of httphead.Option.
	// It returns results of append() to the given slice and a flag that
	// reports whether given header value is wellformed or not.
	//
	// Note that ExtensionCustom may be called multiple times and
	// implementations must track uniqueness of accepted extensions manually.
	//
	// Note that returned options should be valid until Upgrade returns.
	// If ExtensionCustom is set, it used instead of Extension function.
	//ExtensionCustom func([]byte, []httphead.Option) ([]httphead.Option, bool)

	// Negotiate is the callback that is used to negotiate extensions from
	// the client's offer. If this field is set, then the returned non-zero
	// extensions are sent to the client as accepted extensions in the
	// response.
	//
	// The argument is only valid until the Negotiate callback returns.
	//
	// If returned error is non-nil then connection is rejected and response is
	// sent with appropriate HTTP error code and body set to error message.
	//
	// RejectConnectionError could be used to get more control on response.
	//Negotiate func(httphead.Option) (httphead.Option, error)

	// Header is an optional HandshakeHeader instance that could be used to
	// write additional headers to the handshake response.
	//
	// It used instead of any key-value mappings to avoid allocations in user
	// land.
	//
	// Note that if present, it will be written in any result of handshake.
	Header HandshakeHeader

	// OnRequest is a callback that will be called after request line
	// successful parsing.
	//
	// The arguments are only valid until the callback returns.
	//
	// If returned error is non-nil then connection is rejected and response is
	// sent with appropriate HTTP error code and body set to error message.
	//
	// RejectConnectionError could be used to get more control on response.
	OnRequest func(uri []byte) error

	// OnHost is a callback that will be called after "Host" header successful
	// parsing.
	//
	// It is separated from OnHeader callback because the Host header must be
	// present in each request since HTTP/1.1. Thus Host header is non-optional
	// and required for every WebSocket handshake.
	//
	// The arguments are only valid until the callback returns.
	//
	// If returned error is non-nil then connection is rejected and response is
	// sent with appropriate HTTP error code and body set to error message.
	//
	// RejectConnectionError could be used to get more control on response.
	OnHost func(host []byte) error

	// OnHeader is a callback that will be called after successful parsing of
	// header, that is not used during WebSocket handshake procedure. That is,
	// it will be called with non-websocket headers, which could be relevant
	// for application-level logic.
	//
	// The arguments are only valid until the callback returns.
	//
	// If returned error is non-nil then connection is rejected and response is
	// sent with appropriate HTTP error code and body set to error message.
	//
	// RejectConnectionError could be used to get more control on response.
	OnHeader func(key, value []byte) error

	// OnBeforeUpgrade is a callback that will be called before sending
	// successful upgrade response.
	//
	// Setting OnBeforeUpgrade allows user to make final application-level
	// checks and decide whether this connection is allowed to successfully
	// upgrade to WebSocket.
	//
	// It must return non-nil either HandshakeHeader or error and never both.
	//
	// If returned error is non-nil then connection is rejected and response is
	// sent with appropriate HTTP error code and body set to error message.
	//
	// RejectConnectionError could be used to get more control on response.
	//OnBeforeUpgrade func() (header HandshakeHeader, err error)
}