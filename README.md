# BackgroundCrons

BackgroundCrons is a Go application that fetches weather data from an API and stores it in a MongoDB database using cron scheduling.

## Getting Started

### Prerequisites
- Go (1.15+ recommended)
- MongoDB (Make sure it's running and accessible)
- Git

### Installation

-git clone https://github.com/plainsurf-solution/BackgroundCrons.git
-cd BackgroundCrons
-go mod download

## Usage

### Update MongoDB connection parameters and API endpoint in main.go
go run main.go

## Endpoints 
- http://localhost:8080/get - getting weather info
