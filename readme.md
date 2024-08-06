# Getting started 

## Installation

### Tailwind CSS
Make sure you have installed Tailwind CSS. 
I recommand using the standalone version

`https://tailwindcss.com/blog/standalone-cli`

### Go dependencies
Install all required dependencies by running the following command:
`go mod tidy .`

install air base on your environment
`https://github.com/air-verse/air`

install sqlc base on your environment
`https://docs.sqlc.dev/en/stable/overview/install.html`

install templ base on your environment
`go install github.com/a-h/templ/cmd/templ@latest`

## Setting Up
create a log file 
`./logs/app.log`

create a .env file and configure your port and db connection string
`./.env`

### MAKEFILE
The project uses a makefile to run the server and the client.
To run the server, use the following command:

`make dev`

# Enjoy!
