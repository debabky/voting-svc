package handlers

import (
	"github.com/debabky/voting-svc/internal/service/api/requests"
	"github.com/debabky/voting-svc/internal/service/jwt"
	"github.com/debabky/voting-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"
)

func GetAuthData(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewGetAuthDataRequest(r)
	if err != nil {
		Log(r).WithError(err).Error("failed to get auth data request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	verificationRequest, err := MasterQ(r).VerificationRequestsQ().New().
		FilterBy("id", req.ID).
		Get()
	if err != nil {
		Log(r).WithError(err).Error("failed to get verification request")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if verificationRequest == nil {
		ape.RenderErr(w, problems.NotFound())
		return
	}

	access, err := JWTIssuer(r).IssueJWT(
		verificationRequest.VotingID.String(), verificationRequest.Nullifier, jwt.AccessTokenType,
	)
	if err != nil {
		Log(r).WithError(err).Error("failed to issuer JWT token")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	refresh, err := JWTIssuer(r).IssueJWT(
		verificationRequest.VotingID.String(), verificationRequest.Nullifier, jwt.RefreshTokenType,
	)
	if err != nil {
		Log(r).WithError(err).Error("failed to issuer JWT token")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	resp := resources.TokenResponse{
		Data: resources.Token{
			Key: resources.Key{
				ID:   req.ID.String(),
				Type: resources.TOKEN,
			},
			Attributes: resources.TokenAttributes{
				AccessToken: resources.Jwt{
					Token:     access,
					TokenType: string(jwt.AccessTokenType),
				},
				RefreshToken: resources.Jwt{
					Token:     refresh,
					TokenType: string(jwt.RefreshTokenType),
				},
			},
		},
	}

	Cookies(r).SetTokensCookies(w, access, refresh)
	ape.Render(w, resp)
}
