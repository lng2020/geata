package mock

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/spf13/cobra"
)

type Message struct {
	Ref   string `json:"ref"`
	Value string `json:"value"`
}

var server, topic string
var interval time.Duration
var msg Message

func NewMockCmd() *cobra.Command {
	mockCmd := &cobra.Command{
		Use:   "mock",
		Short: "Starts a mock MQTT client that sends messages periodically",
		Run:   runMock,
	}

	mockCmd.Flags().StringVarP(&server, "server", "s", "tcp://localhost:1883", "MQTT server URL")
	mockCmd.Flags().StringVarP(&topic, "topic", "t", "1", "MQTT topic to publish messages")
	mockCmd.Flags().DurationVarP(&interval, "interval", "i", 5*time.Second, "Interval between messages")
	mockCmd.Flags().StringVarP(&msg.Ref, "ref", "r", "ZTPDFMONT/SFGD1$FireState$stVal", "Reference of the message")
	mockCmd.Flags().StringVarP(&msg.Value, "value", "v", "1", "Value of the message")

	return mockCmd
}

func runMock(cmd *cobra.Command, args []string) {
	opts := mqtt.NewClientOptions().AddBroker(server).SetClientID("geata_mock_client")
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println("Error connecting to MQTT broker:", token.Error())
		os.Exit(1)
	}

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		message, err := json.Marshal(msg)
		if err != nil {
			fmt.Println("Error marshalling message:", err)
			continue
		}
		token := client.Publish(topic, 0, false, message)
		token.Wait()
		fmt.Printf("Published message to %s: %s\n", topic, message)
	}
}
