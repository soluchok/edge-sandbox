// Copyright SecureKey Technologies Inc. All Rights Reserved.
//
// SPDX-License-Identifier: Apache-2.0

module github.com/trustbloc/edge-sandbox/test/bdd

go 1.15

require (
	github.com/cucumber/godog v0.8.1
	github.com/google/go-cmp v0.4.0 // indirect
	github.com/hyperledger/fabric-protos-go v0.0.0-20200707132912-fee30f3ccd23
	github.com/pkg/errors v0.9.1
	github.com/spf13/viper v1.4.0
	github.com/trustbloc/edge-core v0.1.5-0.20200902222811-9a73214c780d
	github.com/trustbloc/fabric-peer-test-common v0.1.5
)

replace github.com/hyperledger/fabric-protos-go => github.com/trustbloc/fabric-protos-go-ext v0.1.4-0.20200626180529-18936b36feca
