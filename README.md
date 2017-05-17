# GoLang_Test_NaCL
Testing NaCL Library, using box (asymmetric) and secretbox (symmetric) encryptions

The Test will do the following

1- Generate Private and Public Keys for client and server
2- Generate Nounce that will be used in the both tests, box and secret box
3- Generate Random 32 Bytes AES Key
4- Send the AES Key using Public Key cryoto (box)
5- Client receieve and decrypt, Aquiring the AES Key
6- Client use this key to send the message to the server
7- Server Decrypt the message


