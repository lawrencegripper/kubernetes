package main

// Copyright 2017 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import (
	"path/filepath"
	"testing"

	"github.com/marstr/collection"
)

func Test_versionle(t *testing.T) {
	const dateWithAlpha, dateWithBeta = "2016-02-01-alpha", "2016-02-01-beta"
	const semVer1dot2, semVer1dot3 = "2018-03-03-1.2", "2018-03-03-1.3"

	testCases := []struct {
		left  string
		right string
		want  bool
	}{
		{"2017-12-01", "2018-03-04", true},
		{"2018-03-04", "2017-12-01", false},
		{semVer1dot2, semVer1dot3, true},
		{semVer1dot3, semVer1dot2, false},
		{semVer1dot2, semVer1dot2, true},
		{dateWithAlpha, dateWithBeta, true},
		{dateWithBeta, dateWithAlpha, false},
		{"2016-04-03-preview", "2016-04-03", true},
		{"2016-04-03", "2016-04-03-preview", false},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			t.Logf("\n Left: %s\nRight: %s", tc.left, tc.right)
			if got, err := versionle(tc.left, tc.right); err != nil {
				t.Error(err)
			} else if got != tc.want {
				t.Logf("\n got: %v\nwant: %v", got, tc.want)
				t.Fail()
			}
		})
	}
}

func BenchmarkLatestStrategy_Enumerate(b *testing.B) {
	subject := LatestStrategy{
		Root: filepath.Join("..", "..", "service"),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.Logf("Enumerated %d packages", collection.CountAll(subject))
	}
}
