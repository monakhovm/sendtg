package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

type sendMessageRequest struct {
	ChatID string `json:"chat_id"`
	Text   string `json:"text"`
}

var rootCmd = &cobra.Command{
	Use:   "sendtg [chat_id] [message]",
	Short: "Sends Telegram message",
	Long:  `sendtg — is console utility for Telegram message sending.`,
	Args:  cobra.MaximumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {

		if len(args) == 0 {
			return cmd.Help()
		}

		if len(args) != 2 {
			fmt.Println("Error: chat_id and message must be specified")
			return cmd.Help()
		}

		_ = godotenv.Load()

		token := getBotToken(cmd)
		if token == "" {
			return fmt.Errorf("TELEGRAM_BOT_TOKEN not found. Please provide --token or add it to ENV/.env")
		}

		chatID := args[0]
		message := args[1]

		return sendMessage(token, chatID, message)
	},
}

func init() {

	rootCmd.Flags().StringP("token", "t", "", "Telegram Bot Token (highest priority)")

	rootCmd.AddCommand(completionCmd)
	rootCmd.AddCommand(versionCmd)
}

func getBotToken(cmd *cobra.Command) string {
	flagToken, _ := cmd.Flags().GetString("token")
	if flagToken != "" {
		return flagToken
	}

	envToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if envToken != "" {
		return envToken
	}

	return ""
}

func sendMessage(token, chatID, text string) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)

	body := sendMessageRequest{
		ChatID: chatID,
		Text:   text,
	}
	jsonBody, _ := json.Marshal(body)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Telegram API returned status %d", resp.StatusCode)
	}

	// fmt.Println("✅ Message sent successfully!")
	return nil
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
