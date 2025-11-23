# Sales Management Frontend

A modern Next.js frontend for the Sales Management API with Tailwind CSS and shadcn/ui.

## Features

- **User Management**: Complete CRUD operations for users with access levels
- **Item Management**: Inventory management with stock tracking
- **Dashboard**: Overview of system statistics
- **Responsive Design**: Mobile-friendly interface
- **Modern UI**: Built with shadcn/ui components

## Tech Stack

- **Next.js 14** - React framework
- **TypeScript** - Type safety
- **Tailwind CSS** - Styling
- **shadcn/ui** - UI components
- **Axios** - API client
- **Lucide React** - Icons

## Getting Started

### Development

1. **Install dependencies**
   ```bash
   npm install
   ```

2. **Set environment variables**
   ```bash
   cp .env.example .env
   # Edit NEXT_PUBLIC_API_URL if needed
   ```

3. **Run development server**
   ```bash
   npm run dev
   ```

4. **Open browser**
   Navigate to [http://localhost:3001](http://localhost:3001)

### Production Build

```bash
npm run build
npm start
```

### Docker

```bash
# Build image
docker build -t sales-web .

# Run container
docker run -p 3001:3001 -e NEXT_PUBLIC_API_URL=http://localhost:3000 sales-web
```

## Pages

- `/` - Dashboard with system statistics
- `/users` - User management (list, create, edit, delete)
- `/items` - Item management (list, create, edit, delete)

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| NEXT_PUBLIC_API_URL | Backend API URL | http://localhost:3000 |

## Project Structure

```
web/
├── app/                 # Next.js app directory
│   ├── layout.tsx      # Root layout
│   ├── page.tsx        # Dashboard
│   ├── users/          # Users page
│   └── items/          # Items page
├── components/         # React components
│   ├── ui/            # shadcn/ui components
│   ├── users/         # User-specific components
│   └── items/         # Item-specific components
├── lib/               # Utility functions
├── services/          # API services
├── types/             # TypeScript types
└── public/            # Static files
```

## API Integration

The frontend connects to the Go Fiber backend API. Make sure the backend is running before starting the frontend.

Default API endpoint: `http://localhost:3000/api/v1`

## Development Notes

- The app uses Next.js App Router
- All pages are client-side rendered (`"use client"`)
- API calls are made using Axios
- Form validation is handled in the components
- Toast notifications for user feedback
