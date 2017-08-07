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
	const maxFields1zgensym_6a7ebc1181a5fc8f_2 = 4

	// -- templateDecodeMsg starts here--
	var totalEncodedFields1zgensym_6a7ebc1181a5fc8f_2 uint32
	totalEncodedFields1zgensym_6a7ebc1181a5fc8f_2, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft1zgensym_6a7ebc1181a5fc8f_2 := totalEncodedFields1zgensym_6a7ebc1181a5fc8f_2
	missingFieldsLeft1zgensym_6a7ebc1181a5fc8f_2 := maxFields1zgensym_6a7ebc1181a5fc8f_2 - totalEncodedFields1zgensym_6a7ebc1181a5fc8f_2

	var nextMiss1zgensym_6a7ebc1181a5fc8f_2 int32 = -1
	var found1zgensym_6a7ebc1181a5fc8f_2 [maxFields1zgensym_6a7ebc1181a5fc8f_2]bool
	var curField1zgensym_6a7ebc1181a5fc8f_2 string

doneWithStruct1zgensym_6a7ebc1181a5fc8f_2:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zgensym_6a7ebc1181a5fc8f_2 > 0 || missingFieldsLeft1zgensym_6a7ebc1181a5fc8f_2 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zgensym_6a7ebc1181a5fc8f_2, missingFieldsLeft1zgensym_6a7ebc1181a5fc8f_2, msgp.ShowFound(found1zgensym_6a7ebc1181a5fc8f_2[:]), decodeMsgFieldOrder1zgensym_6a7ebc1181a5fc8f_2)
		if encodedFieldsLeft1zgensym_6a7ebc1181a5fc8f_2 > 0 {
			encodedFieldsLeft1zgensym_6a7ebc1181a5fc8f_2--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField1zgensym_6a7ebc1181a5fc8f_2 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zgensym_6a7ebc1181a5fc8f_2 < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss1zgensym_6a7ebc1181a5fc8f_2 = 0
			}
			for nextMiss1zgensym_6a7ebc1181a5fc8f_2 < maxFields1zgensym_6a7ebc1181a5fc8f_2 && (found1zgensym_6a7ebc1181a5fc8f_2[nextMiss1zgensym_6a7ebc1181a5fc8f_2] || decodeMsgFieldSkip1zgensym_6a7ebc1181a5fc8f_2[nextMiss1zgensym_6a7ebc1181a5fc8f_2]) {
				nextMiss1zgensym_6a7ebc1181a5fc8f_2++
			}
			if nextMiss1zgensym_6a7ebc1181a5fc8f_2 == maxFields1zgensym_6a7ebc1181a5fc8f_2 {
				// filled all the empty fields!
				break doneWithStruct1zgensym_6a7ebc1181a5fc8f_2
			}
			missingFieldsLeft1zgensym_6a7ebc1181a5fc8f_2--
			curField1zgensym_6a7ebc1181a5fc8f_2 = decodeMsgFieldOrder1zgensym_6a7ebc1181a5fc8f_2[nextMiss1zgensym_6a7ebc1181a5fc8f_2]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zgensym_6a7ebc1181a5fc8f_2)
		switch curField1zgensym_6a7ebc1181a5fc8f_2 {
		// -- templateDecodeMsg ends here --

		case "SignMe_zid00_rct":
			found1zgensym_6a7ebc1181a5fc8f_2[0] = true
			const maxFields3zgensym_6a7ebc1181a5fc8f_4 = 2

			// -- templateDecodeMsg starts here--
			var totalEncodedFields3zgensym_6a7ebc1181a5fc8f_4 uint32
			totalEncodedFields3zgensym_6a7ebc1181a5fc8f_4, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			encodedFieldsLeft3zgensym_6a7ebc1181a5fc8f_4 := totalEncodedFields3zgensym_6a7ebc1181a5fc8f_4
			missingFieldsLeft3zgensym_6a7ebc1181a5fc8f_4 := maxFields3zgensym_6a7ebc1181a5fc8f_4 - totalEncodedFields3zgensym_6a7ebc1181a5fc8f_4

			var nextMiss3zgensym_6a7ebc1181a5fc8f_4 int32 = -1
			var found3zgensym_6a7ebc1181a5fc8f_4 [maxFields3zgensym_6a7ebc1181a5fc8f_4]bool
			var curField3zgensym_6a7ebc1181a5fc8f_4 string

		doneWithStruct3zgensym_6a7ebc1181a5fc8f_4:
			// First fill all the encoded fields, then
			// treat the remaining, missing fields, as Nil.
			for encodedFieldsLeft3zgensym_6a7ebc1181a5fc8f_4 > 0 || missingFieldsLeft3zgensym_6a7ebc1181a5fc8f_4 > 0 {
				//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zgensym_6a7ebc1181a5fc8f_4, missingFieldsLeft3zgensym_6a7ebc1181a5fc8f_4, msgp.ShowFound(found3zgensym_6a7ebc1181a5fc8f_4[:]), decodeMsgFieldOrder3zgensym_6a7ebc1181a5fc8f_4)
				if encodedFieldsLeft3zgensym_6a7ebc1181a5fc8f_4 > 0 {
					encodedFieldsLeft3zgensym_6a7ebc1181a5fc8f_4--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						return
					}
					curField3zgensym_6a7ebc1181a5fc8f_4 = msgp.UnsafeString(field)
				} else {
					//missing fields need handling
					if nextMiss3zgensym_6a7ebc1181a5fc8f_4 < 0 {
						// tell the reader to only give us Nils
						// until further notice.
						dc.PushAlwaysNil()
						nextMiss3zgensym_6a7ebc1181a5fc8f_4 = 0
					}
					for nextMiss3zgensym_6a7ebc1181a5fc8f_4 < maxFields3zgensym_6a7ebc1181a5fc8f_4 && (found3zgensym_6a7ebc1181a5fc8f_4[nextMiss3zgensym_6a7ebc1181a5fc8f_4] || decodeMsgFieldSkip3zgensym_6a7ebc1181a5fc8f_4[nextMiss3zgensym_6a7ebc1181a5fc8f_4]) {
						nextMiss3zgensym_6a7ebc1181a5fc8f_4++
					}
					if nextMiss3zgensym_6a7ebc1181a5fc8f_4 == maxFields3zgensym_6a7ebc1181a5fc8f_4 {
						// filled all the empty fields!
						break doneWithStruct3zgensym_6a7ebc1181a5fc8f_4
					}
					missingFieldsLeft3zgensym_6a7ebc1181a5fc8f_4--
					curField3zgensym_6a7ebc1181a5fc8f_4 = decodeMsgFieldOrder3zgensym_6a7ebc1181a5fc8f_4[nextMiss3zgensym_6a7ebc1181a5fc8f_4]
				}
				//fmt.Printf("switching on curField: '%v'\n", curField3zgensym_6a7ebc1181a5fc8f_4)
				switch curField3zgensym_6a7ebc1181a5fc8f_4 {
				// -- templateDecodeMsg ends here --

				case "PubKeyBytes_zid00_bin":
					found3zgensym_6a7ebc1181a5fc8f_4[0] = true
					z.SignMe.PubKeyBytes, err = dc.ReadBytes(z.SignMe.PubKeyBytes)
					if err != nil {
						return
					}
				case "SignedAtTimestamp_zid01_tim":
					found3zgensym_6a7ebc1181a5fc8f_4[1] = true
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
			if nextMiss3zgensym_6a7ebc1181a5fc8f_4 != -1 {
				dc.PopAlwaysNil()
			}

		case "SignatureFormat_zid01_str":
			found1zgensym_6a7ebc1181a5fc8f_2[1] = true
			z.SignatureFormat, err = dc.ReadString()
			if err != nil {
				return
			}
		case "SignatureBlob_zid02_bin":
			found1zgensym_6a7ebc1181a5fc8f_2[2] = true
			z.SignatureBlob, err = dc.ReadBytes(z.SignatureBlob)
			if err != nil {
				return
			}
		case "Others_zid03_slc":
			found1zgensym_6a7ebc1181a5fc8f_2[3] = true
			var zgensym_6a7ebc1181a5fc8f_5 uint32
			zgensym_6a7ebc1181a5fc8f_5, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Others) >= int(zgensym_6a7ebc1181a5fc8f_5) {
				z.Others = (z.Others)[:zgensym_6a7ebc1181a5fc8f_5]
			} else {
				z.Others = make([]*Selfie, zgensym_6a7ebc1181a5fc8f_5)
			}
			for zgensym_6a7ebc1181a5fc8f_0 := range z.Others {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					if z.Others[zgensym_6a7ebc1181a5fc8f_0] != nil {
						dc.PushAlwaysNil()
						err = z.Others[zgensym_6a7ebc1181a5fc8f_0].DecodeMsg(dc)
						if err != nil {
							return
						}
						dc.PopAlwaysNil()
					}
				} else {
					// not Nil, we have something to read

					if z.Others[zgensym_6a7ebc1181a5fc8f_0] == nil {
						z.Others[zgensym_6a7ebc1181a5fc8f_0] = new(Selfie)
					}
					err = z.Others[zgensym_6a7ebc1181a5fc8f_0].DecodeMsg(dc)
					if err != nil {
						return
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
	if nextMiss1zgensym_6a7ebc1181a5fc8f_2 != -1 {
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
var decodeMsgFieldOrder1zgensym_6a7ebc1181a5fc8f_2 = []string{"SignMe_zid00_rct", "SignatureFormat_zid01_str", "SignatureBlob_zid02_bin", "Others_zid03_slc"}

var decodeMsgFieldSkip1zgensym_6a7ebc1181a5fc8f_2 = []bool{false, false, false, false}

// fields of SignedStuff
var decodeMsgFieldOrder3zgensym_6a7ebc1181a5fc8f_4 = []string{"PubKeyBytes_zid00_bin", "SignedAtTimestamp_zid01_tim"}

var decodeMsgFieldSkip3zgensym_6a7ebc1181a5fc8f_4 = []bool{false, false}

// fieldsNotEmpty supports omitempty tags
func (z *Selfie) fieldsNotEmpty(isempty []bool) uint32 {
	if len(isempty) == 0 {
		return 4
	}
	var fieldsInUse uint32 = 4
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
	isempty[3] = (len(z.Others) == 0) // string, omitempty
	if isempty[3] {
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
	var empty_zgensym_6a7ebc1181a5fc8f_6 [4]bool
	fieldsInUse_zgensym_6a7ebc1181a5fc8f_7 := z.fieldsNotEmpty(empty_zgensym_6a7ebc1181a5fc8f_6[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zgensym_6a7ebc1181a5fc8f_7)
	if err != nil {
		return err
	}

	if !empty_zgensym_6a7ebc1181a5fc8f_6[0] {
		// write "SignMe_zid00_rct"
		err = en.Append(0xb0, 0x53, 0x69, 0x67, 0x6e, 0x4d, 0x65, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x72, 0x63, 0x74)
		if err != nil {
			return err
		}

		// honor the omitempty tags
		var empty_zgensym_6a7ebc1181a5fc8f_8 [2]bool
		fieldsInUse_zgensym_6a7ebc1181a5fc8f_9 := z.SignMe.fieldsNotEmpty(empty_zgensym_6a7ebc1181a5fc8f_8[:])

		// map header
		err = en.WriteMapHeader(fieldsInUse_zgensym_6a7ebc1181a5fc8f_9)
		if err != nil {
			return err
		}

		if !empty_zgensym_6a7ebc1181a5fc8f_8[0] {
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

		if !empty_zgensym_6a7ebc1181a5fc8f_8[1] {
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

	if !empty_zgensym_6a7ebc1181a5fc8f_6[1] {
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

	if !empty_zgensym_6a7ebc1181a5fc8f_6[2] {
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

	if !empty_zgensym_6a7ebc1181a5fc8f_6[3] {
		// write "Others_zid03_slc"
		err = en.Append(0xb0, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x73, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x33, 0x5f, 0x73, 0x6c, 0x63)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(z.Others)))
		if err != nil {
			return
		}
		for zgensym_6a7ebc1181a5fc8f_0 := range z.Others {
			if z.Others[zgensym_6a7ebc1181a5fc8f_0] == nil {
				err = en.WriteNil()
				if err != nil {
					return
				}
			} else {
				err = z.Others[zgensym_6a7ebc1181a5fc8f_0].EncodeMsg(en)
				if err != nil {
					return
				}
			}
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
	var empty [4]bool
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

	if !empty[3] {
		// string "Others_zid03_slc"
		o = append(o, 0xb0, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x73, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x33, 0x5f, 0x73, 0x6c, 0x63)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Others)))
		for zgensym_6a7ebc1181a5fc8f_0 := range z.Others {
			if z.Others[zgensym_6a7ebc1181a5fc8f_0] == nil {
				o = msgp.AppendNil(o)
			} else {
				o, err = z.Others[zgensym_6a7ebc1181a5fc8f_0].MarshalMsg(o)
				if err != nil {
					return
				}
			}
		}
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
	const maxFields10zgensym_6a7ebc1181a5fc8f_11 = 4

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

		case "SignMe_zid00_rct":
			found10zgensym_6a7ebc1181a5fc8f_11[0] = true
			const maxFields12zgensym_6a7ebc1181a5fc8f_13 = 2

			// -- templateUnmarshalMsg starts here--
			var totalEncodedFields12zgensym_6a7ebc1181a5fc8f_13 uint32
			if !nbs.AlwaysNil {
				totalEncodedFields12zgensym_6a7ebc1181a5fc8f_13, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
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
				//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft12zgensym_6a7ebc1181a5fc8f_13, missingFieldsLeft12zgensym_6a7ebc1181a5fc8f_13, msgp.ShowFound(found12zgensym_6a7ebc1181a5fc8f_13[:]), unmarshalMsgFieldOrder12zgensym_6a7ebc1181a5fc8f_13)
				if encodedFieldsLeft12zgensym_6a7ebc1181a5fc8f_13 > 0 {
					encodedFieldsLeft12zgensym_6a7ebc1181a5fc8f_13--
					field, bts, err = nbs.ReadMapKeyZC(bts)
					if err != nil {
						return
					}
					curField12zgensym_6a7ebc1181a5fc8f_13 = msgp.UnsafeString(field)
				} else {
					//missing fields need handling
					if nextMiss12zgensym_6a7ebc1181a5fc8f_13 < 0 {
						// set bts to contain just mnil (0xc0)
						bts = nbs.PushAlwaysNil(bts)
						nextMiss12zgensym_6a7ebc1181a5fc8f_13 = 0
					}
					for nextMiss12zgensym_6a7ebc1181a5fc8f_13 < maxFields12zgensym_6a7ebc1181a5fc8f_13 && (found12zgensym_6a7ebc1181a5fc8f_13[nextMiss12zgensym_6a7ebc1181a5fc8f_13] || unmarshalMsgFieldSkip12zgensym_6a7ebc1181a5fc8f_13[nextMiss12zgensym_6a7ebc1181a5fc8f_13]) {
						nextMiss12zgensym_6a7ebc1181a5fc8f_13++
					}
					if nextMiss12zgensym_6a7ebc1181a5fc8f_13 == maxFields12zgensym_6a7ebc1181a5fc8f_13 {
						// filled all the empty fields!
						break doneWithStruct12zgensym_6a7ebc1181a5fc8f_13
					}
					missingFieldsLeft12zgensym_6a7ebc1181a5fc8f_13--
					curField12zgensym_6a7ebc1181a5fc8f_13 = unmarshalMsgFieldOrder12zgensym_6a7ebc1181a5fc8f_13[nextMiss12zgensym_6a7ebc1181a5fc8f_13]
				}
				//fmt.Printf("switching on curField: '%v'\n", curField12zgensym_6a7ebc1181a5fc8f_13)
				switch curField12zgensym_6a7ebc1181a5fc8f_13 {
				// -- templateUnmarshalMsg ends here --

				case "PubKeyBytes_zid00_bin":
					found12zgensym_6a7ebc1181a5fc8f_13[0] = true
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
					found12zgensym_6a7ebc1181a5fc8f_13[1] = true
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
			if nextMiss12zgensym_6a7ebc1181a5fc8f_13 != -1 {
				bts = nbs.PopAlwaysNil()
			}

		case "SignatureFormat_zid01_str":
			found10zgensym_6a7ebc1181a5fc8f_11[1] = true
			z.SignatureFormat, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "SignatureBlob_zid02_bin":
			found10zgensym_6a7ebc1181a5fc8f_11[2] = true
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
		case "Others_zid03_slc":
			found10zgensym_6a7ebc1181a5fc8f_11[3] = true
			if nbs.AlwaysNil {
				(z.Others) = (z.Others)[:0]
			} else {

				var zgensym_6a7ebc1181a5fc8f_14 uint32
				zgensym_6a7ebc1181a5fc8f_14, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.Others) >= int(zgensym_6a7ebc1181a5fc8f_14) {
					z.Others = (z.Others)[:zgensym_6a7ebc1181a5fc8f_14]
				} else {
					z.Others = make([]*Selfie, zgensym_6a7ebc1181a5fc8f_14)
				}
				for zgensym_6a7ebc1181a5fc8f_0 := range z.Others {
					if nbs.AlwaysNil {
						if z.Others[zgensym_6a7ebc1181a5fc8f_0] != nil {
							z.Others[zgensym_6a7ebc1181a5fc8f_0].UnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != z.Others[zgensym_6a7ebc1181a5fc8f_0] {
								z.Others[zgensym_6a7ebc1181a5fc8f_0].UnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if z.Others[zgensym_6a7ebc1181a5fc8f_0] == nil {
								z.Others[zgensym_6a7ebc1181a5fc8f_0] = new(Selfie)
							}
							bts, err = z.Others[zgensym_6a7ebc1181a5fc8f_0].UnmarshalMsg(bts)
							if err != nil {
								return
							}
							if err != nil {
								return
							}
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
	if nextMiss10zgensym_6a7ebc1181a5fc8f_11 != -1 {
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
var unmarshalMsgFieldOrder10zgensym_6a7ebc1181a5fc8f_11 = []string{"SignMe_zid00_rct", "SignatureFormat_zid01_str", "SignatureBlob_zid02_bin", "Others_zid03_slc"}

var unmarshalMsgFieldSkip10zgensym_6a7ebc1181a5fc8f_11 = []bool{false, false, false, false}

// fields of SignedStuff
var unmarshalMsgFieldOrder12zgensym_6a7ebc1181a5fc8f_13 = []string{"PubKeyBytes_zid00_bin", "SignedAtTimestamp_zid01_tim"}

var unmarshalMsgFieldSkip12zgensym_6a7ebc1181a5fc8f_13 = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Selfie) Msgsize() (s int) {
	s = 1 + 17 + 1 + 22 + msgp.BytesPrefixSize + len(z.SignMe.PubKeyBytes) + 28 + msgp.TimeSize + 26 + msgp.StringPrefixSize + len(z.SignatureFormat) + 24 + msgp.BytesPrefixSize + len(z.SignatureBlob) + 17 + msgp.ArrayHeaderSize
	for zgensym_6a7ebc1181a5fc8f_0 := range z.Others {
		if z.Others[zgensym_6a7ebc1181a5fc8f_0] == nil {
			s += msgp.NilSize
		} else {
			s += z.Others[zgensym_6a7ebc1181a5fc8f_0].Msgsize()
		}
	}
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
	const maxFields15zgensym_6a7ebc1181a5fc8f_16 = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields15zgensym_6a7ebc1181a5fc8f_16 uint32
	totalEncodedFields15zgensym_6a7ebc1181a5fc8f_16, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft15zgensym_6a7ebc1181a5fc8f_16 := totalEncodedFields15zgensym_6a7ebc1181a5fc8f_16
	missingFieldsLeft15zgensym_6a7ebc1181a5fc8f_16 := maxFields15zgensym_6a7ebc1181a5fc8f_16 - totalEncodedFields15zgensym_6a7ebc1181a5fc8f_16

	var nextMiss15zgensym_6a7ebc1181a5fc8f_16 int32 = -1
	var found15zgensym_6a7ebc1181a5fc8f_16 [maxFields15zgensym_6a7ebc1181a5fc8f_16]bool
	var curField15zgensym_6a7ebc1181a5fc8f_16 string

doneWithStruct15zgensym_6a7ebc1181a5fc8f_16:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft15zgensym_6a7ebc1181a5fc8f_16 > 0 || missingFieldsLeft15zgensym_6a7ebc1181a5fc8f_16 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft15zgensym_6a7ebc1181a5fc8f_16, missingFieldsLeft15zgensym_6a7ebc1181a5fc8f_16, msgp.ShowFound(found15zgensym_6a7ebc1181a5fc8f_16[:]), decodeMsgFieldOrder15zgensym_6a7ebc1181a5fc8f_16)
		if encodedFieldsLeft15zgensym_6a7ebc1181a5fc8f_16 > 0 {
			encodedFieldsLeft15zgensym_6a7ebc1181a5fc8f_16--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField15zgensym_6a7ebc1181a5fc8f_16 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss15zgensym_6a7ebc1181a5fc8f_16 < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss15zgensym_6a7ebc1181a5fc8f_16 = 0
			}
			for nextMiss15zgensym_6a7ebc1181a5fc8f_16 < maxFields15zgensym_6a7ebc1181a5fc8f_16 && (found15zgensym_6a7ebc1181a5fc8f_16[nextMiss15zgensym_6a7ebc1181a5fc8f_16] || decodeMsgFieldSkip15zgensym_6a7ebc1181a5fc8f_16[nextMiss15zgensym_6a7ebc1181a5fc8f_16]) {
				nextMiss15zgensym_6a7ebc1181a5fc8f_16++
			}
			if nextMiss15zgensym_6a7ebc1181a5fc8f_16 == maxFields15zgensym_6a7ebc1181a5fc8f_16 {
				// filled all the empty fields!
				break doneWithStruct15zgensym_6a7ebc1181a5fc8f_16
			}
			missingFieldsLeft15zgensym_6a7ebc1181a5fc8f_16--
			curField15zgensym_6a7ebc1181a5fc8f_16 = decodeMsgFieldOrder15zgensym_6a7ebc1181a5fc8f_16[nextMiss15zgensym_6a7ebc1181a5fc8f_16]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField15zgensym_6a7ebc1181a5fc8f_16)
		switch curField15zgensym_6a7ebc1181a5fc8f_16 {
		// -- templateDecodeMsg ends here --

		case "PubKeyBytes_zid00_bin":
			found15zgensym_6a7ebc1181a5fc8f_16[0] = true
			z.PubKeyBytes, err = dc.ReadBytes(z.PubKeyBytes)
			if err != nil {
				return
			}
		case "SignedAtTimestamp_zid01_tim":
			found15zgensym_6a7ebc1181a5fc8f_16[1] = true
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
	if nextMiss15zgensym_6a7ebc1181a5fc8f_16 != -1 {
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
var decodeMsgFieldOrder15zgensym_6a7ebc1181a5fc8f_16 = []string{"PubKeyBytes_zid00_bin", "SignedAtTimestamp_zid01_tim"}

var decodeMsgFieldSkip15zgensym_6a7ebc1181a5fc8f_16 = []bool{false, false}

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
	var empty_zgensym_6a7ebc1181a5fc8f_17 [2]bool
	fieldsInUse_zgensym_6a7ebc1181a5fc8f_18 := z.fieldsNotEmpty(empty_zgensym_6a7ebc1181a5fc8f_17[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zgensym_6a7ebc1181a5fc8f_18)
	if err != nil {
		return err
	}

	if !empty_zgensym_6a7ebc1181a5fc8f_17[0] {
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

	if !empty_zgensym_6a7ebc1181a5fc8f_17[1] {
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
	const maxFields19zgensym_6a7ebc1181a5fc8f_20 = 2

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

		case "PubKeyBytes_zid00_bin":
			found19zgensym_6a7ebc1181a5fc8f_20[0] = true
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
			found19zgensym_6a7ebc1181a5fc8f_20[1] = true
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
	if nextMiss19zgensym_6a7ebc1181a5fc8f_20 != -1 {
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
var unmarshalMsgFieldOrder19zgensym_6a7ebc1181a5fc8f_20 = []string{"PubKeyBytes_zid00_bin", "SignedAtTimestamp_zid01_tim"}

var unmarshalMsgFieldSkip19zgensym_6a7ebc1181a5fc8f_20 = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *SignedStuff) Msgsize() (s int) {
	s = 1 + 22 + msgp.BytesPrefixSize + len(z.PubKeyBytes) + 28 + msgp.TimeSize
	return
}
