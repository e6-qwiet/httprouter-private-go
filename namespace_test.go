// Copyright 2013 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package httprouter

import (
	"testing"
)

func TestCreateNamespace(t *testing.T) {
	mockNamespace := Namespace{
		Name: "/api",
		IsRoot: true,
	}

	root := New("/api")
	if root.Name != mockNamespace.Name{
		t.Error("Namespace's name failed")
	}

	if root.IsRoot != true{
		t.Error("Namespace is not root")
	}
}


func TestAddNamespace(t *testing.T) {
	mockNamespace := Namespace{
		Name: "/api/v1",
		IsRoot: false,
	}

	root := New("/api")
	v1 := root.Use("/v1")

	if v1.Name != mockNamespace.Name{
		t.Error("Sub Namespace's name failed")
	}

	if v1.IsRoot != false{
		t.Error("Namespace should not be root")
	}
}

