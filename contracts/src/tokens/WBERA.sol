// SPDX-License-Identifier: AGPL-3.0-only
pragma solidity 0.8.12;

import {ERC20} from "@solmate/tokens/ERC20.sol";

import {SafeTransferLib} from "@solmate/utils/SafeTransferLib.sol";

/// @notice Minimalist and modern Wrapped Bera implementation.
/// @author Inspired by WETH9 (https://github.com/dapphub/ds-weth/blob/master/src/weth9.sol)
contract WBERA is ERC20("Wrapped Bera", "WBERA", 18) {
    using SafeTransferLib for address;

    event Deposit(address indexed from, uint256 amount);

    event Withdrawal(address indexed to, uint256 amount);

    function deposit() public payable {
        _mint(msg.sender, msg.value);

        emit Deposit(msg.sender, msg.value);
    }

    function withdraw(uint256 amount) public {
        _burn(msg.sender, amount);

        emit Withdrawal(msg.sender, amount);

        msg.sender.safeTransferETH(amount);
    }

    receive() external payable {
        deposit();
    }
}
