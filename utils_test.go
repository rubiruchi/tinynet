// Copyright (c) 2017 Che Wei, Lin
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

package tinynet

import (
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
)

func TestUtils_Invalid(t *testing.T) {
	_, err := getAllIPsfromCIDR("1.2.3.4.5./32")
	assert.Error(t, err)
}

func TestInc(t *testing.T) {
	ip, _, _ := net.ParseCIDR("140.113.234.123/30")
	expected := []string{"140.113.234.123", "140.113.234.124", "140.113.234.125"}

	for i, _ := range expected {
		assert.Equal(t, ip.String(), expected[i], "Those two address should be same")
		inc(ip)
	}
}

func TestUtils_Success(t *testing.T) {

	data := []struct {
		desc     string
		cidr     string
		expected []string
	}{
		{"ValidCIDR_29", "140.113.234.123/29", []string{"140.113.234.121", "140.113.234.122", "140.113.234.123", "140.113.234.124", "140.113.234.125", "140.113.234.126"}},
		{"ValidCIDR_30", "140.113.234.123/30", []string{"140.113.234.121", "140.113.234.122"}},
		{"ValidCIDR_31", "140.113.234.123/31", []string{}},
	}
	for _, d := range data {
		t.Run(d.desc, func(t *testing.T) {
			data, err := getAllIPsfromCIDR(d.cidr)
			assert.NoError(t, err)
			assert.Equal(t, len(data), len(d.expected))

			for i, _ := range data {
				assert.Equal(t, data[i], d.expected[i])
			}
		})
	}
}
