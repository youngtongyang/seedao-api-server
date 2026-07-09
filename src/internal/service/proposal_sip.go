package service

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/seedao/seedao-api-server/internal/model"
)

var ErrProposalAssociateSIPMissing = errors.New("associate proposal sip not assigned")

// closeProjectProposalTypes inherit SIP from associate_proposal_id instead of allocating a new number.
var closeProjectProposalTypes = map[string]struct{}{
	"P2项目结项提案": {},
	"P3项目结项提案": {},
	"P2项目关闭提案": {},
	"P3项目关闭提案": {},
}

func isCloseProjectType(proposalType string) bool {
	_, ok := closeProjectProposalTypes[proposalType]
	return ok
}

func sipAssigned(sip int) bool {
	return sip > 0
}

func (s *ProposalService) maybeAssignSIPOnVotingOpen(p *model.Proposal) (*model.Proposal, error) {
	if p == nil {
		return nil, ErrProposalNotFound
	}
	if sipAssigned(p.SIP) || !s.isVotingOpen(p) {
		return p, nil
	}
	if _, err := s.AssignSIP(p.ID); err != nil {
		return p, err
	}
	return s.getProposal(p.ID)
}

// AssignSIP allocates a global SIP when the proposal enters the formal voting window.
// Idempotent: returns existing sip when already assigned.
func (s *ProposalService) AssignSIP(proposalID int64) (int, error) {
	var assigned int
	err := s.db.Transaction(func(tx *gorm.DB) error {
		var p model.Proposal
		q := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&p, proposalID)
		if errors.Is(q.Error, gorm.ErrRecordNotFound) {
			return ErrProposalNotFound
		}
		if q.Error != nil {
			return q.Error
		}
		if sipAssigned(p.SIP) {
			assigned = p.SIP
			return nil
		}

		sipNum, err := s.resolveSIPNumber(tx, &p)
		if err != nil {
			return err
		}

		res := tx.Model(&model.Proposal{}).
			Where("id = ? AND sip = 0", p.ID).
			Update("sip", sipNum)
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			if err := tx.First(&p, proposalID).Error; err != nil {
				return err
			}
			assigned = p.SIP
			return nil
		}
		assigned = sipNum
		return nil
	})
	return assigned, err
}

func (s *ProposalService) resolveSIPNumber(tx *gorm.DB, p *model.Proposal) (int, error) {
	if isCloseProjectType(p.Type) {
		if p.AssociateProposalID == nil || *p.AssociateProposalID <= 0 {
			return 0, &ProposalParamError{Reason: "close-project proposal requires associate_proposal_id"}
		}
		var parent model.Proposal
		if err := tx.First(&parent, *p.AssociateProposalID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return 0, ErrProposalNotFound
			}
			return 0, err
		}
		if !sipAssigned(parent.SIP) {
			return 0, ErrProposalAssociateSIPMissing
		}
		return parent.SIP, nil
	}
	return s.allocateNextSIP(tx)
}

func (s *ProposalService) allocateNextSIP(tx *gorm.DB) (int, error) {
	if tx.Dialector.Name() == "postgres" {
		var sip int64
		if err := tx.Raw("SELECT nextval('proposal_sip_seq')").Scan(&sip).Error; err != nil {
			return 0, fmt.Errorf("allocate sip from sequence: %w", err)
		}
		return int(sip), nil
	}

	var maxSIP int
	if err := tx.Model(&model.Proposal{}).
		Select("COALESCE(MAX(sip), 0)").
		Scan(&maxSIP).Error; err != nil {
		return 0, err
	}
	return maxSIP + 1, nil
}
