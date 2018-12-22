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

var key byte

func init() {
	b := make([]byte, 1)
	n, err := rand.Read(b)
	if err != nil || n != 1 {
		panic("can't init key")
	}
	key = b[0]
}

// GUID is type of [16]byte
type GUID [16]byte

func (id *GUID) setVarient() {
	id[1] = id[1]&^slice1 | (key & slice1)
	id[4] = id[4]&^slice2 | (key & slice2)
	id[9] = id[9]&^slice3 | (key & slice3)
	id[14] = id[14]&^slice4 | (key & slice4)
}

// Verify verify whether id is creare with key ; default key is random
func (id *GUID) Verify() bool {
	return id[1]&slice1|(id[4]&slice2)|(id[9]&slice3)|(id[14]&slice4) == key
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

// SetSercetKey set key if you want to specify key;
// sometimes it doesn't need to call this function
func SetSercetKey(customSerKey byte) {
	key = customSerKey
}
