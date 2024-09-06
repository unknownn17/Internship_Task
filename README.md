# Task Management API

Bu loyiha foydalanuvchilarga topshiriqlarni boshqarish uchun RESTful API xizmatini taqdim etadi. Loyihada **PostgreSQL**, **MongoDB**, va **Redis** kabi texnologiyalar ishlatilgan.

## Texnologiyalar
- **PostgreSQL**: SQL injektsiyasidan himoyalanish uchun **sqlc** texnologiyasi ishlatilgan.
- **MongoDB**: Ma'lumotlarni saqlash uchun ishlatilgan asosiy ma'lumotlar bazasi.
- **Redis**: Foydalanuvchi ro'yxatdan o'tganda va autentifikatsiya jarayonida cache uchun ishlatiladi.

## Autentifikatsiya Jarayoni

1. **Ro'yxatdan o'tish**: Foydalanuvchi ro'yxatdan o'tish ma'lumotlarini kiritgandan so'ng, ushbu ma'lumotlar Redis'ga cache qilinadi. Foydalanuvchi kiritgan elektron pochta manziliga maxfiy kod yuboriladi.
2. **Tasdiqlash**: Foydalanuvchi elektron pochta orqali yuborilgan kodni kiritib, hisobni tasdiqlashi kerak.
3. **Kirish**: Foydalanuvchi ro'yxatdan o'tgan elektron pochta va parolni kiritib, tizimga kiradi va **JWT** token oladi.
4. **Task Xizmatlari**: Autentifikatsiyadan so'ng, foydalanuvchi topshiriqlar xizmatidan erkin foydalanishi mumkin.

## API Endpoints

**Foydalanuvchi xizmatlari:**
- `POST /user/register` – Foydalanuvchini ro'yxatdan o'tkazish
- `POST /user/verify` – Foydalanuvchini tasdiqlash
- `POST /user/login` – Foydalanuvchini tizimga kiritish

**Topshiriq xizmatlari** (JWT autentifikatsiyasi talab qilinadi):
- `POST /task` – Yangi topshiriq yaratish
- `GET /task` – Bitta topshiriqni olish
- `GET /tasks` – Foydalanuvchining barcha topshiriqlarini olish
- `PUT /task` – Topshiriqni yangilash
- `DELETE /task` – Topshiriqni o'chirish

Qo'shimcha ma'lumotlar uchun [handler.go](./internal/connections/handler.go) va [router.go](./internal/router/router.go) fayllarini ko'rib chiqing.

## Swagger Hujjatlari
Loyihada **Swagger** ishlatilgan, API hujjatlarini quyidagi endpoint orqali ko'rishingiz mumkin: 
## http://3.127.221.197:8080/swagger/index.html

