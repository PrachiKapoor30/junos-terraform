// Go NETCONF Client
//
// Copyright (c) 2013-2018, Juniper Networks, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package netconf

import (
	"bytes"
	"crypto/rand"
	"encoding/xml"
	"fmt"
	"io"
)

// RPCMessage represents an RPC Message to be sent.
type RPCMessage struct {
	MessageID string
	Methods   []RPCMethod
}

// NewRPCMessage generates a new RPC Message structure with the provided methods
func NewRPCMessage(methods []RPCMethod) *RPCMessage {
	return &RPCMessage{
		MessageID: msgID(),
		Methods:   methods,
	}
}

// MarshalXML marshals the NETCONF XML data
func (m *RPCMessage) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var buf bytes.Buffer
	for _, method := range m.Methods {
		buf.WriteString(method.MarshalMethod())
	}

	data := struct {
		MessageID string `xml:"message-id,attr"`
		Xmlns     string `xml:"xmlns,attr"`
		Methods   []byte `xml:",innerxml"`
	}{
		m.MessageID,
		"urn:ietf:params:xml:ns:netconf:base:1.0",
		buf.Bytes(),
	}

	// Wrap the raw XML (data) into <rpc>...</rpc> tags
	start.Name.Local = "rpc"
	return e.EncodeElement(data, start)
}

// RPCReply defines a reply to a RPC request
type RPCReplyBody struct {
	Errors   []RPCError `xml:"rpc-error,omitempty"`
	Data     string     `xml:",innerxml"`
	Ok       bool       `xml:"ok,omitempty"`
	RawReply string     `xml:"-"`
}

// RPCReply defines a reply to a RPC request
type RPCReply struct {
	RPCReplyBody
	XMLName xml.Name `xml:"rpc-reply"`
}

// RPCReply defines a reply to a RPC request
type RPCReplyCommitResults struct {
	XMLName       xml.Name     `xml:"rpc-reply"`
	CommitResults RPCReplyBody `xml:"commit-results,omitempty"`
}

// RPCReply defines a reply to a RPC request
type RPCReplyLoadConfigurationResults struct {
	XMLName                  xml.Name     `xml:"rpc-reply"`
	LoadConfigurationResults RPCReplyBody `xml:"load-configuration-results,omitempty"`
}

// NewRPCReply creates a new RPC Reply
func NewRPCReply(rawXML []byte, ErrOnWarning bool) (*RPCReply, error) {
	/*
		f, _ := os.OpenFile("/var/tmp/rpc-reply.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		fmt.Fprintf(f, "RPC reply (%s)\n", string(rawXML))
		f.Close()
	*/
	reply := &RPCReply{}
	reply.RawReply = string(rawXML)

	if err := xml.Unmarshal(rawXML, reply); err != nil {
		return nil, err
	}

	replyCommitResults := &RPCReplyCommitResults{}
	commitResults := RPCReplyBody{}
	commitResults.RawReply = string(rawXML)
	replyCommitResults.CommitResults = commitResults

	if err := xml.Unmarshal(rawXML, replyCommitResults); err != nil {
		return nil, err
	}

	replyLoadConfigurationResults := &RPCReplyLoadConfigurationResults{}
	loadConfigurationResults := RPCReplyBody{}
	loadConfigurationResults.RawReply = string(rawXML)
	replyLoadConfigurationResults.LoadConfigurationResults = loadConfigurationResults

	if err := xml.Unmarshal(rawXML, replyLoadConfigurationResults); err != nil {
		return nil, err
	}
	/*
		f, _ = os.OpenFile("/var/tmp/rpc-reply.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		fmt.Fprintf(f, "RPC unmarshal complete (%+v)\n", replyLoadConfigurationResults)
		f.Close()
	*/

	if reply.Errors != nil {
		for _, rpcErr := range reply.Errors {
			if rpcErr.Severity == "error" || ErrOnWarning {
				return reply, &rpcErr
			}
		}
	}

	if replyCommitResults.CommitResults.Errors != nil {
		for _, rpcErr := range replyCommitResults.CommitResults.Errors {
			if rpcErr.Severity == "error" || ErrOnWarning {
				reply.Errors = replyCommitResults.CommitResults.Errors
				reply.Data = replyCommitResults.CommitResults.Data
				reply.Ok = replyCommitResults.CommitResults.Ok
				return reply, &rpcErr
			}
		}
	}

	if replyLoadConfigurationResults.LoadConfigurationResults.Errors != nil {
		for _, rpcErr := range replyLoadConfigurationResults.LoadConfigurationResults.Errors {
			if rpcErr.Severity == "error" || ErrOnWarning {
				reply.Errors = replyLoadConfigurationResults.LoadConfigurationResults.Errors
				reply.Data = replyLoadConfigurationResults.LoadConfigurationResults.Data
				reply.Ok = replyLoadConfigurationResults.LoadConfigurationResults.Ok
				return reply, &rpcErr
			}
		}
	}

	return reply, nil
}

// RPCError defines an error reply to a RPC request
type RPCError struct {
	Type     string `xml:"error-type,omitempty"`
	Tag      string `xml:"error-tag,omitempty"`
	Severity string `xml:"error-severity,omitempty"`
	Path     string `xml:"error-path,omitempty"`
	Message  string `xml:"error-message,omitempty"`
	Info     string `xml:",innerxml"`
}

// Error generates a string representation of the provided RPC error
func (re *RPCError) Error() string {
	return fmt.Sprintf("netconf rpc [%s] '%s'", re.Severity, re.Message)
}

// RPCMethod defines the interface for creating an RPC method.
type RPCMethod interface {
	MarshalMethod() string
}

// RawMethod defines how a raw text request will be responded to
type RawMethod string

// MarshalMethod converts the method's output into a string
func (r RawMethod) MarshalMethod() string {
	return string(r)
}

// MethodLock files a NETCONF lock target request with the remote host
func MethodLock(target string) RawMethod {
	return RawMethod(fmt.Sprintf("<lock><target><%s/></target></lock>", target))
}

// MethodUnlock files a NETCONF unlock target request with the remote host
func MethodUnlock(target string) RawMethod {
	return RawMethod(fmt.Sprintf("<unlock><target><%s/></target></unlock>", target))
}

// MethodGetConfig files a NETCONF get-config source request with the remote host
func MethodGetConfig(source string) RawMethod {
	return RawMethod(fmt.Sprintf("<get-config><source><%s/></source></get-config>", source))
}

var msgID = uuid

// uuid generates a "good enough" uuid without adding external dependencies
func uuid() string {
	b := make([]byte, 16)
	io.ReadFull(rand.Reader, b)
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
