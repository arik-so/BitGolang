package src

type Wallets struct {
  coin Coin
}

/**

 */
func (wallets Wallets) Get(id string) *Wallet {
  walletUrl := wallets.coin.Url("/wallet/" + id)
  walletData := wallets.coin.bitgo.Get(walletUrl)
  walletObject := new(Wallet)
  walletObject.coin = wallets.coin
  walletObject.data = walletData
  return walletObject
}

/**

 */
func (wallets Wallets) List() []*Wallet {
  walletsUrl := wallets.coin.Url("/wallet")
  walletData := wallets.coin.bitgo.Get(walletsUrl)
  walletList := walletData.Get("wallets").Array()
  var walletArray []*Wallet
  for _, currentWalletData := range walletList{
    currentWalletObject := new(Wallet)
    currentWalletObject.data = currentWalletData
    currentWalletObject.coin = wallets.coin
    walletArray = append(walletArray, currentWalletObject)
  }
  return walletArray
}
