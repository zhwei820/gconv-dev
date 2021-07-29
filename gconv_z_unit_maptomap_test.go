// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/zhwei820/gconv.

package gconv_test

import (
	"testing"

	"github.com/zhwei820/gconv"
	"github.com/zhwei820/gconv/test/gtest"
)

func Test_MapToMap1(t *testing.T) {
	// map[int]int -> map[string]string
	// empty original map.
	gtest.C(t, func(t *gtest.T) {
		m1 := map[int]int{}
		m2 := map[string]string{}
		t.Assert(gconv.MapToMap(m1, &m2), nil)
		t.Assert(len(m1), len(m2))
	})
	// map[int]int -> map[string]string
	gtest.C(t, func(t *gtest.T) {
		m1 := map[int]int{
			1: 100,
			2: 200,
		}
		m2 := map[string]string{}
		t.Assert(gconv.MapToMap(m1, &m2), nil)
		t.Assert(m2["1"], m1[1])
		t.Assert(m2["2"], m1[2])
	})
	// map[string]interface{} -> map[string]string
	gtest.C(t, func(t *gtest.T) {
		m1 := map[string]interface{}{
			"k1": "v1",
			"k2": "v2",
		}
		m2 := map[string]string{}
		t.Assert(gconv.MapToMap(m1, &m2), nil)
		t.Assert(m2["k1"], m1["k1"])
		t.Assert(m2["k2"], m1["k2"])
	})
	// map[string]string -> map[string]interface{}
	gtest.C(t, func(t *gtest.T) {
		m1 := map[string]string{
			"k1": "v1",
			"k2": "v2",
		}
		m2 := map[string]interface{}{}
		t.Assert(gconv.MapToMap(m1, &m2), nil)
		t.Assert(m2["k1"], m1["k1"])
		t.Assert(m2["k2"], m1["k2"])
	})
	// map[string]interface{} -> map[interface{}]interface{}
	gtest.C(t, func(t *gtest.T) {
		m1 := map[string]string{
			"k1": "v1",
			"k2": "v2",
		}
		m2 := map[interface{}]interface{}{}
		t.Assert(gconv.MapToMap(m1, &m2), nil)
		t.Assert(m2["k1"], m1["k1"])
		t.Assert(m2["k2"], m1["k2"])
	})
}

func Test_MapToMap2(t *testing.T) {
	type User struct {
		Id   int
		Name string
	}
	params := map[string]interface{}{
		"key": map[string]interface{}{
			"id":   1,
			"name": "john",
		},
	}
	gtest.C(t, func(t *gtest.T) {
		m := make(map[string]User)
		err := gconv.MapToMap(params, &m)
		t.Assert(err, nil)
		t.Assert(len(m), 1)
		t.Assert(m["key"].Id, 1)
		t.Assert(m["key"].Name, "john")
	})
	gtest.C(t, func(t *gtest.T) {
		m := (map[string]User)(nil)
		err := gconv.MapToMap(params, &m)
		t.Assert(err, nil)
		t.Assert(len(m), 1)
		t.Assert(m["key"].Id, 1)
		t.Assert(m["key"].Name, "john")
	})
	gtest.C(t, func(t *gtest.T) {
		m := make(map[string]*User)
		err := gconv.MapToMap(params, &m)
		t.Assert(err, nil)
		t.Assert(len(m), 1)
		t.Assert(m["key"].Id, 1)
		t.Assert(m["key"].Name, "john")
	})
	gtest.C(t, func(t *gtest.T) {
		m := (map[string]*User)(nil)
		err := gconv.MapToMap(params, &m)
		t.Assert(err, nil)
		t.Assert(len(m), 1)
		t.Assert(m["key"].Id, 1)
		t.Assert(m["key"].Name, "john")
	})
}

func Test_MapToMapDeep(t *testing.T) {
	type Ids struct {
		Id  int
		Uid int
	}
	type Base struct {
		Ids
		Time string
	}
	type User struct {
		Base
		Name string
	}
	params := map[string]interface{}{
		"key": map[string]interface{}{
			"id":   1,
			"name": "john",
		},
	}
	gtest.C(t, func(t *gtest.T) {
		m := (map[string]*User)(nil)
		err := gconv.MapToMap(params, &m)
		t.Assert(err, nil)
		t.Assert(len(m), 1)
		t.Assert(m["key"].Id, 1)
		t.Assert(m["key"].Name, "john")
	})
}

func Test_MapToMaps(t *testing.T) {
	params := []interface{}{
		map[string]interface{}{"id": 1, "name": "john"},
		map[string]interface{}{"id": 2, "name": "smith"},
	}
	gtest.C(t, func(t *gtest.T) {
		var s []map[string]interface{}
		err := gconv.MapToMaps(params, &s)
		t.AssertNil(err)
		t.Assert(len(s), 2)
		t.Assert(s, params)
	})
	gtest.C(t, func(t *gtest.T) {
		var s []*map[string]interface{}
		err := gconv.MapToMaps(params, &s)
		t.AssertNil(err)
		t.Assert(len(s), 2)
		t.Assert(s, params)
	})
}

func Test_MapToMaps_StructParams(t *testing.T) {
	type User struct {
		Id   int
		Name string
	}
	params := []interface{}{
		User{1, "name1"},
		User{2, "name2"},
	}
	gtest.C(t, func(t *gtest.T) {
		var s []map[string]interface{}
		err := gconv.MapToMaps(params, &s)
		t.AssertNil(err)
		t.Assert(len(s), 2)
	})
	gtest.C(t, func(t *gtest.T) {
		var s []*map[string]interface{}
		err := gconv.MapToMaps(params, &s)
		t.AssertNil(err)
		t.Assert(len(s), 2)
	})
}
