package selfie

import (
	"bytes"
	"crypto"
	cry "crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
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

	RsaPriv *rsa.PrivateKey
	RsaPub  *rsa.PublicKey

	Signature *ssh.Signature

	Selfie *Selfie
	Bits   int
	Email  string
}

func NewSSHKey(email string, bits int, storeToFile string) (*SSHKey, error) {
	key := &SSHKey{
		Email:       email,
		Bits:        bits,
		PrivKeyPath: storeToFile,
	}
	rsaPriv, signer, err := sshego.GenRSAKeyPair(storeToFile, bits, email)
	if err != nil {
		return nil, err
	}
	key.PrivKey = signer
	key.RsaPriv = rsaPriv

	key.PubKey = key.PrivKey.PublicKey()
	key.RsaPub = &key.RsaPriv.PublicKey

	return key, nil
}

type SignedStuff struct {
	// PubKeyBytes is the serialized key
	// data in SSH wire format,
	// with the name prefix.
	PubKeyBytes []byte `zid:"0"`

	// SignedAtTimestamp states when
	// the signature occured.
	SignedAtTimestamp time.Time `zid:"1"`

	// Others we vouch for.
	Others []*Selfie `zid:"2"`
}

// Selfie is a self-signed public key.
type Selfie struct {
	SignMe SignedStuff `zid:"0"`

	// SignatureFormat and Signature
	// bytes are the components of the
	// ssh.Signature struct.
	SignatureFormat string `zid:"1"`
	SignatureBlob   []byte `zid:"2"`
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

// SelfSignKey requires s.PrivKey and
// s.PubKey to be established
// already upon entry.
func (s *SSHKey) SelfSignKey() (*Selfie, error) {
	return s.signPublicKeyWith(s.RsaPriv)
}

func (s *SSHKey) signPublicKeyWith(priv *rsa.PrivateKey) (*Selfie, error) {

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

	blob, err := SignSHA256(priv, mm)
	if err != nil {
		return nil, err
	}
	sig := &ssh.Signature{
		Format: "sha256",
		Blob:   blob,
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

	// reserialize the payload field SignMe

	mm, err := selfie.SignMe.MarshalMsg(nil)
	_ = mm
	panicOn(err)

	// ParseAuthorizedKeys parses a public
	// key from an authorized_keys
	// file used in OpenSSH according
	// to the sshd(8) manual page.
	pubkey, err := SshParseRsaPublicKey(selfie.SignMe.PubKeyBytes)
	if err != nil {
		return false, err
	}

	sig := &ssh.Signature{
		Format: selfie.SignatureFormat,
		Blob:   selfie.SignatureBlob,
	}
	_ = sig

	err = VerifySHA256(pubkey, mm, selfie.SignatureBlob)
	if err != nil {
		return false, err
	}

	// recursively validate
	for i, o := range selfie.SignMe.Others {
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

// Sign signs data with rsa-sha256
func SignSHA256(priv *rsa.PrivateKey, data []byte) ([]byte, error) {
	h := sha256.New()
	h.Write(data)
	d := h.Sum(nil)
	return rsa.SignPKCS1v15(cry.Reader, priv, crypto.SHA256, d)
}

// Unsign verifies the message using a rsa-sha256 signature
func VerifySHA256(r *rsa.PublicKey, message []byte, sig []byte) error {
	h := sha256.New()
	h.Write(message)
	d := h.Sum(nil)
	return rsa.VerifyPKCS1v15(r, crypto.SHA256, d, sig)
}
