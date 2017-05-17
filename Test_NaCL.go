package main
import (
  "crypto/rand"
  "fmt"
  "strings"
  "log"
  "encoding/hex"
  "golang.org/x/crypto/nacl/box"
  "golang.org/x/crypto/nacl/secretbox"
)
func main(){
	AESKey_32Bytes := make([]byte, 32)
	_, err := rand.Read(AESKey_32Bytes)
	if err != nil {
		log.Fatal(err)
	}
  var NoOunce [24]byte
  rand.Read(NoOunce[:])
  fmt.Println("Nounce  " + strings.ToUpper(hex.EncodeToString(NoOunce[:])))
  // This part is the Box (Asymetric Key Encryption)

	PublicKeyServer, PrivateKeyServer, _ := box.GenerateKey(rand.Reader)  //check for errors if you want
  fmt.Println("Crypto_Shared_AESKey To be Sent\n" + strings.ToUpper(hex.EncodeToString(AESKey_32Bytes)))
  fmt.Println("Crypto_Server_PublicKey " + strings.ToUpper(hex.EncodeToString(PublicKeyServer[:])))
  fmt.Println("Crypto_Server_PrivateKey " + strings.ToUpper(hex.EncodeToString(PrivateKeyServer[:])))

  PublicKeyClient, PrivateKeyClient, _ := box.GenerateKey(rand.Reader)  //check for errors if you want
  fmt.Println("Crypto_Client_PublicKey " + strings.ToUpper(hex.EncodeToString(PublicKeyClient[:])))
  fmt.Println("Crypto_Client_PrivateKey " + strings.ToUpper(hex.EncodeToString(PrivateKeyClient[:])))

  fmt.Println("Server Encrypting The AES Key")
  encryptedForYou := box.Seal(nil, AESKey_32Bytes, &NoOunce, PublicKeyClient, PrivateKeyServer)  //check for errors if you want
  println(strings.ToUpper(hex.EncodeToString(encryptedForYou)))

  fmt.Println("Client will now Decrypt the Message Received")
  decryptedForYou, _ := box.Open(nil, encryptedForYou, &NoOunce, PublicKeyServer, PrivateKeyClient)  //check for errors if you want
  fmt.Println("The AES Key Decrypted in Message is:")
  fmt.Println(strings.ToUpper(hex.EncodeToString(decryptedForYou)))

// The bottom part is the Secret Box (Symetric Key Encryption)

  fmt.Println("Client will now Encrypt the following text using the received AES Key")
  fmt.Println("The Test Encrypted Message is:")
  Message := "Like any other social media site Facebook has length requirements when it comes to writing on the wall, providing status, messaging and commenting. Understanding how many characters you can use, enables you to more effectively use Facebook as a business or campaign tool"
  fmt.Println(Message)
  var convertDecrypt [32]byte
  copy(convertDecrypt[:],decryptedForYou[:])
  out := secretbox.Seal(nil, []byte(Message), &NoOunce, &convertDecrypt)
  fmt.Println("Encrypted Message:")
  fmt.Println(strings.ToUpper(hex.EncodeToString(out)))

  fmt.Println("Server will try to decrypt the Message:")
  message, _ := secretbox.Open(nil, out, &NoOunce, &convertDecrypt) //check for errors if you want
  fmt.Println(string(message))


}
