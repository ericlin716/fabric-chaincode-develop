export OS_ARCH=$(echo "$(uname -s | tr '[:upper:]' '[:lower:]' | sed 's/mingw64_nt.*/windows/')-$(uname -m | sed 's/x86_64/amd64/g')" | awk '{print tolower($0)}')
export PATH=${PWD}/../bin/${OS_ARCH}:$PATH
export FABRIC_CFG_PATH=$PWD/../config/

export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID="MSP15"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/crypto/crypto-config-dev/peerOrganizations/dev.org1-car-insurance.cathaybc.com/peers/peer0.dev.org1-car-insurance.cathaybc.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/crypto/crypto-config-dev/peerOrganizations/dev.org1-car-insurance.cathaybc.com/users/Admin@dev.org1-car-insurance.cathaybc.com/msp
export CORE_PEER_ADDRESS=localhost:7051

while [[ $# -ge 1 ]] ; do
  key="$1"
  case $key in
  1 ) # CompulsoryCaseExists
    peer chaincode query -C mychannel -n car-insurance -c '{"function":"CompulsoryCaseExists","Args":["A123456789", "zxcv109-10301334"]}'
    shift
    ;;
  2 ) # GetIDHistory
    peer chaincode query -C mychannel -n car-insurance -c '{"function":"GetIDHistory","Args":["A123456789", "109-10"]}'
    shift
    ;;
  3 ) # Invoke
    peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.dev.car-insurance.cathaybc.com --tls --cafile ${PWD}/crypto/crypto-config-dev/ordererOrganizations/dev.car-insurance.cathaybc.com/orderers/orderer.dev.car-insurance.cathaybc.com/msp/tlscacerts/tlsca.dev.car-insurance.cathaybc.com-cert.pem -C mychannel -n car-insurance --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/crypto/crypto-config-dev/peerOrganizations/dev.org1-car-insurance.cathaybc.com/peers/peer0.dev.org1-car-insurance.cathaybc.com/tls/ca.crt --peerAddresses localhost:8051 --tlsRootCertFiles ${PWD}/crypto/crypto-config-dev/peerOrganizations/dev.org2-car-insurance.cathaybc.com/peers/peer0.dev.org2-car-insurance.cathaybc.com/tls/ca.crt -c '{"function":"Function1","Args":["{\"key\":\"key\",\"value\":1,\"message\":\"test complex object\"}"]}'
    shift
    ;;
  4 ) # UpdateCompulsoryCase
    peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.dev.car-insurance.cathaybc.com --tls --cafile ${PWD}/crypto/crypto-config-dev/ordererOrganizations/dev.car-insurance.cathaybc.com/orderers/orderer.dev.car-insurance.cathaybc.com/msp/tlscacerts/tlsca.dev.car-insurance.cathaybc.com-cert.pem -C mychannel -n car-insurance --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/crypto/crypto-config-dev/peerOrganizations/dev.org1-car-insurance.cathaybc.com/peers/peer0.dev.org1-car-insurance.cathaybc.com/tls/ca.crt --peerAddresses localhost:8051 --tlsRootCertFiles ${PWD}/crypto/crypto-config-dev/peerOrganizations/dev.org2-car-insurance.cathaybc.com/peers/peer0.dev.org2-car-insurance.cathaybc.com/tls/ca.crt -c '{"function":"Invoke","Args":["{\"key\":\"key\",\"value\":\"case_key1\",\"apportion_key\":\"apportion_key1\",\"case_number\":\"case_number1\",\"in_car_number\":\"in_car_number1\",\"hit_time\":1,\"insurance_number\":\"insurance_number1\",\"insurance_car_number\":\"insurance_car_number1\",\"applicant_name\":\"applicant_name1\",\"applicant_id_number\":\"applicant_id_number1\",\"applicant_id_number_type\":\"applicant_id_number_type1\",\"applicant_birthday\":\"applicant_birthday1\",\"applicant_type\":\"applicant_type1\",\"amount\":1,\"person_in_charge\":\"person_in_charge1\",\"phone\":\"phone1\",\"extension\":\"extension\",\"file_token\":\"file_token1\",\"created_at\":1,\"charges\":[]}"]}'
    shift
    ;;
  5 ) # ApportionNotified
    peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.dev.car-insurance.cathaybc.com --tls --cafile ${PWD}/crypto/crypto-config-dev/ordererOrganizations/dev.car-insurance.cathaybc.com/orderers/orderer.dev.car-insurance.cathaybc.com/msp/tlscacerts/tlsca.dev.car-insurance.cathaybc.com-cert.pem -C mychannel -n car-insurance --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/crypto/crypto-config-dev/peerOrganizations/dev.org1-car-insurance.cathaybc.com/peers/peer0.dev.org1-car-insurance.cathaybc.com/tls/ca.crt --peerAddresses localhost:9051 --tlsRootCertFiles ${PWD}/crypto/crypto-config-dev/peerOrganizations/dev.org2-car-insurance.cathaybc.com/peers/peer0.dev.org2-car-insurance.cathaybc.com/tls/ca.crt -c '{"function":"ApportionNotified","Args":["zxcv109-10301334", "109-10301341", "NOTIFIEDTIME"]}'
    shift
    ;;
  6 ) # WriteNote
    peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.dev.car-insurance.cathaybc.com --tls --cafile ${PWD}/crypto/crypto-config-dev/ordererOrganizations/dev.car-insurance.cathaybc.com/orderers/orderer.dev.car-insurance.cathaybc.com/msp/tlscacerts/tlsca.dev.car-insurance.cathaybc.com-cert.pem -C mychannel -n car-insurance --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/crypto/crypto-config-dev/peerOrganizations/dev.org1-car-insurance.cathaybc.com/peers/peer0.dev.org1-car-insurance.cathaybc.com/tls/ca.crt --peerAddresses localhost:9051 --tlsRootCertFiles ${PWD}/crypto/crypto-config-dev/peerOrganizations/dev.org2-car-insurance.cathaybc.com/peers/peer0.dev.org2-car-insurance.cathaybc.com/tls/ca.crt -c '{"function":"WriteNote","Args":["zxcv109-10301334", "109-10301341", "Hello World", "109-11031031"]}'
    shift
    ;;
  7 ) # GetNote
    peer chaincode query -C mychannel -n car-insurance -c '{"function":"GetNote","Args":["zxcv109-10301334", "109-10301341"]}'
    shift
    ;;
  8 ) # ApportionCancel
    peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.dev.car-insurance.cathaybc.com --tls --cafile ${PWD}/crypto/crypto-config-dev/ordererOrganizations/dev.car-insurance.cathaybc.com/orderers/orderer.dev.car-insurance.cathaybc.com/msp/tlscacerts/tlsca.dev.car-insurance.cathaybc.com-cert.pem -C mychannel -n car-insurance --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/crypto/crypto-config-dev/peerOrganizations/dev.org1-car-insurance.cathaybc.com/peers/peer0.dev.org1-car-insurance.cathaybc.com/tls/ca.crt --peerAddresses localhost:9051 --tlsRootCertFiles ${PWD}/crypto/crypto-config-dev/peerOrganizations/dev.org2-car-insurance.cathaybc.com/peers/peer0.dev.org2-car-insurance.cathaybc.com/tls/ca.crt -c '{"function":"ApportionCancel","Args":["zxcv109-10301334", "109-10301341", "CANCELTIME"]}'
    shift
    ;;
  9 ) # ApportionCompleted
    peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.dev.car-insurance.cathaybc.com --tls --cafile ${PWD}/crypto/crypto-config-dev/ordererOrganizations/dev.car-insurance.cathaybc.com/orderers/orderer.dev.car-insurance.cathaybc.com/msp/tlscacerts/tlsca.dev.car-insurance.cathaybc.com-cert.pem -C mychannel -n car-insurance --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/crypto/crypto-config-dev/peerOrganizations/dev.org1-car-insurance.cathaybc.com/peers/peer0.dev.org1-car-insurance.cathaybc.com/tls/ca.crt --peerAddresses localhost:9051 --tlsRootCertFiles ${PWD}/crypto/crypto-config-dev/peerOrganizations/dev.org2-car-insurance.cathaybc.com/peers/peer0.dev.org2-car-insurance.cathaybc.com/tls/ca.crt -c '{"function":"ApportionCompleted","Args":["zxcv109-10301334", "109-10301341", "109-12", "COMPLETETIME"]}'
    shift
    ;;
  10 ) # CreateApportion
    peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.dev.car-insurance.cathaybc.com --tls --cafile ${PWD}/crypto/crypto-config-dev/ordererOrganizations/dev.car-insurance.cathaybc.com/orderers/orderer.dev.car-insurance.cathaybc.com/msp/tlscacerts/tlsca.dev.car-insurance.cathaybc.com-cert.pem -C mychannel -n car-insurance --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/crypto/crypto-config-dev/peerOrganizations/dev.org1-car-insurance.cathaybc.com/peers/peer0.dev.org1-car-insurance.cathaybc.com/tls/ca.crt --peerAddresses localhost:9051 --tlsRootCertFiles ${PWD}/crypto/crypto-config-dev/peerOrganizations/dev.org2-car-insurance.cathaybc.com/peers/peer0.dev.org2-car-insurance.cathaybc.com/tls/ca.crt -c '{"function":"CreateApportion","Args":["asdf109-10291754", "109-10291754", "[{\"MSPID\":\"MSP13\",\"InsuranceNumber\":\"AXVBF67485930\",\"CarNumber\":\"AB-0987\",\"Amount\":7700,\"Status\":\"CREATED\"}]", "20000"]}'
    shift
    ;;
  11 ) # AppendKeyToID
    peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.dev.car-insurance.cathaybc.com --tls --cafile ${PWD}/crypto/crypto-config-dev/ordererOrganizations/dev.car-insurance.cathaybc.com/orderers/orderer.dev.car-insurance.cathaybc.com/msp/tlscacerts/tlsca.dev.car-insurance.cathaybc.com-cert.pem -C mychannel -n car-insurance --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/crypto/crypto-config-dev/peerOrganizations/dev.org1-car-insurance.cathaybc.com/peers/peer0.dev.org1-car-insurance.cathaybc.com/tls/ca.crt --peerAddresses localhost:9051 --tlsRootCertFiles ${PWD}/crypto/crypto-config-dev/peerOrganizations/dev.org2-car-insurance.cathaybc.com/peers/peer0.dev.org2-car-insurance.cathaybc.com/tls/ca.crt -c '{"function":"AppendKeyToID","Args":["A123456789", "zxcv109-10301334"]}'
    shift
    ;;
  12 ) # DeleteCompulsoryCase
    peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.dev.car-insurance.cathaybc.com --tls --cafile ${PWD}/crypto/crypto-config-dev/ordererOrganizations/dev.car-insurance.cathaybc.com/orderers/orderer.dev.car-insurance.cathaybc.com/msp/tlscacerts/tlsca.dev.car-insurance.cathaybc.com-cert.pem -C mychannel -n car-insurance --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/crypto/crypto-config-dev/peerOrganizations/dev.org1-car-insurance.cathaybc.com/peers/peer0.dev.org1-car-insurance.cathaybc.com/tls/ca.crt --peerAddresses localhost:9051 --tlsRootCertFiles ${PWD}/crypto/crypto-config-dev/peerOrganizations/dev.org2-car-insurance.cathaybc.com/peers/peer0.dev.org2-car-insurance.cathaybc.com/tls/ca.crt -c '{"function":"DeleteCompulsoryCase","Args":["{\"case_key\":\"case_key1\"}"]}'
    shift
    ;;
  13 ) # GetBalance
    peer chaincode query -C mychannel -n car-insurance -c '{"function":"GetBalance","Args":["109-11"]}'
    shift
    ;;
  14 ) # GetBalance
    peer chaincode query -C mychannel -n car-insurance -c '{"function":"GetBalance","Args":["109-12"]}'
    shift
    ;;
  15 ) # QueryReceivedAndPaid
    peer chaincode query -C mychannel -n car-insurance -c '{"function":"QueryReceivedAndPaid","Args":["109-12"]}'
    shift
    ;;
  * )
    echo
    echo "Unknown flag: $key"
    exit 1
    ;;
  esac
  shift
done


