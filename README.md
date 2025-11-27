# Full-Stack Web App with Go, Svelte, and Docker
This is the final project for the SysAdmin course.  
Our web application is a Mood Tracker, where a user can document whatever mood they are feeling.  
It is a minimal, full-stack web application featuring a Go backend API and a SvelteKit frontend. The entire application is containerized using Docker and orchestrated with Docker Compose for easy, one-command setup and deployment.
## **Key Features**
  - User Registration: New users can create an account.
  - Email Activation: Upon registration, users receive an email with an activation token to verify their account.
  - Token-Based Authentication: Users can log in to receive a JWT for accessing protected resources.
  - Protected Routes: A sample dashboard page is only accessible to authenticated and activated users.
  - Fully Containerized: Both the Go backend and SvelteKit frontend run in isolated Docker containers.
  - One-Command Launch: The entire application stack is managed by a single `docker-compose.yml` file.
## **Technology Stack**  
**Backend:**	Go, httprouter, PostgreSQL, Mailtrap (for SMTP)  
**Frontend:**	SvelteKit, Vite  
**DevOps:**	Docker, Docker Compose  

## Getting Started: The Docker Way (Recommended)
This is the simplest and most reliable way to run the entire application. It ensures the environment is identical for everyone.
### **Prerequisites**
  - Docker installed and running.
  - A running PostgreSQL database accessible from your host machine.
  - A Mailtrap.io account for SMTP testing.

**Step 1: Clone the Repository**
```Bash
git clone <your-repo-url>
cd <your-repo-folder>
```
**Step 2: Configure Your Environment**  
The application requires credentials for the email service (Mailtrap) and your database credentials. This file should be placed in the `Backend folder`. A template file is provided.
1. Rename the example file:
```Bash
# For Windows (Command Prompt)
rename example.envrc .envrc

# For Linux/macOS/Git Bash
mv example.env .env
```
2. Edit the .env file and fill in your actual Mailtrap SMTP username and password; the same thing goes for your database.

**Step 3: Launch the Application**  
Run the following command from the root directory of the project (where the `docker-compose.yml` file is located).
```Bash
docker-compose up --build
```

  - `--build`: This flag tells Docker Compose to build the images for the frontend and backend before starting them. You only need this flag the first time or whenever you make changes to the source code or `Dockerfile`s.

Docker will now build both images and start the containers. You will see the logs from both the backend and frontend in your terminal.  

**Step 4: Access the Application**  

  - Frontend Web App: Open your browser and navigate to http://localhost:5173
  - Backend API: The API is accessible at http://localhost:4000, which the frontend is pre-configured to communicate with.
### How to Stop the Application
To stop all running containers, press `Ctrl + C` in the terminal where Docker Compose is running. To clean up and remove the containers and network, run:
```Bash
docker-compose down
```
# Manual Setup (Without Docker)
This method is for developing and testing the services individually without containerization.
## **Prerequisites**
  - Go (version 1.25+ recommended)
  - Node.js (version 20+ recommended)
  - A running PostgreSQL database.
  - A Mailtrap account.
## **Backend Setup**
1. Set Environment Variables: The Go application is configured via command-line flags. You can set these directly or use a `Makefile`. Ensure your `SYSADMIN_DB_DSN` variable points to your PostgreSQL database.
2. Navigate to the backend folder: `cd backend`
3. Install dependencies: `go mod tidy`
4. Run the server:
```Bash
go run ./cmd/api \
    -db-dsn="postgres://sysadmin:fishsticks@localhost:5432/sysadmin?sslmode=disable" \
    -smtp-host="smtp.mailtrap.io" \
    -smtp-port=2525 \
    -smtp-username="YOUR_USERNAME" \
    -smtp-password="YOUR_PASSWORD" \
    -smtp-sender="Your App <no-reply@yourapp.com>" \
    -cors-trusted-origins="http://localhost:5173"
```
## **Frontend Setup**
1. Navigate to the frontend folder: `cd frontend`
2. Install dependencies: `npm install`
3. Run the development server:
```Bash
npm run dev
```
4. The frontend will now be available at `http://localhost:5173`.
