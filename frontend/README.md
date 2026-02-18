# E-Commerce Sepatu - Frontend

Proyek frontend untuk toko e-commerce sepatu business & sport.

## Tech Stack

- **Next.js 15** (App Router)
- **TypeScript**
- **Tailwind CSS**
- **Prisma** (ORM)
- **PostgreSQL**

## Cara Menjalankan

1. Install dependencies:
```bash
npm install
```

2. Copy environment variables:
```bash
cp .env.local.example .env.local
```

3. Sesuaikan `DATABASE_URL` di `.env.local`

4. Generate Prisma client:
```bash
npx prisma generate
```

5. Jalankan development server:
```bash
npm run dev
```

## Fitur

- Daftar produk sepatu
- Detail produk dengan varian size, warna, dan lebar
- Keranjang belanja
- Checkout & pembayaran
- Admin dashboard untuk mengelola produk