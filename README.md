# News API

## Overview
This API provides access to news stories from Hacker News. It allows you to fetch top, new, and best stories, as well as search for stories based on keywords.

## Base URL
http://localhost:3000


## Endpoints

### Get Top/New/Best Stories
- **URL:** /api/stories
- **Method:** GET
- **Query Parameters:**
  - category (optional, default: top) - Specify the category of stories (top, new, best)
  - page (optional, default: 1) - Specify the page number for pagination
  - limit (optional, default: 10) - Specify the number of stories per page
- **Description:** Fetches the top, new, or best stories from Hacker News.

### Search Stories
- **URL:** /api/search
- **Method:** GET
- **Query Parameters:**
  - keyword (required) - The keyword to search for in stories
- **Description:** Searches for stories containing the specified keyword.

### Get User's Bookmarks
- **URL:** /api/bookmarks
- **Method:** GET
- **Description:** Fetches the stories and comments bookmarked by the user.

### Bookmark a Story or Comment
- **URL:** /api/setbookmarks
- **Method:** POST
- **Request Body:**
  ```json
  {
    "item_id": 123,
  }
