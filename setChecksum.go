package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"

	"github.com/cavaliercoder/grab"
)

func setChecksum(request *grab.Request, hash []byte, deleteOnCheckFail bool, hashType string) {
	switch hashType {
	case "none":
		break
	case "md5":
		request.SetChecksum(
			md5.New(),
			hash,
			deleteOnCheckFail,
		)
		break
	case "sha1":
		request.SetChecksum(
			sha1.New(),
			hash,
			deleteOnCheckFail,
		)
		break
	case "sha256":
		request.SetChecksum(
			sha256.New(),
			hash,
			deleteOnCheckFail,
		)
		break
	case "sha512":
		request.SetChecksum(
			sha512.New(),
			hash,
			deleteOnCheckFail,
		)
		break
	}
}
