package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func handlerChirpsValidate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}
	type returnVals struct {
		CleanedBody string `json:"cleaned_body"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	const maxChirpLength = 140
	if len(params.Body) > maxChirpLength {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return
	}

	cleanResp := strings.ReplaceAll(params.Body, "kerfuffle","****")
	cleanResp = strings.ReplaceAll(cleanResp, "sharbert","****")
	cleanResp = strings.ReplaceAll(cleanResp, "fornax","****")

	cleanResp = strings.ReplaceAll(cleanResp, "Kerfuffle","****")
	cleanResp = strings.ReplaceAll(cleanResp, "Sharbert","****")
	cleanResp = strings.ReplaceAll(cleanResp, "Fornax","****")

	
	respondWithJSON(w, http.StatusOK, returnVals{
		CleanedBody: cleanResp,
	})
}
