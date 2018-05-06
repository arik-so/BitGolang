package BitGo

import (
  "../src"
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestSendMany(t *testing.T) {
  bitgo := src.BitGoSDK("http://localhost", 3080)
  bitgo.Authenticate("arik+test007@bitgo.com", "test@bitgo.com", "0000000")
  wallet := bitgo.Coin("tltc").Wallets().Get("5984e0b26b423fda07a5d6d2d226c87a")
  sendParams := make(map[string]interface{})
  sendParams["walletPassphrase"] = "test@bitgo.com"
  sendParams["otp"] = "0000000"

  recipient1 := make(map[string]interface{})
  recipient1["address"] = "QNc4RFAcbvqmtrR1kR2wbGLCx6tEvojFYE"
  recipient1["amount"] = 600000

  recipient2 := make(map[string]interface{})
  recipient2["address"] = "QPExG8YhRmuMD3W1JM7iVUN8VxH4ugVhZy"
  recipient2["amount"] = 400000

  sendParams["recipients"] = [...]map[string]interface{}{recipient1, recipient2}

  sendData := wallet.SendMany(sendParams)
  txid := sendData.Get("txid").String()
  assert.Equal(t, 64, len(txid))
}
