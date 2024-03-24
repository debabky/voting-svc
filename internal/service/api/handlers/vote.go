package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/debabky/voting-svc/internal/data"
	"github.com/debabky/voting-svc/internal/service/api/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func Vote(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewVoteRequest(r)
	if err != nil {
		Log(r).WithError(err).Error("failed to get vote request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	voting, err := MasterQ(r).VotingsQ().New().
		FilterBy("id", TokenClaims(r).VotingID).
		Get()
	if err != nil {
		Log(r).WithError(err).Error("failed to get voting")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if voting == nil {
		Log(r).Error("voting not found")
		ape.RenderErr(w, problems.BadRequest(errors.New("voting not found"))...)
		return
	}

	if voting.ActiveUntil.Before(time.Now().UTC()) {
		Log(r).Error("voting ended")
		ape.RenderErr(w, problems.BadRequest(errors.New("voting ended"))...)
		return
	}

	// if there are no votes or multiple votes for non-ranked voting
	if len(req.Data.Attributes.Votes) < 1 || (voting.Type != data.RankedVoting && len(req.Data.Attributes.Votes) != 1) {
		Log(r).Error("insufficient number of votes")
		ape.RenderErr(w, problems.BadRequest(errors.New("insufficient number of votes"))...)
		return
	}

	registration, err := MasterQ(r).RegistrationsQ().New().
		FilterBy("voting_id", TokenClaims(r).VotingID).
		FilterBy("nullifier", TokenClaims(r).Nullifier).
		Get()
	if err != nil {
		Log(r).WithError(err).Error("failed to get registration")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if registration == nil {
		Log(r).Error("registration not found")
		ape.RenderErr(w, problems.BadRequest(errors.New("registration not found"))...)
		return
	}

	votesCount, err := MasterQ(r).VotesQ().New().
		FilterBy("voting_id", TokenClaims(r).VotingID).
		FilterBy("nullifier", TokenClaims(r).Nullifier).
		Count()
	if err != nil {
		Log(r).WithError(err).Error("failed to get votes")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if votesCount > 0 {
		Log(r).Error("nullifier has already been used")
		ape.RenderErr(w, problems.BadRequest(errors.New("nullifier has already been used"))...)
		return
	}

	switch voting.Type {
	case data.RankedVoting:
		if err := addRankedVotes(r, voting, req.Data.Attributes.Votes); err != nil {
			Log(r).WithError(err).Error("failed to add ranked votes")
			ape.RenderErr(w, problems.InternalError())
			return
		}
	default:
		Log(r).Error(fmt.Sprintf("%d is not supported voting type", voting.Type))
		ape.RenderErr(w, problems.InternalError())
		return
	}
}

func addRankedVotes(r *http.Request, voting *data.Voting, votes []requests.Vote) error {
	passedOptions := make([]string, 0)
	for _, vote := range votes {
		passedOptions = append(passedOptions, vote.VotingOption)
	}

	options, err := MasterQ(r).VotingOptionsQ().
		FilterBy("voting_id", voting.ID.String()).
		FilterBy("name", passedOptions).
		Select()
	if err != nil {
		return errors.Wrap(err, "failed to get voting options")
	}
	if len(options) == 0 {
		return errors.New("voting options not found")
	}

	if len(options) != len(votes) {
		return errors.New("votes number is not equal to options number")
	}

	if err := MasterQ(r).Transaction(func(db data.MasterQ) error {
		for i, passedOption := range passedOptions {
			rank := int64(i + 1)

			if err := db.VotesQ().Insert(data.Vote{
				VotingID:     TokenClaims(r).VotingID,
				VotingOption: passedOption,
				Nullifier:    TokenClaims(r).Nullifier,
				Rank:         &rank,
			}); err != nil {
				return errors.Wrap(err, "failed to insert vote to the database")
			}
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "failed to perform SQL transaction")
	}

	return nil
}
