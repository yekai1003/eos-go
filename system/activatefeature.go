package system

import (
	"github.com/yekai1003/eos-go"
)

func NewActivateFeature(featureDigest eos.Checksum256) *eos.Action {
	return &eos.Action{
		Account: AN("eosio"),
		Name:    ActN("activate"),
		Authorization: []eos.PermissionLevel{
			{Actor: AN("eosio"), Permission: PN("active")},
		},
		ActionData: eos.NewActionData(Activate{
			FeatureDigest: featureDigest,
		}),
	}
}

// Activate represents a `activate` action on the `eosio` contract.
type Activate struct {
	FeatureDigest eos.Checksum256 `json:"feature_digest"`
}
