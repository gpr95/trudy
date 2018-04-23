package module

import (
	"crypto/tls"
	"encoding/hex"
	"net"
	"strings"

	"github.com/gpr95/trudy/pipe"
)

//Data is a thin wrapper that provides metadata that may be useful when mangling bytes on the network.
type Data struct {
	FromClient bool        //FromClient is true is the data sent is coming from the client (the device you are proxying)
	Bytes      []byte      //Bytes is a byte slice that contians the TCP data
	TLSConfig  *tls.Config //TLSConfig is a TLS server config that contains Trudy's TLS server certficiate.
	ServerAddr net.Addr    //ServerAddr is net.Addr of the server
	ClientAddr net.Addr    //ClientAddr is the net.Addr of the client (the device you are proxying)
}

//DoMangle will return true if Data needs to be sent to the Mangle function.
func (input Data) DoMangle() bool {
	if strings.Contains(input.ServerAddr.String(), "172.217") {
		return true
	}
	return false
}

//Mangle can modify/replace the Bytes values within the Data struct. This can
//be empty if no programmatic mangling needs to be done.
func (input *Data) Mangle() {
	input.ServerAddr, _ = net.ResolveTCPAddr("tcp", "192.168.0.1")
		if len(input.Bytes) > 51 && (input.Bytes[47] & 0x10) == 0x10 {
			input.Bytes[52] = 0xFF
			input.Bytes[53] = 0xFF
		}
}

//Drop will return true if the Data needs to be dropped before going through
//the pipe.
func (input Data) Drop() bool {
	return false
}

//PrettyPrint returns the string representation of the data. This string will
//be the value that is logged to the console.
func (input Data) PrettyPrint() string {
	return hex.Dump(input.Bytes)
}

//DoPrint will return true if the PrettyPrinted version of the Data struct
//needs to be logged to the console.
func (input Data) DoPrint() bool {
	return true
}

//DoIntercept returns true if data should be sent to the Trudy interceptor.
func (input Data) DoIntercept() bool {
	return false
}

//Deserialize should replace the Data struct's Bytes with a deserialized bytes.
//For example, unpacking a HTTP/2 frame would be deserialization.
func (input *Data) Deserialize() {

}

//Serialize should replace the Data struct's Bytes with the serialized form of
//the bytes. The serialized bytes will be sent over the wire.
func (input *Data) Serialize() {

}

//BeforeWriteToClient is a function that will be called before data is sent to
//a client.
func (input *Data) BeforeWriteToClient(p pipe.Pipe) {

}

//AfterWriteToClient is a function that will be called after data is sent to
//a client.
func (input *Data) AfterWriteToClient(p pipe.Pipe) {

}

//BeforeWriteToServer is a function that will be called before data is sent to
//a server.
func (input *Data) BeforeWriteToServer(p pipe.Pipe) {

}

//AfterWriteToServer is a function that will be called after data is sent to
//a server.
func (input *Data) AfterWriteToServer(p pipe.Pipe) {

}
