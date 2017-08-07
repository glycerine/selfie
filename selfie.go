package selfie

import (
	"bytes"
	cry "crypto/rand"
	"fmt"
	"github.com/glycerine/greenpack/msgp"
	"github.com/glycerine/sshego"
	"golang.org/x/crypto/ssh"
	"time"
)

//go:generate greenpack

// In order to communicate identity in
// a location independent manner, we
// will provide our public key along
// with our own signature of that
// public key + signing-time-timestamp,
// to prove that we are in recent
// posession of the private
// key that corresponds to the
// public key.

//msgp:ignore SSHKey

type SSHKey struct {
	PrivKey     ssh.Signer
	PrivKeyPath string

	PubKey      ssh.PublicKey
	PubKeyBytes []byte

	Signature *ssh.Signature

	Selfie *Selfie
}

func NewSSHKey() *SSHKey {
	return &SSHKey{}
}

type SignedStuff struct {
	// PubKeyBytes is the serialized key
	// data in SSH wire format,
	// with the name prefix.
	PubKeyBytes []byte `zid:"0"`

	// SignedAtTimestamp states when
	// the signature occured.
	SignedAtTimestamp time.Time `zid:"1"`
}

// Selfie is a self-signed public key.
type Selfie struct {
	SignMe SignedStuff `zid:"0"`

	// SignatureFormat and Signature
	// bytes are the components of the
	// ssh.Signature struct.
	SignatureFormat string `zid:"1"`
	SignatureBlob   []byte `zid:"2"`

	// Others we vouch for.
	Others []*Selfie `zid:"3"`
}

func (s *Selfie) JSON() string {
	by, err := s.MarshalMsg(nil)
	panicOn(err)
	var out bytes.Buffer
	in := bytes.NewBuffer(by)
	_, err = msgp.CopyToJSON(&out, in)
	panicOn(err)
	return out.String()
}

// ReadKey will replace s.PrivKey and s.PrivKeyPath
// with a key loaded from path
func (s *SSHKey) ReadKey(path string) error {
	priv, err := sshego.LoadRSAPrivateKey(path)
	if err != nil {
		return err
	}
	return s.LoadPriv(priv, path)
}

func (s *SSHKey) LoadPriv(priv ssh.Signer, path string) error {
	s.PrivKey = priv
	s.PrivKeyPath = path

	pub := s.PrivKey.PublicKey()
	s.PubKey = pub
	return nil
}

// SelfSignKey requires s.PrivKey and
// s.PubKey to be established
// already upon entry.
func (s *SSHKey) SelfSignKey() (*Selfie, error) {
	return s.signPublicKeyWith(s.PrivKey)
}

func (s *SSHKey) signPublicKeyWith(priv ssh.Signer) (*Selfie, error) {

	// Marshal returns the serialized
	// key data in SSH wire format,
	// with the name prefix.
	s.PubKeyBytes = s.PubKey.Marshal()

	// Sign returns raw signature
	// for the given data. This method
	// will apply the hash specified
	// for the keytype to the data.

	signme := SignedStuff{
		PubKeyBytes:       s.PubKeyBytes,
		SignedAtTimestamp: time.Now(),
	}
	mm, err := signme.MarshalMsg(nil)
	panicOn(err)

	sig, err := priv.Sign(cry.Reader, mm)
	if err != nil {
		return nil, err
	}
	s.Signature = sig
	selfie := &Selfie{
		SignMe:          signme,
		SignatureFormat: sig.Format,
		SignatureBlob:   sig.Blob,
	}
	s.Selfie = selfie
	return selfie, nil
}

func ValidateSelfSignedKey(selfie *Selfie) (valid bool, err error) {

	// reinflate the public key
	// and timestamp

	mm, err := selfie.SignMe.MarshalMsg(nil)
	panicOn(err)

	// ParseAuthorizedKeys parses a public
	// key from an authorized_keys
	// file used in OpenSSH according
	// to the sshd(8) manual page.
	pubkey, err := ssh.ParsePublicKey(selfie.SignMe.PubKeyBytes)
	panicOn(err)
	if err != nil {
		return false, err
	}

	sig := &ssh.Signature{
		Format: selfie.SignatureFormat,
		Blob:   selfie.SignatureBlob,
	}

	// Verify that sig is a signature on
	// the given data using this
	// key. This function will hash the
	// data appropriately first.
	err = pubkey.Verify(mm, sig)
	if err != nil {
		return false, err
	}

	// recursively validate
	for i, o := range selfie.Others {
		ov, err := ValidateSelfSignedKey(o)
		if err != nil {
			return false, err
		}
		if !ov {
			return false, fmt.Errorf("the i-th other was invalid", i)
		}
	}

	return true, nil
}
