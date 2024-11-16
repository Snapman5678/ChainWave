# ChainWave

**ChainWave** is a full-stack supply chain management platform built with Next.js and Go, featuring real-time collaboration and secure user management. It is designed to optimize logistics, streamline operations, and enhance resource allocation. The platform connects businesses, suppliers, and consumers in a dynamic network, ensuring real-time transparency and efficiency.

## Tech Stack

- **Frontend**: Next.js, TypeScript, Tailwind CSS
- **Backend**: Go, Gin
- **Database**: PostgreSQL
- **Authentication**: JWT
- **Containerization**: Docker, Docker Compose

## Features

- User authentication and authorization
- Role-based access control
- Profile management
- Real-time updates
- Responsive design
- Containerized development environment

## Prerequisites

- Docker and Docker Compose
- Node.js (for local development)
- Go 1.21+ (for local development)
- PostgreSQL (for local development)

## Quick Start

1. **Clone the repository**
```bash
git clone https://github.com/Snapman5678/ChainWave.git
cd chainwave
```

2. **Environment Setup**

Create a .env file in the root directory:
```env
POSTGRES_USER=your_user
POSTGRES_PASSWORD=your_password
POSTGRES_DB=chainwave
```

3. **Run with Docker Compose**
```bash
docker-compose up --build
```

The application will be available at:
- Frontend: http://localhost:3001
- Backend API: http://localhost:8000

## Local Development

### Frontend

```bash
cd frontend
npm install
npm run dev
```

### Backend

```bash
cd backend
go mod download
go run cmd/main.go
```

## Project Structure

```
├── frontend/                # Next.js frontend
│   ├── src/
│   ├── public/
│   └── next.config.mjs
├── backend/                 # Go backend
│   ├── cmd/
│   ├── internal/
│   └── go.mod
├── docker-compose.yml
└── README.md
```

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contact

Aamir Mohammed - [@Snapman5678](https://github.com/Snapman5678)
A Arshad Khan - [@ArshdKhan](https://github.com/ArshdKhan/)

Project Link: [https://github.com/Snapman5678/ChainWave](https://github.com/Snapman5678/ChainWave)