/////////////////////////////////////////////
//    THE BLOCKCHAIN PKI EXPERIMENT     ////
///////////////////////////////////////////
/*
	This is the fabpki, a chaincode that implements a Public Key Infrastructure (PKI)
	for measuring instruments. It runs in Hyperledger Fabric 1.4.
	He was created as part of the PKI Experiment. You can invoke its methods
	to store measuring instruments public keys in the ledger, and also to verify
	digital signatures that are supposed to come from these instruments.

	@author: Wilson S. Melo Jr.
	@date: Oct/2019
*/
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

type SmartContract struct {
}

type Veiculos struct {
	Veiculos []Veiculum `json:"Veiculo"`
}

type Veiculum struct {
	Categoria string `json:"Categoria"`
	Marca     string `json:"Marca"`
	Versao    string `json:"Versao"`
	Modelo    string `json:"Modelo"`
	Emissao   int    `json:"Emissao"`
	Codigo    string `json:"Codigo"`
}

func (s *SmartContract) Init(stub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(stub shim.ChaincodeStubInterface) sc.Response {
	fn, args := stub.GetFunctionAndParameters()

	if fn == "registrarBanco" {
		return s.registrarBanco(stub, args)

	} else if fn == "registrarUsuario" {
		return s.registrarUsuario(stub, args)

	} else if fn == "sleepTest" {
		return s.sleepTest(stub, args)

	} else if fn == "countHistory" {
		return s.countHistory(stub, args)

	} else if fn == "countLedger" {
		return s.countLedger(stub)

	} else if fn == "queryLedger" {
		return s.queryLedger(stub, args)
	}

	return shim.Error("Chaincode não suporta essa função.")
}

func (s *SmartContract) registrarBanco(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Não foi encontrado nenhum argumento. Tente novamente!")
	}

	banco := args[0]

	var bancoInformacoes Veiculos

	bancoAsBytes, _ := json.Marshal(bancoInformacoes)

	stub.PutState(banco, bancoAsBytes)

	fmt.Println("Registrando seu banco de veiculos...")

	//notify procedure success
	return shim.Success(nil)
}

func (s *SmartContract) registrarUsuario(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 3 {
		return shim.Error("Eram esperados 3 argumentos... Tente novamente!")
	}
	banco := args[0]
	dadosUsuario := args[1]

	bancoAsBytes, err := stub.GetState(banco)

	if err != nil || bancoAsBytes == nil {
		return shim.Error("Erro na validação dos dados de veículos!")
	}

	bancoUsr := Veiculum{}

	fmt.Println("Comparando com banco original...")

	json.Unmarshal(bancoAsBytes, &bancoUsr)

	//Continuar com a captura de dados do usuario com base no arquivo py
}

func (s *SmartContract) sleepTest(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	//validate args vector lenght
	if len(args) != 1 {
		return shim.Error("It was expected 1 parameter: <sleeptime>")
	}

	//gets the parameter associated with the meter ID and the incremental measurement
	sleeptime, err := strconv.Atoi(args[0])

	//test if we receive a valid meter ID
	if err != nil {
		return shim.Error("Error on retrieving sleep time")
	}

	//tests if sleeptime is a valid value
	if sleeptime > 0 {
		//stops during sleeptime seconds
		time.Sleep(time.Duration(sleeptime) * time.Second)
	}

	//return payload with bytes related to the meter state
	return shim.Success(nil)
}

/*
   This method brings the changing history of a specific meter asset. It can be useful to
   query all the changes that happened with a meter value.
   - args[0] - asset key (or meter ID)
*/
func (s *SmartContract) queryHistory(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	//validate args vector lenght
	if len(args) != 1 {
		return shim.Error("It was expected 1 parameter: <key>")
	}

	historyIer, err := stub.GetHistoryForKey(args[0])

	//verifies if the history exists
	if err != nil {
		//fmt.Println(errMsg)
		return shim.Error("Fail on getting ledger history")
	}

	// buffer is a JSON array containing records
	var buffer bytes.Buffer
	var counter = 0
	buffer.WriteString("[")
	bArrayMemberAlreadyWritten := false
	for historyIer.HasNext() {
		//increments iterator
		queryResponse, err := historyIer.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}

		//generates a formated result
		buffer.WriteString("{\"Value\":")
		buffer.WriteString("\"")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("\"")
		buffer.WriteString(", \"Counter\":")
		buffer.WriteString(strconv.Itoa(counter))
		//buffer.WriteString(queryResponse.Timestamp)
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true

		//increases counter
		counter++
	}
	buffer.WriteString("]")
	historyIer.Close()

	//loging...
	fmt.Printf("Consulting ledger history, found %d\n records", counter)

	//notify procedure success
	return shim.Success(buffer.Bytes())
}

/*
   This method brings the number of times that a meter asset was modified in the ledger.
   It performs faster than queryHistory() method once it does not retrive any information,
   it only counts the changes.
   - args[0] - asset key (or meter ID)
*/
func (s *SmartContract) countHistory(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	//validate args vector lenght
	if len(args) != 1 {
		return shim.Error("It was expected 1 parameter: <key>")
	}

	historyIer, err := stub.GetHistoryForKey(args[0])

	//verifies if the history exists
	if err != nil {
		//fmt.Println(errMsg)
		return shim.Error("Fail on getting ledger history")
	}

	//creates a counter
	var counter int64
	counter = 0

	for historyIer.HasNext() {
		//increments iterator
		_, err := historyIer.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		//increases counter
		counter++

		fmt.Printf("Consulting ledger history, found %d\n records", counter)
	}
	// buffer is a JSON array containing records
	var buffer bytes.Buffer
	buffer.WriteString("[")
	buffer.WriteString("\"Counter\":")
	buffer.WriteString(strconv.FormatInt(counter, 10))
	buffer.WriteString("]")

	historyIer.Close()

	//loging...
	fmt.Printf("Consulting ledger history, found %d\n records", counter)

	//notify procedure success
	return shim.Success(buffer.Bytes())
}

/*
   This method counts the total of well succeeded transactions in the ledger.
*/
func (s *SmartContract) countLedger(stub shim.ChaincodeStubInterface) sc.Response {

	//use a range of keys, assuming that the max key value is 999999,
	resultsIterator, err := stub.GetStateByRange("0", "999999")
	if err != nil {
		return shim.Error(err.Error())
	}

	//defer iterator closes at the end of the function
	defer resultsIterator.Close()

	//creates a counter
	var counter int64
	var keys int64
	counter = 0
	keys = 0

	//the interator checks all the valid keys
	for resultsIterator.HasNext() {

		//increments iterator
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		//busca historico da proxima key
		historyIer, err := stub.GetHistoryForKey(queryResponse.Key)

		//verifies if the history exists
		if err != nil {
			//fmt.Println(errMsg)
			return shim.Error(err.Error())
		}

		defer historyIer.Close()

		for historyIer.HasNext() {
			//increments iterator
			_, err := historyIer.Next()
			if err != nil {
				return shim.Error(err.Error())
			}

			//increases counter
			counter++
		}
		fmt.Printf("Consulting ledger history, found key %s\n", queryResponse.Key)

		keys++
	}
	// buffer is a JSON array containing records
	var buffer bytes.Buffer
	buffer.WriteString("[")
	buffer.WriteString("\"Counter\":")
	buffer.WriteString(strconv.FormatInt(counter, 10))
	buffer.WriteString("\"Keys\":")
	buffer.WriteString(strconv.FormatInt(keys, 10))
	buffer.WriteString("]")

	//loging...
	fmt.Printf("Consulting ledger history, found %d transactions in %d keys\n", counter, keys)

	//notify procedure success
	return shim.Success(buffer.Bytes())
}

/*
   This method executes a free query on the ledger, returning a vector of meter assets.
   The query string must be a query expression supported by CouchDB servers.
   - args[0] - query string.
*/
func (s *SmartContract) queryLedger(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	//validate args vector lenght
	if len(args) != 1 {
		return shim.Error("It was expected 1 parameter: <query string>")
	}

	//using auxiliar variable
	queryString := args[0]

	//loging...
	fmt.Printf("Executing the following query: %s\n", queryString)

	//try to execute query and obtain records iterator
	resultsIterator, err := stub.GetQueryResult(queryString)
	//test if iterator is valid
	if err != nil {
		return shim.Error(err.Error())
	}
	//defer iterator closes at the end of the function
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer
	buffer.WriteString("[")
	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		//increments iterator
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}

		//generates a formated result
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")
		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	//loging...
	fmt.Printf("Obtained the following fill up records: %s\n", buffer.String())

	//notify procedure success
	return shim.Success(buffer.Bytes())
}

/*
 * The main function starts up the chaincode in the container during instantiate
 */
func main() {

	////////////////////////////////////////////////////////
	// USE THIS BLOCK TO COMPILE THE CHAINCODE
	if err := shim.Start(new(SmartContract)); err != nil {
		fmt.Printf("Erro ao compilar Smart Contract: %s\n", err)
	}
	////////////////////////////////////////////////////////

	////////////////////////////////////////////////////////
	// USE THIS BLOCK TO PERFORM ANY TEST WITH THE CHAINCODE

	// //create pair of keys
	// privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	// if err != nil {
	// 	panic(err)
	// }

	// //marshal the keys in a buffer
	// e, err := json.Marshal(privateKey)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// _ = ioutil.WriteFile("ecdsa-keys.json", e, 0644)

	// //read the saved key
	// file, _ := ioutil.ReadFile("ecdsa-keys.json")

	// myPrivKey := ecdsa.PrivateKey{}
	// //myPubKey := ecdsa.PublicKey{}

	// _ = json.Unmarshal([]byte(file), &myPrivKey)

	// fmt.Println("Essa é minha chave privada:")
	// fmt.Println(myPrivKey)

	// myPubKey := myPrivKey.PublicKey

	// //test digital signature verifying
	// msg := "message"
	// hash := sha256.Sum256([]byte(msg))
	// fmt.Println("hash: ", hash)

	// r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("signature: (0x%x, 0x%x)\n", r, s)

	// myPubKey.Curve = elliptic.P256()

	// fmt.Println("Essa é minha chave publica:")
	// fmt.Println(myPubKey)

	// valid := ecdsa.Verify(&myPubKey, hash[:], r, s)
	// fmt.Println("signature verified:", valid)

	// otherpk := "-----BEGIN PUBLIC KEY-----\nMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE6NXETwtkAKGWBcIsI6/OYE0EwsVj\n3Fc4hHTaReNfq6Hz2UEzsJKCYN0stjPCXbpdUlYtETC1a3EcS3SUVYX6qA==\n-----END PUBLIC KEY-----\n"

	// newkey := PublicKeyDecodePEM(otherpk)
	// myPubKey.Curve = elliptic.P256()

	// //valid = ecdsa.Verify(newkey, hash[:], r, s)
	// //fmt.Println("signature verified:", valid)

	// mysign := "MEYCIQCY16jbdY222oEpFiSRwXPi1kS7c4wuwxYXeWJOoAjnVgIhAJQTM+itbm1mQyd40Ug0xr2/AvjZmFSdoc/iSSHA6nRI"

	// // first decode the signature to extract the DER-encoded byte string
	// der, err := base64.StdEncoding.DecodeString(mysign)
	// if err != nil {
	// 	panic(err)
	// }

	// // unmarshal the R and S components of the ASN.1-encoded signature into our
	// // signature data structure
	// sig := &ECDSASignature{}
	// _, err = asn1.Unmarshal(der, sig)
	// if err != nil {
	// 	panic(err)
	// }

	// valid = ecdsa.Verify(&newkey, hash[:], sig.R, sig.S)
	// fmt.Println("signature verified:", valid)

	// fmt.Println("Curve: ", newkey.Curve.Params())

	////////////////////////////////////////////////////////

}
