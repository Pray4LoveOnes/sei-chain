#!/usr/bin/env bash
# Usage: ./x402.sh receipts.json

jq -r '.[] | select(.status=="owed" or .status=="pending" or (.settled==false))
       | "\(.txid)\t\(.amount)\t\(.currency)\t\(.memo_private // "")"' "$1" \
| awk -F'\t' 'BEGIN{printf "%-66s %-12s %-6s %-42s\n","TXID","AMOUNT","CCY","PRIVATE"}
              {printf "%-66s %-12s %-6s %-42s\n",$1,$2,$3,$4; sum+=$2}
              END{print "\nTOTAL OWED: " sum " " $3}'
