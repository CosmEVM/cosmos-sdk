package types

import (
	"math"
	"time"

	"github.com/tendermint/tendermint/crypto/tmhash"
	tmbits "github.com/tendermint/tendermint/proto/libs/bits"
	tmproto "github.com/tendermint/tendermint/proto/types"
	tmtypes "github.com/tendermint/tendermint/types"
	"github.com/tendermint/tendermint/version"
)

// Copied unimported test functions from tmtypes to use them here
func MakeBlockID(hash []byte, partSetSize uint32, partSetHash []byte) tmtypes.BlockID {
	return tmtypes.BlockID{
		Hash: hash,
		PartsHeader: tmtypes.PartSetHeader{
			Total: partSetSize,
			Hash:  partSetHash,
		},
	}
}

// CreateTestHeader creates a mock header for testing only.
func CreateTestHeader(chainID string, height int64, timestamp time.Time, valSet *tmtypes.ValidatorSet, signers []tmtypes.PrivValidator) Header {
	vsetHash := valSet.Hash()
	tmHeader := &tmtypes.Header{
		Version:            version.Consensus{Block: 2, App: 2},
		ChainID:            chainID,
		Height:             height,
		Time:               timestamp,
		LastBlockID:        MakeBlockID(make([]byte, tmhash.Size), math.MaxInt32, make([]byte, tmhash.Size)),
		LastCommitHash:     tmhash.Sum([]byte("last_commit_hash")),
		DataHash:           tmhash.Sum([]byte("data_hash")),
		ValidatorsHash:     vsetHash,
		NextValidatorsHash: vsetHash,
		ConsensusHash:      tmhash.Sum([]byte("consensus_hash")),
		AppHash:            tmhash.Sum([]byte("app_hash")),
		LastResultsHash:    tmhash.Sum([]byte("last_results_hash")),
		EvidenceHash:       tmhash.Sum([]byte("evidence_hash")),
		ProposerAddress:    valSet.Proposer.Address,
	}

	blockID := MakeBlockID(tmHeader.Hash(), 3, tmhash.Sum([]byte("part_set")))
	voteSet := tmtypes.NewVoteSet(chainID, height, 1, tmproto.PrecommitType, valSet)
	commit, err := tmtypes.MakeCommit(blockID, height, 1, voteSet, signers, timestamp)
	if err != nil {
		panic(err)
	}

	commitSigs := make([]tmproto.CommitSig, len(commit.Signatures))

	for i := range commit.Signatures {
		commitSigs[i] = tmproto.CommitSig{
			BlockIdFlag:      commit.Signatures[i].BlockIDFlag,
			ValidatorAddress: commit.Signatures[i].ValidatorAddress,
			Timestamp:        commit.Signatures[i].Timestamp,
			Signature:        commit.Signatures[i].Signature,
		}
	}

	abciBlockID := tmtypes.TM2PB.BlockID(blockID)
	abciHeader := tmtypes.TM2PB.Header(tmHeader)
	bitArray := commit.BitArray()

	signedHeader := tmproto.SignedHeader{
		Header: &abciHeader,
		Commit: &tmproto.Commit{
			Height: commit.Height,
			Round:  int32(commit.Round),
			BlockID: tmproto.BlockID{
				Hash: abciBlockID.Hash,
				PartsHeader: tmproto.PartSetHeader{
					Total: abciBlockID.PartsHeader.Total,
					Hash:  abciBlockID.PartsHeader.Hash,
				},
			},
			Signatures: commitSigs,
			Hash:       commit.Hash(),
			BitArray: &tmbits.BitArray{
				Bits:  int64(bitArray.Bits),
				Elems: bitArray.Elems,
			},
		},
	}

	return Header{
		SignedHeader: signedHeader,
		ValidatorSet: valSet,
	}
}
