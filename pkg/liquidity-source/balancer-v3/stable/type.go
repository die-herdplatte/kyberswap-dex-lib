package stable

import (
	"math/big"

	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/balancer-v3/shared"
	"github.com/holiman/uint256"
)

type Extra struct {
	HooksConfig                shared.HooksConfig `json:"hooksConfig"`
	StaticSwapFeePercentage    *uint256.Int       `json:"staticSwapFeePercentage"`
	AggregateSwapFeePercentage *uint256.Int       `json:"aggregateSwapFeePercentage"`
	AmplificationParameter     *uint256.Int       `json:"amplificationParameter"`
	BalancesLiveScaled18       []*uint256.Int     `json:"balancesLiveScaled18"`
	DecimalScalingFactors      []*uint256.Int     `json:"decimalScalingFactors"`
	TokenRates                 []*uint256.Int     `json:"tokenRates"`
	IsVaultPaused              bool               `json:"isVaultPaused"`
	IsPoolPaused               bool               `json:"isPoolPaused"`
	IsPoolInRecoveryMode       bool               `json:"isPoolInRecoveryMode"`
}

type StaticExtra struct {
	Vault             string `json:"vault"`
	DefaultHook       string `json:"defaultHook"`
	IsPoolInitialized bool   `json:"isPoolInitialized"`
}

type AmplificationParameter struct {
	Value      *big.Int
	IsUpdating bool
	Precision  *big.Int
}

type PoolMetaInfo struct {
	Vault       string `json:"vault"`
	BlockNumber uint64 `json:"blockNumber"`
}

type RpcResult struct {
	HooksConfig                shared.HooksConfig
	BalancesRaw                []*big.Int
	BalancesLiveScaled18       []*big.Int
	TokenRates                 []*big.Int
	DecimalScalingFactors      []*big.Int
	AmplificationParameter     *big.Int
	StaticSwapFeePercentage    *big.Int
	AggregateSwapFeePercentage *big.Int
	IsVaultPaused              bool
	IsPoolPaused               bool
	IsPoolInRecoveryMode       bool
	BlockNumber                uint64
}

type SwapInfo struct {
	AggregateFee *big.Int `json:"aggregateFee"`
}
