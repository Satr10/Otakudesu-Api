# Otakudesu API (Tidak Resmi)

> **⚠️ Peringatan**: Ini adalah API tidak resmi yang mengambil data dari situs web Otakudesu menggunakan teknik scraping. API ini dapat berhenti bekerja jika ada perubahan signifikan pada struktur HTML situs tersebut. Gunakan dengan risiko Anda sendiri.

API tidak resmi untuk mengambil data dari Otakudesu yang dibangun dengan Go (Golang) dan framework Fiber.

## Daftar Isi

- [URL Dasar](#url-dasar)
- [Struktur Respons](#struktur-respons)
- [Endpoint](#endpoint)
- [Struktur Data](#struktur-data)
- [Penanganan Error](#penanganan-error)

## URL Dasar

Semua URL yang direferensikan dalam dokumentasi ini memiliki URL dasar berikut:

```
http://localhost:5001/api
```

Ganti `localhost:5001` dengan domain dan port server Anda jika di-hosting di tempat lain.

## Struktur Respons

Semua respons API mengikuti struktur JSON berikut:

```json
{
    "status": "success" | "failed",
    "message": "Pesan deskriptif atau string kosong jika berhasil tanpa pesan khusus",
    "data": null | {} | []
}
```

### Field Respons

| Field | Tipe | Deskripsi |
|-------|------|-----------|
| `status` | string | Status permintaan (`success` atau `failed`) |
| `message` | string | Informasi tambahan tentang status, terutama untuk error |
| `data` | mixed | Data yang diminta (bisa `null`, objek JSON, atau array) |

## Endpoint

### Pesan Selamat Datang

**GET** `/`

Mengembalikan pesan selamat datang untuk API.

**Contoh Permintaan:**
```
GET http://localhost:5001/api/
```

**Contoh Respons:**
```json
{
    "status": "success",
    "message": "Welcome to Otakudesu API, (THIS IS AN UNOFFICIAL API)",
    "data": null
}
```

---

### Halaman Beranda

**GET** `/home`

Mengambil daftar anime yang ditampilkan di halaman beranda Otakudesu.

**Contoh Permintaan:**
```
GET http://localhost:5001/api/home
```

**Contoh Respons:**
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

---

### Anime Sedang Tayang

**GET** `/ongoing` atau `/ongoing/:page`

Mengambil daftar anime yang sedang tayang. Tanpa parameter halaman, mengembalikan halaman 1.

**Parameter:**
- `page` (opsional): Nomor halaman (integer)

**Contoh Permintaan:**
```
GET http://localhost:5001/api/ongoing
GET http://localhost:5001/api/ongoing/1
```

**Contoh Respons:**
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

---

### Anime Selesai

**GET** `/completed` atau `/completed/:page`

Mengambil daftar anime yang sudah selesai. Tanpa parameter halaman, mengembalikan halaman 1.

**Parameter:**
- `page` (opsional): Nomor halaman (integer)

**Contoh Permintaan:**
```
GET http://localhost:5001/api/completed
GET http://localhost:5001/api/completed/1
```

**Contoh Respons:**
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

---

### Pencarian Anime

**GET** `/search/:query`

Mencari anime berdasarkan kata kunci.

**Parameter:**
- `query` (wajib): Kata kunci pencarian (gunakan URL encode jika mengandung spasi atau karakter khusus)

**Contoh Permintaan:**
```
GET http://localhost:5001/api/search/boku%20dake
```

**Contoh Respons:**
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

---

### Detail Anime

**GET** `/anime/:slug`

Mengambil informasi detail tentang anime tertentu.

**Parameter:**
- `slug` (wajib): Slug unik anime (biasanya bagian terakhir dari URL Otakudesu)

**Contoh Permintaan:**
```
GET http://localhost:5001/api/anime/boku-dake-inai-machi-sub-indo
```

**Contoh Respons:**
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

---

### Detail Episode

**GET** `/episode/:slug`

Mengambil detail episode dan link unduhan.

**Parameter:**
- `slug` (wajib): Slug unik episode

**Contoh Permintaan:**
```
GET http://localhost:5001/api/episode/bokmacih-episode-1-sub-indo
```

**Contoh Respons:**
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
            },
            {
                "quality": "480p",
                "size": "50.2 MB",
                "downloads": [
                    {
                        "provider": "Mega",
                        "downloadUrl": "https://desustream.com/safelink/link/?id=..."
                    }
                ]
            }
        ]
    }
}
```

---

### Daftar Genre

**GET** `/genre-list`

Mengambil semua genre yang tersedia di Otakudesu.

**Contoh Permintaan:**
```
GET http://localhost:5001/api/genre-list
```

**Contoh Respons:**
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

---

### Anime Berdasarkan Genre

**GET** `/genre/:slug` atau `/genre/:slug/:page`

Mengambil daftar anime berdasarkan genre. Tanpa parameter halaman, mengembalikan halaman 1.

**Parameter:**
- `slug` (wajib): Slug genre (dapat diperoleh dari `/genre-list`)
- `page` (opsional): Nomor halaman (integer)

**Contoh Permintaan:**
```
GET http://localhost:5001/api/genre/romance
GET http://localhost:5001/api/genre/romance/2
```

**Contoh Respons:**
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

## Struktur Data

### Item Anime

Digunakan dalam endpoint daftar (`/home`, `/ongoing`, `/completed`, `/search`, `/genre`):

```typescript
{
    title: string,           // Judul anime
    episode: string,         // Info episode (mis., "Episode 8" atau "12 Episode")
    status: string,          // "Ongoing", "Completed", atau kosong
    schedule: string | null, // Hari tayang atau null
    rating: string | null,   // Rating atau null
    date: string,           // Tanggal update/rilis
    slug: string,           // Slug unik anime
    image: string,          // URL gambar thumbnail
    url: string             // URL halaman anime di Otakudesu
}
```

### Detail Anime

Digunakan dalam endpoint `/anime/:slug`:

```typescript
{
    title: string,           // Judul anime
    japaneseTitle: string,   // Judul Jepang
    rating: string,          // Rating
    producer: string,        // Daftar produser
    type: string,           // Jenis anime (mis., "BD", "TV")
    status: string,         // Status (mis., "Completed", "Ongoing")
    episodeTotal: string,   // Total episode
    duration: string,       // Durasi per episode
    releaseDate: string,    // Tanggal rilis
    studio: string,         // Studio animasi
    genre: string,          // Daftar genre
    synopsis: string,       // Sinopsis anime
    episodes: Array<{       // Daftar episode
        episodeTitle: string,
        slug: string,
        url: string
    }>
}
```

### Detail Episode

Digunakan dalam endpoint `/episode/:slug`:

```typescript
{
    downloads: Array<{      // Opsi unduhan
        quality: string,    // Kualitas video (mis., "360p", "480p")
        size: string,       // Ukuran file (mis., "26.5 MB")
        downloads: Array<{  // Penyedia unduhan
            provider: string,    // Nama penyedia (mis., "ZippyShare")
            downloadUrl: string  // URL unduhan
        }>
    }>
}
```

### Daftar Genre

Digunakan dalam endpoint `/genre-list`:

```typescript
{
    genres: Array<{         // Daftar genre
        title: string,      // Nama genre
        slug: string,       // Slug genre
        url: string         // URL halaman genre di Otakudesu
    }>
}
```

## Penanganan Error

### Respons Error Umum

```json
{
    "status": "failed",
    "message": "Deskripsi error",
    "data": null
}
```

### Kode Status HTTP

- **200 OK**: Permintaan berhasil
- **404 Not Found**: Endpoint atau resource tidak ditemukan
- **500 Internal Server Error**: Error server (biasanya masalah scraping)

## Catatan

- Semua URL dalam respons adalah link langsung ke Otakudesu
- Format informasi episode dapat bervariasi (mis., "Episode 8" vs "12 Episode")
- Beberapa field mungkin `null` atau string kosong tergantung ketersediaan
- URL unduhan mungkin berupa link langsung atau safelink
- API ini mengandalkan web scraping, sehingga perubahan struktur di Otakudesu dapat merusak fungsionalitas

---

**Selamat coding! 🚀**