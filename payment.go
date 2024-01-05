package rozetkapay

import "time"

type PaymentMode string

const (
	PaymentModeDirect PaymentMode = "direct"
	PaymentModeHosted PaymentMode = "hosted"
)

type PaymentStatus string

const (
	// The transaction has been created but the processing has not yet started.
	PaymentStatusInit PaymentStatus = "init"

	// Transaction in processing.
	PaymentStatusPending PaymentStatus = "pending"

	// The transaction is successful.
	PaymentStatusSuccess PaymentStatus = "success"

	// Transaction is not successful
	PaymentStatusFailure PaymentStatus = "failure"
)

// PaymentStatusCode represents the custom string type for error codes.
type PaymentStatusCode string

const (
	StatusCodeAuthorizationFailed                               PaymentStatusCode = "authorization_failed"
	StatusCodeCustomerAuthNotFound                              PaymentStatusCode = "customer_auth_not_found"
	StatusCodeRequestFailed                                     PaymentStatusCode = "request_failed"
	StatusCodeInternalError                                     PaymentStatusCode = "internal_error"
	StatusCodeAccessNotAllowed                                  PaymentStatusCode = "access_not_allowed"
	StatusCodeInvalidRequestBody                                PaymentStatusCode = "invalid_request_body"
	StatusCodePaymentSettingsNotFound                           PaymentStatusCode = "payment_settings_not_found"
	StatusCodeTransactionAlreadyPaid                            PaymentStatusCode = "transaction_already_paid"
	StatusCodeActionNotAllowed                                  PaymentStatusCode = "action_not_allowed"
	StatusCodeActionAlreadyDone                                 PaymentStatusCode = "action_already_done"
	StatusCodeTransactionSuccessPrimaryNotFound                 PaymentStatusCode = "transaction_success_primary_not_found"
	StatusCodePaymentMethodNotAllowed                           PaymentStatusCode = "payment_method_not_allowed"
	StatusCodeWalletNotConfigured                               PaymentStatusCode = "wallet_not_configured"
	StatusCodePaymentMethodAlreadyConfirmed                     PaymentStatusCode = "payment_method_already_confirmed"
	StatusCodePaymentMethodNotFound                             PaymentStatusCode = "payment_method_not_found"
	StatusCodeInvalidCardToken                                  PaymentStatusCode = "invalid_card_token"
	StatusCodeCustomerAuthTokenExpiredOrInvalid                 PaymentStatusCode = "customer_auth_token_expired_or_invalid"
	StatusCodeCustomerProfileNotFound                           PaymentStatusCode = "customer_profile_not_found"
	StatusCodeCustomerIDNotPassed                               PaymentStatusCode = "customer_id_not_passed"
	StatusCodeTransactionNotFound                               PaymentStatusCode = "transaction_not_found"
	StatusCodeWaitingForVerification                            PaymentStatusCode = "waiting_for_verification"
	StatusCodeTransactionAmountLimit                            PaymentStatusCode = "transaction_amount_limit"
	StatusCodeInvalidData                                       PaymentStatusCode = "invalid_data"
	StatusCodeTransactionDeclined                               PaymentStatusCode = "transaction_declined"
	StatusCodeAuthorizationError                                PaymentStatusCode = "authorization_error"
	StatusCodeTransactionRejected                               PaymentStatusCode = "transaction_rejected"
	StatusCodeTransactionSuccessful                             PaymentStatusCode = "transaction_successful"
	StatusCodeAntiFraudCheck                                    PaymentStatusCode = "anti_fraud_check"
	StatusCodeCardNotSupported                                  PaymentStatusCode = "card_not_supported"
	StatusCodeConfirmationTimeout                               PaymentStatusCode = "confirmation_timeout"
	StatusCodeInvalidCardData                                   PaymentStatusCode = "invalid_card_data"
	StatusCodeInvalidCurrency                                   PaymentStatusCode = "invalid_currency"
	StatusCodePending                                           PaymentStatusCode = "pending"
	StatusCodeWaitingForComplete                                PaymentStatusCode = "waiting_for_complete"
	StatusCodeAccessError                                       PaymentStatusCode = "access_error"
	StatusCodeCardExpired                                       PaymentStatusCode = "card_expired"
	StatusCodeReceiverInfoError                                 PaymentStatusCode = "receiver_info_error"
	StatusCodeTransactionLimitExceeded                          PaymentStatusCode = "transaction_limit_exceeded"
	StatusCodeTransactionNotSupported                           PaymentStatusCode = "transaction_not_supported"
	StatusCodeThreeDSNotSupported                               PaymentStatusCode = "3ds_not_supported"
	StatusCodeThreeDSRequired                                   PaymentStatusCode = "3ds_required"
	StatusCodeFailedToCreateTransaction                         PaymentStatusCode = "failed_to_create_transaction"
	StatusCodeFailedToFinishTransaction                         PaymentStatusCode = "failed_to_finish_transaction"
	StatusCodeInsufficientFunds                                 PaymentStatusCode = "insufficient_funds"
	StatusCodeInvalidPhoneNumber                                PaymentStatusCode = "invalid_phone_number"
	StatusCodeCardHasConstraints                                PaymentStatusCode = "card_has_constraints"
	StatusCodePINTRIESExceeded                                  PaymentStatusCode = "pin_tries_exceeded"
	StatusCodeSessionExpired                                    PaymentStatusCode = "session_expired"
	StatusCodeTimeout                                           PaymentStatusCode = "timeout"
	StatusCodeTransactionCreated                                PaymentStatusCode = "transaction_created"
	StatusCodeWaitingForRedirect                                PaymentStatusCode = "waiting_for_redirect"
	StatusCodeWrongAmount                                       PaymentStatusCode = "wrong_amount"
	StatusCodeTestTransaction                                   PaymentStatusCode = "test_transaction"
	StatusCodeSubscriptionSuccessful                            PaymentStatusCode = "subscription_successful"
	StatusCodeUnsubscribedSuccessfully                          PaymentStatusCode = "unsubscribed_successfully"
	StatusCodeWrongPIN                                          PaymentStatusCode = "wrong_pin"
	StatusCodeWrongAuthorizationCode                            PaymentStatusCode = "wrong_authorization_code"
	StatusCodeWrongCAVV                                         PaymentStatusCode = "wrong_cavv"
	StatusCodeWrongCVV                                          PaymentStatusCode = "wrong_cvv"
	StatusCodeWrongAccountNumber                                PaymentStatusCode = "wrong_account_number"
	StatusCodeConfirmRequired                                   PaymentStatusCode = "confirm_required"
	StatusCodeCVVIsRequired                                     PaymentStatusCode = "cvv_is_required"
	StatusCodeConfirmationRequired                              PaymentStatusCode = "confirmation_required"
	StatusCodeSenderInfoRequired                                PaymentStatusCode = "sender_info_required"
	StatusCodeMissedPayoutMethodData                            PaymentStatusCode = "missed_payout_method_data"
	StatusCodeCardVerificationRequired                          PaymentStatusCode = "card_verification_required"
	StatusCodeIncorrectRefundSumOrCurrency                      PaymentStatusCode = "incorrect_refund_sum_or_currency"
	StatusCodePaymentCardHasInvalidStatus                       PaymentStatusCode = "payment_card_has_invalid_status"
	StatusCodeWrongCardNumber                                   PaymentStatusCode = "wrong_card_number"
	StatusCodeUserNotFound                                      PaymentStatusCode = "user_not_found"
	StatusCodeFailedToSendSMS                                   PaymentStatusCode = "failed_to_send_sms"
	StatusCodeWrongSMSPassword                                  PaymentStatusCode = "wrong_sms_password"
	StatusCodeCardNotFound                                      PaymentStatusCode = "card_not_found"
	StatusCodePaymentSystemNotSupported                         PaymentStatusCode = "payment_system_not_supported"
	StatusCodeCountryNotSupported                               PaymentStatusCode = "country_not_supported"
	StatusCodeNoDiscountFound                                   PaymentStatusCode = "no_discount_found"
	StatusCodeFailedToLoadWallet                                PaymentStatusCode = "failed_to_load_wallet"
	StatusCodeInvalidVerificationCode                           PaymentStatusCode = "invalid_verification_code"
	StatusCodeAdditionalInformationIsPending                    PaymentStatusCode = "additional_information_is_pending"
	StatusCodeTransactionIsNotRecurring                         PaymentStatusCode = "transaction_is_not_recurring"
	StatusCodeConfirmAmountCannotBeMoreThanTheTransactionAmount PaymentStatusCode = "confirm_amount_cannot_be_more_than_the_transaction_amount"
	StatusCodeCardBINNotFound                                   PaymentStatusCode = "card_bin_not_found"
	StatusCodeCurrencyRateNotFound                              PaymentStatusCode = "currency_rate_not_found"
	StatusCodeInvalidRecipientName                              PaymentStatusCode = "invalid_recipient_name"
	StatusCodeDailyCardUsageLimitReached                        PaymentStatusCode = "daily_card_usage_limit_reached"
	StatusCodeInvalidTransactionAmount                          PaymentStatusCode = "invalid_transaction_amount"
	StatusCodeCardTypeIsNotSupported                            PaymentStatusCode = "card_type_is_not_supported"
	StatusCodeStoreIsBlocked                                    PaymentStatusCode = "store_is_blocked"
	StatusCodeStoreIsNotActive                                  PaymentStatusCode = "store_is_not_active"
	StatusCodeTransactionCannotBeProcessed                      PaymentStatusCode = "transaction_cannot_be_processed"
	StatusCodeInvalidTransactionStatus                          PaymentStatusCode = "invalid_transaction_status"
	StatusCodePublicKeyNotFound                                 PaymentStatusCode = "public_key_not_found"
	StatusCodeTerminalNotFound                                  PaymentStatusCode = "terminal_not_found"
	StatusCodeFeeNotFound                                       PaymentStatusCode = "fee_not_found"
	StatusCodeFailedToVerifyCard                                PaymentStatusCode = "failed_to_verify_card"
	StatusCodeInvalidTransactionType                            PaymentStatusCode = "invalid_transaction_type"
	StatusCodeRestrictedIP                                      PaymentStatusCode = "restricted_ip"
	StatusCodeInvalidToken                                      PaymentStatusCode = "invalid_token"
	StatusCodePreauthNotAllowed                                 PaymentStatusCode = "preauth_not_allowed"
	StatusCodeTokenDoesNotExist                                 PaymentStatusCode = "token_does_not_exist"
	StatusCodeReachedTheLimitOfAttemptsForIP                    PaymentStatusCode = "reached_the_limit_of_attempts_for_ip"
	StatusCodeCardBranchIsBlocked                               PaymentStatusCode = "card_branch_is_blocked"
	StatusCodeCardBranchDailyLimitReached                       PaymentStatusCode = "card_branch_daily_limit_reached"
	StatusCodeCompletionLimitReached                            PaymentStatusCode = "completion_limit_reached"
	StatusCodeRecurringTransactionsNotAllowed                   PaymentStatusCode = "recurring_transactions_not_allowed"
	StatusCodeTransactionIsCanceledByPayer                      PaymentStatusCode = "transaction_is_canceled_by_payer"
	StatusCodePaymentWasRefunded                                PaymentStatusCode = "payment_was_refunded"
)

type (
	BrowserFingerprint struct {
		BrowserAcceptHeader   string `json:"browser_accept_header,omitempty"`
		BrowserColorDepth     string `json:"browser_color_depth,omitempty"`
		BrowserIPAddress      string `json:"browser_ip_address,omitempty"`
		BrowserJavaEnabled    string `json:"browser_java_enabled,omitempty"`
		BrowserLanguage       string `json:"browser_language,omitempty"`
		BrowserScreenHeight   string `json:"browser_screen_height,omitempty"`
		BrowserTimeZone       string `json:"browser_time_zone,omitempty"`
		BrowserTimeZoneOffset string `json:"browser_time_zone_offset,omitempty"`
		BrowserUserAgent      string `json:"browser_user_agent,omitempty"`
	}

	PaymentMethodAdditions struct {
		BrowserFingerPrint BrowserFingerprint `json:"browser_fingerprint,omitempty"`
		Token              string             `json:"token" validate:"required"`
		Use3DSFlow         bool               `json:"use_3ds_flow,omitempty"`
	}

	WalletRequestPaymentMethod struct {
		BrowserFingerPrint BrowserFingerprint `json:"browser_fingerprint,omitempty"`
		OptionID           string             `json:"option_id" validate:"required"`
		Use3DSFlow         bool               `json:"use_3ds_flow,omitempty"`
	}

	CCToken struct {
		BrowserFingerprint BrowserFingerprint `json:"browser_fingerprint"`
		Token              string             `json:"token"`
		Use3DSFlow         bool               `json:"use_3ds_flow"`
	}

	ApplePay struct {
		BrowserFingerprint BrowserFingerprint `json:"browser_fingerprint"`
		Token              string             `json:"token"`
		Use3DSFlow         bool               `json:"use_3ds_flow"`
	}

	GooglePay struct {
		BrowserFingerprint BrowserFingerprint `json:"browser_fingerprint"`
		Token              string             `json:"token"`
		Use3DSFlow         bool               `json:"use_3ds_flow"`
	}

	Wallet struct {
		BrowserFingerprint BrowserFingerprint `json:"browser_fingerprint"`
		OptionID           string             `json:"option_id"`
		Use3DSFlow         bool               `json:"use_3ds_flow"`
	}

	Product struct {
		Category    string  `json:"category,omitempty"`
		Currency    string  `json:"currency,omitempty"`
		Description string  `json:"description,omitempty"`
		ID          string  `json:"id,omitempty"`
		Image       string  `json:"image,omitempty"`
		Name        string  `json:"name,omitempty"`
		NetAmount   float64 `json:"net_amount,omitempty"`
		Quantity    string  `json:"quantity,omitempty"`
		URL         string  `json:"url,omitempty"`
		VATAmount   float64 `json:"vat_amount,omitempty"`
	}

	Recipient struct {
		Address       string        `json:"address,omitempty" validate:"lte=50"`
		City          string        `json:"city,omitempty"`
		Country       string        `json:"country,omitempty"`
		Email         string        `json:"email,omitempty"`
		ExternalID    string        `json:"external_id,omitempty"`
		FirstName     string        `json:"first_name,omitempty"`
		LastName      string        `json:"last_name,omitempty"`
		Patronym      string        `json:"patronym,omitempty"`
		PaymentMethod PaymentMethod `json:"payment_method,omitempty"`
		Phone         string        `json:"phone,omitempty"`
		PostalCode    string        `json:"postal_code,omitempty"`
	}

	PaymentMethod struct {
		ApplePay  ApplePay          `json:"apple_pay,omitempty"`
		CCToken   CCToken           `json:"cc_token,omitempty"`
		GooglePay GooglePay         `json:"google_pay,omitempty"`
		Type      PaymentMethodType `json:"type,omitempty"`
		Wallet    Wallet            `json:"wallet,omitempty"`
	}
)

type (
	PaymentResponse struct {
		ID string `json:"id"`

		// Block that will be filled in if the parameters are "action_required": True.
		// Possible types: url.
		// In this case, you need to initiate a redirect to the address specified in the value parameter to complete the operation (3ds verification).
		// The parameter will be at host2host requests (type: direct) payment method.
		Action PaymentResponseAction `json:"action"`

		// Indicator of whether additional steps are required to complete the operation.
		ActionRequired bool `json:"action_required"`

		Details PaymentResponseDetails `json:"details"`

		// Unique order number
		ExternalID    string                  `json:"external_id"`
		IsSuccess     bool                    `json:"is_success"`
		ReceiptURL    string                  `json:"receipt_url"`
		PaymentMethod PaymentMethod           `json:"payment_method"`
		Customer      PaymentResponseCustomer `json:"customer"`
	}

	PaymentResponseAction struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	}

	PaymentResponseDetails struct {
		Amount            string                           `json:"amount"`
		BillingOrderID    string                           `json:"billing_order_id"`
		CreatedAt         time.Time                        `json:"created_at"`
		Currency          string                           `json:"currency"`
		Description       string                           `json:"description"`
		GatewayOrderID    string                           `json:"gateway_order_id"`
		Payload           string                           `json:"payload"`
		PaymentID         string                           `json:"payment_id"`
		ProcessedAt       time.Time                        `json:"processed_at"`
		Properties        PaymentResponseDetailsProperties `json:"properties"`
		RRN               string                           `json:"rrn"`
		Status            PaymentStatus                    `json:"status"`
		StatusCode        PaymentStatusCode                `json:"status_code"`
		StatusDescription string                           `json:"status_description"`
		TransactionID     string                           `json:"transaction_id"`
		AuthCode          string                           `json:"auth_code"`
		Fee               PaymentResponseDetailsFee        `json:"fee"`
		TerminalName      string                           `json:"terminal_name"`
	}

	PaymentResponseDetailsProperties struct {
		Property1 string `json:"property1"`
		Property2 string `json:"property2"`
	}

	PaymentResponseDetailsFee struct {
		Amount   string `json:"amount"`
		Currency string `json:"currency"`
	}

	PaymentResponsePaymentMethod struct {
		CCToken CCToken `json:"cc_token"`
		Type    string  `json:"type"`
	}

	PaymentResponseCustomer struct {
		BrowserUserAgent string `json:"browser_user_agent"`
		Email            string `json:"email"`
		ExternalID       string `json:"external_id"`
		FirstName        string `json:"first_name"`
		IPAddress        string `json:"ip_address"`
		LastName         string `json:"last_name"`
		Patronym         string `json:"patronym"`
		Phone            string `json:"phone"`
	}
)

// Hosted mode usally returns payment url.
// Direct mode usally returns url for second factor (3DS or other).
func (p *PaymentResponse) GetLink() string {
	return p.Action.Value
}
