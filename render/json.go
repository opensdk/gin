// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package render

import (
	"github.com/opentoolkit/gotools/json"
	"net/http"
)

var LowerFirstChar = true

type (
	JSON struct {
		Data interface{}
	}

	IndentedJSON struct {
		Data interface{}
	}
)

var jsonContentType = []string{"application/json; charset=utf-8"}

func (r JSON) Render(w http.ResponseWriter) error {
	return WriteJSON(w, r.Data)
}

func (r IndentedJSON) Render(w http.ResponseWriter) error {
	writeContentType(w, jsonContentType)
	jsonBytes, err := jsonutils.MarshalIndent(r.Data, "", "    ", lowerFirstChar)
	if err != nil {
		return err
	}
	w.Write(jsonBytes)
	return nil
}

func WriteJSON(w http.ResponseWriter, obj interface{}) error {
	writeContentType(w, jsonContentType)
	return jsonutils.NewEncoder(w).Encode(obj)
}
