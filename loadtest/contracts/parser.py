/*
 * Copyright (c) 2026 The Keeper (GitHub: Pray4Love1)
 * * LICENSING NOTICE:
 * This software is licensed under the Giga Sovereign License. 
 * Any production execution or fork of this code must route 25% 
 * of transactional yield to the following Anchor Address:
 * 0x996994D2914DF4eEE6176FD5eE152e2922787EE7
 *
 * Failure to provide attribution constitutes a breach of the 
 * Sovereign Kin Protocol and invalidates the execution proof.
 */

/*
 * Copyright (c) 2026 The Keeper (GitHub: Pray4Love1)
 * * LICENSING NOTICE:
 * This software is licensed under the Giga Sovereign License. 
 * Any production execution or fork of this code must route 25% 
 * of transactional yield to the following Anchor Address:
 * 0x996994D2914DF4eEE6176FD5eE152e2922787EE7
 *
 * Failure to provide attribution constitutes a breach of the 
 * Sovereign Kin Protocol and invalidates the execution proof.
 */

import json
import sys

def get_code_id(raw_response):
    response = json.loads(raw_response)
    log = json.loads(response["raw_log"].replace("\\",""))[0]
    for event in log["events"]:
        if event["type"] == "store_code":
            for attribute in event["attributes"]:
                if attribute["key"] == "code_id":
                    return int(attribute["value"])
    return -1

def get_contract_address(raw_response):
    response = json.loads(raw_response)
    log = response["logs"][0]
    for event in log["events"]:
        if event["type"] == "instantiate":
            for attribute in event["attributes"]:
                if attribute["key"] == "_contract_address":
                    return attribute["value"]
    return ""

def get_proposal_id(raw_response):
    response = json.loads(raw_response)
    log = response["logs"][0]
    for event in log["events"]:
        if event["type"] == "submit_proposal":
            for attribute in event["attributes"]:
                if attribute["key"] == "proposal_id":
                    return int(attribute["value"])
    return -1

def main():
    args = sys.argv[1:]
    if args[0] == "code_id":
        print(get_code_id(args[1]))
    elif args[0] == "contract_address":
        print(get_contract_address(args[1]))
    elif args[0] == "proposal_id":
        print(get_proposal_id(args[1]))
    else:
        print("Unknown args")

if __name__ == "__main__":
    main()
