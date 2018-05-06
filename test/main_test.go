package BitGo

import (
  "github.com/stretchr/testify/assert"
  "../src"
  "testing"
)


func TestBuild(t *testing.T) {
 bitgo := src.BitGoSDK("https://localhost", 3000)
 bitgoUrl := bitgo.Url("/user/login")
 assert.Equal(t, "https://localhost:3000/api/v1/user/login", bitgoUrl)
}

func TestAuthentication(t *testing.T) {
 bitgo := src.BitGoSDK("http://localhost", 3080)
 authentication := bitgo.Authenticate("arik+test007@bitgo.com", "test@bitgo.com", "0000000")
 assert.Equal(t, "bearer", authentication.Get("token_type").String())
 assert.Equal(t, int64(3600), authentication.Get("expires_in").Int())
 assert.Equal(t, "arik+test007@bitgo.com", authentication.Get("user.username").String())
 assert.Equal(t, "US/Pacific", authentication.Get("user.timezone").String())
}

func TestGenesisBlock(t *testing.T) {
  bitgo := src.BitGoSDK("https://test.bitgo.com", 443)
  genesisBlock := bitgo.Get(bitgo.UrlV2("/tbtc/public/block/000000000933ea01ad0ee984209779baaec3ced90fa3f408719526f8d77f4943"))
  assert.Equal(t, "000000000933ea01ad0ee984209779baaec3ced90fa3f408719526f8d77f4943", genesisBlock.Get("id").String())
  assert.Equal(t, "4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b", genesisBlock.Get("merkleRoot").String())
}

/*
func TestTransactionSigning(t *testing.T) {
  bitgo := src.BitGoSDK("http://localhost", 3080)

  postBody := make(map[string]interface{})
  // postBody["tx"] = ""
  postBody["wallet_passphrase"] = ""
  postBody["address"] = ""
  postBody["amount"] = ""
  postBody["fee_rate"] = ""

  // wallet.send("")

  // genesisBlock := bitgo.Post(.Url("/wallet/5abcdefgh/sendcoins"), postBody)
  assert.Equal(t, "000000000933ea01ad0ee984209779baaec3ced90fa3f408719526f8d77f4943", genesisBlock.Get("id").String())
  assert.Equal(t, "4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b", genesisBlock.Get("merkleRoot").String())
}
*/
