package pump

import "github.com/gagliardetto/solana-go"

// BondingCurve represents the on-chain bonding curve account data
// for a pump.fun token. Fields map to the account layout.
type BondingCurve struct {
	VTR, VSR, RTR, RSR, TTS uint64
	Complete                 bool
	Creator                  solana.PublicKey
	IsCashbackCoin           bool
}
