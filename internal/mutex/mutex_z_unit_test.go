// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/zhwei820/g.

package mutex_test

import (
	"testing"
	"time"

	"github.com/zhwei820/g/internal/mutex"
	"github.com/zhwei820/g/test/gtest"
)

func TestMutexIsSafe(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		lock := mutex.New()
		t.Assert(lock.IsSafe(), false)

		lock = mutex.New(false)
		t.Assert(lock.IsSafe(), false)

		lock = mutex.New(false, false)
		t.Assert(lock.IsSafe(), false)

		lock = mutex.New(true, false)
		t.Assert(lock.IsSafe(), true)

		lock = mutex.New(true, true)
		t.Assert(lock.IsSafe(), true)

		lock = mutex.New(true)
		t.Assert(lock.IsSafe(), true)
	})
}

func TestSafeMutex(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		safeLock := mutex.New(true)
		array := []int{}

		go func() {
			safeLock.Lock()
			array = append(array, 1)
			time.Sleep(100 * time.Millisecond)
			array = append(array, 1)
			safeLock.Unlock()
		}()
		go func() {
			time.Sleep(10 * time.Millisecond)
			safeLock.Lock()
			array = append(array, 1)
			time.Sleep(200 * time.Millisecond)
			array = append(array, 1)
			safeLock.Unlock()
		}()
		time.Sleep(50 * time.Millisecond)
		t.Assert(len(array), 1)
		time.Sleep(80 * time.Millisecond)
		t.Assert(len(array), 3)
		time.Sleep(100 * time.Millisecond)
		t.Assert(len(array), 3)
		time.Sleep(100 * time.Millisecond)
		t.Assert(len(array), 4)
	})
}

func TestUnsafeMutex(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		unsafeLock := mutex.New()
		array := []int{}

		go func() {
			unsafeLock.Lock()
			array = append(array, 1)
			time.Sleep(100 * time.Millisecond)
			array = append(array, 1)
			unsafeLock.Unlock()
		}()
		go func() {
			time.Sleep(10 * time.Millisecond)
			unsafeLock.Lock()
			array = append(array, 1)
			time.Sleep(200 * time.Millisecond)
			array = append(array, 1)
			unsafeLock.Unlock()
		}()
		time.Sleep(50 * time.Millisecond)
		t.Assert(len(array), 2)
		time.Sleep(100 * time.Millisecond)
		t.Assert(len(array), 3)
		time.Sleep(50 * time.Millisecond)
		t.Assert(len(array), 3)
		time.Sleep(100 * time.Millisecond)
		t.Assert(len(array), 4)
	})
}
