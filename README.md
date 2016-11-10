# oncrpc: An ONC RPC v2 implementation for Go

The `oncrpc` package generates Go code that implements the data types, encoders, and decoders, for protocols described in the XDR data description language. The language is defined by [RFC 5531](https://tools.ietf.org/html/rfc5531), which turn is an extension of [RFC 4506](https://tools.ietf.org/html/rfc4506)).

This is accomplished by using an XDR parser that has been generated with [ANTLR](http://www.antlr.org/), and can be easily re-generated if the grammar needs to be updated or extended. The parser generates a Go implementation that is tailored to encoding and decoding data described by the used-provided XDR language file.

For example, to create a complete Go encoder and decoder for the NFSv3 protocol, we can feed the parser with the following inputs, which are distilled from the relevant RFCs:
- rpc2.x: RPCv2, [RFC 5531](https://tools.ietf.org/html/rfc5531)
- portmap.x: Portmapper, [RFC 1833](https://tools.ietf.org/html/rfc1833)
- mount.x: Mount, [RFC 1813](https://tools.ietf.org/html/rfc1813)
- nfs.x: NFSv3, [RFC 1813](https://tools.ietf.org/html/rfc1813)

The implementation leverages the Go type system to create code that is statically typed. This makes the code efficient (no reflection is used). More importantly, it allows writing typesafe code that detects errors at compile time rather than at run time.

## Planned Tasks

1. Implement decoding (partially complete)
1. Add tests
1. Replace `go-xdr` with native functions based on `encoding/binary`
1. Implement an NFSv3 client

## References

- [davecgh/go-xdr](https://github.com/davecgh/go-xdr): Implements the data representation part of RFC 4506.
- [Packetbeat NFS decoder](https://github.com/elastic/beats/tree/master/packetbeat/protos/nfs): Partially decodes NFS3/4 packets for network sniffing.
- [calmh/xdr](https://github.com/calmh/xdr)
