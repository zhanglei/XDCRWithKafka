package aof

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import "github.com/tinylib/msgp/msgp"

// DecodeMsg implements msgp.Decodable
func (z *Operation) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "command":
			z.Command, err = dc.ReadString()
			if err != nil {
				return
			}
		case "subop":
			z.SubOp, err = dc.ReadString()
			if err != nil {
				return
			}
		case "key":
			z.Key, err = dc.ReadString()
			if err != nil {
				return
			}
		case "arguments":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Arguments) >= int(zb0002) {
				z.Arguments = (z.Arguments)[:zb0002]
			} else {
				z.Arguments = make([]string, zb0002)
			}
			for za0001 := range z.Arguments {
				z.Arguments[za0001], err = dc.ReadString()
				if err != nil {
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Operation) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "command"
	err = en.Append(0x84, 0xa7, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Command)
	if err != nil {
		return
	}
	// write "subop"
	err = en.Append(0xa5, 0x73, 0x75, 0x62, 0x6f, 0x70)
	if err != nil {
		return err
	}
	err = en.WriteString(z.SubOp)
	if err != nil {
		return
	}
	// write "key"
	err = en.Append(0xa3, 0x6b, 0x65, 0x79)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Key)
	if err != nil {
		return
	}
	// write "arguments"
	err = en.Append(0xa9, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Arguments)))
	if err != nil {
		return
	}
	for za0001 := range z.Arguments {
		err = en.WriteString(z.Arguments[za0001])
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Operation) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "command"
	o = append(o, 0x84, 0xa7, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64)
	o = msgp.AppendString(o, z.Command)
	// string "subop"
	o = append(o, 0xa5, 0x73, 0x75, 0x62, 0x6f, 0x70)
	o = msgp.AppendString(o, z.SubOp)
	// string "key"
	o = append(o, 0xa3, 0x6b, 0x65, 0x79)
	o = msgp.AppendString(o, z.Key)
	// string "arguments"
	o = append(o, 0xa9, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Arguments)))
	for za0001 := range z.Arguments {
		o = msgp.AppendString(o, z.Arguments[za0001])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Operation) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "command":
			z.Command, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "subop":
			z.SubOp, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "key":
			z.Key, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "arguments":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Arguments) >= int(zb0002) {
				z.Arguments = (z.Arguments)[:zb0002]
			} else {
				z.Arguments = make([]string, zb0002)
			}
			for za0001 := range z.Arguments {
				z.Arguments[za0001], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Operation) Msgsize() (s int) {
	s = 1 + 8 + msgp.StringPrefixSize + len(z.Command) + 6 + msgp.StringPrefixSize + len(z.SubOp) + 4 + msgp.StringPrefixSize + len(z.Key) + 10 + msgp.ArrayHeaderSize
	for za0001 := range z.Arguments {
		s += msgp.StringPrefixSize + len(z.Arguments[za0001])
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *UnexpectedEOF) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z UnexpectedEOF) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 0
	err = en.Append(0x80)
	if err != nil {
		return err
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z UnexpectedEOF) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 0
	o = append(o, 0x80)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *UnexpectedEOF) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z UnexpectedEOF) Msgsize() (s int) {
	s = 1
	return
}
