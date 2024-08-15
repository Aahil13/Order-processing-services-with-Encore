# Microservices with encore.dev

This project is a straightforward demonstration of a microservices architecture using [encore.dev](https://encore.dev).

It features two services: **Service A** (Order Service) and **Service B** (Notification Service). These services communicate easily to process orders and send notifications.

## Prerequisites

To use this project, make sure you have the following:

- **Go**: Ensure you have Go installed. You can download it [here](https://golang.org/dl/).
- **encore.dev CLI**: Install the Encore CLI by running:

  ```bash
  brew install encoredev/tap/encore
  ```

  The command above works for macos, but you can check out the [Encore documentation](https://encore.dev/docs/install) for one that matches your OS.

## Getting Started

Follow the steps below to get started with this project:

1. **Clone the repository**:

   ```bash
   git clone ttps://github.com/Aahil13/Order-processing-services-with-Encore.git
   cd Order-processing-services-with-Encore
   ```

2. **Run the services**: Start the Encore development server:

   ```bash
   encore run
   ```

   This will spin up your two services locally.

3. **Test the endpoints**:

   - You can place an order by sending a POST request to Service A:

     ```bash
     curl -X POST http://localhost:4000/serviceA/placeorder -d '{"ProductID": "123", "Quantity": 2}'
     ```

        OR

   - Trigger a notification by calling Service B, which internally calls Service A:

     ```bash
     curl -X POST http://localhost:4000/serviceB/sendnotification -d '{"ProductID": "123", "Quantity": 2}'
     ```

**NOTE**: Instead of doing this via your commnand line, you can navigate to `localhost:4000` to access a user friendly dashboard.

From here, you can make these API calls and also observe tracing on them.