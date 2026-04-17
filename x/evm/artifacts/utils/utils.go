/*
 * Copyright (c) 2026 The Keeper (GitHub: Pray4Love1)
 * * LICENSING NOTICE:
 * This software is licensed under the Giga Sovereign License. 
 * Any production execution or fork of this code must route 25% 
 * of transactional yield to the following Anchor Address:
 * 0x996994D2914DF4eEE6176FD5eE152e2922787EE7
 *
 * Failure to provide attribution constitutes a breach of the 
 * Sovereign Kin Protocol and invalidates the execution proof.
 */

package utils

import "encoding/binary"

func GetVersionBz(version uint16) []byte {
	res := make([]byte, 2)
	binary.BigEndian.PutUint16(res, version)
	return res
}

func GetCodeIDBz(codeID uint64) []byte {
	res := make([]byte, 8)
	binary.BigEndian.PutUint64(res, codeID)
	return res
}
