package src

import "github.com/tidwall/gjson"

type Wallet struct {
  coin Coin
  data gjson.Result
}

func (wallet Wallet) Id() string {
  return wallet.data.Get("id").Str
}

func (wallet Wallet) Url(extension string) string {
  return wallet.coin.Url("/wallet/" + wallet.Id() + extension)
}

func (wallet Wallet) Balance() string {
  return wallet.data.Get("balanceString").Str
}

func (wallet Wallet) ConfirmedBalance() string {
  return wallet.data.Get("confirmedBalanceString").Str
}

func (wallet Wallet) SpendableBalance() string {
  return wallet.data.Get("spendableBalanceString").Str
}

func (wallet Wallet) SendMany(params map[string]interface{}) gjson.Result {
  walletUrl := wallet.Url("/sendmany/")
  println(walletUrl)
  sendResult := wallet.coin.bitgo.Post(walletUrl, params)
  return sendResult
}

func (wallet Wallet) Data() gjson.Result {
  return wallet.data
}
