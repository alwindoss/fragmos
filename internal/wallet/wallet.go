package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"

	"github.com/btcsuite/btcd/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
)

func New() (*Wallet, error) {
	w := &Wallet{}
	privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		err = fmt.Errorf("unable to generate key: %w", err)
		return nil, err
	}
	w.privateKey = privKey
	w.publicKey = &privKey.PublicKey
	// 2. Perform SHA256 hasing on the public key (32 bytes)
	h2 := sha256.New()
	h2.Write(w.publicKey.X.Bytes())
	h2.Write(w.publicKey.Y.Bytes())
	digest2 := h2.Sum(nil)

	// 3. perform RIPEMD160 hashing on the result of the SHA256 (20 bytes)
	h3 := ripemd160.New()
	h3.Write(digest2)
	digest3 := h3.Sum(nil)

	// 4. Add version byte in front of RIPEMD160 hash (0x00 for Main Network)
	vd4 := make([]byte, 21)
	vd4[0] = 0x00
	copy(vd4[1:], digest3[:])

	// 5. Perform SHA256 on the versioned hash
	h5 := sha256.New()
	h5.Write(vd4)
	digest5 := h5.Sum(nil)

	// 6. Perform SHA256 hash on the result of the previous SHA256
	h6 := sha256.New()
	h6.Write(digest5)
	digest6 := h6.Sum(nil)

	// 7. Take the first 4 bytes of the second SHA256 hash for checksum
	checksum := digest6[:4]

	// 8. Add the 4 checksum bytes from 7 at the end of extended RIPEMD160 hash from 4 (25 bytes)
	d8 := make([]byte, 25)
	copy(d8[:21], vd4[:])
	copy(d8[21:], checksum[:])
	// 9. Convert the result from a byte string into base58
	blkChnAddr := base58.Encode(d8)
	w.blockchainAddress = blkChnAddr

	return w, nil
}

type Wallet struct {
	privateKey        *ecdsa.PrivateKey
	publicKey         *ecdsa.PublicKey
	blockchainAddress string
}

func (w *Wallet) PrivateKey() *ecdsa.PrivateKey {
	return w.privateKey
}

func (w *Wallet) PublicKey() *ecdsa.PublicKey {
	return w.publicKey
}

func (w *Wallet) PublicKeyString() string {
	return fmt.Sprintf("%x", w.privateKey.D.Bytes())
}

func (w *Wallet) PrivateKeyString() string {
	return fmt.Sprintf("%x%x", w.privateKey.X.Bytes(), w.privateKey.Y.Bytes())
}

func (w *Wallet) BlockchainAddress() string {
	return w.blockchainAddress
}
