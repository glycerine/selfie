package selfie

// NOTE: THIS FILE WAS PRODUCED BY THE
// GREENPACK CODE GENERATION TOOL (github.com/glycerine/greenpack)
// DO NOT EDIT

import (
	"github.com/glycerine/greenpack/msgp"
)

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *Selfie) DecodeMsg(dc *msgp.Reader) (err error) {
	var sawTopNil bool
	if dc.IsNil() {
		sawTopNil = true
		err = dc.ReadNil()
		if err != nil {
			return
		}
		dc.PushAlwaysNil()
	}

	var field []byte
	_ = field
	const maxFields0zgensym_6a7ebc1181a5fc8f_1 = 3

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0zgensym_6a7ebc1181a5fc8f_1 uint32
	totalEncodedFields0zgensym_6a7ebc1181a5fc8f_1, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zgensym_6a7ebc1181a5fc8f_1 := totalEncodedFields0zgensym_6a7ebc1181a5fc8f_1
	missingFieldsLeft0zgensym_6a7ebc1181a5fc8f_1 := maxFields0zgensym_6a7ebc1181a5fc8f_1 - totalEncodedFields0zgensym_6a7ebc1181a5fc8f_1

	var nextMiss0zgensym_6a7ebc1181a5fc8f_1 int32 = -1
	var found0zgensym_6a7ebc1181a5fc8f_1 [maxFields0zgensym_6a7ebc1181a5fc8f_1]bool
	var curField0zgensym_6a7ebc1181a5fc8f_1 string

doneWithStruct0zgensym_6a7ebc1181a5fc8f_1:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zgensym_6a7ebc1181a5fc8f_1 > 0 || missingFieldsLeft0zgensym_6a7ebc1181a5fc8f_1 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zgensym_6a7ebc1181a5fc8f_1, missingFieldsLeft0zgensym_6a7ebc1181a5fc8f_1, msgp.ShowFound(found0zgensym_6a7ebc1181a5fc8f_1[:]), decodeMsgFieldOrder0zgensym_6a7ebc1181a5fc8f_1)
		if encodedFieldsLeft0zgensym_6a7ebc1181a5fc8f_1 > 0 {
			encodedFieldsLeft0zgensym_6a7ebc1181a5fc8f_1--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0zgensym_6a7ebc1181a5fc8f_1 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zgensym_6a7ebc1181a5fc8f_1 < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zgensym_6a7ebc1181a5fc8f_1 = 0
			}
			for nextMiss0zgensym_6a7ebc1181a5fc8f_1 < maxFields0zgensym_6a7ebc1181a5fc8f_1 && (found0zgensym_6a7ebc1181a5fc8f_1[nextMiss0zgensym_6a7ebc1181a5fc8f_1] || decodeMsgFieldSkip0zgensym_6a7ebc1181a5fc8f_1[nextMiss0zgensym_6a7ebc1181a5fc8f_1]) {
				nextMiss0zgensym_6a7ebc1181a5fc8f_1++
			}
			if nextMiss0zgensym_6a7ebc1181a5fc8f_1 == maxFields0zgensym_6a7ebc1181a5fc8f_1 {
				// filled all the empty fields!
				break doneWithStruct0zgensym_6a7ebc1181a5fc8f_1
			}
			missingFieldsLeft0zgensym_6a7ebc1181a5fc8f_1--
			curField0zgensym_6a7ebc1181a5fc8f_1 = decodeMsgFieldOrder0zgensym_6a7ebc1181a5fc8f_1[nextMiss0zgensym_6a7ebc1181a5fc8f_1]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zgensym_6a7ebc1181a5fc8f_1)
		switch curField0zgensym_6a7ebc1181a5fc8f_1 {
		// -- templateDecodeMsg ends here --

		case "SignMe_zid00_rct":
			found0zgensym_6a7ebc1181a5fc8f_1[0] = true
			err = z.SignMe.DecodeMsg(dc)
			if err != nil {
				return
			}
		case "SignatureFormat_zid01_str":
			found0zgensym_6a7ebc1181a5fc8f_1[1] = true
			z.SignatureFormat, err = dc.ReadString()
			if err != nil {
				return
			}
		case "SignatureBlob_zid02_bin":
			found0zgensym_6a7ebc1181a5fc8f_1[2] = true
			z.SignatureBlob, err = dc.ReadBytes(z.SignatureBlob)
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	if nextMiss0zgensym_6a7ebc1181a5fc8f_1 != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	if p, ok := interface{}(z).(msgp.PostLoad); ok {
		p.PostLoadHook()
	}

	return
}

// fields of Selfie
var decodeMsgFieldOrder0zgensym_6a7ebc1181a5fc8f_1 = []string{"SignMe_zid00_rct", "SignatureFormat_zid01_str", "SignatureBlob_zid02_bin"}

var decodeMsgFieldSkip0zgensym_6a7ebc1181a5fc8f_1 = []bool{false, false, false}

// fieldsNotEmpty supports omitempty tags
func (z *Selfie) fieldsNotEmpty(isempty []bool) uint32 {
	if len(isempty) == 0 {
		return 3
	}
	var fieldsInUse uint32 = 3
	isempty[0] = false
	if isempty[0] {
		fieldsInUse--
	}
	isempty[1] = (len(z.SignatureFormat) == 0) // string, omitempty
	if isempty[1] {
		fieldsInUse--
	}
	isempty[2] = (len(z.SignatureBlob) == 0) // string, omitempty
	if isempty[2] {
		fieldsInUse--
	}

	return fieldsInUse
}

// EncodeMsg implements msgp.Encodable
func (z *Selfie) EncodeMsg(en *msgp.Writer) (err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	// honor the omitempty tags
	var empty_zgensym_6a7ebc1181a5fc8f_2 [3]bool
	fieldsInUse_zgensym_6a7ebc1181a5fc8f_3 := z.fieldsNotEmpty(empty_zgensym_6a7ebc1181a5fc8f_2[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zgensym_6a7ebc1181a5fc8f_3)
	if err != nil {
		return err
	}

	if !empty_zgensym_6a7ebc1181a5fc8f_2[0] {
		// write "SignMe_zid00_rct"
		err = en.Append(0xb0, 0x53, 0x69, 0x67, 0x6e, 0x4d, 0x65, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x72, 0x63, 0x74)
		if err != nil {
			return err
		}
		err = z.SignMe.EncodeMsg(en)
		if err != nil {
			return
		}
	}

	if !empty_zgensym_6a7ebc1181a5fc8f_2[1] {
		// write "SignatureFormat_zid01_str"
		err = en.Append(0xb9, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x31, 0x5f, 0x73, 0x74, 0x72)
		if err != nil {
			return err
		}
		err = en.WriteString(z.SignatureFormat)
		if err != nil {
			return
		}
	}

	if !empty_zgensym_6a7ebc1181a5fc8f_2[2] {
		// write "SignatureBlob_zid02_bin"
		err = en.Append(0xb7, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x42, 0x6c, 0x6f, 0x62, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x32, 0x5f, 0x62, 0x69, 0x6e)
		if err != nil {
			return err
		}
		err = en.WriteBytes(z.SignatureBlob)
		if err != nil {
			return
		}
	}

	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Selfie) MarshalMsg(b []byte) (o []byte, err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	o = msgp.Require(b, z.Msgsize())

	// honor the omitempty tags
	var empty [3]bool
	fieldsInUse := z.fieldsNotEmpty(empty[:])
	o = msgp.AppendMapHeader(o, fieldsInUse)

	if !empty[0] {
		// string "SignMe_zid00_rct"
		o = append(o, 0xb0, 0x53, 0x69, 0x67, 0x6e, 0x4d, 0x65, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x72, 0x63, 0x74)
		o, err = z.SignMe.MarshalMsg(o)
		if err != nil {
			return
		}
	}

	if !empty[1] {
		// string "SignatureFormat_zid01_str"
		o = append(o, 0xb9, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x31, 0x5f, 0x73, 0x74, 0x72)
		o = msgp.AppendString(o, z.SignatureFormat)
	}

	if !empty[2] {
		// string "SignatureBlob_zid02_bin"
		o = append(o, 0xb7, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x42, 0x6c, 0x6f, 0x62, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x32, 0x5f, 0x62, 0x69, 0x6e)
		o = msgp.AppendBytes(o, z.SignatureBlob)
	}

	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Selfie) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *Selfie) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields4zgensym_6a7ebc1181a5fc8f_5 = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields4zgensym_6a7ebc1181a5fc8f_5 uint32
	if !nbs.AlwaysNil {
		totalEncodedFields4zgensym_6a7ebc1181a5fc8f_5, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft4zgensym_6a7ebc1181a5fc8f_5 := totalEncodedFields4zgensym_6a7ebc1181a5fc8f_5
	missingFieldsLeft4zgensym_6a7ebc1181a5fc8f_5 := maxFields4zgensym_6a7ebc1181a5fc8f_5 - totalEncodedFields4zgensym_6a7ebc1181a5fc8f_5

	var nextMiss4zgensym_6a7ebc1181a5fc8f_5 int32 = -1
	var found4zgensym_6a7ebc1181a5fc8f_5 [maxFields4zgensym_6a7ebc1181a5fc8f_5]bool
	var curField4zgensym_6a7ebc1181a5fc8f_5 string

doneWithStruct4zgensym_6a7ebc1181a5fc8f_5:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zgensym_6a7ebc1181a5fc8f_5 > 0 || missingFieldsLeft4zgensym_6a7ebc1181a5fc8f_5 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zgensym_6a7ebc1181a5fc8f_5, missingFieldsLeft4zgensym_6a7ebc1181a5fc8f_5, msgp.ShowFound(found4zgensym_6a7ebc1181a5fc8f_5[:]), unmarshalMsgFieldOrder4zgensym_6a7ebc1181a5fc8f_5)
		if encodedFieldsLeft4zgensym_6a7ebc1181a5fc8f_5 > 0 {
			encodedFieldsLeft4zgensym_6a7ebc1181a5fc8f_5--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField4zgensym_6a7ebc1181a5fc8f_5 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zgensym_6a7ebc1181a5fc8f_5 < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss4zgensym_6a7ebc1181a5fc8f_5 = 0
			}
			for nextMiss4zgensym_6a7ebc1181a5fc8f_5 < maxFields4zgensym_6a7ebc1181a5fc8f_5 && (found4zgensym_6a7ebc1181a5fc8f_5[nextMiss4zgensym_6a7ebc1181a5fc8f_5] || unmarshalMsgFieldSkip4zgensym_6a7ebc1181a5fc8f_5[nextMiss4zgensym_6a7ebc1181a5fc8f_5]) {
				nextMiss4zgensym_6a7ebc1181a5fc8f_5++
			}
			if nextMiss4zgensym_6a7ebc1181a5fc8f_5 == maxFields4zgensym_6a7ebc1181a5fc8f_5 {
				// filled all the empty fields!
				break doneWithStruct4zgensym_6a7ebc1181a5fc8f_5
			}
			missingFieldsLeft4zgensym_6a7ebc1181a5fc8f_5--
			curField4zgensym_6a7ebc1181a5fc8f_5 = unmarshalMsgFieldOrder4zgensym_6a7ebc1181a5fc8f_5[nextMiss4zgensym_6a7ebc1181a5fc8f_5]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zgensym_6a7ebc1181a5fc8f_5)
		switch curField4zgensym_6a7ebc1181a5fc8f_5 {
		// -- templateUnmarshalMsg ends here --

		case "SignMe_zid00_rct":
			found4zgensym_6a7ebc1181a5fc8f_5[0] = true
			bts, err = z.SignMe.UnmarshalMsg(bts)
			if err != nil {
				return
			}
			if err != nil {
				return
			}
		case "SignatureFormat_zid01_str":
			found4zgensym_6a7ebc1181a5fc8f_5[1] = true
			z.SignatureFormat, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "SignatureBlob_zid02_bin":
			found4zgensym_6a7ebc1181a5fc8f_5[2] = true
			if nbs.AlwaysNil || msgp.IsNil(bts) {
				if !nbs.AlwaysNil {
					bts = bts[1:]
				}
				z.SignatureBlob = z.SignatureBlob[:0]
			} else {
				z.SignatureBlob, bts, err = nbs.ReadBytesBytes(bts, z.SignatureBlob)

				if err != nil {
					return
				}
			}
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	if nextMiss4zgensym_6a7ebc1181a5fc8f_5 != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	if p, ok := interface{}(z).(msgp.PostLoad); ok {
		p.PostLoadHook()
	}

	return
}

// fields of Selfie
var unmarshalMsgFieldOrder4zgensym_6a7ebc1181a5fc8f_5 = []string{"SignMe_zid00_rct", "SignatureFormat_zid01_str", "SignatureBlob_zid02_bin"}

var unmarshalMsgFieldSkip4zgensym_6a7ebc1181a5fc8f_5 = []bool{false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Selfie) Msgsize() (s int) {
	s = 1 + 17 + z.SignMe.Msgsize() + 26 + msgp.StringPrefixSize + len(z.SignatureFormat) + 24 + msgp.BytesPrefixSize + len(z.SignatureBlob)
	return
}

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *SignedStuff) DecodeMsg(dc *msgp.Reader) (err error) {
	var sawTopNil bool
	if dc.IsNil() {
		sawTopNil = true
		err = dc.ReadNil()
		if err != nil {
			return
		}
		dc.PushAlwaysNil()
	}

	var field []byte
	_ = field
	const maxFields7zgensym_6a7ebc1181a5fc8f_8 = 3

	// -- templateDecodeMsg starts here--
	var totalEncodedFields7zgensym_6a7ebc1181a5fc8f_8 uint32
	totalEncodedFields7zgensym_6a7ebc1181a5fc8f_8, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft7zgensym_6a7ebc1181a5fc8f_8 := totalEncodedFields7zgensym_6a7ebc1181a5fc8f_8
	missingFieldsLeft7zgensym_6a7ebc1181a5fc8f_8 := maxFields7zgensym_6a7ebc1181a5fc8f_8 - totalEncodedFields7zgensym_6a7ebc1181a5fc8f_8

	var nextMiss7zgensym_6a7ebc1181a5fc8f_8 int32 = -1
	var found7zgensym_6a7ebc1181a5fc8f_8 [maxFields7zgensym_6a7ebc1181a5fc8f_8]bool
	var curField7zgensym_6a7ebc1181a5fc8f_8 string

doneWithStruct7zgensym_6a7ebc1181a5fc8f_8:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft7zgensym_6a7ebc1181a5fc8f_8 > 0 || missingFieldsLeft7zgensym_6a7ebc1181a5fc8f_8 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft7zgensym_6a7ebc1181a5fc8f_8, missingFieldsLeft7zgensym_6a7ebc1181a5fc8f_8, msgp.ShowFound(found7zgensym_6a7ebc1181a5fc8f_8[:]), decodeMsgFieldOrder7zgensym_6a7ebc1181a5fc8f_8)
		if encodedFieldsLeft7zgensym_6a7ebc1181a5fc8f_8 > 0 {
			encodedFieldsLeft7zgensym_6a7ebc1181a5fc8f_8--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField7zgensym_6a7ebc1181a5fc8f_8 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss7zgensym_6a7ebc1181a5fc8f_8 < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss7zgensym_6a7ebc1181a5fc8f_8 = 0
			}
			for nextMiss7zgensym_6a7ebc1181a5fc8f_8 < maxFields7zgensym_6a7ebc1181a5fc8f_8 && (found7zgensym_6a7ebc1181a5fc8f_8[nextMiss7zgensym_6a7ebc1181a5fc8f_8] || decodeMsgFieldSkip7zgensym_6a7ebc1181a5fc8f_8[nextMiss7zgensym_6a7ebc1181a5fc8f_8]) {
				nextMiss7zgensym_6a7ebc1181a5fc8f_8++
			}
			if nextMiss7zgensym_6a7ebc1181a5fc8f_8 == maxFields7zgensym_6a7ebc1181a5fc8f_8 {
				// filled all the empty fields!
				break doneWithStruct7zgensym_6a7ebc1181a5fc8f_8
			}
			missingFieldsLeft7zgensym_6a7ebc1181a5fc8f_8--
			curField7zgensym_6a7ebc1181a5fc8f_8 = decodeMsgFieldOrder7zgensym_6a7ebc1181a5fc8f_8[nextMiss7zgensym_6a7ebc1181a5fc8f_8]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField7zgensym_6a7ebc1181a5fc8f_8)
		switch curField7zgensym_6a7ebc1181a5fc8f_8 {
		// -- templateDecodeMsg ends here --

		case "PubKeyBytes_zid00_bin":
			found7zgensym_6a7ebc1181a5fc8f_8[0] = true
			z.PubKeyBytes, err = dc.ReadBytes(z.PubKeyBytes)
			if err != nil {
				return
			}
		case "SignedAtTimestamp_zid01_tim":
			found7zgensym_6a7ebc1181a5fc8f_8[1] = true
			z.SignedAtTimestamp, err = dc.ReadTime()
			if err != nil {
				return
			}
		case "Others_zid02_slc":
			found7zgensym_6a7ebc1181a5fc8f_8[2] = true
			var zgensym_6a7ebc1181a5fc8f_9 uint32
			zgensym_6a7ebc1181a5fc8f_9, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Others) >= int(zgensym_6a7ebc1181a5fc8f_9) {
				z.Others = (z.Others)[:zgensym_6a7ebc1181a5fc8f_9]
			} else {
				z.Others = make([]*Selfie, zgensym_6a7ebc1181a5fc8f_9)
			}
			for zgensym_6a7ebc1181a5fc8f_6 := range z.Others {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					z.Others[zgensym_6a7ebc1181a5fc8f_6] = nil
				} else {
					if z.Others[zgensym_6a7ebc1181a5fc8f_6] == nil {
						z.Others[zgensym_6a7ebc1181a5fc8f_6] = new(Selfie)
					}
					const maxFields10zgensym_6a7ebc1181a5fc8f_11 = 3

					// -- templateDecodeMsg starts here--
					var totalEncodedFields10zgensym_6a7ebc1181a5fc8f_11 uint32
					totalEncodedFields10zgensym_6a7ebc1181a5fc8f_11, err = dc.ReadMapHeader()
					if err != nil {
						return
					}
					encodedFieldsLeft10zgensym_6a7ebc1181a5fc8f_11 := totalEncodedFields10zgensym_6a7ebc1181a5fc8f_11
					missingFieldsLeft10zgensym_6a7ebc1181a5fc8f_11 := maxFields10zgensym_6a7ebc1181a5fc8f_11 - totalEncodedFields10zgensym_6a7ebc1181a5fc8f_11

					var nextMiss10zgensym_6a7ebc1181a5fc8f_11 int32 = -1
					var found10zgensym_6a7ebc1181a5fc8f_11 [maxFields10zgensym_6a7ebc1181a5fc8f_11]bool
					var curField10zgensym_6a7ebc1181a5fc8f_11 string

				doneWithStruct10zgensym_6a7ebc1181a5fc8f_11:
					// First fill all the encoded fields, then
					// treat the remaining, missing fields, as Nil.
					for encodedFieldsLeft10zgensym_6a7ebc1181a5fc8f_11 > 0 || missingFieldsLeft10zgensym_6a7ebc1181a5fc8f_11 > 0 {
						//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft10zgensym_6a7ebc1181a5fc8f_11, missingFieldsLeft10zgensym_6a7ebc1181a5fc8f_11, msgp.ShowFound(found10zgensym_6a7ebc1181a5fc8f_11[:]), decodeMsgFieldOrder10zgensym_6a7ebc1181a5fc8f_11)
						if encodedFieldsLeft10zgensym_6a7ebc1181a5fc8f_11 > 0 {
							encodedFieldsLeft10zgensym_6a7ebc1181a5fc8f_11--
							field, err = dc.ReadMapKeyPtr()
							if err != nil {
								return
							}
							curField10zgensym_6a7ebc1181a5fc8f_11 = msgp.UnsafeString(field)
						} else {
							//missing fields need handling
							if nextMiss10zgensym_6a7ebc1181a5fc8f_11 < 0 {
								// tell the reader to only give us Nils
								// until further notice.
								dc.PushAlwaysNil()
								nextMiss10zgensym_6a7ebc1181a5fc8f_11 = 0
							}
							for nextMiss10zgensym_6a7ebc1181a5fc8f_11 < maxFields10zgensym_6a7ebc1181a5fc8f_11 && (found10zgensym_6a7ebc1181a5fc8f_11[nextMiss10zgensym_6a7ebc1181a5fc8f_11] || decodeMsgFieldSkip10zgensym_6a7ebc1181a5fc8f_11[nextMiss10zgensym_6a7ebc1181a5fc8f_11]) {
								nextMiss10zgensym_6a7ebc1181a5fc8f_11++
							}
							if nextMiss10zgensym_6a7ebc1181a5fc8f_11 == maxFields10zgensym_6a7ebc1181a5fc8f_11 {
								// filled all the empty fields!
								break doneWithStruct10zgensym_6a7ebc1181a5fc8f_11
							}
							missingFieldsLeft10zgensym_6a7ebc1181a5fc8f_11--
							curField10zgensym_6a7ebc1181a5fc8f_11 = decodeMsgFieldOrder10zgensym_6a7ebc1181a5fc8f_11[nextMiss10zgensym_6a7ebc1181a5fc8f_11]
						}
						//fmt.Printf("switching on curField: '%v'\n", curField10zgensym_6a7ebc1181a5fc8f_11)
						switch curField10zgensym_6a7ebc1181a5fc8f_11 {
						// -- templateDecodeMsg ends here --

						case "SignMe_zid00_rct":
							found10zgensym_6a7ebc1181a5fc8f_11[0] = true
							err = z.Others[zgensym_6a7ebc1181a5fc8f_6].SignMe.DecodeMsg(dc)
							if err != nil {
								return
							}
						case "SignatureFormat_zid01_str":
							found10zgensym_6a7ebc1181a5fc8f_11[1] = true
							z.Others[zgensym_6a7ebc1181a5fc8f_6].SignatureFormat, err = dc.ReadString()
							if err != nil {
								return
							}
						case "SignatureBlob_zid02_bin":
							found10zgensym_6a7ebc1181a5fc8f_11[2] = true
							z.Others[zgensym_6a7ebc1181a5fc8f_6].SignatureBlob, err = dc.ReadBytes(z.Others[zgensym_6a7ebc1181a5fc8f_6].SignatureBlob)
							if err != nil {
								return
							}
						default:
							err = dc.Skip()
							if err != nil {
								return
							}
						}
					}
					if nextMiss10zgensym_6a7ebc1181a5fc8f_11 != -1 {
						dc.PopAlwaysNil()
					}

				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	if nextMiss7zgensym_6a7ebc1181a5fc8f_8 != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	if p, ok := interface{}(z).(msgp.PostLoad); ok {
		p.PostLoadHook()
	}

	return
}

// fields of SignedStuff
var decodeMsgFieldOrder7zgensym_6a7ebc1181a5fc8f_8 = []string{"PubKeyBytes_zid00_bin", "SignedAtTimestamp_zid01_tim", "Others_zid02_slc"}

var decodeMsgFieldSkip7zgensym_6a7ebc1181a5fc8f_8 = []bool{false, false, false}

// fields of Selfie
var decodeMsgFieldOrder10zgensym_6a7ebc1181a5fc8f_11 = []string{"SignMe_zid00_rct", "SignatureFormat_zid01_str", "SignatureBlob_zid02_bin"}

var decodeMsgFieldSkip10zgensym_6a7ebc1181a5fc8f_11 = []bool{false, false, false}

// fieldsNotEmpty supports omitempty tags
func (z *SignedStuff) fieldsNotEmpty(isempty []bool) uint32 {
	if len(isempty) == 0 {
		return 3
	}
	var fieldsInUse uint32 = 3
	isempty[0] = (len(z.PubKeyBytes) == 0) // string, omitempty
	if isempty[0] {
		fieldsInUse--
	}
	isempty[1] = (z.SignedAtTimestamp.IsZero()) // time.Time, omitempty
	if isempty[1] {
		fieldsInUse--
	}
	isempty[2] = (len(z.Others) == 0) // string, omitempty
	if isempty[2] {
		fieldsInUse--
	}

	return fieldsInUse
}

// EncodeMsg implements msgp.Encodable
func (z *SignedStuff) EncodeMsg(en *msgp.Writer) (err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	// honor the omitempty tags
	var empty_zgensym_6a7ebc1181a5fc8f_12 [3]bool
	fieldsInUse_zgensym_6a7ebc1181a5fc8f_13 := z.fieldsNotEmpty(empty_zgensym_6a7ebc1181a5fc8f_12[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zgensym_6a7ebc1181a5fc8f_13)
	if err != nil {
		return err
	}

	if !empty_zgensym_6a7ebc1181a5fc8f_12[0] {
		// write "PubKeyBytes_zid00_bin"
		err = en.Append(0xb5, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x42, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x62, 0x69, 0x6e)
		if err != nil {
			return err
		}
		err = en.WriteBytes(z.PubKeyBytes)
		if err != nil {
			return
		}
	}

	if !empty_zgensym_6a7ebc1181a5fc8f_12[1] {
		// write "SignedAtTimestamp_zid01_tim"
		err = en.Append(0xbb, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x41, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x31, 0x5f, 0x74, 0x69, 0x6d)
		if err != nil {
			return err
		}
		err = en.WriteTime(z.SignedAtTimestamp)
		if err != nil {
			return
		}
	}

	if !empty_zgensym_6a7ebc1181a5fc8f_12[2] {
		// write "Others_zid02_slc"
		err = en.Append(0xb0, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x73, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x32, 0x5f, 0x73, 0x6c, 0x63)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(z.Others)))
		if err != nil {
			return
		}
		for zgensym_6a7ebc1181a5fc8f_6 := range z.Others {
			if z.Others[zgensym_6a7ebc1181a5fc8f_6] == nil {
				err = en.WriteNil()
				if err != nil {
					return
				}
			} else {

				// honor the omitempty tags
				var empty_zgensym_6a7ebc1181a5fc8f_14 [3]bool
				fieldsInUse_zgensym_6a7ebc1181a5fc8f_15 := z.Others[zgensym_6a7ebc1181a5fc8f_6].fieldsNotEmpty(empty_zgensym_6a7ebc1181a5fc8f_14[:])

				// map header
				err = en.WriteMapHeader(fieldsInUse_zgensym_6a7ebc1181a5fc8f_15)
				if err != nil {
					return err
				}

				if !empty_zgensym_6a7ebc1181a5fc8f_14[0] {
					// write "SignMe_zid00_rct"
					err = en.Append(0xb0, 0x53, 0x69, 0x67, 0x6e, 0x4d, 0x65, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x72, 0x63, 0x74)
					if err != nil {
						return err
					}
					err = z.Others[zgensym_6a7ebc1181a5fc8f_6].SignMe.EncodeMsg(en)
					if err != nil {
						return
					}
				}

				if !empty_zgensym_6a7ebc1181a5fc8f_14[1] {
					// write "SignatureFormat_zid01_str"
					err = en.Append(0xb9, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x31, 0x5f, 0x73, 0x74, 0x72)
					if err != nil {
						return err
					}
					err = en.WriteString(z.Others[zgensym_6a7ebc1181a5fc8f_6].SignatureFormat)
					if err != nil {
						return
					}
				}

				if !empty_zgensym_6a7ebc1181a5fc8f_14[2] {
					// write "SignatureBlob_zid02_bin"
					err = en.Append(0xb7, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x42, 0x6c, 0x6f, 0x62, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x32, 0x5f, 0x62, 0x69, 0x6e)
					if err != nil {
						return err
					}
					err = en.WriteBytes(z.Others[zgensym_6a7ebc1181a5fc8f_6].SignatureBlob)
					if err != nil {
						return
					}
				}

			}
		}
	}

	return
}

// MarshalMsg implements msgp.Marshaler
func (z *SignedStuff) MarshalMsg(b []byte) (o []byte, err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	o = msgp.Require(b, z.Msgsize())

	// honor the omitempty tags
	var empty [3]bool
	fieldsInUse := z.fieldsNotEmpty(empty[:])
	o = msgp.AppendMapHeader(o, fieldsInUse)

	if !empty[0] {
		// string "PubKeyBytes_zid00_bin"
		o = append(o, 0xb5, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x42, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x62, 0x69, 0x6e)
		o = msgp.AppendBytes(o, z.PubKeyBytes)
	}

	if !empty[1] {
		// string "SignedAtTimestamp_zid01_tim"
		o = append(o, 0xbb, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x41, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x31, 0x5f, 0x74, 0x69, 0x6d)
		o = msgp.AppendTime(o, z.SignedAtTimestamp)
	}

	if !empty[2] {
		// string "Others_zid02_slc"
		o = append(o, 0xb0, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x73, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x32, 0x5f, 0x73, 0x6c, 0x63)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Others)))
		for zgensym_6a7ebc1181a5fc8f_6 := range z.Others {
			if z.Others[zgensym_6a7ebc1181a5fc8f_6] == nil {
				o = msgp.AppendNil(o)
			} else {

				// honor the omitempty tags
				var empty [3]bool
				fieldsInUse := z.Others[zgensym_6a7ebc1181a5fc8f_6].fieldsNotEmpty(empty[:])
				o = msgp.AppendMapHeader(o, fieldsInUse)

				if !empty[0] {
					// string "SignMe_zid00_rct"
					o = append(o, 0xb0, 0x53, 0x69, 0x67, 0x6e, 0x4d, 0x65, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x72, 0x63, 0x74)
					o, err = z.Others[zgensym_6a7ebc1181a5fc8f_6].SignMe.MarshalMsg(o)
					if err != nil {
						return
					}
				}

				if !empty[1] {
					// string "SignatureFormat_zid01_str"
					o = append(o, 0xb9, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x31, 0x5f, 0x73, 0x74, 0x72)
					o = msgp.AppendString(o, z.Others[zgensym_6a7ebc1181a5fc8f_6].SignatureFormat)
				}

				if !empty[2] {
					// string "SignatureBlob_zid02_bin"
					o = append(o, 0xb7, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x42, 0x6c, 0x6f, 0x62, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x32, 0x5f, 0x62, 0x69, 0x6e)
					o = msgp.AppendBytes(o, z.Others[zgensym_6a7ebc1181a5fc8f_6].SignatureBlob)
				}

			}
		}
	}

	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *SignedStuff) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *SignedStuff) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields16zgensym_6a7ebc1181a5fc8f_17 = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields16zgensym_6a7ebc1181a5fc8f_17 uint32
	if !nbs.AlwaysNil {
		totalEncodedFields16zgensym_6a7ebc1181a5fc8f_17, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft16zgensym_6a7ebc1181a5fc8f_17 := totalEncodedFields16zgensym_6a7ebc1181a5fc8f_17
	missingFieldsLeft16zgensym_6a7ebc1181a5fc8f_17 := maxFields16zgensym_6a7ebc1181a5fc8f_17 - totalEncodedFields16zgensym_6a7ebc1181a5fc8f_17

	var nextMiss16zgensym_6a7ebc1181a5fc8f_17 int32 = -1
	var found16zgensym_6a7ebc1181a5fc8f_17 [maxFields16zgensym_6a7ebc1181a5fc8f_17]bool
	var curField16zgensym_6a7ebc1181a5fc8f_17 string

doneWithStruct16zgensym_6a7ebc1181a5fc8f_17:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft16zgensym_6a7ebc1181a5fc8f_17 > 0 || missingFieldsLeft16zgensym_6a7ebc1181a5fc8f_17 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft16zgensym_6a7ebc1181a5fc8f_17, missingFieldsLeft16zgensym_6a7ebc1181a5fc8f_17, msgp.ShowFound(found16zgensym_6a7ebc1181a5fc8f_17[:]), unmarshalMsgFieldOrder16zgensym_6a7ebc1181a5fc8f_17)
		if encodedFieldsLeft16zgensym_6a7ebc1181a5fc8f_17 > 0 {
			encodedFieldsLeft16zgensym_6a7ebc1181a5fc8f_17--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField16zgensym_6a7ebc1181a5fc8f_17 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss16zgensym_6a7ebc1181a5fc8f_17 < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss16zgensym_6a7ebc1181a5fc8f_17 = 0
			}
			for nextMiss16zgensym_6a7ebc1181a5fc8f_17 < maxFields16zgensym_6a7ebc1181a5fc8f_17 && (found16zgensym_6a7ebc1181a5fc8f_17[nextMiss16zgensym_6a7ebc1181a5fc8f_17] || unmarshalMsgFieldSkip16zgensym_6a7ebc1181a5fc8f_17[nextMiss16zgensym_6a7ebc1181a5fc8f_17]) {
				nextMiss16zgensym_6a7ebc1181a5fc8f_17++
			}
			if nextMiss16zgensym_6a7ebc1181a5fc8f_17 == maxFields16zgensym_6a7ebc1181a5fc8f_17 {
				// filled all the empty fields!
				break doneWithStruct16zgensym_6a7ebc1181a5fc8f_17
			}
			missingFieldsLeft16zgensym_6a7ebc1181a5fc8f_17--
			curField16zgensym_6a7ebc1181a5fc8f_17 = unmarshalMsgFieldOrder16zgensym_6a7ebc1181a5fc8f_17[nextMiss16zgensym_6a7ebc1181a5fc8f_17]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField16zgensym_6a7ebc1181a5fc8f_17)
		switch curField16zgensym_6a7ebc1181a5fc8f_17 {
		// -- templateUnmarshalMsg ends here --

		case "PubKeyBytes_zid00_bin":
			found16zgensym_6a7ebc1181a5fc8f_17[0] = true
			if nbs.AlwaysNil || msgp.IsNil(bts) {
				if !nbs.AlwaysNil {
					bts = bts[1:]
				}
				z.PubKeyBytes = z.PubKeyBytes[:0]
			} else {
				z.PubKeyBytes, bts, err = nbs.ReadBytesBytes(bts, z.PubKeyBytes)

				if err != nil {
					return
				}
			}
			if err != nil {
				return
			}
		case "SignedAtTimestamp_zid01_tim":
			found16zgensym_6a7ebc1181a5fc8f_17[1] = true
			z.SignedAtTimestamp, bts, err = nbs.ReadTimeBytes(bts)

			if err != nil {
				return
			}
		case "Others_zid02_slc":
			found16zgensym_6a7ebc1181a5fc8f_17[2] = true
			if nbs.AlwaysNil {
				(z.Others) = (z.Others)[:0]
			} else {

				var zgensym_6a7ebc1181a5fc8f_18 uint32
				zgensym_6a7ebc1181a5fc8f_18, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.Others) >= int(zgensym_6a7ebc1181a5fc8f_18) {
					z.Others = (z.Others)[:zgensym_6a7ebc1181a5fc8f_18]
				} else {
					z.Others = make([]*Selfie, zgensym_6a7ebc1181a5fc8f_18)
				}
				for zgensym_6a7ebc1181a5fc8f_6 := range z.Others {
					// default gPtr logic.
					if nbs.PeekNil(bts) && z.Others[zgensym_6a7ebc1181a5fc8f_6] == nil {
						// consume the nil
						bts, err = nbs.ReadNilBytes(bts)
						if err != nil {
							return
						}
					} else {
						// read as-if the wire has bytes, letting nbs take care of nils.

						if z.Others[zgensym_6a7ebc1181a5fc8f_6] == nil {
							z.Others[zgensym_6a7ebc1181a5fc8f_6] = new(Selfie)
						}
						const maxFields19zgensym_6a7ebc1181a5fc8f_20 = 3

						// -- templateUnmarshalMsg starts here--
						var totalEncodedFields19zgensym_6a7ebc1181a5fc8f_20 uint32
						if !nbs.AlwaysNil {
							totalEncodedFields19zgensym_6a7ebc1181a5fc8f_20, bts, err = nbs.ReadMapHeaderBytes(bts)
							if err != nil {
								return
							}
						}
						encodedFieldsLeft19zgensym_6a7ebc1181a5fc8f_20 := totalEncodedFields19zgensym_6a7ebc1181a5fc8f_20
						missingFieldsLeft19zgensym_6a7ebc1181a5fc8f_20 := maxFields19zgensym_6a7ebc1181a5fc8f_20 - totalEncodedFields19zgensym_6a7ebc1181a5fc8f_20

						var nextMiss19zgensym_6a7ebc1181a5fc8f_20 int32 = -1
						var found19zgensym_6a7ebc1181a5fc8f_20 [maxFields19zgensym_6a7ebc1181a5fc8f_20]bool
						var curField19zgensym_6a7ebc1181a5fc8f_20 string

					doneWithStruct19zgensym_6a7ebc1181a5fc8f_20:
						// First fill all the encoded fields, then
						// treat the remaining, missing fields, as Nil.
						for encodedFieldsLeft19zgensym_6a7ebc1181a5fc8f_20 > 0 || missingFieldsLeft19zgensym_6a7ebc1181a5fc8f_20 > 0 {
							//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft19zgensym_6a7ebc1181a5fc8f_20, missingFieldsLeft19zgensym_6a7ebc1181a5fc8f_20, msgp.ShowFound(found19zgensym_6a7ebc1181a5fc8f_20[:]), unmarshalMsgFieldOrder19zgensym_6a7ebc1181a5fc8f_20)
							if encodedFieldsLeft19zgensym_6a7ebc1181a5fc8f_20 > 0 {
								encodedFieldsLeft19zgensym_6a7ebc1181a5fc8f_20--
								field, bts, err = nbs.ReadMapKeyZC(bts)
								if err != nil {
									return
								}
								curField19zgensym_6a7ebc1181a5fc8f_20 = msgp.UnsafeString(field)
							} else {
								//missing fields need handling
								if nextMiss19zgensym_6a7ebc1181a5fc8f_20 < 0 {
									// set bts to contain just mnil (0xc0)
									bts = nbs.PushAlwaysNil(bts)
									nextMiss19zgensym_6a7ebc1181a5fc8f_20 = 0
								}
								for nextMiss19zgensym_6a7ebc1181a5fc8f_20 < maxFields19zgensym_6a7ebc1181a5fc8f_20 && (found19zgensym_6a7ebc1181a5fc8f_20[nextMiss19zgensym_6a7ebc1181a5fc8f_20] || unmarshalMsgFieldSkip19zgensym_6a7ebc1181a5fc8f_20[nextMiss19zgensym_6a7ebc1181a5fc8f_20]) {
									nextMiss19zgensym_6a7ebc1181a5fc8f_20++
								}
								if nextMiss19zgensym_6a7ebc1181a5fc8f_20 == maxFields19zgensym_6a7ebc1181a5fc8f_20 {
									// filled all the empty fields!
									break doneWithStruct19zgensym_6a7ebc1181a5fc8f_20
								}
								missingFieldsLeft19zgensym_6a7ebc1181a5fc8f_20--
								curField19zgensym_6a7ebc1181a5fc8f_20 = unmarshalMsgFieldOrder19zgensym_6a7ebc1181a5fc8f_20[nextMiss19zgensym_6a7ebc1181a5fc8f_20]
							}
							//fmt.Printf("switching on curField: '%v'\n", curField19zgensym_6a7ebc1181a5fc8f_20)
							switch curField19zgensym_6a7ebc1181a5fc8f_20 {
							// -- templateUnmarshalMsg ends here --

							case "SignMe_zid00_rct":
								found19zgensym_6a7ebc1181a5fc8f_20[0] = true
								bts, err = z.Others[zgensym_6a7ebc1181a5fc8f_6].SignMe.UnmarshalMsg(bts)
								if err != nil {
									return
								}
								if err != nil {
									return
								}
							case "SignatureFormat_zid01_str":
								found19zgensym_6a7ebc1181a5fc8f_20[1] = true
								z.Others[zgensym_6a7ebc1181a5fc8f_6].SignatureFormat, bts, err = nbs.ReadStringBytes(bts)

								if err != nil {
									return
								}
							case "SignatureBlob_zid02_bin":
								found19zgensym_6a7ebc1181a5fc8f_20[2] = true
								if nbs.AlwaysNil || msgp.IsNil(bts) {
									if !nbs.AlwaysNil {
										bts = bts[1:]
									}
									z.Others[zgensym_6a7ebc1181a5fc8f_6].SignatureBlob = z.Others[zgensym_6a7ebc1181a5fc8f_6].SignatureBlob[:0]
								} else {
									z.Others[zgensym_6a7ebc1181a5fc8f_6].SignatureBlob, bts, err = nbs.ReadBytesBytes(bts, z.Others[zgensym_6a7ebc1181a5fc8f_6].SignatureBlob)

									if err != nil {
										return
									}
								}
								if err != nil {
									return
								}
							default:
								bts, err = msgp.Skip(bts)
								if err != nil {
									return
								}
							}
						}
						if nextMiss19zgensym_6a7ebc1181a5fc8f_20 != -1 {
							bts = nbs.PopAlwaysNil()
						}

					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	if nextMiss16zgensym_6a7ebc1181a5fc8f_17 != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	if p, ok := interface{}(z).(msgp.PostLoad); ok {
		p.PostLoadHook()
	}

	return
}

// fields of SignedStuff
var unmarshalMsgFieldOrder16zgensym_6a7ebc1181a5fc8f_17 = []string{"PubKeyBytes_zid00_bin", "SignedAtTimestamp_zid01_tim", "Others_zid02_slc"}

var unmarshalMsgFieldSkip16zgensym_6a7ebc1181a5fc8f_17 = []bool{false, false, false}

// fields of Selfie
var unmarshalMsgFieldOrder19zgensym_6a7ebc1181a5fc8f_20 = []string{"SignMe_zid00_rct", "SignatureFormat_zid01_str", "SignatureBlob_zid02_bin"}

var unmarshalMsgFieldSkip19zgensym_6a7ebc1181a5fc8f_20 = []bool{false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *SignedStuff) Msgsize() (s int) {
	s = 1 + 22 + msgp.BytesPrefixSize + len(z.PubKeyBytes) + 28 + msgp.TimeSize + 17 + msgp.ArrayHeaderSize
	for zgensym_6a7ebc1181a5fc8f_6 := range z.Others {
		if z.Others[zgensym_6a7ebc1181a5fc8f_6] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 17 + z.Others[zgensym_6a7ebc1181a5fc8f_6].SignMe.Msgsize() + 26 + msgp.StringPrefixSize + len(z.Others[zgensym_6a7ebc1181a5fc8f_6].SignatureFormat) + 24 + msgp.BytesPrefixSize + len(z.Others[zgensym_6a7ebc1181a5fc8f_6].SignatureBlob)
		}
	}
	return
}
