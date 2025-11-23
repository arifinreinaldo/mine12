# Sales API - Go Fiber Application

A RESTful API for sales management built with Go, Fiber framework, and PostgreSQL.

## Features

- ✅ User management with access level control (Admin, Manager, Staff, Read-only)
- ✅ Master item/product management
- ✅ CRUD operations for both users and items
- ✅ PostgreSQL database with GORM
- ✅ Docker and Docker Compose support
- ✅ Soft delete functionality
- ✅ Input validation
- ✅ Password hashing with bcrypt

## Tech Stack

- **Framework**: Go Fiber v2
- **Database**: PostgreSQL 15
- **ORM**: GORM
- **Containerization**: Docker & Docker Compose
- **Validation**: go-playground/validator
- **Password Hashing**: bcrypt

## Project Structure

```
.
├── config/          # Configuration management
├── database/        # Database connection and migrations
├── handlers/        # HTTP request handlers
├── models/          # Data models and DTOs
├── main.go          # Application entry point
├── Dockerfile       # Docker image configuration
├── docker-compose.yml
├── .env.example
└── README.md
```

## Getting Started

### Prerequisites

- Go 1.21 or higher
- Docker and Docker Compose (for containerized deployment)
- PostgreSQL 15 (if running without Docker)

### Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd mine12
   ```

2. **Set up environment variables**
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

3. **Install dependencies**
   ```bash
   go mod download
   ```

## Running the Application

### Option 1: Using Docker Compose (Recommended)

This is the easiest way to run the application with all dependencies.

```bash
# Start all services (app + database)
docker-compose up -d

# View logs
docker-compose logs -f

# Stop all services
docker-compose down

# Stop and remove volumes (clean slate)
docker-compose down -v
```

The API will be available at `http://localhost:3000`

### Option 2: Running Locally

1. **Start PostgreSQL** (skip if using Docker for DB only)
   ```bash
   docker-compose up -d postgres
   ```

2. **Run the application**
   ```bash
   go run main.go
   ```

3. **Or build and run**
   ```bash
   go build -o sales-api
   ./sales-api
   ```

## API Endpoints

### Health Check
- `GET /` - API health check

### Users API

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/v1/users` | Get all users |
| GET | `/api/v1/users/:id` | Get user by ID |
| POST | `/api/v1/users` | Create new user |
| PUT | `/api/v1/users/:id` | Update user |
| DELETE | `/api/v1/users/:id` | Delete user (soft delete) |

### Items API

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/v1/items` | Get all items |
| GET | `/api/v1/items/:id` | Get item by ID |
| POST | `/api/v1/items` | Create new item |
| PUT | `/api/v1/items/:id` | Update item |
| DELETE | `/api/v1/items/:id` | Delete item (soft delete) |

## API Usage Examples

### Create User

```bash
curl -X POST http://localhost:3000/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "username": "john_doe",
    "email": "john@example.com",
    "password": "secure123",
    "full_name": "John Doe",
    "access_level": "staff"
  }'
```

**Access Levels**: `admin`, `manager`, `staff`, `readonly`

### Create Item

```bash
curl -X POST http://localhost:3000/api/v1/items \
  -H "Content-Type: application/json" \
  -d '{
    "code": "ITEM001",
    "name": "Sample Product",
    "description": "This is a sample product",
    "price": 99.99,
    "stock": 100,
    "unit": "pcs"
  }'
```

### Get All Users

```bash
curl http://localhost:3000/api/v1/users
```

### Update User

```bash
curl -X PUT http://localhost:3000/api/v1/users/1 \
  -H "Content-Type: application/json" \
  -d '{
    "full_name": "John Updated Doe",
    "access_level": "manager"
  }'
```

### Delete Item

```bash
curl -X DELETE http://localhost:3000/api/v1/items/1
```

## Deployment Guide

### Deploy to Production Server

#### 1. Deploy with Docker Compose (Recommended)

```bash
# On your production server
git clone <repository-url>
cd mine12

# Set production environment variables
cp .env.example .env
nano .env  # Edit with production values

# Build and start services
docker-compose up -d

# Check logs
docker-compose logs -f app
```

#### 2. Deploy to Cloud Platforms

##### AWS EC2 / DigitalOcean / VPS

```bash
# Install Docker and Docker Compose on your server
sudo apt update
sudo apt install docker.io docker-compose -y

# Clone and run
git clone <repository-url>
cd mine12
cp .env.example .env
nano .env  # Configure for production

docker-compose up -d

# Set up reverse proxy (nginx)
sudo apt install nginx -y
```

**Nginx Configuration** (`/etc/nginx/sites-available/sales-api`):
```nginx
server {
    listen 80;
    server_name your-domain.com;

    location / {
        proxy_pass http://localhost:3000;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection 'upgrade';
        proxy_set_header Host $host;
        proxy_cache_bypass $http_upgrade;
    }
}
```

```bash
# Enable site and reload nginx
sudo ln -s /etc/nginx/sites-available/sales-api /etc/nginx/sites-enabled/
sudo nginx -t
sudo systemctl reload nginx

# Optional: Set up SSL with Let's Encrypt
sudo apt install certbot python3-certbot-nginx -y
sudo certbot --nginx -d your-domain.com
```

##### Heroku

```bash
# Install Heroku CLI
# Login to Heroku
heroku login

# Create new app
heroku create your-app-name

# Add PostgreSQL addon
heroku addons:create heroku-postgresql:mini

# Set config vars
heroku config:set APP_PORT=8080

# Deploy
git push heroku main

# Open app
heroku open
```

##### Google Cloud Run

```bash
# Install gcloud CLI
# Authenticate
gcloud auth login

# Set project
gcloud config set project YOUR_PROJECT_ID

# Build and push image
gcloud builds submit --tag gcr.io/YOUR_PROJECT_ID/sales-api

# Deploy to Cloud Run
gcloud run deploy sales-api \
  --image gcr.io/YOUR_PROJECT_ID/sales-api \
  --platform managed \
  --region us-central1 \
  --allow-unauthenticated \
  --add-cloudsql-instances YOUR_PROJECT_ID:REGION:INSTANCE_NAME \
  --set-env-vars DB_HOST=/cloudsql/YOUR_PROJECT_ID:REGION:INSTANCE_NAME
```

##### Railway.app

```bash
# Install Railway CLI
npm i -g @railway/cli

# Login
railway login

# Initialize project
railway init

# Add PostgreSQL
railway add

# Deploy
railway up
```

#### 3. Deploy with Docker Swarm

```bash
# Initialize swarm
docker swarm init

# Deploy stack
docker stack deploy -c docker-compose.yml sales-app

# Check services
docker service ls

# Scale app service
docker service scale sales-app_app=3
```

#### 4. Deploy with Kubernetes

Create `k8s-deployment.yml`:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: sales-api
spec:
  replicas: 3
  selector:
    matchLabels:
      app: sales-api
  template:
    metadata:
      labels:
        app: sales-api
    spec:
      containers:
      - name: sales-api
        image: your-registry/sales-api:latest
        ports:
        - containerPort: 3000
        env:
        - name: DB_HOST
          value: "postgres-service"
        - name: DB_PORT
          value: "5432"
---
apiVersion: v1
kind: Service
metadata:
  name: sales-api-service
spec:
  selector:
    app: sales-api
  ports:
  - port: 80
    targetPort: 3000
  type: LoadBalancer
```

Deploy:
```bash
kubectl apply -f k8s-deployment.yml
```

### Production Checklist

- [ ] Set strong database passwords
- [ ] Configure firewall rules
- [ ] Set up SSL/TLS certificates
- [ ] Configure CORS properly
- [ ] Set up monitoring (Prometheus, Grafana)
- [ ] Configure logging
- [ ] Set up database backups
- [ ] Configure health checks
- [ ] Set resource limits in Docker
- [ ] Use environment-specific configurations
- [ ] Set up CI/CD pipeline

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| DB_HOST | PostgreSQL host | localhost |
| DB_PORT | PostgreSQL port | 5432 |
| DB_USER | Database user | postgres |
| DB_PASSWORD | Database password | postgres |
| DB_NAME | Database name | sales_db |
| APP_PORT | Application port | 3000 |

## Database Schema

### Users Table
- `id` - Primary key
- `username` - Unique username
- `email` - Unique email
- `password` - Hashed password
- `full_name` - User's full name
- `access_level` - User role (admin/manager/staff/readonly)
- `is_active` - Active status
- `created_at` - Creation timestamp
- `updated_at` - Update timestamp
- `deleted_at` - Soft delete timestamp

### Items Table
- `id` - Primary key
- `code` - Unique item code
- `name` - Item name
- `description` - Item description
- `price` - Item price
- `stock` - Available stock
- `unit` - Unit of measurement
- `is_active` - Active status
- `created_at` - Creation timestamp
- `updated_at` - Update timestamp
- `deleted_at` - Soft delete timestamp

## Development

### Run Tests
```bash
go test ./... -v
```

### Build for Production
```bash
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
```

## Troubleshooting

### Database Connection Issues
```bash
# Check if PostgreSQL is running
docker-compose ps

# Check logs
docker-compose logs postgres

# Restart services
docker-compose restart
```

### Port Already in Use
```bash
# Change APP_PORT in .env file
# Or stop the conflicting service
lsof -ti:3000 | xargs kill -9
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License.

## Support

For issues and questions, please open an issue in the repository.
