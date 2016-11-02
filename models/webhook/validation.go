package webhook

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

var allowedEvents = []string{
	"Account.Created", "Account.Deleted", "Account.Updated", "Alert.Notification", "Server.Created", "Server.Deleted", "Server.Updated",
	"User.Created", "User.Deleted", "User.Updated"}

// ValidateTargetURI provides validation checks for the targetUri of a webhook
func ValidateTargetURI(targetURI string) error {

	validURL := govalidator.IsURL(targetURI)
	if !validURL {
		return fmt.Errorf("TargetUri: invalid URI.")
	}

	url, _ := url.Parse(targetURI)
	if strings.ToLower(url.Scheme) != "https" {
		return fmt.Errorf("TargetUri: must be an HTTPS endpoint.")
	}

	return nil
}

// ValidateEvent provides validation checks for the event name of a webhook
func ValidateEvent(eventName string) error {

	if stringInSlice(eventName, allowedEvents) == false {
		return fmt.Errorf("Event: %s is an invalid event. Allowed events are: %s", eventName, allowedEvents)
	}

	return nil
}

func stringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}
