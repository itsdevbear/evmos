// SPDX-License-Identifier: APGL-3.0
pragma solidity 0.8.12;

import {ERC20} from "@solmate/tokens/ERC20.sol";

/**
 * @dev reeeee better on gas than openzepplinnn $er!!!
 */
contract CosmosRelayedERC20 is ERC20 {
    /*///////////////////////////////////////////////////////////////
                        VALIDATOR ADDRESS STORAGE
    //////////////////////////////////////////////////////////////*/

    address immutable AUTH;

    /*///////////////////////////////////////////////////////////////
                               CONSTRUCTOR
    //////////////////////////////////////////////////////////////*/

    constructor(
        string memory name,
        string memory symbol,
        uint8 decimals_
    ) ERC20(name, symbol, decimals_) {
        AUTH = msg.sender;
    }

    /*///////////////////////////////////////////////////////////////
                       EXTERNAL MINT/BURN LOGIC
    //////////////////////////////////////////////////////////////*/

    function mint(address to, uint256 amount) external {
        require(
            msg.sender == AUTH,
            "CosmosRelayedERC20: must have minter role to mint"
        );
        _mint(to, amount);
    }

    function burn(uint256 amount) external {
        _burn(msg.sender, amount);
    }

    function burnFrom(address account, uint256 amount) external {
        allowance[account][msg.sender] = amount;
        emit Approval(account, msg.sender, allowance[account][msg.sender]);
        _burn(account, amount);
    }

    function burnCoins(address from, uint256 amount) external {
        require(
            msg.sender == AUTH,
            "CosmosRelayedERC20: must have burner role to burn"
        );
        _burn(from, amount);
    }
}
