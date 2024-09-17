//go:build skip_wowometer_send
// +build skip_wowometer_send

package parameters

import "log"

const SKIP_SEND = true

func init() {
	log.Println("[i] Wowometer will NOT send to Google Form")
}
