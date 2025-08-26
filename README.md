# ClickBait.ai 🚀

*A satirical URL shortener that transforms boring links into irresistibly clickable tech buzzword masterpieces*

## 🎯 Project Overview

ClickBait.ai is a full-stack web application that shortens URLs while wrapping them in gloriously over-the-top tech industry buzzwords. Built as a learning project to demonstrate modern web development skills and best practices.

**Live Demo**: `your-domain.com/make-money-fast/abc123` → redirects to your boring corporate website

## ✨ Features

- **Dynamic URL Shortening**: Generate unique short codes with cryptographically secure randomization
- **Satirical Branding**: Transforms mundane links into "EXPONENTIALLY DISRUPTIVE" clickbait
- **Expiration Management**: Configurable URL lifespans with automatic cleanup
- **Responsive Design**: Mobile-first UI built with Tailwind CSS
- **Real-time Feedback**: Loading states, error handling, and copy-to-clipboard functionality
- **Database Persistence**: PostgreSQL storage with proper indexing and relationships

## 🛠 Technical Stack

### Frontend
- **React 18** with TypeScript
- **Tailwind CSS** for responsive styling
- **Vite** for fast development and building
- **Lucide React** for consistent iconography

### Backend
- **Go 1.23** with clean architecture patterns
- **Chi Router** for HTTP routing and middleware
- **PostgreSQL** with pgx driver for database operations
- **Environment-based configuration**

### Architecture Highlights
- **Component-based frontend** with proper prop typing
- **RESTful API design** with structured error handling
- **Database connection pooling** and prepared statements
- **CORS middleware** for cross-origin requests
- **Graceful error handling** throughout the stack

## 📁 Project Structure

```
shrimpy/
├── client/                 # React frontend
│   ├── src/
│   │   ├── components/     # Reusable UI components
│   │   └── App.tsx        # Main application logic
│   └── package.json
└── server/                # Go backend
    ├── cmd/shrimpy/       # Application entry point
    ├── internal/
    │   ├── handlers/      # HTTP request handlers
    │   ├── shortener/     # URL shortening logic
    │   └── storage/       # Database operations
    └── go.mod
```

## 🔧 Key Technical Implementations

### Secure Random Key Generation
- Cryptographically secure random string generation using `crypto/rand`
- Base62 encoding for URL-safe short codes
- Collision detection and retry logic

### Database Design
```sql
CREATE TABLE urls (
    id SERIAL PRIMARY KEY,
    old_url TEXT NOT NULL,
    new_url VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    expires_at TIMESTAMP NOT NULL
);
```

### API Endpoints
- `POST /shorten` - Create shortened URL with expiration
- `GET /make-money-fast/{key}` - Redirect to original URL
- `GET /preview/{key}` - Preview URL details without redirect

### Error Handling
- Comprehensive input validation
- Graceful database error recovery
- User-friendly error messages
- Proper HTTP status codes

## 🎨 UI/UX Features

- **Animated Background**: CSS gradient animations with floating elements
- **Progressive Enhancement**: Works without JavaScript for basic functionality
- **Accessibility**: Proper ARIA labels and keyboard navigation
- **Mobile Optimization**: Touch-friendly buttons and responsive breakpoints

## 🚀 Development Learnings

This project demonstrates proficiency in:

- **Full-stack development** with modern toolchains
- **Database design** and optimization techniques  
- **RESTful API** architecture and implementation
- **Frontend state management** and component composition
- **Error handling** and user experience considerations
- **Security best practices** for web applications
- **Environment configuration** and deployment preparation

## 🎭 The Satire

Beyond the technical implementation, this project pokes fun at:
- Tech industry buzzword culture
- Clickbait marketing tactics
- Over-engineered "AI-powered" solutions
- The startup ecosystem's tendency toward hyperbole

*Because sometimes the best way to learn serious development skills is through a completely unserious project.*

---

**Tech Stack**: React • TypeScript • Go • PostgreSQL • Tailwind CSS  
**Focus Areas**: Full-stack development • Database design • API architecture •