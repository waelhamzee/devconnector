# DevConnector

DevConnector is a minimal backend social network application for developers, built using Go (Golang).

## Features

- User authentication with JWT
- Create and manage developer profiles
- Post, like, and comment on posts
- Connect with other developers

## Getting Started

### Prerequisites

- Go (Golang)
- MongoDB

### Installation

```bash
git clone https://github.com/yourusername/devconnector.git
cd devconnector
go mod download
```

### Environment Variables

Create a `.env` file in the root directory and add:

```
JWT_SECRET=your_jwt_secret
```

### Running the App

```bash
go run main.go
```

The API runs on [http://localhost:8080](http://localhost:8080).