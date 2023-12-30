package keys

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	VkToken        string
	OwnerId        int64
	GroupDomain    string
	Filter         string
	Count          int
	TelegramChatId string
)

func LoadEnvVariables(filename string) {
	if err := godotenv.Load(filename); err != nil {
		log.Println("No .env files found")
		return
	}

	VkToken, _ = os.LookupEnv("VK_API_TOKEN")
	ownerIDStr, _ := os.LookupEnv("OWNER_ID")
	OwnerId, _ = strconv.ParseInt(ownerIDStr, 10, 64)
	GroupDomain, _ = os.LookupEnv("GROUP_DOMAIN")
	Filter, _ = os.LookupEnv("FILTER")
	countString, _ := os.LookupEnv("COUNT")
	count, _ := strconv.ParseInt(countString, 10, 32)
	Count = int(count)
	TelegramChatId, _ = os.LookupEnv("chat_id")

	fmt.Printf("%T", OwnerId)
}
