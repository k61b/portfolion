# Portfolion - Portfolio Tracking Application

Portfolion is a portfolio tracking application that allows you to consolidate your savings from different places into a single platform, making it easy to keep track of your investments and savings.

## Table of Contents
- [Introduction](#introduction)
- [Features](#features)
- [Tech Stack](#tech-stack)
- [Installation](#installation)
- [Usage](#usage)
- [Environment Variables](#environment-variables)
- [Contributing](#contributing)
- [License](#license)

## Introduction

Portfolion is designed to provide users with a convenient way to monitor their savings and investments across various platforms. The application is built using React and Golang, making it efficient and scalable.

## Features

- Consolidate savings and investments from different sources
- User authentication and authorization
- View portfolio holdings and performance
- Support for various investment types (e.g., stocks, cryptocurrencies, bonds)
- Real-time data updates from integrated APIs
- Secure and encrypted data storage

## Tech Stack

- Frontend: React, React Router
- Backend: Golang (Go)
- Database: MongoDB
- Authentication: JSON Web Tokens (JWT)
- External APIs: [List any external APIs you integrate with, e.g., AVATAR_API]

## Installation

To set up the Portfolion application on your local machine, follow these steps:

1. Clone the repository:

```bash
git clone https://github.com/kayraberktuncer/portfolion.git
cd portfolion
```

2. Install dependencies for the frontend:

```bash
cd client
npm install
```

3. Create a `.env` file in the `client` folder and provide the necessary environment variable(s):

```env
VITE_PUBLIC_API_URL=YOUR_API_URL
```

4. Install dependencies for the backend (v1):

```bash
cd ../v1
go mod download
```

5. Create a `.env` file in the `v1` folder and provide the necessary environment variable(s):

```env
MONGO_URI=YOUR_MONGODB_CONNECTION_STRING
MONGO_DB=YOUR_MONGODB_DATABASE_NAME
USERS_COLLECTION=YOUR_MONGODB_USERS_COLLECTION
SYMBOLS_COLLECTION=YOUR_MONGODB_SYMBOLS_COLLECTION
PORT=YOUR_BACKEND_PORT
JWT_SECRET=YOUR_JWT_SECRET_KEY
API_KEY=YOUR_EXTERNAL_API_KEY
ALLOWED_ORIGINS=YOUR_ALLOWED_ORIGINS
AVATAR_API=YOUR_AVATAR_API_URL
AVATAR_API_OPTIONS=YOUR_AVATAR_API_OPTIONS
```

## Usage

1. Start the backend server:

```bash
cd v1
go run main.go
```

2. Start the frontend development server:

```bash
cd client
npm run dev
```

3. Open your browser and visit `http://localhost:3000` to access the Portfolion application.

## Environment Variables

Here is a list of environment variables used in the application:

- `VITE_PUBLIC_API_URL`: The public API URL for the frontend.
- `MONGO_URI`: The MongoDB connection string.
- `MONGO_DB`: The MongoDB database name.
- `USERS_COLLECTION`: The MongoDB collection for storing user data.
- `SYMBOLS_COLLECTION`: The MongoDB collection for storing investment symbols/data.
- `PORT`: The backend server port number.
- `JWT_SECRET`: Secret key for JSON Web Tokens.
- `API_KEY`: API key for external API integration.
- `ALLOWED_ORIGINS`: Comma-separated list of allowed origins for CORS.
- `AVATAR_API`: URL for the avatar API service.
- `AVATAR_API_OPTIONS`: Options for the avatar API service.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please feel free to create a pull request.

## License

[MIT](LICENSE)
