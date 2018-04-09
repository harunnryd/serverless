package serverless

/**
 * ----------------------
 * NOTE : response body
 * ----------------------
 */
type Response struct {

	/**
	 * ------------------------
	 * NOTE : response #1
	 * developer
	 * ------------------------
	 */
	IncomingWebhook struct {
		Endpoint      string `json:"end_point"`
		Authorization string `json:"authorization"`
		Method        string `json:"method"`
	} `json:"incoming_webhook"`

	/**
	 * -----------------------------
	 * NOTE : response #2
	 * https://cloud.openalpr.com
	 * -----------------------------
	 */
	ALPR struct {
		UUID           string `json:"uuid"`
		DataType       string `json:"data_type"`
		EpochTime      int64  `json:"epoch_time"`
		ProcessingTime struct {
			Plates float64 `json:"plates"`
			Total  float64 `json:"total"`
		} `json:"processing_time"`
		ImgHeight int `json:"img_height"`
		ImgWidth  int `json:"img_width"`
		Results   []struct {
			Plate            string  `json:"plate"`
			Confidence       float64 `json:"confidence"`
			RegionConfidence int     `json:"region_confidence"`
			VehicleRegion    struct {
				Y      int `json:"y"`
				X      int `json:"x"`
				Height int `json:"height"`
				Width  int `json:"width"`
			} `json:"vehicle_region"`
			Region           string  `json:"region"`
			PlateIndex       int     `json:"plate_index"`
			ProcessingTimeMs float64 `json:"processing_time_ms"`
			Candidates       []struct {
				MatchesTemplate int     `json:"matches_template"`
				Plate           string  `json:"plate"`
				Confidence      float64 `json:"confidence"`
			} `json:"candidates"`
			Coordinates []struct {
				Y int `json:"y"`
				X int `json:"x"`
			} `json:"coordinates"`
			MatchesTemplate int `json:"matches_template"`
			RequestedTopn   int `json:"requested_topn"`
		} `json:"results"`
		CreditsMonthlyUsed  int  `json:"credits_monthly_used"`
		Version             int  `json:"version"`
		CreditsMonthlyTotal int  `json:"credits_monthly_total"`
		Error               bool `json:"error"`
		RegionsOfInterest   []struct {
			Y      int `json:"y"`
			X      int `json:"x"`
			Height int `json:"height"`
			Width  int `json:"width"`
		} `json:"regions_of_interest"`
		CreditCost int `json:"credit_cost"`
	}
}
