# Project E-Commerce Sepatu

## Deskripsi
Proyek e-commerce sepatu yang dirancang untuk membantu pengguna menjual sepatu dengan kategori seperti sepatu business, sport, dan casual. Aplikasi ini mencakup frontend berbasis Next.js, backend berbasis Golang, dan database PostgreSQL dengan Prisma sebagai ORM.

---

## Fitur Utama
1. **Frontend (UI/UX):**
   - Dibangun dengan **Next.js** menggunakan TypeScript & Tailwind CSS.
   - Desain responsif untuk pengalaman pengguna yang optimal di desktop dan mobile.

2. **Backend (API):**
   - **Golang (Fiber):** API yang cepat dan ringan untuk menangani logika bisnis.
   - Koneksi ke database menggunakan Prisma atau SQL murni.

3. **Database:**
   - Skema database dengan entitas utama: **Product**, **Variant**, **Order**, dan **OrderItem**.
   - Dibangun untuk mempermudah pengelolaan stok, pesanan, dan pembayaran.

4. **Fitur Tambahan:**
   - Sistem autentikasi pengguna.
   - Integrasi payment gateway (**Midtrans** atau lainnya).
   - Dashboard Admin untuk mengelola produk dan pesanan.

---

## Tech Stack
1. **Frontend:**
   - [Next.js](https://nextjs.org/): Framework React modern.
   - **TypeScript** untuk keamanan tipe.
   - **Tailwind CSS** untuk styling cepat & konsisten.
2. **Backend:**
   - [Golang](https://golang.org/): Backend powerful untuk performa tinggi.
   - **Fiber** untuk framework API.
3. **Database:**
   - PostgreSQL dengan Prisma ORM.

---

## Cara Menjalankan Project

### 1. Clone Repository
```bash
git clone git@github.com:andikapp98/project-e-commerce.git
cd project-e-commerce
```

### 2. Setup Frontend
```bash
cd frontend
npm install
cp .env.local.example .env.local
npx prisma generate
npm run dev
```

### 3. Backend (Opsional, jika diperlukan sekarang)
- Setup backend Golang di direktori backend (belum dibuat).

---

## Tim Proyek
- **Nama Developer:** Andika (GitHub: [@andikapp98](https://github.com/andikapp98))
- **Asisten AI:** Upilbot 🤖

---

## Catatan Tambahan
Ini adalah proyek open-source yang dapat dikembangkan lebih lanjut sesuai kebutuhan. Jangan ragu untuk menghubungi saya jika butuh bantuan lebih lanjut!