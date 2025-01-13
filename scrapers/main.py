from typing import List, Dict
import aiohttp
from bs4 import BeautifulSoup

BASE_URL = "https://otakudesu.cloud"


async def fetch_html(session: aiohttp.ClientSession, url: str) -> BeautifulSoup:
    """Mengambil konten HTML dari URL yang diberikan."""
    async with session.get(url) as response:
        html = await response.text()
        return BeautifulSoup(html, "html.parser")


async def extract_anime_data(anime, status: str) -> Dict:
    """Mengekstrak data anime dari elemen BeautifulSoup."""
    return {
        "title": anime.find("h2").text,
        "episode": anime.find("div", class_="epz").text,
        "schedule": (
            anime.find("div", class_="epztipe").text if status == "Ongoing" else None
        ),
        "rating": (
            anime.find("div", class_="epztipe").text if status == "Completed" else None
        ),
        "date": anime.find("div", class_="newnime").text,
        "slug": anime.find("a")["href"].split("/")[-2],
        "image": anime.find("img")["src"],
        "url": anime.find("a")["href"],
    }


async def scrape_airing_anime(page: int = 1) -> List[Dict]:
    """Mengambil daftar anime yang sedang tayang."""
    url = f"{BASE_URL}/ongoing-anime/page/{page}/"
    async with aiohttp.ClientSession() as session:
        html_soup = await fetch_html(session, url)
        venz = html_soup.find("div", class_="venz")
        anime_list = venz.find_all("div", class_="detpost")
        return [await extract_anime_data(anime, "Ongoing") for anime in anime_list]


async def scrape_completed_anime(page: int = 1) -> List[Dict]:
    """Mengambil daftar anime yang sudah selesai."""
    url = f"{BASE_URL}/complete-anime/page/{page}/"
    async with aiohttp.ClientSession() as session:
        html_soup = await fetch_html(session, url)
        anime_list = html_soup.find_all("div", class_="detpost")
        return [await extract_anime_data(anime, "Completed") for anime in anime_list]


async def scrape_search_anime(query: str) -> List[Dict]:
    """Mengambil daftar anime berdasarkan query pencarian."""
    url = f"{BASE_URL}/?s={query}&post_type=anime"
    async with aiohttp.ClientSession() as session:
        html_soup = await fetch_html(session, url)
        chivsrc = html_soup.find("ul", class_="chivsrc")
        return [
            {
                "title": anime.find("h2").text,
                "status": anime.find_all("div")[1].text.replace("Status : ", ""),
                "rating": anime.find_all("div")[2].text.replace("Rating : ", ""),
                "slug": anime.find("a")["href"].split("/")[-2],
                "image": anime.find("img")["src"],
                "url": anime.find("a")["href"],
            }
            for anime in chivsrc.find_all("li")
        ]


async def scrape_anime(slug: str):
    url = f"https://otakudesu.cloud/anime/{slug}/"

    async with aiohttp.ClientSession() as session:
        async with session.get(url) as response:
            html_soup = BeautifulSoup(await response.text(), "html.parser")
            venser = html_soup.find("div", class_="venser")

            # Extract basic information
            info_texts = [p.text.split(": ", 1)[1] for p in venser.find_all("p")[1:11]]
            keys = [
                "japanese_title",
                "rating",
                "producer",
                "type",
                "status",
                "episode_total",
                "duration",
                "release_date",
                "studio",
                "genre",
            ]
            data = dict(zip(keys, info_texts))

            data.update(
                {
                    "title": venser.find("h1").text,
                    "image": venser.find("img")["src"],
                }
            )

            # Extract episode list
            episode_list = html_soup.find_all("ul")[3]
            data["episode_list"] = [
                {
                    "episode": ep.find("a").text,
                    "slug": ep.find("a")["href"].split("/")[-2],
                    "url": ep.find("a")["href"],
                }
                for ep in episode_list.find_all("li")
            ]

            # Extract batch if completed
            if data["status"] == "Completed":
                batch_list = venser.find_all("ul")[2]
                data["batch"] = batch_list.find("a")["href"]

            return data


async def scrape_anime_episodes(slug: str) -> Dict:
    """Mengambil daftar episode dari anime tertentu."""
    # Fungsi ini tampaknya redundan dengan scrape_anime().
    # Sebaiknya gunakan scrape_anime() dan ambil data episode dari sana.
    return (await scrape_anime(slug))["episode_list"]


async def scrape_single_episode(slug: str) -> Dict:
    """Mengambil detail dan link download dari episode tertentu."""
    url = f"{BASE_URL}/episode/{slug}/"
    async with aiohttp.ClientSession() as session:
        html_soup = await fetch_html(session, url)
        download_section = html_soup.find("div", class_="download")
        episode_title = download_section.find("h4").text
        video_quality = download_section.find("ul")

        data = {
            "episode_title": episode_title,
            "download_links": [
                {
                    quality.find("strong").text: {
                        link.text: link["href"] for link in quality.find_all("a")
                    }
                }
                for quality in video_quality.find_all("li")
            ],
        }
        return data


async def scrape_genres_list() -> Dict:
    """Mengambil daftar genre dan URL-nya."""
    data = []
    url = f"{BASE_URL}/genre-list/"
    async with aiohttp.ClientSession() as session:
        html_soup = await fetch_html(session, url)
        genres_section = html_soup.find("ul", class_="genres")
        genres = genres_section.find("li")

        for genre in genres.find_all("a"):
            data.append(
                {
                    "name": genre.text,
                    "slug": genre["href"].split("/")[-2],
                    "url": BASE_URL + genre["href"],
                }
            )
    return data


async def scrape_single_genre(slug: str, page: int = 1) -> Dict:
    """Mengambil detail dan link download dari episode tertentu."""
    data = []
    url = f"{BASE_URL}/genres/{slug}/page/{page}/"
    async with aiohttp.ClientSession() as session:
        html_soup = await fetch_html(session, url)
        venser = html_soup.find("div", class_="venser")
        for anime in venser.find_all("div", class_="col-anime"):
            data.append(
                {
                    "title": anime.find("div", class_="col-anime-title").text,
                    "slug": anime.find("a")["href"].split("/")[-2],
                    "image": anime.find("img")["src"],
                    "rating": anime.find("div", class_="col-anime-rating").text,
                    "episode": anime.find("div", class_="col-anime-eps").text,
                    "season": anime.find("div", class_="col-anime-date").text,
                    "studio": anime.find("div", class_="col-anime-studio").text,
                    "url": anime.find("a")["href"],
                }
            )

    return data


if __name__ == "__main__":
    import asyncio

    data = asyncio.run(scrape_anime_episodes("salaryman-isekai-sub-indo"))
    print(data)
