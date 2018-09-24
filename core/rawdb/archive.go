package rawdb

import (
	"encoding/binary"

	"github.com/gochain-io/gochain/common"
)

// DBArchivePrefixes is the set of key prefixes which are eligible for archival.
var DBArchivePrefixes = [...]byte{blockBodyPrefix[0], blockReceiptsPrefix[0], headerPrefix[0]}

// DBArchiveKey checks if a key is archivable, and returns its parts if so.
func DBArchiveKey(key []byte) (bool, byte, uint64, common.Hash) {
	if len(key) != 41 {
		return false, 0, 0, common.Hash{}
	}
	switch key[0] {
	case headerPrefix[0], blockBodyPrefix[0], blockReceiptsPrefix[0]:
		return true, key[0], binary.BigEndian.Uint64(key[1:]), common.BytesToHash(key[9:])
	}
	return false, 0, 0, common.Hash{}
}
