package server

import (
	"airdrop/types"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"testing"
)

func TestGetDelegateAttestationRequest(t *testing.T) {
	data := types.ZKPassProofDetail{
		BalanceLE: 1,
		Proof: types.ZKPassProof{
			TaskId:             "b8f72ead91ab426b822106d1a4184cfb",
			PublicFields:       []interface{}{},
			AllocatorAddress:   "0x19a567b3b212a5b35bA0E3B600FbEd5c2eE9083d",
			PublicFieldsHash:   "0xc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc6",
			AllocatorSignature: "0x82958010ff1ef42c09632fdbec8b6f11cb5f87d0369b2063b59aabafbebc1eed32af60d5b27d44f7e91326425cecc68d7b996d71ffba67335a2946b6400b46c01c",
			UHash:              "0x511f42913644f4f70713bcc0191711df8e7dfd9f3b4119e0102cae1bf54c1274",
			ValidatorAddress:   "0xb1C4C1E1Cdd5Cf69E27A3A08C8f51145c2E12C6a",
			ValidatorSignature: "0xb88821c157a775c1522864ac2777b0301ad122c997136ea1d9fb6160d82ff0187da6d5593b9c75eb75b68f89a604d7a3c69c53edbd47c30e6b6ac8513d3f86b61b",
		},
		Recipient:  "0x471543A3bd04486008c8a38c5C00543B73F1769e",
		Expiration: 1723507077,
	}

	signer, err := crypto.HexToECDSA("19dfa30d6165181386c6a706f065bd841bf41fb9457a7f4a7a9b8c5df5d4de89")

	//fmt.Println(crypto.PubkeyToAddress(signer.PublicKey).String())
	if err != nil {
		panic("err")
	}
	bData, err := json.Marshal(data)
	if err != nil {
		panic("err")
	}

	t.Log("jsonb:", string(bData))

	hash := sha256.Sum256(bData)
	t.Log("Hash:", hex.EncodeToString(hash[:]))
	sig, err := crypto.Sign(hash[:], signer)
	if err != nil {
		panic("err")
	}

	t.Log(common.Bytes2Hex(sig))
}
