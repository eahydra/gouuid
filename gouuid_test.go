//
// Used to Generate UUID. It's standard by RFC4122
// Current package use VariantRFC4122 as Default Variant
//
// Copyright (C) 2013 by Hydra <xxeaglenet@gmail.com>
//
package gouuid

import (
	"testing"
)

func TestTimestampUUID(t *testing.T) {
	u := NewUUID()
	if Compare(u, u) != 0 {
		t.Fatal("")
	}
	if u.Variant() != VariantRFC4122 {
		t.Fatal("")
	}

	if u.Version() != VerTimeStamp {
		t.Fatal("")
	}

	nu := Parse(u.Byte())
	if Compare(nu, u) != 0 {
		t.Fatal("")
	}
}

func TestRandomUUID(t *testing.T) {
	u := NewUUIDByRandom()
	if Compare(u, u) != 0 {
		t.Fatal("")
	}
	if u.Version() != VerRandom {
		t.Fatal("")
	}
	if u.Variant() != VariantRFC4122 {
		t.Fatal("")
	}
	nu := Parse(u.Byte())
	if Compare(nu, u) != 0 {
		t.Fatal("")
	}
}

func TestMD5UUID(t *testing.T) {
	u := NewUUIDByMd5(NamespaceDNS, "www.widgets.com")
	if u.String() != "{3d813cbb-47fb-32ba-91df-831e1593ac29}" {
		println("md5=", u.String())
		t.Fatal("")
	}

	if Compare(u, u) != 0 {
		t.Fatal("")
	}

	if u.Variant() != VariantRFC4122 {
		t.Fatal("")
	}

	if u.Version() != VerNameBasedMD5 {
		t.Fatal("")
	}

	nu := Parse(u.Byte())
	if nu.String() != "{3d813cbb-47fb-32ba-91df-831e1593ac29}" {
		t.Fatal("")
	}
	if Compare(nu, u) != 0 {
		t.Fatal("")
	}
}

func TestSHA1UUID(t *testing.T) {
	u := NewUUIDBySHA1(NamespaceURL, "www.golang.org")
	if u.String() != "{0899ebe6-1d6f-57eb-9da9-d0b30b620430}" {
		t.Fatal("")
	}
	if Compare(u, u) != 0 {
		t.Fatal("")
	}

	if u.Variant() != VariantRFC4122 {
		t.Fatal("")
	}

	if u.Version() != VerNameBasedSHA1 {
		t.Fatal("sha1-ver=", u.Version())
	}

	nu := Parse(u.Byte())
	if nu.String() != "{0899ebe6-1d6f-57eb-9da9-d0b30b620430}" {
		t.Fatal("")
	}
	if Compare(nu, u) != 0 {
		t.Fatal("")
	}
}

func BenchmarkTimestamp(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NewUUIDByTime()
	}
}

func BenchmarkRandom(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NewUUIDByRandom()
	}
}

func BenchmarkMD5(b *testing.B) {
	b.ResetTimer()
	s := "www.widgets.com"
	for i := 0; i < b.N; i++ {
		NewUUIDByMd5(NamespaceDNS, s)
	}
}

func BenchmarkSHA1(b *testing.B) {
	b.ResetTimer()
	s := "www.golang.org"
	for i := 0; i < b.N; i++ {
		NewUUIDBySHA1(NamespaceURL, s)
	}
}
