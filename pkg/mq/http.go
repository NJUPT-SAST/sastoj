package mq

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type Client struct {
	URL    string
	User   string
	Passwd string
}

// QueueStats represents the statistics of a RabbitMQ queue
type QueueStats struct {
	Name                string  `json:"name"`
	VHost               string  `json:"vhost"`
	Messages            int64   `json:"messages"`
	MessagesReady       int64   `json:"messages_ready"`
	MessagesUnacked     int64   `json:"messages_unacknowledged"`
	Consumers           int     `json:"consumers"`
	Memory              int64   `json:"memory"`
	State               string  `json:"state"`
	MessageRateIn       float64 `json:"message_rate_in,omitempty"`
	MessageRateOut      float64 `json:"message_rate_out,omitempty"`
	MessageRateAck      float64 `json:"message_rate_ack,omitempty"`
	MessageBytesReady   int64   `json:"message_bytes_ready,omitempty"`
	MessageBytesUnacked int64   `json:"message_bytes_unacknowledged,omitempty"`
	MessageBytes        int64   `json:"message_bytes,omitempty"`
}

// RabbitMQOverview represents the overview statistics from RabbitMQ
type RabbitMQOverview struct {
	QueueTotals struct {
		Messages        int64 `json:"messages"`
		MessagesReady   int64 `json:"messages_ready"`
		MessagesUnacked int64 `json:"messages_unacknowledged"`
	} `json:"queue_totals"`
	ObjectTotals struct {
		Consumers   int `json:"consumers"`
		Queues      int `json:"queues"`
		Exchanges   int `json:"exchanges"`
		Connections int `json:"connections"`
		Channels    int `json:"channels"`
	} `json:"object_totals"`
	MessageStats struct {
		Publish        int64 `json:"publish"`
		PublishDetails struct {
			Rate float64 `json:"rate"`
		} `json:"publish_details"`
		DeliverGet        int64 `json:"deliver_get"`
		DeliverGetDetails struct {
			Rate float64 `json:"rate"`
		} `json:"deliver_get_details"`
	} `json:"message_stats"`
}

// GetQueueStats retrieves statistics for all queues from RabbitMQ
func (c *Client) GetQueueStats(ctx context.Context) ([]QueueStats, error) {
	logger := log.NewHelper(log.With(log.GetLogger(), "module", "mq"))

	// Construct the management API URL for queues
	apiURL := fmt.Sprintf("http://%s/api/queues", c.URL)

	// Create a new request
	req, err := http.NewRequestWithContext(ctx, "GET", apiURL, nil)
	if err != nil {
		logger.Errorf("failed to create request: %v", err)
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set basic auth credentials (using default guest/guest for simplicity)
	// In production, you should use proper credentials from configuration
	req.SetBasicAuth(c.User, c.Passwd)

	// Send the request
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		logger.Errorf("failed to send request: %v", err)
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		logger.Errorf("API request failed with status: %d", resp.StatusCode)
		return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Errorf("failed to read response body: %v", err)
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse the JSON response
	var queues []struct {
		Name            string `json:"name"`
		VHost           string `json:"vhost"`
		Messages        int64  `json:"messages"`
		MessagesReady   int64  `json:"messages_ready"`
		MessagesUnacked int64  `json:"messages_unacknowledged"`
		Consumers       int    `json:"consumers"`
		Memory          int64  `json:"memory"`
		State           string `json:"state"`
		MessageStats    *struct {
			PublishDetails struct {
				Rate float64 `json:"rate"`
			} `json:"publish_details,omitempty"`
			DeliverGetDetails struct {
				Rate float64 `json:"rate"`
			} `json:"deliver_get_details,omitempty"`
			AckDetails struct {
				Rate float64 `json:"rate"`
			} `json:"ack_details,omitempty"`
		} `json:"message_stats,omitempty"`
		MessageBytesReady   int64 `json:"message_bytes_ready,omitempty"`
		MessageBytesUnacked int64 `json:"message_bytes_unacknowledged,omitempty"`
		MessageBytes        int64 `json:"message_bytes,omitempty"`
	}

	if err := json.Unmarshal(body, &queues); err != nil {
		logger.Errorf("failed to unmarshal response: %v", err)
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	// Convert to QueueStats
	result := make([]QueueStats, 0, len(queues))
	for _, q := range queues {
		stats := QueueStats{
			Name:                q.Name,
			VHost:               q.VHost,
			Messages:            q.Messages,
			MessagesReady:       q.MessagesReady,
			MessagesUnacked:     q.MessagesUnacked,
			Consumers:           q.Consumers,
			Memory:              q.Memory,
			State:               q.State,
			MessageBytesReady:   q.MessageBytesReady,
			MessageBytesUnacked: q.MessageBytesUnacked,
			MessageBytes:        q.MessageBytes,
		}

		if q.MessageStats != nil {
			stats.MessageRateIn = q.MessageStats.PublishDetails.Rate
			stats.MessageRateOut = q.MessageStats.DeliverGetDetails.Rate
			stats.MessageRateAck = q.MessageStats.AckDetails.Rate
		}

		result = append(result, stats)
	}

	return result, nil
}

// GetMQOverview retrieves overview statistics from RabbitMQ
func (c *Client) GetMQOverview(ctx context.Context) (*RabbitMQOverview, error) {
	logger := log.NewHelper(log.With(log.GetLogger(), "module", "mq"))

	// Construct the management API URL for overview
	apiURL := fmt.Sprintf("http://%s/api/overview", c.URL)

	// Create a new request
	req, err := http.NewRequestWithContext(ctx, "GET", apiURL, nil)
	if err != nil {
		logger.Errorf("failed to create request: %v", err)
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set basic auth credentials (using default guest/guest for simplicity)
	// In production, you should use proper credentials from configuration
	req.SetBasicAuth(c.User, c.Passwd)

	// Send the request
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		logger.Errorf("failed to send request: %v", err)
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		logger.Errorf("API request failed with status: %d", resp.StatusCode)
		return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Errorf("failed to read response body: %v", err)
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse the JSON response
	var overview RabbitMQOverview
	if err := json.Unmarshal(body, &overview); err != nil {
		logger.Errorf("failed to unmarshal response: %v", err)
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &overview, nil
}

// GetQueueStatsForVHost retrieves statistics for all queues in a specific virtual host
func (c *Client) GetQueueStatsForVHost(ctx context.Context, vhost string) ([]QueueStats, error) {
	logger := log.NewHelper(log.With(log.GetLogger(), "module", "mq"))

	// Construct the management API URL for queues in the specific vhost
	apiURL := fmt.Sprintf("http://%s/api/queues/%s", c.URL, vhost)

	// Create a new request
	req, err := http.NewRequestWithContext(ctx, "GET", apiURL, nil)
	if err != nil {
		logger.Errorf("failed to create request: %v", err)
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set basic auth credentials (using default guest/guest for simplicity)
	// In production, you should use proper credentials from configuration
	req.SetBasicAuth(c.User, c.Passwd)

	// Send the request
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		logger.Errorf("failed to send request: %v", err)
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		logger.Errorf("API request failed with status: %d", resp.StatusCode)
		return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Errorf("failed to read response body: %v", err)
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse the JSON response
	var queues []struct {
		Name            string `json:"name"`
		VHost           string `json:"vhost"`
		Messages        int64  `json:"messages"`
		MessagesReady   int64  `json:"messages_ready"`
		MessagesUnacked int64  `json:"messages_unacknowledged"`
		Consumers       int    `json:"consumers"`
		Memory          int64  `json:"memory"`
		State           string `json:"state"`
		MessageStats    *struct {
			PublishDetails struct {
				Rate float64 `json:"rate"`
			} `json:"publish_details,omitempty"`
			DeliverGetDetails struct {
				Rate float64 `json:"rate"`
			} `json:"deliver_get_details,omitempty"`
			AckDetails struct {
				Rate float64 `json:"rate"`
			} `json:"ack_details,omitempty"`
		} `json:"message_stats,omitempty"`
		MessageBytesReady   int64 `json:"message_bytes_ready,omitempty"`
		MessageBytesUnacked int64 `json:"message_bytes_unacknowledged,omitempty"`
		MessageBytes        int64 `json:"message_bytes,omitempty"`
	}

	if err := json.Unmarshal(body, &queues); err != nil {
		logger.Errorf("failed to unmarshal response: %v", err)
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	// Convert to QueueStats
	result := make([]QueueStats, 0, len(queues))
	for _, q := range queues {
		stats := QueueStats{
			Name:                q.Name,
			VHost:               q.VHost,
			Messages:            q.Messages,
			MessagesReady:       q.MessagesReady,
			MessagesUnacked:     q.MessagesUnacked,
			Consumers:           q.Consumers,
			Memory:              q.Memory,
			State:               q.State,
			MessageBytesReady:   q.MessageBytesReady,
			MessageBytesUnacked: q.MessageBytesUnacked,
			MessageBytes:        q.MessageBytes,
		}

		if q.MessageStats != nil {
			stats.MessageRateIn = q.MessageStats.PublishDetails.Rate
			stats.MessageRateOut = q.MessageStats.DeliverGetDetails.Rate
			stats.MessageRateAck = q.MessageStats.AckDetails.Rate
		}

		result = append(result, stats)
	}

	return result, nil
}

func NewClient(url string, user string, passwd string) *Client {
	return &Client{
		URL:    url,
		User:   user,
		Passwd: passwd,
	}
}
