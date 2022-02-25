// SPDX-License-Identifier: AGPL-3.0-only
pragma solidity 0.8.12;

import {ERC20} from "@solmate/tokens/ERC20.sol";
import {SafeTransferLib} from "@solmate/utils/SafeTransferLib.sol";

contract CosmosSynapse {
    using SafeTransferLib for ERC20;

    address immutable SYNAPSE_MODULE;

    event Cosmos__Synapse__BridgeOut(
        address token,
        uint256 amount,
        address sender,
        address receiver,
        string destChain
    );

    constructor(address _synapseModule) {
        SYNAPSE_MODULE = _synapseModule;
    }

    function bridgeOut(
        address token,
        uint256 amount,
        string calldata destChain
    ) external {
        _bridgeOut(token, amount, destChain, msg.sender, msg.sender);
    }

    function bridgeOut(
        address token,
        uint256 amount,
        string calldata destChain,
        address receiver
    ) internal {
        _bridgeOut(token, amount, destChain, msg.sender, receiver);
    }

    function bridgeOut(
        address token,
        uint256 amount,
        string calldata destChain,
        address sender,
        address receiver
    ) internal {
        _bridgeOut(token, amount, destChain, sender, receiver);
    }

    function _bridgeOut(
        address token,
        uint256 amount,
        string calldata destChain,
        address sender,
        address receiver
    ) internal {
        ERC20(token).safeTransferFrom(sender, SYNAPSE_MODULE, amount);
        // Emitting this without doing a safety check on sender is okay, because
        // safeTransferFrom will revert if sender has not allowed CosmosSynapse to
        // transfer `amount` tokens
        emit Cosmos__Synapse__BridgeOut(
            token,
            amount,
            sender,
            receiver,
            destChain
        );
    }
}
