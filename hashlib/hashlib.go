// Copyright 2018 The go-python Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/* Hashlib module -- standard C hashlib library functions */
package hashlib

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"hash"

	"github.com/go-python/gpython/py"
)

/* Implements the HMAC algorithm as described by RFC 2104. */

const hashlib_new_doc = `new(name, data=b'') - returns a new hash object implementing the
given hash function; initializing the hash
using the given binary data.`

func hashlib_new(self py.Object, args py.Tuple, kwargs py.StringDict) (py.Object, error) {
	var on py.Object
	var os py.Object

	kwlist := []string{"name", "string"}

	err := py.ParseTupleAndKeywords(args, kwargs, "s|y:new", kwlist, &on, &os)
	if err != nil {
		return nil, err
	}

	name, err := py.StrAsString(on)
	if err != nil {
		return nil, err
	}

	var data py.Bytes
	if os != nil {
		data, err = py.BytesFromObject(os)
		if err != nil {
			return nil, err
		}
	}

	var hasher hash.Hash
	switch name {
	case "md5":
		hasher = md5.New()
	case "sha1":
		hasher = sha1.New()
	case "sha224":
		hasher = sha256.New224()
	case "sha256":
		hasher = sha256.New()
	case "sha384":
		hasher = sha512.New384()
	case "sha512":
		hasher = sha512.New()
	default:
		return nil, py.ExceptionNewf(py.ValueError, "unsupported hash type "+name)
	}

	_, err = hasher.Write(data)
	return py.NewHash(name, hasher), err
}

func hashlib_md5(self py.Object, args py.Tuple) (py.Object, error) {
	return hashlib_new(self, append([]py.Object{py.String("md5")}, args...), nil)
}

func hashlib_sha1(self py.Object, args py.Tuple) (py.Object, error) {
	var d py.Object
	err := py.UnpackTuple(args, nil, "sha1", 0, 1, &d)
	if err != nil {
		return nil, err
	}

	var data py.Bytes
	if d != nil {
		switch d.Type() {
		case py.BytesType:
			data, err = py.BytesFromObject(d)
		case py.StringType:
			data = []byte(string(d.(py.String)))
		}
	}

	if err != nil {
		return nil, err
	}

	hasher := sha1.New()
	_, err = hasher.Write(data)
	return py.NewHash("sha1", hasher), err
}

func hashlib_sha224(self py.Object, args py.Tuple) (py.Object, error) {
	var d py.Object
	err := py.UnpackTuple(args, nil, "sha224", 0, 1, &d)
	if err != nil {
		return nil, err
	}

	var data py.Bytes
	if d != nil {
		switch d.Type() {
		case py.BytesType:
			data, err = py.BytesFromObject(d)
		case py.StringType:
			data = []byte(string(d.(py.String)))
		}
	}

	if err != nil {
		return nil, err
	}

	hasher := sha256.New224()
	_, err = hasher.Write(data)
	return py.NewHash("sha224", hasher), err
}

func hashlib_sha256(self py.Object, arg py.Object) (py.Object, error) {
	data, err := py.BytesFromObject(arg)
	if err != nil {
		return nil, err
	}
	hasher := sha256.New()
	_, err = hasher.Write(data)
	return py.NewHash("sha256", hasher), err
}

func hashlib_sha384(self py.Object, arg py.Object) (py.Object, error) {
	data, err := py.BytesFromObject(arg)
	if err != nil {
		return nil, err
	}
	hasher := sha512.New384()
	_, err = hasher.Write(data)
	return py.NewHash("sha384", hasher), err
}

func hashlib_sha512(self py.Object, arg py.Object) (py.Object, error) {
	data, err := py.BytesFromObject(arg)
	if err != nil {
		return nil, err
	}
	hasher := sha512.New()
	_, err = hasher.Write(data)
	return py.NewHash("sha512", hasher), err
}

const hashlib_doc = `hashlib module - A common interface to many hash functions.
new(name, data=b'') - returns a new hash object implementing the
                      given hash function; initializing the hash
                      using the given binary data.
Named constructor functions are also available, these are faster
than using new(name):
md5(), sha1(), sha224(), sha256(), sha384(), and sha512()
More algorithms may be available on your platform but the above are guaranteed
to exist.  See the algorithms_guaranteed and algorithms_available attributes
to find out what algorithm names can be passed to new().
NOTE: If you want the adler32 or crc32 hash functions they are available in
the zlib module.
Choose your hash function wisely.  Some have known collision weaknesses.
sha384 and sha512 will be slow on 32 bit platforms.
Hash objects have these methods:
 - update(arg): Update the hash object with the bytes in arg. Repeated calls
                are equivalent to a single call with the concatenation of all
                the arguments.
 - digest():    Return the digest of the bytes passed to the update() method
                so far.
 - hexdigest(): Like digest() except the digest is returned as a unicode
                object of double length, containing only hexadecimal digits.
 - copy():      Return a copy (clone) of the hash object. This can be used to
                efficiently compute the digests of strings that share a common
                initial substring.
For example, to obtain the digest of the string 'Nobody inspects the
spammish repetition':
    >>> import hashlib
    >>> m = hashlib.md5()
    >>> m.update(b"Nobody inspects")
    >>> m.update(b" the spammish repetition")
    >>> m.digest()
    b'\\xbbd\\x9c\\x83\\xdd\\x1e\\xa5\\xc9\\xd9\\xde\\xc9\\xa1\\x8d\\xf0\\xff\\xe9'
More condensed:
    >>> hashlib.sha224(b"Nobody inspects the spammish repetition").hexdigest()
    'a4337bc45a8fc544c03f52dc550cd6e1e87021bc896588bd79e901e2'
`

// Initialise the module
func init() {
	methods := []*py.Method{
		py.MustNewMethod("new", hashlib_new, 0, hashlib_new_doc),
		py.MustNewMethod("md5", hashlib_md5, 0, "Returns a md5 hash object; optionally initialized with a string"),
		py.MustNewMethod("sha1", hashlib_sha1, 0, "Returns a sha1 hash object; optionally initialized with a string"),
		py.MustNewMethod("sha224", hashlib_sha224, 0, "Returns a sha224 hash object; optionally initialized with a string"),
		py.MustNewMethod("sha256", hashlib_sha256, 0, "Returns a sha256 hash object; optionally initialized with a string"),
		py.MustNewMethod("sha384", hashlib_sha384, 0, "Returns a sha384 hash object; optionally initialized with a string"),
		py.MustNewMethod("sha512", hashlib_sha512, 0, "Returns a sha512 hash object; optionally initialized with a string"),
	}
	py.NewModule("hashlib", hashlib_doc, methods, nil)
}
