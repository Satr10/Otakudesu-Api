from fastapi import FastAPI, HTTPException
from typing import Optional

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
        return {
            "status": "Ok",
            "data": {
                "airing": {"0": "Title", "1": "Title"},
                "completed": {"0": "Title", "1": "Title"},
            },
        }
    except Exception as e:
        return {"status": "Error", "message": str(e)}


@app.get("/ongoing-anime")
async def ongoing_anime(page: Optional[int] = None):
    if page is None:
        return {"status": "Error", "message": "Parameter 'page' dibutuhkan."}
    try:
        return {
            "status": "Ok",
            "data": {"0": "Title", "1": "Title"},
        }
    except Exception as e:
        return {"status": "Error", "message": str(e)}


@app.get("/completed-anime")
async def completed_anime(page: Optional[int] = None):
    if page is None:
        return {"status": "Error", "message": "Parameter 'page' dibutuhkan."}
    try:
        return {
            "status": "Ok",
            "data": {"0": "Title", "1": "Title"},
        }
    except Exception as e:
        return {"status": "Error", "message": str(e)}


@app.get("/search")
async def search(query: str):
    try:
        return {
            "status": "Ok",
            "data": {"Your Search: ": query},
        }
    except Exception as e:
        return {"status": "Error", "message": str(e)}


@app.get("/anime/{slug}")
async def anime(slug: str):
    try:
        return {
            "status": "Ok",
            "data": {"Anime: ": "Anime"},
        }
    except Exception as e:
        return {"status": "Error", "message": str(e)}


@app.get("/anime/{slug}/episodes")
async def episodes(slug: str):
    try:
        return {
            "status": "Ok",
            "data": {"Episode: ": "Episode"},
        }
    except Exception as e:
        return {"status": "Error", "message": str(e)}


@app.get("/anime/{slug}/episodes/{episode}")
async def episode(slug: str, episode: str):
    try:
        return {
            "status": "Ok",
            "data": {"Episode: ": "Episode"},
        }
    except Exception as e:
        return {"status": "Error", "message": str(e)}


@app.get("/v1/episode/{episode}")
async def episode(episode: str):
    try:
        return {
            "status": "Ok",
            "data": {"Episode: ": "Episode"},
        }
    except Exception as e:
        return {"status": "Error", "message": str(e)}


@app.get("/genres")
def genres():
    try:
        return {
            "status": "Ok",
            "data": {"Genre: ": "Genre"},
        }
    except Exception as e:
        return {"status": "Error", "message": str(e)}


@app.get("/genre/{genre}")
async def genre(genre: str):
    try:
        return {
            "status": "Ok",
            "data": {"Genre: ": "Genre"},
        }
    except Exception as e:
        return {"status": "Error", "message": str(e)}


if __name__ == "__main__":
    import uvicorn

    uvicorn.run("main:app", host="0.0.0.0", port=8000, reload=True)
