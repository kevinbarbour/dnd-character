# D&D Character Creator - Vercel Deployment Guide

This guide will help you deploy the D&D Character Creator to Vercel with PostgreSQL database.

## Prerequisites

1. **Vercel Account**: Sign up at [vercel.com](https://vercel.com)
2. **Vercel CLI**: Install with `npm i -g vercel`
3. **Git Repository**: Push your code to GitHub/GitLab/Bitbucket

## Step 1: Prepare Your Project

The project has been restructured for Vercel deployment:

```
/
├── api/                     # Vercel serverless functions
│   ├── index.go            # Home page
│   ├── characters/
│   │   ├── index.go        # List characters
│   │   └── new.go          # Create character
│   ├── generate-abilities.go
│   └── spells.go
├── public/                  # Static assets
│   ├── static/             # CSS, JS files
│   └── templates/          # HTML templates
├── internal/               # Go packages
├── vercel.json            # Vercel configuration
└── go.mod
```

## Step 2: Create Vercel Postgres Database

1. **Login to Vercel Dashboard**
   ```bash
   vercel login
   ```

2. **Create New Project**
   ```bash
   vercel
   ```
   - Link to your Git repository
   - Choose project name (e.g., `dnd-character-creator`)

3. **Add Postgres Database**
   - Go to your project dashboard on vercel.com
   - Navigate to **Storage** tab
   - Click **Create Database**
   - Select **Postgres**
   - Choose a database name (e.g., `dnd-characters`)
   - Select region closest to your users

4. **Copy Environment Variables**
   - After database creation, copy these environment variables:
     - `POSTGRES_URL`
     - `POSTGRES_PRISMA_URL` 
     - `POSTGRES_URL_NON_POOLING`

## Step 3: Configure Environment Variables

1. **In Vercel Dashboard:**
   - Go to **Settings** → **Environment Variables**
   - Add the PostgreSQL environment variables from Step 2
   - Make sure they're available for **Production**, **Preview**, and **Development**

2. **Required Environment Variables:**
   ```
   POSTGRES_URL=postgresql://username:password@host:port/database
   POSTGRES_PRISMA_URL=postgresql://username:password@host:port/database?pgbouncer=true&connect_timeout=15
   POSTGRES_URL_NON_POOLING=postgresql://username:password@host:port/database
   ```

## Step 4: Deploy

1. **Deploy to Production**
   ```bash
   vercel --prod
   ```

2. **Or Deploy via Git**
   - Push your code to your Git repository
   - Vercel will automatically deploy on every push to main branch

## Step 5: Verify Deployment

1. **Check Your Live Site**
   - Visit your Vercel URL (e.g., `https://your-project.vercel.app`)
   - Test character creation functionality
   - Verify spell selection works for spellcasting classes

2. **Monitor Database**
   - Check Vercel dashboard for database metrics
   - Verify tables are created automatically on first visit

## Features Available

✅ **Complete D&D 5e Character Creator**
✅ **Spell Selection System** (6 spellcasting classes)
✅ **Character Persistence** with PostgreSQL
✅ **Beautiful Web Interface**
✅ **Responsive Design**
✅ **Global CDN** via Vercel
✅ **Automatic HTTPS**

## API Endpoints

- `GET /` - Home page
- `GET /characters` - List all characters
- `GET /characters/new` - Character creation form
- `POST /characters/new` - Create new character
- `GET /characters/{id}` - View character details
- `POST /api/generate-abilities` - Generate ability scores
- `GET /api/spells?class={className}` - Get spells for class

## Troubleshooting

### Database Connection Issues
- Verify environment variables are set correctly
- Check database URL format
- Ensure database is in same region as functions

### Template Loading Issues
- Verify templates are in `public/templates/` directory
- Check file paths in API handlers

### Build Errors
- Run `go mod tidy` to ensure dependencies are correct
- Check Go version compatibility (requires Go 1.19+)

### Function Timeout
- Vercel functions have 10-second timeout on Hobby plan
- Consider upgrading to Pro plan for longer timeouts

## Custom Domain (Optional)

1. **Add Custom Domain**
   - Go to **Settings** → **Domains**
   - Add your domain name
   - Configure DNS records as instructed

2. **SSL Certificate**
   - Automatically provided by Vercel
   - No additional configuration needed

## Monitoring

- **Analytics**: Available in Vercel dashboard
- **Function Logs**: Real-time logs for debugging
- **Performance**: Core Web Vitals tracking
- **Database Metrics**: Query performance and usage

Your D&D Character Creator is now live on Vercel! 🎉

## Support

For issues with:
- **Vercel Platform**: Check [Vercel Documentation](https://vercel.com/docs)
- **PostgreSQL**: Check [Vercel Postgres Docs](https://vercel.com/docs/storage/vercel-postgres)
- **Application**: Check the application logs in Vercel dashboard
