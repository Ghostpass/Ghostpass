// Provides the high level definition for the `Field` struct, which encapsulates and performs AEAD
// on secrets in order to return for storage back into the secret store mapping.
package ghostpass

import (
	"errors"
	"github.com/awnumar/memguard"
	"strings"
)

// Represents a strongly typed field, a struct that encapsulates a secret attribute that represents
// an encrypted username and password combination. Given a deniable combo pair, the secret can be
// mutated through a one-time pad and a deniable key can be derived for plausible deniability
type Field struct {

	// auth credentials are securely stored for fast retrieval in memory when deserialized, but
	// will never show up in persistent storage for security.
	Username *memguard.Enclave `json:"-"`
	Pwd      *memguard.Enclave `json:"-"`

	// encrypted secret of auth combo is persistently stored, and used to recover the pair
	// once deserialized back to memory securely.
	AuthPair []byte `json:"authpair"`

	// stores n number of deniable authpairs that can revealed from a generated key
	DeniablePairs [][]byte `json:"-"`
}

// Given a key, service key and auth combination, create a completely new field that is encrypted.
func NewField(key []byte, username string, pwd *memguard.Enclave) (*Field, error) {

	// unseal the password
	clearpwd, err := pwd.Open()
	if err != nil {
		return nil, err
	}

	// TODO: symmetrically encrypt pwd once first

	// initialize the secret by concating: `username:pwdstr`.
	var secretstr strings.Builder
	secretstr.WriteString(username)
	secretstr.WriteString(":")
	secretstr.WriteString(string(clearpwd.Bytes()))

	// encrypt the secret with the key
	secret, err := BoxEncrypt(key, []byte(secretstr.String()))
	if err != nil {
		return nil, err
	}

	// memguard pwdstr and username
	user_enclave := memguard.NewBufferFromBytes([]byte(username))

	return &Field{
		Username:      user_enclave.Seal(),
		Pwd:           pwd,
		AuthPair:      secret,
		DeniablePairs: nil,
	}, nil
}

// Given a compressed secret, reconstruct a `Field` by decrypting it with a symmetric key, and re-deriving
// the username and password securely from them. This is used if the store being deserialized is from a plainsight
// state, where no field structure is JSONified and needs to be reconstructed completely.
func ReconstructField(key []byte, compressed []byte) (*Field, error) {

	// create empty field, and partially initialize
	var field Field
	field.AuthPair = compressed

	// rederive auth pair with symmetric key
	err := field.RederiveAuthPair(key)
	if err != nil {
		return nil, err
	}

	// return populated field
	return &field, nil
}

// Given a partially initialized Field, like one being deserialized from a stationary store, rederive the
// user and encrypted password for retrieval by a user in-memory.
func (f *Field) RederiveAuthPair(key []byte) error {

	// sanity checks
	if f.AuthPair == nil {
		return errors.New("No secret in field")
	}

	// decrypt the secret field in order to recover username and pwd
	plaintext, err := BoxDecrypt(key, f.AuthPair)
	if err != nil {
		return err
	}

	// split by by colon and return substrings
	creds := strings.Split(string(plaintext), ":")
	user, pwd := creds[0], creds[1]

	// memguard username and encrypted password
	// if a key generated by a deniable pair is used, the bogus user and password will be set instead
	user_enclave := memguard.NewBufferFromBytes([]byte(user))
	pwd_enclave := memguard.NewBufferFromBytes([]byte(pwd))

	// we now reinitialize the field with the cleartext username, encrypted password,
	// and a secret checksum representing their resultant encryption.
	f.Username = user_enclave.Seal()
	f.Pwd = pwd_enclave.Seal()
	return nil
}

// Given a bogus and deniable auth combo, generate a secret like with the original pair and store it for
// deniable key generation later. (TODO)
func (f *Field) AddDeniableSecret(username string, pwd *memguard.Enclave) error {
	// unseal the password
	clearpwd, err := pwd.Open()
	if err != nil {
		return err
	}

	// TODO: symmetrically encrypt pwd once first

	// initialize the secret by concating: `username:pwdstr`.
	var secretstr strings.Builder
	secretstr.WriteString(username)
	secretstr.WriteString(":")
	secretstr.WriteString(string(clearpwd.Bytes()))

	// add bogus deniable pair
	f.DeniablePairs = append(f.DeniablePairs, []byte(secretstr.String()))
	return nil
}
