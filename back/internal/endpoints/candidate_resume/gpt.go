package candidate_resume

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
)

const (
	filesURL     = "https://api.openai.com/v1/files"
	responsesURL = "https://api.openai.com/v1/responses"
	model        = "gpt-4.1"
)

// ===== upload file =====

func UploadFile(file io.Reader, filename, apiKey string) (string, error) {

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	part, _ := writer.CreateFormFile("file", filename)
	io.Copy(part, file)

	writer.WriteField("purpose", "assistants")
	writer.Close()

	req, _ := http.NewRequest("POST", filesURL, &body)
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var res struct {
		ID string `json:"id"`
	}

	json.NewDecoder(resp.Body).Decode(&res)
	return res.ID, nil
}

// ===== GPT =====

type GPTResponse struct {
	Output []struct {
		Content []struct {
			Text string `json:"text"`
		} `json:"content"`
	} `json:"output"`
}

func askGPT(fileID, vacancy, apiKey string) (bool, error) {
	payload := map[string]interface{}{
		"model": model,
		"input": []map[string]interface{}{
			{
				"role": "system",
				"content": []map[string]string{
					{
						"type": "input_text",
						"text": "Ты HR ассистент. Отвечай строго от 1 - 10 (целочисленый тип).",
					},
				},
			},
			{
				"role": "user",
				"content": []map[string]string{
					{
						"type": "input_text",
						"text": "Вакансия:\n" + vacancy + "\n\nПодходит ли кандидат?",
					},
					{
						"type":    "input_file",
						"file_id": fileID,
					},
				},
			},
		},
	}

	data, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", responsesURL, bytes.NewBuffer(data))
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var gptResp GPTResponse
	if err := json.Unmarshal(body, &gptResp); err != nil {
		return false, err
	}

	if len(gptResp.Output) == 0 ||
		len(gptResp.Output[0].Content) == 0 {
		return false, fmt.Errorf("empty GPT response")
	}

	result := strings.ToLower(
		strings.TrimSpace(gptResp.Output[0].Content[0].Text),
	)

	resultScore, _ := strconv.Atoi(result)

	if resultScore < 7 {
		return false, nil
	}

	return true, nil
}
