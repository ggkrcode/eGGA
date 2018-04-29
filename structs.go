
package main

import "time"

type WebHookResp struct {
	Speech string `json:"speech"`
	Displaytext string `json:"displayText"`
}

type WebHookRequest struct {
	OriginalRequest struct {
		Source  string `json:"source"`
		Version string `json:"version"`
		Data    struct {
			IsInSandbox bool `json:"isInSandbox"`
			Surface     struct {
				Capabilities []struct {
					Name string `json:"name"`
				} `json:"capabilities"`
			} `json:"surface"`
			Inputs []struct {
				RawInputs []struct {
					Query     string `json:"query"`
					InputType string `json:"inputType"`
				} `json:"rawInputs"`
				Arguments []struct {
					RawText   string `json:"rawText"`
					TextValue string `json:"textValue"`
					Name      string `json:"name"`
				} `json:"arguments"`
				Intent string `json:"intent"`
			} `json:"inputs"`
			User struct {
				LastSeen time.Time `json:"lastSeen"`
				Locale   string    `json:"locale"`
				UserID   string    `json:"userId"`
			} `json:"user"`
			Conversation struct {
				ConversationID    string `json:"conversationId"`
				Type              string `json:"type"`
				ConversationToken string `json:"conversationToken"`
			} `json:"conversation"`
			AvailableSurfaces []struct {
				Capabilities []struct {
					Name string `json:"name"`
				} `json:"capabilities"`
			} `json:"availableSurfaces"`
		} `json:"data"`
	} `json:"originalRequest"`
	ID        string    `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	Lang      string    `json:"lang"`
	Result    struct {
		Source           string `json:"source"`
		ResolvedQuery    string `json:"resolvedQuery"`
		Speech           string `json:"speech"`
		Action           string `json:"action"`
		ActionIncomplete bool   `json:"actionIncomplete"`
		Parameters       struct {
			Status  string `json:"status"`
			Network string `json:"network"`
		} `json:"parameters"`
		Contexts []struct {
			Name       string `json:"name"`
			Parameters struct {
				NetworkOriginal string `json:"network.original"`
				StatusOriginal  string `json:"status.original"`
				Network         string `json:"network"`
				Status          string `json:"status"`
			} `json:"parameters"`
			Lifespan int `json:"lifespan"`
		} `json:"contexts"`
		Metadata struct {
			MatchedParameters []struct {
				DataType string `json:"dataType"`
				Name     string `json:"name"`
				Value    string `json:"value"`
				IsList   bool   `json:"isList"`
			} `json:"matchedParameters"`
			IntentName                string `json:"intentName"`
			IsResponseToSlotfilling   bool   `json:"isResponseToSlotfilling"`
			IntentID                  string `json:"intentId"`
			WebhookUsed               string `json:"webhookUsed"`
			WebhookForSlotFillingUsed string `json:"webhookForSlotFillingUsed"`
			NluResponseTime           int    `json:"nluResponseTime"`
		} `json:"metadata"`
		Fulfillment struct {
			Speech   string `json:"speech"`
			Messages []struct {
				Type   int    `json:"type"`
				Speech string `json:"speech"`
			} `json:"messages"`
		} `json:"fulfillment"`
		Score float64 `json:"score"`
	} `json:"result"`
	Status struct {
		Code            int    `json:"code"`
		ErrorType       string `json:"errorType"`
		WebhookTimedOut bool   `json:"webhookTimedOut"`
	} `json:"status"`
	SessionID string `json:"sessionId"`
}