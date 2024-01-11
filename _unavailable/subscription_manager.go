package rozetkapay

import (
	"net/http"
)

// Create subscription
func (c *Client) CreateSubscription(schema *CreateSubscriptionSchema) (
	*CreateSubscriptionResponse, error,
) {
	req, err := c.NewRequest(
		http.MethodPost,
		c.Config.APIURL+"subscriptions/v1/subscriptions",
		schema,
		nil,
	)
	if err != nil {
		return nil, err
	}
	resp := &CreateSubscriptionResponse{}
	if err := c.Send(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Deactivate specified subscription
func (c *Client) DeactivateSubscription(subscriptionID string) (
	*SingleSubscriptionMessageResponse, error,
) {
	req, err := c.NewRequest(
		http.MethodDelete,
		c.Config.APIURL+"subscriptions/v1/subscriptions/"+subscriptionID,
		nil,
		nil,
	)
	if err != nil {
		return nil, err
	}
	resp := &SingleSubscriptionMessageResponse{}
	if err := c.Send(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Retrieve subscription details
func (c *Client) GetSubscriptionInfo(subscriptionID string) (
	*GetSubscriptionInfoResponse, error,
) {
	req, err := c.NewRequest(
		http.MethodGet,
		c.Config.APIURL+"subscriptions/v1/subscriptions/"+subscriptionID,
		nil,
		nil,
	)
	if err != nil {
		return nil, err
	}
	resp := &GetSubscriptionInfoResponse{}
	if err := c.Send(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Update specific subscription by id
func (c *Client) UpdateSubscription(subscriptionID string, autoRenew bool) (
	*SingleSubscriptionMessageResponse, error,
) {
	type Schema struct {
		AutoRenew bool `json:"auto_renew"`
	}
	req, err := c.NewRequest(
		http.MethodPatch,
		c.Config.APIURL+"subscriptions/v1/subscriptions/"+subscriptionID,
		&Schema{AutoRenew: autoRenew},
		nil,
	)
	if err != nil {
		return nil, err
	}
	resp := &SingleSubscriptionMessageResponse{}
	if err := c.Send(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Get subscriptions for specific customer.
func (c *Client) GetCustomerSubscriptions(customerID string) (
	[]GetSubscriptionInfoResponse, error,
) {
	req, err := c.NewRequest(
		http.MethodGet,
		c.Config.APIURL+"subscriptions/v1/customer-subscriptions",
		nil,
		map[string]string{"customer_id": customerID},
	)
	if err != nil {
		return nil, err
	}
	resp := []GetSubscriptionInfoResponse{}
	if err := c.Send(req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Get all plans for the platform
func (c *Client) CreateSubscriptionPlan(schema *CreateSubscriptionPlanSchema) (
	*CreateSubscriptionPlanResponse, error,
) {
	err := c.Validator.Struct(schema)
	if err != nil {
		return nil, err
	}
	req, err := c.NewRequest(
		http.MethodPost,
		c.Config.APIURL+"subscriptions/v1/plans",
		schema,
		nil,
	)
	if err != nil {
		return nil, err
	}
	resp := &CreateSubscriptionPlanResponse{}
	if err := c.Send(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Deactivate specific plan
func (c *Client) DeactivateSubscriptionPlan(planID string) (
	*SingleSubscriptionMessageResponse, error,
) {
	req, err := c.NewRequest(
		http.MethodDelete,
		c.Config.APIURL+"subscriptions/v1/plans/"+planID,
		nil,
		nil,
	)
	if err != nil {
		return nil, err
	}
	resp := &SingleSubscriptionMessageResponse{}
	if err := c.Send(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Get specific plan by id
func (c *Client) GetSubscriptionPlan(planID string) (
	*GetSubscriptionPlanResponse, error,
) {
	req, err := c.NewRequest(
		http.MethodGet,
		c.Config.APIURL+"subscriptions/v1/plans/"+planID,
		nil,
		nil,
	)
	if err != nil {
		return nil, err
	}
	resp := &GetSubscriptionPlanResponse{}
	if err := c.Send(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Get all plans for the platform
func (c *Client) GetSubscriptionPlans(platform string) (
	[]GetSubscriptionPlanResponse, error,
) {
	req, err := c.NewRequest(
		http.MethodGet,
		c.Config.APIURL+"subscriptions/v1/plans",
		nil,
		map[string]string{"platform": platform},
	)
	if err != nil {
		return nil, err
	}
	resp := []GetSubscriptionPlanResponse{}
	if err := c.Send(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Update specific plan by id
func (c *Client) UpdateSubscriptionPlan(planID string, schema *UpdateSubscriptionPlanSchema) (
	*SingleSubscriptionMessageResponse, error,
) {
	req, err := c.NewRequest(
		http.MethodPatch,
		c.Config.APIURL+"subscriptions/v1/plans/"+planID,
		schema,
		nil,
	)
	if err != nil {
		return nil, err
	}
	resp := &SingleSubscriptionMessageResponse{}
	if err := c.Send(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Returns all payments for the specified subscription.
func (c *Client) GetSubscriptionPayments(subscriptionID string) (
	*GetSubscriptionPaymentsResponse, error,
) {
	req, err := c.NewRequest(
		http.MethodGet,
		c.Config.APIURL+"subscriptions/v1/subscriptions/"+subscriptionID+"/payments",
		nil,
		nil,
	)
	if err != nil {
		return nil, err
	}
	resp := &GetSubscriptionPaymentsResponse{}
	if err := c.Send(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
