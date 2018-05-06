package BitGo

import (
  "github.com/stretchr/testify/assert"
  "../src"
  "testing"
)


func TestCoinUrl(t *testing.T) {
 bitgo := src.BitGoSDK("https://localhost", 3000)
 coin := bitgo.Coin("tbtc")
 coinUrl := coin.Url("/public/block/latest")
 assert.Equal(t, "https://localhost:3000/api/v2/tbtc/public/block/latest", coinUrl)
}
