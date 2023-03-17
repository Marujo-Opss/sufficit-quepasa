package controllers

import (
	"net/http"
	"sort"
	"strconv"
	"time"

	models "github.com/sufficit/sufficit-quepasa/models"
	whatsapp "github.com/sufficit/sufficit-quepasa/whatsapp"
)

func GetTimestamp(timestamp string) (result int64, err error) {
	if len(timestamp) > 0 {
		result, err = strconv.ParseInt(timestamp, 10, 64)
		if err != nil {
			if len(timestamp) > 0 {
				return
			} else {
				result = 0
			}
		}
	}
	return
}

// Retrieve messages with timestamp parameter
// Sorting then
func GetMessages(server *models.QpWhatsappServer, timestamp int64) (messages []whatsapp.WhatsappMessage) {
	searchTime := time.Unix(timestamp, 0)
	messages = server.GetMessages(searchTime)
	sort.Sort(whatsapp.ByTimestamp(messages))
	return
}

/*
<summary>
	Find a system track identifier to follow the message
	Getting from PATH => QUERY => HEADER
</summary>
*/
func GetTrackId(r *http.Request) string {
	return models.GetRequestParameter(r, "trackid")
}

/*
<summary>
	Get Picture Identifier of contact
	Getting from PATH => QUERY => HEADER
</summary>
*/
func GetPictureId(r *http.Request) string {
	return models.GetRequestParameter(r, "pictureid")
}

/*
<summary>
	Get Token From Http Request
	Getting from PATH => QUERY => HEADER
</summary>
*/
func GetToken(r *http.Request) string {
	return models.GetRequestParameter(r, "token")
}

/*
<summary>
	Get Message Id From Http Request
	Getting from PATH => QUERY => HEADER
</summary>
*/
func GetMessageId(r *http.Request) string {
	messageid := models.GetRequestParameter(r, "messageid")
	if len(messageid) == 0 {

		// compatibility with V3
		messageid = models.QueryGetValue(r.URL, "id")
	}
	return messageid
}

/*
<summary>
	Get File Name From Http Request
	Getting from PATH => QUERY => HEADER
</summary>
*/
func GetFileName(r *http.Request) string {
	return models.GetRequestParameter(r, "filename")
}

/*
<summary>
	Get Text Label From Http Request
	Getting from PATH => QUERY => HEADER
</summary>
*/
func GetTextParameter(r *http.Request) string {
	return models.GetRequestParameter(r, "text")
}

/*
<summary>
	Get a boolean indicating that cache should be used, From Http Request
	Getting from PATH => QUERY => HEADER
</summary>
*/
func GetCache(r *http.Request) bool {
	return models.ToBoolean(models.GetRequestParameter(r, "cache"))
}