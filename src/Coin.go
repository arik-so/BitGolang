package src

type Coin struct {
  bitgo BitGo
  coin string
}

func (coin Coin) Url(extension string) string {
  return coin.bitgo.UrlV2("/" + coin.coin + extension)
}

func (coin Coin) Wallets() *Wallets{
  wallets := new(Wallets)
  wallets.coin = coin
  return wallets
}
