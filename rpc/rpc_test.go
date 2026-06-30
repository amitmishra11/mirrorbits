// Copyright (c) 2014-2019 Ludovic Fauvet
// Licensed under the MIT license

package rpc

import (
	"sort"
	"testing"
)

func names(mirrors []*MirrorID) []string {
	var out []string
	for _, m := range mirrors {
		out = append(out, m.Name)
	}
	sort.Strings(out)
	return out
}

func TestMatchMirrorsByPatternExactMatchAmongSubstrings(t *testing.T) {
	// Regression test for https://github.com/etix/mirrorbits/issues/134
	// A mirror whose ID is a substring of other mirror names could never
	// be matched on its own, because every other mirror containing it as
	// a substring would also show up as a match.
	mirrors := map[int]string{
		1: "fcix.net",
		2: "mirror.fcix.net",
		3: "paducahix.mm.fcix.net",
		4: "forksystems.mm.fcix.net",
	}

	got := matchMirrorsByPattern(mirrors, "fcix.net")
	want := []string{"fcix.net"}

	if gotNames := names(got); len(gotNames) != len(want) || gotNames[0] != want[0] {
		t.Fatalf("matchMirrorsByPattern(%q) = %v, want %v", "fcix.net", gotNames, want)
	}
}

func TestMatchMirrorsByPatternExactMatchIsCaseInsensitive(t *testing.T) {
	mirrors := map[int]string{
		1: "FCIX.net",
		2: "mirror.fcix.net",
	}

	got := matchMirrorsByPattern(mirrors, "fcix.net")
	want := []string{"FCIX.net"}

	if gotNames := names(got); len(gotNames) != len(want) || gotNames[0] != want[0] {
		t.Fatalf("matchMirrorsByPattern(%q) = %v, want %v", "fcix.net", gotNames, want)
	}
}

func TestMatchMirrorsByPatternMultipleSubstringMatches(t *testing.T) {
	mirrors := map[int]string{
		1: "mirror.fcix.net",
		2: "paducahix.mm.fcix.net",
	}

	got := matchMirrorsByPattern(mirrors, "fcix.net")
	want := []string{"mirror.fcix.net", "paducahix.mm.fcix.net"}

	gotNames := names(got)
	if len(gotNames) != len(want) {
		t.Fatalf("matchMirrorsByPattern(%q) = %v, want %v", "fcix.net", gotNames, want)
	}
	for i := range want {
		if gotNames[i] != want[i] {
			t.Fatalf("matchMirrorsByPattern(%q) = %v, want %v", "fcix.net", gotNames, want)
		}
	}
}

func TestMatchMirrorsByPatternNoMatch(t *testing.T) {
	mirrors := map[int]string{
		1: "alpha",
		2: "beta",
	}

	got := matchMirrorsByPattern(mirrors, "gamma")
	if len(got) != 0 {
		t.Fatalf("matchMirrorsByPattern(%q) = %v, want empty", "gamma", names(got))
	}
}

func TestMatchMirrorsByPatternSingleSubstringMatch(t *testing.T) {
	mirrors := map[int]string{
		1: "mirror.example.com",
		2: "other.example.org",
	}

	got := matchMirrorsByPattern(mirrors, "example.com")
	want := []string{"mirror.example.com"}

	if gotNames := names(got); len(gotNames) != len(want) || gotNames[0] != want[0] {
		t.Fatalf("matchMirrorsByPattern(%q) = %v, want %v", "example.com", gotNames, want)
	}
}
