# \LoopringDEXRestfulAPIApi

All URIs are relative to *http://uat2.loopring.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyApiKeyV2**](LoopringDEXRestfulAPIApi.md#ApplyApiKeyV2) | **Post** /api/v2/apiKey | Update user&#39;s ApiKey
[**ApplyApiKeyV3**](LoopringDEXRestfulAPIApi.md#ApplyApiKeyV3) | **Post** /api/v3/apiKey | Update user&#39;s ApiKey
[**BatchCancelOrdersByClientOrderIdV2**](LoopringDEXRestfulAPIApi.md#BatchCancelOrdersByClientOrderIdV2) | **Delete** /api/v2/orders/byClientOrderId | Cancel multiple orders by clientOrderId
[**BatchCancelOrdersByClientOrderIdV3**](LoopringDEXRestfulAPIApi.md#BatchCancelOrdersByClientOrderIdV3) | **Delete** /api/v3/orders/byClientOrderId | 
[**BatchCancelOrdersByOrderHashV2**](LoopringDEXRestfulAPIApi.md#BatchCancelOrdersByOrderHashV2) | **Delete** /api/v2/orders/byHash | Cancel multiple orders by hash
[**BatchCancelOrdersByOrderHashV3**](LoopringDEXRestfulAPIApi.md#BatchCancelOrdersByOrderHashV3) | **Delete** /api/v3/orders/byHash | 
[**BatchSubmitOrderV2**](LoopringDEXRestfulAPIApi.md#BatchSubmitOrderV2) | **Post** /api/v2/batchOrders | Submit multiple orders
[**CanMintV3**](LoopringDEXRestfulAPIApi.md#CanMintV3) | **Get** /api/v3/nft/info/canMint | check if a minter is authorized on a nft token contract
[**CancelOrderV3**](LoopringDEXRestfulAPIApi.md#CancelOrderV3) | **Delete** /api/v3/order | Cancel order
[**CancelOrdersV2**](LoopringDEXRestfulAPIApi.md#CancelOrdersV2) | **Delete** /api/v2/orders | api.CancelOrders.value
[**ClaimLuckyTokenV3**](LoopringDEXRestfulAPIApi.md#ClaimLuckyTokenV3) | **Post** /api/v3/luckyToken/claimLuckyToken | api.claimLuckyToken.value
[**ComputeTokenAddressV3**](LoopringDEXRestfulAPIApi.md#ComputeTokenAddressV3) | **Get** /api/v3/nft/info/computeTokenAddress | 
[**ConfirmHebaoPayV3Hebao**](LoopringDEXRestfulAPIApi.md#ConfirmHebaoPayV3Hebao) | **Post** /api/v3/hebao/confirmPay | 
[**CrawlPriceFromBinanceV2**](LoopringDEXRestfulAPIApi.md#CrawlPriceFromBinanceV2) | **Get** /api/v2/crawlPriceFromBinance | 
[**CrawlSiteV3**](LoopringDEXRestfulAPIApi.md#CrawlSiteV3) | **Get** /api/v3/crawl | 
[**DeployTokenAddressV3**](LoopringDEXRestfulAPIApi.md#DeployTokenAddressV3) | **Post** /api/v3/nft/deployTokenAddress | 
[**GetAccountServicesV3**](LoopringDEXRestfulAPIApi.md#GetAccountServicesV3) | **Get** /api/v3/spi/getAccountServices | 
[**GetAccountV2**](LoopringDEXRestfulAPIApi.md#GetAccountV2) | **Get** /api/v2/account | Query user information
[**GetAccountV3**](LoopringDEXRestfulAPIApi.md#GetAccountV3) | **Get** /api/v3/account | Query user information
[**GetActivityRulesV3**](LoopringDEXRestfulAPIApi.md#GetActivityRulesV3) | **Get** /api/v3/sidecar/activityRules | 
[**GetAllowanceV3**](LoopringDEXRestfulAPIApi.md#GetAllowanceV3) | **Get** /api/v3/eth/allowances | Get user token allowance
[**GetAmmAssetsV3**](LoopringDEXRestfulAPIApi.md#GetAmmAssetsV3) | **Get** /api/v3/amm/assets | 
[**GetAmmDepthV2**](LoopringDEXRestfulAPIApi.md#GetAmmDepthV2) | **Get** /api/v2/amm/depth | Get market orderbook
[**GetAmmDepthV3**](LoopringDEXRestfulAPIApi.md#GetAmmDepthV3) | **Get** /api/v3/amm/depth | Get market orderbook
[**GetAmmMarketsV2**](LoopringDEXRestfulAPIApi.md#GetAmmMarketsV2) | **Get** /api/v2/amm/markets | api.getAmmMarkets.value
[**GetAmmPoolBalanceV3**](LoopringDEXRestfulAPIApi.md#GetAmmPoolBalanceV3) | **Get** /api/v3/amm/balance | Get AMM pool balance snapshot
[**GetAmmPoolsBalancesV3**](LoopringDEXRestfulAPIApi.md#GetAmmPoolsBalancesV3) | **Get** /api/v3/amm/balances | api.getAmmPoolsBalances.value
[**GetAmmPoolsV3**](LoopringDEXRestfulAPIApi.md#GetAmmPoolsV3) | **Get** /api/v3/amm/pools | Get AMM pool configurations
[**GetAmmSnapshotV2**](LoopringDEXRestfulAPIApi.md#GetAmmSnapshotV2) | **Get** /api/v2/amm/snapshot | api.getAmmSnapshot.value
[**GetAmmSnapshotsV2**](LoopringDEXRestfulAPIApi.md#GetAmmSnapshotsV2) | **Get** /api/v2/amm/snapshots | api.getAmmSnapshots.value
[**GetAmmTradesV2**](LoopringDEXRestfulAPIApi.md#GetAmmTradesV2) | **Get** /api/v2/amm/trades | get AMM pool trade transactions
[**GetAmmTradesV3**](LoopringDEXRestfulAPIApi.md#GetAmmTradesV3) | **Get** /api/v3/amm/trades | get AMM pool trade transactions
[**GetAmmTransactionsV2**](LoopringDEXRestfulAPIApi.md#GetAmmTransactionsV2) | **Get** /api/v2/amm/transactions | get AMM involved join/exit transactions
[**GetAmmTransactionsV3**](LoopringDEXRestfulAPIApi.md#GetAmmTransactionsV3) | **Get** /api/v3/amm/transactions | get AMM involved join/exit transactions
[**GetAmmUserTransactionsV2**](LoopringDEXRestfulAPIApi.md#GetAmmUserTransactionsV2) | **Get** /api/v2/amm/user/transactions | User&#39;s AMM join/exit transactions
[**GetAmmUserTransactionsV3**](LoopringDEXRestfulAPIApi.md#GetAmmUserTransactionsV3) | **Get** /api/v3/amm/user/transactions | User&#39;s AMM join/exit transactions
[**GetApiKeyV2**](LoopringDEXRestfulAPIApi.md#GetApiKeyV2) | **Get** /api/v2/apiKey | Get user ApiKey
[**GetApiKeyV3**](LoopringDEXRestfulAPIApi.md#GetApiKeyV3) | **Get** /api/v3/apiKey | Get user ApiKey
[**GetAvailableBrokerV3**](LoopringDEXRestfulAPIApi.md#GetAvailableBrokerV3) | **Get** /api/v3/getAvailableBroker | 
[**GetAvailableBrokerV3Hebao**](LoopringDEXRestfulAPIApi.md#GetAvailableBrokerV3Hebao) | **Get** /api/v3/hebao/getAvailableBroker | 
[**GetBlockV3**](LoopringDEXRestfulAPIApi.md#GetBlockV3) | **Get** /api/v3/block/getBlock | Get L2 block info
[**GetBusinessFee2V2**](LoopringDEXRestfulAPIApi.md#GetBusinessFee2V2) | **Get** /api/v2/user/offchainFee | Query current fee amount
[**GetBusinessFee2V3**](LoopringDEXRestfulAPIApi.md#GetBusinessFee2V3) | **Get** /api/v3/user/offchainFee | Query current fee amount
[**GetCandlestickV2**](LoopringDEXRestfulAPIApi.md#GetCandlestickV2) | **Get** /api/v2/candlestick | Get market candlestick
[**GetCandlestickV3**](LoopringDEXRestfulAPIApi.md#GetCandlestickV3) | **Get** /api/v3/candlestick | Get market candlestick
[**GetClaimHistoryV3**](LoopringDEXRestfulAPIApi.md#GetClaimHistoryV3) | **Get** /api/v3/luckyToken/user/claimHistory | 
[**GetClaimLuckyTokenRecordsV3**](LoopringDEXRestfulAPIApi.md#GetClaimLuckyTokenRecordsV3) | **Get** /api/v3/luckyToken/user/luckyTokenDetail | 
[**GetClaimLuckyTokenV3**](LoopringDEXRestfulAPIApi.md#GetClaimLuckyTokenV3) | **Get** /api/v3/luckyToken/user/claimedLuckyTokens | 
[**GetCommissionRewardRankV3**](LoopringDEXRestfulAPIApi.md#GetCommissionRewardRankV3) | **Get** /api/v3/sidecar/commissionRewardRank | 
[**GetCommissionRewardV3**](LoopringDEXRestfulAPIApi.md#GetCommissionRewardV3) | **Get** /api/v3/sidecar/commissionReward | 
[**GetCommissionTotalRewardV3**](LoopringDEXRestfulAPIApi.md#GetCommissionTotalRewardV3) | **Get** /api/v3/sidecar/commissionRewardTotal | 
[**GetCounterFactualInfoV3**](LoopringDEXRestfulAPIApi.md#GetCounterFactualInfoV3) | **Get** /api/v3/counterFactualInfo | api.getCounterFactualInfo.value
[**GetCurrencyRatesV2**](LoopringDEXRestfulAPIApi.md#GetCurrencyRatesV2) | **Get** /api/v2/currencyRates | 
[**GetDaoSquareAccountsV3**](LoopringDEXRestfulAPIApi.md#GetDaoSquareAccountsV3) | **Get** /api/v3/user/daoSquareAccounts | 
[**GetDepthV2**](LoopringDEXRestfulAPIApi.md#GetDepthV2) | **Get** /api/v2/depth | Get market orderbook
[**GetDepthV3**](LoopringDEXRestfulAPIApi.md#GetDepthV3) | **Get** /api/v3/depth | Get market orderbook
[**GetEthBalanceV3**](LoopringDEXRestfulAPIApi.md#GetEthBalanceV3) | **Get** /api/v3/eth/balances | Get user&#39;s Ether balance on Ethereum mainnet
[**GetEthNonceV3**](LoopringDEXRestfulAPIApi.md#GetEthNonceV3) | **Get** /api/v3/eth/nonce | Get user&#39;s next Ethereum nonce
[**GetExchangeFeeInfoV2**](LoopringDEXRestfulAPIApi.md#GetExchangeFeeInfoV2) | **Get** /api/v2/exchange/feeInfo | Get exchange configurations
[**GetExchangeFeeInfoV3**](LoopringDEXRestfulAPIApi.md#GetExchangeFeeInfoV3) | **Get** /api/v3/exchange/feeInfo | Get exchange configurations
[**GetExchangeInfoV2**](LoopringDEXRestfulAPIApi.md#GetExchangeInfoV2) | **Get** /api/v2/exchange/info | Get exchange configurations
[**GetExchangeInfoV3**](LoopringDEXRestfulAPIApi.md#GetExchangeInfoV3) | **Get** /api/v3/exchange/info | Get exchange configurations
[**GetFinanceIncomeV3**](LoopringDEXRestfulAPIApi.md#GetFinanceIncomeV3) | **Get** /api/v3/user/financeIncome | 
[**GetLatestTokenPricesV3**](LoopringDEXRestfulAPIApi.md#GetLatestTokenPricesV3) | **Get** /api/v3/datacenter/getLatestTokenPrices | 
[**GetLiquidityMiningConfV3**](LoopringDEXRestfulAPIApi.md#GetLiquidityMiningConfV3) | **Get** /api/v3/sidecar/liquidityMiningConf | 
[**GetLiquidityMiningRankV3**](LoopringDEXRestfulAPIApi.md#GetLiquidityMiningRankV3) | **Get** /api/v3/sidecar/liquidityMiningRank | 
[**GetLiquidityMiningRewardTotalV3**](LoopringDEXRestfulAPIApi.md#GetLiquidityMiningRewardTotalV3) | **Get** /api/v3/sidecar/liquidityMiningTotal | 
[**GetLiquidityMiningRewardsV3**](LoopringDEXRestfulAPIApi.md#GetLiquidityMiningRewardsV3) | **Get** /api/v3/sidecar/liquidityMining | 
[**GetLiquidityMiningUserHistoryV3**](LoopringDEXRestfulAPIApi.md#GetLiquidityMiningUserHistoryV3) | **Get** /api/v3/sidecar/liquidityMiningUserHistory | 
[**GetLuckyTokenAgentsV3**](LoopringDEXRestfulAPIApi.md#GetLuckyTokenAgentsV3) | **Get** /api/v3/luckyToken/agents | api.getLuckyTokenAgents.value
[**GetLuckyTokenSignersV3**](LoopringDEXRestfulAPIApi.md#GetLuckyTokenSignersV3) | **Get** /api/v3/luckyToken/authorizedSigners | api.getLuckyTokenSigners.value
[**GetLuckyTokenSummaryV3**](LoopringDEXRestfulAPIApi.md#GetLuckyTokenSummaryV3) | **Get** /api/v3/luckyToken/user/summary | 
[**GetLuckyTokenUserBalancesV3**](LoopringDEXRestfulAPIApi.md#GetLuckyTokenUserBalancesV3) | **Get** /api/v3/luckyToken/user/balances | api.getLuckyTokenBalance.value
[**GetLuckyTokenV3**](LoopringDEXRestfulAPIApi.md#GetLuckyTokenV3) | **Get** /api/v3/luckyToken/user/luckyTokens | 
[**GetLuckyTokenWithdrawV3**](LoopringDEXRestfulAPIApi.md#GetLuckyTokenWithdrawV3) | **Get** /api/v3/luckyToken/user/withdraws | 
[**GetMarketTradeV2**](LoopringDEXRestfulAPIApi.md#GetMarketTradeV2) | **Get** /api/v2/trade | Get market recent trades
[**GetMarketTradeV3**](LoopringDEXRestfulAPIApi.md#GetMarketTradeV3) | **Get** /api/v3/trade | Get market recent trades
[**GetMarketsV2**](LoopringDEXRestfulAPIApi.md#GetMarketsV2) | **Get** /api/v2/exchange/markets | Get market configurations
[**GetMarketsV3**](LoopringDEXRestfulAPIApi.md#GetMarketsV3) | **Get** /api/v3/exchange/markets | Get market configurations
[**GetMixedCandlestickV2**](LoopringDEXRestfulAPIApi.md#GetMixedCandlestickV2) | **Get** /api/v2/mix/candlestick | Get market candlestick
[**GetMixedCandlestickV3**](LoopringDEXRestfulAPIApi.md#GetMixedCandlestickV3) | **Get** /api/v3/mix/candlestick | Get mixed market candlestick
[**GetMixedDepthV2**](LoopringDEXRestfulAPIApi.md#GetMixedDepthV2) | **Get** /api/v2/mix/depth | Get market orderbook
[**GetMixedDepthV3**](LoopringDEXRestfulAPIApi.md#GetMixedDepthV3) | **Get** /api/v3/mix/depth | Get mixed market orderbook
[**GetMixedMarketsV2**](LoopringDEXRestfulAPIApi.md#GetMixedMarketsV2) | **Get** /api/v2/mix/markets | Get market configurations
[**GetMixedMarketsV3**](LoopringDEXRestfulAPIApi.md#GetMixedMarketsV3) | **Get** /api/v3/mix/markets | Get mixed market configurations
[**GetMixedTickerV2**](LoopringDEXRestfulAPIApi.md#GetMixedTickerV2) | **Get** /api/v2/mix/ticker | Get market ticker
[**GetMixedTickerV3**](LoopringDEXRestfulAPIApi.md#GetMixedTickerV3) | **Get** /api/v3/mix/ticker | Get mixed market ticker
[**GetNewestTokenPricesV3**](LoopringDEXRestfulAPIApi.md#GetNewestTokenPricesV3) | **Get** /api/v3/datacenter/getNewestTokenPrices | 
[**GetNextStorageIdV2**](LoopringDEXRestfulAPIApi.md#GetNextStorageIdV2) | **Get** /api/v2/storageId | Get next storage ID
[**GetNextStorageIdV3**](LoopringDEXRestfulAPIApi.md#GetNextStorageIdV3) | **Get** /api/v3/storageId | Get next storage ID
[**GetNftDataV3**](LoopringDEXRestfulAPIApi.md#GetNftDataV3) | **Get** /api/v3/nft/info/nftData | Query nftDatas by minter, tokenAddress and nftID.
[**GetNftDepositsV3**](LoopringDEXRestfulAPIApi.md#GetNftDepositsV3) | **Get** /api/v3/user/nft/deposits | Get user NFT deposit history
[**GetNftHoldersV3**](LoopringDEXRestfulAPIApi.md#GetNftHoldersV3) | **Get** /api/v3/nft/info/nftHolders | Query NFT holders by looprings nftData
[**GetNftMintersV3**](LoopringDEXRestfulAPIApi.md#GetNftMintersV3) | **Get** /api/v3/nft/info/minters | Get minters
[**GetNftOrderUserRateAmountV3**](LoopringDEXRestfulAPIApi.md#GetNftOrderUserRateAmountV3) | **Get** /api/v3/nft/info/orderUserRateAmount | Query current token minimum amount to place order based on users VIP level and max fee bips
[**GetNftRequestFeesV3**](LoopringDEXRestfulAPIApi.md#GetNftRequestFeesV3) | **Get** /api/v3/user/nft/offchainFee | Query current NFT requests fee amount
[**GetNftTokenAddressesV3**](LoopringDEXRestfulAPIApi.md#GetNftTokenAddressesV3) | **Get** /api/v3/nft/info/nftTokenAddresses | Get Nft token addresses
[**GetNftTradesV3**](LoopringDEXRestfulAPIApi.md#GetNftTradesV3) | **Get** /api/v3/user/nft/trades | Users Nft trade list
[**GetNftTransfersV3**](LoopringDEXRestfulAPIApi.md#GetNftTransfersV3) | **Get** /api/v3/user/nft/transfers | Get user NFT transfer history
[**GetNftWithdrawalsV3**](LoopringDEXRestfulAPIApi.md#GetNftWithdrawalsV3) | **Get** /api/v3/user/nft/withdrawals | Get user NFT withdraw history
[**GetNftsInfoV3**](LoopringDEXRestfulAPIApi.md#GetNftsInfoV3) | **Get** /api/v3/nft/info/nfts | Query NFT info by looprings nftData
[**GetNotWithdrawContractTokensV3**](LoopringDEXRestfulAPIApi.md#GetNotWithdrawContractTokensV3) | **Get** /api/v3/exchange/notWithdrawContractTokens | api.getNotWithdrawContractTokens.value
[**GetOrderAmountV2**](LoopringDEXRestfulAPIApi.md#GetOrderAmountV2) | **Get** /api/v2/user/orderAmount | Query current token minimum amount to place order
[**GetOrderAmountV3**](LoopringDEXRestfulAPIApi.md#GetOrderAmountV3) | **Get** /api/v3/user/orderAmount | Query current token minimum amount to place order
[**GetOrderDetailV2**](LoopringDEXRestfulAPIApi.md#GetOrderDetailV2) | **Get** /api/v2/order | Get order details
[**GetOrderDetailV3**](LoopringDEXRestfulAPIApi.md#GetOrderDetailV3) | **Get** /api/v3/order | Get order details
[**GetOrderUserRateAmountV2**](LoopringDEXRestfulAPIApi.md#GetOrderUserRateAmountV2) | **Get** /api/v2/user/orderUserRateAmount | Query current token minimum amount to place order
[**GetOrderUserRateAmountV3**](LoopringDEXRestfulAPIApi.md#GetOrderUserRateAmountV3) | **Get** /api/v3/user/orderUserRateAmount | Query current token minimum amount to place order based on users VIP level and max fee bips
[**GetOrdersV2**](LoopringDEXRestfulAPIApi.md#GetOrdersV2) | **Get** /api/v2/orders | Get multiple orders
[**GetOrdersV3**](LoopringDEXRestfulAPIApi.md#GetOrdersV3) | **Get** /api/v3/orders | Get multiple orders
[**GetPendingRequestsV3**](LoopringDEXRestfulAPIApi.md#GetPendingRequestsV3) | **Get** /api/v3/block/getPendingRequests | Get pending txs
[**GetPoolsStatsAmmV3**](LoopringDEXRestfulAPIApi.md#GetPoolsStatsAmmV3) | **Get** /api/v3/amm/poolsStats | api.getPoolsStats.value
[**GetPoolsStatsV2**](LoopringDEXRestfulAPIApi.md#GetPoolsStatsV2) | **Get** /api/v2/amm/poolsStats | api.getPoolsStats.value
[**GetPoolsStatsV3**](LoopringDEXRestfulAPIApi.md#GetPoolsStatsV3) | **Get** /api/v3/poolsStats | api.getPoolsStats.value
[**GetPriceV2**](LoopringDEXRestfulAPIApi.md#GetPriceV2) | **Get** /api/v2/price | Get token fiat prices
[**GetPriceV3**](LoopringDEXRestfulAPIApi.md#GetPriceV3) | **Get** /api/v3/price | Get token fiat prices
[**GetProfitSharingStatusV3**](LoopringDEXRestfulAPIApi.md#GetProfitSharingStatusV3) | **Get** /api/v3/sidecar/profitShareReward | 
[**GetProtocolPortraitV3**](LoopringDEXRestfulAPIApi.md#GetProtocolPortraitV3) | **Get** /api/v3/sidecar/ProtocolPortrait | 
[**GetRecommendedGasPriceRangeV3**](LoopringDEXRestfulAPIApi.md#GetRecommendedGasPriceRangeV3) | **Get** /api/v3/eth/recommendedGasPriceRange | Get recommended gas price
[**GetRecommendedGasPriceV3**](LoopringDEXRestfulAPIApi.md#GetRecommendedGasPriceV3) | **Get** /api/v3/eth/recommendedGasPrice | Get recommended gas price
[**GetRecommendedMarketsV3**](LoopringDEXRestfulAPIApi.md#GetRecommendedMarketsV3) | **Get** /api/v3/exchange/recommended | api.getRecommendedMarkets.value
[**GetRequestInBlockV3**](LoopringDEXRestfulAPIApi.md#GetRequestInBlockV3) | **Get** /api/v3/requestInBlock | Get a txs block
[**GetTickerV2**](LoopringDEXRestfulAPIApi.md#GetTickerV2) | **Get** /api/v2/ticker | Get market ticker
[**GetTickerV3**](LoopringDEXRestfulAPIApi.md#GetTickerV3) | **Get** /api/v3/ticker | Get market ticker
[**GetTimestampV2**](LoopringDEXRestfulAPIApi.md#GetTimestampV2) | **Get** /api/v2/timestamp | Get relayer&#39;s current time
[**GetTimestampV3**](LoopringDEXRestfulAPIApi.md#GetTimestampV3) | **Get** /api/v3/timestamp | Get relayer&#39;s current time
[**GetTokenBalanceV3**](LoopringDEXRestfulAPIApi.md#GetTokenBalanceV3) | **Get** /api/v3/eth/tokenBalances | Get user token balance
[**GetTokenDecimalV2**](LoopringDEXRestfulAPIApi.md#GetTokenDecimalV2) | **Get** /api/v2/token/decimal | 
[**GetTokenInfoV3**](LoopringDEXRestfulAPIApi.md#GetTokenInfoV3) | **Get** /api/v3/spi/tokenInfo | api.getTokenInfo.value
[**GetTokenPricesV3**](LoopringDEXRestfulAPIApi.md#GetTokenPricesV3) | **Get** /api/v3/datacenter/getTokenPrices | 
[**GetTokensV2**](LoopringDEXRestfulAPIApi.md#GetTokensV2) | **Get** /api/v2/exchange/tokens | Get token configurations
[**GetTokensV3**](LoopringDEXRestfulAPIApi.md#GetTokensV3) | **Get** /api/v3/exchange/tokens | Get token configurations
[**GetTournamentRankV3**](LoopringDEXRestfulAPIApi.md#GetTournamentRankV3) | **Get** /api/v3/game/rank | 
[**GetUserAssetsV3**](LoopringDEXRestfulAPIApi.md#GetUserAssetsV3) | **Get** /api/v3/datacenter/getUserAssets | 
[**GetUserBalancesV2**](LoopringDEXRestfulAPIApi.md#GetUserBalancesV2) | **Get** /api/v2/user/balances | Get user exchange balances
[**GetUserBalancesV3**](LoopringDEXRestfulAPIApi.md#GetUserBalancesV3) | **Get** /api/v3/user/balances | Get user exchange balances
[**GetUserBillV3**](LoopringDEXRestfulAPIApi.md#GetUserBillV3) | **Get** /api/v3/user/bills | 
[**GetUserCreateV2**](LoopringDEXRestfulAPIApi.md#GetUserCreateV2) | **Get** /api/v2/user/createInfo | Get user registration transactions
[**GetUserCreateV3**](LoopringDEXRestfulAPIApi.md#GetUserCreateV3) | **Get** /api/v3/user/createInfo | Get user registration transactions
[**GetUserDepositsV2**](LoopringDEXRestfulAPIApi.md#GetUserDepositsV2) | **Get** /api/v2/user/deposits | Get user deposit history
[**GetUserDepositsV3**](LoopringDEXRestfulAPIApi.md#GetUserDepositsV3) | **Get** /api/v3/user/deposits | Get user deposit history
[**GetUserFeeRates2V2**](LoopringDEXRestfulAPIApi.md#GetUserFeeRates2V2) | **Get** /api/v2/user/orderFee | Query user place order fee rate
[**GetUserFeeRates2V3**](LoopringDEXRestfulAPIApi.md#GetUserFeeRates2V3) | **Get** /api/v3/user/orderFee | Query user place order fee rate
[**GetUserNftBalancesV3**](LoopringDEXRestfulAPIApi.md#GetUserNftBalancesV3) | **Get** /api/v3/user/nft/balances | Get users NFT balance
[**GetUserNftMintsV3**](LoopringDEXRestfulAPIApi.md#GetUserNftMintsV3) | **Get** /api/v3/user/nft/mints | Get user NFT mint history
[**GetUserNftOrderFeeRatesV3**](LoopringDEXRestfulAPIApi.md#GetUserNftOrderFeeRatesV3) | **Get** /api/v3/user/nft/orderFee | Query user place order fee rate
[**GetUserNftTransactionsV3**](LoopringDEXRestfulAPIApi.md#GetUserNftTransactionsV3) | **Get** /api/v3/user/nft/transactions | api.getUserNftTransactions.value
[**GetUserRewardsV2**](LoopringDEXRestfulAPIApi.md#GetUserRewardsV2) | **Get** /api/v2/amm/user/rewards | api.getPoolsStats.value
[**GetUserRewardsV3**](LoopringDEXRestfulAPIApi.md#GetUserRewardsV3) | **Get** /api/v3/amm/user/rewards | api.getPoolsStats.value
[**GetUserTournamentRankV3**](LoopringDEXRestfulAPIApi.md#GetUserTournamentRankV3) | **Get** /api/v3/game/user/rank | 
[**GetUserTradeAmountV3**](LoopringDEXRestfulAPIApi.md#GetUserTradeAmountV3) | **Get** /api/v3/datacenter/getUserTradeAmount | 
[**GetUserTradesV2**](LoopringDEXRestfulAPIApi.md#GetUserTradesV2) | **Get** /api/v2/user/trades | Get user trade history
[**GetUserTradesV3**](LoopringDEXRestfulAPIApi.md#GetUserTradesV3) | **Get** /api/v3/user/trades | Get user trade history
[**GetUserTransactionsV3**](LoopringDEXRestfulAPIApi.md#GetUserTransactionsV3) | **Get** /api/v3/user/transactions | 
[**GetUserTransfersV2**](LoopringDEXRestfulAPIApi.md#GetUserTransfersV2) | **Get** /api/v2/user/transfers | Get user transfer list
[**GetUserTransfersV3**](LoopringDEXRestfulAPIApi.md#GetUserTransfersV3) | **Get** /api/v3/user/transfers | Get user transfer list
[**GetUserUpdateV2**](LoopringDEXRestfulAPIApi.md#GetUserUpdateV2) | **Get** /api/v2/user/updateInfo | Get password reset transactions
[**GetUserUpdateV3**](LoopringDEXRestfulAPIApi.md#GetUserUpdateV3) | **Get** /api/v3/user/updateInfo | Get password reset transactions
[**GetUserVipInfoV3**](LoopringDEXRestfulAPIApi.md#GetUserVipInfoV3) | **Get** /api/v3/user/vipInfo | api.getUserVipInfo.value
[**GetUserWithdrawalsV2**](LoopringDEXRestfulAPIApi.md#GetUserWithdrawalsV2) | **Get** /api/v2/user/withdrawals | Get user onchain withdrawal history
[**GetUserWithdrawalsV3**](LoopringDEXRestfulAPIApi.md#GetUserWithdrawalsV3) | **Get** /api/v3/user/withdrawals | Get user onchain withdrawal history
[**GetWithdrawalAgentsV3**](LoopringDEXRestfulAPIApi.md#GetWithdrawalAgentsV3) | **Get** /api/v3/exchange/withdrawalAgents | api.getWithdrawalAgents.value
[**OpenHebaoAccountV3Hebao**](LoopringDEXRestfulAPIApi.md#OpenHebaoAccountV3Hebao) | **Post** /api/v3/hebao/openAccount | 
[**PrepareHebaoPayV3Hebao**](LoopringDEXRestfulAPIApi.md#PrepareHebaoPayV3Hebao) | **Post** /api/v3/hebao/preparePay | 
[**RefundHebaoPayV3Hebao**](LoopringDEXRestfulAPIApi.md#RefundHebaoPayV3Hebao) | **Post** /api/v3/hebao/refundPay | 
[**SearchTokenMetadataV2**](LoopringDEXRestfulAPIApi.md#SearchTokenMetadataV2) | **Get** /api/v2/token/metadata | 
[**SearchTokenV2**](LoopringDEXRestfulAPIApi.md#SearchTokenV2) | **Get** /api/v2/tokens/search | 
[**SendLuckyTokenV3**](LoopringDEXRestfulAPIApi.md#SendLuckyTokenV3) | **Post** /api/v3/luckyToken/sendLuckyToken | api.sendLuckyToken.value
[**SendTransactionV3**](LoopringDEXRestfulAPIApi.md#SendTransactionV3) | **Post** /api/v3/eth/sendTx | Send a raw Ethereum transaction
[**SetReferrerV3**](LoopringDEXRestfulAPIApi.md#SetReferrerV3) | **Post** /api/v3/refer | 
[**SubmitAmmPoolExitV2**](LoopringDEXRestfulAPIApi.md#SubmitAmmPoolExitV2) | **Post** /api/v2/amm/exit | Exit an AMM pool
[**SubmitAmmPoolExitV3**](LoopringDEXRestfulAPIApi.md#SubmitAmmPoolExitV3) | **Post** /api/v3/amm/exit | Exit an AMM pool
[**SubmitAmmPoolJoinV2**](LoopringDEXRestfulAPIApi.md#SubmitAmmPoolJoinV2) | **Post** /api/v2/amm/join | Join into AMM pool
[**SubmitAmmPoolJoinV3**](LoopringDEXRestfulAPIApi.md#SubmitAmmPoolJoinV3) | **Post** /api/v3/amm/join | Join into AMM pool
[**SubmitDualAuthTransferV2**](LoopringDEXRestfulAPIApi.md#SubmitDualAuthTransferV2) | **Post** /api/v2/dualAuthTransfer | Submit dual authority transfer
[**SubmitMintNftV3**](LoopringDEXRestfulAPIApi.md#SubmitMintNftV3) | **Post** /api/v3/nft/mint | Mint a NFT token in Loopring L2
[**SubmitNftTradeV3**](LoopringDEXRestfulAPIApi.md#SubmitNftTradeV3) | **Post** /api/v3/nft/trade | Settle down an NFT trade which has two matched orders
[**SubmitNftTransferV3**](LoopringDEXRestfulAPIApi.md#SubmitNftTransferV3) | **Post** /api/v3/nft/transfer | Submit internal NFT transfer
[**SubmitOffChainNftWithdrawalV3**](LoopringDEXRestfulAPIApi.md#SubmitOffChainNftWithdrawalV3) | **Post** /api/v3/nft/withdrawal | Withdraw a NFT token
[**SubmitOffChainWithdrawalV2**](LoopringDEXRestfulAPIApi.md#SubmitOffChainWithdrawalV2) | **Post** /api/v2/user/withdrawals | Submit offchain withdraw request
[**SubmitOffChainWithdrawalV3**](LoopringDEXRestfulAPIApi.md#SubmitOffChainWithdrawalV3) | **Post** /api/v3/user/withdrawals | Submit offchain withdraw request
[**SubmitOrderV2V2**](LoopringDEXRestfulAPIApi.md#SubmitOrderV2V2) | **Post** /api/v2/order | api.submitOrderV2.value
[**SubmitOrderV3V3**](LoopringDEXRestfulAPIApi.md#SubmitOrderV3V3) | **Post** /api/v3/order | Submit an order
[**SubmitRecommendedMarketsV3**](LoopringDEXRestfulAPIApi.md#SubmitRecommendedMarketsV3) | **Post** /api/v3/exchange/recommended | api.getRecommendedMarkets.value
[**SubmitRewardV3**](LoopringDEXRestfulAPIApi.md#SubmitRewardV3) | **Post** /api/v3/reward | 
[**SubmitTransferV2**](LoopringDEXRestfulAPIApi.md#SubmitTransferV2) | **Post** /api/v2/transfer | Submit internal transfer
[**SubmitTransferV3**](LoopringDEXRestfulAPIApi.md#SubmitTransferV3) | **Post** /api/v3/transfer | Submit internal transfer
[**SubmitUpdateAccountV2**](LoopringDEXRestfulAPIApi.md#SubmitUpdateAccountV2) | **Post** /api/v2/account | Update account EDDSA key
[**SubmitUpdateAccountV3**](LoopringDEXRestfulAPIApi.md#SubmitUpdateAccountV3) | **Post** /api/v3/account | Update account EDDSA key
[**TokenMetadatasV2**](LoopringDEXRestfulAPIApi.md#TokenMetadatasV2) | **Get** /api/v2/tokens/metadata | 
[**ValidateNftOrderV3**](LoopringDEXRestfulAPIApi.md#ValidateNftOrderV3) | **Post** /api/v3/nft/validateOrder | Validate a NFT order
[**VerifyAllEcdsaV3**](LoopringDEXRestfulAPIApi.md#VerifyAllEcdsaV3) | **Post** /api/v3/verifyAllEcdsa | 
[**WithdrawLuckyTokenV3**](LoopringDEXRestfulAPIApi.md#WithdrawLuckyTokenV3) | **Post** /api/v3/luckyToken/user/withdrawals | api.withdrawLuckyToken.value



## ApplyApiKeyV2

> GetApiKeyResponseV2 ApplyApiKeyV2(ctx).Body(body).Execute()

Update user's ApiKey



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewGetApiKeyRequest(int64(1)) // GetApiKeyRequest | api.applyApiKey.implicit.value

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.ApplyApiKeyV2(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.ApplyApiKeyV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplyApiKeyV2`: GetApiKeyResponseV2
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.ApplyApiKeyV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApplyApiKeyV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GetApiKeyRequest**](GetApiKeyRequest.md) | api.applyApiKey.implicit.value | 

### Return type

[**GetApiKeyResponseV2**](GetApiKeyResponseV2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplyApiKeyV3

> GetApiKeyResponseV3 ApplyApiKeyV3(ctx).Body(body).Execute()

Update user's ApiKey



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewGetApiKeyRequest(int64(1)) // GetApiKeyRequest | api.applyApiKey.implicit.value

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.ApplyApiKeyV3(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.ApplyApiKeyV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplyApiKeyV3`: GetApiKeyResponseV3
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.ApplyApiKeyV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApplyApiKeyV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GetApiKeyRequest**](GetApiKeyRequest.md) | api.applyApiKey.implicit.value | 

### Return type

[**GetApiKeyResponseV3**](GetApiKeyResponseV3.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchCancelOrdersByClientOrderIdV2

> BatchCancelOrderResponse BatchCancelOrdersByClientOrderIdV2(ctx).AccountId(accountId).ClientOrderId(clientOrderId).Execute()

Cancel multiple orders by clientOrderId



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := int32(1) // int32 | request.cancelOrders.accountId
    clientOrderId := "20200318000000001010,20200318000000001011" // string | ClientOrderIds of orders to be cancelled.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.BatchCancelOrdersByClientOrderIdV2(context.Background()).AccountId(accountId).ClientOrderId(clientOrderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.BatchCancelOrdersByClientOrderIdV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchCancelOrdersByClientOrderIdV2`: BatchCancelOrderResponse
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.BatchCancelOrdersByClientOrderIdV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchCancelOrdersByClientOrderIdV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32** | request.cancelOrders.accountId | 
 **clientOrderId** | **string** | ClientOrderIds of orders to be cancelled. | 

### Return type

[**BatchCancelOrderResponse**](BatchCancelOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchCancelOrdersByClientOrderIdV3

> map[string]interface{} BatchCancelOrdersByClientOrderIdV3(ctx).AccountId(accountId).ClientOrderId(clientOrderId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := int32(1) // int32 | request.cancelOrders.accountId
    clientOrderId := "20200318000000001010,20200318000000001011" // string | ClientOrderIds of orders to be cancelled.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.BatchCancelOrdersByClientOrderIdV3(context.Background()).AccountId(accountId).ClientOrderId(clientOrderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.BatchCancelOrdersByClientOrderIdV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchCancelOrdersByClientOrderIdV3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.BatchCancelOrdersByClientOrderIdV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchCancelOrdersByClientOrderIdV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32** | request.cancelOrders.accountId | 
 **clientOrderId** | **string** | ClientOrderIds of orders to be cancelled. | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchCancelOrdersByOrderHashV2

> BatchCancelOrderResponse BatchCancelOrdersByOrderHashV2(ctx).AccountId(accountId).OrderHash(orderHash).Execute()

Cancel multiple orders by hash



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := int32(1) // int32 | request.cancelOrders.accountId
    orderHash := "13375450901292179417154974849571793069911517354720397125027633242680470075859,13375450901292179417154974849571793069911517354720397125027633242680470075860" // string | Hash of order to be canceled,separate multiple hashes with commas.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.BatchCancelOrdersByOrderHashV2(context.Background()).AccountId(accountId).OrderHash(orderHash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.BatchCancelOrdersByOrderHashV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchCancelOrdersByOrderHashV2`: BatchCancelOrderResponse
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.BatchCancelOrdersByOrderHashV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchCancelOrdersByOrderHashV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32** | request.cancelOrders.accountId | 
 **orderHash** | **string** | Hash of order to be canceled,separate multiple hashes with commas. | 

### Return type

[**BatchCancelOrderResponse**](BatchCancelOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchCancelOrdersByOrderHashV3

> map[string]interface{} BatchCancelOrdersByOrderHashV3(ctx).AccountId(accountId).OrderHash(orderHash).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := int32(1) // int32 | request.cancelOrders.accountId
    orderHash := "13375450901292179417154974849571793069911517354720397125027633242680470075859,13375450901292179417154974849571793069911517354720397125027633242680470075860" // string | Hash of order to be canceled,separate multiple hashes with commas.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.BatchCancelOrdersByOrderHashV3(context.Background()).AccountId(accountId).OrderHash(orderHash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.BatchCancelOrdersByOrderHashV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchCancelOrdersByOrderHashV3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.BatchCancelOrdersByOrderHashV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchCancelOrdersByOrderHashV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32** | request.cancelOrders.accountId | 
 **orderHash** | **string** | Hash of order to be canceled,separate multiple hashes with commas. | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchSubmitOrderV2

> BatchSubmitOrderResponse BatchSubmitOrderV2(ctx).Body(body).Execute()

Submit multiple orders



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewBatchSubmitOrderRequest([]openapiclient.SubmitOrderRequest{*openapiclient.NewSubmitOrderRequest("1", int32(1), int64(1), int32(0), int32(2), "1000000000000000000", "1000000000000000000", "true", int64(1567053142), int32(20), "13375450901292179417154974849571793069911517354720397125027633242680470075859")}) // BatchSubmitOrderRequest | Body of batch submit orders.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.BatchSubmitOrderV2(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.BatchSubmitOrderV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchSubmitOrderV2`: BatchSubmitOrderResponse
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.BatchSubmitOrderV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchSubmitOrderV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**BatchSubmitOrderRequest**](BatchSubmitOrderRequest.md) | Body of batch submit orders. | 

### Return type

[**BatchSubmitOrderResponse**](BatchSubmitOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CanMintV3

> bool CanMintV3(ctx).Minter(minter).NftTokenAddress(nftTokenAddress).Execute()

check if a minter is authorized on a nft token contract



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    minter := "minter_example" // string | minter address
    nftTokenAddress := "nftTokenAddress_example" // string | nft token address

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.CanMintV3(context.Background()).Minter(minter).NftTokenAddress(nftTokenAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.CanMintV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CanMintV3`: bool
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.CanMintV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCanMintV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **minter** | **string** | minter address | 
 **nftTokenAddress** | **string** | nft token address | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelOrderV3

> SubmitOrderResponseItemV3 CancelOrderV3(ctx).AccountId(accountId).OrderHash(orderHash).ClientOrderId(clientOrderId).Execute()

Cancel order



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := int64(1) // int64 | Account ID
    orderHash := "13375450901292179417154974849571793069911517354720397125027633242680470075859" // string | Order HASH (optional)
    clientOrderId := "20200318000000001010" // string | The unique order ID of the client (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.CancelOrderV3(context.Background()).AccountId(accountId).OrderHash(orderHash).ClientOrderId(clientOrderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.CancelOrderV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelOrderV3`: SubmitOrderResponseItemV3
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.CancelOrderV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCancelOrderV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int64** | Account ID | 
 **orderHash** | **string** | Order HASH | 
 **clientOrderId** | **string** | The unique order ID of the client | 

### Return type

[**SubmitOrderResponseItemV3**](SubmitOrderResponseItemV3.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelOrdersV2

> CancelOrdersResponseData CancelOrdersV2(ctx).AccountId(accountId).OrderHash(orderHash).ClientOrderId(clientOrderId).Execute()

api.CancelOrders.value



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := int64(1) // int64 | request.cancelOrders.accountId
    orderHash := "13375450901292179417154974849571793069911517354720397125027633242680470075859" // string | request.cancelOrders.orderHash (optional)
    clientOrderId := "20200318000000001010" // string | request.cancelOrders.clientOrderId (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.CancelOrdersV2(context.Background()).AccountId(accountId).OrderHash(orderHash).ClientOrderId(clientOrderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.CancelOrdersV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelOrdersV2`: CancelOrdersResponseData
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.CancelOrdersV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCancelOrdersV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int64** | request.cancelOrders.accountId | 
 **orderHash** | **string** | request.cancelOrders.orderHash | 
 **clientOrderId** | **string** | request.cancelOrders.clientOrderId | 

### Return type

[**CancelOrdersResponseData**](CancelOrdersResponseData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClaimLuckyTokenV3

> GetLuckyTokenClaimAmountResponse ClaimLuckyTokenV3(ctx).Body(body).Execute()

api.claimLuckyToken.value



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewClaimLuckyTokenRequestV3("0xf7c932351186c3a9053f313eefa16209c018f7f1dba8aa8ca7100400f7c31085", "0xbbbbca6a901c926f240b89eacb641d8aec7aeafd") // ClaimLuckyTokenRequestV3 | api.claimLuckyToken.implicit.value

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.ClaimLuckyTokenV3(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.ClaimLuckyTokenV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClaimLuckyTokenV3`: GetLuckyTokenClaimAmountResponse
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.ClaimLuckyTokenV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiClaimLuckyTokenV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ClaimLuckyTokenRequestV3**](ClaimLuckyTokenRequestV3.md) | api.claimLuckyToken.implicit.value | 

### Return type

[**GetLuckyTokenClaimAmountResponse**](GetLuckyTokenClaimAmountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComputeTokenAddressV3

> map[string]interface{} ComputeTokenAddressV3(ctx).NftFactory(nftFactory).NftOwner(nftOwner).NftBaseUri(nftBaseUri).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    nftFactory := "nftFactory_example" // string |  (optional)
    nftOwner := "nftOwner_example" // string |  (optional)
    nftBaseUri := "nftBaseUri_example" // string | request.computeTokenAddress.nftBaseUri (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.ComputeTokenAddressV3(context.Background()).NftFactory(nftFactory).NftOwner(nftOwner).NftBaseUri(nftBaseUri).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.ComputeTokenAddressV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComputeTokenAddressV3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.ComputeTokenAddressV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiComputeTokenAddressV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nftFactory** | **string** |  | 
 **nftOwner** | **string** |  | 
 **nftBaseUri** | **string** | request.computeTokenAddress.nftBaseUri | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfirmHebaoPayV3Hebao

> map[string]interface{} ConfirmHebaoPayV3Hebao(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.ConfirmHebaoPayV3Hebao(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.ConfirmHebaoPayV3Hebao``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConfirmHebaoPayV3Hebao`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.ConfirmHebaoPayV3Hebao`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiConfirmHebaoPayV3HebaoRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CrawlPriceFromBinanceV2

> map[string]interface{} CrawlPriceFromBinanceV2(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.CrawlPriceFromBinanceV2(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.CrawlPriceFromBinanceV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CrawlPriceFromBinanceV2`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.CrawlPriceFromBinanceV2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCrawlPriceFromBinanceV2Request struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CrawlSiteV3

> map[string]interface{} CrawlSiteV3(ctx).Site(site).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    site := "site_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.CrawlSiteV3(context.Background()).Site(site).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.CrawlSiteV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CrawlSiteV3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.CrawlSiteV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCrawlSiteV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **site** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeployTokenAddressV3

> map[string]interface{} DeployTokenAddressV3(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.DeployTokenAddressV3(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.DeployTokenAddressV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeployTokenAddressV3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.DeployTokenAddressV3`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeployTokenAddressV3Request struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountServicesV3

> map[string]interface{} GetAccountServicesV3(ctx).Phone(phone).Email(email).Wallet(wallet).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    phone := "phone_example" // string |  (optional)
    email := "email_example" // string |  (optional)
    wallet := "wallet_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetAccountServicesV3(context.Background()).Phone(phone).Email(email).Wallet(wallet).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetAccountServicesV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountServicesV3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetAccountServicesV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountServicesV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **phone** | **string** |  | 
 **email** | **string** |  | 
 **wallet** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountV2

> GetAccountResponseV2 GetAccountV2(ctx).Owner(owner).Execute()

Query user information



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "0x123456" // string | Ethereum address

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetAccountV2(context.Background()).Owner(owner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetAccountV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountV2`: GetAccountResponseV2
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetAccountV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | Ethereum address | 

### Return type

[**GetAccountResponseV2**](GetAccountResponseV2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountV3

> DexAccountV3 GetAccountV3(ctx).Owner(owner).AccountId(accountId).Execute()

Query user information



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "0x123456" // string | Ethereum address, either owner or accountId should be presented. (optional)
    accountId := int64(10003) // int64 | AccountID, if owner is presented, it must be align with the owners accountId, otherwise an error occurs. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetAccountV3(context.Background()).Owner(owner).AccountId(accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetAccountV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountV3`: DexAccountV3
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetAccountV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | Ethereum address, either owner or accountId should be presented. | 
 **accountId** | **int64** | AccountID, if owner is presented, it must be align with the owners accountId, otherwise an error occurs. | [default to 0]

### Return type

[**DexAccountV3**](DexAccountV3.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActivityRulesV3

> map[string]interface{} GetActivityRulesV3(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetActivityRulesV3(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetActivityRulesV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetActivityRulesV3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetActivityRulesV3`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetActivityRulesV3Request struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllowanceV3

> GetAllowancesV2Response GetAllowanceV3(ctx).Owner(owner).Token(token).Execute()

Get user token allowance



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "0x0000000000000000000000000000000000000000" // string | Ethereum address
    token := "0x0000000000000000000000000000000000000000" // string | The token address to query

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetAllowanceV3(context.Background()).Owner(owner).Token(token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetAllowanceV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllowanceV3`: GetAllowancesV2Response
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetAllowanceV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllowanceV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | Ethereum address | 
 **token** | **string** | The token address to query | 

### Return type

[**GetAllowancesV2Response**](GetAllowancesV2Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAmmAssetsV3

> map[string]interface{} GetAmmAssetsV3(ctx).PoolAddress(poolAddress).Limit(limit).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    poolAddress := "poolAddress_example" // string |  (optional)
    limit := int32(56) // int32 |  (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetAmmAssetsV3(context.Background()).PoolAddress(poolAddress).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetAmmAssetsV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAmmAssetsV3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetAmmAssetsV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAmmAssetsV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **poolAddress** | **string** |  | 
 **limit** | **int32** |  | [default to 0]

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAmmDepthV2

> GetDepthResponseV2 GetAmmDepthV2(ctx).PoolAddress(poolAddress).Level(level).Limit(limit).Execute()

Get market orderbook



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    poolAddress := "0xbbbbca6a901c926f240b89eacb641d8aec7aeafd" // string | request.getAmmDepth.market
    level := int32(2) // int32 | Order book aggregation level, larger value means further price aggregation.
    limit := int32(50) // int32 | Maximum numbers of bids/asks. (optional) (default to 50)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetAmmDepthV2(context.Background()).PoolAddress(poolAddress).Level(level).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetAmmDepthV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAmmDepthV2`: GetDepthResponseV2
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetAmmDepthV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAmmDepthV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **poolAddress** | **string** | request.getAmmDepth.market | 
 **level** | **int32** | Order book aggregation level, larger value means further price aggregation. | 
 **limit** | **int32** | Maximum numbers of bids/asks. | [default to 50]

### Return type

[**GetDepthResponseV2**](GetDepthResponseV2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAmmDepthV3

> DepthV3 GetAmmDepthV3(ctx).PoolAddress(poolAddress).Level(level).Limit(limit).Execute()

Get market orderbook



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    poolAddress := "0xbbbbca6a901c926f240b89eacb641d8aec7aeafd" // string | request.getAmmDepth.market
    level := int32(2) // int32 | Order book aggregation level, larger value means further price aggregation.
    limit := int32(50) // int32 | Maximum numbers of bids/asks. (optional) (default to 50)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetAmmDepthV3(context.Background()).PoolAddress(poolAddress).Level(level).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetAmmDepthV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAmmDepthV3`: DepthV3
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetAmmDepthV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAmmDepthV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **poolAddress** | **string** | request.getAmmDepth.market | 
 **level** | **int32** | Order book aggregation level, larger value means further price aggregation. | 
 **limit** | **int32** | Maximum numbers of bids/asks. | [default to 50]

### Return type

[**DepthV3**](DepthV3.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAmmMarketsV2

> GetAmmMarketInfoResponse GetAmmMarketsV2(ctx).Execute()

api.getAmmMarkets.value



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetAmmMarketsV2(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetAmmMarketsV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAmmMarketsV2`: GetAmmMarketInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetAmmMarketsV2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAmmMarketsV2Request struct via the builder pattern


### Return type

[**GetAmmMarketInfoResponse**](GetAmmMarketInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAmmPoolBalanceV3

> AmmPoolBalanceV3 GetAmmPoolBalanceV3(ctx).PoolAddress(poolAddress).Execute()

Get AMM pool balance snapshot



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    poolAddress := "0xbbbbca6a901c926f240b89eacb641d8aec7aeafd" // string | input AMM pool address

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetAmmPoolBalanceV3(context.Background()).PoolAddress(poolAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetAmmPoolBalanceV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAmmPoolBalanceV3`: AmmPoolBalanceV3
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetAmmPoolBalanceV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAmmPoolBalanceV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **poolAddress** | **string** | input AMM pool address | 

### Return type

[**AmmPoolBalanceV3**](AmmPoolBalanceV3.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAmmPoolsBalancesV3

> AmmPoolBalanceV3 GetAmmPoolsBalancesV3(ctx).Execute()

api.getAmmPoolsBalances.value



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetAmmPoolsBalancesV3(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetAmmPoolsBalancesV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAmmPoolsBalancesV3`: AmmPoolBalanceV3
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetAmmPoolsBalancesV3`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAmmPoolsBalancesV3Request struct via the builder pattern


### Return type

[**AmmPoolBalanceV3**](AmmPoolBalanceV3.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAmmPoolsV3

> GetAmmPoolsResponse GetAmmPoolsV3(ctx).Execute()

Get AMM pool configurations



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetAmmPoolsV3(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetAmmPoolsV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAmmPoolsV3`: GetAmmPoolsResponse
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetAmmPoolsV3`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAmmPoolsV3Request struct via the builder pattern


### Return type

[**GetAmmPoolsResponse**](GetAmmPoolsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAmmSnapshotV2

> GetAmmSnapshotResponse GetAmmSnapshotV2(ctx).PoolAddress(poolAddress).Execute()

api.getAmmSnapshot.value



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    poolAddress := "0xbbbbca6a901c926f240b89eacb641d8aec7aeafd" // string | request.getAmmSnapshot.address

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetAmmSnapshotV2(context.Background()).PoolAddress(poolAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetAmmSnapshotV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAmmSnapshotV2`: GetAmmSnapshotResponse
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetAmmSnapshotV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAmmSnapshotV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **poolAddress** | **string** | request.getAmmSnapshot.address | 

### Return type

[**GetAmmSnapshotResponse**](GetAmmSnapshotResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAmmSnapshotsV2

> GetAmmSnapshotResponse GetAmmSnapshotsV2(ctx).Execute()

api.getAmmSnapshots.value



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetAmmSnapshotsV2(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetAmmSnapshotsV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAmmSnapshotsV2`: GetAmmSnapshotResponse
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetAmmSnapshotsV2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAmmSnapshotsV2Request struct via the builder pattern


### Return type

[**GetAmmSnapshotResponse**](GetAmmSnapshotResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAmmTradesV2

> GetAmmTradesResponseV2 GetAmmTradesV2(ctx).AmmPoolAddress(ammPoolAddress).Limit(limit).Offset(offset).Execute()

get AMM pool trade transactions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ammPoolAddress := "ammPoolAddress_example" // string |  (optional)
    limit := int32(56) // int32 |  (optional) (default to 0)
    offset := int64(789) // int64 |  (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetAmmTradesV2(context.Background()).AmmPoolAddress(ammPoolAddress).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetAmmTradesV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAmmTradesV2`: GetAmmTradesResponseV2
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetAmmTradesV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAmmTradesV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ammPoolAddress** | **string** |  | 
 **limit** | **int32** |  | [default to 0]
 **offset** | **int64** |  | [default to 0]

### Return type

[**GetAmmTradesResponseV2**](GetAmmTradesResponseV2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAmmTradesV3

> AmmTradeDataListV3 GetAmmTradesV3(ctx).AmmPoolAddress(ammPoolAddress).Limit(limit).Offset(offset).Execute()

get AMM pool trade transactions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ammPoolAddress := "0xbbbbca6a901c926f240b89eacb641d8aec7aeafd" // string | The address of the pool on which the swap was submitted.
    limit := int32(50) // int32 | Used to limit the number of returned records. Useful in implementing pagination. (optional) (default to 0)
    offset := int64(0) // int64 | Used to apply an offset when looking for valid records. Useful in implementing (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetAmmTradesV3(context.Background()).AmmPoolAddress(ammPoolAddress).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetAmmTradesV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAmmTradesV3`: AmmTradeDataListV3
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetAmmTradesV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAmmTradesV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ammPoolAddress** | **string** | The address of the pool on which the swap was submitted. | 
 **limit** | **int32** | Used to limit the number of returned records. Useful in implementing pagination. | [default to 0]
 **offset** | **int64** | Used to apply an offset when looking for valid records. Useful in implementing | [default to 0]

### Return type

[**AmmTradeDataListV3**](AmmTradeDataListV3.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAmmTransactionsV2

> GetAmmTransactionResponseV2 GetAmmTransactionsV2(ctx).PoolAddress(poolAddress).BillType(billType).Start(start).End(end).Limit(limit).Offset(offset).TokenId(tokenId).Income(income).TransferAddress(transferAddress).FromAddress(fromAddress).Execute()

get AMM involved join/exit transactions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    poolAddress := "poolAddress_example" // string |  (optional)
    billType := int32(56) // int32 |  (optional) (default to -1)
    start := int64(789) // int64 |  (optional) (default to 0)
    end := int64(789) // int64 |  (optional) (default to 0)
    limit := int32(56) // int32 |  (optional) (default to 20)
    offset := int32(56) // int32 |  (optional) (default to 0)
    tokenId := int32(56) // int32 |  (optional) (default to -1)
    income := true // bool |  (optional) (default to false)
    transferAddress := "transferAddress_example" // string |  (optional)
    fromAddress := "fromAddress_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetAmmTransactionsV2(context.Background()).PoolAddress(poolAddress).BillType(billType).Start(start).End(end).Limit(limit).Offset(offset).TokenId(tokenId).Income(income).TransferAddress(transferAddress).FromAddress(fromAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetAmmTransactionsV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAmmTransactionsV2`: GetAmmTransactionResponseV2
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetAmmTransactionsV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAmmTransactionsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **poolAddress** | **string** |  | 
 **billType** | **int32** |  | [default to -1]
 **start** | **int64** |  | [default to 0]
 **end** | **int64** |  | [default to 0]
 **limit** | **int32** |  | [default to 20]
 **offset** | **int32** |  | [default to 0]
 **tokenId** | **int32** |  | [default to -1]
 **income** | **bool** |  | [default to false]
 **transferAddress** | **string** |  | 
 **fromAddress** | **string** |  | 

### Return type

[**GetAmmTransactionResponseV2**](GetAmmTransactionResponseV2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAmmTransactionsV3

> UserBillList GetAmmTransactionsV3(ctx).PoolAddress(poolAddress).Start(start).End(end).Limit(limit).Offset(offset).TxTypes(txTypes).TxStatus(txStatus).Execute()

get AMM involved join/exit transactions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    poolAddress := "0xbbbbca6a901c926f240b89eacb641d8aec7aeafd" // string | The address of the pool on which the swap was submitted.
    start := int64(1578558098000) // int64 | Date from which to start fetching AMM transactions. (optional) (default to 0)
    end := int64(1578558098000) // int64 | End Date of the query (optional) (default to 0)
    limit := int32(50) // int32 | Used to limit the number of returned records. Useful in implementing pagination. (optional) (default to 0)
    offset := int64(0) // int64 | Used to apply an offset when looking for valid records. Useful in implementing (optional) (default to 0)
    txTypes := "join,exit" // string | Transaction type: join or exit (optional)
    txStatus := "processing" // string | The AMM transaction status. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetAmmTransactionsV3(context.Background()).PoolAddress(poolAddress).Start(start).End(end).Limit(limit).Offset(offset).TxTypes(txTypes).TxStatus(txStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetAmmTransactionsV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAmmTransactionsV3`: UserBillList
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetAmmTransactionsV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAmmTransactionsV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **poolAddress** | **string** | The address of the pool on which the swap was submitted. | 
 **start** | **int64** | Date from which to start fetching AMM transactions. | [default to 0]
 **end** | **int64** | End Date of the query | [default to 0]
 **limit** | **int32** | Used to limit the number of returned records. Useful in implementing pagination. | [default to 0]
 **offset** | **int64** | Used to apply an offset when looking for valid records. Useful in implementing | [default to 0]
 **txTypes** | **string** | Transaction type: join or exit | 
 **txStatus** | **string** | The AMM transaction status. | 

### Return type

[**UserBillList**](UserBillList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAmmUserTransactionsV2

> GetAmmTransactionResponseV2 GetAmmUserTransactionsV2(ctx).AccountId(accountId).Start(start).End(end).Limit(limit).Offset(offset).TxTypes(txTypes).TxStatus(txStatus).AmmPoolAddress(ammPoolAddress).Execute()

User's AMM join/exit transactions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := int64(789) // int64 |  (optional)
    start := int64(789) // int64 |  (optional) (default to 0)
    end := int64(789) // int64 |  (optional) (default to 0)
    limit := int32(56) // int32 |  (optional) (default to 0)
    offset := int64(789) // int64 |  (optional) (default to 0)
    txTypes := "txTypes_example" // string |  (optional)
    txStatus := "txStatus_example" // string |  (optional)
    ammPoolAddress := "ammPoolAddress_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetAmmUserTransactionsV2(context.Background()).AccountId(accountId).Start(start).End(end).Limit(limit).Offset(offset).TxTypes(txTypes).TxStatus(txStatus).AmmPoolAddress(ammPoolAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetAmmUserTransactionsV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAmmUserTransactionsV2`: GetAmmTransactionResponseV2
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetAmmUserTransactionsV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAmmUserTransactionsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int64** |  | 
 **start** | **int64** |  | [default to 0]
 **end** | **int64** |  | [default to 0]
 **limit** | **int32** |  | [default to 0]
 **offset** | **int64** |  | [default to 0]
 **txTypes** | **string** |  | 
 **txStatus** | **string** |  | 
 **ammPoolAddress** | **string** |  | 

### Return type

[**GetAmmTransactionResponseV2**](GetAmmTransactionResponseV2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAmmUserTransactionsV3

> AmmTransactionDataListV3 GetAmmUserTransactionsV3(ctx).AccountId(accountId).Start(start).End(end).Limit(limit).Offset(offset).TxTypes(txTypes).TxStatus(txStatus).AmmPoolAddress(ammPoolAddress).Hashes(hashes).Execute()

User's AMM join/exit transactions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := int64(1578558098000) // int64 | Looprings account identifier. (optional)
    start := int64(1578558098000) // int64 | Date from which to start fetching AMM transactions. (optional) (default to 0)
    end := int64(1578558098000) // int64 | End Date of the query (optional) (default to 0)
    limit := int32(50) // int32 | Used to limit the number of returned records. Useful in implementing pagination. (optional) (default to 0)
    offset := int64(0) // int64 | Used to apply an offset when looking for valid records. Useful in implementing (optional) (default to 0)
    txTypes := "0" // string | Transaction type: join or exit (optional)
    txStatus := "0" // string | The AMM transaction status. (optional)
    ammPoolAddress := "0" // string | The address of the pool on which the swap was submitted. (optional)
    hashes := "0" // string | request.getAmmTransactions.hashes (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetAmmUserTransactionsV3(context.Background()).AccountId(accountId).Start(start).End(end).Limit(limit).Offset(offset).TxTypes(txTypes).TxStatus(txStatus).AmmPoolAddress(ammPoolAddress).Hashes(hashes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetAmmUserTransactionsV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAmmUserTransactionsV3`: AmmTransactionDataListV3
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetAmmUserTransactionsV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAmmUserTransactionsV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int64** | Looprings account identifier. | 
 **start** | **int64** | Date from which to start fetching AMM transactions. | [default to 0]
 **end** | **int64** | End Date of the query | [default to 0]
 **limit** | **int32** | Used to limit the number of returned records. Useful in implementing pagination. | [default to 0]
 **offset** | **int64** | Used to apply an offset when looking for valid records. Useful in implementing | [default to 0]
 **txTypes** | **string** | Transaction type: join or exit | 
 **txStatus** | **string** | The AMM transaction status. | 
 **ammPoolAddress** | **string** | The address of the pool on which the swap was submitted. | 
 **hashes** | **string** | request.getAmmTransactions.hashes | 

### Return type

[**AmmTransactionDataListV3**](AmmTransactionDataListV3.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiKeyV2

> GetApiKeyResponseV2 GetApiKeyV2(ctx).AccountId(accountId).Execute()

Get user ApiKey



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := int64(10) // int64 | AccountID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetApiKeyV2(context.Background()).AccountId(accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetApiKeyV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiKeyV2`: GetApiKeyResponseV2
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetApiKeyV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiKeyV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int64** | AccountID | 

### Return type

[**GetApiKeyResponseV2**](GetApiKeyResponseV2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiKeyV3

> GetApiKeyResponseV3 GetApiKeyV3(ctx).AccountId(accountId).Execute()

Get user ApiKey



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := int64(10) // int64 | AccountID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetApiKeyV3(context.Background()).AccountId(accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetApiKeyV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiKeyV3`: GetApiKeyResponseV3
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetApiKeyV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiKeyV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int64** | AccountID | 

### Return type

[**GetApiKeyResponseV3**](GetApiKeyResponseV3.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAvailableBrokerV3

> map[string]interface{} GetAvailableBrokerV3(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetAvailableBrokerV3(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetAvailableBrokerV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAvailableBrokerV3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetAvailableBrokerV3`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAvailableBrokerV3Request struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAvailableBrokerV3Hebao

> map[string]interface{} GetAvailableBrokerV3Hebao(ctx).Organization(organization).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organization := "organization_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetAvailableBrokerV3Hebao(context.Background()).Organization(organization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetAvailableBrokerV3Hebao``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAvailableBrokerV3Hebao`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetAvailableBrokerV3Hebao`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAvailableBrokerV3HebaoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organization** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockV3

> BlockResp GetBlockV3(ctx).Id(id).Execute()

Get L2 block info



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "1234" // string | The block id, could be finalized, confirmed, or block_idx_num (optional) (default to "finalized")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetBlockV3(context.Background()).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetBlockV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBlockV3`: BlockResp
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetBlockV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The block id, could be finalized, confirmed, or block_idx_num | [default to &quot;finalized&quot;]

### Return type

[**BlockResp**](BlockResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBusinessFee2V2

> GetOffchainFee2Response GetBusinessFee2V2(ctx).AccountId(accountId).RequestType(requestType).TokenSymbol(tokenSymbol).Amount(amount).Execute()

Query current fee amount



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := int32(0) // int32 | Account ID
    requestType := int32(1) // int32 | Off-chain request type
    tokenSymbol := "LRC" // string | The token to withdraw (optional)
    amount := "10000000000" // string | The amount to withdraw (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetBusinessFee2V2(context.Background()).AccountId(accountId).RequestType(requestType).TokenSymbol(tokenSymbol).Amount(amount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetBusinessFee2V2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBusinessFee2V2`: GetOffchainFee2Response
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetBusinessFee2V2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBusinessFee2V2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32** | Account ID | 
 **requestType** | **int32** | Off-chain request type | 
 **tokenSymbol** | **string** | The token to withdraw | 
 **amount** | **string** | The amount to withdraw | 

### Return type

[**GetOffchainFee2Response**](GetOffchainFee2Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBusinessFee2V3

> OffFeeInfo2 GetBusinessFee2V3(ctx).AccountId(accountId).RequestType(requestType).TokenSymbol(tokenSymbol).Amount(amount).Market(market).Execute()

Query current fee amount



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := int32(0) // int32 | Account ID
    requestType := int32(1) // int32 | Off-chain request type
    tokenSymbol := "LRC" // string | The token to withdraw (optional)
    amount := "10000000000" // string | The amount to withdraw (optional)
    market := "LRC-USDT" // string | The market (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetBusinessFee2V3(context.Background()).AccountId(accountId).RequestType(requestType).TokenSymbol(tokenSymbol).Amount(amount).Market(market).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetBusinessFee2V3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBusinessFee2V3`: OffFeeInfo2
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetBusinessFee2V3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBusinessFee2V3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32** | Account ID | 
 **requestType** | **int32** | Off-chain request type | 
 **tokenSymbol** | **string** | The token to withdraw | 
 **amount** | **string** | The amount to withdraw | 
 **market** | **string** | The market | 

### Return type

[**OffFeeInfo2**](OffFeeInfo2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCandlestickV2

> GetCandlestickResponseV2 GetCandlestickV2(ctx).Market(market).Interval(interval).Start(start).End(end).Limit(limit).Execute()

Get market candlestick



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    market := "LRC-ETH" // string | Trading pair ID, multi-market is not supported
    interval := "5min" // string | Candlestick interval, Supported values are: 1min, 5min, 15min, 30min, 1hr, 2hr, 4hr, 12hr, 1d, 1w
    start := int64(1584479105000) // int64 | Start time (optional) (default to 0)
    end := int64(1584565505000) // int64 | End time (optional) (default to 0)
    limit := int32(120) // int32 | Number of data points. If more data points are available, the API will only return the first 'limit' data points. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetCandlestickV2(context.Background()).Market(market).Interval(interval).Start(start).End(end).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetCandlestickV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCandlestickV2`: GetCandlestickResponseV2
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetCandlestickV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCandlestickV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **market** | **string** | Trading pair ID, multi-market is not supported | 
 **interval** | **string** | Candlestick interval, Supported values are: 1min, 5min, 15min, 30min, 1hr, 2hr, 4hr, 12hr, 1d, 1w | 
 **start** | **int64** | Start time | [default to 0]
 **end** | **int64** | End time | [default to 0]
 **limit** | **int32** | Number of data points. If more data points are available, the API will only return the first &#39;limit&#39; data points. | [default to 0]

### Return type

[**GetCandlestickResponseV2**](GetCandlestickResponseV2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCandlestickV3

> GetCandlestickResponseV3 GetCandlestickV3(ctx).Market(market).Interval(interval).Start(start).End(end).Limit(limit).Execute()

Get market candlestick



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    market := "LRC-ETH" // string | Trading pair ID, multi-market is not supported
    interval := "5min" // string | Candlestick interval, Supported values are: 1min, 5min, 15min, 30min, 1hr, 2hr, 4hr, 12hr, 1d, 1w
    start := int64(1584479105000) // int64 | Start time (optional) (default to 0)
    end := int64(1584565505000) // int64 | End time (optional) (default to 0)
    limit := int32(120) // int32 | Number of data points. If more data points are available, the API will only return the first 'limit' data points. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetCandlestickV3(context.Background()).Market(market).Interval(interval).Start(start).End(end).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetCandlestickV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCandlestickV3`: GetCandlestickResponseV3
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetCandlestickV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCandlestickV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **market** | **string** | Trading pair ID, multi-market is not supported | 
 **interval** | **string** | Candlestick interval, Supported values are: 1min, 5min, 15min, 30min, 1hr, 2hr, 4hr, 12hr, 1d, 1w | 
 **start** | **int64** | Start time | [default to 0]
 **end** | **int64** | End time | [default to 0]
 **limit** | **int32** | Number of data points. If more data points are available, the API will only return the first &#39;limit&#39; data points. | [default to 0]

### Return type

[**GetCandlestickResponseV3**](GetCandlestickResponseV3.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClaimHistoryV3

> map[string]interface{} GetClaimHistoryV3(ctx).FromId(fromId).Limit(limit).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    fromId := int64(789) // int64 |  (optional) (default to 0)
    limit := int32(56) // int32 |  (optional) (default to 50)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetClaimHistoryV3(context.Background()).FromId(fromId).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetClaimHistoryV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClaimHistoryV3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetClaimHistoryV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetClaimHistoryV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fromId** | **int64** |  | [default to 0]
 **limit** | **int32** |  | [default to 50]

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClaimLuckyTokenRecordsV3

> map[string]interface{} GetClaimLuckyTokenRecordsV3(ctx).Hash(hash).FromId(fromId).Limit(limit).ShowHelper(showHelper).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    hash := "hash_example" // string |  (optional)
    fromId := int64(789) // int64 |  (optional) (default to 0)
    limit := int32(56) // int32 |  (optional) (default to 50)
    showHelper := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetClaimLuckyTokenRecordsV3(context.Background()).Hash(hash).FromId(fromId).Limit(limit).ShowHelper(showHelper).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetClaimLuckyTokenRecordsV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClaimLuckyTokenRecordsV3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetClaimLuckyTokenRecordsV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetClaimLuckyTokenRecordsV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hash** | **string** |  | 
 **fromId** | **int64** |  | [default to 0]
 **limit** | **int32** |  | [default to 50]
 **showHelper** | **bool** |  | [default to false]

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClaimLuckyTokenV3

> map[string]interface{} GetClaimLuckyTokenV3(ctx).Hashes(hashes).FromId(fromId).Limit(limit).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    hashes := "hashes_example" // string |  (optional)
    fromId := int64(789) // int64 |  (optional) (default to 0)
    limit := int32(56) // int32 |  (optional) (default to 50)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetClaimLuckyTokenV3(context.Background()).Hashes(hashes).FromId(fromId).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetClaimLuckyTokenV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClaimLuckyTokenV3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetClaimLuckyTokenV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetClaimLuckyTokenV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hashes** | **string** |  | 
 **fromId** | **int64** |  | [default to 0]
 **limit** | **int32** |  | [default to 50]

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommissionRewardRankV3

> map[string]interface{} GetCommissionRewardRankV3(ctx).TokenId(tokenId).Top(top).RewardType(rewardType).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tokenId := int32(56) // int32 |  (optional)
    top := int32(56) // int32 |  (optional)
    rewardType := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetCommissionRewardRankV3(context.Background()).TokenId(tokenId).Top(top).RewardType(rewardType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetCommissionRewardRankV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCommissionRewardRankV3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetCommissionRewardRankV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCommissionRewardRankV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenId** | **int32** |  | 
 **top** | **int32** |  | 
 **rewardType** | **int32** |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommissionRewardV3

> map[string]interface{} GetCommissionRewardV3(ctx).AccountId(accountId).TokenId(tokenId).Start(start).Size(size).RewardType(rewardType).Taker(taker).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := int64(789) // int64 |  (optional)
    tokenId := int32(56) // int32 |  (optional)
    start := int64(789) // int64 |  (optional)
    size := int32(56) // int32 |  (optional)
    rewardType := int32(56) // int32 |  (optional)
    taker := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetCommissionRewardV3(context.Background()).AccountId(accountId).TokenId(tokenId).Start(start).Size(size).RewardType(rewardType).Taker(taker).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetCommissionRewardV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCommissionRewardV3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetCommissionRewardV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCommissionRewardV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int64** |  | 
 **tokenId** | **int32** |  | 
 **start** | **int64** |  | 
 **size** | **int32** |  | 
 **rewardType** | **int32** |  | 
 **taker** | **int32** |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommissionTotalRewardV3

> map[string]interface{} GetCommissionTotalRewardV3(ctx).AccountId(accountId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := int64(789) // int64 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetCommissionTotalRewardV3(context.Background()).AccountId(accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetCommissionTotalRewardV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCommissionTotalRewardV3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetCommissionTotalRewardV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCommissionTotalRewardV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int64** |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCounterFactualInfoV3

> AccountCounterFactualInfo GetCounterFactualInfoV3(ctx).AccountId(accountId).Execute()

api.getCounterFactualInfo.value



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := "accountId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetCounterFactualInfoV3(context.Background()).AccountId(accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetCounterFactualInfoV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCounterFactualInfoV3`: AccountCounterFactualInfo
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetCounterFactualInfoV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCounterFactualInfoV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** |  | 

### Return type

[**AccountCounterFactualInfo**](AccountCounterFactualInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrencyRatesV2

> map[string]interface{} GetCurrencyRatesV2(ctx).Currencies(currencies).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    currencies := "USD,CNY" // string | Currency symbol to be queried

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetCurrencyRatesV2(context.Background()).Currencies(currencies).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetCurrencyRatesV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCurrencyRatesV2`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetCurrencyRatesV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrencyRatesV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currencies** | **string** | Currency symbol to be queried | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDaoSquareAccountsV3

> map[string]interface{} GetDaoSquareAccountsV3(ctx).StartTime(startTime).EndTime(endTime).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    startTime := int64(789) // int64 |  (optional) (default to 0)
    endTime := int64(789) // int64 |  (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetDaoSquareAccountsV3(context.Background()).StartTime(startTime).EndTime(endTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetDaoSquareAccountsV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDaoSquareAccountsV3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetDaoSquareAccountsV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDaoSquareAccountsV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTime** | **int64** |  | [default to 0]
 **endTime** | **int64** |  | [default to 0]

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDepthV2

> GetDepthResponseV2 GetDepthV2(ctx).Market(market).Level(level).Limit(limit).Execute()

Get market orderbook



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    market := "LRC-ETH" // string | The ID of a trading pair.
    level := int32(2) // int32 | Order book aggregation level, larger value means further price aggregation.
    limit := int32(50) // int32 | Maximum numbers of bids/asks. (optional) (default to 50)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetDepthV2(context.Background()).Market(market).Level(level).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetDepthV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDepthV2`: GetDepthResponseV2
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetDepthV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDepthV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **market** | **string** | The ID of a trading pair. | 
 **level** | **int32** | Order book aggregation level, larger value means further price aggregation. | 
 **limit** | **int32** | Maximum numbers of bids/asks. | [default to 50]

### Return type

[**GetDepthResponseV2**](GetDepthResponseV2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDepthV3

> DepthV3 GetDepthV3(ctx).Market(market).Level(level).Limit(limit).Execute()

Get market orderbook



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    market := "LRC-ETH" // string | The ID of a trading pair.
    level := int32(2) // int32 | Order book aggregation level, larger value means further price aggregation.
    limit := int32(50) // int32 | Maximum numbers of bids/asks. (optional) (default to 50)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetDepthV3(context.Background()).Market(market).Level(level).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetDepthV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDepthV3`: DepthV3
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetDepthV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDepthV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **market** | **string** | The ID of a trading pair. | 
 **level** | **int32** | Order book aggregation level, larger value means further price aggregation. | 
 **limit** | **int32** | Maximum numbers of bids/asks. | [default to 50]

### Return type

[**DepthV3**](DepthV3.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEthBalanceV3

> GetEthBalancesV2Response GetEthBalanceV3(ctx).Owner(owner).Execute()

Get user's Ether balance on Ethereum mainnet



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "0x0000000000000000000000000000000000000000" // string | Ethereum address

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetEthBalanceV3(context.Background()).Owner(owner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetEthBalanceV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEthBalanceV3`: GetEthBalancesV2Response
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetEthBalanceV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEthBalanceV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | Ethereum address | 

### Return type

[**GetEthBalancesV2Response**](GetEthBalancesV2Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEthNonceV3

> GetEthNonceV2Response GetEthNonceV3(ctx).Owner(owner).Execute()

Get user's next Ethereum nonce



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "0x0000000000000000000000000000000000000000" // string | Ethereum address

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetEthNonceV3(context.Background()).Owner(owner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetEthNonceV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEthNonceV3`: GetEthNonceV2Response
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetEthNonceV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEthNonceV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | Ethereum address | 

### Return type

[**GetEthNonceV2Response**](GetEthNonceV2Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExchangeFeeInfoV2

> GetExchangeFeeInfoResponse GetExchangeFeeInfoV2(ctx).Execute()

Get exchange configurations



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetExchangeFeeInfoV2(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetExchangeFeeInfoV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExchangeFeeInfoV2`: GetExchangeFeeInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetExchangeFeeInfoV2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetExchangeFeeInfoV2Request struct via the builder pattern


### Return type

[**GetExchangeFeeInfoResponse**](GetExchangeFeeInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExchangeFeeInfoV3

> GetExchangeFeeInfoResponseData GetExchangeFeeInfoV3(ctx).Execute()

Get exchange configurations



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetExchangeFeeInfoV3(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetExchangeFeeInfoV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExchangeFeeInfoV3`: GetExchangeFeeInfoResponseData
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetExchangeFeeInfoV3`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetExchangeFeeInfoV3Request struct via the builder pattern


### Return type

[**GetExchangeFeeInfoResponseData**](GetExchangeFeeInfoResponseData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExchangeInfoV2

> GetExchangeInfoResponse GetExchangeInfoV2(ctx).Execute()

Get exchange configurations



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetExchangeInfoV2(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetExchangeInfoV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExchangeInfoV2`: GetExchangeInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetExchangeInfoV2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetExchangeInfoV2Request struct via the builder pattern


### Return type

[**GetExchangeInfoResponse**](GetExchangeInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExchangeInfoV3

> ExchangeInfo GetExchangeInfoV3(ctx).Execute()

Get exchange configurations



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetExchangeInfoV3(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetExchangeInfoV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExchangeInfoV3`: ExchangeInfo
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetExchangeInfoV3`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetExchangeInfoV3Request struct via the builder pattern


### Return type

[**ExchangeInfo**](ExchangeInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceIncomeV3

> map[string]interface{} GetFinanceIncomeV3(ctx).Address(address).FinanceType(financeType).TokenAddress(tokenAddress).Offset(offset).Limit(limit).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    address := "address_example" // string |  (optional)
    financeType := "financeType_example" // string |  (optional)
    tokenAddress := "tokenAddress_example" // string |  (optional)
    offset := int64(789) // int64 |  (optional) (default to 0)
    limit := int32(56) // int32 |  (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetFinanceIncomeV3(context.Background()).Address(address).FinanceType(financeType).TokenAddress(tokenAddress).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetFinanceIncomeV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFinanceIncomeV3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetFinanceIncomeV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceIncomeV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **address** | **string** |  | 
 **financeType** | **string** |  | 
 **tokenAddress** | **string** |  | 
 **offset** | **int64** |  | [default to 0]
 **limit** | **int32** |  | [default to 0]

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLatestTokenPricesV3

> map[string]interface{} GetLatestTokenPricesV3(ctx).Tokens(tokens).Currency(currency).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tokens := "tokens_example" // string |  (optional)
    currency := "currency_example" // string |  (optional) (default to "USD")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetLatestTokenPricesV3(context.Background()).Tokens(tokens).Currency(currency).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetLatestTokenPricesV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLatestTokenPricesV3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetLatestTokenPricesV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLatestTokenPricesV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokens** | **string** |  | 
 **currency** | **string** |  | [default to &quot;USD&quot;]

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLiquidityMiningConfV3

> map[string]interface{} GetLiquidityMiningConfV3(ctx).Market(market).Running(running).PageIndex(pageIndex).Size(size).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    market := "market_example" // string |  (optional)
    running := true // bool |  (optional) (default to false)
    pageIndex := int32(56) // int32 |  (optional)
    size := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetLiquidityMiningConfV3(context.Background()).Market(market).Running(running).PageIndex(pageIndex).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetLiquidityMiningConfV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLiquidityMiningConfV3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetLiquidityMiningConfV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLiquidityMiningConfV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **market** | **string** |  | 
 **running** | **bool** |  | [default to false]
 **pageIndex** | **int32** |  | 
 **size** | **int32** |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLiquidityMiningRankV3

> map[string]interface{} GetLiquidityMiningRankV3(ctx).Market(market).Top(top).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    market := "market_example" // string |  (optional)
    top := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetLiquidityMiningRankV3(context.Background()).Market(market).Top(top).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetLiquidityMiningRankV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLiquidityMiningRankV3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetLiquidityMiningRankV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLiquidityMiningRankV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **market** | **string** |  | 
 **top** | **int32** |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLiquidityMiningRewardTotalV3

> map[string]interface{} GetLiquidityMiningRewardTotalV3(ctx).AccountId(accountId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := int64(789) // int64 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetLiquidityMiningRewardTotalV3(context.Background()).AccountId(accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetLiquidityMiningRewardTotalV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLiquidityMiningRewardTotalV3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetLiquidityMiningRewardTotalV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLiquidityMiningRewardTotalV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int64** |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLiquidityMiningRewardsV3

> map[string]interface{} GetLiquidityMiningRewardsV3(ctx).AccountId(accountId).Market(market).Start(start).Timestamp(timestamp).Size(size).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := int64(789) // int64 |  (optional)
    market := "market_example" // string |  (optional)
    start := int64(789) // int64 |  (optional) (default to 0)
    timestamp := int64(789) // int64 |  (optional)
    size := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetLiquidityMiningRewardsV3(context.Background()).AccountId(accountId).Market(market).Start(start).Timestamp(timestamp).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetLiquidityMiningRewardsV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLiquidityMiningRewardsV3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetLiquidityMiningRewardsV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLiquidityMiningRewardsV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int64** |  | 
 **market** | **string** |  | 
 **start** | **int64** |  | [default to 0]
 **timestamp** | **int64** |  | 
 **size** | **int32** |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLiquidityMiningUserHistoryV3

> map[string]interface{} GetLiquidityMiningUserHistoryV3(ctx).AccountId(accountId).Markets(markets).Start(start).End(end).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := int64(789) // int64 |  (optional)
    markets := "markets_example" // string |  (optional)
    start := int64(789) // int64 |  (optional)
    end := int64(789) // int64 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetLiquidityMiningUserHistoryV3(context.Background()).AccountId(accountId).Markets(markets).Start(start).End(end).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetLiquidityMiningUserHistoryV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLiquidityMiningUserHistoryV3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetLiquidityMiningUserHistoryV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLiquidityMiningUserHistoryV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int64** |  | 
 **markets** | **string** |  | 
 **start** | **int64** |  | 
 **end** | **int64** |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLuckyTokenAgentsV3

> GetLuckyTokenAgentsResponse GetLuckyTokenAgentsV3(ctx).Execute()

api.getLuckyTokenAgents.value



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetLuckyTokenAgentsV3(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetLuckyTokenAgentsV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLuckyTokenAgentsV3`: GetLuckyTokenAgentsResponse
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetLuckyTokenAgentsV3`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLuckyTokenAgentsV3Request struct via the builder pattern


### Return type

[**GetLuckyTokenAgentsResponse**](GetLuckyTokenAgentsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLuckyTokenSignersV3

> GetLuckyTokenAgentsResponse GetLuckyTokenSignersV3(ctx).Execute()

api.getLuckyTokenSigners.value



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetLuckyTokenSignersV3(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetLuckyTokenSignersV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLuckyTokenSignersV3`: GetLuckyTokenAgentsResponse
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetLuckyTokenSignersV3`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLuckyTokenSignersV3Request struct via the builder pattern


### Return type

[**GetLuckyTokenAgentsResponse**](GetLuckyTokenAgentsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLuckyTokenSummaryV3

> map[string]interface{} GetLuckyTokenSummaryV3(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetLuckyTokenSummaryV3(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetLuckyTokenSummaryV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLuckyTokenSummaryV3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetLuckyTokenSummaryV3`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLuckyTokenSummaryV3Request struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLuckyTokenUserBalancesV3

> BalanceV3 GetLuckyTokenUserBalancesV3(ctx).AccountId(accountId).Tokens(tokens).Execute()

api.getLuckyTokenBalance.value



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := int32(1) // int32 | AccountID
    tokens := "0,1" // string | Query tokens (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetLuckyTokenUserBalancesV3(context.Background()).AccountId(accountId).Tokens(tokens).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetLuckyTokenUserBalancesV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLuckyTokenUserBalancesV3`: BalanceV3
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetLuckyTokenUserBalancesV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLuckyTokenUserBalancesV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32** | AccountID | 
 **tokens** | **string** | Query tokens | 

### Return type

[**BalanceV3**](BalanceV3.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLuckyTokenV3

> map[string]interface{} GetLuckyTokenV3(ctx).SenderId(senderId).Hash(hash).Partitions(partitions).Nodes(nodes).Scopes(scopes).Statuses(statuses).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).Official(official).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    senderId := int64(789) // int64 |  (optional) (default to 0)
    hash := "hash_example" // string |  (optional)
    partitions := "partitions_example" // string |  (optional)
    nodes := "nodes_example" // string |  (optional)
    scopes := "scopes_example" // string |  (optional)
    statuses := "statuses_example" // string |  (optional)
    startTime := int64(789) // int64 |  (optional) (default to 0)
    endTime := int64(789) // int64 |  (optional) (default to 0)
    fromId := int64(789) // int64 |  (optional) (default to 0)
    limit := int32(56) // int32 |  (optional) (default to 50)
    official := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetLuckyTokenV3(context.Background()).SenderId(senderId).Hash(hash).Partitions(partitions).Nodes(nodes).Scopes(scopes).Statuses(statuses).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).Official(official).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetLuckyTokenV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLuckyTokenV3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetLuckyTokenV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLuckyTokenV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **senderId** | **int64** |  | [default to 0]
 **hash** | **string** |  | 
 **partitions** | **string** |  | 
 **nodes** | **string** |  | 
 **scopes** | **string** |  | 
 **statuses** | **string** |  | 
 **startTime** | **int64** |  | [default to 0]
 **endTime** | **int64** |  | [default to 0]
 **fromId** | **int64** |  | [default to 0]
 **limit** | **int32** |  | [default to 50]
 **official** | **bool** |  | [default to false]

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLuckyTokenWithdrawV3

> map[string]interface{} GetLuckyTokenWithdrawV3(ctx).Statuses(statuses).TokenId(tokenId).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    statuses := "statuses_example" // string |  (optional)
    tokenId := int32(56) // int32 |  (optional) (default to -1)
    startTime := int64(789) // int64 |  (optional) (default to 0)
    endTime := int64(789) // int64 |  (optional) (default to 0)
    fromId := int64(789) // int64 |  (optional) (default to 0)
    limit := int32(56) // int32 |  (optional) (default to 50)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetLuckyTokenWithdrawV3(context.Background()).Statuses(statuses).TokenId(tokenId).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetLuckyTokenWithdrawV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLuckyTokenWithdrawV3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetLuckyTokenWithdrawV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLuckyTokenWithdrawV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **statuses** | **string** |  | 
 **tokenId** | **int32** |  | [default to -1]
 **startTime** | **int64** |  | [default to 0]
 **endTime** | **int64** |  | [default to 0]
 **fromId** | **int64** |  | [default to 0]
 **limit** | **int32** |  | [default to 50]

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketTradeV2

> GetMarketTradesV2Response GetMarketTradeV2(ctx).Market(market).Limit(limit).Execute()

Get market recent trades



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    market := "LRC-USDT" // string | Single market to query
    limit := int32(20) // int32 | Number of queries (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetMarketTradeV2(context.Background()).Market(market).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetMarketTradeV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMarketTradeV2`: GetMarketTradesV2Response
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetMarketTradeV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketTradeV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **market** | **string** | Single market to query | 
 **limit** | **int32** | Number of queries | [default to 0]

### Return type

[**GetMarketTradesV2Response**](GetMarketTradesV2Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketTradeV3

> MarketTradesV3 GetMarketTradeV3(ctx).Market(market).Limit(limit).FillTypes(fillTypes).Execute()

Get market recent trades



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    market := "LRC-USDT" // string | Single market to query
    limit := int32(20) // int32 | Number of queries (optional) (default to 0)
    fillTypes := "dex" // string | request.getUserTxs.fillTypes (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetMarketTradeV3(context.Background()).Market(market).Limit(limit).FillTypes(fillTypes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetMarketTradeV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMarketTradeV3`: MarketTradesV3
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetMarketTradeV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketTradeV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **market** | **string** | Single market to query | 
 **limit** | **int32** | Number of queries | [default to 0]
 **fillTypes** | **string** | request.getUserTxs.fillTypes | 

### Return type

[**MarketTradesV3**](MarketTradesV3.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketsV2

> GetMarketInfoResponse GetMarketsV2(ctx).Execute()

Get market configurations



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetMarketsV2(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetMarketsV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMarketsV2`: GetMarketInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetMarketsV2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketsV2Request struct via the builder pattern


### Return type

[**GetMarketInfoResponse**](GetMarketInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketsV3

> GetMarketsResponseV3 GetMarketsV3(ctx).Execute()

Get market configurations



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetMarketsV3(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetMarketsV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMarketsV3`: GetMarketsResponseV3
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetMarketsV3`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketsV3Request struct via the builder pattern


### Return type

[**GetMarketsResponseV3**](GetMarketsResponseV3.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMixedCandlestickV2

> GetCandlestickResponseV2 GetMixedCandlestickV2(ctx).Market(market).Interval(interval).Start(start).End(end).Limit(limit).Execute()

Get market candlestick



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    market := "LRC-ETH" // string | Trading pair ID, multi-market is not supported
    interval := "5min" // string | Candlestick interval, Supported values are: 1min, 5min, 15min, 30min, 1hr, 2hr, 4hr, 12hr, 1d, 1w
    start := int64(1584479105000) // int64 | Start time (optional) (default to 0)
    end := int64(1584565505000) // int64 | End time (optional) (default to 0)
    limit := int32(120) // int32 | Number of data points. If more data points are available, the API will only return the first 'limit' data points. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetMixedCandlestickV2(context.Background()).Market(market).Interval(interval).Start(start).End(end).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetMixedCandlestickV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMixedCandlestickV2`: GetCandlestickResponseV2
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetMixedCandlestickV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMixedCandlestickV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **market** | **string** | Trading pair ID, multi-market is not supported | 
 **interval** | **string** | Candlestick interval, Supported values are: 1min, 5min, 15min, 30min, 1hr, 2hr, 4hr, 12hr, 1d, 1w | 
 **start** | **int64** | Start time | [default to 0]
 **end** | **int64** | End time | [default to 0]
 **limit** | **int32** | Number of data points. If more data points are available, the API will only return the first &#39;limit&#39; data points. | [default to 0]

### Return type

[**GetCandlestickResponseV2**](GetCandlestickResponseV2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMixedCandlestickV3

> GetCandlestickResponseV3 GetMixedCandlestickV3(ctx).Market(market).Interval(interval).Start(start).End(end).Limit(limit).Execute()

Get mixed market candlestick



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    market := "LRC-ETH" // string | Trading pair ID, multi-market is not supported
    interval := "5min" // string | Candlestick interval, Supported values are: 1min, 5min, 15min, 30min, 1hr, 2hr, 4hr, 12hr, 1d, 1w
    start := int64(1584479105000) // int64 | Start time (optional) (default to 0)
    end := int64(1584565505000) // int64 | End time (optional) (default to 0)
    limit := int32(120) // int32 | Number of data points. If more data points are available, the API will only return the first 'limit' data points. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetMixedCandlestickV3(context.Background()).Market(market).Interval(interval).Start(start).End(end).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetMixedCandlestickV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMixedCandlestickV3`: GetCandlestickResponseV3
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetMixedCandlestickV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMixedCandlestickV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **market** | **string** | Trading pair ID, multi-market is not supported | 
 **interval** | **string** | Candlestick interval, Supported values are: 1min, 5min, 15min, 30min, 1hr, 2hr, 4hr, 12hr, 1d, 1w | 
 **start** | **int64** | Start time | [default to 0]
 **end** | **int64** | End time | [default to 0]
 **limit** | **int32** | Number of data points. If more data points are available, the API will only return the first &#39;limit&#39; data points. | [default to 0]

### Return type

[**GetCandlestickResponseV3**](GetCandlestickResponseV3.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMixedDepthV2

> GetDepthResponseV2 GetMixedDepthV2(ctx).Market(market).Level(level).Limit(limit).ShowOverlap(showOverlap).Execute()

Get market orderbook



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    market := "LRC-ETH" // string | The ID of a trading pair.
    level := int32(2) // int32 | Order book aggregation level, larger value means further price aggregation.
    limit := int32(50) // int32 | Maximum numbers of bids/asks. (optional) (default to 50)
    showOverlap := false // bool | request.getDepth.showOverlap (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetMixedDepthV2(context.Background()).Market(market).Level(level).Limit(limit).ShowOverlap(showOverlap).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetMixedDepthV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMixedDepthV2`: GetDepthResponseV2
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetMixedDepthV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMixedDepthV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **market** | **string** | The ID of a trading pair. | 
 **level** | **int32** | Order book aggregation level, larger value means further price aggregation. | 
 **limit** | **int32** | Maximum numbers of bids/asks. | [default to 50]
 **showOverlap** | **bool** | request.getDepth.showOverlap | [default to false]

### Return type

[**GetDepthResponseV2**](GetDepthResponseV2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMixedDepthV3

> DepthV3 GetMixedDepthV3(ctx).Market(market).Level(level).Limit(limit).ShowOverlap(showOverlap).Execute()

Get mixed market orderbook



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    market := "LRC-ETH" // string | The ID of a trading pair.
    level := int32(2) // int32 | Order book aggregation level, larger value means further price aggregation.
    limit := int32(50) // int32 | Maximum numbers of bids/asks. (optional) (default to 50)
    showOverlap := false // bool | request.getDepth.showOverlap (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetMixedDepthV3(context.Background()).Market(market).Level(level).Limit(limit).ShowOverlap(showOverlap).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetMixedDepthV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMixedDepthV3`: DepthV3
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetMixedDepthV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMixedDepthV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **market** | **string** | The ID of a trading pair. | 
 **level** | **int32** | Order book aggregation level, larger value means further price aggregation. | 
 **limit** | **int32** | Maximum numbers of bids/asks. | [default to 50]
 **showOverlap** | **bool** | request.getDepth.showOverlap | [default to false]

### Return type

[**DepthV3**](DepthV3.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMixedMarketsV2

> GetMixedMarketInfoResponse GetMixedMarketsV2(ctx).Execute()

Get market configurations



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetMixedMarketsV2(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetMixedMarketsV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMixedMarketsV2`: GetMixedMarketInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetMixedMarketsV2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMixedMarketsV2Request struct via the builder pattern


### Return type

[**GetMixedMarketInfoResponse**](GetMixedMarketInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMixedMarketsV3

> GetMixedMarketsResponseV3 GetMixedMarketsV3(ctx).Execute()

Get mixed market configurations



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetMixedMarketsV3(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetMixedMarketsV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMixedMarketsV3`: GetMixedMarketsResponseV3
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetMixedMarketsV3`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMixedMarketsV3Request struct via the builder pattern


### Return type

[**GetMixedMarketsResponseV3**](GetMixedMarketsResponseV3.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMixedTickerV2

> GetTickerResponseV2 GetMixedTickerV2(ctx).Market(market).Execute()

Get market ticker



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    market := "LRC-ETH,LRC-USDT" // string | Market pair, support multiple markets

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetMixedTickerV2(context.Background()).Market(market).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetMixedTickerV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMixedTickerV2`: GetTickerResponseV2
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetMixedTickerV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMixedTickerV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **market** | **string** | Market pair, support multiple markets | 

### Return type

[**GetTickerResponseV2**](GetTickerResponseV2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMixedTickerV3

> GetTickerResponseV3 GetMixedTickerV3(ctx).Market(market).Execute()

Get mixed market ticker



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    market := "LRC-ETH,LRC-USDT" // string | Market pair, support multiple markets

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetMixedTickerV3(context.Background()).Market(market).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetMixedTickerV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMixedTickerV3`: GetTickerResponseV3
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetMixedTickerV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMixedTickerV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **market** | **string** | Market pair, support multiple markets | 

### Return type

[**GetTickerResponseV3**](GetTickerResponseV3.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNewestTokenPricesV3

> map[string]interface{} GetNewestTokenPricesV3(ctx).Tokens(tokens).Currency(currency).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tokens := "tokens_example" // string |  (optional)
    currency := "currency_example" // string |  (optional) (default to "USD")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetNewestTokenPricesV3(context.Background()).Tokens(tokens).Currency(currency).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetNewestTokenPricesV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNewestTokenPricesV3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetNewestTokenPricesV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNewestTokenPricesV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokens** | **string** |  | 
 **currency** | **string** |  | [default to &quot;USD&quot;]

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNextStorageIdV2

> GetNextStorageIdResponseData GetNextStorageIdV2(ctx).AccountId(accountId).TokenSId(tokenSId).Execute()

Get next storage ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := int64(1) // int64 | Looprings account identifier
    tokenSId := int32(0) // int32 | request.getNextStorageId.tokenSId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetNextStorageIdV2(context.Background()).AccountId(accountId).TokenSId(tokenSId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetNextStorageIdV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNextStorageIdV2`: GetNextStorageIdResponseData
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetNextStorageIdV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNextStorageIdV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int64** | Looprings account identifier | 
 **tokenSId** | **int32** | request.getNextStorageId.tokenSId | 

### Return type

[**GetNextStorageIdResponseData**](GetNextStorageIdResponseData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNextStorageIdV3

> NextStorageIdResponse GetNextStorageIdV3(ctx).AccountId(accountId).SellTokenId(sellTokenId).MaxNext(maxNext).Execute()

Get next storage ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := int64(1) // int64 | Looprings account identifier
    sellTokenId := int32(0) // int32 | The unique identifier of the token which the user wants to sell in the next order.
    maxNext := false // bool | Return the max of the next available storageId, so any storageId > returned value is avaliable, to help user manage storageId by themselves. for example, if [20, 60, 100] is avaliable, all other ids < 100 is used before, user gets 20 if flag is false (and 60 in next run), but gets 100 if flag is true, so he can use 102, 104 freely (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetNextStorageIdV3(context.Background()).AccountId(accountId).SellTokenId(sellTokenId).MaxNext(maxNext).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetNextStorageIdV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNextStorageIdV3`: NextStorageIdResponse
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetNextStorageIdV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNextStorageIdV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int64** | Looprings account identifier | 
 **sellTokenId** | **int32** | The unique identifier of the token which the user wants to sell in the next order. | 
 **maxNext** | **bool** | Return the max of the next available storageId, so any storageId &gt; returned value is avaliable, to help user manage storageId by themselves. for example, if [20, 60, 100] is avaliable, all other ids &lt; 100 is used before, user gets 20 if flag is false (and 60 in next run), but gets 100 if flag is true, so he can use 102, 104 freely | [default to false]

### Return type

[**NextStorageIdResponse**](NextStorageIdResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNftDataV3

> NftTokenInfo GetNftDataV3(ctx).Minter(minter).TokenAddress(tokenAddress).NftId(nftId).Execute()

Query nftDatas by minter, tokenAddress and nftID.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    minter := "10001" // string | The minters accountId.
    tokenAddress := "0xbbbbca6a901c926f240b89eacb641d8aec7aeafd" // string | The NFT token address
    nftId := "0x01" // string | The NFT_ID of the NFT token

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetNftDataV3(context.Background()).Minter(minter).TokenAddress(tokenAddress).NftId(nftId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetNftDataV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNftDataV3`: NftTokenInfo
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetNftDataV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNftDataV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **minter** | **string** | The minters accountId. | 
 **tokenAddress** | **string** | The NFT token address | 
 **nftId** | **string** | The NFT_ID of the NFT token | 

### Return type

[**NftTokenInfo**](NftTokenInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNftDepositsV3

> GetUserNftDepositResponse GetNftDepositsV3(ctx).AccountId(accountId).NftData(nftData).Start(start).End(end).StartId(startId).Limit(limit).TxStatus(txStatus).Hashes(hashes).Execute()

Get user NFT deposit history



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := int64(10001) // int64 | users accountId.
    nftData := "0xf7c932351186c3a9053f313eefa16209c018f7f1dba8aa8ca7100400f7c31085" // string | the nftData of the NFT token (optional)
    start := int64(1567053142) // int64 | Start time in milliseconds (optional) (default to 0)
    end := int64(1567053142) // int64 | End time in milliseconds (optional) (default to 0)
    startId := int64(1234) // int64 | The begin id of the query. (optional) (default to 0)
    limit := int32(50) // int32 | Number of records to return (optional) (default to 0)
    txStatus := "processing" // string | The deposit status (optional)
    hashes := "0xf7c932351186c3a9053f313eefa16209c018f7f1dba8aa8ca7100400f7c31085" // string | The hashes (split by ,) of the transactions, normally its L2 tx hash, except the deposit which uses L1 tx hash. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetNftDepositsV3(context.Background()).AccountId(accountId).NftData(nftData).Start(start).End(end).StartId(startId).Limit(limit).TxStatus(txStatus).Hashes(hashes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetNftDepositsV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNftDepositsV3`: GetUserNftDepositResponse
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetNftDepositsV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNftDepositsV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int64** | users accountId. | 
 **nftData** | **string** | the nftData of the NFT token | 
 **start** | **int64** | Start time in milliseconds | [default to 0]
 **end** | **int64** | End time in milliseconds | [default to 0]
 **startId** | **int64** | The begin id of the query. | [default to 0]
 **limit** | **int32** | Number of records to return | [default to 0]
 **txStatus** | **string** | The deposit status | 
 **hashes** | **string** | The hashes (split by ,) of the transactions, normally its L2 tx hash, except the deposit which uses L1 tx hash. | 

### Return type

[**GetUserNftDepositResponse**](GetUserNftDepositResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNftHoldersV3

> NftHolders GetNftHoldersV3(ctx).NftData(nftData).Offset(offset).Limit(limit).Execute()

Query NFT holders by looprings nftData



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    nftData := "0xf7c932351186c3a9053f313eefa16209c018f7f1dba8aa8ca7100400f7c31085" // string | The Loopring's NFT token data identifier which is a hash string of NFT token address and NFT_ID
    offset := int32(0) // int32 | Number of records to skip (optional) (default to 0)
    limit := int32(100) // int32 | Number of records to return (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetNftHoldersV3(context.Background()).NftData(nftData).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetNftHoldersV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNftHoldersV3`: NftHolders
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetNftHoldersV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNftHoldersV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nftData** | **string** | The Loopring&#39;s NFT token data identifier which is a hash string of NFT token address and NFT_ID | 
 **offset** | **int32** | Number of records to skip | [default to 0]
 **limit** | **int32** | Number of records to return | [default to 0]

### Return type

[**NftHolders**](NftHolders.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNftMintersV3

> GetNftMintersResponse GetNftMintersV3(ctx).Channel(channel).AuthContractAddress(authContractAddress).Execute()

Get minters



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    channel := int32(1) // int32 | channel id, if you dont know your channel id, please contact with Loopring team.
    authContractAddress := "authContractAddress_example" // string | auth contract address (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetNftMintersV3(context.Background()).Channel(channel).AuthContractAddress(authContractAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetNftMintersV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNftMintersV3`: GetNftMintersResponse
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetNftMintersV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNftMintersV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channel** | **int32** | channel id, if you dont know your channel id, please contact with Loopring team. | 
 **authContractAddress** | **string** | auth contract address | 

### Return type

[**GetNftMintersResponse**](GetNftMintersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNftOrderUserRateAmountV3

> GetUserNftOrderRatesAmountsResponse GetNftOrderUserRateAmountV3(ctx).AccountId(accountId).NftTokenAddress(nftTokenAddress).FeeTokenSymbol(feeTokenSymbol).Execute()

Query current token minimum amount to place order based on users VIP level and max fee bips



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := int64(10086) // int64 | Account ID, some hash query APIs doesnt need it if in hash query mode, check require flag of each API to see if its a must.
    nftTokenAddress := "0xbbbbca6a901c926f240b89eacb641d8aec7aeafd" // string | NFT token address of order
    feeTokenSymbol := "ETH" // string | Fee token symbol (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetNftOrderUserRateAmountV3(context.Background()).AccountId(accountId).NftTokenAddress(nftTokenAddress).FeeTokenSymbol(feeTokenSymbol).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetNftOrderUserRateAmountV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNftOrderUserRateAmountV3`: GetUserNftOrderRatesAmountsResponse
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetNftOrderUserRateAmountV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNftOrderUserRateAmountV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int64** | Account ID, some hash query APIs doesnt need it if in hash query mode, check require flag of each API to see if its a must. | 
 **nftTokenAddress** | **string** | NFT token address of order | 
 **feeTokenSymbol** | **string** | Fee token symbol | 

### Return type

[**GetUserNftOrderRatesAmountsResponse**](GetUserNftOrderRatesAmountsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNftRequestFeesV3

> OffFeeInfo2 GetNftRequestFeesV3(ctx).AccountId(accountId).RequestType(requestType).TokenAddress(tokenAddress).Amount(amount).Execute()

Query current NFT requests fee amount



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := int32(10001) // int32 | Account ID
    requestType := int32(9) // int32 | Off-chain request type
    tokenAddress := "0xbbbbca6a901c926f240b89eacb641d8aec7aeafd" // string | the NFT tokenAddress (optional)
    amount := "10000000000" // string | The amount to withdraw (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetNftRequestFeesV3(context.Background()).AccountId(accountId).RequestType(requestType).TokenAddress(tokenAddress).Amount(amount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetNftRequestFeesV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNftRequestFeesV3`: OffFeeInfo2
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetNftRequestFeesV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNftRequestFeesV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32** | Account ID | 
 **requestType** | **int32** | Off-chain request type | 
 **tokenAddress** | **string** | the NFT tokenAddress | 
 **amount** | **string** | The amount to withdraw | 

### Return type

[**OffFeeInfo2**](OffFeeInfo2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNftTokenAddressesV3

> GetNftTokenAddressesResponse GetNftTokenAddressesV3(ctx).Channel(channel).AuthContractAddress(authContractAddress).Offset(offset).Limit(limit).Execute()

Get Nft token addresses



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    channel := int32(1) // int32 | channel id, if you dont know your channel id, please contact with Loopring team.
    authContractAddress := "authContractAddress_example" // string | auth contract address (optional)
    offset := int64(0) // int64 | amount to skip(default: 0) (optional) (default to 0)
    limit := int32(100) // int32 | Max amount of token addresses to return(default: 500) (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetNftTokenAddressesV3(context.Background()).Channel(channel).AuthContractAddress(authContractAddress).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetNftTokenAddressesV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNftTokenAddressesV3`: GetNftTokenAddressesResponse
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetNftTokenAddressesV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNftTokenAddressesV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channel** | **int32** | channel id, if you dont know your channel id, please contact with Loopring team. | 
 **authContractAddress** | **string** | auth contract address | 
 **offset** | **int64** | amount to skip(default: 0) | [default to 0]
 **limit** | **int32** | Max amount of token addresses to return(default: 500) | [default to 0]

### Return type

[**GetNftTokenAddressesResponse**](GetNftTokenAddressesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNftTradesV3

> NftTradeList GetNftTradesV3(ctx).AccountId(accountId).NftData(nftData).OrderHash(orderHash).Start(start).End(end).StartId(startId).Limit(limit).Execute()

Users Nft trade list



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := int64(10001) // int64 | users accountId.
    nftData := "0xf7c932351186c3a9053f313eefa16209c018f7f1dba8aa8ca7100400f7c31085" // string | the nftData of the NFT token (optional)
    orderHash := "0xf7c932351186c3a9053f313eefa16209c018f7f1dba8aa8ca7100400f7c31085" // string | request.getNftTrades.orderHash (optional)
    start := int64(0xf7c932351186c3a9053f313eefa16209c018f7f1dba8aa8ca7100400f7c31085) // int64 | request.getNftTrades.tradeHash (optional) (default to 0)
    end := int64(1567053142) // int64 | Start time in milliseconds (optional) (default to 0)
    startId := int64(1567053142) // int64 | End time in milliseconds (optional) (default to 0)
    limit := int32(1234) // int32 | The begin id of the query. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetNftTradesV3(context.Background()).AccountId(accountId).NftData(nftData).OrderHash(orderHash).Start(start).End(end).StartId(startId).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetNftTradesV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNftTradesV3`: NftTradeList
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetNftTradesV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNftTradesV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int64** | users accountId. | 
 **nftData** | **string** | the nftData of the NFT token | 
 **orderHash** | **string** | request.getNftTrades.orderHash | 
 **start** | **int64** | request.getNftTrades.tradeHash | [default to 0]
 **end** | **int64** | Start time in milliseconds | [default to 0]
 **startId** | **int64** | End time in milliseconds | [default to 0]
 **limit** | **int32** | The begin id of the query. | [default to 0]

### Return type

[**NftTradeList**](NftTradeList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNftTransfersV3

> GetUserNftTransfersResponse GetNftTransfersV3(ctx).AccountId(accountId).NftData(nftData).Start(start).End(end).StartId(startId).Limit(limit).TxStatus(txStatus).Hashes(hashes).Execute()

Get user NFT transfer history



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := int64(10001) // int64 | users accountId.
    nftData := "0xf7c932351186c3a9053f313eefa16209c018f7f1dba8aa8ca7100400f7c31085" // string | the nftData of the NFT token (optional)
    start := int64(1567053142) // int64 | Start time in milliseconds (optional) (default to 0)
    end := int64(1567053142) // int64 | End time in milliseconds (optional) (default to 0)
    startId := int64(1234) // int64 | The begin id of the query. (optional) (default to 0)
    limit := int32(50) // int32 | Number of records to return (optional) (default to 0)
    txStatus := "processing" // string | The transfer status (optional)
    hashes := "0xf7c932351186c3a9053f313eefa16209c018f7f1dba8aa8ca7100400f7c31085" // string | The hashes (split by ,) of the transactions, normally its L2 tx hash, except the deposit which uses L1 tx hash. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetNftTransfersV3(context.Background()).AccountId(accountId).NftData(nftData).Start(start).End(end).StartId(startId).Limit(limit).TxStatus(txStatus).Hashes(hashes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetNftTransfersV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNftTransfersV3`: GetUserNftTransfersResponse
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetNftTransfersV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNftTransfersV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int64** | users accountId. | 
 **nftData** | **string** | the nftData of the NFT token | 
 **start** | **int64** | Start time in milliseconds | [default to 0]
 **end** | **int64** | End time in milliseconds | [default to 0]
 **startId** | **int64** | The begin id of the query. | [default to 0]
 **limit** | **int32** | Number of records to return | [default to 0]
 **txStatus** | **string** | The transfer status | 
 **hashes** | **string** | The hashes (split by ,) of the transactions, normally its L2 tx hash, except the deposit which uses L1 tx hash. | 

### Return type

[**GetUserNftTransfersResponse**](GetUserNftTransfersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNftWithdrawalsV3

> GetUserNftWithdrawalResponse GetNftWithdrawalsV3(ctx).AccountId(accountId).NftData(nftData).Start(start).End(end).StartId(startId).Limit(limit).TxStatus(txStatus).Hashes(hashes).Execute()

Get user NFT withdraw history



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := int64(10001) // int64 | users accountId.
    nftData := "0xf7c932351186c3a9053f313eefa16209c018f7f1dba8aa8ca7100400f7c31085" // string | the nftData of the NFT token (optional)
    start := int64(1567053142) // int64 | Start time in milliseconds (optional) (default to 0)
    end := int64(1567053142) // int64 | End time in milliseconds (optional) (default to 0)
    startId := int64(1234) // int64 | The begin id of the query. (optional) (default to 0)
    limit := int32(50) // int32 | Number of records to return (optional) (default to 0)
    txStatus := "processing" // string | The withdrawal status (optional)
    hashes := "0xf7c932351186c3a9053f313eefa16209c018f7f1dba8aa8ca7100400f7c31085" // string | The hashes (split by ,) of the transactions, normally its L2 tx hash, except the deposit which uses L1 tx hash. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetNftWithdrawalsV3(context.Background()).AccountId(accountId).NftData(nftData).Start(start).End(end).StartId(startId).Limit(limit).TxStatus(txStatus).Hashes(hashes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetNftWithdrawalsV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNftWithdrawalsV3`: GetUserNftWithdrawalResponse
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetNftWithdrawalsV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNftWithdrawalsV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int64** | users accountId. | 
 **nftData** | **string** | the nftData of the NFT token | 
 **start** | **int64** | Start time in milliseconds | [default to 0]
 **end** | **int64** | End time in milliseconds | [default to 0]
 **startId** | **int64** | The begin id of the query. | [default to 0]
 **limit** | **int32** | Number of records to return | [default to 0]
 **txStatus** | **string** | The withdrawal status | 
 **hashes** | **string** | The hashes (split by ,) of the transactions, normally its L2 tx hash, except the deposit which uses L1 tx hash. | 

### Return type

[**GetUserNftWithdrawalResponse**](GetUserNftWithdrawalResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNftsInfoV3

> NftTokenInfo GetNftsInfoV3(ctx).NftDatas(nftDatas).Execute()

Query NFT info by looprings nftData



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    nftDatas := "0xf7c932351186c3a9053f313eefa16209c018f7f1dba8aa8ca7100400f7c31085,0xc3a9053f313eef7c932351ca7100400f7c186fa16209c018f7f1dba8aa831085" // string | The Loopring's NFT token data identifier which is a hash string of NFT token address and NFT_ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetNftsInfoV3(context.Background()).NftDatas(nftDatas).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetNftsInfoV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNftsInfoV3`: NftTokenInfo
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetNftsInfoV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNftsInfoV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nftDatas** | **string** | The Loopring&#39;s NFT token data identifier which is a hash string of NFT token address and NFT_ID | 

### Return type

[**NftTokenInfo**](NftTokenInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotWithdrawContractTokensV3

> TokenInfoV3 GetNotWithdrawContractTokensV3(ctx).Execute()

api.getNotWithdrawContractTokens.value



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetNotWithdrawContractTokensV3(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetNotWithdrawContractTokensV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotWithdrawContractTokensV3`: TokenInfoV3
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetNotWithdrawContractTokensV3`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotWithdrawContractTokensV3Request struct via the builder pattern


### Return type

[**TokenInfoV3**](TokenInfoV3.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrderAmountV2

> GetOrderAmountResponse GetOrderAmountV2(ctx).TokenSymbol(tokenSymbol).Execute()

Query current token minimum amount to place order



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tokenSymbol := "tokenSymbol_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetOrderAmountV2(context.Background()).TokenSymbol(tokenSymbol).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetOrderAmountV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrderAmountV2`: GetOrderAmountResponse
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetOrderAmountV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderAmountV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenSymbol** | **string** |  | 

### Return type

[**GetOrderAmountResponse**](GetOrderAmountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrderAmountV3

> GetOrderAmountResponseData GetOrderAmountV3(ctx).TokenSymbol(tokenSymbol).Execute()

Query current token minimum amount to place order



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tokenSymbol := "LRC" // string | Token symbol to place order

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetOrderAmountV3(context.Background()).TokenSymbol(tokenSymbol).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetOrderAmountV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrderAmountV3`: GetOrderAmountResponseData
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetOrderAmountV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderAmountV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenSymbol** | **string** | Token symbol to place order | 

### Return type

[**GetOrderAmountResponseData**](GetOrderAmountResponseData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrderDetailV2

> GetOrderDetailResponseV2 GetOrderDetailV2(ctx).AccountId(accountId).OrderHash(orderHash).Execute()

Get order details



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := int64(1) // int64 | Account ID, could be empty if hash query is presented.
    orderHash := "13375450901292179417154974849571793069911517354720397125027633242680470075859" // string | Order hash

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetOrderDetailV2(context.Background()).AccountId(accountId).OrderHash(orderHash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetOrderDetailV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrderDetailV2`: GetOrderDetailResponseV2
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetOrderDetailV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderDetailV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int64** | Account ID, could be empty if hash query is presented. | 
 **orderHash** | **string** | Order hash | 

### Return type

[**GetOrderDetailResponseV2**](GetOrderDetailResponseV2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrderDetailV3

> OrderDetailV3 GetOrderDetailV3(ctx).AccountId(accountId).OrderHash(orderHash).Execute()

Get order details



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := int64(1) // int64 | Account ID, could be empty if hash query is presented.
    orderHash := "13375450901292179417154974849571793069911517354720397125027633242680470075859" // string | Order hash

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetOrderDetailV3(context.Background()).AccountId(accountId).OrderHash(orderHash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetOrderDetailV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrderDetailV3`: OrderDetailV3
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetOrderDetailV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderDetailV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int64** | Account ID, could be empty if hash query is presented. | 
 **orderHash** | **string** | Order hash | 

### Return type

[**OrderDetailV3**](OrderDetailV3.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrderUserRateAmountV2

> GetOrderGroupAmountResponse GetOrderUserRateAmountV2(ctx).AccountId(accountId).Market(market).Execute()

Query current token minimum amount to place order



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := int64(789) // int64 |  (optional)
    market := "market_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetOrderUserRateAmountV2(context.Background()).AccountId(accountId).Market(market).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetOrderUserRateAmountV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrderUserRateAmountV2`: GetOrderGroupAmountResponse
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetOrderUserRateAmountV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderUserRateAmountV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int64** |  | 
 **market** | **string** |  | 

### Return type

[**GetOrderGroupAmountResponse**](GetOrderGroupAmountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrderUserRateAmountV3

> GetOrderGroupAmountData GetOrderUserRateAmountV3(ctx).AccountId(accountId).Market(market).Execute()

Query current token minimum amount to place order based on users VIP level and max fee bips



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := int64(10086) // int64 | Account ID, some hash query APIs doesnt need it if in hash query mode, check require flag of each API to see if its a must.
    market := "LRC-ETH" // string | Trading pair

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetOrderUserRateAmountV3(context.Background()).AccountId(accountId).Market(market).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetOrderUserRateAmountV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrderUserRateAmountV3`: GetOrderGroupAmountData
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetOrderUserRateAmountV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderUserRateAmountV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int64** | Account ID, some hash query APIs doesnt need it if in hash query mode, check require flag of each API to see if its a must. | 
 **market** | **string** | Trading pair | 

### Return type

[**GetOrderGroupAmountData**](GetOrderGroupAmountData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrdersV2

> GetOrdersResponseV2 GetOrdersV2(ctx).AccountId(accountId).Market(market).Start(start).End(end).Side(side).Status(status).Limit(limit).Offset(offset).OrderTypes(orderTypes).TradeChannels(tradeChannels).Execute()

Get multiple orders



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := int32(1) // int32 | Account ID
    market := "LRC-ETH" // string | Trading pair (optional)
    start := int64(1567053142000) // int64 | Lower bound of order's creation timestamp in millisecond (optional) (default to 0)
    end := int64(1567053242000) // int64 | Upper bound of order's creation timestamp in millisecond (optional) (default to 0)
    side := "BUY" // string | \"BUY\" or \"SELL\" (optional)
    status := "processing,processed" // string | Order status. You can specify one of the following values: (optional)
    limit := int32(50) // int32 | Limit of orders (default 50) (optional) (default to 0)
    offset := int64(0) // int64 | Offset of orders (default 0) (optional) (default to 0)
    orderTypes := "LIMIT_ORDER" // string | request.getOrders.orderTypes (optional)
    tradeChannels := "ORDER_BOOK" // string | field.SubmitOrderRequest.tradeChannel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetOrdersV2(context.Background()).AccountId(accountId).Market(market).Start(start).End(end).Side(side).Status(status).Limit(limit).Offset(offset).OrderTypes(orderTypes).TradeChannels(tradeChannels).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetOrdersV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrdersV2`: GetOrdersResponseV2
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetOrdersV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrdersV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32** | Account ID | 
 **market** | **string** | Trading pair | 
 **start** | **int64** | Lower bound of order&#39;s creation timestamp in millisecond | [default to 0]
 **end** | **int64** | Upper bound of order&#39;s creation timestamp in millisecond | [default to 0]
 **side** | **string** | \&quot;BUY\&quot; or \&quot;SELL\&quot; | 
 **status** | **string** | Order status. You can specify one of the following values: | 
 **limit** | **int32** | Limit of orders (default 50) | [default to 0]
 **offset** | **int64** | Offset of orders (default 0) | [default to 0]
 **orderTypes** | **string** | request.getOrders.orderTypes | 
 **tradeChannels** | **string** | field.SubmitOrderRequest.tradeChannel | 

### Return type

[**GetOrdersResponseV2**](GetOrdersResponseV2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrdersV3

> OrdersDetailV3 GetOrdersV3(ctx).AccountId(accountId).Market(market).Start(start).End(end).Side(side).Status(status).Limit(limit).Offset(offset).OrderTypes(orderTypes).TradeChannels(tradeChannels).Execute()

Get multiple orders



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := int32(1) // int32 | Account ID
    market := "LRC-ETH" // string | Trading pair (optional)
    start := int64(1567053142000) // int64 | Lower bound of order's creation timestamp in millisecond (optional) (default to 0)
    end := int64(1567053242000) // int64 | Upper bound of order's creation timestamp in millisecond (optional) (default to 0)
    side := "BUY" // string | \"BUY\" or \"SELL\" (optional)
    status := "processing,processed" // string | Order status. You can specify one of the following values: (optional)
    limit := int32(50) // int32 | Limit of orders (default 50) (optional) (default to 0)
    offset := int64(0) // int64 | Offset of orders (default 0) (optional) (default to 0)
    orderTypes := "LIMIT_ORDER" // string | request.getOrders.orderTypes (optional)
    tradeChannels := "ORDER_BOOK,AMM_POOL" // string | field.SubmitOrderRequest.tradeChannel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetOrdersV3(context.Background()).AccountId(accountId).Market(market).Start(start).End(end).Side(side).Status(status).Limit(limit).Offset(offset).OrderTypes(orderTypes).TradeChannels(tradeChannels).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetOrdersV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrdersV3`: OrdersDetailV3
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetOrdersV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrdersV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32** | Account ID | 
 **market** | **string** | Trading pair | 
 **start** | **int64** | Lower bound of order&#39;s creation timestamp in millisecond | [default to 0]
 **end** | **int64** | Upper bound of order&#39;s creation timestamp in millisecond | [default to 0]
 **side** | **string** | \&quot;BUY\&quot; or \&quot;SELL\&quot; | 
 **status** | **string** | Order status. You can specify one of the following values: | 
 **limit** | **int32** | Limit of orders (default 50) | [default to 0]
 **offset** | **int64** | Offset of orders (default 0) | [default to 0]
 **orderTypes** | **string** | request.getOrders.orderTypes | 
 **tradeChannels** | **string** | field.SubmitOrderRequest.tradeChannel | 

### Return type

[**OrdersDetailV3**](OrdersDetailV3.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPendingRequestsV3

> TransactionBlock GetPendingRequestsV3(ctx).Execute()

Get pending txs



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetPendingRequestsV3(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetPendingRequestsV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPendingRequestsV3`: TransactionBlock
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetPendingRequestsV3`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPendingRequestsV3Request struct via the builder pattern


### Return type

[**TransactionBlock**](TransactionBlock.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPoolsStatsAmmV3

> GetAmmPoolsStatsResponse GetPoolsStatsAmmV3(ctx).AmmPoolMarkets(ammPoolMarkets).Execute()

api.getPoolsStats.value



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ammPoolMarkets := "ammPoolMarkets_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetPoolsStatsAmmV3(context.Background()).AmmPoolMarkets(ammPoolMarkets).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetPoolsStatsAmmV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPoolsStatsAmmV3`: GetAmmPoolsStatsResponse
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetPoolsStatsAmmV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPoolsStatsAmmV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ammPoolMarkets** | **string** |  | 

### Return type

[**GetAmmPoolsStatsResponse**](GetAmmPoolsStatsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPoolsStatsV2

> GetAmmPoolsStatsResponse GetPoolsStatsV2(ctx).AmmPoolMarkets(ammPoolMarkets).Execute()

api.getPoolsStats.value



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ammPoolMarkets := "ammPoolMarkets_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetPoolsStatsV2(context.Background()).AmmPoolMarkets(ammPoolMarkets).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetPoolsStatsV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPoolsStatsV2`: GetAmmPoolsStatsResponse
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetPoolsStatsV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPoolsStatsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ammPoolMarkets** | **string** |  | 

### Return type

[**GetAmmPoolsStatsResponse**](GetAmmPoolsStatsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPoolsStatsV3

> AmmPoolStatistics GetPoolsStatsV3(ctx).AmmPoolMarkets(ammPoolMarkets).Execute()

api.getPoolsStats.value



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ammPoolMarkets := "ammPoolMarkets_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetPoolsStatsV3(context.Background()).AmmPoolMarkets(ammPoolMarkets).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetPoolsStatsV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPoolsStatsV3`: AmmPoolStatistics
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetPoolsStatsV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPoolsStatsV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ammPoolMarkets** | **string** |  | 

### Return type

[**AmmPoolStatistics**](AmmPoolStatistics.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPriceV2

> GetPriceV2Response GetPriceV2(ctx).Legal(legal).Execute()

Get token fiat prices



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    legal := "USD" // string | request.getPrice.legal

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetPriceV2(context.Background()).Legal(legal).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetPriceV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPriceV2`: GetPriceV2Response
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetPriceV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPriceV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **legal** | **string** | request.getPrice.legal | 

### Return type

[**GetPriceV2Response**](GetPriceV2Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPriceV3

> GetPriceResponseV3 GetPriceV3(ctx).Legal(legal).Execute()

Get token fiat prices



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    legal := "USD" // string | The fiat currency to uses. Currently the following values are supported: USD,CNY,JPY,EUR,GBP,HKD

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetPriceV3(context.Background()).Legal(legal).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetPriceV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPriceV3`: GetPriceResponseV3
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetPriceV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPriceV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **legal** | **string** | The fiat currency to uses. Currently the following values are supported: USD,CNY,JPY,EUR,GBP,HKD | 

### Return type

[**GetPriceResponseV3**](GetPriceResponseV3.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProfitSharingStatusV3

> map[string]interface{} GetProfitSharingStatusV3(ctx).AccountId(accountId).TokenId(tokenId).Start(start).Size(size).RewardType(rewardType).Taker(taker).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := int64(789) // int64 |  (optional)
    tokenId := int32(56) // int32 |  (optional)
    start := int64(789) // int64 |  (optional)
    size := int32(56) // int32 |  (optional)
    rewardType := int32(56) // int32 |  (optional)
    taker := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetProfitSharingStatusV3(context.Background()).AccountId(accountId).TokenId(tokenId).Start(start).Size(size).RewardType(rewardType).Taker(taker).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetProfitSharingStatusV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProfitSharingStatusV3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetProfitSharingStatusV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProfitSharingStatusV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int64** |  | 
 **tokenId** | **int32** |  | 
 **start** | **int64** |  | 
 **size** | **int32** |  | 
 **rewardType** | **int32** |  | 
 **taker** | **int32** |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProtocolPortraitV3

> map[string]interface{} GetProtocolPortraitV3(ctx).Timestamp(timestamp).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    timestamp := int64(789) // int64 |  (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetProtocolPortraitV3(context.Background()).Timestamp(timestamp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetProtocolPortraitV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProtocolPortraitV3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetProtocolPortraitV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProtocolPortraitV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | [default to 0]

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRecommendedGasPriceRangeV3

> GetRecommendedGasPriceResponseRangeV2 GetRecommendedGasPriceRangeV3(ctx).Execute()

Get recommended gas price



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetRecommendedGasPriceRangeV3(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetRecommendedGasPriceRangeV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRecommendedGasPriceRangeV3`: GetRecommendedGasPriceResponseRangeV2
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetRecommendedGasPriceRangeV3`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRecommendedGasPriceRangeV3Request struct via the builder pattern


### Return type

[**GetRecommendedGasPriceResponseRangeV2**](GetRecommendedGasPriceResponseRangeV2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRecommendedGasPriceV3

> GetRecommendedGasPriceResponseV2 GetRecommendedGasPriceV3(ctx).Execute()

Get recommended gas price



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetRecommendedGasPriceV3(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetRecommendedGasPriceV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRecommendedGasPriceV3`: GetRecommendedGasPriceResponseV2
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetRecommendedGasPriceV3`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRecommendedGasPriceV3Request struct via the builder pattern


### Return type

[**GetRecommendedGasPriceResponseV2**](GetRecommendedGasPriceResponseV2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRecommendedMarketsV3

> GetExchangeFeeInfoResponseData GetRecommendedMarketsV3(ctx).Size(size).Execute()

api.getRecommendedMarkets.value



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    size := int32(56) // int32 |  (optional) (default to 4)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetRecommendedMarketsV3(context.Background()).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetRecommendedMarketsV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRecommendedMarketsV3`: GetExchangeFeeInfoResponseData
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetRecommendedMarketsV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRecommendedMarketsV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32** |  | [default to 4]

### Return type

[**GetExchangeFeeInfoResponseData**](GetExchangeFeeInfoResponseData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRequestInBlockV3

> RequestInBlockResponse GetRequestInBlockV3(ctx).Hashes(hashes).Execute()

Get a txs block



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    hashes := "0xf7c932351186c3a9053f313eefa16209c018f7f1dba8aa8ca7100400f7c31085" // string | L2 tx hash (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetRequestInBlockV3(context.Background()).Hashes(hashes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetRequestInBlockV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRequestInBlockV3`: RequestInBlockResponse
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetRequestInBlockV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRequestInBlockV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hashes** | **string** | L2 tx hash | 

### Return type

[**RequestInBlockResponse**](RequestInBlockResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTickerV2

> GetTickerResponseV2 GetTickerV2(ctx).Market(market).Execute()

Get market ticker



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    market := "LRC-ETH,LRC-USDT" // string | Market pair, support multiple markets

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetTickerV2(context.Background()).Market(market).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetTickerV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTickerV2`: GetTickerResponseV2
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetTickerV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTickerV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **market** | **string** | Market pair, support multiple markets | 

### Return type

[**GetTickerResponseV2**](GetTickerResponseV2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTickerV3

> GetTickerResponseV3 GetTickerV3(ctx).Market(market).Execute()

Get market ticker



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    market := "LRC-ETH,LRC-USDT" // string | Market pair, support multiple markets

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetTickerV3(context.Background()).Market(market).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetTickerV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTickerV3`: GetTickerResponseV3
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetTickerV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTickerV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **market** | **string** | Market pair, support multiple markets | 

### Return type

[**GetTickerResponseV3**](GetTickerResponseV3.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTimestampV2

> GetTimestampV2Response GetTimestampV2(ctx).Execute()

Get relayer's current time



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetTimestampV2(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetTimestampV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTimestampV2`: GetTimestampV2Response
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetTimestampV2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTimestampV2Request struct via the builder pattern


### Return type

[**GetTimestampV2Response**](GetTimestampV2Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTimestampV3

> TimestampV3 GetTimestampV3(ctx).Execute()

Get relayer's current time



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetTimestampV3(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetTimestampV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTimestampV3`: TimestampV3
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetTimestampV3`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTimestampV3Request struct via the builder pattern


### Return type

[**TimestampV3**](TimestampV3.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTokenBalanceV3

> GetTokenBalancesV2Response GetTokenBalanceV3(ctx).Owner(owner).Token(token).Execute()

Get user token balance



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "0x0000000000000000000000000000000000000000" // string | Ethereum address
    token := "0x0000000000000000000000000000000000000000" // string | Token's ERC20 address

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetTokenBalanceV3(context.Background()).Owner(owner).Token(token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetTokenBalanceV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTokenBalanceV3`: GetTokenBalancesV2Response
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetTokenBalanceV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenBalanceV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | Ethereum address | 
 **token** | **string** | Token&#39;s ERC20 address | 

### Return type

[**GetTokenBalancesV2Response**](GetTokenBalancesV2Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTokenDecimalV2

> map[string]interface{} GetTokenDecimalV2(ctx).TokenAddresses(tokenAddresses).Network(network).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tokenAddresses := "0x0000000000000000000000000000000000000000,0x0000000000000000000000000000000000000001" // string | token address to be queried
    network := "ETHEREUM or ARBITRUM" // string | network

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetTokenDecimalV2(context.Background()).TokenAddresses(tokenAddresses).Network(network).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetTokenDecimalV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTokenDecimalV2`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetTokenDecimalV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenDecimalV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenAddresses** | **string** | token address to be queried | 
 **network** | **string** | network | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTokenInfoV3

> GetUserBalancesResponseV2 GetTokenInfoV3(ctx).TokenAddress(tokenAddress).Execute()

api.getTokenInfo.value



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tokenAddress := "ETH" // string | token address to be queried (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetTokenInfoV3(context.Background()).TokenAddress(tokenAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetTokenInfoV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTokenInfoV3`: GetUserBalancesResponseV2
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetTokenInfoV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenInfoV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenAddress** | **string** | token address to be queried | 

### Return type

[**GetUserBalancesResponseV2**](GetUserBalancesResponseV2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTokenPricesV3

> map[string]interface{} GetTokenPricesV3(ctx).Token(token).Limit(limit).Interval(interval).Currency(currency).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    token := "token_example" // string |  (optional)
    limit := int32(56) // int32 |  (optional) (default to 30)
    interval := "interval_example" // string |  (optional) (default to "DAY")
    currency := "currency_example" // string |  (optional) (default to "USD")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetTokenPricesV3(context.Background()).Token(token).Limit(limit).Interval(interval).Currency(currency).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetTokenPricesV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTokenPricesV3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetTokenPricesV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenPricesV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** |  | 
 **limit** | **int32** |  | [default to 30]
 **interval** | **string** |  | [default to &quot;DAY&quot;]
 **currency** | **string** |  | [default to &quot;USD&quot;]

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTokensV2

> GetTokenInfoResponse GetTokensV2(ctx).Execute()

Get token configurations



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetTokensV2(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetTokensV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTokensV2`: GetTokenInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetTokensV2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokensV2Request struct via the builder pattern


### Return type

[**GetTokenInfoResponse**](GetTokenInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTokensV3

> TokenInfoV3 GetTokensV3(ctx).Execute()

Get token configurations



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetTokensV3(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetTokensV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTokensV3`: TokenInfoV3
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetTokensV3`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokensV3Request struct via the builder pattern


### Return type

[**TokenInfoV3**](TokenInfoV3.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTournamentRankV3

> map[string]interface{} GetTournamentRankV3(ctx).AmmPoolMarket(ammPoolMarket).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ammPoolMarket := "ammPoolMarket_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetTournamentRankV3(context.Background()).AmmPoolMarket(ammPoolMarket).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetTournamentRankV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTournamentRankV3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetTournamentRankV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTournamentRankV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ammPoolMarket** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserAssetsV3

> map[string]interface{} GetUserAssetsV3(ctx).Address(address).Currency(currency).AssetTypes(assetTypes).Token(token).Limit(limit).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    address := "address_example" // string |  (optional)
    currency := "currency_example" // string |  (optional) (default to "USD")
    assetTypes := "assetTypes_example" // string |  (optional)
    token := "token_example" // string |  (optional)
    limit := int32(56) // int32 |  (optional) (default to 30)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetUserAssetsV3(context.Background()).Address(address).Currency(currency).AssetTypes(assetTypes).Token(token).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetUserAssetsV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserAssetsV3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetUserAssetsV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserAssetsV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **address** | **string** |  | 
 **currency** | **string** |  | [default to &quot;USD&quot;]
 **assetTypes** | **string** |  | 
 **token** | **string** |  | 
 **limit** | **int32** |  | [default to 30]

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserBalancesV2

> GetUserBalancesResponseV2 GetUserBalancesV2(ctx).AccountId(accountId).Tokens(tokens).Execute()

Get user exchange balances



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := int32(1) // int32 | AccountID
    tokens := "0,1" // string | Query tokens (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetUserBalancesV2(context.Background()).AccountId(accountId).Tokens(tokens).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetUserBalancesV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserBalancesV2`: GetUserBalancesResponseV2
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetUserBalancesV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserBalancesV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32** | AccountID | 
 **tokens** | **string** | Query tokens | 

### Return type

[**GetUserBalancesResponseV2**](GetUserBalancesResponseV2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserBalancesV3

> BalanceV3 GetUserBalancesV3(ctx).AccountId(accountId).Tokens(tokens).Execute()

Get user exchange balances



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := int32(1) // int32 | AccountID
    tokens := "0,1" // string | Query tokens (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetUserBalancesV3(context.Background()).AccountId(accountId).Tokens(tokens).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetUserBalancesV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserBalancesV3`: BalanceV3
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetUserBalancesV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserBalancesV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32** | AccountID | 
 **tokens** | **string** | Query tokens | 

### Return type

[**BalanceV3**](BalanceV3.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserBillV3

> map[string]interface{} GetUserBillV3(ctx).AccountId(accountId).BillType(billType).Start(start).End(end).Limit(limit).Offset(offset).TokenId(tokenId).Income(income).TransferAddress(transferAddress).FromAddress(fromAddress).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := int32(56) // int32 |  (optional)
    billType := "billType_example" // string |  (optional)
    start := int64(789) // int64 |  (optional) (default to 0)
    end := int64(789) // int64 |  (optional) (default to 0)
    limit := int32(56) // int32 |  (optional) (default to 20)
    offset := int32(56) // int32 |  (optional) (default to 0)
    tokenId := int32(56) // int32 |  (optional) (default to -1)
    income := true // bool |  (optional) (default to false)
    transferAddress := "transferAddress_example" // string |  (optional)
    fromAddress := "fromAddress_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetUserBillV3(context.Background()).AccountId(accountId).BillType(billType).Start(start).End(end).Limit(limit).Offset(offset).TokenId(tokenId).Income(income).TransferAddress(transferAddress).FromAddress(fromAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetUserBillV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserBillV3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetUserBillV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserBillV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32** |  | 
 **billType** | **string** |  | 
 **start** | **int64** |  | [default to 0]
 **end** | **int64** |  | [default to 0]
 **limit** | **int32** |  | [default to 20]
 **offset** | **int32** |  | [default to 0]
 **tokenId** | **int32** |  | [default to -1]
 **income** | **bool** |  | [default to false]
 **transferAddress** | **string** |  | 
 **fromAddress** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserCreateV2

> GetUserCreateResponseV2 GetUserCreateV2(ctx).AccountId(accountId).Start(start).End(end).Status(status).Limit(limit).Offset(offset).Execute()

Get user registration transactions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := int32(1) // int32 | Account ID, some hash query APIs doesnt need it if in hash query mode, check require flag of each API to see if its a must.
    start := int64(1578558098000) // int64 | Start time in milliseconds (optional) (default to 0)
    end := int64(1578558098000) // int64 | End time in milliseconds (optional) (default to 0)
    status := "processing,processed" // string | Comma separated status values (optional)
    limit := int32(50) // int32 | Number of records to return (optional) (default to 0)
    offset := int64(1) // int64 | Number of records to skip (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetUserCreateV2(context.Background()).AccountId(accountId).Start(start).End(end).Status(status).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetUserCreateV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserCreateV2`: GetUserCreateResponseV2
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetUserCreateV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserCreateV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32** | Account ID, some hash query APIs doesnt need it if in hash query mode, check require flag of each API to see if its a must. | 
 **start** | **int64** | Start time in milliseconds | [default to 0]
 **end** | **int64** | End time in milliseconds | [default to 0]
 **status** | **string** | Comma separated status values | 
 **limit** | **int32** | Number of records to return | [default to 0]
 **offset** | **int64** | Number of records to skip | [default to 0]

### Return type

[**GetUserCreateResponseV2**](GetUserCreateResponseV2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserCreateV3

> UserCreateDataList GetUserCreateV3(ctx).AccountId(accountId).Start(start).End(end).Status(status).Limit(limit).Offset(offset).Execute()

Get user registration transactions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := int32(1) // int32 | Account ID, some hash query APIs doesnt need it if in hash query mode, check require flag of each API to see if its a must.
    start := int64(1578558098000) // int64 | Start time in milliseconds (optional) (default to 0)
    end := int64(1578558098000) // int64 | End time in milliseconds (optional) (default to 0)
    status := "processing,processed" // string | Comma separated status values (optional)
    limit := int32(50) // int32 | Number of records to return (optional) (default to 0)
    offset := int64(1) // int64 | Number of records to skip (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetUserCreateV3(context.Background()).AccountId(accountId).Start(start).End(end).Status(status).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetUserCreateV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserCreateV3`: UserCreateDataList
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetUserCreateV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserCreateV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32** | Account ID, some hash query APIs doesnt need it if in hash query mode, check require flag of each API to see if its a must. | 
 **start** | **int64** | Start time in milliseconds | [default to 0]
 **end** | **int64** | End time in milliseconds | [default to 0]
 **status** | **string** | Comma separated status values | 
 **limit** | **int32** | Number of records to return | [default to 0]
 **offset** | **int64** | Number of records to skip | [default to 0]

### Return type

[**UserCreateDataList**](UserCreateDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserDepositsV2

> GetUserDepositResponseV2 GetUserDepositsV2(ctx).AccountId(accountId).Start(start).End(end).Status(status).Limit(limit).TokenSymbol(tokenSymbol).Offset(offset).Execute()

Get user deposit history



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := int32(1) // int32 | Account ID, some hash query APIs doesnt need it if in hash query mode, check require flag of each API to see if its a must.
    start := int64(1578558098000) // int64 | Start time in milliseconds (optional) (default to 0)
    end := int64(1578558098000) // int64 | End time in milliseconds (optional) (default to 0)
    status := "processing,processed" // string | Comma separated status values (optional)
    limit := int32(50) // int32 | Number of records to return (optional) (default to 0)
    tokenSymbol := "ETH" // string | Token to filter. If you want to return deposit records for all tokens, omit this parameter (optional)
    offset := int64(1) // int64 | Number of records to skip (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetUserDepositsV2(context.Background()).AccountId(accountId).Start(start).End(end).Status(status).Limit(limit).TokenSymbol(tokenSymbol).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetUserDepositsV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserDepositsV2`: GetUserDepositResponseV2
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetUserDepositsV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserDepositsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32** | Account ID, some hash query APIs doesnt need it if in hash query mode, check require flag of each API to see if its a must. | 
 **start** | **int64** | Start time in milliseconds | [default to 0]
 **end** | **int64** | End time in milliseconds | [default to 0]
 **status** | **string** | Comma separated status values | 
 **limit** | **int32** | Number of records to return | [default to 0]
 **tokenSymbol** | **string** | Token to filter. If you want to return deposit records for all tokens, omit this parameter | 
 **offset** | **int64** | Number of records to skip | [default to 0]

### Return type

[**GetUserDepositResponseV2**](GetUserDepositResponseV2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserDepositsV3

> DepositDataList GetUserDepositsV3(ctx).AccountId(accountId).Start(start).End(end).Status(status).Limit(limit).TokenSymbol(tokenSymbol).Offset(offset).Hashes(hashes).Execute()

Get user deposit history



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := int64(1) // int64 | Account ID, some hash query APIs doesnt need it if in hash query mode, check require flag of each API to see if its a must. (optional) (default to 0)
    start := int64(1578558098000) // int64 | Start time in milliseconds (optional) (default to 0)
    end := int64(1578558098000) // int64 | End time in milliseconds (optional) (default to 0)
    status := "processing,processed" // string | Comma separated status values (optional)
    limit := int32(50) // int32 | Number of records to return (optional) (default to 0)
    tokenSymbol := "ETH" // string | Token to filter. If you want to return deposit records for all tokens, omit this parameter (optional)
    offset := int64(1) // int64 | Number of records to skip (optional) (default to 0)
    hashes := "0xf7c932351186c3a9053f313eefa16209c018f7f1dba8aa8ca7100400f7c31085" // string | The hashes (split by ,) of the transactions, normally its L2 tx hash, except the deposit which uses L1 tx hash. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetUserDepositsV3(context.Background()).AccountId(accountId).Start(start).End(end).Status(status).Limit(limit).TokenSymbol(tokenSymbol).Offset(offset).Hashes(hashes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetUserDepositsV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserDepositsV3`: DepositDataList
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetUserDepositsV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserDepositsV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int64** | Account ID, some hash query APIs doesnt need it if in hash query mode, check require flag of each API to see if its a must. | [default to 0]
 **start** | **int64** | Start time in milliseconds | [default to 0]
 **end** | **int64** | End time in milliseconds | [default to 0]
 **status** | **string** | Comma separated status values | 
 **limit** | **int32** | Number of records to return | [default to 0]
 **tokenSymbol** | **string** | Token to filter. If you want to return deposit records for all tokens, omit this parameter | 
 **offset** | **int64** | Number of records to skip | [default to 0]
 **hashes** | **string** | The hashes (split by ,) of the transactions, normally its L2 tx hash, except the deposit which uses L1 tx hash. | 

### Return type

[**DepositDataList**](DepositDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserFeeRates2V2

> GetUserFeeRates2Request GetUserFeeRates2V2(ctx).AccountId(accountId).Market(market).TokenB(tokenB).AmountB(amountB).Execute()

Query user place order fee rate



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := int32(1) // int32 | Account ID
    market := "LRC-ETH" // string | List of markets to be queried separated by \",\"
    tokenB := int32(0) // int32 | Token ID
    amountB := "10000000000000000" // string | Amount to buy

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetUserFeeRates2V2(context.Background()).AccountId(accountId).Market(market).TokenB(tokenB).AmountB(amountB).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetUserFeeRates2V2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserFeeRates2V2`: GetUserFeeRates2Request
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetUserFeeRates2V2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserFeeRates2V2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32** | Account ID | 
 **market** | **string** | List of markets to be queried separated by \&quot;,\&quot; | 
 **tokenB** | **int32** | Token ID | 
 **amountB** | **string** | Amount to buy | 

### Return type

[**GetUserFeeRates2Request**](GetUserFeeRates2Request.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserFeeRates2V3

> GetUserFeeRatesResponseData GetUserFeeRates2V3(ctx).AccountId(accountId).Market(market).TokenB(tokenB).AmountB(amountB).Execute()

Query user place order fee rate



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := int32(1) // int32 | Account ID
    market := "LRC-ETH" // string | List of markets to be queried separated by \",\"
    tokenB := int32(0) // int32 | Token ID
    amountB := "10000000000000000" // string | Amount to buy

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetUserFeeRates2V3(context.Background()).AccountId(accountId).Market(market).TokenB(tokenB).AmountB(amountB).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetUserFeeRates2V3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserFeeRates2V3`: GetUserFeeRatesResponseData
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetUserFeeRates2V3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserFeeRates2V3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32** | Account ID | 
 **market** | **string** | List of markets to be queried separated by \&quot;,\&quot; | 
 **tokenB** | **int32** | Token ID | 
 **amountB** | **string** | Amount to buy | 

### Return type

[**GetUserFeeRatesResponseData**](GetUserFeeRatesResponseData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserNftBalancesV3

> GetUserNftBalancesResponse GetUserNftBalancesV3(ctx).AccountId(accountId).NftDatas(nftDatas).TokenAddrs(tokenAddrs).TokenIds(tokenIds).Offset(offset).Limit(limit).NonZero(nonZero).Execute()

Get users NFT balance



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := int64(10001) // int64 | AccountID
    nftDatas := "0xf7c932351186c3a9053f313eefa16209c018f7f1dba8aa8ca7100400f7c31085,0xc3a9053f313eef7c932351ca7100400f7c186fa16209c018f7f1dba8aa831085" // string | The Loopring's NFT token data identifier which is a hash string of NFT token address and NFT_ID (optional)
    tokenAddrs := "0xf7c932351186c3a9053f313eefa16209c018f7f1dba8aa8ca7100400f7c31085,0xc3a9053f313eef7c932351ca7100400f7c186fa16209c018f7f1dba8aa831085" // string | NFT token address (optional)
    tokenIds := "32768,32769,32770" // string | The token slot ID in loopring DEX. (optional)
    offset := int64(12345) // int64 | Number of records to skip (optional) (default to 0)
    limit := int32(20) // int32 | Number of records to return (optional) (default to 0)
    nonZero := true // bool | Hide 0 balance NFT token, default is true (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetUserNftBalancesV3(context.Background()).AccountId(accountId).NftDatas(nftDatas).TokenAddrs(tokenAddrs).TokenIds(tokenIds).Offset(offset).Limit(limit).NonZero(nonZero).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetUserNftBalancesV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserNftBalancesV3`: GetUserNftBalancesResponse
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetUserNftBalancesV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserNftBalancesV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int64** | AccountID | 
 **nftDatas** | **string** | The Loopring&#39;s NFT token data identifier which is a hash string of NFT token address and NFT_ID | 
 **tokenAddrs** | **string** | NFT token address | 
 **tokenIds** | **string** | The token slot ID in loopring DEX. | 
 **offset** | **int64** | Number of records to skip | [default to 0]
 **limit** | **int32** | Number of records to return | [default to 0]
 **nonZero** | **bool** | Hide 0 balance NFT token, default is true | [default to false]

### Return type

[**GetUserNftBalancesResponse**](GetUserNftBalancesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserNftMintsV3

> GetUserNftMintsResponse GetUserNftMintsV3(ctx).AccountId(accountId).NftData(nftData).Start(start).End(end).StartId(startId).Limit(limit).TxStatus(txStatus).Hashes(hashes).Execute()

Get user NFT mint history



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := int64(10001) // int64 | users accountId.
    nftData := "0xf7c932351186c3a9053f313eefa16209c018f7f1dba8aa8ca7100400f7c31085" // string | the nftData of the NFT token (optional)
    start := int64(1567053142) // int64 | Start time in milliseconds (optional) (default to 0)
    end := int64(1567053142) // int64 | End time in milliseconds (optional) (default to 0)
    startId := int64(1234) // int64 | The begin id of the query. (optional) (default to 0)
    limit := int32(50) // int32 | Number of records to return (optional) (default to 0)
    txStatus := "processing" // string | The mint status (optional)
    hashes := "0xf7c932351186c3a9053f313eefa16209c018f7f1dba8aa8ca7100400f7c31085" // string | The hashes (split by ,) of the transactions, normally its L2 tx hash, except the deposit which uses L1 tx hash. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetUserNftMintsV3(context.Background()).AccountId(accountId).NftData(nftData).Start(start).End(end).StartId(startId).Limit(limit).TxStatus(txStatus).Hashes(hashes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetUserNftMintsV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserNftMintsV3`: GetUserNftMintsResponse
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetUserNftMintsV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserNftMintsV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int64** | users accountId. | 
 **nftData** | **string** | the nftData of the NFT token | 
 **start** | **int64** | Start time in milliseconds | [default to 0]
 **end** | **int64** | End time in milliseconds | [default to 0]
 **startId** | **int64** | The begin id of the query. | [default to 0]
 **limit** | **int32** | Number of records to return | [default to 0]
 **txStatus** | **string** | The mint status | 
 **hashes** | **string** | The hashes (split by ,) of the transactions, normally its L2 tx hash, except the deposit which uses L1 tx hash. | 

### Return type

[**GetUserNftMintsResponse**](GetUserNftMintsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserNftOrderFeeRatesV3

> GetUserNftFeeRatesResponse GetUserNftOrderFeeRatesV3(ctx).AccountId(accountId).NftTokenAddress(nftTokenAddress).QuoteToken(quoteToken).QuoteAmount(quoteAmount).Execute()

Query user place order fee rate



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := int32(1) // int32 | Account ID
    nftTokenAddress := "0xf7c932351186c3a9053f313eefa16209c018f7f1dba8aa8ca7100400f7c31085" // string | NFT token address of order
    quoteToken := int32(0) // int32 | the quote token of the NFT-Quote market
    quoteAmount := "10000000000000000" // string | request.getUserFeeRates.quoteAmount

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetUserNftOrderFeeRatesV3(context.Background()).AccountId(accountId).NftTokenAddress(nftTokenAddress).QuoteToken(quoteToken).QuoteAmount(quoteAmount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetUserNftOrderFeeRatesV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserNftOrderFeeRatesV3`: GetUserNftFeeRatesResponse
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetUserNftOrderFeeRatesV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserNftOrderFeeRatesV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32** | Account ID | 
 **nftTokenAddress** | **string** | NFT token address of order | 
 **quoteToken** | **int32** | the quote token of the NFT-Quote market | 
 **quoteAmount** | **string** | request.getUserFeeRates.quoteAmount | 

### Return type

[**GetUserNftFeeRatesResponse**](GetUserNftFeeRatesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserNftTransactionsV3

> GetUserNftTxsResponseV2 GetUserNftTransactionsV3(ctx).AccountId(accountId).Start(start).End(end).Limit(limit).Offset(offset).NftData(nftData).TxStatus(txStatus).Types(types).Hashes(hashes).Execute()

api.getUserNftTransactions.value



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := int64(10001) // int64 | request.getUserNftTransactions.accountId
    start := int64(1) // int64 | request.getUserNftTransactions.start (optional) (default to 0)
    end := int64(100) // int64 | request.getUserNftTransactions.end (optional) (default to 0)
    limit := int32(50) // int32 | request.getUserNftTransactions.limit (optional) (default to 0)
    offset := int64(1) // int64 | request.getUserNftTransactions.offset (optional) (default to 0)
    nftData := "0xf7c932351186c3a9053f313eefa16209c018f7f1dba8aa8ca7100400f7c31085" // string | request.getUserNftTransactions.nftHash (optional)
    txStatus := "processing" // string | request.getUserNftTransactions.txStatus (optional)
    types := "mint,deposit,transfer,onchain_withdrawal,offchain_withdrawal" // string | request.getUserNftTransactions.types (optional)
    hashes := "0xf7c932351186c3a9053f313eefa16209c018f7f1dba8aa8ca7100400f7c31085" // string | request.getUserNftTransactions.hashes (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetUserNftTransactionsV3(context.Background()).AccountId(accountId).Start(start).End(end).Limit(limit).Offset(offset).NftData(nftData).TxStatus(txStatus).Types(types).Hashes(hashes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetUserNftTransactionsV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserNftTransactionsV3`: GetUserNftTxsResponseV2
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetUserNftTransactionsV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserNftTransactionsV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int64** | request.getUserNftTransactions.accountId | 
 **start** | **int64** | request.getUserNftTransactions.start | [default to 0]
 **end** | **int64** | request.getUserNftTransactions.end | [default to 0]
 **limit** | **int32** | request.getUserNftTransactions.limit | [default to 0]
 **offset** | **int64** | request.getUserNftTransactions.offset | [default to 0]
 **nftData** | **string** | request.getUserNftTransactions.nftHash | 
 **txStatus** | **string** | request.getUserNftTransactions.txStatus | 
 **types** | **string** | request.getUserNftTransactions.types | 
 **hashes** | **string** | request.getUserNftTransactions.hashes | 

### Return type

[**GetUserNftTxsResponseV2**](GetUserNftTxsResponseV2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserRewardsV2

> GetAmmPoolsStatsResponse GetUserRewardsV2(ctx).Owner(owner).AmmPoolMarkets(ammPoolMarkets).Execute()

api.getPoolsStats.value



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := int64(789) // int64 |  (optional)
    ammPoolMarkets := "ammPoolMarkets_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetUserRewardsV2(context.Background()).Owner(owner).AmmPoolMarkets(ammPoolMarkets).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetUserRewardsV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserRewardsV2`: GetAmmPoolsStatsResponse
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetUserRewardsV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRewardsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **int64** |  | 
 **ammPoolMarkets** | **string** |  | 

### Return type

[**GetAmmPoolsStatsResponse**](GetAmmPoolsStatsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserRewardsV3

> GetAmmPoolsStatsResponse GetUserRewardsV3(ctx).Owner(owner).AmmPoolMarkets(ammPoolMarkets).Execute()

api.getPoolsStats.value



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := int64(789) // int64 |  (optional)
    ammPoolMarkets := "ammPoolMarkets_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetUserRewardsV3(context.Background()).Owner(owner).AmmPoolMarkets(ammPoolMarkets).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetUserRewardsV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserRewardsV3`: GetAmmPoolsStatsResponse
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetUserRewardsV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRewardsV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **int64** |  | 
 **ammPoolMarkets** | **string** |  | 

### Return type

[**GetAmmPoolsStatsResponse**](GetAmmPoolsStatsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserTournamentRankV3

> map[string]interface{} GetUserTournamentRankV3(ctx).Owner(owner).AmmPoolMarket(ammPoolMarket).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string |  (optional)
    ammPoolMarket := "ammPoolMarket_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetUserTournamentRankV3(context.Background()).Owner(owner).AmmPoolMarket(ammPoolMarket).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetUserTournamentRankV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserTournamentRankV3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetUserTournamentRankV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserTournamentRankV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** |  | 
 **ammPoolMarket** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserTradeAmountV3

> map[string]interface{} GetUserTradeAmountV3(ctx).AccountId(accountId).Markets(markets).Limit(limit).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := int64(789) // int64 |  (optional)
    markets := "markets_example" // string |  (optional)
    limit := int32(56) // int32 |  (optional) (default to 30)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetUserTradeAmountV3(context.Background()).AccountId(accountId).Markets(markets).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetUserTradeAmountV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserTradeAmountV3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetUserTradeAmountV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserTradeAmountV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int64** |  | 
 **markets** | **string** |  | 
 **limit** | **int32** |  | [default to 30]

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserTradesV2

> GetUserTradesResponseV2 GetUserTradesV2(ctx).AccountId(accountId).Market(market).OrderHash(orderHash).Offset(offset).Limit(limit).FromId(fromId).FillTypes(fillTypes).Execute()

Get user trade history



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := int32(1) // int32 | Account ID, some hash query APIs doesnt need it if in hash query mode, check require flag of each API to see if its a must.
    market := "LRC-ETH" // string | Trading pair (optional)
    orderHash := "0x9d114267e8b261457d567093c13cf3deea5f14c9235be26c6fa833dba12a9632" // string | If a value is provided, only returns the trades associated with the given order (optional)
    offset := int32(1) // int32 | Number of records to skip (optional) (default to 0)
    limit := int32(50) // int32 | Number of records to return (optional) (default to 0)
    fromId := int64(1) // int64 | Pagination of data to return records earlier than the requested ID (optional) (default to 0)
    fillTypes := "dex" // string | request.getUserTrades.fillTypes (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetUserTradesV2(context.Background()).AccountId(accountId).Market(market).OrderHash(orderHash).Offset(offset).Limit(limit).FromId(fromId).FillTypes(fillTypes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetUserTradesV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserTradesV2`: GetUserTradesResponseV2
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetUserTradesV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserTradesV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32** | Account ID, some hash query APIs doesnt need it if in hash query mode, check require flag of each API to see if its a must. | 
 **market** | **string** | Trading pair | 
 **orderHash** | **string** | If a value is provided, only returns the trades associated with the given order | 
 **offset** | **int32** | Number of records to skip | [default to 0]
 **limit** | **int32** | Number of records to return | [default to 0]
 **fromId** | **int64** | Pagination of data to return records earlier than the requested ID | [default to 0]
 **fillTypes** | **string** | request.getUserTrades.fillTypes | 

### Return type

[**GetUserTradesResponseV2**](GetUserTradesResponseV2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserTradesV3

> TradeList GetUserTradesV3(ctx).AccountId(accountId).Market(market).OrderHash(orderHash).Offset(offset).Limit(limit).FromId(fromId).FillTypes(fillTypes).Execute()

Get user trade history



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := int32(1) // int32 | Account ID, some hash query APIs doesnt need it if in hash query mode, check require flag of each API to see if its a must.
    market := "LRC-ETH" // string | Trading pair (optional)
    orderHash := "0x9d114267e8b261457d567093c13cf3deea5f14c9235be26c6fa833dba12a9632" // string | Order hash (optional)
    offset := int32(1) // int32 | Number of records to skip (optional) (default to 0)
    limit := int32(50) // int32 | Number of records to return (optional) (default to 0)
    fromId := int64(1) // int64 | The begin id of the query (optional) (default to 0)
    fillTypes := "dex" // string | request.getUserTxs.fillTypes (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetUserTradesV3(context.Background()).AccountId(accountId).Market(market).OrderHash(orderHash).Offset(offset).Limit(limit).FromId(fromId).FillTypes(fillTypes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetUserTradesV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserTradesV3`: TradeList
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetUserTradesV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserTradesV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32** | Account ID, some hash query APIs doesnt need it if in hash query mode, check require flag of each API to see if its a must. | 
 **market** | **string** | Trading pair | 
 **orderHash** | **string** | Order hash | 
 **offset** | **int32** | Number of records to skip | [default to 0]
 **limit** | **int32** | Number of records to return | [default to 0]
 **fromId** | **int64** | The begin id of the query | [default to 0]
 **fillTypes** | **string** | request.getUserTxs.fillTypes | 

### Return type

[**TradeList**](TradeList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserTransactionsV3

> map[string]interface{} GetUserTransactionsV3(ctx).AccountId(accountId).Start(start).End(end).Limit(limit).TokenSymbol(tokenSymbol).Offset(offset).Types(types).Hashes(hashes).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := int32(1) // int32 | Account ID, some hash query APIs doesnt need it if in hash query mode, check require flag of each API to see if its a must.
    start := int64(1578558098000) // int64 | Start time in milliseconds (optional) (default to 0)
    end := int64(1578558098000) // int64 | End time in milliseconds (optional) (default to 0)
    limit := int32(50) // int32 | Number of records to return (optional) (default to 0)
    tokenSymbol := "LRC" // string | Token to filter. If you want to return deposit records for all tokens, omit this parameter (optional)
    offset := int64(1) // int64 | Number of records to skip (optional) (default to 0)
    types := "deposit" // string | The type of the transactions to be queried (optional)
    hashes := "0xf7c932351186c3a9053f313eefa16209c018f7f1dba8aa8ca7100400f7c31085" // string | The hashes (split by ,) of the transactions, normally its L2 tx hash, except the deposit which uses L1 tx hash. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetUserTransactionsV3(context.Background()).AccountId(accountId).Start(start).End(end).Limit(limit).TokenSymbol(tokenSymbol).Offset(offset).Types(types).Hashes(hashes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetUserTransactionsV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserTransactionsV3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetUserTransactionsV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserTransactionsV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32** | Account ID, some hash query APIs doesnt need it if in hash query mode, check require flag of each API to see if its a must. | 
 **start** | **int64** | Start time in milliseconds | [default to 0]
 **end** | **int64** | End time in milliseconds | [default to 0]
 **limit** | **int32** | Number of records to return | [default to 0]
 **tokenSymbol** | **string** | Token to filter. If you want to return deposit records for all tokens, omit this parameter | 
 **offset** | **int64** | Number of records to skip | [default to 0]
 **types** | **string** | The type of the transactions to be queried | 
 **hashes** | **string** | The hashes (split by ,) of the transactions, normally its L2 tx hash, except the deposit which uses L1 tx hash. | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserTransfersV2

> GetUserTransferResponse GetUserTransfersV2(ctx).AccountId(accountId).Start(start).End(end).Status(status).Limit(limit).TokenSymbol(tokenSymbol).Offset(offset).TransferTypes(transferTypes).Execute()

Get user transfer list



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := int32(1) // int32 | Account ID, some hash query APIs doesnt need it if in hash query mode, check require flag of each API to see if its a must.
    start := int64(1578558098000) // int64 | Start time in milliseconds (optional) (default to 0)
    end := int64(1578558098000) // int64 | End time in milliseconds (optional) (default to 0)
    status := "processing,processed" // string | Comma separated status values (optional)
    limit := int32(50) // int32 | Number of records to return (optional) (default to 0)
    tokenSymbol := "LRC" // string | Token to filter. If you want to return deposit records for all tokens, omit this parameter (optional)
    offset := int64(1) // int64 | Number of records to skip (optional) (default to 0)
    transferTypes := "transfer, transfer_red" // string | request.getUserTxs.transferTypes (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetUserTransfersV2(context.Background()).AccountId(accountId).Start(start).End(end).Status(status).Limit(limit).TokenSymbol(tokenSymbol).Offset(offset).TransferTypes(transferTypes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetUserTransfersV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserTransfersV2`: GetUserTransferResponse
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetUserTransfersV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserTransfersV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32** | Account ID, some hash query APIs doesnt need it if in hash query mode, check require flag of each API to see if its a must. | 
 **start** | **int64** | Start time in milliseconds | [default to 0]
 **end** | **int64** | End time in milliseconds | [default to 0]
 **status** | **string** | Comma separated status values | 
 **limit** | **int32** | Number of records to return | [default to 0]
 **tokenSymbol** | **string** | Token to filter. If you want to return deposit records for all tokens, omit this parameter | 
 **offset** | **int64** | Number of records to skip | [default to 0]
 **transferTypes** | **string** | request.getUserTxs.transferTypes | 

### Return type

[**GetUserTransferResponse**](GetUserTransferResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserTransfersV3

> TransferDataList GetUserTransfersV3(ctx).AccountId(accountId).Start(start).End(end).Status(status).Limit(limit).TokenSymbol(tokenSymbol).Offset(offset).TransferTypes(transferTypes).Hashes(hashes).Execute()

Get user transfer list



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := int64(1) // int64 | Account ID, some hash query APIs doesnt need it if in hash query mode, check require flag of each API to see if its a must. (optional) (default to 0)
    start := int64(1578558098000) // int64 | Start time in milliseconds (optional) (default to 0)
    end := int64(1578558098000) // int64 | End time in milliseconds (optional) (default to 0)
    status := "processing,processed" // string | Comma separated status values (optional)
    limit := int32(50) // int32 | Number of records to return (optional) (default to 0)
    tokenSymbol := "LRC" // string | Token to filter. If you want to return deposit records for all tokens, omit this parameter (optional)
    offset := int64(1) // int64 | Number of records to skip (optional) (default to 0)
    transferTypes := "transfer, transfer_red" // string | request.getUserTxs.transferTypes (optional)
    hashes := "0xf7c932351186c3a9053f313eefa16209c018f7f1dba8aa8ca7100400f7c31085" // string | The hashes (split by ,) of the transactions, normally its L2 tx hash, except the deposit which uses L1 tx hash. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetUserTransfersV3(context.Background()).AccountId(accountId).Start(start).End(end).Status(status).Limit(limit).TokenSymbol(tokenSymbol).Offset(offset).TransferTypes(transferTypes).Hashes(hashes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetUserTransfersV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserTransfersV3`: TransferDataList
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetUserTransfersV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserTransfersV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int64** | Account ID, some hash query APIs doesnt need it if in hash query mode, check require flag of each API to see if its a must. | [default to 0]
 **start** | **int64** | Start time in milliseconds | [default to 0]
 **end** | **int64** | End time in milliseconds | [default to 0]
 **status** | **string** | Comma separated status values | 
 **limit** | **int32** | Number of records to return | [default to 0]
 **tokenSymbol** | **string** | Token to filter. If you want to return deposit records for all tokens, omit this parameter | 
 **offset** | **int64** | Number of records to skip | [default to 0]
 **transferTypes** | **string** | request.getUserTxs.transferTypes | 
 **hashes** | **string** | The hashes (split by ,) of the transactions, normally its L2 tx hash, except the deposit which uses L1 tx hash. | 

### Return type

[**TransferDataList**](TransferDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserUpdateV2

> GetUserAccountUpdateResponseV2 GetUserUpdateV2(ctx).AccountId(accountId).Start(start).End(end).Status(status).Limit(limit).Offset(offset).Execute()

Get password reset transactions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := int32(1) // int32 | Account ID, some hash query APIs doesnt need it if in hash query mode, check require flag of each API to see if its a must.
    start := int64(1578558098000) // int64 | Start time in milliseconds (optional) (default to 0)
    end := int64(1578558098000) // int64 | End time in milliseconds (optional) (default to 0)
    status := "processing,processed" // string | Comma separated status values (optional)
    limit := int32(50) // int32 | Number of records to return (optional) (default to 0)
    offset := int64(1) // int64 | Number of records to skip (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetUserUpdateV2(context.Background()).AccountId(accountId).Start(start).End(end).Status(status).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetUserUpdateV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserUpdateV2`: GetUserAccountUpdateResponseV2
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetUserUpdateV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserUpdateV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32** | Account ID, some hash query APIs doesnt need it if in hash query mode, check require flag of each API to see if its a must. | 
 **start** | **int64** | Start time in milliseconds | [default to 0]
 **end** | **int64** | End time in milliseconds | [default to 0]
 **status** | **string** | Comma separated status values | 
 **limit** | **int32** | Number of records to return | [default to 0]
 **offset** | **int64** | Number of records to skip | [default to 0]

### Return type

[**GetUserAccountUpdateResponseV2**](GetUserAccountUpdateResponseV2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserUpdateV3

> UserAccountUpdateDataList GetUserUpdateV3(ctx).AccountId(accountId).Start(start).End(end).Status(status).Limit(limit).Offset(offset).Execute()

Get password reset transactions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := int32(1) // int32 | Account ID, some hash query APIs doesnt need it if in hash query mode, check require flag of each API to see if its a must.
    start := int64(1578558098000) // int64 | Start time in milliseconds (optional) (default to 0)
    end := int64(1578558098000) // int64 | End time in milliseconds (optional) (default to 0)
    status := "processing,processed" // string | Comma separated status values (optional)
    limit := int32(50) // int32 | Number of records to return (optional) (default to 0)
    offset := int64(1) // int64 | Number of records to skip (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetUserUpdateV3(context.Background()).AccountId(accountId).Start(start).End(end).Status(status).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetUserUpdateV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserUpdateV3`: UserAccountUpdateDataList
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetUserUpdateV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserUpdateV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32** | Account ID, some hash query APIs doesnt need it if in hash query mode, check require flag of each API to see if its a must. | 
 **start** | **int64** | Start time in milliseconds | [default to 0]
 **end** | **int64** | End time in milliseconds | [default to 0]
 **status** | **string** | Comma separated status values | 
 **limit** | **int32** | Number of records to return | [default to 0]
 **offset** | **int64** | Number of records to skip | [default to 0]

### Return type

[**UserAccountUpdateDataList**](UserAccountUpdateDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserVipInfoV3

> BalanceV3 GetUserVipInfoV3(ctx).UserAddress(userAddress).Execute()

api.getUserVipInfo.value



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    userAddress := "userAddress_example" // string | request.getUserVipInfo.address

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetUserVipInfoV3(context.Background()).UserAddress(userAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetUserVipInfoV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserVipInfoV3`: BalanceV3
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetUserVipInfoV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserVipInfoV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userAddress** | **string** | request.getUserVipInfo.address | 

### Return type

[**BalanceV3**](BalanceV3.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserWithdrawalsV2

> GetUserOnchainWithdrawalResponseV2 GetUserWithdrawalsV2(ctx).AccountId(accountId).Start(start).End(end).Status(status).Limit(limit).TokenSymbol(tokenSymbol).Offset(offset).WithdrawalTypes(withdrawalTypes).Execute()

Get user onchain withdrawal history



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := int32(1) // int32 | Account ID, some hash query APIs doesnt need it if in hash query mode, check require flag of each API to see if its a must.
    start := int64(1578558098000) // int64 | Start time in milliseconds (optional) (default to 0)
    end := int64(1578558098000) // int64 | End time in milliseconds (optional) (default to 0)
    status := "processing,processed" // string | Comma separated status values (optional)
    limit := int32(50) // int32 | Number of records to return (optional) (default to 0)
    tokenSymbol := "LRC" // string | Token to filter. If you want to return deposit records for all tokens, omit this parameter (optional)
    offset := int64(1) // int64 | Number of records to skip (optional) (default to 0)
    withdrawalTypes := "force_withdrawal" // string | request.getUserTxs.withdrawalTypes (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetUserWithdrawalsV2(context.Background()).AccountId(accountId).Start(start).End(end).Status(status).Limit(limit).TokenSymbol(tokenSymbol).Offset(offset).WithdrawalTypes(withdrawalTypes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetUserWithdrawalsV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserWithdrawalsV2`: GetUserOnchainWithdrawalResponseV2
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetUserWithdrawalsV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserWithdrawalsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32** | Account ID, some hash query APIs doesnt need it if in hash query mode, check require flag of each API to see if its a must. | 
 **start** | **int64** | Start time in milliseconds | [default to 0]
 **end** | **int64** | End time in milliseconds | [default to 0]
 **status** | **string** | Comma separated status values | 
 **limit** | **int32** | Number of records to return | [default to 0]
 **tokenSymbol** | **string** | Token to filter. If you want to return deposit records for all tokens, omit this parameter | 
 **offset** | **int64** | Number of records to skip | [default to 0]
 **withdrawalTypes** | **string** | request.getUserTxs.withdrawalTypes | 

### Return type

[**GetUserOnchainWithdrawalResponseV2**](GetUserOnchainWithdrawalResponseV2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserWithdrawalsV3

> OnchainWithdrawalDataList GetUserWithdrawalsV3(ctx).AccountId(accountId).Start(start).End(end).Status(status).Limit(limit).TokenSymbol(tokenSymbol).Offset(offset).WithdrawalTypes(withdrawalTypes).Hashes(hashes).Execute()

Get user onchain withdrawal history



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := int64(1) // int64 | Account ID, some hash query APIs doesnt need it if in hash query mode, check require flag of each API to see if its a must. (default to 0)
    start := int64(1578558098000) // int64 | Start time in milliseconds (optional) (default to 0)
    end := int64(1578558098000) // int64 | End time in milliseconds (optional) (default to 0)
    status := "processing,processed" // string | Comma separated status values (optional)
    limit := int32(50) // int32 | Number of records to return (optional) (default to 0)
    tokenSymbol := "LRC" // string | Token to filter. If you want to return deposit records for all tokens, omit this parameter (optional)
    offset := int64(1) // int64 | Number of records to skip (optional) (default to 0)
    withdrawalTypes := "force_withdrawal" // string | request.getUserTxs.withdrawalTypes (optional)
    hashes := "0xf7c932351186c3a9053f313eefa16209c018f7f1dba8aa8ca7100400f7c31085" // string | The hashes (split by ,) of the transactions, normally its L2 tx hash, except the deposit which uses L1 tx hash. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetUserWithdrawalsV3(context.Background()).AccountId(accountId).Start(start).End(end).Status(status).Limit(limit).TokenSymbol(tokenSymbol).Offset(offset).WithdrawalTypes(withdrawalTypes).Hashes(hashes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetUserWithdrawalsV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserWithdrawalsV3`: OnchainWithdrawalDataList
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetUserWithdrawalsV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserWithdrawalsV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int64** | Account ID, some hash query APIs doesnt need it if in hash query mode, check require flag of each API to see if its a must. | [default to 0]
 **start** | **int64** | Start time in milliseconds | [default to 0]
 **end** | **int64** | End time in milliseconds | [default to 0]
 **status** | **string** | Comma separated status values | 
 **limit** | **int32** | Number of records to return | [default to 0]
 **tokenSymbol** | **string** | Token to filter. If you want to return deposit records for all tokens, omit this parameter | 
 **offset** | **int64** | Number of records to skip | [default to 0]
 **withdrawalTypes** | **string** | request.getUserTxs.withdrawalTypes | 
 **hashes** | **string** | The hashes (split by ,) of the transactions, normally its L2 tx hash, except the deposit which uses L1 tx hash. | 

### Return type

[**OnchainWithdrawalDataList**](OnchainWithdrawalDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWithdrawalAgentsV3

> AgentInfo GetWithdrawalAgentsV3(ctx).TokenId(tokenId).Amount(amount).Execute()

api.getWithdrawalAgents.value



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tokenId := int32(56) // int32 |  (optional)
    amount := "amount_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.GetWithdrawalAgentsV3(context.Background()).TokenId(tokenId).Amount(amount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.GetWithdrawalAgentsV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWithdrawalAgentsV3`: AgentInfo
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.GetWithdrawalAgentsV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWithdrawalAgentsV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenId** | **int32** |  | 
 **amount** | **string** |  | 

### Return type

[**AgentInfo**](AgentInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpenHebaoAccountV3Hebao

> map[string]interface{} OpenHebaoAccountV3Hebao(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.OpenHebaoAccountV3Hebao(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.OpenHebaoAccountV3Hebao``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OpenHebaoAccountV3Hebao`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.OpenHebaoAccountV3Hebao`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOpenHebaoAccountV3HebaoRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrepareHebaoPayV3Hebao

> map[string]interface{} PrepareHebaoPayV3Hebao(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.PrepareHebaoPayV3Hebao(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.PrepareHebaoPayV3Hebao``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrepareHebaoPayV3Hebao`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.PrepareHebaoPayV3Hebao`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPrepareHebaoPayV3HebaoRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefundHebaoPayV3Hebao

> map[string]interface{} RefundHebaoPayV3Hebao(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.RefundHebaoPayV3Hebao(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.RefundHebaoPayV3Hebao``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RefundHebaoPayV3Hebao`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.RefundHebaoPayV3Hebao`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRefundHebaoPayV3HebaoRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchTokenMetadataV2

> map[string]interface{} SearchTokenMetadataV2(ctx).TokenAddresses(tokenAddresses).Network(network).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tokenAddresses := "0x0000000000000000000000000000000000000000,0x0000000000000000000000000000000000000001" // string | token address to be queried
    network := "ETHEREUM or ARBITRUM" // string | network

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.SearchTokenMetadataV2(context.Background()).TokenAddresses(tokenAddresses).Network(network).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.SearchTokenMetadataV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchTokenMetadataV2`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.SearchTokenMetadataV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchTokenMetadataV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenAddresses** | **string** | token address to be queried | 
 **network** | **string** | network | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchTokenV2

> map[string]interface{} SearchTokenV2(ctx).Network(network).Key(key).Offset(offset).Limit(limit).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    network := "network_example" // string |  (optional)
    key := "key_example" // string |  (optional)
    offset := int64(789) // int64 |  (optional)
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.SearchTokenV2(context.Background()).Network(network).Key(key).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.SearchTokenV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchTokenV2`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.SearchTokenV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchTokenV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **network** | **string** |  | 
 **key** | **string** |  | 
 **offset** | **int64** |  | 
 **limit** | **int32** |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendLuckyTokenV3

> SubmitOffChainRequestItem SendLuckyTokenV3(ctx).Body(body).Execute()

api.sendLuckyToken.value



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewSendLuckyTokenRequestV3("1", int64(5), "Lucky Draw for YOU", "true", int64(1), int32(1519217383), int64(1)) // SendLuckyTokenRequestV3 | api.sendLuckyToken.implicit.value

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.SendLuckyTokenV3(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.SendLuckyTokenV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendLuckyTokenV3`: SubmitOffChainRequestItem
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.SendLuckyTokenV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendLuckyTokenV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SendLuckyTokenRequestV3**](SendLuckyTokenRequestV3.md) | api.sendLuckyToken.implicit.value | 

### Return type

[**SubmitOffChainRequestItem**](SubmitOffChainRequestItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendTransactionV3

> ForwardEthTxResponse SendTransactionV3(ctx).Body(body).Execute()

Send a raw Ethereum transaction



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewForwardEthTxRequest("0xf8aa0a8502540be400830186a0949032dbf5669341c3d95bc02b4bde90e4e051db3580b844095ea7b30000000000000000000000001d307532a97879b866a6fe33bf4a517bd28de854ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff25a071af95d2a3eb01bb11f84d9c66d09924d97a2b3ce38e0079a2d56dfb13035df0a07404b3a48c40cc3e7b9893d88512ef526a67bf3c274548ec09f94386e0062fc2") // ForwardEthTxRequest | Body of send raw transaction

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.SendTransactionV3(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.SendTransactionV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendTransactionV3`: ForwardEthTxResponse
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.SendTransactionV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendTransactionV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ForwardEthTxRequest**](ForwardEthTxRequest.md) | Body of send raw transaction | 

### Return type

[**ForwardEthTxResponse**](ForwardEthTxResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetReferrerV3

> map[string]interface{} SetReferrerV3(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.SetReferrerV3(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.SetReferrerV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetReferrerV3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.SetReferrerV3`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSetReferrerV3Request struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitAmmPoolExitV2

> SubmitOffChainRequestResponse SubmitAmmPoolExitV2(ctx).Body(body).Execute()

Exit an AMM pool



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewAmmPoolExitRequest("0xe25c94ba036d91b48833acb637f719038f07372d", "0xe25c94ba036d91b48833acb637f719038f07372d", "100000000000000", int64(1), "1000000000000000000", "1000000000000000000", int32(1598431481), "0x30f2f3cf420ccd2a49533cf38b27b7cf06019019d8d7216f79ca5efccfbab66671388be422ac509aed5d76315862cd625e994865d4737ea5d93b712ddb9e36871b", "0x30f2f3cf420ccd2a49533cf38b27b7cf06019019d8d7216f79ca5efccfbab66671388be422ac509aed5d76315862cd625e994865d4737ea5d93b712ddb9e36871b") // AmmPoolExitRequest | api.submitAmmPoolExit.implicit.value

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.SubmitAmmPoolExitV2(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.SubmitAmmPoolExitV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitAmmPoolExitV2`: SubmitOffChainRequestResponse
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.SubmitAmmPoolExitV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitAmmPoolExitV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AmmPoolExitRequest**](AmmPoolExitRequest.md) | api.submitAmmPoolExit.implicit.value | 

### Return type

[**SubmitOffChainRequestResponse**](SubmitOffChainRequestResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitAmmPoolExitV3

> SubmitOffChainRequestItem SubmitAmmPoolExitV3(ctx).Body(body).Execute()

Exit an AMM pool



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewAmmPoolExitRequestV3("0xe25c94ba036d91b48833acb637f719038f07372d", "0xe25c94ba036d91b48833acb637f719038f07372d", *openapiclient.NewAmmPoolExitTokens([]openapiclient.TokenVolumeV3{*openapiclient.NewTokenVolumeV3(int32(0), "1000000000000")}, "1000000000"), int64(1), "1000000000000000000", int32(1598431481)) // AmmPoolExitRequestV3 | AMM exit request params

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.SubmitAmmPoolExitV3(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.SubmitAmmPoolExitV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitAmmPoolExitV3`: SubmitOffChainRequestItem
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.SubmitAmmPoolExitV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitAmmPoolExitV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AmmPoolExitRequestV3**](AmmPoolExitRequestV3.md) | AMM exit request params | 

### Return type

[**SubmitOffChainRequestItem**](SubmitOffChainRequestItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitAmmPoolJoinV2

> SubmitOffChainRequestResponse SubmitAmmPoolJoinV2(ctx).Body(body).Execute()

Join into AMM pool



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewAmmPoolJoinRequest("0xe25c94ba036d91b48833acb637f719038f07372d", "0xe25c94ba036d91b48833acb637f719038f07372d", "10000000000,20000000000", "1", "1000000000", "1000000000000000000", int32(1598431481), "0x30f2f3cf420ccd2a49533cf38b27b7cf06019019d8d7216f79ca5efccfbab66671388be422ac509aed5d76315862cd625e994865d4737ea5d93b712ddb9e36871b", "0x30f2f3cf420ccd2a49533cf38b27b7cf06019019d8d7216f79ca5efccfbab66671388be422ac509aed5d76315862cd625e994865d4737ea5d93b712ddb9e36871b") // AmmPoolJoinRequest | api.AmmPoolJoinRequest.implicit.value

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.SubmitAmmPoolJoinV2(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.SubmitAmmPoolJoinV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitAmmPoolJoinV2`: SubmitOffChainRequestResponse
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.SubmitAmmPoolJoinV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitAmmPoolJoinV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AmmPoolJoinRequest**](AmmPoolJoinRequest.md) | api.AmmPoolJoinRequest.implicit.value | 

### Return type

[**SubmitOffChainRequestResponse**](SubmitOffChainRequestResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitAmmPoolJoinV3

> SubmitOffChainRequestItem SubmitAmmPoolJoinV3(ctx).Body(body).Execute()

Join into AMM pool



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewAmmPoolJoinRequestV3("0xe25c94ba036d91b48833acb637f719038f07372d", "0xe25c94ba036d91b48833acb637f719038f07372d", *openapiclient.NewAmmPoolJoinTokens([]openapiclient.TokenVolumeV3{*openapiclient.NewTokenVolumeV3(int32(0), "1000000000000")}, *openapiclient.NewTokenVolumeV3(int32(0), "1000000000000")), "1", "1000000000000000000", int32(1598431481)) // AmmPoolJoinRequestV3 | AMM join request params

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.SubmitAmmPoolJoinV3(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.SubmitAmmPoolJoinV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitAmmPoolJoinV3`: SubmitOffChainRequestItem
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.SubmitAmmPoolJoinV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitAmmPoolJoinV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AmmPoolJoinRequestV3**](AmmPoolJoinRequestV3.md) | AMM join request params | 

### Return type

[**SubmitOffChainRequestItem**](SubmitOffChainRequestItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitDualAuthTransferV2

> SubmitOffChainRequestResponse SubmitDualAuthTransferV2(ctx).Body(body).Execute()

Submit dual authority transfer



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewDualAuthTransferRequest("0xABCD", int32(1), "0xABCD", int32(1), "1000000000000000000", int32(1), "1000000000000000000", int64(1), int32(1598431481), "0xX..{64}..XY..{64}..YZ..{64}..Z", "13375450901292179417154974849571793069911517354720397125027633242680470075859", "13375450901292179417154974849571793069911517354720397125027633242680470075859", "0xX..{64}..XY..{64}..YZ..{64}..Z", int64(1), "0xABCD") // DualAuthTransferRequest | Submit dual authority transfer post message body

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.SubmitDualAuthTransferV2(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.SubmitDualAuthTransferV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitDualAuthTransferV2`: SubmitOffChainRequestResponse
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.SubmitDualAuthTransferV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitDualAuthTransferV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DualAuthTransferRequest**](DualAuthTransferRequest.md) | Submit dual authority transfer post message body | 

### Return type

[**SubmitOffChainRequestResponse**](SubmitOffChainRequestResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitMintNftV3

> SubmitMintNftResponseItem SubmitMintNftV3(ctx).Body(body).Execute()

Mint a NFT token in Loopring L2



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewSubmitNftMintRequest("1", int64(1), "0xbbbbca6a901c926f240b89eacb641d8aec7aeafd", int64(1), int32(123), "0xbbbbca6a901c926f240b89eacb641d8aec7aeafd", "0xf7c932351186c3a9053f313eefa16209c018f7f1dba8aa8ca7100400f7c31085", "1", int64(1567053142), int32(1), *openapiclient.NewTokenAmountInfo(int32(0), "1")) // SubmitNftMintRequest | Mint a NFT token in Loopring L2

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.SubmitMintNftV3(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.SubmitMintNftV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitMintNftV3`: SubmitMintNftResponseItem
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.SubmitMintNftV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitMintNftV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SubmitNftMintRequest**](SubmitNftMintRequest.md) | Mint a NFT token in Loopring L2 | 

### Return type

[**SubmitMintNftResponseItem**](SubmitMintNftResponseItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitNftTradeV3

> NftTradeResponse SubmitNftTradeV3(ctx).Body(body).Execute()

Settle down an NFT trade which has two matched orders



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewSubmitNftTradeRequestV3(*openapiclient.NewSubmitNftOrderRequestV3("1", int64(1), int32(1), *openapiclient.NewNftTokenAmountInfo(int32(0), "1"), *openapiclient.NewNftTokenAmountInfo(int32(0), "1"), int64(1567053142), int32(20), "13375450901292179417154974849571793069911517354720397125027633242680470075859"), int32(100), *openapiclient.NewSubmitNftOrderRequestV3("1", int64(1), int32(1), *openapiclient.NewNftTokenAmountInfo(int32(0), "1"), *openapiclient.NewNftTokenAmountInfo(int32(0), "1"), int64(1567053142), int32(20), "13375450901292179417154974849571793069911517354720397125027633242680470075859"), int32(100)) // SubmitNftTradeRequestV3 | Settle down an NFT trade which has two matched orders

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.SubmitNftTradeV3(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.SubmitNftTradeV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitNftTradeV3`: NftTradeResponse
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.SubmitNftTradeV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitNftTradeV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SubmitNftTradeRequestV3**](SubmitNftTradeRequestV3.md) | Settle down an NFT trade which has two matched orders | 

### Return type

[**NftTradeResponse**](NftTradeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitNftTransferV3

> SubmitOffChainRequestItem SubmitNftTransferV3(ctx).Body(body).Execute()

Submit internal NFT transfer



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewNftTransferRequest("1", int64(1), "0xABCD", int64(1), "0xCDEF", *openapiclient.NewNftTokenAmountInfo(int32(0), "1"), *openapiclient.NewTokenAmountInfo(int32(0), "1"), int64(1), int32(1598431481)) // NftTransferRequest | api.submitNftTransfer.implicit.value

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.SubmitNftTransferV3(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.SubmitNftTransferV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitNftTransferV3`: SubmitOffChainRequestItem
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.SubmitNftTransferV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitNftTransferV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NftTransferRequest**](NftTransferRequest.md) | api.submitNftTransfer.implicit.value | 

### Return type

[**SubmitOffChainRequestItem**](SubmitOffChainRequestItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitOffChainNftWithdrawalV3

> SubmitOffChainRequestItem SubmitOffChainNftWithdrawalV3(ctx).Body(body).Execute()

Withdraw a NFT token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewNftOffChainWithdrawalRequestV3("1", int64(1003), "0xbbbbca6a901c926f240b89eacb641d8aec7aeafd", *openapiclient.NewNftTokenAmountInfo(int32(0), "1"), *openapiclient.NewTokenAmountInfo(int32(0), "1"), int64(1), int32(1519217383), "0x12345678") // NftOffChainWithdrawalRequestV3 | Withdraw a NFT token

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.SubmitOffChainNftWithdrawalV3(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.SubmitOffChainNftWithdrawalV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitOffChainNftWithdrawalV3`: SubmitOffChainRequestItem
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.SubmitOffChainNftWithdrawalV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitOffChainNftWithdrawalV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NftOffChainWithdrawalRequestV3**](NftOffChainWithdrawalRequestV3.md) | Withdraw a NFT token | 

### Return type

[**SubmitOffChainRequestItem**](SubmitOffChainRequestItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitOffChainWithdrawalV2

> SubmitOffChainRequestResponse SubmitOffChainWithdrawalV2(ctx).Body(body).Execute()

Submit offchain withdraw request



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewOffChainWithdrawalRequest("1", int64(1), "1", int32(1), "1000000000000000000", int32(1), "1000000000000000000", int64(1), int32(1519217383), int32(1519217383), "0x12345678") // OffChainWithdrawalRequest | Submit offchain withdraw request post message body

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.SubmitOffChainWithdrawalV2(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.SubmitOffChainWithdrawalV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitOffChainWithdrawalV2`: SubmitOffChainRequestResponse
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.SubmitOffChainWithdrawalV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitOffChainWithdrawalV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OffChainWithdrawalRequest**](OffChainWithdrawalRequest.md) | Submit offchain withdraw request post message body | 

### Return type

[**SubmitOffChainRequestResponse**](SubmitOffChainRequestResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitOffChainWithdrawalV3

> SubmitOffChainRequestItem SubmitOffChainWithdrawalV3(ctx).Body(body).Execute()

Submit offchain withdraw request



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewOffChainWithdrawalRequestV3("1", int64(1003), "0xbbbbca6a901c926f240b89eacb641d8aec7aeafd", *openapiclient.NewTokenVolumeV3(int32(0), "1000000000000"), *openapiclient.NewTokenVolumeV3(int32(0), "1000000000000"), int64(1), int32(1519217383), "0x12345678") // OffChainWithdrawalRequestV3 | Submit offchain withdraw request post message body

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.SubmitOffChainWithdrawalV3(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.SubmitOffChainWithdrawalV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitOffChainWithdrawalV3`: SubmitOffChainRequestItem
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.SubmitOffChainWithdrawalV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitOffChainWithdrawalV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OffChainWithdrawalRequestV3**](OffChainWithdrawalRequestV3.md) | Submit offchain withdraw request post message body | 

### Return type

[**SubmitOffChainRequestItem**](SubmitOffChainRequestItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitOrderV2V2

> SubmitOrderResponseV2 SubmitOrderV2V2(ctx).Body(body).Execute()

api.submitOrderV2.value



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewSubmitOrderRequest("1", int32(1), int64(1), int32(0), int32(2), "1000000000000000000", "1000000000000000000", "true", int64(1567053142), int32(20), "13375450901292179417154974849571793069911517354720397125027633242680470075859") // SubmitOrderRequest | api.submitOrderV2.implicit.value

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.SubmitOrderV2V2(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.SubmitOrderV2V2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitOrderV2V2`: SubmitOrderResponseV2
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.SubmitOrderV2V2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitOrderV2V2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SubmitOrderRequest**](SubmitOrderRequest.md) | api.submitOrderV2.implicit.value | 

### Return type

[**SubmitOrderResponseV2**](SubmitOrderResponseV2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitOrderV3V3

> SubmitOrderResponseItemV3 SubmitOrderV3V3(ctx).Body(body).Execute()

Submit an order



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewSubmitOrderRequestV3("1", int64(1), int32(1), *openapiclient.NewTokenVolumeV3(int32(0), "1000000000000"), *openapiclient.NewTokenVolumeV3(int32(0), "1000000000000"), "false", "false", int64(1567053142), int32(20), "13375450901292179417154974849571793069911517354720397125027633242680470075859") // SubmitOrderRequestV3 | Submit order message body

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.SubmitOrderV3V3(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.SubmitOrderV3V3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitOrderV3V3`: SubmitOrderResponseItemV3
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.SubmitOrderV3V3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitOrderV3V3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SubmitOrderRequestV3**](SubmitOrderRequestV3.md) | Submit order message body | 

### Return type

[**SubmitOrderResponseItemV3**](SubmitOrderResponseItemV3.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitRecommendedMarketsV3

> GetExchangeFeeInfoResponseData SubmitRecommendedMarketsV3(ctx).Execute()

api.getRecommendedMarkets.value



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.SubmitRecommendedMarketsV3(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.SubmitRecommendedMarketsV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitRecommendedMarketsV3`: GetExchangeFeeInfoResponseData
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.SubmitRecommendedMarketsV3`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSubmitRecommendedMarketsV3Request struct via the builder pattern


### Return type

[**GetExchangeFeeInfoResponseData**](GetExchangeFeeInfoResponseData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitRewardV3

> map[string]interface{} SubmitRewardV3(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.SubmitRewardV3(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.SubmitRewardV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitRewardV3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.SubmitRewardV3`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSubmitRewardV3Request struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitTransferV2

> SubmitOffChainRequestResponse SubmitTransferV2(ctx).Body(body).Execute()

Submit internal transfer



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewOriginTransferRequest("1", int64(1), "0xABCD", int64(1), "0xCDEF", int32(1), "1000000000000000000", int32(1), "1000000000000000000", int64(1), int32(1598431481)) // OriginTransferRequest | Submit internal transfer post message body

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.SubmitTransferV2(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.SubmitTransferV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitTransferV2`: SubmitOffChainRequestResponse
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.SubmitTransferV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitTransferV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OriginTransferRequest**](OriginTransferRequest.md) | Submit internal transfer post message body | 

### Return type

[**SubmitOffChainRequestResponse**](SubmitOffChainRequestResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitTransferV3

> SubmitOffChainRequestItem SubmitTransferV3(ctx).Body(body).Execute()

Submit internal transfer



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewOriginTransferRequestV3("1", int64(1), "0xABCD", int64(1), "0xCDEF", *openapiclient.NewTokenVolumeV3(int32(0), "1000000000000"), *openapiclient.NewTokenVolumeV3(int32(0), "1000000000000"), int64(1), int32(1598431481)) // OriginTransferRequestV3 | Submit internal transfer post message body

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.SubmitTransferV3(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.SubmitTransferV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitTransferV3`: SubmitOffChainRequestItem
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.SubmitTransferV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitTransferV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OriginTransferRequestV3**](OriginTransferRequestV3.md) | Submit internal transfer post message body | 

### Return type

[**SubmitOffChainRequestItem**](SubmitOffChainRequestItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitUpdateAccountV2

> SubmitOffChainRequestResponse SubmitUpdateAccountV2(ctx).Body(body).Execute()

Update account EDDSA key



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewUpdateAccountRequest("1", "0xB4A70168340C75119523019f79F5Ffd9c60DceC7", int64(1), int32(1583183141), int32(1), "0xABCD", "0xABCD", int32(1), "1234") // UpdateAccountRequest | api.submitUpdateAccount.implicit.value

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.SubmitUpdateAccountV2(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.SubmitUpdateAccountV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitUpdateAccountV2`: SubmitOffChainRequestResponse
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.SubmitUpdateAccountV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitUpdateAccountV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateAccountRequest**](UpdateAccountRequest.md) | api.submitUpdateAccount.implicit.value | 

### Return type

[**SubmitOffChainRequestResponse**](SubmitOffChainRequestResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitUpdateAccountV3

> SubmitOffChainRequestItem SubmitUpdateAccountV3(ctx).Body(body).Execute()

Update account EDDSA key



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewUpdateAccountRequestV3("1", "0xB4A70168340C75119523019f79F5Ffd9c60DceC7", int64(1), *openapiclient.NewPublicKey("0x241707bcc6d7a4ccf10304be248d343a527e85f61b45d721544d027cc1f2fb5f", "0x302f3a521dbdd1d0eb1944c8323d4ac3b3e9c9201f4aa43a2565054886369d9c"), *openapiclient.NewTokenVolumeV3(int32(0), "1000000000000"), int32(1583183141), int32(1)) // UpdateAccountRequestV3 | api.submitUpdateAccount.implicit.value

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.SubmitUpdateAccountV3(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.SubmitUpdateAccountV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitUpdateAccountV3`: SubmitOffChainRequestItem
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.SubmitUpdateAccountV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitUpdateAccountV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateAccountRequestV3**](UpdateAccountRequestV3.md) | api.submitUpdateAccount.implicit.value | 

### Return type

[**SubmitOffChainRequestItem**](SubmitOffChainRequestItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TokenMetadatasV2

> map[string]interface{} TokenMetadatasV2(ctx).Network(network).Offset(offset).Limit(limit).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    network := "network_example" // string |  (optional)
    offset := int64(789) // int64 |  (optional)
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.TokenMetadatasV2(context.Background()).Network(network).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.TokenMetadatasV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TokenMetadatasV2`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.TokenMetadatasV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTokenMetadatasV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **network** | **string** |  | 
 **offset** | **int64** |  | 
 **limit** | **int32** |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateNftOrderV3

> SubmitOffChainRequestItem ValidateNftOrderV3(ctx).Body(body).Execute()

Validate a NFT order



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewSubmitNftOrderRequestV3("1", int64(1), int32(1), *openapiclient.NewNftTokenAmountInfo(int32(0), "1"), *openapiclient.NewNftTokenAmountInfo(int32(0), "1"), int64(1567053142), int32(20), "13375450901292179417154974849571793069911517354720397125027633242680470075859") // SubmitNftOrderRequestV3 | Validate a NFT order

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.ValidateNftOrderV3(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.ValidateNftOrderV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ValidateNftOrderV3`: SubmitOffChainRequestItem
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.ValidateNftOrderV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateNftOrderV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SubmitNftOrderRequestV3**](SubmitNftOrderRequestV3.md) | Validate a NFT order | 

### Return type

[**SubmitOffChainRequestItem**](SubmitOffChainRequestItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyAllEcdsaV3

> map[string]interface{} VerifyAllEcdsaV3(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.VerifyAllEcdsaV3(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.VerifyAllEcdsaV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VerifyAllEcdsaV3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.VerifyAllEcdsaV3`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyAllEcdsaV3Request struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WithdrawLuckyTokenV3

> SubmitOffChainRequestItem WithdrawLuckyTokenV3(ctx).Body(body).Execute()

api.withdrawLuckyToken.value



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewWithdrawLuckyTokenRequestV3(int32(123), int32(123), "0xbbbbca6a901c926f240b89eacb641d8aec7aeafd", "0xbbbbca6a901c926f240b89eacb641d8aec7aeafd") // WithdrawLuckyTokenRequestV3 | api.withdrawLuckyToken.implicit.value

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoopringDEXRestfulAPIApi.WithdrawLuckyTokenV3(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoopringDEXRestfulAPIApi.WithdrawLuckyTokenV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WithdrawLuckyTokenV3`: SubmitOffChainRequestItem
    fmt.Fprintf(os.Stdout, "Response from `LoopringDEXRestfulAPIApi.WithdrawLuckyTokenV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWithdrawLuckyTokenV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**WithdrawLuckyTokenRequestV3**](WithdrawLuckyTokenRequestV3.md) | api.withdrawLuckyToken.implicit.value | 

### Return type

[**SubmitOffChainRequestItem**](SubmitOffChainRequestItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

