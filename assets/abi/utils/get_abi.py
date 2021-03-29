#!/usr/bin/python
import sys
import json

import requests

ABI_ENDPOINT = 'https://api.etherscan.io/api?module=contract&action=getabi&address='

def get_abi(address):
    response_json = requests.get(ABI_ENDPOINT + address).json()
    return json.loads(response_json['result'])


if __name__ == "__main__":
    address, save_to = sys.argv[1:]

    with open(save_to, 'w') as f:
        json.dump(get_abi(address), f)
