package main

import "encoding/json"

// CopyAndPasteReq defines the action specific fields
type CopyAndPasteReq struct {
	Pasted bool
	FormID string `json:"formId"`
}

func copyAndPaste(body []byte, sessionData *Data) error {
	var data CopyAndPasteReq
	err := json.Unmarshal(body, &data)
	if err != nil {
		return err
	}

	sessionData.CopyAndPaste[data.FormID] = data.Pasted
	return nil
}

// ResizeWindowReq defines the action specific fields
type ResizeWindowReq struct {
	ResizeFrom Dimension
	ResizeTo   Dimension
}

func resizeWindow(body []byte, sessionData *Data) error {
	var data ResizeWindowReq
	err := json.Unmarshal(body, &data)
	if err != nil {
		return err
	}

	sessionData.ResizeFrom = data.ResizeFrom
	sessionData.ResizeTo = data.ResizeTo

	return nil
}

// TimeTakenReq defines the action specific fields
type TimeTakenReq struct {
	Time int
}

func timeTaken(body []byte, sessionData *Data) error {
	var data TimeTakenReq
	err := json.Unmarshal(body, &data)
	if err != nil {
		return err
	}

	sessionData.FormCompletionTime = data.Time
	return nil
}
