// SPDX-License-Identifier: AGPL-3.0-only
pragma solidity 0.8.12;

contract CosmosERC20Relayer {
    // NOTE: ibc/CDC458787... is a hash of the denom, the IBC port, and the channel.

    event Cosmos__ConvertCoin(
        address sender,
        address receiver,
        uint256 amount,
        address tokenAddress
    );

    event Cosmos__ConvertERC20(
        address sender,
        address receiver,
        uint256 amount,
        address tokenAddress
    );
    event Cosmos__SendERC20ViaIBCEthAddress(
        address sender,
        address receiver,
        uint256 amount,
        address tokenAddress,
        string sourcePort,
        string sourceChannel
    );
    event Cosmos__SendERC20ViaIBCBech32Address(
        address sender,
        string receiver,
        uint256 amount,
        address tokenAddress,
        string sourcePort,
        string sourceChannel
    );

    function ICS20toERC20(
        address receiver,
        uint256 amount,
        address tokenAddress
    ) public {
        emit Cosmos__ConvertCoin(msg.sender, receiver, amount, tokenAddress);
    }

    function ERC20toICS20(
        address receiver,
        uint256 amount,
        address tokenAddress
    ) public {
        emit Cosmos__ConvertERC20(msg.sender, receiver, amount, tokenAddress);
    }

    function sendTokensViaIBCEthAddress(
        address receiver,
        uint256 amount,
        address tokenAddress,
        string calldata sourcePort,
        string calldata sourceChannel
    ) public {
        emit Cosmos__SendERC20ViaIBCEthAddress(
            msg.sender,
            receiver,
            amount,
            tokenAddress,
            sourcePort,
            sourceChannel
        );
    }

    function sendTokensViaIBCBech32Address(
        string calldata receiver,
        uint256 amount,
        address tokenAddress,
        string calldata sourcePort,
        string calldata sourceChannel
    ) public {
        emit Cosmos__SendERC20ViaIBCBech32Address(
            msg.sender,
            receiver,
            amount,
            tokenAddress,
            sourcePort,
            sourceChannel
        );
    }
}
