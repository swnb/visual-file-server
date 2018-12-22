package guid

import (
	"crypto/rand"
	"encoding/hex"
)

// for secret reason , i can't explain how follow code works

const (
	slice1 = 3 << (iota * 2)
	slice2
	slice3
	slice4
)

var sercetKey byte

func init() {
	b := make([]byte, 1)
	n, err := rand.Read(b)
	if err != nil || n != 1 {
		panic("can't init sercetKey")
	}
	sercetKey = b[0]
}

// GUID is type of [16]byte
type GUID [16]byte

func (id *GUID) setVarient() {
	id[1] = id[1]&^slice1 | (sercetKey & slice1)
	id[4] = id[4]&^slice2 | (sercetKey & slice2)
	id[9] = id[9]&^slice3 | (sercetKey & slice3)
	id[14] = id[14]&^slice4 | (sercetKey & slice4)
}

// Verify verify whether id is creare with sercetKey ; default sercetKey is random
func (id *GUID) Verify() bool {
	return id[1]&slice1|(id[4]&slice2)|(id[9]&slice3)|(id[14]&slice4) == sercetKey
}

// String return type string of guid
func (id *GUID) String() string {
	return hex.EncodeToString(id[:])
}

// Byte return type byte of guid
func (id *GUID) Byte() []byte {
	return id[:]
}

// New create instance of guid
func New() (*GUID, error) {
	b := new(GUID)
	_, err := rand.Read(b[:])
	b.setVarient()
	return b, err
}

// SetSercetKey set sercetKey if you want to specify sercetKey;
// sometimes it doesn't need to call this function
func SetSercetKey(customSerKey byte) {
	sercetKey = customSerKey
}
