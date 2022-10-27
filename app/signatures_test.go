package app

import "testing"

// examples from the circleci docs:
// Body	Secret Key	Signature
// hello world	secret	734cc62f32841568f45715aeb9f4d7891324e6d948e4c6c60c0621cdac48623a
// lalala	another-secret	daa220016c8f29a8b214fbfc3671aeec2145cfb1e6790184ffb38b6d0425fa00
// an-important-request-payload	hunter123	9be2242094a9a8c00c64306f382a7f9d691de910b4a266f67bd314ef18ac49fa
func TestHashes(t *testing.T) {
	if !VerifySignature([]byte("hello world"), []byte("secret"), "734cc62f32841568f45715aeb9f4d7891324e6d948e4c6c60c0621cdac48623a") {
		t.Fatalf(`verify failed: "hello world"`)
	}
	if !VerifySignature([]byte("lalala"), []byte("another-secret"), "daa220016c8f29a8b214fbfc3671aeec2145cfb1e6790184ffb38b6d0425fa00") {
		t.Fatalf(`verify failed: "lalala"`)
	}
	if !VerifySignature([]byte("an-important-request-payload"), []byte("hunter123"), "9be2242094a9a8c00c64306f382a7f9d691de910b4a266f67bd314ef18ac49fa") {
		t.Fatalf(`verify failed: "an-important-request-payload"`)
	}
}
