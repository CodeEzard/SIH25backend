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


try:
    with open(f"{path}/contracts/build/mint.bin") as f:
        bytecode = f.read()
        print("Successfully loaded bytecode: ")
        print("---------------------------------")
except Exception as e:
    text = f"Some error occured when reading from bin file: {e}"
    print(text)
    print("-" * len(text))

try:
    contract = w3.eth.contract(abi=abi, bytecode=bytecode)
except Exception as e:
    text = f"Some error occurred when connecting to contract: {e}"
    print(text)
    print("-" * len(text))

nonce = w3.eth.get_transaction_count(account.address)
print(f"Nonce for {account.address}: {nonce}")


tx = contract.constructor().build_transaction({
    "from": account.address,
    "nonce": nonce,
    "gas": 3000000,
    "gasPrice": w3.eth.gas_price,
    "chainId": 11155111,
 })



signed_tx = w3.eth.account.sign_transaction(tx, pvt_key)
tx_hash = w3.eth.send_raw_transaction(signed_tx.raw_transaction)
print("Transaction sent: ", tx_hash.hex())
print("---------------------------------")

receipt = w3.eth.wait_for_transaction_receipt(tx_hash)
print("Contract deployed at:", receipt.contractAddress)
print("---------------------------------")
