// Copyright 2020 FastWeGo
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

package miniprogram_test

import (
	"fmt"

	"github.com/fastwego/offiaccount"
	"github.com/fastwego/offiaccount/apis/poi/miniprogram"
)

func ExampleGetMerchantCategory() {
	var ctx *offiaccount.OffiAccount

	resp, err := miniprogram.GetMerchantCategory(ctx)

	fmt.Println(resp, err)
}

func ExampleApplyMerchant() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := miniprogram.ApplyMerchant(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetMerchantAuditInfo() {
	var ctx *offiaccount.OffiAccount

	resp, err := miniprogram.GetMerchantAuditInfo(ctx)

	fmt.Println(resp, err)
}

func ExampleModifyMerchant() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := miniprogram.ModifyMerchant(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetDistrict() {
	var ctx *offiaccount.OffiAccount

	resp, err := miniprogram.GetDistrict(ctx)

	fmt.Println(resp, err)
}

func ExampleSearchMapPoi() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := miniprogram.SearchMapPoi(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleCreateMapPoi() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := miniprogram.CreateMapPoi(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleAddStore() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := miniprogram.AddStore(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleUpdateStore() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := miniprogram.UpdateStore(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleCardStorewxaGet() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := miniprogram.CardStorewxaGet(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetStoreInfo() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := miniprogram.GetStoreInfo(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetStoreList() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := miniprogram.GetStoreList(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDelStore() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := miniprogram.DelStore(ctx, payload)

	fmt.Println(resp, err)
}
