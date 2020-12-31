package types

import (
	"fmt"
)

// GenesisState - all nameservice state that must be provided at genesis
type GenesisState struct {
	// TODO: Fill out what is needed by the module for genesis
	WhoisRecords []Whois `json:"whois_records"`
}

// NewGenesisState creates a new GenesisState object
func NewGenesisState( /* TODO: Fill out with what is needed for genesis state */ ) GenesisState {
	return GenesisState{
		// TODO: Fill out according to your genesis state
		WhoisRecords: nil,
	}
}

// DefaultGenesisState - default GenesisState used by Cosmos Hub
func DefaultGenesisState() GenesisState {
	return GenesisState{
		WhoisRecords: []Whois{},
	}
}

// ValidateGenesis validates the nameservice genesis parameters
func ValidateGenesis(data GenesisState) error {
	// TODO: Create a sanity check to make sure the state conforms to the modules needs
	for _, record := range data.WhoisRecords {
		if record.Owner == nil {
			return fmt.Errorf("invalid WhoisRecord: Owner: %s. Error: Missing Owner", record.Owner)
		}
		if record.Value == "" {
			return fmt.Errorf("invalid WhoisRecord: Value: %s. Error: Missing Value", record.Value)
		}
		if record.Price == nil {
			return fmt.Errorf("invalid WhoisRecord: Price: %s. Error: Missing Price", record.Price)
		}
	}
	return nil
}

//// InitGenesis initialize default parameters
//// and the keeper's address to pubkey map
//func InitGenesis(ctx sdk.Context, keeper keeper.Keeper, data GenesisState) {
//	for _, record := range data.WhoisRecords {
//		keeper.SetWhois(ctx, record.Value, record)
//	}
//}
//
//// ExportGenesis writes the current store values
//// to a genesis file, which can be imported again
//// with InitGenesis
//func ExportGenesis(ctx sdk.Context, k keeper.Keeper) GenesisState {
//	var records []Whois
//	iterator := k.GetNamesIterator(ctx)
//	for ; iterator.Valid(); iterator.Next() {
//
//		name := string(iterator.Key())
//		whois, _ := k.GetWhois(ctx, name)
//		records = append(records, whois)
//
//	}
//	return GenesisState{WhoisRecords: records}
//}
