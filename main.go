package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
	"strings"
	"errors"
)

func main() {
	fmt.Println("Accepting requests on localhost:8080...")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {

	/* Read json http body into struct */
	clientMsg, _ := ioutil.ReadAll(r.Body)
	var jsonMsg struct {
		ArmoredMessage string
		ExpectedSignatureKeyId string
	}
	json.Unmarshal(clientMsg, &jsonMsg)

	/* Validate signature and extract armored message */
	signatureIsValid, contents, err := GpgSignatureIsValid(jsonMsg.ArmoredMessage, jsonMsg.ExpectedSignatureKeyId)

	/* Prepare response */
	var response map[string]interface{}
	if err != nil {
		response = map[string]interface{} {
			"error": err.Error(),
		}
	} else {
		response = map[string]interface{} {
			"messageContents": contents,
			"signatureIsValid": signatureIsValid,
		}
	}

	/* Send response back to client */
	converted, err := json.Marshal(response)
	fmt.Fprintf(w, string(converted))
}

func GpgSignatureIsValid(armoredMessage string, expectedKeyId string) (bool, string, error) {
	/* Decoded amored message */
	decoded, err := armor.Decode(strings.NewReader(armoredMessage))
	if err != nil {
		return false, "", errors.New("Error decoding armored message")
	}

	/* Read message details from decoded armored message */
	msg, err := openpgp.ReadMessage(decoded.Body, openpgp.EntityList{}, nil, nil)
	if err != nil {
		return false, "", errors.New("Error reading message details from decoded armored message")
	}

	/* Check if signature matches expected key */
	strKey := fmt.Sprintf("%X", msg.SignedByKeyId)
	keyLen := len(strKey)
	contents, _ := ioutil.ReadAll(msg.UnverifiedBody)
	if keyLen > 0 && len(expectedKeyId) >= keyLen && expectedKeyId[len(expectedKeyId)-keyLen:] == strKey {
		return true, string(contents), nil
	}

	/* Signature does not match */
	return false, string(contents), nil
}
