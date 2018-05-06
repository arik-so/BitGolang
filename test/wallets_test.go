package BitGo

import (
  "../src"
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestGetWallet(t *testing.T) {
  bitgo := src.BitGoSDK("http://localhost", 3080)
  bitgo.Authenticate("arik+test007@bitgo.com", "test@bitgo.com", "0000000")
  wallet := bitgo.Coin("tltc").Wallets().Get("5984e0b26b423fda07a5d6d2d226c87a")
  assert.Equal(t, "5984e0b26b423fda07a5d6d2d226c87a", wallet.Id())
}

func TestListWallets(t *testing.T) {
  bitgo := src.BitGoSDK("http://localhost", 3080)
  bitgo.Authenticate("arik+test007@bitgo.com", "test@bitgo.com", "0000000")
  wallets := bitgo.Coin("tltc").Wallets().List()
  assert.Equal(t, 14, len(wallets))
}
