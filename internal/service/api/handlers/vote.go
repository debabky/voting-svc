package handlers

import (
	"github.com/debabky/voting-svc/internal/data"
	"github.com/debabky/voting-svc/internal/service/api/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
	"time"
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
		ape.RenderErr(w, problems.BadRequest(errors.New("voting not found"))...)
		return
	}

	if voting.ActiveUntil.Before(time.Now().UTC()) {
		ape.RenderErr(w, problems.BadRequest(errors.New("voting ended"))...)
		return
	}

	// if there are no votes or multiple votes for non-ranked voting
	if len(req.Data.Votes) < 1 || (voting.Type != data.RankedVoting && len(req.Data.Votes) != 1) {
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
		ape.RenderErr(w, problems.BadRequest(errors.New("nullifier has already been used"))...)
		return
	}

	// TODO validate rank
	for _, vote := range req.Data.Votes {
		if err := MasterQ(r).VotesQ().Insert(data.Vote{
			//VotingID:     req.Data.VotingID, // FIXME
			VotingOption: vote.VotingOption,
			Nullifier:    TokenClaims(r).Nullifier,
			Rank:         vote.Rank,
		}); err != nil {
			Log(r).WithError(err).Error("failed to insert vote to the database")
			ape.RenderErr(w, problems.InternalError())
			return
		}
	}
}
