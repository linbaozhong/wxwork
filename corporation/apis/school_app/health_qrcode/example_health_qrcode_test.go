// Copyright 2021 FastWeGo
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

package health_qrcode_test

import (
	"fmt"

	"github.com/linbaozhong/wxwork/corporation"
	"github.com/linbaozhong/wxwork/corporation/apis/school_app/health_qrcode"
)

func ExampleGetTeacherCustomizeHealthInfo() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := health_qrcode.GetTeacherCustomizeHealthInfo(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetStudentCustomizeHealthInfo() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := health_qrcode.GetStudentCustomizeHealthInfo(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetHealthQrcode() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := health_qrcode.GetHealthQrcode(ctx, payload)

	fmt.Println(resp, err)
}
