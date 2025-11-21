# Backend Setup Guide for SysAdmin Final Project
This guide provides the necessary steps to build and run the Go backend application using Docker on both Windows and Linux.

### Prerequisites
1. Project Files: You need the `Backend` folder from the project.
2. PostgreSQL Database: A running PostgreSQL server that is accessible from your machine. The application will need:
    - A database named `sysadmin`.
    - A user named `sysadmin` with the password `fishsticks`.
3. Mailtrap Credentials: You will need your own Mailtrap SMTP **Username** and **Password**.

## Windows Setup Instructions (Command Prompt or PowerShell)
Follow these steps in a terminal on your Windows machine.

- Step 1: Navigate to the Backend Directory  
Open your terminal and change the directory to the project's `Backend` folder.

```Cmd
cd path\to\your\SysAdmin-FinalProject\Backend
```
- Step 2: Build the Docker Image  
This command packages the Go application into a runnable image named `sysadmin-backend`.

```Cmd
docker build -t sysadmin-backend .
```
- Step 3: Run the Docker Container  
This command starts the container. It connects your local port `4000` to the container's port `4000` and passes all the necessary database and email configurations to the application.  

**Important:** Replace `YOUR_MAILTRAP_USERNAME` and `YOUR_MAILTRAP_PASSWORD` with your actual Mailtrap credentials.

```Cmd
docker run -d -p 4000:4000 --name my-backend-container sysadmin-backend ./server -db-dsn="postgres://sysadmin:fishsticks@host.docker.internal:5432/sysadmin?sslmode=disable" -smtp-host="smtp.mailtrap.io" -smtp-port=2525 -smtp-username="YOUR_MAILTRAP_USERNAME" -smtp-password="YOUR_MAILTRAP_PASSWORD" -smtp-sender="SysAdmin App <no-reply@sysadmin.com>"
```

## Linux Setup Instructions (Bash Terminal)
Follow these steps in a terminal on your Linux machine.  
- Step 1: Navigate to the Backend Directory  
Open your terminal and change the directory to the project's `Backend` folder.

```Bash
cd /path/to/your/SysAdmin-FinalProject/Backend
```

- Step 2: Build the Docker Image  
This command is identical to the Windows command. It packages the Go application into a runnable image named `sysadmin-backend`.  

```Bash
docker build -t sysadmin-backend .
```

- Step 3: Run the Docker Container  
This command is nearly identical to the Windows version, but includes the `--add-host` flag. This flag is necessary on Linux to allow the container to connect to services running on the host machine (like your database) via the name `host.docker.internal`.  

**Important:** Replace `YOUR_MAILTRAP_USERNAME` and `YOUR_MAILTRAP_PASSWORD` with your actual Mailtrap credentials.

```Bash
docker run -d -p 4000:4000 --name my-backend-container --add-host=host.docker.internal:host-gateway sysadmin-backend ./server -db-dsn="postgres://sysadmin:fishsticks@host.docker.internal:5432/sysadmin?sslmode=disable" -smtp-host="smtp.mailtrap.io" -smtp-port=2525 -smtp-username="YOUR_MAILTRAP_USERNAME" -smtp-password="YOUR_MAILTRAP_PASSWORD" -smtp-sender="SysAdmin App <no-reply@sysadmin.com>"
```

### Testing and Troubleshooting (For Both Platforms)
Once the container is running, use these commands to test and debug.
1. Test User Creation with `curl`  
Run this command in your terminal to create a new user.

```Bash
curl -i -X POST -H "Content-Type: application/json" -d '{"name": "Test User", "email": "test-from-docker@example.com", "password": "password123"}' http://localhost:4000/v1/users
```

*   **Expected Success:** You should see an `HTTP/1.1 201 Created` or `HTTP/1.1 200 OK` response. Then, check your Mailtrap inbox for the activation email.
*   **Expected Failure:** If you see `curl: (7) Failed to connect...`, it means your container is not running.

2. How to Check if the Container is Running

```bash
docker ps
```
If the container started successfully, you will see `my-backend-container` in the list with a status of Up.  

3. How to Check Logs if the Container Fails to Start  
If `docker ps` shows nothing, the container likely crashed. Check its logs to see the error message (like a bad password or firewall issue).

```Bash
docker logs my-backend-container
```

4. How to Stop and Remove the Container  
If you need to restart the container with new settings, you must first stop and remove the old one.

```Bash
# Stop the running container
docker stop my-backend-container

# Remove the stopped container
docker rm my-backend-container
```

**NOTE:** You can also do both at once by forcing removal:  
`docker rm -f my-backend-container`.