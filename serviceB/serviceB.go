package serviceB

import (
    "context"
    "encore.app/serviceA" 
)

// NotificationResponse defines the response structure
type NotificationResponse struct {
    NotificationMessage string
}

// SendOrderNotification handles requests to /serviceB/sendnotification.
//encore:api public method=POST path=/serviceB/sendnotification
func SendOrderNotification(ctx context.Context, req *serviceA.OrderRequest) (*NotificationResponse, error) {
    // Call Service A to place an order
    orderRes, err := serviceA.PlaceOrder(ctx, req)
    if err != nil {
        return nil, err
    }

    notificationMessage := "Notification: " + orderRes.Message + " with Order ID: " + orderRes.OrderID

    return &NotificationResponse{NotificationMessage: notificationMessage}, nil
}
