/*
Copyright 2009 The Go Authors. All rights reserved.

Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

Package gob manages streams of gobs - binary values exchanged between an Encoder (transmitter) and a Decoder (receiver).
A typical use is transporting arguments and results of remote procedure calls (RPCs) such as those provided by package
"net/rpc".

The implementation compiles a custom codec for each data type in the stream and is most efficient when a single Encoder
is used to transmit a stream of values, amortizing the cost of compilation.

Basics

A stream of gobs is self-describing. Each data item in the stream is preceded by a specification of its type, expressed
in terms of a small set of predefined types. Pointers are not transmitted, but the things they point to are transmitted;
that is, the values are flattened. Nil pointers are not permitted, as they have no value. Recursive types work fine, but
recursive values (data with cycles) are problematic. This may change.

To use gobs, create an Encoder and present it with a series of data items as values or addresses that can be
dereferenced to values. The Encoder makes sure all type information is sent before it is needed. At the receive side, a
Decoder retrieves values from the encoded stream and unpacks them into local variables.

Types and Values

The source and destination values/types need not correspond exactly. For structs, fields (identified by name) that are
in the source but absent from the receiving variable will be ignored. Fields that are in the receiving variable but
missing from the transmitted type or value will be ignored in the destination. If a field with the same name is present
in both, their types must be compatible. Both the receiver and transmitter will do all necessary indirection and
dereferencing to convert between gobs and actual Go values. For instance, a gob type that is schematically,

	struct { A, B int }

can be sent from or received into any of these Go types:

	struct { A, B int }	    // the same
	*struct { A, B int }    // extra indirection of the struct
	struct { *A, **B int }  // extra indirection of the fields
	struct { A, B int64 }   // different concrete value type; see below

It may also be received into any of these:

	struct { A, B int }	   // the same
	struct { B, A int }    // ordering doesn't matter; matching is by name
	struct { A, B, C int } // extra field (C) ig
*/
package main

import "fmt"

func main() {
	fmt.Println("Hello", "World!")
}
