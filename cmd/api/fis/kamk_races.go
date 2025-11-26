package fisapi

import (
	"fmt"
	"net/http"
	"sort"
	"strings"

	"github.com/DeRuina/KUHA-REST-API/internal/auth/authz"
	"github.com/DeRuina/KUHA-REST-API/internal/store/cache"
	"github.com/DeRuina/KUHA-REST-API/internal/store/fis"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
)

type RaceSearchItem struct {
	Sectorcode     string  `json:"sectorcode" example:"CC"`
	Gender         *string `json:"gender,omitempty" example:"M"`
	Raceid         int32   `json:"raceid" example:"123456"`
	Racedate       *string `json:"racedate,omitempty" example:"2025-02-15"` // YYYY-MM-DD
	Catcode        *string `json:"catcode,omitempty" example:"WC"`
	Description    *string `json:"description,omitempty" example:"World Cup Sprint"`
	Place          *string `json:"place,omitempty" example:"Lahti"`
	Nationcode     *string `json:"nationcode,omitempty" example:"FIN"`
	Disciplinecode *string `json:"disciplinecode,omitempty" example:"DSPR"`
}

type RaceSearchHandler struct {
	raceCC fis.Racecc
	raceJP fis.Racejp
	raceNK fis.Racenk
	cache  *cache.Storage
}

func NewRaceSearchHandler(
	raceCC fis.Racecc,
	raceJP fis.Racejp,
	raceNK fis.Racenk,
	cache *cache.Storage,
) *RaceSearchHandler {
	return &RaceSearchHandler{
		raceCC: raceCC,
		raceJP: raceJP,
		raceNK: raceNK,
		cache:  cache,
	}
}

// SearchRaces godoc
//
//	@Summary		Search races across sectors
//	@Description	Gets all Races with optional filters (Nationcode, Seasoncode, Gender, Catcode) for each selected sector (NK, CC, JP). If sector is omitted, all sectors are used.
//	@Tags			FIS - KAMK
//	@Accept			json
//	@Produce		json
//	@Param			sector		query		[]string	false	"Sector code (CC,JP,NK – repeat or comma-separated; if omitted, all sectors are used)"
//	@Param			seasoncode	query		int32		false	"Season code"
//	@Param			nationcode	query		string		false	"Nation code (e.g. FIN)"
//	@Param			gender		query		string		false	"Gender (M/W)"
//	@Param			catcode		query		string		false	"Category code (e.g. WC)"
//	@Success		200			{object}	swagger.FISRacesSearchResponse
//	@Failure		400			{object}	swagger.ValidationErrorResponse
//	@Failure		401			{object}	swagger.UnauthorizedResponse
//	@Failure		403			{object}	swagger.ForbiddenResponse
//	@Failure		500			{object}	swagger.InternalServerErrorResponse
//	@Failure		503			{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/fis/races/search [get]
func (h *RaceSearchHandler) SearchRaces(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	if err := utils.ValidateParams(r, []string{
		"sector", "seasoncode", "nationcode", "gender", "catcode",
	}); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	parseList := func(key string) []string {
		vals := r.URL.Query()[key]
		if len(vals) == 1 && strings.Contains(vals[0], ",") {
			return strings.Split(vals[0], ",")
		}
		return vals
	}

	sectorsRaw := parseList("sector")

	sectorSet := make(map[string]struct{})

	for _, s := range sectorsRaw {
		s = strings.TrimSpace(strings.ToUpper(s))
		if s == "" {
			continue
		}
		switch s {
		case "CC", "JP", "NK":
			sectorSet[s] = struct{}{}
		default:
			utils.BadRequestResponse(w, r, fmt.Errorf("invalid sector: %s", s))
			return
		}
	}

	if len(sectorSet) == 0 {
		sectorSet["CC"] = struct{}{}
		sectorSet["JP"] = struct{}{}
		sectorSet["NK"] = struct{}{}
	}

	sectors := make([]string, 0, len(sectorSet))
	for _, s := range []string{"CC", "JP", "NK"} {
		if _, ok := sectorSet[s]; ok {
			sectors = append(sectors, s)
		}
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

	var nationPtr *string
	if nv := strings.TrimSpace(r.URL.Query().Get("nationcode")); nv != "" {
		nv = strings.ToUpper(nv)
		nationPtr = &nv
	}

	var genderPtr *string
	if gv := strings.TrimSpace(r.URL.Query().Get("gender")); gv != "" {
		gv = strings.ToUpper(gv)
		genderPtr = &gv
	}

	var catPtr *string
	if cv := strings.TrimSpace(r.URL.Query().Get("catcode")); cv != "" {
		cv = strings.ToUpper(cv)
		catPtr = &cv
	}

	var results []RaceSearchItem

	for _, sector := range sectors {
		switch sector {
		case "CC":
			if h.raceCC != nil {
				rows, err := h.raceCC.SearchRacesCC(r.Context(), seasonPtr, nationPtr, genderPtr, catPtr)
				if err != nil {
					utils.InternalServerError(w, r, err)
					return
				}
				for _, row := range rows {
					results = append(results, RaceSearchItem{
						Sectorcode:     row.Sectorcode, // "CC"
						Gender:         utils.StringPtrOrNil(row.Gender),
						Raceid:         row.Raceid,
						Racedate:       utils.FormatDatePtr(row.Racedate),
						Catcode:        utils.StringPtrOrNil(row.Catcode),
						Description:    utils.StringPtrOrNil(row.Description),
						Place:          utils.StringPtrOrNil(row.Place),
						Nationcode:     utils.StringPtrOrNil(row.Nationcode),
						Disciplinecode: utils.StringPtrOrNil(row.Disciplinecode),
					})
				}
			}

		case "JP":
			if h.raceJP != nil {
				rows, err := h.raceJP.SearchRacesJP(r.Context(), seasonPtr, nationPtr, genderPtr, catPtr)
				if err != nil {
					utils.InternalServerError(w, r, err)
					return
				}
				for _, row := range rows {
					results = append(results, RaceSearchItem{
						Sectorcode:     row.Sectorcode, // "JP"
						Gender:         utils.StringPtrOrNil(row.Gender),
						Raceid:         row.Raceid,
						Racedate:       utils.FormatDatePtr(row.Racedate),
						Catcode:        utils.StringPtrOrNil(row.Catcode),
						Description:    utils.StringPtrOrNil(row.Description),
						Place:          utils.StringPtrOrNil(row.Place),
						Nationcode:     utils.StringPtrOrNil(row.Nationcode),
						Disciplinecode: utils.StringPtrOrNil(row.Disciplinecode),
					})
				}
			}

		case "NK":
			if h.raceNK != nil {
				rows, err := h.raceNK.SearchRacesNK(r.Context(), seasonPtr, nationPtr, genderPtr, catPtr)
				if err != nil {
					utils.InternalServerError(w, r, err)
					return
				}
				for _, row := range rows {
					results = append(results, RaceSearchItem{
						Sectorcode:     row.Sectorcode, // "NK"
						Gender:         utils.StringPtrOrNil(row.Gender),
						Raceid:         row.Raceid,
						Racedate:       utils.FormatDatePtr(row.Racedate),
						Catcode:        utils.StringPtrOrNil(row.Catcode),
						Description:    utils.StringPtrOrNil(row.Description),
						Place:          utils.StringPtrOrNil(row.Place),
						Nationcode:     utils.StringPtrOrNil(row.Nationcode),
						Disciplinecode: utils.StringPtrOrNil(row.Disciplinecode),
					})
				}
			}
		}
	}

	body := map[string]any{
		"races": results,
	}

	utils.WriteJSON(w, http.StatusOK, body)
}

// GetRacesByIDs godoc
//
//	@Summary		Get races by IDs
//	@Description	Gets race(s) for a given sector and one or more race IDs.
//	@Tags			FIS - KAMK
//	@Accept			json
//	@Produce		json
//	@Param			sector	query		string	true	"Sector code (CC,JP,NK)"
//	@Param			raceid	query		[]int32	true	"Race ID(s) – repeat or comma-separated (e.g. raceid=123&raceid=456 or raceid=123,456)"
//	@Success		200		{object}	swagger.FISRacesByIDsResponse
//	@Failure		400		{object}	swagger.ValidationErrorResponse
//	@Failure		401		{object}	swagger.UnauthorizedResponse
//	@Failure		403		{object}	swagger.ForbiddenResponse
//	@Failure		500		{object}	swagger.InternalServerErrorResponse
//	@Failure		503		{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/fis/races/by-ids [get]
func (h *RaceSearchHandler) GetRacesByIDs(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	if err := utils.ValidateParams(r, []string{
		"sector", "raceid",
	}); err != nil {
		utils.BadRequestResponse(w, r, err)
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

	parseList := func(key string) []string {
		vals := r.URL.Query()[key]
		if len(vals) == 1 && strings.Contains(vals[0], ",") {
			return strings.Split(vals[0], ",")
		}
		return vals
	}

	raceIDStrs := parseList("raceid")
	if len(raceIDStrs) == 0 {
		utils.BadRequestResponse(w, r, fmt.Errorf("missing required query param: raceid"))
		return
	}

	var raceIDs []int32
	for _, raw := range raceIDStrs {
		raw = strings.TrimSpace(raw)
		if raw == "" {
			continue
		}
		id, err := utils.ParsePositiveInt32(raw)
		if err != nil {
			utils.BadRequestResponse(w, r, fmt.Errorf("invalid raceid: %s", raw))
			return
		}
		raceIDs = append(raceIDs, id)
	}

	if len(raceIDs) == 0 {
		utils.BadRequestResponse(w, r, fmt.Errorf("no valid raceid values provided"))
		return
	}

	var races []any

	switch sector {
	case "CC":
		if h.raceCC == nil {
			utils.InternalServerError(w, r, fmt.Errorf("raceCC store not configured"))
			return
		}
		rows, err := h.raceCC.GetRacesByIDsCC(r.Context(), raceIDs)
		if err != nil {
			utils.InternalServerError(w, r, err)
			return
		}
		for _, row := range rows {
			races = append(races, FISRaceCCFullFromSqlc(row))
		}

	case "JP":
		if h.raceJP == nil {
			utils.InternalServerError(w, r, fmt.Errorf("raceJP store not configured"))
			return
		}
		rows, err := h.raceJP.GetRacesByIDsJP(r.Context(), raceIDs)
		if err != nil {
			utils.InternalServerError(w, r, err)
			return
		}
		for _, row := range rows {
			races = append(races, FISRaceJPFullFromSqlc(row))
		}

	case "NK":
		if h.raceNK == nil {
			utils.InternalServerError(w, r, fmt.Errorf("raceNK store not configured"))
			return
		}
		rows, err := h.raceNK.GetRacesByIDsNK(r.Context(), raceIDs)
		if err != nil {
			utils.InternalServerError(w, r, err)
			return
		}
		for _, row := range rows {
			races = append(races, FISRaceNKFullFromSqlc(row))
		}
	}

	body := map[string]any{
		"sector": sector,
		"races":  races,
	}

	utils.WriteJSON(w, http.StatusOK, body)
}

// GetRaceCategoryCounts godoc
//
//	@Summary		Get race counts by category
//	@Description	Gets counts of races grouped by category code (Catcode) for a given season and selected sector(s). Optional filters: Nationcode, Gender.
//	@Tags			FIS - KAMK
//	@Accept			json
//	@Produce		json
//	@Param			seasoncode	query		int32		true	"Season code"
//	@Param			sector		query		[]string	true	"Sector code(s) (CC,JP,NK – repeat or comma-separated, e.g. sector=CC&sector=JP or sector=CC,JP)"
//	@Param			nationcode	query		string		false	"Nation code filter (e.g. FIN)"
//	@Param			gender		query		string		false	"Gender filter (e.g. M/W)"
//	@Success		200			{object}	swagger.FISRacesCategoryCountsResponse
//	@Failure		400			{object}	swagger.ValidationErrorResponse
//	@Failure		401			{object}	swagger.UnauthorizedResponse
//	@Failure		403			{object}	swagger.ForbiddenResponse
//	@Failure		500			{object}	swagger.InternalServerErrorResponse
//	@Failure		503			{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/fis/races/count-by-category [get]
func (h *RaceSearchHandler) GetRaceCategoryCounts(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	if err := utils.ValidateParams(r, []string{
		"seasoncode", "sector", "nationcode", "gender",
	}); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	rawSeason := strings.TrimSpace(r.URL.Query().Get("seasoncode"))
	if rawSeason == "" {
		utils.BadRequestResponse(w, r, fmt.Errorf("missing required query param: seasoncode"))
		return
	}
	season, err := utils.ParsePositiveInt32(rawSeason)
	if err != nil {
		utils.BadRequestResponse(w, r, fmt.Errorf("invalid seasoncode: %s", rawSeason))
		return
	}

	parseList := func(key string) []string {
		vals := r.URL.Query()[key]
		if len(vals) == 1 && strings.Contains(vals[0], ",") {
			return strings.Split(vals[0], ",")
		}
		return vals
	}

	sectorsRaw := parseList("sector")
	if len(sectorsRaw) == 0 {
		utils.BadRequestResponse(w, r, fmt.Errorf("missing required query param: sector"))
		return
	}

	sectorSet := make(map[string]struct{})
	for _, s := range sectorsRaw {
		s = strings.TrimSpace(strings.ToUpper(s))
		if s == "" {
			continue
		}
		switch s {
		case "CC", "JP", "NK":
			sectorSet[s] = struct{}{}
		default:
			utils.BadRequestResponse(w, r, fmt.Errorf("invalid sector: %s", s))
			return
		}
	}
	if len(sectorSet) == 0 {
		utils.BadRequestResponse(w, r, fmt.Errorf("no valid sector values provided"))
		return
	}

	sectors := make([]string, 0, len(sectorSet))
	for _, s := range []string{"CC", "JP", "NK"} {
		if _, ok := sectorSet[s]; ok {
			sectors = append(sectors, s)
		}
	}

	var nationPtr *string
	if nv := strings.TrimSpace(r.URL.Query().Get("nationcode")); nv != "" {
		nv = strings.ToUpper(nv)
		nationPtr = &nv
	}

	var genderPtr *string
	if gv := strings.TrimSpace(r.URL.Query().Get("gender")); gv != "" {
		gv = strings.ToUpper(gv)
		genderPtr = &gv
	}

	type aggItem struct {
		Catcode *string `json:"catcode,omitempty"`
		Total   int64   `json:"total"`
	}

	totals := make(map[string]int64)
	var nullTotal int64
	var hasNull bool

	for _, sector := range sectors {
		switch sector {
		case "CC":
			if h.raceCC == nil {
				utils.InternalServerError(w, r, fmt.Errorf("raceCC store not configured"))
				return
			}
			rows, err := h.raceCC.GetRaceCountsByCategoryCC(r.Context(), season, nationPtr, genderPtr)
			if err != nil {
				utils.InternalServerError(w, r, err)
				return
			}
			for _, row := range rows {
				if !row.Catcode.Valid {
					hasNull = true
					nullTotal += row.Total
					continue
				}
				key := strings.TrimSpace(row.Catcode.String)
				totals[key] += row.Total
			}

		case "JP":
			if h.raceJP == nil {
				utils.InternalServerError(w, r, fmt.Errorf("raceJP store not configured"))
				return
			}
			rows, err := h.raceJP.GetRaceCountsByCategoryJP(r.Context(), season, nationPtr, genderPtr)
			if err != nil {
				utils.InternalServerError(w, r, err)
				return
			}
			for _, row := range rows {
				if !row.Catcode.Valid {
					hasNull = true
					nullTotal += row.Total
					continue
				}
				key := strings.TrimSpace(row.Catcode.String)
				totals[key] += row.Total
			}

		case "NK":
			if h.raceNK == nil {
				utils.InternalServerError(w, r, fmt.Errorf("raceNK store not configured"))
				return
			}
			rows, err := h.raceNK.GetRaceCountsByCategoryNK(r.Context(), season, nationPtr, genderPtr)
			if err != nil {
				utils.InternalServerError(w, r, err)
				return
			}
			for _, row := range rows {
				if !row.Catcode.Valid {
					hasNull = true
					nullTotal += row.Total
					continue
				}
				key := strings.TrimSpace(row.Catcode.String)
				totals[key] += row.Total
			}
		}
	}

	var items []aggItem

	for cat, total := range totals {
		c := cat
		items = append(items, aggItem{
			Catcode: &c,
			Total:   total,
		})
	}

	if hasNull && nullTotal > 0 {
		items = append(items, aggItem{
			Catcode: nil,
			Total:   nullTotal,
		})
	}

	sort.Slice(items, func(i, j int) bool {
		if items[i].Total != items[j].Total {
			return items[i].Total > items[j].Total
		}
		if items[i].Catcode == nil && items[j].Catcode == nil {
			return false
		}
		if items[i].Catcode == nil {
			return false
		}
		if items[j].Catcode == nil {
			return true
		}
		return *items[i].Catcode < *items[j].Catcode
	})

	body := map[string]any{
		"seasoncode": season,
		"sectors":    sectors,
		"nationcode": nationPtr,
		"gender":     genderPtr,
		"categories": items,
	}

	utils.WriteJSON(w, http.StatusOK, body)
}

type RaceNationCountItem struct {
	Nationcode string `json:"nationcode" example:"FIN"`
	Total      int64  `json:"total" example:"42"`
}

// GetRaceCountsByNation godoc
//
//	@Summary		Get race counts by nation
//	@Description	Get counts of races grouped by Nationcodes for a given season and one or more sectors (CC, JP, NK). Optional filters: gender, catcode.
//	@Tags			FIS - KAMK
//	@Accept			json
//	@Produce		json
//	@Param			sector		query		[]string	true	"Sector code(s) (CC,JP,NK – repeat or comma-separated, e.g. sector=CC&sector=JP or sector=CC,JP)"
//	@Param			seasoncode	query		int32		true	"Season code"
//	@Param			gender		query		string		false	"Gender (M/W)"
//	@Param			catcode		query		string		false	"Category code (e.g. WC)"
//	@Success		200			{object}	swagger.FISRacesNationCountsResponse
//	@Failure		400			{object}	swagger.ValidationErrorResponse
//	@Failure		401			{object}	swagger.UnauthorizedResponse
//	@Failure		403			{object}	swagger.ForbiddenResponse
//	@Failure		500			{object}	swagger.InternalServerErrorResponse
//	@Failure		503			{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/fis/races/count-by-nation [get]
func (h *RaceSearchHandler) GetRaceCountsByNation(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	if err := utils.ValidateParams(r, []string{
		"sector", "seasoncode", "gender", "catcode",
	}); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	parseList := func(key string) []string {
		vals := r.URL.Query()[key]
		if len(vals) == 1 && strings.Contains(vals[0], ",") {
			return strings.Split(vals[0], ",")
		}
		return vals
	}

	rawSeason := strings.TrimSpace(r.URL.Query().Get("seasoncode"))
	if rawSeason == "" {
		utils.BadRequestResponse(w, r, fmt.Errorf("missing required query param: seasoncode"))
		return
	}
	season, err := utils.ParsePositiveInt32(rawSeason)
	if err != nil {
		utils.BadRequestResponse(w, r, fmt.Errorf("invalid seasoncode: %s", rawSeason))
		return
	}

	sectorsRaw := parseList("sector")
	if len(sectorsRaw) == 0 {
		utils.BadRequestResponse(w, r, fmt.Errorf("missing required query param: sector"))
		return
	}

	sectorSet := make(map[string]struct{})
	for _, s := range sectorsRaw {
		s = strings.TrimSpace(strings.ToUpper(s))
		if s == "" {
			continue
		}
		switch s {
		case "CC", "JP", "NK":
			sectorSet[s] = struct{}{}
		default:
			utils.BadRequestResponse(w, r, fmt.Errorf("invalid sector: %s", s))
			return
		}
	}
	if len(sectorSet) == 0 {
		utils.BadRequestResponse(w, r, fmt.Errorf("no valid sector values provided"))
		return
	}

	sectors := make([]string, 0, len(sectorSet))
	for _, s := range []string{"CC", "JP", "NK"} {
		if _, ok := sectorSet[s]; ok {
			sectors = append(sectors, s)
		}
	}

	var genderPtr *string
	if gv := strings.TrimSpace(r.URL.Query().Get("gender")); gv != "" {
		gv = strings.ToUpper(gv)
		genderPtr = &gv
	}

	var catPtr *string
	if cv := strings.TrimSpace(r.URL.Query().Get("catcode")); cv != "" {
		cv = strings.ToUpper(cv)
		catPtr = &cv
	}

	counts := make(map[string]int64)

	for _, sector := range sectors {
		switch sector {
		case "CC":
			if h.raceCC == nil {
				utils.InternalServerError(w, r, fmt.Errorf("raceCC store not configured"))
				return
			}
			rows, err := h.raceCC.GetRaceCountsByNationCC(r.Context(), season, catPtr, genderPtr)
			if err != nil {
				utils.InternalServerError(w, r, err)
				return
			}
			for _, row := range rows {
				if !row.Nationcode.Valid || row.Nationcode.String == "" {
					continue
				}
				counts[row.Nationcode.String] += row.Total
			}

		case "JP":
			if h.raceJP == nil {
				utils.InternalServerError(w, r, fmt.Errorf("raceJP store not configured"))
				return
			}
			rows, err := h.raceJP.GetRaceCountsByNationJP(r.Context(), season, catPtr, genderPtr)
			if err != nil {
				utils.InternalServerError(w, r, err)
				return
			}
			for _, row := range rows {
				if !row.Nationcode.Valid || row.Nationcode.String == "" {
					continue
				}
				counts[row.Nationcode.String] += row.Total
			}

		case "NK":
			if h.raceNK == nil {
				utils.InternalServerError(w, r, fmt.Errorf("raceNK store not configured"))
				return
			}
			rows, err := h.raceNK.GetRaceCountsByNationNK(r.Context(), season, catPtr, genderPtr)
			if err != nil {
				utils.InternalServerError(w, r, err)
				return
			}
			for _, row := range rows {
				if !row.Nationcode.Valid || row.Nationcode.String == "" {
					continue
				}
				counts[row.Nationcode.String] += row.Total
			}
		}
	}

	items := make([]RaceNationCountItem, 0, len(counts))
	for code, total := range counts {
		items = append(items, RaceNationCountItem{
			Nationcode: code,
			Total:      total,
		})
	}

	sort.Slice(items, func(i, j int) bool {
		if items[i].Total == items[j].Total {
			return items[i].Nationcode < items[j].Nationcode
		}
		return items[i].Total > items[j].Total
	})

	body := map[string]any{
		"seasoncode": season,
		"sectors":    sectors,
		"gender":     genderPtr,
		"catcode":    catPtr,
		"nations":    items,
	}

	utils.WriteJSON(w, http.StatusOK, body)
}

// GetRaceTotals godoc
//
//	@Summary		Get total race count
//	@Description	Gets total count of races for a given season and one or more sectors (CC, JP, NK). Optional filters: Catcode, Gender.
//	@Tags			FIS - KAMK
//	@Accept			json
//	@Produce		json
//	@Param			seasoncode	query		int32		true	"Season code"
//	@Param			sector		query		[]string	true	"Sector code(s) (CC,JP,NK – repeat or comma-separated, e.g. sector=CC&sector=JP or sector=CC,JP)"
//	@Param			catcode		query		string		false	"Category code filter (e.g. WC)"
//	@Param			gender		query		string		false	"Gender filter (e.g. M/W)"
//	@Success		200			{object}	swagger.FISRacesTotalsResponse
//	@Failure		400			{object}	swagger.ValidationErrorResponse
//	@Failure		401			{object}	swagger.UnauthorizedResponse
//	@Failure		403			{object}	swagger.ForbiddenResponse
//	@Failure		500			{object}	swagger.InternalServerErrorResponse
//	@Failure		503			{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/fis/races/count-total [get]
func (h *RaceSearchHandler) GetRaceTotals(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	if err := utils.ValidateParams(r, []string{
		"seasoncode", "sector", "catcode", "gender",
	}); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	rawSeason := strings.TrimSpace(r.URL.Query().Get("seasoncode"))
	if rawSeason == "" {
		utils.BadRequestResponse(w, r, fmt.Errorf("missing required query param: seasoncode"))
		return
	}
	season, err := utils.ParsePositiveInt32(rawSeason)
	if err != nil {
		utils.BadRequestResponse(w, r, fmt.Errorf("invalid seasoncode: %s", rawSeason))
		return
	}

	parseList := func(key string) []string {
		vals := r.URL.Query()[key]
		if len(vals) == 1 && strings.Contains(vals[0], ",") {
			return strings.Split(vals[0], ",")
		}
		return vals
	}

	sectorsRaw := parseList("sector")
	if len(sectorsRaw) == 0 {
		utils.BadRequestResponse(w, r, fmt.Errorf("missing required query param: sector"))
		return
	}

	sectorSet := make(map[string]struct{})
	for _, s := range sectorsRaw {
		s = strings.TrimSpace(strings.ToUpper(s))
		if s == "" {
			continue
		}
		switch s {
		case "CC", "JP", "NK":
			sectorSet[s] = struct{}{}
		default:
			utils.BadRequestResponse(w, r, fmt.Errorf("invalid sector: %s", s))
			return
		}
	}
	if len(sectorSet) == 0 {
		utils.BadRequestResponse(w, r, fmt.Errorf("no valid sector values provided"))
		return
	}

	sectors := make([]string, 0, len(sectorSet))
	for _, s := range []string{"CC", "JP", "NK"} {
		if _, ok := sectorSet[s]; ok {
			sectors = append(sectors, s)
		}
	}

	var catPtr *string
	if cv := strings.TrimSpace(r.URL.Query().Get("catcode")); cv != "" {
		cv = strings.ToUpper(cv)
		catPtr = &cv
	}

	var genderPtr *string
	if gv := strings.TrimSpace(r.URL.Query().Get("gender")); gv != "" {
		gv = strings.ToUpper(gv)
		genderPtr = &gv
	}

	var total int64

	for _, sector := range sectors {
		switch sector {
		case "CC":
			if h.raceCC == nil {
				utils.InternalServerError(w, r, fmt.Errorf("raceCC store not configured"))
				return
			}
			t, err := h.raceCC.GetRaceTotalCC(r.Context(), season, catPtr, genderPtr)
			if err != nil {
				utils.InternalServerError(w, r, err)
				return
			}
			total += t

		case "JP":
			if h.raceJP == nil {
				utils.InternalServerError(w, r, fmt.Errorf("raceJP store not configured"))
				return
			}
			t, err := h.raceJP.GetRaceTotalJP(r.Context(), season, catPtr, genderPtr)
			if err != nil {
				utils.InternalServerError(w, r, err)
				return
			}
			total += t

		case "NK":
			if h.raceNK == nil {
				utils.InternalServerError(w, r, fmt.Errorf("raceNK store not configured"))
				return
			}
			t, err := h.raceNK.GetRaceTotalNK(r.Context(), season, catPtr, genderPtr)
			if err != nil {
				utils.InternalServerError(w, r, err)
				return
			}
			total += t
		}
	}

	body := map[string]any{
		"seasoncode": season,
		"sectors":    sectors,
		"catcode":    catPtr,
		"gender":     genderPtr,
		"total":      total,
	}

	utils.WriteJSON(w, http.StatusOK, body)
}
