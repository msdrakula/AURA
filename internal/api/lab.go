package api

import "net/http"

// Same wording as intel.Engine.RunStage so the UI can treat one message.
const errNeedLabConfirm = "active probe requires confirmation that you are testing your own target"

func rejectUnlessAuthorized(w http.ResponseWriter, authorized bool) bool {
	if authorized {
		return false
	}
	writeErr(w, 400, errNeedLabConfirm)
	return true
}
