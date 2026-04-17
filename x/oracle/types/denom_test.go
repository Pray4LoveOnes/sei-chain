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

import "testing"

func TestDenomListContains(t *testing.T) {
	tests := []struct {
		name      string
		denomList DenomList
		denom     string
		want      bool
	}{
		{
			name: "denomination present",
			denomList: DenomList{
				{Name: "USD"},
				{Name: "EUR"},
				{Name: "INR"},
			},
			denom: "EUR",
			want:  true,
		},
		{
			name: "denomination absent",
			denomList: DenomList{
				{Name: "USD"},
				{Name: "EUR"},
				{Name: "INR"},
			},
			denom: "JPY",
			want:  false,
		},
		{
			name:      "empty list",
			denomList: DenomList{},
			denom:     "USD",
			want:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.denomList.Contains(tt.denom); got != tt.want {
				t.Errorf("DenomList.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
