package handler

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ory/fosite"
	"github.com/ory/fosite/compose"
	"github.com/ory/fosite/storage"
	"github.com/ory/fosite/token/jwt"
	"golang.org/x/oauth2"
)

var (
	clientConf = oauth2.Config{
		ClientID:     "my-client",
		ClientSecret: "foobar",
		RedirectURL:  "http://localhost:3846/callback",
		Scopes:       []string{"offline", "openid", "photos"},
		Endpoint: oauth2.Endpoint{
			TokenURL: "http://localhost:3846/token",
			AuthURL:  "http://localhost:3846/auth",
		},
	}

	config = &fosite.Config{
		// AccessTokenLifespan:  time.Minute * 10,
		// RefreshTokenLifespan: time.Minute * 30,
		GlobalSecret: []byte("some-cool-secret-that-is-32bytes"),
	}
	store = storage.NewExampleStore()

	privateKey, _ = rsa.GenerateKey(rand.Reader, 2048)
	strategy      = compose.CommonStrategy{
		CoreStrategy: compose.NewOAuth2HMACStrategy(config),
		OpenIDConnectTokenStrategy: compose.NewOpenIDConnectStrategy(func(ctx context.Context) (interface{}, error) {
			return privateKey, nil
		}, config),
		Signer: &jwt.DefaultSigner{GetPrivateKey: func(ctx context.Context) (interface{}, error) {
			return privateKey, nil
		}},
	}

	factories = []compose.Factory{
		compose.OAuth2AuthorizeExplicitFactory,
		compose.OAuth2AuthorizeImplicitFactory,
		compose.OAuth2RefreshTokenGrantFactory,
		compose.OAuth2ClientCredentialsGrantFactory,
		compose.OAuth2TokenIntrospectionFactory,
		compose.OAuth2TokenRevocationFactory,
		compose.OAuth2PKCEFactory,
	}

	oauth2Compose = compose.Compose(config, store, strategy, factories...)
)

type OAuth2Handler interface {
	HomeEndpoint(ginCtx *gin.Context)
	AuthEndpoint(ginCtx *gin.Context)
	CallbackEndpoint(ginCtx *gin.Context)
	TokenEndpoint(ginCtx *gin.Context)
	RefreshTokenEndpoint(ginCtx *gin.Context)
}

func (h *handler) HomeEndpoint(ginCtx *gin.Context) {
	ginCtx.Redirect(http.StatusFound, clientConf.AuthCodeURL("some-random-state-foobar"+"&nonce=some-random-nonce"))
}

func (h *handler) AuthEndpoint(ginCtx *gin.Context) {
	ctx := context.Background()
	ar, err := oauth2Compose.NewAuthorizeRequest(ctx, ginCtx.Request)
	if err != nil {
		oauth2Compose.WriteAuthorizeError(ctx, ginCtx.Writer, ar, err)
		return
	}

	for _, scope := range clientConf.Scopes {
		ar.GrantScope(scope)
	}

	sessionData := &fosite.DefaultSession{
		Subject: "peter",
		Extra:   make(map[string]interface{}),
	}
	response, err := oauth2Compose.NewAuthorizeResponse(ctx, ar, sessionData)
	if err != nil {
		oauth2Compose.WriteAuthorizeError(ctx, ginCtx.Writer, ar, err)
		return
	}

	oauth2Compose.WriteAuthorizeResponse(ctx, ginCtx.Writer, ar, response)
}

func (h *handler) CallbackEndpoint(ginCtx *gin.Context) {

	code := ginCtx.Request.URL.Query().Get("code")
	ctx := context.Background()

	var opts []oauth2.AuthCodeOption

	token, err := clientConf.Exchange(ctx, code, opts...)
	if err != nil {
		ginCtx.AbortWithStatusJSON(500, gin.H{"error": fmt.Sprintf("I tried to exchange the authorize code for an access token but it did not work but got error: %s", err.Error())})
		return
	}

	ginCtx.JSON(200, gin.H{
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
		"token_type":    token.TokenType,
		"expiry":        token.Expiry,
	})
}

func (h *handler) TokenEndpoint(ginCtx *gin.Context) {
	ctx := context.Background()
	mySessionData := &fosite.DefaultSession{
		Subject:   "peter",
		ExpiresAt: map[fosite.TokenType]time.Time{
			// fosite.AccessToken:  time.Now().Add(time.Hour),
			// fosite.RefreshToken: time.Now().Add(time.Hour * 24),
		},
		Extra: make(map[string]interface{}),
	}

	accessRequest, err := oauth2Compose.NewAccessRequest(ctx, ginCtx.Request, mySessionData)
	if err != nil {
		oauth2Compose.WriteAccessError(ctx, ginCtx.Writer, accessRequest, err)
		return
	}

	if accessRequest.GetGrantTypes().ExactOne("client_credentials") {
		for _, scope := range accessRequest.GetRequestedScopes() {
			accessRequest.GrantScope(scope)
		}
	}

	response, err := oauth2Compose.NewAccessResponse(ctx, accessRequest)
	if err != nil {
		log.Printf("Error occurred in NewAccessResponse: %+v", err)
		oauth2Compose.WriteAccessError(ctx, ginCtx.Writer, accessRequest, err)
		return
	}

	oauth2Compose.WriteAccessResponse(ctx, ginCtx.Writer, accessRequest, response)
}

func (h *handler) RefreshTokenEndpoint(ginCtx *gin.Context) {
	refreshToken := ginCtx.Request.URL.Query().Get("refresh")
	ctx := context.Background()

	oldToken := &oauth2.Token{
		RefreshToken: refreshToken,
	}
	tokenSource := clientConf.TokenSource(ctx, oldToken)
	token, err := tokenSource.Token()
	if err != nil {
		ginCtx.JSON(500, gin.H{"error": fmt.Sprintf("Could not refresh token %s", err)})
		return
	}

	ginCtx.JSON(200, gin.H{
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
		"token_type":    token.TokenType,
	})
}
