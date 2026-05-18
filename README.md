setting for backend local
1. make folder storage/documents on permit-system-be
2. create database on local computer
3. go mod tidy
4. create .env 

APP_PORT=

DB_HOST=
DB_PORT=
DB_USER=
DB_PASSWORD=
DB_NAME=

JWT_SECRET=

MAIL_HOST=
MAIL_PORT=
MAIL_USERNAME=
MAIL_PASSWORD= #must app password#
MAIL_FROM=
APP_URL=

if step 1 and step 2 finish, run cmd/main.go, database auto migrate

setting frontend 
1. npm install
2. after node_module finish install for run in local, running "npm run dev"
