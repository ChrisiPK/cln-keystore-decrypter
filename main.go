package main

import (
	"fmt"
	"os"
)

func (kr *keyRing) printPubKeys() {
	for _, CSAKey := range kr.CSA {
		fmt.Printf("Unlocked CSA key! Public key: %s, Private key: 0x%x\n", CSAKey.ID(), []byte(CSAKey.Raw()))
	}
	for _, ETHKey := range kr.Eth {
		fmt.Printf("Unlocked ETH key! Address: %s, Private key: 0x%064x\n", ETHKey.ID(), []byte(ETHKey.Raw()))
	}
	for _, OCRKey := range kr.OCR {
		fmt.Printf("Unlocked OCR key! Public key: %s, Private key: 0x%x\n", OCRKey.ID(), []byte(OCRKey.Raw()))
	}
	for _, OCR2Key := range kr.OCR2 {
		fmt.Printf("Unlocked OCR2 key! Public key: %s, Private key: 0x%x\n", OCR2Key.ID(), []byte(OCR2Key.Raw()))
	}
	for _, P2PKey := range kr.P2P {
		fmt.Printf("Unlocked P2P key! Public key: %s, Private key: 0x%x\n", P2PKey.ID(), []byte(P2PKey.Raw()))
	}
	for _, solanaKey := range kr.Solana {
		fmt.Printf("Unlocked Solana key! Address: %s, Private key: 0x%x\n", solanaKey.ID(), []byte(solanaKey.Raw()))
	}
	for _, terraKey := range kr.Terra {
		fmt.Printf("Unlocked Terra key! Address: %s, Private key: 0x%x\n", terraKey.ID(), []byte(terraKey.Raw()))
	}
	for _, VRFKey := range kr.VRF {
		fmt.Printf("Unlocked VRF key! Public key: %s, Private key: 0x%x\n", VRFKey.ID(), []byte(VRFKey.Raw()))
	}
}

func main() {
	keyfile, err := os.ReadFile("keyring")
	if err != nil {
		panic(err)
	}
	passfile, err := os.ReadFile("password")
	if err != nil {
		panic(err)
	}
	ekr := encryptedKeyRing{
		EncryptedKeys: keyfile,
	}
	kr, err := ekr.Decrypt(string(passfile))
	if err != nil {
		panic(err)
	}
	kr.printPubKeys()
}
