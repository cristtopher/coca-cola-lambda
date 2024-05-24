package handlers

import (
	"encoding/json"
	"net/http"

	"coca-cola-lambda/internal/adapters/s3"
	"coca-cola-lambda/internal/adapters/secretmanager"
	"coca-cola-lambda/internal/domain"

	"github.com/aws/aws-lambda-go/events"
)

type HTTPHandler struct {
	clientService        *domain.ClientService
	s3Adapter            *s3.S3Adapter
	secretManagerAdapter *secretmanager.SecretManagerAdapter
}

func NewHTTPHandler(clientService *domain.ClientService, s3Adapter *s3.S3Adapter, secretManagerAdapter *secretmanager.SecretManagerAdapter) *HTTPHandler {
	return &HTTPHandler{
		clientService:        clientService,
		s3Adapter:            s3Adapter,
		secretManagerAdapter: secretManagerAdapter,
	}
}

func (h *HTTPHandler) HandleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {
	case "GET":
		return h.handleGetClients()
	case "POST":
		return h.handleCreateClient(req)
	default:
		return events.APIGatewayProxyResponse{StatusCode: http.StatusMethodNotAllowed}, nil
	}
}

func (h *HTTPHandler) handleGetClients() (events.APIGatewayProxyResponse, error) {
	clients, err := h.clientService.GetAllClients()
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, nil
	}
	body, err := json.Marshal(clients)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, nil
	}
	return events.APIGatewayProxyResponse{StatusCode: http.StatusOK, Body: string(body)}, nil
}

func (h *HTTPHandler) handleCreateClient(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var client domain.Client
	err := json.Unmarshal([]byte(req.Body), &client)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest}, nil
	}
	err = h.clientService.CreateClient(client)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, nil
	}
	return events.APIGatewayProxyResponse{StatusCode: http.StatusCreated}, nil
}
