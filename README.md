# Artpaper
REST API of site with articles
## Stack
- Go with [gin](https://github.com/gin-gonic/gin)
- PostgreSQL
- Redis
- Docker
## Running
### Set environment variables in .env file
```bash
GIN_MODE=debug
APP_PORT=8080
JWT_SIGN_KEY=
PASSWORD_SALT=

POSTGRES_HOST=db
POSTGRES_PORT=5432
POSTGRES_DB=
POSTGRES_USER=
POSTGRES_PASSWORD=

TOKEN_STORAGE_HOST=token_storage
TOKEN_STORAGE_PORT=6379
TOKEN_STORAGE_PASSWORD=

CACHE_HOST=cache
CACHE_PORT=6379
CACHE_PASSWORD=
```
### Clone
```bash
git clone https://github.com/gelerum/artpaper.git
```
### Build containers
```bash
cd artpaper
docker-compose build
```
### Up containers
```bash
docker-compose up
```
## Documentation
For documentation open `0.0.0.0:8080/swagger/index.html` in browser while running.
## License
Usage is provided under the [MIT License](https://opensource.org/licenses/mit-license.php). See LICENSE for the full details.
