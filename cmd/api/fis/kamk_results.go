package fisapi

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/DeRuina/KUHA-REST-API/internal/auth/authz"
	"github.com/DeRuina/KUHA-REST-API/internal/store/cache"
	"github.com/DeRuina/KUHA-REST-API/internal/store/fis"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
)

// Handles competitor/result-related KAMK endpoints.
type ResultKAMKHandler struct {
	resultCC fis.Resultcc
	resultJP fis.Resultjp
	resultNK fis.Resultnk
	cache    *cache.Storage
}

func NewResultKAMKHandler(
	resultCC fis.Resultcc,
	resultJP fis.Resultjp,
	resultNK fis.Resultnk,
	cache *cache.Storage,
) *ResultKAMKHandler {
	return &ResultKAMKHandler{
		resultCC: resultCC,
		resultJP: resultJP,
		resultNK: resultNK,
		cache:    cache,
	}
}

// GetCompetitorSeasonsCatcodes godoc
//
//	@Summary		Get competitor seasons and categories
//	@Description	Gets distinct (Seasoncode, Catcode) combinations for the races a competitor has actually competed in, for a given sector (CC, JP, NK). Uses both race and result tables.
//	@Tags			FIS - KAMK
//	@Accept			json
//	@Produce		json
//	@Param			competitorid	query		int32	true	"Competitor ID"
//	@Param			sector			query		string	true	"Sector code (CC,JP,NK)"
//	@Success		200				{object}	swagger.FISCompetitorSeasonsCatcodesResponse
//	@Failure		400				{object}	swagger.ValidationErrorResponse
//	@Failure		401				{object}	swagger.UnauthorizedResponse
//	@Failure		403				{object}	swagger.ForbiddenResponse
//	@Failure		500				{object}	swagger.InternalServerErrorResponse
//	@Failure		503				{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/fis/competitor/seasons-catcodes [get]
func (h *ResultKAMKHandler) GetCompetitorSeasonsCatcodes(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	if err := utils.ValidateParams(r, []string{
		"competitorid", "sector",
	}); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	rawComp := strings.TrimSpace(r.URL.Query().Get("competitorid"))
	if rawComp == "" {
		utils.BadRequestResponse(w, r, fmt.Errorf("missing required query param: competitorid"))
		return
	}
	competitorID, err := utils.ParsePositiveInt32(rawComp)
	if err != nil {
		utils.BadRequestResponse(w, r, fmt.Errorf("invalid competitorid: %s", rawComp))
		return
	}

	sector := strings.TrimSpace(strings.ToUpper(r.URL.Query().Get("sector")))
	switch sector {
	case "CC", "JP", "NK":
	default:
		if sector == "" {
			utils.BadRequestResponse(w, r, fmt.Errorf("missing required query param: sector"))
		} else {
			utils.BadRequestResponse(w, r, fmt.Errorf("invalid sector: %s", sector))
		}
		return
	}

	type item struct {
		Seasoncode int32   `json:"seasoncode"`
		Catcode    *string `json:"catcode,omitempty"`
	}

	var items []item

	switch sector {
	case "CC":
		if h.resultCC == nil {
			utils.InternalServerError(w, r, fmt.Errorf("resultCC store not configured"))
			return
		}
		rows, err := h.resultCC.GetSeasonsCatcodesCCByCompetitor(r.Context(), competitorID)
		if err != nil {
			utils.InternalServerError(w, r, err)
			return
		}
		for _, row := range rows {
			if !row.Seasoncode.Valid {
				continue
			}
			var cat *string
			if row.Catcode.Valid && strings.TrimSpace(row.Catcode.String) != "" {
				c := strings.TrimSpace(row.Catcode.String)
				cat = &c
			}
			items = append(items, item{
				Seasoncode: row.Seasoncode.Int32,
				Catcode:    cat,
			})
		}

	case "JP":
		if h.resultJP == nil {
			utils.InternalServerError(w, r, fmt.Errorf("resultJP store not configured"))
			return
		}
		rows, err := h.resultJP.GetSeasonsCatcodesJPByCompetitor(r.Context(), competitorID)
		if err != nil {
			utils.InternalServerError(w, r, err)
			return
		}
		for _, row := range rows {
			if !row.Seasoncode.Valid {
				continue
			}
			var cat *string
			if row.Catcode.Valid && strings.TrimSpace(row.Catcode.String) != "" {
				c := strings.TrimSpace(row.Catcode.String)
				cat = &c
			}
			items = append(items, item{
				Seasoncode: row.Seasoncode.Int32,
				Catcode:    cat,
			})
		}

	case "NK":
		if h.resultNK == nil {
			utils.InternalServerError(w, r, fmt.Errorf("resultNK store not configured"))
			return
		}
		rows, err := h.resultNK.GetSeasonsCatcodesNKByCompetitor(r.Context(), competitorID)
		if err != nil {
			utils.InternalServerError(w, r, err)
			return
		}
		for _, row := range rows {
			if !row.Seasoncode.Valid {
				continue
			}
			var cat *string
			if row.Catcode.Valid && strings.TrimSpace(row.Catcode.String) != "" {
				c := strings.TrimSpace(row.Catcode.String)
				cat = &c
			}
			items = append(items, item{
				Seasoncode: row.Seasoncode.Int32,
				Catcode:    cat,
			})
		}
	}

	body := map[string]any{
		"competitorid": competitorID,
		"sector":       sector,
		"items":        items,
	}

	utils.WriteJSON(w, http.StatusOK, body)
}

// GetCompetitorLatestResults godoc
//
//	@Summary		Get competitor latest race results
//	@Description	Gets an athlete's latest race results filtered by sector and optional season and catcodes. Results are ordered by Racedate DESC.
//	@Tags			FIS - KAMK
//	@Accept			json
//	@Produce		json
//	@Param			competitorid	query		int32		true	"Competitor ID"
//	@Param			sector			query		string		true	"Sector code (CC,JP,NK)"
//	@Param			seasoncode		query		int32		false	"Season code filter"
//	@Param			catcode			query		[]string	false	"Category code(s) – repeat or comma-separated (e.g. catcode=WC&catcode=COC or catcode=WC,COC)"
//	@Param			limit			query		int32		false	"Maximum number of results to return (default 50)"
//	@Success		200				{object}	swagger.FISLatestResultsResponse
//	@Failure		400				{object}	swagger.ValidationErrorResponse
//	@Failure		401				{object}	swagger.UnauthorizedResponse
//	@Failure		403				{object}	swagger.ForbiddenResponse
//	@Failure		500				{object}	swagger.InternalServerErrorResponse
//	@Failure		503				{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/fis/competitor/latest-results [get]
func (h *ResultKAMKHandler) GetCompetitorLatestResults(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	if err := utils.ValidateParams(r, []string{
		"competitorid", "sector", "seasoncode", "catcode", "limit",
	}); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	rawComp := strings.TrimSpace(r.URL.Query().Get("competitorid"))
	if rawComp == "" {
		utils.BadRequestResponse(w, r, fmt.Errorf("missing required query param: competitorid"))
		return
	}
	competitorID, err := utils.ParsePositiveInt32(rawComp)
	if err != nil {
		utils.BadRequestResponse(w, r, fmt.Errorf("invalid competitorid: %s", rawComp))
		return
	}

	sector := strings.TrimSpace(strings.ToUpper(r.URL.Query().Get("sector")))
	switch sector {
	case "CC", "JP", "NK":
	default:
		if sector == "" {
			utils.BadRequestResponse(w, r, fmt.Errorf("missing required query param: sector"))
		} else {
			utils.BadRequestResponse(w, r, fmt.Errorf("invalid sector: %s", sector))
		}
		return
	}

	var seasonPtr *int32
	if sv := strings.TrimSpace(r.URL.Query().Get("seasoncode")); sv != "" {
		season, err := utils.ParsePositiveInt32(sv)
		if err != nil {
			utils.BadRequestResponse(w, r, fmt.Errorf("invalid seasoncode: %s", sv))
			return
		}
		seasonPtr = &season
	}

	parseList := func(key string) []string {
		vals := r.URL.Query()[key]
		if len(vals) == 1 && strings.Contains(vals[0], ",") {
			return strings.Split(vals[0], ",")
		}
		return vals
	}

	rawCats := parseList("catcode")
	var catcodes []string
	for _, c := range rawCats {
		c = strings.TrimSpace(strings.ToUpper(c))
		if c == "" {
			continue
		}
		catcodes = append(catcodes, c)
	}
	if len(catcodes) == 0 {
		catcodes = nil
	}

	const defaultLimit int32 = 50
	limit := defaultLimit
	if lv := strings.TrimSpace(r.URL.Query().Get("limit")); lv != "" {
		parsed, err := utils.ParsePositiveInt32(lv)
		if err != nil {
			utils.BadRequestResponse(w, r, fmt.Errorf("invalid limit: %s", lv))
			return
		}
		if parsed > 0 {
			limit = parsed
		}
	}

	type latestResultItem struct {
		Raceid         *int32  `json:"raceid,omitempty"`
		Racedate       *string `json:"racedate,omitempty"`
		Seasoncode     *int32  `json:"seasoncode,omitempty"`
		Disciplinecode *string `json:"disciplinecode,omitempty"`
		Catcode        *string `json:"catcode,omitempty"`
		Place          *string `json:"place,omitempty"`
		Nationcode     *string `json:"nationcode,omitempty"`
		Position       *string `json:"position,omitempty"`
		Timetot        *string `json:"timetot,omitempty"`
		Distance       *string `json:"distance,omitempty"`
		Hill           *int32  `json:"hill,omitempty"`
		Posr1          *string `json:"posr1,omitempty"`
		Speedr1        *string `json:"speedr1,omitempty"`
		Distr1         *string `json:"distr1,omitempty"`
		Judptsr1       *string `json:"judptsr1,omitempty"`
		Windr1         *string `json:"windr1,omitempty"`
		Windptsr1      *string `json:"windptsr1,omitempty"`
		Gater1         *string `json:"gater1,omitempty"`
		Poscc          *string `json:"poscc,omitempty"`
		Timetotint     *int32  `json:"timetotint,omitempty"`
		Pointsjump     *string `json:"pointsjump,omitempty"`
	}

	results := make([]latestResultItem, 0)

	switch sector {
	case "CC":
		if h.resultCC == nil {
			utils.InternalServerError(w, r, fmt.Errorf("resultCC store not configured"))
			return
		}
		rows, err := h.resultCC.GetLatestResultsCC(r.Context(), competitorID, seasonPtr, catcodes, &limit)
		if err != nil {
			utils.InternalServerError(w, r, err)
			return
		}

		for _, row := range rows {
			// raceid
			var raceIDPtr *int32
			if row.Raceid.Valid {
				v := row.Raceid.Int32
				raceIDPtr = &v
			}

			// seasoncode
			var seasonOut *int32
			if row.Seasoncode.Valid {
				v := row.Seasoncode.Int32
				seasonOut = &v
			}

			item := latestResultItem{
				Raceid:         raceIDPtr,
				Racedate:       utils.FormatDatePtr(row.Racedate),
				Seasoncode:     seasonOut,
				Disciplinecode: utils.StringPtrOrNil(row.Disciplinecode),
				Catcode:        utils.StringPtrOrNil(row.Catcode),
				Place:          utils.StringPtrOrNil(row.Place),
				Nationcode:     utils.StringPtrOrNil(row.Nationcode),
				Position: func() *string {
					if row.Position.Valid {
						p := strings.TrimSpace(row.Position.String)
						if p != "" {
							return &p
						}
					}
					return nil
				}(),
				Timetot: func() *string {
					if row.Timetot.Valid {
						t := strings.TrimSpace(row.Timetot.String)
						if t != "" {
							return &t
						}
					}
					return nil
				}(),
			}

			results = append(results, item)
		}

	case "JP":
		if h.resultJP == nil {
			utils.InternalServerError(w, r, fmt.Errorf("resultJP store not configured"))
			return
		}
		rows, err := h.resultJP.GetLatestResultsJP(r.Context(), competitorID, seasonPtr, catcodes, &limit)
		if err != nil {
			utils.InternalServerError(w, r, err)
			return
		}

		for _, row := range rows {
			var raceIDPtr *int32
			if row.Raceid.Valid {
				v := row.Raceid.Int32
				raceIDPtr = &v
			}

			var seasonOut *int32
			if row.Seasoncode.Valid {
				v := row.Seasoncode.Int32
				seasonOut = &v
			}

			var posPtr *string
			if row.Position.Valid {
				p := fmt.Sprintf("%d", row.Position.Int32)
				posPtr = &p
			}

			item := latestResultItem{
				Raceid:         raceIDPtr,
				Racedate:       utils.FormatDatePtr(row.Racedate),
				Seasoncode:     seasonOut,
				Disciplinecode: utils.StringPtrOrNil(row.Disciplinecode),
				Catcode:        utils.StringPtrOrNil(row.Catcode),
				Place:          utils.StringPtrOrNil(row.Place),
				Nationcode:     utils.StringPtrOrNil(row.Nationcode),
				Position:       posPtr,

				Posr1:     utils.StringPtrOrNil(row.Posr1),
				Speedr1:   utils.StringPtrOrNil(row.Speedr1),
				Distr1:    utils.StringPtrOrNil(row.Distr1),
				Judptsr1:  utils.StringPtrOrNil(row.Judptsr1),
				Windr1:    utils.StringPtrOrNil(row.Windr1),
				Windptsr1: utils.StringPtrOrNil(row.Windptsr1),
				Gater1:    utils.StringPtrOrNil(row.Gater1),
			}

			results = append(results, item)
		}

	case "NK":
		if h.resultNK == nil {
			utils.InternalServerError(w, r, fmt.Errorf("resultNK store not configured"))
			return
		}
		rows, err := h.resultNK.GetLatestResultsNK(r.Context(), competitorID, seasonPtr, catcodes, &limit)
		if err != nil {
			utils.InternalServerError(w, r, err)
			return
		}

		for _, row := range rows {
			var raceIDPtr *int32
			if row.Raceid.Valid {
				v := row.Raceid.Int32
				raceIDPtr = &v
			}

			var seasonOut *int32
			if row.Seasoncode.Valid {
				v := row.Seasoncode.Int32
				seasonOut = &v
			}

			var hillPtr *int32
			if row.Hill.Valid {
				v := row.Hill.Int32
				hillPtr = &v
			}

			var posPtr *string
			if row.Position.Valid {
				p := fmt.Sprintf("%d", row.Position.Int32)
				posPtr = &p
			}

			var timetotintPtr *int32
			if row.Timetotint.Valid {
				v := row.Timetotint.Int32
				timetotintPtr = &v
			}

			item := latestResultItem{
				Raceid:         raceIDPtr,
				Racedate:       utils.FormatDatePtr(row.Racedate),
				Seasoncode:     seasonOut,
				Disciplinecode: utils.StringPtrOrNil(row.Disciplinecode),
				Catcode:        utils.StringPtrOrNil(row.Catcode),
				Place:          utils.StringPtrOrNil(row.Place),
				Nationcode:     utils.StringPtrOrNil(row.Nationcode),
				Position:       posPtr,
				Timetot: func() *string {
					if row.Timetot.Valid {
						t := strings.TrimSpace(row.Timetot.String)
						if t != "" {
							return &t
						}
					}
					return nil
				}(),

				Distance:   utils.StringPtrOrNil(row.Distance),
				Hill:       hillPtr,
				Posr1:      utils.StringPtrOrNil(row.Posr1),
				Speedr1:    utils.StringPtrOrNil(row.Speedr1),
				Distr1:     utils.StringPtrOrNil(row.Distr1),
				Judptsr1:   utils.StringPtrOrNil(row.Judptsr1),
				Windr1:     utils.StringPtrOrNil(row.Windr1),
				Windptsr1:  utils.StringPtrOrNil(row.Windptsr1),
				Gater1:     utils.StringPtrOrNil(row.Gater1),
				Poscc:      utils.StringPtrOrNil(row.Poscc),
				Timetotint: timetotintPtr,
				Pointsjump: utils.StringPtrOrNil(row.Pointsjump),
			}

			results = append(results, item)
		}
	}

	body := map[string]any{
		"competitorid": competitorID,
		"sector":       sector,
		"seasoncode":   seasonPtr,
		"catcodes":     catcodes,
		"limit":        limit,
		"results":      results,
	}

	utils.WriteJSON(w, http.StatusOK, body)
}
