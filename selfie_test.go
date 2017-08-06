package selfie

import (
	"fmt"

	cv "github.com/glycerine/goconvey/convey"
	"github.com/glycerine/sshego"
	"testing"
)

func Test001SelfSignedKeysValidate(t *testing.T) {

	cv.Convey("Given a self-signed public key, we should be able to correctly validate the signature", t, func() {

		storeToFile := ""
		bits := 1024
		email := "example@example.com"
		_, signer, err := sshego.GenRSAKeyPair(storeToFile, bits, email)
		panicOn(err)

		key := NewSSHKey()
		key.LoadPriv(signer, "")

		selfie, err := key.SelfSignKey()
		panicOn(err)

		valid, err := ValidateSelfSignedKey(selfie)
		panicOn(err)
		cv.So(valid, cv.ShouldBeTrue)
	})
}

func Test002NonSelfSignedKeysInvalid(t *testing.T) {

	cv.Convey("Given a public key signed by someone else, we should be able to correctly detect the invalid signature", t, func() {

		storeToFile := ""
		bits := 1024
		email := "example@example.com"
		_, signer0, err := sshego.GenRSAKeyPair(storeToFile, bits, email)
		panicOn(err)

		_, signer1, err := sshego.GenRSAKeyPair(storeToFile, bits, email)
		panicOn(err)

		key := NewSSHKey()
		key.LoadPriv(signer0, "")

		// sign with a different key than our own.
		selfie, err := key.signPublicKeyWith(signer1)
		panicOn(err)

		valid, err := ValidateSelfSignedKey(selfie)
		cv.So(err, cv.ShouldNotBeNil)
		cv.So(valid, cv.ShouldBeFalse)
	})
}

func Test003JSONserializationOfSelfie(t *testing.T) {

	cv.Convey("Given a Selfie, JSON serialization should work", t, func() {

		storeToFile := ""
		bits := 1024
		email := "example@example.com"
		_, signer, err := sshego.GenRSAKeyPair(storeToFile, bits, email)
		panicOn(err)

		key := NewSSHKey()
		key.LoadPriv(signer, "")

		selfie, err := key.SelfSignKey()
		panicOn(err)

		fmt.Printf("\n ----> json of selfie: '%s'\n", selfie.JSON())
	})
}
