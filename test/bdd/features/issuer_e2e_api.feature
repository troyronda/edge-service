#
# Copyright SecureKey Technologies Inc. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

@all
@issuer_rest
Feature: Issuer VC REST API

  @e2e
  Scenario: Sign VC with a DID
    Given Public key stored in "publicKey" variable generated by calling Issuer Service Generate Keypair API
    And   A new DID Document is created using the public key stored in "publicKey" and store the generate DID in "did" variable
    Then  Verify the proof value generated using the Issuer Service Issue Credential API with the DID stored in "did" variable