// Copyright 2016 fatedier, fatedier@gmail.com
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

package pcrypto

import (
	"fmt"
	"testing"
)

func TestEncrypt(t *testing.T) {
	pp := new(Pcrypto)
	pp.Init([]byte("Hana"))
	res, err := pp.Encrypt([]byte("Just One Test!"))
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("[%x]\n", res)
}

func TestDecrypt(t *testing.T) {
	pp := new(Pcrypto)
	pp.Init([]byte("Hana"))
	res, err := pp.Encrypt([]byte("Just One Test!"))
	if err != nil {
		t.Fatal(err)
	}

	res, err = pp.Decrypt(res)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("[%s]\n", string(res))
}
