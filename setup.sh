#!/bin/bash

echo "==================================="
echo "E-Commerce Sepatu - Setup Script"
echo "==================================="
echo ""

# Check if .env exists
if [ ! -f .env ]; then
  echo "⚠️  .env file not found!"
  echo "Creating .env from .env.example..."
  cp .env.example .env
  echo "✅ .env created. Please update with your values."
fi

echo ""
echo "Starting Docker containers..."
docker-compose up -d

echo ""
echo "Waiting for PostgreSQL to be ready..."
sleep 5

echo ""
echo "Installing frontend dependencies..."
cd frontend
npm install

echo ""
echo "Generating Prisma Client..."
npx prisma generate

echo ""
echo "Running database migrations..."
npx prisma migrate dev --name init

cd ..

echo ""
echo "==================================="
echo "✅ Setup Complete!"
echo "==================================="
echo ""
echo "Frontend: http://localhost/"
echo "Backend API: http://localhost/api/"
echo "Database: localhost:5432"
echo ""
echo "Next steps:"
echo "1. Update .env with proper secrets (JWT_SECRET, NEXTAUTH_SECRET)"
echo "2. npm run dev (in frontend folder) for local development"
echo "3. Create admin user and start adding products!"
echo ""
