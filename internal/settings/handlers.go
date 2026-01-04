package settings

import (
	"net/http"

	"github.com/knr1997/quiz-tracker-backend/internal/json"
)

func GetSettings(w http.ResponseWriter, r *http.Request) {
	settings := map[string]any{
		"id": 1,
		"options": map[string]any{
			"deliveryTime": []map[string]string{
				{
					"title":       "Express Delivery",
					"description": "90 min express delivery",
				},
				{
					"title":       "Morning",
					"description": "8.00 AM - 11.00 AM",
				},
				{
					"title":       "Noon",
					"description": "11.00 AM - 2.00 PM",
				},
				{
					"title":       "Afternoon",
					"description": "2.00 PM - 5.00 PM",
				},
				{
					"title":       "Evening",
					"description": "5.00 PM - 8.00 PM",
				},
			},
			"isProductReview":       false,
			"useGoogleMap":          false,
			"enableTerms":           true,
			"isMultiCommissionRate": false,
			"enableCoupons":         true,
			"enableReviewPopup":     true,
		},
	}

	json.Write(w, http.StatusOK, settings)
}
