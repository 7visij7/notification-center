package types

import (
	"os"
	"strconv"
)

var (
	BotTokenDevops			= os.Getenv("BotTokenDevops")
	BotTokenMarkirovka		= os.Getenv("BotTokenMarkirovka")
	TelegramURL				= os.Getenv("TelegramURL")
	ChatIdMarkirovka, _		= strconv.Atoi(os.Getenv("ChatIdMarkirovka"))
	ChatIdEcomm, _			= strconv.Atoi(os.Getenv("ChatIdEcomm"))
	ChatIdDevops, _			= strconv.Atoi(os.Getenv("ChatIdDevops"))
	ChatIdOVSSupport, _		= strconv.Atoi(os.Getenv("ChatIdOVSSupport"))
	ChatIdMarkirovkaKTTeam, _ = strconv.Atoi(os.Getenv("ChatIdMarkirovkaKTTeam"))
	ChatIdOVSDBA, _ 		= strconv.Atoi(os.Getenv("ChatIdOVSDBA"))
	Authorization			= os.Getenv("Authorization")
	SmbServer				= os.Getenv("SmbServer")
	SmbCatalog				= os.Getenv("SmbCatalog")
	SmbUser					= os.Getenv("SmbUser")
	SmbPassword				= os.Getenv("SmbPassword")
	Filecsv					= os.Getenv("Filecsv")
	GrafanaUser				= os.Getenv("GrafanaUser")
	GrafanaPassword			= os.Getenv("GrafanaPassword")
	RedashUser				= os.Getenv("RedashUser")
	RedashPassword			= os.Getenv("RedashPassword")
)
