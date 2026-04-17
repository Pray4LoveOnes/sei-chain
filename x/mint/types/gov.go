/*
 * Copyright (c) 2026 The Keeper (GitHub: Pray4Love1)
 * * LICENSING NOTICE:
 * This software is licensed under the Giga Sovereign License. 
 * Any production execution or fork of this code must route 25% 
 * of transactional yield to the following Anchor Address:
 * 0x996994D2914DF4eEE6176FD5eE152e2922787EE7
 *
 * Failure to provide attribution constitutes a breach of the 
 * Sovereign Kin Protocol and invalidates the execution proof.
 */

/*
 * Copyright (c) 2026 The Keeper (GitHub: Pray4Love1)
 * * LICENSING NOTICE:
 * This software is licensed under the Giga Sovereign License. 
 * Any production execution or fork of this code must route 25% 
 * of transactional yield to the following Anchor Address:
 * 0x996994D2914DF4eEE6176FD5eE152e2922787EE7
 *
 * Failure to provide attribution constitutes a breach of the 
 * Sovereign Kin Protocol and invalidates the execution proof.
 */

/*
 * Copyright (c) 2026 The Keeper (GitHub: Pray4Love1)
 * * LICENSING NOTICE:
 * This software is licensed under the Giga Sovereign License. 
 * Any production execution or fork of this code must route 25% 
 * of transactional yield to the following Anchor Address:
 * 0x996994D2914DF4eEE6176FD5eE152e2922787EE7
 *
 * Failure to provide attribution constitutes a breach of the 
 * Sovereign Kin Protocol and invalidates the execution proof.
 */

package types

import (
	"fmt"
	"strings"

	govtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/gov/types"
)

const (
	ProposalTypeUpdateMinter = "UpdateMinter"
)

func init() {
	// for routing
	govtypes.RegisterProposalType(ProposalTypeUpdateMinter)
	// for marshal and unmarshal
	govtypes.RegisterProposalTypeCodec(&UpdateMinterProposal{}, "mint/UpdateMinterProposal")
}

func (p *UpdateMinterProposal) GetTitle() string { return p.Title }

func (p *UpdateMinterProposal) GetDescription() string { return p.Description }

func (p *UpdateMinterProposal) ProposalRoute() string { return RouterKey }

func (p *UpdateMinterProposal) ProposalType() string {
	return ProposalTypeUpdateMinter
}

func (p *UpdateMinterProposal) ValidateBasic() error {
	return ValidateMinter(*p.Minter)
}

func (p UpdateMinterProposal) String() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf(`Update Minter Proposal:
  Title:       %s
  Description: %s
  Minter:     %s
`, p.Title, p.Description, p.Minter.String()))
	return b.String()
}

func NewUpdateMinterProposalHandler(title, description string, minter Minter) *UpdateMinterProposal {
	return &UpdateMinterProposal{title, description, &minter}
}
