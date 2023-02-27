// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

contract VoteContract {

    struct Voter {
        uint weight; 
        bool voted; 
        uint vote;
    }

    struct Candidat {
        bytes32 name;  
        uint voteCount; 
    }

    address public VoteCreator;

    mapping(address => Voter) public voters;

    Candidat[] public Candidats;

    uint public TotalCandidats;

    constructor(bytes32[] memory CandidatNames) {
        VoteCreator = msg.sender;
        voters[VoteCreator].weight = 1;

        for (uint i = 0; i < CandidatNames.length; i++) {
         Candidats.push(Candidat({
                name: CandidatNames[i],
                voteCount: 0
            }));
        }

        TotalCandidats = Candidats.length;
    }


    function giveRightToVote(address voter) public {
        require(
            msg.sender == VoteCreator,
            "Only VoteCreator can give right to vote."
        );
        require(
            !voters[voter].voted,
            "The voter already voted."
        );
        require(voters[voter].weight == 0);
        voters[voter].weight = 1;
    }

    function vote(uint Candidat) public {
        Voter storage sender = voters[msg.sender];
        require(sender.weight != 0, "Has no right to vote");
        require(!sender.voted, "Already voted.");
        sender.voted = true;
        sender.vote = Candidat;

     Candidats[Candidat].voteCount += sender.weight;
    }


    function winninCandidat() public view
            returns (uint winninCandidat_)
    {
        uint winningVoteCount = 0;
        for (uint p = 0; p < Candidats.length; p++) {
            if  (Candidats[p].voteCount > winningVoteCount) {
                winningVoteCount = Candidats[p].voteCount;
                winninCandidat_ = p;
            }
        }
    }

    function winnerName() public view
            returns (bytes32 winnerName_)
    {
        winnerName_ = Candidats[winninCandidat()].name;
    }
}