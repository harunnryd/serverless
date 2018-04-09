package serverless

/**
 * --------------------------
 * NOTE : payload request
 * --------------------------
 */
type Request struct {

	/**
	 * -----------------------
	 * NOTE : payload #1
	 * developer
	 * -----------------------
	 */
	IncomingWebhook struct {
		Endpoint      string `json:"end_point"`
		Authorization string `json:"authorization"`
		Method        string `json:"method"`
	} `json:"incoming_webhook"`

	/**
	 * -----------------------------
	 * NOTE : payload #2
	 * https://cloud.openalpr.com
	 * -----------------------------
	 */
	ALPR struct {
		ImageURL  string `json:"image_url"`
		SecretKey string `json:"secret_key"`
		Type      string `json:"type"`
	} `json:"alpr"`
}
