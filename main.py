from fastapi import FastAPI, HTTPException
from typing import Optional
from scrapers.main import *
import asyncio

app = FastAPI()


@app.get("/")
async def root():
    try:
        return {"status": "Ok", "message": "Unofficial Otakudesu API made with fastAPI"}
    except Exception as e:
        return {"status": "Error", "message": str(e)}


@app.get("/home")
async def home():
    try:
        airing, completed = await asyncio.gather(
            scrape_airing_anime(), scrape_completed_anime()
        )
        return {
            "status": "Ok",
            "data": {
                "airing": [
                    {
                        "title": anime["title"],
                        "episode": anime["episode"],
                        "schedule": anime["schedule"],
                        "date": anime["date"],
                        "slug": anime["slug"],
                        "url": anime["url"],
                    }
                    for anime in airing
                ],
                "completed": [
                    {
                        "title": anime["title"],
                        "episode": anime["episode"],
                        "rating": anime["rating"],
                        "date": anime["date"],
                        "slug": anime["slug"],
                        "url": anime["url"],
                    }
                    for anime in completed
                ],
            },
        }
    except Exception as e:
        return {"status": "Error", "message": str(e)}


@app.get("/ongoing-anime")
async def ongoing_anime(page: Optional[int] = 1):
    try:
        airing = await scrape_airing_anime(page)
        return {
            "status": "Ok",
            "data": [
                {
                    "title": anime["title"],
                    "episode": anime["episode"],
                    "schedule": anime["schedule"],
                    "date": anime["date"],
                    "slug": anime["slug"],
                    "image": anime["image"],
                    "url": anime["url"],
                }
                for anime in airing
            ],
        }
    except Exception as e:
        return {"status": "Error", "message": str(e)}


@app.get("/completed-anime")
async def completed_anime(page: Optional[int] = 1):
    try:
        completed = await scrape_completed_anime(page)
        return {
            "status": "Ok",
            "data": [
                {
                    "title": anime["title"],
                    "episode": anime["episode"],
                    "rating": anime["rating"],
                    "date": anime["date"],
                    "slug": anime["slug"],
                    "image": anime["image"],
                    "url": anime["url"],
                }
                for anime in completed
            ],
        }
    except Exception as e:
        return {"status": "Error", "message": str(e)}


@app.get("/search")
async def search(query: str):
    try:
        result = await scrape_search_anime(query)
        return {
            "status": "Ok",
            "data": [
                {
                    "title": anime["title"],
                    "status": anime["status"],
                    "rating": anime["rating"],
                    "slug": anime["slug"],
                    "image": anime["image"],
                    "url": anime["url"],
                }
                for anime in result
            ],
        }
    except Exception as e:
        return {"status": "Error", "message": str(e)}


@app.get("/anime/{slug}")
async def anime(slug: str):
    try:
        data = await scrape_anime(slug)
        return {
            "status": "Ok",
            "data": {
                "title": data["title"],
                "image": data["image"],
                "japanese_title": data["japanese_title"],
                "rating": data["rating"],
                "producer": data["producer"],
                "type": data["type"],
                "status": data["status"],
                "episode_total": data["episode_total"],
                "duration": data["duration"],
                "release_date": data["release_date"],
                "studio": data["studio"],
                "genre": data["genre"],
                "episode_list": data["episode_list"],
                "batch": data["batch"],
            },
        }
    except Exception as e:
        return {"status": "Error", "message": str(e)}


@app.get("/anime/{slug}/episodes")
async def episodes(slug: str):
    try:
        data = await scrape_anime(slug)
        return {
            "status": "Ok",
            "data": {
                "episode_list": data["episode_list"],
            },
        }
    except Exception as e:
        return {"status": "Error", "message": str(e)}


# tidak bisa diaplikasikan karna url setiap episode berbeda
# @app.get("/anime/{slug}/episodes/{episode}")
# async def episode(slug: str, episode: str):
#     try:
#         return {
#             "status": "Ok",
#             "data": {"Episode: ": "Episode"},
#         }
#     except Exception as e:
#         return {"status": "Error", "message": str(e)}


@app.get("/episode/{episode_slug}")
async def episode(episode_slug: str):
    try:
        data = await scrape_single_episode(episode_slug)
        return {
            "status": "Ok",
            "data": data,
        }
    except Exception as e:
        return {"status": "Error", "message": str(e)}


@app.get("/genres")
async def genres():
    try:
        data = await scrape_genres_list()
        return {
            "status": "Ok",
            "data": {"data: ": data},
        }
    except Exception as e:
        return {"status": "Error", "message": str(e)}


@app.get("/genres/{slug}")
async def genre(slug: str):
    try:
        data = await scrape_single_genre(slug)
        return {
            "status": "Ok",
            "data": {"Genre: ": slug, "anime": data},
        }
    except Exception as e:
        return {"status": "Error", "message": str(e)}


if __name__ == "__main__":
    import uvicorn

    uvicorn.run("main:app", host="0.0.0.0", port=8000, reload=True)
