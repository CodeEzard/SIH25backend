from web3 import Web3
import json
import os
import pathlib

path = pathlib.Path.cwd() # get the current fucking path

w3 = Web3(Web3.HTTPProvider("https://sepolia.infura.io/v3/2a159ca7304a4df4ade0d0ccf9ce70ef"))

if not w3.is_connected():
    print("Failed to conncet to sepolia.")
else:
    print("Connected to sepolia successfully.")
    print("---------------------------------")
    
pvt_key = os.getenv("pvt_key")

try:
    account = w3.eth.account.from_key(pvt_key)
except:
    print("Some error occurred while connecting to account.")
    print("---------------------------------")



try:
    with open(f"{path}/contracts/build/mint.abi", "r") as f:
        abi = json.load(f)
        text = f"Successfully loaded abi."
        print(text)
        print("-" * len(text))
        
except Exception as e:
    text = f"Some error occured when reading from abi file: {e}"
    print(text)
    print("-" * len(text))

contract_address = "0x6DA6E82143802EBAC3100E05aB4cd1b8A96554D2"

try:
    contract = w3.eth.contract(address=contract_address,abi=abi)
    print("Contract: ", contract)
except Exception as e:
    print("Some error occurred when connecting to contract: ", e)
    print("---------------------------------")


print("Token ID: ", contract.functions.tokenURI(0).call())

# nonce = w3.eth.get_transaction_count(account.address)
# print(f"Nonce for {account.address}: {nonce}")

# tx = contract.functions.mintDoc("0x8d0BB74e37ab644964AcA2f3Fbe12b9147f9d841", "https://chocolate-electoral-parrot-267.mypinata.cloud/ipfs/bafkreige7clmdi67kbrtyouk3ilcxrkjzcjb65s5ccplllek4helrmmw3u").build_transaction({
#     "from": account.address,
#     "nonce": nonce,
#     "gas": 3000000,
#     "gasPrice": w3.eth.gas_price,
# })

# try:
#     signed_txn = w3.eth.account.sign_transaction(tx, private_key=pvt_key)
# except Exception as e:
#     print("Error while signing txn: ", e)

# try:
#     tx_hash = w3.eth.send_raw_transaction(signed_txn.raw_transaction)
#     print("Transaction sent: ", tx_hash.hex())
# except Exception as e:
#     print("Error when sending raw txn: ", e)

# try:
#     receipt = w3.eth.wait_for_transaction_receipt(tx_hash)
#     print("Receipt:", receipt.transactionHash)
# except Exception as e:
#     print("Error while receiving receipt: ", e)



