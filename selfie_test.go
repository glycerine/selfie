package selfie

import (
	"fmt"

	cv "github.com/glycerine/goconvey/convey"
	"github.com/glycerine/sshego"
	"testing"
)

func Test001SelfSignedKeysValidate(t *testing.T) {

	cv.Convey("Given a self-signed public key, we should be able to correctly validate the signature", t, func() {

		selfie := NewTestSelfie("example@example.com")

		valid, err := ValidateSelfSignedKey(selfie)
		panicOn(err)
		cv.So(valid, cv.ShouldBeTrue)

		// corrupt the key and be sure we notice.
		selfie.SignMe.PubKeyBytes[0] = ^selfie.SignMe.PubKeyBytes[0]

		valid2, err := ValidateSelfSignedKey(selfie)
		cv.So(err, cv.ShouldNotBeNil)
		cv.So(valid2, cv.ShouldBeFalse)
	})
}

func Test002NonSelfSignedKeysInvalid(t *testing.T) {

	cv.Convey("Given a public key signed by someone else, we should be able to correctly detect the invalid signature", t, func() {

		storeToFile := ""
		bits := 1024
		email := "example@example.com"

		priv1, _, err := sshego.GenRSAKeyPair(storeToFile, bits, email)
		panicOn(err)

		key, err := NewSSHKey(email, bits, storeToFile)
		panicOn(err)

		// sign with a different key than our own.
		selfie, err := key.signPublicKeyWith(priv1)
		panicOn(err)

		valid, err := ValidateSelfSignedKey(selfie)
		cv.So(err, cv.ShouldNotBeNil)
		cv.So(valid, cv.ShouldBeFalse)
	})
}

func Test003JSONserializationOfSelfie(t *testing.T) {

	cv.Convey("Given a Selfie, JSON serialization should work", t, func() {

		s1 := NewTestSelfie("example@example.com")
		s2 := NewTestSelfie("dos@example.com")
		s1.SignMe.Others = []*Selfie{s2}

		fmt.Printf("\n ----> json of selfie: '%s'\n", s1.JSON())
	})
}

func NewTestSelfie(email string) *Selfie {

	storeToFile := ""

	// speed up tests, use fewer bits
	bits := 1024

	// for production, use
	//bits := 4096

	key, err := NewSSHKey(email, bits, storeToFile)
	panicOn(err)

	selfie, err := key.SelfSignKey()
	panicOn(err)
	return selfie
}
