/*
   Copyright (c) 2016 VMware, Inc. All Rights Reserved.
   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package utils

import (
	"encoding/base64"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
}

func TestParseRepository(t *testing.T) {
	repository := "library/ubuntu"
	project, rest := ParseRepository(repository)
	if project != "library" {
		t.Errorf("unexpected project: %s != %s", project, "library")
	}
	if rest != "ubuntu" {
		t.Errorf("unexpected rest: %s != %s", rest, "ubuntu")
	}

	repository = "library/test/ubuntu"
	project, rest = ParseRepository(repository)
	if project != "library/test" {
		t.Errorf("unexpected project: %s != %s", project, "library/test")
	}
	if rest != "ubuntu" {
		t.Errorf("unexpected rest: %s != %s", rest, "ubuntu")
	}

	repository = "ubuntu"
	project, rest = ParseRepository(repository)
	if project != "" {
		t.Errorf("unexpected project: [%s] != [%s]", project, "")
	}

	if rest != "ubuntu" {
		t.Errorf("unexpected rest: %s != %s", rest, "ubuntu")
	}

	repository = ""
	project, rest = ParseRepository(repository)
	if project != "" {
		t.Errorf("unexpected project: [%s] != [%s]", project, "")
	}

	if rest != "" {
		t.Errorf("unexpected rest: [%s] != [%s]", rest, "")
	}
}

func TestReversibleEncrypt(t *testing.T) {
	password := "password"
	key := "1234567890123456"
	encrypted, err := ReversibleEncrypt(password, key)
	if err != nil {
		t.Errorf("Failed to encrypt: %v", err)
	}
	t.Logf("Encrypted password: %s", encrypted)
	if encrypted == password {
		t.Errorf("Encrypted password is identical to the original")
	}
	if !strings.HasPrefix(encrypted, EncryptHeaderV1) {
		t.Errorf("Encrypted password does not have v1 header")
	}
	decrypted, err := ReversibleDecrypt(encrypted, key)
	if err != nil {
		t.Errorf("Failed to decrypt: %v", err)
	}
	if decrypted != password {
		t.Errorf("decrypted password: %s, is not identical to original", decrypted)
	}
	//Test b64 for backward compatibility
	b64password := base64.StdEncoding.EncodeToString([]byte(password))
	decrypted, err = ReversibleDecrypt(b64password, key)
	if err != nil {
		t.Errorf("Failed to decrypt: %v", err)
	}
	if decrypted != password {
		t.Errorf("decrypted password: %s, is not identical to original", decrypted)
	}
}