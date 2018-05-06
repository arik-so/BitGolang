package src

import (
  "strconv"
  "encoding/json"
  "fmt"
  "net/http"
  "bytes"
  "io/ioutil"
  "github.com/tidwall/gjson"
)

type BitGo struct {
  baseUrl string
  accessToken string
}

func BitGoSDK(baseUrl string, port int) *BitGo {
  bitgo := new(BitGo)
  bitgo.baseUrl = baseUrl + ":" + strconv.Itoa(port)
  return bitgo
}

func (bitgo *BitGo) Authenticate(username string, password string, otp string) gjson.Result {
  loginUrl := bitgo.Url("/user/login")

  // build a post request
  jsonBody := make(map[string]interface{})
  jsonBody["email"] = username
  jsonBody["password"] = password
  jsonBody["otp"] = otp

  // configure the request
  loginResponse := bitgo.makeHttpRequest("POST", loginUrl, jsonBody)

  // print(loginResponse)
  accessToken := loginResponse.Get("access_token").Str
  if len(accessToken) > 0 {
    bitgo.accessToken = accessToken
  }

  return loginResponse
}

func (bitgo BitGo) Get(url string) gjson.Result {
  return bitgo.makeHttpRequest("GET", url, nil)
}

func (bitgo BitGo) Post(url string, data map[string]interface{}) gjson.Result {
  return bitgo.makeHttpRequest("POST", url, data)
}

func (bitgo BitGo) Url(extension string) string {
  return bitgo.baseUrl + "/api/v1" + extension
}

func (bitgo BitGo) UrlV2(extension string) string {
  return bitgo.baseUrl + "/api/v2" + extension
}

func (bitgo BitGo) Coin(coin string) *Coin {
  coinObject := new(Coin)
  coinObject.bitgo = bitgo
  coinObject.coin = coin
  return coinObject
}


func (bitgo BitGo) makeHttpRequest(method string, url string, jsonBody map[string]interface{}) gjson.Result {
  jsonBuffer, err := json.Marshal(jsonBody)

  req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonBuffer))

  if jsonBody != nil {
    req.Header.Set("Content-Type", "application/json")
  }else{
    req, err = http.NewRequest(method, url, nil)
  }

  if len(bitgo.accessToken) > 0 {
    req.Header.Set("Authorization", "Bearer " + bitgo.accessToken)
  }

  // initiate the request
  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    panic(err)
  }
  // await the response
  defer resp.Body.Close()
  // read the response
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    fmt.Println(err)
    panic(err)
  }

  return gjson.ParseBytes(body)
}
