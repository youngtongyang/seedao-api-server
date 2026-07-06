package service

import "errors"

var (
	ErrProposalNotFound      = errors.New("proposal not found")
	ErrProposalForbidden     = errors.New("proposal forbidden")
	ErrProposalInvalidStatus = errors.New("proposal invalid status")
	ErrProposalVoteClosed    = errors.New("proposal vote closed")
	ErrProposalAlreadyVoted  = errors.New("proposal already voted")
	ErrProposalNotVoter      = errors.New("proposal not voter")
	ErrProposalSNSRequired   = errors.New("proposal sns required")
	ErrProposalVoteNotFound  = errors.New("proposal vote not found")
)
