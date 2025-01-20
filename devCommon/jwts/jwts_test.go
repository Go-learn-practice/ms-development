package jwts

import (
	"encoding/json"
	"testing"
	"time"
)

func TestCreateToken(t *testing.T) {
	token := CreateToken("123", time.Hour, "ms_project", time.Hour*24, "ms_project")
	marshal, _ := json.Marshal(token)
	t.Log(string(marshal))
}

func TestParseToken(t *testing.T) {
	// accessToken
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzcxODc2MjYsInRva2VuIjoiMTIzIn0.csfPDql0lgR5UlFrfLqQD8l0UFlezFjfP5YwRdklPng"
	ParseToken(tokenString, "ms_project")
}
