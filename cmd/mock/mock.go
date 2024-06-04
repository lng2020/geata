package mock

import (
	"fmt"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/spf13/cobra"
)

var server, topic string
var interval time.Duration

func NewMockCmd() *cobra.Command {
	mockCmd := &cobra.Command{
		Use:   "mock",
		Short: "Starts a mock MQTT client that sends messages periodically",
		Run:   runMock,
	}

	mockCmd.Flags().StringVarP(&server, "server", "s", "tcp://localhost:1883", "MQTT server URL")
	mockCmd.Flags().StringVarP(&topic, "topic", "t", "1", "MQTT topic to publish messages")
	mockCmd.Flags().DurationVarP(&interval, "interval", "i", 5*time.Second, "Interval between messages")

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
		message := fmt.Sprintf("Mock message at %s", time.Now().Format(time.RFC3339))
		token := client.Publish(topic, 0, false, message)
		token.Wait()
		fmt.Printf("Published message to %s: %s\n", topic, message)
	}
}
