package main

import (
	"bytes"
	"log"
	"testing"

	xdr "github.com/davecgh/go-xdr/xdr2"
	"github.com/gavrie/oncrpc/rpc2"
)

func TestRpcMsg(t *testing.T) {
	var in, out rpc2.RpcMsg
	var buf bytes.Buffer
	var n int
	var err error

	encoder := xdr.NewEncoder(&buf)
	decoder := xdr.NewDecoder(&buf)

	in = rpc2.RpcMsg{
		Xid: 0x12345678,
		Body: rpc2.RpcBody{
			Mtype: rpc2.Call,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100005, // nfs.ProgramMountprogmountV3,
				Vers:    1,
				Proc:    5, // Mountproc3Export,
				Cred:    rpc2.OpaqueAuth{Flavor: rpc2.AuthUnix, Body: []uint8{0x41, 0x42, 0x43, 0x44}},
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AuthNull},
			},
		},
	}

	log.Printf("encode")
	log.Printf("in: %+v", in)
	err = in.Encode(encoder)
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("n: %v, err: %v", n, err)

	log.Printf("buf: %v", buf.Bytes())

	_ = out
	_ = decoder
	/*
		log.Printf("decode")
		err = out.Decode(decoder)
		if err != nil {
			t.Fatal(err)
		}
		log.Printf("n: %v, err: %v", n, err)
		log.Printf("out: %+v", out)
	*/
}
