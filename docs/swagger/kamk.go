package swagger

type KamkAddInjuryRequest struct {
	UserID      int32   `json:"user_id" example:"27353728"`
	InjuryType  int32   `json:"injury_type" example:"1"`
	Severity    *int32  `json:"severity,omitempty" example:"3"`
	PainLevel   *int32  `json:"pain_level,omitempty" example:"7"`
	Description *string `json:"description,omitempty" example:"Left ankle sprain during training"`
	InjuryID    int32   `json:"injury_id" example:"2"`
	Meta        *string `json:"meta,omitempty" example:"phase=preseason"`
}

type KamkMarkRecoveredRequest struct {
	UserID   int32 `json:"user_id" example:"27353728"`
	InjuryID int32 `json:"injury_id" example:"2"`
}

type KamkInjuryItem struct {
	CompetitorID int32   `json:"competitor_id" example:"27353728"`
	InjuryType   int32   `json:"injury_type" example:"1"`
	Severity     *int32  `json:"severity,omitempty" example:"3"`
	PainLevel    *int32  `json:"pain_level,omitempty" example:"7"`
	Description  *string `json:"description,omitempty" example:"Left ankle sprain during training"`
	DateStart    string  `json:"date_start" example:"2025-01-10T09:30:00Z"`
	Status       int32   `json:"status" example:"0"`
	DateEnd      *string `json:"date_end,omitempty" example:"2025-01-17T12:00:00Z"`
	InjuryID     *int32  `json:"injury_id,omitempty" example:"2"`
	Meta         *string `json:"meta,omitempty" example:"phase=preseason"`
}

type KamkInjuriesListResponse struct {
	Injuries []KamkInjuryItem `json:"injuries"`
}

type KamkMaxInjuryIDResponse struct {
	ID int32 `json:"id" example:"3"`
}

type KamkAddQuestionnaireRequest struct {
	UserID    int32   `json:"user_id" example:"27353728"`
	QueryType *int32  `json:"query_type,omitempty" example:"1"`
	Answers   *string `json:"answers,omitempty" example:"{\"mood\":7,\"sleep\":8}"`
	Comment   *string `json:"comment,omitempty" example:"Felt good this morning"`
	Meta      *string `json:"meta,omitempty" example:"build=ios,v=2.3.1"`
}

type KamkCreateQuestionnaireResponse struct {
	ID int64 `json:"id" example:"123"`
}

type KamkQuestionnaireItem struct {
	ID           int64   `json:"id" example:"123"`
	CompetitorID int32   `json:"competitor_id" example:"27353728"`
	QueryType    *int32  `json:"query_type,omitempty" example:"1"`
	Answers      *string `json:"answers,omitempty" example:"{\"mood\":7,\"sleep\":8}"`
	Comment      *string `json:"comment,omitempty" example:"Felt good this morning"`
	Timestamp    string  `json:"timestamp" example:"2025-10-20T09:30:00Z"`
	Meta         *string `json:"meta,omitempty" example:"build=ios,v=2.3.1"`
}

type KamkQuestionnairesListResponse struct {
	Questionnaires []KamkQuestionnaireItem `json:"questionnaires"`
}

type KamkUpdateQuestionnaireBody struct {
	Answers string  `json:"answers" example:"{\"mood\":5,\"sleep\":6}"`
	Comment *string `json:"comment,omitempty" example:"A bit tired today"`
}
