package rozetkapay

type SingleSubscriptionMessageResponse struct {
	Message string `json:"message"`
}

// Get subscription info
type (
	GetSubscriptionInfoResponse struct {
		AutoRenew            bool   `json:"auto_renew"`
		CallbackURL          string `json:"callback_url"`
		CreatedAt            string `json:"created_at"`
		Currency             string `json:"currency"`
		Description          string `json:"description"`
		ID                   string `json:"id"`
		NextNotificationDate string `json:"next_notification_date"`
		NextPaymentDate      string `json:"next_payment_date"`
		PlanID               string `json:"plan_id"`
		Price                int    `json:"price"`
		ResultURL            string `json:"result_url"`
		StartDate            string `json:"start_date"`
		State                State  `json:"state"`
		UpdatedAt            string `json:"updated_at"`
	}
)

// Create subscription
type (
	CreateSubscriptionSchema struct {
		AutoRenew   bool      `json:"auto_renew"`
		CallbackURL string    `json:"callback_url" validate:"required"`
		Customer    Recipient `json:"customer" validate:"required"`
		Description string    `json:"description"`
		PlanID      string    `json:"plan_id" validate:"required"`
		Price       int       `json:"price"`
		ResultURL   string    `json:"result_url" validate:"required"`
		StartDate   string    `json:"start_date" validate:"required"`
	}

	PaymentDetails struct {
		Amount             int    `json:"amount"`
		CreatedAt          string `json:"created_at"`
		Currency           string `json:"currency"`
		Description        string `json:"description"`
		NextProcessingDate string `json:"next_processing_date"`
		ProcessedAt        string `json:"processed_at"`
		RetryCount         int    `json:"retry_count"`
		Status             string `json:"status"`
		StatusCode         string `json:"status_code"`
		StatusDescription  string `json:"status_description"`
		UpdatedAt          string `json:"updated_at"`
	}

	Payment struct {
		Details        PaymentDetails    `json:"details"`
		ID             string            `json:"id"`
		SubscriptionID string            `json:"subscription_id"`
		UserAction     PaymentUserAction `json:"user_action"`
	}

	Subscription struct {
		AutoRenew            bool   `json:"auto_renew"`
		CallbackURL          string `json:"callback_url"`
		CreatedAt            string `json:"created_at"`
		Currency             string `json:"currency"`
		Description          string `json:"description"`
		ID                   string `json:"id"`
		NextNotificationDate string `json:"next_notification_date"`
		NextPaymentDate      string `json:"next_payment_date"`
		PlanID               string `json:"plan_id"`
		Price                int    `json:"price"`
		ResultURL            string `json:"result_url"`
		StartDate            string `json:"start_date"`
		State                string `json:"state"`
		UpdatedAt            string `json:"updated_at"`
	}

	CreateSubscriptionResponse struct {
		Payment      Payment      `json:"payment"`
		Subscription Subscription `json:"subscription"`
	}
)

// Create subscription plan

type FrequencyType string

const (
	FrequencyTypeDaily   FrequencyType = "daily"
	FrequencyTypeWeekly  FrequencyType = "weekly"
	FrequencyTypeMonthly FrequencyType = "monthly"
	FrequencyTypeYearly  FrequencyType = "yearly"
)

type State string

const (
	StateActive   State = "active"
	StateInActive State = "inactive"
)

type (
	CreateSubscriptionPlanSchema struct {
		Name          string        `json:"name" validate:"required"`
		Description   string        `json:"description"`
		Price         int           `json:"price" validate:"required"`
		Currency      string        `json:"currency" validate:"required"`
		Platforms     []string      `json:"platforms" validate:"required"`
		FrequencyType FrequencyType `json:"frequency_type" validate:"required"`

		// frequency_type: daily, frequency: 2 - subscription debits occur every two days;
		// frequency_type: monthly, frequency: 1 - subscription payments are made every month;
		// frequency_type: monthly, frequency: 6 - subscription is charged every six months.
		Frequency int `json:"frequency" validate:"required"`

		// The start_date controls when the plan will become active:
		// if this date is the same as the date of plan creation, it becomes active immediately, if later, it becomes active on the specified date.
		// Only dates are taken, time is ignored.
		StartDate string `json:"start_date" validate:"required"`

		// If you pass an end_date, the plan is not eternal, but expires automatically on the specified date.
		// It will be impossible to create new subscriptions for this plan, but existing subscriptions will remain active even after the plan's end_date.
		// If you want to deactivate all the subscriptions of a plan that has already reached the end_date, you need to call https://cdn.evopay.com.ua/public-docs/index.html#tag/subscriptions/operation/deactivatePlan.
		// Only dates are taken, time is ignored.
		EndDate string `json:"end_date"`
	}

	CreateSubscriptionPlanResponse struct {
		ID            string        `json:"id"`
		CreatedAt     string        `json:"created_at"`
		Currency      string        `json:"currency"`
		Description   string        `json:"description"`
		EndDate       string        `json:"end_date"`
		FrequencyType FrequencyType `json:"frequency_type"`
		Frequency     int           `json:"frequency"`
		Name          string        `js\son:"name"`
		Platforms     []string      `json:"platforms"`
		Price         int           `json:"price"`
		StartDate     string        `json:"start_date"`
		UpdatedAt     string        `json:"updated_at"`
		State         State         `json:"state"`
	}
)

// Get subscription plan response
type (
	GetSubscriptionPlanResponse struct {
		CreatedAt     string   `json:"created_at"`
		Currency      string   `json:"currency"`
		Description   string   `json:"description"`
		EndDate       string   `json:"end_date"`
		Frequency     int      `json:"frequency"`
		FrequencyType string   `json:"frequency_type"`
		ID            string   `json:"id"`
		State         string   `json:"state"`
		Name          string   `json:"name"`
		Platforms     []string `json:"platforms"`
		Price         int      `json:"price"`
		StartDate     string   `json:"start_date"`
		UpdatedAt     string   `json:"updated_at"`
	}
)

// Get subscription payments
type (
	GetSubscriptionPaymentDetails struct {
		Amount             string `json:"amount"`
		CreatedAt          string `json:"created_at"`
		Currency           string `json:"currency"`
		Description        string `json:"description"`
		NextProcessingDate string `json:"next_processing_date"`
		ProcessedAt        string `json:"processed_at"`
		RetryCount         int    `json:"retry_count"`
		Status             string `json:"status"`
		StatusCode         string `json:"status_code"`
		StatusDescription  string `json:"status_description"`
		UpdatedAt          string `json:"updated_at"`
	}

	GetSubscriptionPaymentsResponse struct {
		ID             string                        `json:"id"`
		SubscriptionID string                        `json:"subscription_id"`
		Details        GetSubscriptionPaymentDetails `json:"details"`
		UserAction     PaymentUserAction             `json:"user_action"`
	}
)

// Update subscription plan
type (
	UpdateSubscriptionPlanSchema struct {
		Name        string `json:"name,omitempty"`
		Description string `json:"description,omitempty"`
	}
)
