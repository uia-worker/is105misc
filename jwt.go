package main
import (
  "crypto/hmac"
  "crypto/sha256"
  "fmt"
)
// JSON Web Tokens
// consist of three parts separated by dots ( . )
// Header.Payload.Signature
// xxxxx.yyyyy.zzzzz
// -------
// Header consists of two parts
// the type of token (JWT) and hashing algorithm in use (HMAC SHA256 RSA)
// {
//   "alg": "HS256",
//   "typ": "JWT"
// }
// JSON is Base64Url encoded to form the first part of JWT
// ------
// Payload contains the claims
// Claims are statements about the entity, usually the user
// They also include some additional metadata
// There are three types of claims (reserved, public and private)
// Reserved claims: not mandatory but recommended
//     iss (issuer), exp (expiration time), sub (subject), aud (audience ++)
// Public claims: can be defined at will by those using JWT
//     Should be defined in IANA JSON Web Token Registry of as URI to avoid
//     collisions
// Private claims: the custom claims created to share information between
//     parties that agree on using them
// {
//   "sub": "1234567890"
//   "name": "John Doe"
//   "admin": true
// }
// The payload is then Base64Url encoded
// -------
// Signature, - take the encoded header, the encoded payload, a secret,
// the algorithm specified in header, and sign that.
// HMACSHA256(
//    base64UrlEncode(header) + "." +
//    base64UrlEncode(payload), secret)
// In this way the sender of the JWT can be verified and integrity
// (message was not changed along the way) checked
// ----
// User agent sends the JWT, typically in the Authorization header
// using Bearer schema.
// Authorization: Bearer <token>
func main() {
  // Make a hash with SHA256
  h := sha256.New()
  mac := hmac.New(sha256.New, []byte("secret123"))
  h.Write([]byte("A secret message"))
  fmt.Printf("%x\n", h.Sum(nil))


  mac.Write([]byte("A secret message"))
  messageMAC := mac.Sum(nil)
  fmt.Printf("%x\n", messageMAC)

  fmt.Println(CheckMAC([]byte("A secret message"), messageMAC, []byte("secret122")))

}

// CheckMAC reports whether messageMAC is a valid HMAC tag for message.
func CheckMAC(message, messageMAC, key []byte) bool {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	expectedMAC := mac.Sum(nil)
	return hmac.Equal(messageMAC, expectedMAC)
}
