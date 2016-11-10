// package rpc2 is automatically generated from rpc2.x file by ...

package rpc2

import (
	"fmt"

	xdr "github.com/davecgh/go-xdr/xdr2"
	"github.com/gavrie/oncrpc/types"
)

// Prevent unused imports
var _ = types.ErrInvalidEnumOpt
var _ = xdr.ErrBadArguments

type enum int32 // RFC 1014

const MaxArrayLength = 0x7fffffff

// enum msg_type
type MsgType enum

const (
	Call  MsgType = 0
	Reply MsgType = 1
)

func (v MsgType) String() string {
	switch v {
	case Call:
		return "Call"
	case Reply:
		return "Reply"
	}
	return fmt.Sprintf("MsgType(%d)", v)
}

func (v MsgType) fieldCheck() error {
	switch v {
	case Call:
	case Reply:
	default:
		return types.ErrInvalidEnumOpt
	}
	return nil
}

func (v *MsgType) Encode(e *xdr.Encoder) error {
	if err := v.fieldCheck(); err != nil {
		return err
	}
	if _, err := e.EncodeInt(int32(*v)); err != nil {
		return err
	}
	return nil
}

func (v *MsgType) Decode(d *xdr.Decoder) error {
	intValue, _, err := d.DecodeInt()
	if err != nil {
		return err
	}
	*v = MsgType(intValue)
	return v.fieldCheck()
}

// enum reply_stat
type ReplyStat enum

const (
	MsgAccepted ReplyStat = 0
	MsgDenied   ReplyStat = 1
)

func (v ReplyStat) String() string {
	switch v {
	case MsgAccepted:
		return "MsgAccepted"
	case MsgDenied:
		return "MsgDenied"
	}
	return fmt.Sprintf("ReplyStat(%d)", v)
}

func (v ReplyStat) fieldCheck() error {
	switch v {
	case MsgAccepted:
	case MsgDenied:
	default:
		return types.ErrInvalidEnumOpt
	}
	return nil
}

func (v *ReplyStat) Encode(e *xdr.Encoder) error {
	if err := v.fieldCheck(); err != nil {
		return err
	}
	if _, err := e.EncodeInt(int32(*v)); err != nil {
		return err
	}
	return nil
}

func (v *ReplyStat) Decode(d *xdr.Decoder) error {
	intValue, _, err := d.DecodeInt()
	if err != nil {
		return err
	}
	*v = ReplyStat(intValue)
	return v.fieldCheck()
}

// enum accept_stat
type AcceptStat enum

const (
	Success      AcceptStat = 0
	ProgUnavail  AcceptStat = 1
	ProgMismatch AcceptStat = 2
	ProcUnavail  AcceptStat = 3
	GarbageArgs  AcceptStat = 4
)

func (v AcceptStat) String() string {
	switch v {
	case Success:
		return "Success"
	case ProgUnavail:
		return "ProgUnavail"
	case ProgMismatch:
		return "ProgMismatch"
	case ProcUnavail:
		return "ProcUnavail"
	case GarbageArgs:
		return "GarbageArgs"
	}
	return fmt.Sprintf("AcceptStat(%d)", v)
}

func (v AcceptStat) fieldCheck() error {
	switch v {
	case Success:
	case ProgUnavail:
	case ProgMismatch:
	case ProcUnavail:
	case GarbageArgs:
	default:
		return types.ErrInvalidEnumOpt
	}
	return nil
}

func (v *AcceptStat) Encode(e *xdr.Encoder) error {
	if err := v.fieldCheck(); err != nil {
		return err
	}
	if _, err := e.EncodeInt(int32(*v)); err != nil {
		return err
	}
	return nil
}

func (v *AcceptStat) Decode(d *xdr.Decoder) error {
	intValue, _, err := d.DecodeInt()
	if err != nil {
		return err
	}
	*v = AcceptStat(intValue)
	return v.fieldCheck()
}

// enum reject_stat
type RejectStat enum

const (
	RpcMismatch RejectStat = 0
	AuthError   RejectStat = 1
)

func (v RejectStat) String() string {
	switch v {
	case RpcMismatch:
		return "RpcMismatch"
	case AuthError:
		return "AuthError"
	}
	return fmt.Sprintf("RejectStat(%d)", v)
}

func (v RejectStat) fieldCheck() error {
	switch v {
	case RpcMismatch:
	case AuthError:
	default:
		return types.ErrInvalidEnumOpt
	}
	return nil
}

func (v *RejectStat) Encode(e *xdr.Encoder) error {
	if err := v.fieldCheck(); err != nil {
		return err
	}
	if _, err := e.EncodeInt(int32(*v)); err != nil {
		return err
	}
	return nil
}

func (v *RejectStat) Decode(d *xdr.Decoder) error {
	intValue, _, err := d.DecodeInt()
	if err != nil {
		return err
	}
	*v = RejectStat(intValue)
	return v.fieldCheck()
}

// enum auth_stat
type AuthStat enum

const (
	AuthBadcred      AuthStat = 1
	AuthRejectedcred AuthStat = 2
	AuthBadverf      AuthStat = 3
	AuthRejectedverf AuthStat = 4
	AuthTooweak      AuthStat = 5
)

func (v AuthStat) String() string {
	switch v {
	case AuthBadcred:
		return "AuthBadcred"
	case AuthRejectedcred:
		return "AuthRejectedcred"
	case AuthBadverf:
		return "AuthBadverf"
	case AuthRejectedverf:
		return "AuthRejectedverf"
	case AuthTooweak:
		return "AuthTooweak"
	}
	return fmt.Sprintf("AuthStat(%d)", v)
}

func (v AuthStat) fieldCheck() error {
	switch v {
	case AuthBadcred:
	case AuthRejectedcred:
	case AuthBadverf:
	case AuthRejectedverf:
	case AuthTooweak:
	default:
		return types.ErrInvalidEnumOpt
	}
	return nil
}

func (v *AuthStat) Encode(e *xdr.Encoder) error {
	if err := v.fieldCheck(); err != nil {
		return err
	}
	if _, err := e.EncodeInt(int32(*v)); err != nil {
		return err
	}
	return nil
}

func (v *AuthStat) Decode(d *xdr.Decoder) error {
	intValue, _, err := d.DecodeInt()
	if err != nil {
		return err
	}
	*v = AuthStat(intValue)
	return v.fieldCheck()
}

// enum auth_flavor
type AuthFlavor enum

const (
	AuthNull  AuthFlavor = 0
	AuthUnix  AuthFlavor = 1
	AuthShort AuthFlavor = 2
	AuthDes   AuthFlavor = 3
)

func (v AuthFlavor) String() string {
	switch v {
	case AuthNull:
		return "AuthNull"
	case AuthUnix:
		return "AuthUnix"
	case AuthShort:
		return "AuthShort"
	case AuthDes:
		return "AuthDes"
	}
	return fmt.Sprintf("AuthFlavor(%d)", v)
}

func (v AuthFlavor) fieldCheck() error {
	switch v {
	case AuthNull:
	case AuthUnix:
	case AuthShort:
	case AuthDes:
	default:
		return types.ErrInvalidEnumOpt
	}
	return nil
}

func (v *AuthFlavor) Encode(e *xdr.Encoder) error {
	if err := v.fieldCheck(); err != nil {
		return err
	}
	if _, err := e.EncodeInt(int32(*v)); err != nil {
		return err
	}
	return nil
}

func (v *AuthFlavor) Decode(d *xdr.Decoder) error {
	intValue, _, err := d.DecodeInt()
	if err != nil {
		return err
	}
	*v = AuthFlavor(intValue)
	return v.fieldCheck()
}

// struct opaque_auth
type OpaqueAuth struct {
	Flavor AuthFlavor
	Body   []byte // Max length: 400

}

func (s *OpaqueAuth) Encode(e *xdr.Encoder) error {

	if err := s.Flavor.Encode(e); err != nil {
		return err
	}

	// s.Body: []byte <400>
	if len(s.Body) > 400 {
		return types.ErrArrayTooLarge
	}
	if _, err := e.EncodeOpaque(s.Body); err != nil {
		return err
	}

	return nil
}

// struct call_body
type CallBody struct {
	Rpcvers uint32
	Prog    uint32
	Vers    uint32
	Proc    uint32
	Cred    OpaqueAuth
	Verf    OpaqueAuth
}

func (s *CallBody) Encode(e *xdr.Encoder) error {
	if _, err := e.EncodeUint(s.Rpcvers); err != nil {
		return err
	}
	if _, err := e.EncodeUint(s.Prog); err != nil {
		return err
	}
	if _, err := e.EncodeUint(s.Vers); err != nil {
		return err
	}
	if _, err := e.EncodeUint(s.Proc); err != nil {
		return err
	}

	if err := s.Cred.Encode(e); err != nil {
		return err
	}

	if err := s.Verf.Encode(e); err != nil {
		return err
	}

	return nil
}

// struct mismatch_info_body
type MismatchInfoBody struct {
	Low  uint32
	High uint32
}

func (s *MismatchInfoBody) Encode(e *xdr.Encoder) error {
	if _, err := e.EncodeUint(s.Low); err != nil {
		return err
	}
	if _, err := e.EncodeUint(s.High); err != nil {
		return err
	}

	return nil
}

// union reply_data_body

type ReplyDataBody struct {
	Stat  AcceptStat // Arbitrator
	Union replyDataBodyUnion
}

type replyDataBodyUnion interface {
	isReplyDataBodyUnion()
}

func (MismatchInfoBody) isReplyDataBodyUnion() {}

func (s *ReplyDataBody) Encode(e *xdr.Encoder) error {
	// Arbitrator

	if err := s.Stat.Encode(e); err != nil {
		return err
	}

	switch s.Stat {
	case Success:

		// Voidcase ProgMismatch:
		u, ok := s.Union.(MismatchInfoBody)
		if !ok {
			return types.ErrArbitratorValueMismatch
		}
		if err := u.Encode(e); err != nil {
			return err
		}
	default:
		// Void
	}

	return nil
}

func (s *ReplyDataBody) Decode(d *xdr.Decoder) error {
	return nil
}

// struct accepted_reply
type AcceptedReply struct {
	Verf      OpaqueAuth
	ReplyData ReplyDataBody
}

func (s *AcceptedReply) Encode(e *xdr.Encoder) error {

	if err := s.Verf.Encode(e); err != nil {
		return err
	}

	if err := s.ReplyData.Encode(e); err != nil {
		return err
	}

	return nil
}

// union rejected_reply

type RejectedReply struct {
	Stat  RejectStat // Arbitrator
	Union rejectedReplyUnion
}

type rejectedReplyUnion interface {
	isRejectedReplyUnion()
}

func (MismatchInfoBody) isRejectedReplyUnion() {}

func (AuthStat) isRejectedReplyUnion() {}

func (s *RejectedReply) Encode(e *xdr.Encoder) error {
	// Arbitrator

	if err := s.Stat.Encode(e); err != nil {
		return err
	}

	switch s.Stat {
	case RpcMismatch:
		u, ok := s.Union.(MismatchInfoBody)
		if !ok {
			return types.ErrArbitratorValueMismatch
		}
		if err := u.Encode(e); err != nil {
			return err
		}
	case AuthError:
		u, ok := s.Union.(AuthStat)
		if !ok {
			return types.ErrArbitratorValueMismatch
		}
		if err := u.Encode(e); err != nil {
			return err
		}

	}

	return nil
}

func (s *RejectedReply) Decode(d *xdr.Decoder) error {
	return nil
}

// union reply_body

type ReplyBody struct {
	Stat  ReplyStat // Arbitrator
	Union replyBodyUnion
}

type replyBodyUnion interface {
	isReplyBodyUnion()
}

func (AcceptedReply) isReplyBodyUnion() {}

func (RejectedReply) isReplyBodyUnion() {}

func (s *ReplyBody) Encode(e *xdr.Encoder) error {
	// Arbitrator

	if err := s.Stat.Encode(e); err != nil {
		return err
	}

	switch s.Stat {
	case MsgAccepted:
		u, ok := s.Union.(AcceptedReply)
		if !ok {
			return types.ErrArbitratorValueMismatch
		}
		if err := u.Encode(e); err != nil {
			return err
		}
	case MsgDenied:
		u, ok := s.Union.(RejectedReply)
		if !ok {
			return types.ErrArbitratorValueMismatch
		}
		if err := u.Encode(e); err != nil {
			return err
		}

	}

	return nil
}

func (s *ReplyBody) Decode(d *xdr.Decoder) error {
	return nil
}

// union rpc_body

type RpcBody struct {
	Mtype MsgType // Arbitrator
	Union rpcBodyUnion
}

type rpcBodyUnion interface {
	isRpcBodyUnion()
}

func (CallBody) isRpcBodyUnion() {}

func (ReplyBody) isRpcBodyUnion() {}

func (s *RpcBody) Encode(e *xdr.Encoder) error {
	// Arbitrator

	if err := s.Mtype.Encode(e); err != nil {
		return err
	}

	switch s.Mtype {
	case Call:
		u, ok := s.Union.(CallBody)
		if !ok {
			return types.ErrArbitratorValueMismatch
		}
		if err := u.Encode(e); err != nil {
			return err
		}
	case Reply:
		u, ok := s.Union.(ReplyBody)
		if !ok {
			return types.ErrArbitratorValueMismatch
		}
		if err := u.Encode(e); err != nil {
			return err
		}

	}

	return nil
}

func (s *RpcBody) Decode(d *xdr.Decoder) error {
	return nil
}

// struct rpc_msg
type RpcMsg struct {
	Xid  uint32
	Body RpcBody
}

func (s *RpcMsg) Encode(e *xdr.Encoder) error {
	if _, err := e.EncodeUint(s.Xid); err != nil {
		return err
	}

	if err := s.Body.Encode(e); err != nil {
		return err
	}

	return nil
}

// struct auth_unix_data
type AuthUnixData struct {
	Stamp       uint32
	Machinename string // Max length: 255
	Uid         uint32
	Gid         uint32
	Gids        []uint32 // Max length: 16

}

func (s *AuthUnixData) Encode(e *xdr.Encoder) error {
	if _, err := e.EncodeUint(s.Stamp); err != nil {
		return err
	}

	// s.Machinename: string <255>
	if len(s.Machinename) > 255 {
		return types.ErrArrayTooLarge
	}
	if _, err := e.EncodeString(s.Machinename); err != nil {
		return err
	}

	if _, err := e.EncodeUint(s.Uid); err != nil {
		return err
	}
	if _, err := e.EncodeUint(s.Gid); err != nil {
		return err
	}

	// s.Gids: []uint32 <16>
	{
		dataLength := uint32(len(s.Gids))

		if dataLength > 16 {
			return types.ErrArrayTooLarge
		}

		if _, err := e.EncodeUint(dataLength); err != nil {
			return err
		}

		for _, value := range s.Gids {
			if _, err := e.EncodeUint(value); err != nil {
				return err
			}
		}
	}

	return nil
}
