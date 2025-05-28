# Otakudesu-Api

API dan Scraper Otakudesu yang dibuat menggunakan Golang dengan Fiber framework.

## 📋 Deskripsi
> **⚠️ Peringatan**: Ini adalah API tidak resmi yang mengambil data dari situs web Otakudesu menggunakan teknik scraping. API ini dapat berhenti bekerja jika ada perubahan signifikan pada struktur HTML situs tersebut. Gunakan dengan risiko Anda sendiri.

API tidak resmi untuk mengambil data dari Otakudesu yang dibangun dengan Go (Golang) dan framework Fiber.


## 🚀 Fitur

- Scraping data anime dari Otakudesu
- REST API dengan Golang Fiber
- Performa tinggi dan efisien
- Mudah digunakan

## 💻 Teknologi yang Digunakan

- [Go](https://golang.org/) - Bahasa pemrograman utama
- [Fiber](https://gofiber.io/) - Web framework yang cepat untuk Golang
- [Colly](https://go-colly.org/) - Web scraping framework untuk Golang

## ⚙️ Instalasi

1. Pastikan Go sudah terinstal di sistem Anda
2. Clone repositori ini
   ```bash
   git clone https://github.com/Satr10/Otakudesu-Api.git
   ```
3. Masuk ke direktori proyek
   ```bash
   cd Otakudesu-Api
   ```
4. Install dependensi
   ```bash
   go mod tidy
   ```
5. Jalankan aplikasi
   ```bash
   go run main.go
   ```

## 🛠️ Penggunaan / Dokumentasi Endpoint API

Berikut adalah daftar endpoint yang tersedia beserta contoh request dan responsnya.

### GET /api/

Menampilkan pesan selamat datang untuk API.

**Contoh Request:**
```
GET http://127.0.0.1:5001/api/
```

**Contoh Response Sukses (200 OK):**
```json
{
  "status": "success",
  "message": "Welcome to Otakudesu API, (THIS IS AN UNOFFICIAL API)",
  "data": null
}
```

### GET /api/home

Mengambil daftar anime yang ada di halaman utama Otakudesu.

**Contoh Request:**
```
GET http://127.0.0.1:5001/api/home
```

**Contoh Response Sukses (200 OK):**
```json
{
  "status": "success",
  "message": "",
  "data": [
    {
      "title": "Kidou Senshi Gundam: GQuuuuuuX",
      "episode": "Episode 8",
      "status": "Ongoing",
      "schedule": "Rabu",
      "rating": null,
      "date": "28 Mei",
      "slug": "kidou-gundam-gquuuuuux-sub-indo",
      "image": "https://otakudesu.cloud/wp-content/uploads/2025/04/bx185213-RKQN0qxKlqd6.jpg",
      "url": "https://otakudesu.cloud/anime/kidou-gundam-gquuuuuux-sub-indo/"
    },
    {
      "title": "Aru Majo ga Shinu Made",
      "episode": "Episode 9",
      "status": "Ongoing",
      "schedule": "Selasa",
      "rating": null,
      "date": "28 Mei",
      "slug": "aru-majo-shinu-made-sub-indo",
      "image": "https://otakudesu.cloud/wp-content/uploads/2025/03/148221.jpg",
      "url": "https://otakudesu.cloud/anime/aru-majo-shinu-made-sub-indo/"
    }
  ]
}
```

**Contoh Response Gagal (Error):**
```json
{
  "status": "failed",
  "message": "failed to fetch animes",
  "data": null
}
```

### GET /api/ongoing
### GET /api/ongoing/:page

Mengambil daftar anime yang sedang tayang (ongoing), dengan paginasi.

**Path Parameters:**
- `page` (opsional): Nomor halaman. Default ke 1 jika tidak disertakan.

**Contoh Request:**
```
GET http://127.0.0.1:5001/api/ongoing/1
```
atau
```
GET http://127.0.0.1:5001/api/ongoing
```

**Contoh Response Sukses (200 OK):**
```json
{
  "status": "success",
  "message": "",
  "data": [
    {
      "title": "Kidou Senshi Gundam: GQuuuuuuX",
      "episode": "Episode 8",
      "status": "Ongoing",
      "schedule": "Rabu",
      "rating": null,
      "date": "28 Mei",
      "slug": "kidou-gundam-gquuuuuux-sub-indo",
      "image": "https://otakudesu.cloud/wp-content/uploads/2025/04/bx185213-RKQN0qxKlqd6.jpg",
      "url": "https://otakudesu.cloud/anime/kidou-gundam-gquuuuuux-sub-indo/"
    }
  ]
}
```

### GET /api/completed
### GET /api/completed/:page

Mengambil daftar anime yang sudah tamat (completed), dengan paginasi.

**Path Parameters:**
- `page` (opsional): Nomor halaman. Default ke 1 jika tidak disertakan.

**Contoh Request:**
```
GET http://127.0.0.1:5001/api/completed/1
```
atau
```
GET http://127.0.0.1:5001/api/completed
```

**Contoh Response Sukses (200 OK):**
```json
{
  "status": "success",
  "message": "",
  "data": [
    {
      "title": "Hana wa Saku, Shura no Gotoku",
      "episode": "12 Episode",
      "status": "Completed",
      "schedule": null,
      "rating": "7.14",
      "date": "22 Mei",
      "slug": "hana-saku-shura-gotoku-sub-indo",
      "image": "https://otakudesu.cloud/wp-content/uploads/2025/05/Hana-wa-Saku-Shura-no-Gotoku-Sub-Indo.jpg",
      "url": "https://otakudesu.cloud/anime/hana-saku-shura-gotoku-sub-indo/"
    }
  ]
}
```

### GET /api/search/:query

Mencari anime berdasarkan kata kunci.

**Path Parameters:**
- `query` (wajib): Kata kunci pencarian.

**Contoh Request:**
```
GET http://127.0.0.1:5001/api/search/boku%20dake
```

**Contoh Response Sukses (200 OK):**
```json
{
  "status": "success",
  "message": "",
  "data": [
    {
      "title": "Boku dake ga Inai Machi BD (Episode 1 – 12) Sub Indo",
      "episode": "",
      "status": "Completed",
      "schedule": null,
      "rating": "8.61",
      "date": "",
      "slug": "boku-dake-inai-machi-sub-indo",
      "image": "https://otakudesu.cloud/wp-content/uploads/2017/01/Boku-dake-ga-Inai-Machi-Sub-Indo.jpg",
      "url": "https://otakudesu.cloud/anime/boku-dake-inai-machi-sub-indo/"
    }
  ]
}
```

### GET /api/anime/:slug

Mengambil detail informasi suatu anime berdasarkan slug-nya.

**Path Parameters:**
- `slug` (wajib): Slug dari anime.

**Contoh Request:**
```
GET http://127.0.0.1:5001/api/anime/boku-dake-inai-machi-sub-indo
```

**Contoh Response Sukses (200 OK):**
```json
{
  "status": "success",
  "message": "",
  "data": {
    "title": "Boku dake ga Inai Machi",
    "japaneseTitle": "僕だけがいない街",
    "rating": "8.61",
    "producer": "Aniplex, Dentsu, Kadokawa Shoten, Fuji TV, DAX Production, other.",
    "type": "BD",
    "status": "Completed",
    "episodeTotal": "12",
    "duration": "23 Menit / Episode",
    "releaseDate": "08 Januari, 2016 Sampai 25 Maret, 2016",
    "studio": "A-1 Pictures",
    "genre": "Mystery, Psychological, Seinen, Supernatural",
    "synopsis": "Satoru Fujinuma, ia adalah seorang mangaka yang karirnya tidak terlalu sukses...",
    "episodes": [
      {
        "episodeTitle": "Boku Dake ga Inai Machi [BATCH] Subtitle Indonesia",
        "slug": "bdginm-batch-sub-indo",
        "url": "https://otakudesu.cloud/batch/bdginm-batch-sub-indo/"
      }
    ]
  }
}
```

### GET /api/episode/:slug

Mengambil detail informasi dan link download suatu episode berdasarkan slug-nya.

**Path Parameters:**
- `slug` (wajib): Slug dari episode.

**Contoh Request:**
```
GET http://127.0.0.1:5001/api/episode/bokmacih-episode-1-sub-indo/
```

**Contoh Response Sukses (200 OK):**
```json
{
  "status": "success",
  "message": "",
  "data": {
    "downloads": [
      {
        "quality": "360p",
        "size": "26.5 MB",
        "downloads": [
          {
            "provider": "ZippyShare",
            "downloadUrl": "https://desustream.com/safelink/link/?id=..."
          }
        ]
      }
    ]
  }
}
```

### GET /api/genre-list

Mengambil daftar semua genre yang tersedia.

**Contoh Request:**
```
GET http://127.0.0.1:5001/api/genre-list
```

**Contoh Response Sukses (200 OK):**
```json
{
  "status": "success",
  "message": "",
  "data": {
    "genres": [
      {
        "title": "Action",
        "slug": "action",
        "url": "https://otakudesu.cloud/genres/action/"
      },
      {
        "title": "Adventure",
        "slug": "adventure",
        "url": "https://otakudesu.cloud/genres/adventure/"
      }
    ]
  }
}
```

### GET /api/genre/:slug
### GET /api/genre/:slug/:page

Mengambil daftar anime berdasarkan genre tertentu, dengan paginasi.

**Path Parameters:**
- `slug` (wajib): Slug dari genre.
- `page` (opsional): Nomor halaman. Default ke 1 jika tidak disertakan.

**Contoh Request:**
```
GET http://127.0.0.1:5001/api/genre/romance/2
```
atau
```
GET http://127.0.0.1:5001/api/genre/romance
```

**Contoh Response Sukses (200 OK):**
```json
{
  "status": "success",
  "message": "",
  "data": [
    {
      "title": "Hyakkano Season 2",
      "episode": "Unknown",
      "status": "",
      "schedule": null,
      "rating": "7.81",
      "date": "",
      "slug": "hyakkano-s2-sub-indo",
      "image": "https://otakudesu.cloud/wp-content/uploads/2025/01/145470.jpg",
      "url": "https://otakudesu.cloud/anime/hyakkano-s2-sub-indo/"
    }
  ]
}
```

## 📝 Catatan

- Proyek ini adalah hasil penulisan ulang dari versi Python+FastAPI ke Golang+Fiber
- Gunakan API ini dengan bijak dan sesuai dengan ketentuan yang berlaku
- Hormati rate limiting dan kebijakan penggunaan

## 📜 Lisensi

[LISENSI](LICENSE)

Copyright © 2025 Satrio <personal.satrio@protonmail.com>
This work is free. You can redistribute it and/or modify it under the
terms of the Do What The Fuck You Want To Public License, Version 2,
as published by Sam Hocevar. See http://www.wtfpl.net/ for more details.

## 🤝 Kontribusi

Kontribusi selalu diterima dengan senang hati! Berikut langkah-langkah untuk berkontribusi:

1. Fork repositori ini
2. Buat branch baru (`git checkout -b fitur-baru`)
3. Commit perubahan Anda (`git commit -m 'Menambahkan fitur baru'`)
4. Push ke branch tersebut (`git push origin fitur-baru`)
5. Buat Pull Request

## ⚠️ Disclaimer

Proyek ini dibuat untuk tujuan pembelajaran dan pengembangan. Pastikan untuk menggunakan API ini sesuai dengan ketentuan dan peraturan yang berlaku.

## 📞 Kontak

- GitHub: [@Satr10](https://github.com/Satr10)

---

⭐ **Jika Anda menyukai proyek ini, berikan bintang!** ⭐