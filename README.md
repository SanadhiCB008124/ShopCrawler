# ShopCrawler

**ShopCrawler** is a  product aggregator that crawls and scrapes data from multiple e-commerce websites to provide real-time product comparisons. It helps users find the best deals, track price changes, and make informed shopping decisions.

## Features

- **Multi-Site Crawling:** Discover products from various e-commerce platforms.
- **Real-Time Scraping:** Extract product details like name, price, image, and description.
- **Price Comparison:** Compare prices of similar products across different sites.
- **Search & Filter:** Advanced search options with filters for price, category, and more.
- **Data Deduplication:** Avoid duplicate products when crawling overlapping sources.
- **Responsive Dashboard:** Clean UI for browsing, comparing, and analyzing products.
- **Scheduled Crawling:** Automate data collection at regular intervals.

## Technologies Used

- **Backend:** Go (Golang) with [Colly](https://github.com/gocolly/colly) for crawling and scraping
- **Database:** MongoDB for flexible product data storage
- **API:** RESTful API built with Go for data access
- **Frontend:** Angular for a dynamic and responsive user interface
- **Authentication:** JWT (JSON Web Tokens) for secure user authentication
- **Deployment:** Docker for containerization and easy deployment

## Setting Up

1. **Clone the Repository:**
   ```bash
   git clone https://github.com/SanadhiCB008124/shopcrawler.git
   cd shopcrawler
   ```

2. **Backend Setup:**
   ```bash
   cd backend
   go mod tidy
   go run main.go
   ```

3. **Frontend Setup:**
   ```bash
   cd frontend
   npm install
   ng serve
   ```

4. **Database:**
   - Ensure MongoDB is running locally or provide your cloud connection string.

## API Endpoints

- `GET /products` - Fetch all products
- `GET /products/:id` - Get details of a specific product
- `POST /crawl` - Trigger a new crawl session
- `GET /compare?productName=xyz` - Compare product prices across sites

## Roadmap

- Add user accounts with saved product tracking
- Implement price drop alerts via email or SMS
- Support more e-commerce platforms
- Introduce AI-powered product recommendations



  
