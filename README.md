# Fiscal Microservices Project

## Overview

This project is a collection of microservices built with **Go** and the **Fiber** web framework, designed to handle various fiscal operations (e.g., tax calculations, invoicing, financial reporting). The services are containerized using **Docker** and leverage **PostgreSQL** as the relational database. The architecture follows microservices principles, allowing each service to be deployed, scaled, and managed independently.

## Technologies Used
- **Go**: Backend language for building the microservices, known for performance and simplicity.
- **Fiber**: A high-performance web framework for Go, inspired by Express.js.
- **Docker**: Used for containerization to ensure that each microservice can be isolated, deployed, and managed in different environments.
- **PostgreSQL**: The relational database for persistent data storage.
- **Docker Compose**: For managing multi-container Docker applications (used in development).
- **Swagger**: For API documentation.

## Services

### 1. Tax Calculation Service
- Handles all tax-related computations based on various fiscal rules.
- Exposes RESTful endpoints for clients to retrieve tax details.

### 2. Invoicing Service
- Manages invoice generation, including handling different fiscal policies for invoicing.
- Stores invoice data and provides APIs to generate and retrieve invoices.

### 3. Financial Reporting Service
- Aggregates financial data for reporting purposes.
- Provides APIs for generating financial reports and analytics.

## Project Structure

```bash
.
├── tax-service/
│   ├── main.go
│   ├── Dockerfile
│   ├── go.mod
│   └── ...
├── invoicing-service/
│   ├── main.go
│   ├── Dockerfile
│   ├── go.mod
│   └── ...
├── reporting-service/
│   ├── main.go
│   ├── Dockerfile
│   ├── go.mod
│   └── ...
├── docker-compose.yml
├── README.md
└── .env
