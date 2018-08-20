// Copyright 2014 izigo Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package utils

import "testing"

func TestMail(t *testing.T) {
	config := `{"username":"dotiendiep@gmail.com","password":"diepdt","host":"smtp.gmail.com","port":587}`
	mail := NewEMail(config)
	if mail.Username != "dotiendiep@gmail.com" {
		t.Fatal("email parse get username error")
	}
	if mail.Password != "diepdt" {
		t.Fatal("email parse get password error")
	}
	if mail.Host != "smtp.gmail.com" {
		t.Fatal("email parse get host error")
	}
	if mail.Port != 587 {
		t.Fatal("email parse get port error")
	}
	mail.To = []string{"xiemengjun@gmail.com"}
	mail.From = "dotiendiep@gmail.com"
	mail.Subject = "hi, just from izigo!"
	mail.Text = "Text Body is, of course, supported!"
	mail.HTML = "<h1>Fancy Html is supported, too!</h1>"
	mail.AttachFile("/Users/diepdt/github/izigo/izigo.go")
	mail.Send()
}