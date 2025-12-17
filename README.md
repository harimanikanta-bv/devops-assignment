# DevOps Reverse Proxy Project

This project sets up two backend services (Alpha and Beta) with an Nginx reverse proxy using Docker Compose.

 
Setup Instructions
1. Clone the repo:
2. Run with Docker Compose:

How Routing Works
- Nginx listens on port `8081`.
- Requests to `/alpha` go to the Alpha service (`port 5000`).
- Requests to `/beta` go to the Beta service (`port 5001`).

Bonus
- Clean separation of services.
- Simple reverse proxy config using Nginx.
