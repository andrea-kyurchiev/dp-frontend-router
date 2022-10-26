package analytics

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/ONSdigital/log.go/v2/log"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

var _ ServiceBackend = &sqsBackend{}

//go:generate moq -out analyticstest/sqsclient.go -pkg analyticstest . SQSClient
type SQSClient interface {
	SendMessage(ctx context.Context, params *sqs.SendMessageInput, optFns ...func(*sqs.Options)) (*sqs.SendMessageOutput, error)
}

type sqsBackend struct {
	sqsClient SQSClient
	queueURL  string
}

// NewSQSBackend creates a new SQS backend for storing analytics data
func NewSQSBackend(ctx context.Context, queueURL string) (ServiceBackend, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	sqsClient := sqs.NewFromConfig(cfg)
	return &sqsBackend{
		sqsClient,
		queueURL,
	}, nil
}

func (b *sqsBackend) Store(req *http.Request, url, term, listType, gaID, gID string, pageIndex, linkIndex, pageSize float64) {
	var data = map[string]interface{}{
		"created":   time.Now().Format(time.RFC3339),
		"url":       url,
		"term":      term,
		"listType":  listType,
		"gaID":      gaID, // 2 year expiration cookie (_ga)
		"gID":       gID,  // 24 hour expiration cookie (_gid)
		"pageIndex": pageIndex,
		"linkIndex": linkIndex,
		"pageSize":  pageSize,
	}

	jb, err := json.Marshal(&data)
	if err != nil {
		log.Error(req.Context(), "error marshaling json", err)
		return
	}

	strJSON := string(jb)
	smi := &sqs.SendMessageInput{
		MessageBody: &strJSON,
		QueueUrl:    &b.queueURL,
	}

	smo, err := b.sqsClient.SendMessage(req.Context(), smi)
	if err != nil {
		log.Error(req.Context(), "error sending sqs message", err)
		return
	}

	log.Info(req.Context(), "stored analytics data in SQS", log.Data{"message_id": *smo.MessageId})
}
