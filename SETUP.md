# Setup Guide - Project E-Commerce Sepatu

## Prerequisites

- Docker & Docker Compose installed
- Node.js 20+ (untuk development lokal)
- Go 1.23+ (untuk backend development)

## Quick Start dengan Docker

### 1. Clone & Enter Directory

```bash
cd project-e-commerce
```

### 2. Copy Environment File

```bash
cp .env.example .env
```

Update nilai `JWT_SECRET` dan `NEXTAUTH_SECRET` dengan random string yang aman:

```bash
# Generate JWT_SECRET
openssl rand -base64 32

# Generate NEXTAUTH_SECRET
openssl rand -base64 32
```

### 3. Setup Database

```bash
# Jalankan docker-compose
docker-compose up -d

# Tunggu PostgreSQL ready, kemudian jalankan migration
docker exec project-ecommerce-frontend npx prisma migrate dev --name init
```

### 4. Akses Aplikasi

- **Frontend:** http://localhost/
- **Backend API:** http://localhost/api/
- **Database:** localhost:5432

---

## Development Setup (Tanpa Docker)

### Frontend

```bash
cd frontend

# Install dependencies
npm install

# Setup environment
cp .env.example .env.local

# Update DATABASE_URL di .env.local
DATABASE_URL="postgresql://your-user:your-password@localhost:5432/your-db"

# Generate Prisma Client
npx prisma generate

# Migration (jika belum)
npx prisma migrate dev --name init

# Run development server
npm run dev
```

**Akses:** http://localhost:3000

### Backend

```bash
cd backend

# Generate go.sum
go mod tidy

# Update .env file di root
export DATABASE_URL="postgresql://your-user:your-password@localhost:5432/your-db"

# Run server
go run main.go
```

**Server runs at:** http://localhost:8080

---

## Authentication

### Register

**URL:** `POST http://localhost/api/auth/register`

Request:

```json
{
  "email": "user@example.com",
  "password": "password123",
  "name": "John Doe"
}
```

### Login

**URL:** `GET http://localhost/auth/login`

Atau via API:
**URL:** `POST http://localhost/api/auth/callback/credentials`

### Logout

**URL:** `http://localhost/api/auth/logout`

---

## Project Structure

```
project-e-commerce/
├── frontend/                 # Next.js Frontend
│   ├── app/
│   │   ├── api/auth/        # Authentication API routes
│   │   ├── auth/            # Login & Register pages
│   │   └── page.tsx         # Home page
│   ├── lib/
│   │   └── prisma.ts        # Prisma client
│   ├── prisma/
│   │   └── schema.prisma    # Database schema
│   └── package.json
│
├── backend/                  # Go Fiber Backend
│   ├── main.go              # API routes & handlers
│   ├── go.mod
│   ├── go.sum
│   └── Dockerfile
│
├── nginx/                    # Nginx configuration
│   └── nginx.conf           # Reverse proxy setup
│
├── postgres-init/           # PostgreSQL initialization
│   └── 01-init.sql         # Database setup script
│
├── docker-compose.yml       # Docker orchestration
├── .env                     # Environment variables
└── README.md
```

---

## Database Schema

### User Table

```sql
CREATE TABLE "User" (
  id VARCHAR(255) PRIMARY KEY,
  email VARCHAR(255) UNIQUE NOT NULL,
  name VARCHAR(255),
  password VARCHAR(255) NOT NULL,
  role VARCHAR(50) DEFAULT 'customer',
  isActive BOOLEAN DEFAULT true,
  createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

- **role:** 'customer' atau 'admin'
- **isActive:** Flag untuk nonaktifkan user tanpa delete

---

## API Routes

### Authentication

- `POST /api/auth/register` - Register user baru
- `POST /api/auth/callback/credentials` - Login
- `POST /api/auth/logout` - Logout

### Products (Backend)

- `GET /products` - Daftar produk
- `GET /products/:id` - Detail produk
- `POST /products` - Buat produk (admin)
- `PUT /products/:id` - Update produk (admin)
- `DELETE /products/:id` - Delete produk (admin)

### Orders (Backend)

- `GET /orders` - Daftar pesanan
- `GET /orders/:id` - Detail pesanan
- `POST /orders` - Buat pesanan
- `PUT /orders/:id` - Update pesanan

---

## Environment Variables

| Variable          | Purpose                | Default            |
| ----------------- | ---------------------- | ------------------ |
| `DB_USER`         | PostgreSQL username    | ecommerce_user     |
| `DB_PASSWORD`     | PostgreSQL password    | ecommerce_password |
| `DB_NAME`         | Database name          | ecommerce_db       |
| `DB_PORT`         | Database port          | 5432               |
| `DATABASE_URL`    | Full connection string | (auto-generated)   |
| `NEXTAUTH_SECRET` | NextAuth secret key    | (must set!)        |
| `NEXTAUTH_URL`    | Frontend URL           | http://localhost   |
| `JWT_SECRET`      | Backend JWT secret     | (must set!)        |
| `NODE_ENV`        | Environment            | development        |

---

## Troubleshooting

### Database Connection Error

```bash
# Check PostgreSQL status
docker-compose ps

# View PostgreSQL logs
docker-compose logs postgres

# Restart PostgreSQL
docker-compose restart postgres
```

### Port Already in Use

```bash
# Change ports di docker-compose.yml
# Contoh: ganti "5432:5432" jadi "5433:5432"
```

### Migration Issues

```bash
# Reset database (HATI-HATI: akan delete semua data!)
docker exec project-ecommerce-frontend npx prisma migrate reset

# Atau manual:
docker-compose down -v  # Remove volumes
docker-compose up -d    # Restart
```

### NextAuth Session Issues

- Pastikan `NEXTAUTH_SECRET` dan `NEXTAUTH_URL` sudah set di `.env.local`
- Restart frontend container jika ubah environment
- Clear browser cookies

---

## Next Steps

1. ✅ Setup login/register
2. ⏳ Implement product listing
3. ⏳ Setup shopping cart
4. ⏳ Payment integration (Midtrans)
5. ⏳ Admin dashboard

---

**Happy coding! 🚀**
