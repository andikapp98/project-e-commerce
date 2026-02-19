# 🚀 Quick Start - Login System

## Requirements

- Docker & Docker Compose
- Node.js 20+ (untuk dev lokal)
- Go 1.23+ (untuk backend dev)

---

## Jalankan Dengan Docker (Paling Mudah!)

### Step 1: Generate Secrets

```bash
# Generate JWT_SECRET (copy hasilnya)
openssl rand -base64 32

# Generate NEXTAUTH_SECRET (copy hasilnya)
openssl rand -base64 32
```

### Step 2: Update .env

Edit `.env` di root folder:

```env
JWT_SECRET=<paste-hasil-openssl-di-sini>
NEXTAUTH_SECRET=<paste-hasil-openssl-di-sini>
```

### Step 3: Jalankan Docker

```bash
docker-compose up -d
```

### Step 4: Setup Database

```bash
# Tunggu PostgreSQL ready (5-10 detik)
docker exec project-ecommerce-frontend npx prisma migrate dev --name init
```

✅ **Selesai!**

- Frontend: http://localhost/
- Backend: http://localhost/api/health
- Database: localhost:5432

---

## Coba Login System

### 1. Register User Baru

Buka: http://localhost/auth/register

Isi form:

- Email: `test@example.com`
- Password: `password123`

### 2. Login

Buka: http://localhost/auth/login

Gunakan email & password yang sudah di-register

### 3. API Testing

Register via API:

```bash
curl -X POST http://localhost/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "password123",
    "name": "John Doe"
  }'
```

---

## Development Lokal (Tanpa Docker)

### Frontend

```bash
cd frontend
npm install
npx prisma generate
npx prisma migrate dev --name init
npm run dev  # http://localhost:3000
```

### Backend

```bash
cd backend
go mod tidy
export DATABASE_URL="postgresql://ecommerce_user:ecommerce_password@localhost:5432/ecommerce_db"
go run main.go  # http://localhost:8080
```

---

## Struktur Login System

```
┌─ User Register (POST /api/auth/register)
│  └─ Hash password with bcryptjs
│  └─ Save to database
│
├─ User Login (NextAuth credentials)
│  └─ Find user by email
│  └─ Verify password
│  └─ Create JWT session
│
└─ Protected Pages
   └─ Use middleware for /dashboard/* routes
```

---

## Database Schema

```sql
CREATE TABLE "User" (
  id STRING PRIMARY KEY,
  email STRING UNIQUE NOT NULL,
  name STRING,
  password STRING NOT NULL,
  role STRING DEFAULT 'customer',
  isActive BOOLEAN DEFAULT true,
  createdAt TIMESTAMP,
  updatedAt TIMESTAMP
);
```

---

## Common Issues & Fix

### ❌ "Connection refused" saat docker up

```bash
# Make sure Docker daemon running
docker --version

# Check containers
docker-compose ps
```

### ❌ Port 80 atau 5432 sudah terpakai

Edit `docker-compose.yml`:

```yaml
services:
  postgres:
    ports:
      - "5433:5432" # Change 5432 to 5433

  nginx:
    ports:
      - "8080:80" # Change 80 to 8080
```

### ❌ Kode register tapi error "Email already registered"

User sudah ada di database. Gunakan email yang lain atau bersihkan database:

```bash
docker-compose down -v
docker-compose up -d
docker exec project-ecommerce-frontend npx prisma migrate dev
```

### ❌ Login gagal "User not found"

Pastikan sudah register terlebih dahulu atau check database:

```bash
docker exec -it project-ecommerce-db psql -U ecommerce_user -d ecommerce_db
SELECT * FROM "User";
```

---

## Useful Commands

```bash
# View logs
docker-compose logs -f frontend      # Frontend logs
docker-compose logs -f backend       # Backend logs
docker-compose logs -f postgres      # Database logs
docker-compose logs -f nginx         # Nginx logs

# Stop containers
docker-compose down

# Remove everything (hati-hati!)
docker-compose down -v

# Restart service
docker-compose restart frontend
docker-compose restart backend

# Reset database
docker-compose exec frontend npx prisma migrate reset
```

---

## Next Steps

Setelah login berhasil:

1. ✅ Implement product listing
2. ✅ Add shopping cart
3. ✅ Payment integration
4. ✅ Admin dashboard

---

**Butuh bantuan? Baca SETUP.md untuk dokumentasi lengkap** 📖
