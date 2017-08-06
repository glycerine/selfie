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
			const maxFields2zgensym_6a7ebc1181a5fc8f_3 = 2

			// -- templateDecodeMsg starts here--
			var totalEncodedFields2zgensym_6a7ebc1181a5fc8f_3 uint32
			totalEncodedFields2zgensym_6a7ebc1181a5fc8f_3, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			encodedFieldsLeft2zgensym_6a7ebc1181a5fc8f_3 := totalEncodedFields2zgensym_6a7ebc1181a5fc8f_3
			missingFieldsLeft2zgensym_6a7ebc1181a5fc8f_3 := maxFields2zgensym_6a7ebc1181a5fc8f_3 - totalEncodedFields2zgensym_6a7ebc1181a5fc8f_3

			var nextMiss2zgensym_6a7ebc1181a5fc8f_3 int32 = -1
			var found2zgensym_6a7ebc1181a5fc8f_3 [maxFields2zgensym_6a7ebc1181a5fc8f_3]bool
			var curField2zgensym_6a7ebc1181a5fc8f_3 string

		doneWithStruct2zgensym_6a7ebc1181a5fc8f_3:
			// First fill all the encoded fields, then
			// treat the remaining, missing fields, as Nil.
			for encodedFieldsLeft2zgensym_6a7ebc1181a5fc8f_3 > 0 || missingFieldsLeft2zgensym_6a7ebc1181a5fc8f_3 > 0 {
				//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zgensym_6a7ebc1181a5fc8f_3, missingFieldsLeft2zgensym_6a7ebc1181a5fc8f_3, msgp.ShowFound(found2zgensym_6a7ebc1181a5fc8f_3[:]), decodeMsgFieldOrder2zgensym_6a7ebc1181a5fc8f_3)
				if encodedFieldsLeft2zgensym_6a7ebc1181a5fc8f_3 > 0 {
					encodedFieldsLeft2zgensym_6a7ebc1181a5fc8f_3--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						return
					}
					curField2zgensym_6a7ebc1181a5fc8f_3 = msgp.UnsafeString(field)
				} else {
					//missing fields need handling
					if nextMiss2zgensym_6a7ebc1181a5fc8f_3 < 0 {
						// tell the reader to only give us Nils
						// until further notice.
						dc.PushAlwaysNil()
						nextMiss2zgensym_6a7ebc1181a5fc8f_3 = 0
					}
					for nextMiss2zgensym_6a7ebc1181a5fc8f_3 < maxFields2zgensym_6a7ebc1181a5fc8f_3 && (found2zgensym_6a7ebc1181a5fc8f_3[nextMiss2zgensym_6a7ebc1181a5fc8f_3] || decodeMsgFieldSkip2zgensym_6a7ebc1181a5fc8f_3[nextMiss2zgensym_6a7ebc1181a5fc8f_3]) {
						nextMiss2zgensym_6a7ebc1181a5fc8f_3++
					}
					if nextMiss2zgensym_6a7ebc1181a5fc8f_3 == maxFields2zgensym_6a7ebc1181a5fc8f_3 {
						// filled all the empty fields!
						break doneWithStruct2zgensym_6a7ebc1181a5fc8f_3
					}
					missingFieldsLeft2zgensym_6a7ebc1181a5fc8f_3--
					curField2zgensym_6a7ebc1181a5fc8f_3 = decodeMsgFieldOrder2zgensym_6a7ebc1181a5fc8f_3[nextMiss2zgensym_6a7ebc1181a5fc8f_3]
				}
				//fmt.Printf("switching on curField: '%v'\n", curField2zgensym_6a7ebc1181a5fc8f_3)
				switch curField2zgensym_6a7ebc1181a5fc8f_3 {
				// -- templateDecodeMsg ends here --

				case "PubKeyBytes_zid00_bin":
					found2zgensym_6a7ebc1181a5fc8f_3[0] = true
					z.SignMe.PubKeyBytes, err = dc.ReadBytes(z.SignMe.PubKeyBytes)
					if err != nil {
						return
					}
				case "SignedAtTimestamp_zid01_tim":
					found2zgensym_6a7ebc1181a5fc8f_3[1] = true
					z.SignMe.SignedAtTimestamp, err = dc.ReadTime()
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
			if nextMiss2zgensym_6a7ebc1181a5fc8f_3 != -1 {
				dc.PopAlwaysNil()
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

// fields of SignedStuff
var decodeMsgFieldOrder2zgensym_6a7ebc1181a5fc8f_3 = []string{"PubKeyBytes_zid00_bin", "SignedAtTimestamp_zid01_tim"}

var decodeMsgFieldSkip2zgensym_6a7ebc1181a5fc8f_3 = []bool{false, false}

// fieldsNotEmpty supports omitempty tags
func (z *Selfie) fieldsNotEmpty(isempty []bool) uint32 {
	if len(isempty) == 0 {
		return 3
	}
	var fieldsInUse uint32 = 3
	isempty[0] = false // struct values are never empty
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
	var empty_zgensym_6a7ebc1181a5fc8f_4 [3]bool
	fieldsInUse_zgensym_6a7ebc1181a5fc8f_5 := z.fieldsNotEmpty(empty_zgensym_6a7ebc1181a5fc8f_4[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zgensym_6a7ebc1181a5fc8f_5)
	if err != nil {
		return err
	}

	if !empty_zgensym_6a7ebc1181a5fc8f_4[0] {
		// write "SignMe_zid00_rct"
		err = en.Append(0xb0, 0x53, 0x69, 0x67, 0x6e, 0x4d, 0x65, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x72, 0x63, 0x74)
		if err != nil {
			return err
		}

		// honor the omitempty tags
		var empty_zgensym_6a7ebc1181a5fc8f_6 [2]bool
		fieldsInUse_zgensym_6a7ebc1181a5fc8f_7 := z.SignMe.fieldsNotEmpty(empty_zgensym_6a7ebc1181a5fc8f_6[:])

		// map header
		err = en.WriteMapHeader(fieldsInUse_zgensym_6a7ebc1181a5fc8f_7)
		if err != nil {
			return err
		}

		if !empty_zgensym_6a7ebc1181a5fc8f_6[0] {
			// write "PubKeyBytes_zid00_bin"
			err = en.Append(0xb5, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x42, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x62, 0x69, 0x6e)
			if err != nil {
				return err
			}
			err = en.WriteBytes(z.SignMe.PubKeyBytes)
			if err != nil {
				return
			}
		}

		if !empty_zgensym_6a7ebc1181a5fc8f_6[1] {
			// write "SignedAtTimestamp_zid01_tim"
			err = en.Append(0xbb, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x41, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x31, 0x5f, 0x74, 0x69, 0x6d)
			if err != nil {
				return err
			}
			err = en.WriteTime(z.SignMe.SignedAtTimestamp)
			if err != nil {
				return
			}
		}

	}

	if !empty_zgensym_6a7ebc1181a5fc8f_4[1] {
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

	if !empty_zgensym_6a7ebc1181a5fc8f_4[2] {
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

		// honor the omitempty tags
		var empty [2]bool
		fieldsInUse := z.SignMe.fieldsNotEmpty(empty[:])
		o = msgp.AppendMapHeader(o, fieldsInUse)

		if !empty[0] {
			// string "PubKeyBytes_zid00_bin"
			o = append(o, 0xb5, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x42, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x62, 0x69, 0x6e)
			o = msgp.AppendBytes(o, z.SignMe.PubKeyBytes)
		}

		if !empty[1] {
			// string "SignedAtTimestamp_zid01_tim"
			o = append(o, 0xbb, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x41, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x31, 0x5f, 0x74, 0x69, 0x6d)
			o = msgp.AppendTime(o, z.SignMe.SignedAtTimestamp)
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
	const maxFields8zgensym_6a7ebc1181a5fc8f_9 = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields8zgensym_6a7ebc1181a5fc8f_9 uint32
	if !nbs.AlwaysNil {
		totalEncodedFields8zgensym_6a7ebc1181a5fc8f_9, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft8zgensym_6a7ebc1181a5fc8f_9 := totalEncodedFields8zgensym_6a7ebc1181a5fc8f_9
	missingFieldsLeft8zgensym_6a7ebc1181a5fc8f_9 := maxFields8zgensym_6a7ebc1181a5fc8f_9 - totalEncodedFields8zgensym_6a7ebc1181a5fc8f_9

	var nextMiss8zgensym_6a7ebc1181a5fc8f_9 int32 = -1
	var found8zgensym_6a7ebc1181a5fc8f_9 [maxFields8zgensym_6a7ebc1181a5fc8f_9]bool
	var curField8zgensym_6a7ebc1181a5fc8f_9 string

doneWithStruct8zgensym_6a7ebc1181a5fc8f_9:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft8zgensym_6a7ebc1181a5fc8f_9 > 0 || missingFieldsLeft8zgensym_6a7ebc1181a5fc8f_9 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft8zgensym_6a7ebc1181a5fc8f_9, missingFieldsLeft8zgensym_6a7ebc1181a5fc8f_9, msgp.ShowFound(found8zgensym_6a7ebc1181a5fc8f_9[:]), unmarshalMsgFieldOrder8zgensym_6a7ebc1181a5fc8f_9)
		if encodedFieldsLeft8zgensym_6a7ebc1181a5fc8f_9 > 0 {
			encodedFieldsLeft8zgensym_6a7ebc1181a5fc8f_9--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField8zgensym_6a7ebc1181a5fc8f_9 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss8zgensym_6a7ebc1181a5fc8f_9 < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss8zgensym_6a7ebc1181a5fc8f_9 = 0
			}
			for nextMiss8zgensym_6a7ebc1181a5fc8f_9 < maxFields8zgensym_6a7ebc1181a5fc8f_9 && (found8zgensym_6a7ebc1181a5fc8f_9[nextMiss8zgensym_6a7ebc1181a5fc8f_9] || unmarshalMsgFieldSkip8zgensym_6a7ebc1181a5fc8f_9[nextMiss8zgensym_6a7ebc1181a5fc8f_9]) {
				nextMiss8zgensym_6a7ebc1181a5fc8f_9++
			}
			if nextMiss8zgensym_6a7ebc1181a5fc8f_9 == maxFields8zgensym_6a7ebc1181a5fc8f_9 {
				// filled all the empty fields!
				break doneWithStruct8zgensym_6a7ebc1181a5fc8f_9
			}
			missingFieldsLeft8zgensym_6a7ebc1181a5fc8f_9--
			curField8zgensym_6a7ebc1181a5fc8f_9 = unmarshalMsgFieldOrder8zgensym_6a7ebc1181a5fc8f_9[nextMiss8zgensym_6a7ebc1181a5fc8f_9]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField8zgensym_6a7ebc1181a5fc8f_9)
		switch curField8zgensym_6a7ebc1181a5fc8f_9 {
		// -- templateUnmarshalMsg ends here --

		case "SignMe_zid00_rct":
			found8zgensym_6a7ebc1181a5fc8f_9[0] = true
			const maxFields10zgensym_6a7ebc1181a5fc8f_11 = 2

			// -- templateUnmarshalMsg starts here--
			var totalEncodedFields10zgensym_6a7ebc1181a5fc8f_11 uint32
			if !nbs.AlwaysNil {
				totalEncodedFields10zgensym_6a7ebc1181a5fc8f_11, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
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
				//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft10zgensym_6a7ebc1181a5fc8f_11, missingFieldsLeft10zgensym_6a7ebc1181a5fc8f_11, msgp.ShowFound(found10zgensym_6a7ebc1181a5fc8f_11[:]), unmarshalMsgFieldOrder10zgensym_6a7ebc1181a5fc8f_11)
				if encodedFieldsLeft10zgensym_6a7ebc1181a5fc8f_11 > 0 {
					encodedFieldsLeft10zgensym_6a7ebc1181a5fc8f_11--
					field, bts, err = nbs.ReadMapKeyZC(bts)
					if err != nil {
						return
					}
					curField10zgensym_6a7ebc1181a5fc8f_11 = msgp.UnsafeString(field)
				} else {
					//missing fields need handling
					if nextMiss10zgensym_6a7ebc1181a5fc8f_11 < 0 {
						// set bts to contain just mnil (0xc0)
						bts = nbs.PushAlwaysNil(bts)
						nextMiss10zgensym_6a7ebc1181a5fc8f_11 = 0
					}
					for nextMiss10zgensym_6a7ebc1181a5fc8f_11 < maxFields10zgensym_6a7ebc1181a5fc8f_11 && (found10zgensym_6a7ebc1181a5fc8f_11[nextMiss10zgensym_6a7ebc1181a5fc8f_11] || unmarshalMsgFieldSkip10zgensym_6a7ebc1181a5fc8f_11[nextMiss10zgensym_6a7ebc1181a5fc8f_11]) {
						nextMiss10zgensym_6a7ebc1181a5fc8f_11++
					}
					if nextMiss10zgensym_6a7ebc1181a5fc8f_11 == maxFields10zgensym_6a7ebc1181a5fc8f_11 {
						// filled all the empty fields!
						break doneWithStruct10zgensym_6a7ebc1181a5fc8f_11
					}
					missingFieldsLeft10zgensym_6a7ebc1181a5fc8f_11--
					curField10zgensym_6a7ebc1181a5fc8f_11 = unmarshalMsgFieldOrder10zgensym_6a7ebc1181a5fc8f_11[nextMiss10zgensym_6a7ebc1181a5fc8f_11]
				}
				//fmt.Printf("switching on curField: '%v'\n", curField10zgensym_6a7ebc1181a5fc8f_11)
				switch curField10zgensym_6a7ebc1181a5fc8f_11 {
				// -- templateUnmarshalMsg ends here --

				case "PubKeyBytes_zid00_bin":
					found10zgensym_6a7ebc1181a5fc8f_11[0] = true
					if nbs.AlwaysNil || msgp.IsNil(bts) {
						if !nbs.AlwaysNil {
							bts = bts[1:]
						}
						z.SignMe.PubKeyBytes = z.SignMe.PubKeyBytes[:0]
					} else {
						z.SignMe.PubKeyBytes, bts, err = nbs.ReadBytesBytes(bts, z.SignMe.PubKeyBytes)

						if err != nil {
							return
						}
					}
					if err != nil {
						return
					}
				case "SignedAtTimestamp_zid01_tim":
					found10zgensym_6a7ebc1181a5fc8f_11[1] = true
					z.SignMe.SignedAtTimestamp, bts, err = nbs.ReadTimeBytes(bts)

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
			if nextMiss10zgensym_6a7ebc1181a5fc8f_11 != -1 {
				bts = nbs.PopAlwaysNil()
			}

		case "SignatureFormat_zid01_str":
			found8zgensym_6a7ebc1181a5fc8f_9[1] = true
			z.SignatureFormat, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "SignatureBlob_zid02_bin":
			found8zgensym_6a7ebc1181a5fc8f_9[2] = true
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
	if nextMiss8zgensym_6a7ebc1181a5fc8f_9 != -1 {
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
var unmarshalMsgFieldOrder8zgensym_6a7ebc1181a5fc8f_9 = []string{"SignMe_zid00_rct", "SignatureFormat_zid01_str", "SignatureBlob_zid02_bin"}

var unmarshalMsgFieldSkip8zgensym_6a7ebc1181a5fc8f_9 = []bool{false, false, false}

// fields of SignedStuff
var unmarshalMsgFieldOrder10zgensym_6a7ebc1181a5fc8f_11 = []string{"PubKeyBytes_zid00_bin", "SignedAtTimestamp_zid01_tim"}

var unmarshalMsgFieldSkip10zgensym_6a7ebc1181a5fc8f_11 = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Selfie) Msgsize() (s int) {
	s = 1 + 17 + 1 + 22 + msgp.BytesPrefixSize + len(z.SignMe.PubKeyBytes) + 28 + msgp.TimeSize + 26 + msgp.StringPrefixSize + len(z.SignatureFormat) + 24 + msgp.BytesPrefixSize + len(z.SignatureBlob)
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
	const maxFields12zgensym_6a7ebc1181a5fc8f_13 = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields12zgensym_6a7ebc1181a5fc8f_13 uint32
	totalEncodedFields12zgensym_6a7ebc1181a5fc8f_13, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft12zgensym_6a7ebc1181a5fc8f_13 := totalEncodedFields12zgensym_6a7ebc1181a5fc8f_13
	missingFieldsLeft12zgensym_6a7ebc1181a5fc8f_13 := maxFields12zgensym_6a7ebc1181a5fc8f_13 - totalEncodedFields12zgensym_6a7ebc1181a5fc8f_13

	var nextMiss12zgensym_6a7ebc1181a5fc8f_13 int32 = -1
	var found12zgensym_6a7ebc1181a5fc8f_13 [maxFields12zgensym_6a7ebc1181a5fc8f_13]bool
	var curField12zgensym_6a7ebc1181a5fc8f_13 string

doneWithStruct12zgensym_6a7ebc1181a5fc8f_13:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft12zgensym_6a7ebc1181a5fc8f_13 > 0 || missingFieldsLeft12zgensym_6a7ebc1181a5fc8f_13 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft12zgensym_6a7ebc1181a5fc8f_13, missingFieldsLeft12zgensym_6a7ebc1181a5fc8f_13, msgp.ShowFound(found12zgensym_6a7ebc1181a5fc8f_13[:]), decodeMsgFieldOrder12zgensym_6a7ebc1181a5fc8f_13)
		if encodedFieldsLeft12zgensym_6a7ebc1181a5fc8f_13 > 0 {
			encodedFieldsLeft12zgensym_6a7ebc1181a5fc8f_13--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField12zgensym_6a7ebc1181a5fc8f_13 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss12zgensym_6a7ebc1181a5fc8f_13 < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss12zgensym_6a7ebc1181a5fc8f_13 = 0
			}
			for nextMiss12zgensym_6a7ebc1181a5fc8f_13 < maxFields12zgensym_6a7ebc1181a5fc8f_13 && (found12zgensym_6a7ebc1181a5fc8f_13[nextMiss12zgensym_6a7ebc1181a5fc8f_13] || decodeMsgFieldSkip12zgensym_6a7ebc1181a5fc8f_13[nextMiss12zgensym_6a7ebc1181a5fc8f_13]) {
				nextMiss12zgensym_6a7ebc1181a5fc8f_13++
			}
			if nextMiss12zgensym_6a7ebc1181a5fc8f_13 == maxFields12zgensym_6a7ebc1181a5fc8f_13 {
				// filled all the empty fields!
				break doneWithStruct12zgensym_6a7ebc1181a5fc8f_13
			}
			missingFieldsLeft12zgensym_6a7ebc1181a5fc8f_13--
			curField12zgensym_6a7ebc1181a5fc8f_13 = decodeMsgFieldOrder12zgensym_6a7ebc1181a5fc8f_13[nextMiss12zgensym_6a7ebc1181a5fc8f_13]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField12zgensym_6a7ebc1181a5fc8f_13)
		switch curField12zgensym_6a7ebc1181a5fc8f_13 {
		// -- templateDecodeMsg ends here --

		case "PubKeyBytes_zid00_bin":
			found12zgensym_6a7ebc1181a5fc8f_13[0] = true
			z.PubKeyBytes, err = dc.ReadBytes(z.PubKeyBytes)
			if err != nil {
				return
			}
		case "SignedAtTimestamp_zid01_tim":
			found12zgensym_6a7ebc1181a5fc8f_13[1] = true
			z.SignedAtTimestamp, err = dc.ReadTime()
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
	if nextMiss12zgensym_6a7ebc1181a5fc8f_13 != -1 {
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
var decodeMsgFieldOrder12zgensym_6a7ebc1181a5fc8f_13 = []string{"PubKeyBytes_zid00_bin", "SignedAtTimestamp_zid01_tim"}

var decodeMsgFieldSkip12zgensym_6a7ebc1181a5fc8f_13 = []bool{false, false}

// fieldsNotEmpty supports omitempty tags
func (z *SignedStuff) fieldsNotEmpty(isempty []bool) uint32 {
	if len(isempty) == 0 {
		return 2
	}
	var fieldsInUse uint32 = 2
	isempty[0] = (len(z.PubKeyBytes) == 0) // string, omitempty
	if isempty[0] {
		fieldsInUse--
	}
	isempty[1] = (z.SignedAtTimestamp.IsZero()) // time.Time, omitempty
	if isempty[1] {
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
	var empty_zgensym_6a7ebc1181a5fc8f_14 [2]bool
	fieldsInUse_zgensym_6a7ebc1181a5fc8f_15 := z.fieldsNotEmpty(empty_zgensym_6a7ebc1181a5fc8f_14[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zgensym_6a7ebc1181a5fc8f_15)
	if err != nil {
		return err
	}

	if !empty_zgensym_6a7ebc1181a5fc8f_14[0] {
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

	if !empty_zgensym_6a7ebc1181a5fc8f_14[1] {
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

	return
}

// MarshalMsg implements msgp.Marshaler
func (z *SignedStuff) MarshalMsg(b []byte) (o []byte, err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	o = msgp.Require(b, z.Msgsize())

	// honor the omitempty tags
	var empty [2]bool
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
	const maxFields16zgensym_6a7ebc1181a5fc8f_17 = 2

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
var unmarshalMsgFieldOrder16zgensym_6a7ebc1181a5fc8f_17 = []string{"PubKeyBytes_zid00_bin", "SignedAtTimestamp_zid01_tim"}

var unmarshalMsgFieldSkip16zgensym_6a7ebc1181a5fc8f_17 = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *SignedStuff) Msgsize() (s int) {
	s = 1 + 22 + msgp.BytesPrefixSize + len(z.PubKeyBytes) + 28 + msgp.TimeSize
	return
}
