/*
Package rfc6979 is an implementation of RFC 6979's deterministic DSA.

	Such signatures are compatible with standard Digital Signature Algorithm
	(DSA) and Elliptic Curve Digital Signature Algorithm (ECDSA) digital
	signatures and can be processed with unmodified verifiers, which need not be
	aware of the procedure described therein.  Deterministic signatures retain
	the cryptographic security features associated with digital signatures but
	can be more easily implemented in various environments, since they do not
	need access to a source of high-quality randomness.

(https://tools.ietf.org/html/rfc6979)

Provides functions similar to crypto/dsa and crypto/ecdsa.
*/
package rfc6979

import (
	"hash"
	"crypto/hmac"
	"math/big"
	"bytes"
)

// mac returns an HMAC of the given key and message.
func mac(alg func() hash.Hash, k, m, buf []byte) []byte {
	h := hmac.New(alg, k)
	h.Write(m)
	return h.Sum(buf[:0])
}

// https://tools.ietf.org/html/rfc6979#section-2.3.2
func bits2int(in []byte, qlen int) *big.Int {
	vlen := len(in) * 8
	v := new(big.Int).SetBytes(in)
	if vlen > qlen {
		v = new(big.Int).Rsh(v, uint(vlen-qlen))
	}
	return v
}

// https://tools.ietf.org/html/rfc6979#section-2.3.3
func int2octets(v *big.Int, rolen int) []byte {
	out := v.Bytes()

	//pad with zeros if it's too short
	if len(out) < rolen {
		out2 := make([]byte, rolen)
		copy(out2[rolen-len(out):], out)
		return out2
	}

	//drop most significant bytes if it's too long
	if len(out) > rolen {
		out2 := make([]byte, rolen)
		copy(out2, out[len(out)-rolen:])
		return out2
	}

	return out
}

// https://tools.ietf.org/html/rfc6979#section-2.3.4
func bits2octets(in []byte, q *big.Int, qlen, rolen int) []byte {
	z1 := bits2int(in, qlen)
	z2 := new(big.Int).Sub(z1, q)
	if z2.Sign() < 0 {
		return int2octets(z1, rolen)
	}
	return int2octets(z2, rolen)
}

var one = big.NewInt(1)

// https://tools.ietf.org/html/rfc6979#section-3.2
func generateSecret(q, x *big.Int, alg func() hash.Hash, hash []byte, test func(*big.Int) bool) {
	qlen := q.BitLen()
	holen := alg().Size()
	rolen := (qlen + 7) >> 3
	bx := append(int2octets(x, rolen), bits2octets(hash, q, qlen, rolen)...)

	//step B
	v := bytes.Repeat([]byte{0x01}, holen)

	//step C
	k := bytes.Repeat([]byte{0x00}, holen)

	//step D
	k = mac(alg, k, append(append(v, 0x00), bx...), k)

	//step E
	v = mac(alg, k, v, v)

	//step F
	k = mac(alg, k , append(append(v, 0x01), bx...), k)

	//step G
	v= mac(alg, k, v, v)

	//step H
	for {
		//step H1
		var t []byte

		//step H2
		for len(t) < qlen/8 {
			v = mac(alg, k, v, v)
			t = append(t, v...)
		}

		//step H3
		secret := bits2int(t, qlen)
		if secret.Cmp(one) >= 0 && secret.Cmp(q) < 0 && test(secret) {
			return
		}
		k = mac(alg, k, append(v, 0x00), k)
		v = mac(alg, k, v, v)
	}
}













