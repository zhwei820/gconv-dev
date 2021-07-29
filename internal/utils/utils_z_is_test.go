// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/zhwei820/gconv.

package utils_test

import (
	"testing"

	"github.com/zhwei820/gconv/internal/utils"
	"github.com/zhwei820/gconv/test/gtest"
)

func TestVar_IsNil(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(utils.IsNil(0), false)
		t.Assert(utils.IsNil(nil), true)
		// t.Assert(utils.IsNil(map[string]interface{}{}), false)
		// t.Assert(utils.IsNil([]interface{}{}), false)
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(utils.IsNil(1), false)
		t.Assert(utils.IsNil(0.1), false)
		// t.Assert(utils.IsNil(map[string]interface{}{"k": "v"}), false)
		// t.Assert(utils.IsNil([]interface{}{0}), false)
	})
}

func TestVar_IsEmpty(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(utils.IsEmpty(0), true)
		t.Assert(utils.IsEmpty(nil), true)
		// t.Assert(utils.IsEmpty(map[string]interface{}{}), true)
		// t.Assert(utils.IsEmpty([]interface{}{}), true)
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(utils.IsEmpty(1), false)
		t.Assert(utils.IsEmpty(0.1), false)
		// t.Assert(utils.IsEmpty(map[string]interface{}{"k": "v"}), false)
		// t.Assert(utils.IsEmpty([]interface{}{0}), false)
	})
}

func TestVar_IsInt(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(utils.IsInt(0), true)
		t.Assert(utils.IsInt(nil), false)
		// t.Assert(utils.IsInt(map[string]interface{}{}), false)
		// t.Assert(utils.IsInt([]interface{}{}), false)
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(utils.IsInt(1), true)
		t.Assert(utils.IsInt(-1), true)
		t.Assert(utils.IsInt(0.1), false)
		// t.Assert(utils.IsInt(map[string]interface{}{"k": "v"}), false)
		// t.Assert(utils.IsInt([]interface{}{0}), false)
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(utils.IsInt(int8(1)), true)
		t.Assert(utils.IsInt(uint8(1)), false)
	})
}

func TestVar_IsUint(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(utils.IsUint(0), false)
		t.Assert(utils.IsUint(nil), false)
		// t.Assert(utils.IsUint(map[string]interface{}{}), false)
		// t.Assert(utils.IsUint([]interface{}{}), false)
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(utils.IsUint(1), false)
		t.Assert(utils.IsUint(-1), false)
		t.Assert(utils.IsUint(0.1), false)
		// t.Assert(utils.IsUint(map[string]interface{}{"k": "v"}), false)
		// t.Assert(utils.IsUint([]interface{}{0}), false)
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(utils.IsUint(int8(1)), false)
		t.Assert(utils.IsUint(uint8(1)), true)
	})
}

func TestVar_IsFloat(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(utils.IsFloat(0), false)
		t.Assert(utils.IsFloat(nil), false)
		// t.Assert(utils.IsFloat(map[string]interface{}{}), false)
		// t.Assert(utils.IsFloat([]interface{}{}), false)
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(utils.IsFloat(1), false)
		t.Assert(utils.IsFloat(-1), false)
		t.Assert(utils.IsFloat(0.1), true)
		t.Assert(utils.IsFloat(float64(1)), true)
		// t.Assert(utils.IsFloat(map[string]interface{}{"k": "v"}), false)
		// t.Assert(utils.IsFloat([]interface{}{0}), false)
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(utils.IsFloat(int8(1)), false)
		t.Assert(utils.IsFloat(uint8(1)), false)
	})
}

func TestVar_IsSlice(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(utils.IsSlice(0), false)
		t.Assert(utils.IsSlice(nil), false)
		// t.Assert(utils.IsSlice(map[string]interface{}{}), false)
		// t.Assert(utils.IsSlice([]interface{}{}), true)
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(utils.IsSlice(1), false)
		t.Assert(utils.IsSlice(-1), false)
		t.Assert(utils.IsSlice(0.1), false)
		t.Assert(utils.IsSlice(float64(1)), false)
		// t.Assert(utils.IsSlice(map[string]interface{}{"k": "v"}), false)
		// t.Assert(utils.IsSlice([]interface{}{0}), true)
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(utils.IsSlice(int8(1)), false)
		t.Assert(utils.IsSlice(uint8(1)), false)
	})
}

func TestVar_IsMap(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(utils.IsMap(0), false)
		t.Assert(utils.IsMap(nil), false)
		// t.Assert(utils.IsMap(map[string]interface{}{}), true)
		// t.Assert(utils.IsMap([]interface{}{}), false)
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(utils.IsMap(1), false)
		t.Assert(utils.IsMap(-1), false)
		t.Assert(utils.IsMap(0.1), false)
		t.Assert(utils.IsMap(float64(1)), false)
		// t.Assert(utils.IsMap(map[string]interface{}{"k": "v"}), true)
		// t.Assert(utils.IsMap([]interface{}{0}), false)
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(utils.IsMap(int8(1)), false)
		t.Assert(utils.IsMap(uint8(1)), false)
	})
}

func TestVar_IsStruct(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(utils.IsStruct(0), false)
		t.Assert(utils.IsStruct(nil), false)
		// t.Assert(utils.IsStruct(map[string]interface{}{}), false)
		// t.Assert(utils.IsStruct([]interface{}{}), false)
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(utils.IsStruct(1), false)
		t.Assert(utils.IsStruct(-1), false)
		t.Assert(utils.IsStruct(0.1), false)
		t.Assert(utils.IsStruct(float64(1)), false)
		// t.Assert(utils.IsStruct(map[string]interface{}{"k": "v"}), false)
		// t.Assert(utils.IsStruct([]interface{}{0}), false)
	})
	gtest.C(t, func(t *gtest.T) {
		a := &struct {
		}{}
		t.Assert(utils.IsStruct(a), true)
		t.Assert(utils.IsStruct(*a), true)
		t.Assert(utils.IsStruct(&a), true)
	})
}
