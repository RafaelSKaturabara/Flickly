package auth

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"net/url"

// 	"github.com/rkaturabara/flickly/internal/domain/users/entities"
// 	"github.com/rkaturabara/flickly/internal/domain/users/repositories"

// 	"github.com/gin-gonic/gin"
// )

// type OAuth2Service struct {
// 	config     *AuthConfig
// 	repository repositories.IUserRepository
// }

// func NewOAuth2Service(config *AuthConfig, repository repositories.IUserRepository) *OAuth2Service {
// 	return &OAuth2Service{
// 		config:     config,
// 		repository: repository,
// 	}
// }

// func (s *OAuth2Service) GetAuthURL() string {
// 	scopes := "email profile"
// 	return fmt.Sprintf(
// 		"https://accounts.google.com/o/oauth2/v2/auth?client_id=%s&response_type=code&scope=%s",
// 		s.config.ClientID,
// 		scopes,
// 	)
// }

// func (s *OAuth2Service) HandleCallback(c *gin.Context) (*entities.User, error) {
// 	code := c.Query("code")
// 	if code == "" {
// 		return nil, fmt.Errorf("código de autorização não fornecido")
// 	}

// 	// Trocar o código por um token de acesso
// 	token, err := s.exchangeCodeForToken(code)
// 	if err != nil {
// 		return nil, fmt.Errorf("erro ao trocar código por token: %v", err)
// 	}

// 	// Obter informações do usuário
// 	userInfo, err := s.getUserInfo(token)
// 	if err != nil {
// 		return nil, fmt.Errorf("erro ao obter informações do usuário: %v", err)
// 	}

// 	// Verificar se o usuário já existe
// 	existingUser, err := s.repository.GetUserByProviderID(c.Request.Context(), "google", userInfo.ID)
// 	if err != nil {
// 		// Criar novo usuário
// 		user := entities.NewUser(
// 			userInfo.Name,
// 			userInfo.Email,
// 			"google",
// 			userInfo.ID,
// 		)
// 		user.GivenName = userInfo.GivenName
// 		user.FamilyName = userInfo.FamilyName
// 		user.VerifiedEmail = userInfo.VerifiedEmail
// 		user.Picture = userInfo.Picture

// 		if err := s.repository.CreateUser(c.Request.Context(), user); err != nil {
// 			return nil, fmt.Errorf("erro ao criar usuário: %v", err)
// 		}

// 		return user, nil
// 	}

// 	// Atualizar informações do usuário existente
// 	existingUser.GivenName = userInfo.GivenName
// 	existingUser.FamilyName = userInfo.FamilyName
// 	existingUser.VerifiedEmail = userInfo.VerifiedEmail
// 	existingUser.Picture = userInfo.Picture

// 	if err := s.repository.UpdateUser(c.Request.Context(), existingUser); err != nil {
// 		return nil, fmt.Errorf("erro ao atualizar usuário: %v", err)
// 	}

// 	return existingUser, nil
// }

// func (s *OAuth2Service) exchangeCodeForToken(code string) (string, error) {
// 	data := url.Values{}
// 	data.Set("client_id", s.config.ClientID)
// 	data.Set("client_secret", s.config.ClientSecret)
// 	data.Set("code", code)
// 	data.Set("grant_type", "authorization_code")

// 	resp, err := http.PostForm("https://oauth2.googleapis.com/token", data)
// 	if err != nil {
// 		return "", fmt.Errorf("erro ao fazer requisição para o Google: %v", err)
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		return "", fmt.Errorf("erro ao trocar código por token: %s", resp.Status)
// 	}

// 	var tokenResp struct {
// 		AccessToken string `json:"access_token"`
// 		TokenType   string `json:"token_type"`
// 		ExpiresIn   int    `json:"expires_in"`
// 	}

// 	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
// 		return "", fmt.Errorf("erro ao decodificar resposta: %v", err)
// 	}

// 	return tokenResp.AccessToken, nil
// }

// func (s *OAuth2Service) getUserInfo(token string) (*GoogleUserInfo, error) {
// 	req, err := http.NewRequest("GET", "https://www.googleapis.com/oauth2/v2/userinfo", nil)
// 	if err != nil {
// 		return nil, fmt.Errorf("erro ao criar requisição: %v", err)
// 	}

// 	req.Header.Set("Authorization", "Bearer "+token)

// 	resp, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		return nil, fmt.Errorf("erro ao fazer requisição para o Google: %v", err)
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		return nil, fmt.Errorf("erro ao obter informações do usuário: %s", resp.Status)
// 	}

// 	var userInfo GoogleUserInfo
// 	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
// 		return nil, fmt.Errorf("erro ao decodificar resposta: %v", err)
// 	}

// 	return &userInfo, nil
// }

// type GoogleUserInfo struct {
// 	ID            string `json:"id"`
// 	Email         string `json:"email"`
// 	VerifiedEmail bool   `json:"verified_email"`
// 	Name          string `json:"name"`
// 	GivenName     string `json:"given_name"`
// 	FamilyName    string `json:"family_name"`
// 	Picture       string `json:"picture"`
// }
