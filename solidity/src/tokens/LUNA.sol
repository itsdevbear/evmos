// SPDX-License-Identifier: AGPL-3.0-only
pragma solidity 0.8.12;

import {ERC20} from "@solmate/tokens/ERC20.sol";

contract LUNA is ERC20("Luna", "LUNA", 18) {
    address private immutable VALIDATOR_ERC20_MODULE;

    constructor(address erc20ModuleAddress) {
        VALIDATOR_ERC20_MODULE = erc20ModuleAddress;
    }

    function mint(address to, uint256 amount) public virtual {
        require(msg.sender == VALIDATOR_ERC20_MODULE, "LUNA: Only Validators");
        _mint(to, amount);
    }

    function burnCoins(address from, uint256 amount) public virtual {
        require(msg.sender == VALIDATOR_ERC20_MODULE, "LUNA: Only Validators");
        _burn(from, amount);
    }
}
