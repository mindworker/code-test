package main

import "encoding/json"

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
