# Search
- This is the recipe search page. This URL will return all results and show breadcrumb of next page. 
  - First: https://icecreamfromscratch.com/?s=
  - Next: https://icecreamfromscratch.com/page/2/?s=
  - Last: https://icecreamfromscratch.com/page/23/?s

# Feeds URLs:
- Content: https://icecreamfromscratch.com/feed/
- Comments: https://icecreamfromscratch.com/comments/feed/
- RSS: https://icecreamfromscratch.com/search/feed/rss2/

# Pipeline
- Get page data then parse it.
## Get
- RSS: This should be used to get the latest recipes from the page daily. 
  - I think all we need are the links to the page themselves. 
- Search: This should be used to get all old the recipes from the site that are not shown in the RSS feed.
  - Process pages using Go routines: https://medium.com/@arnesh07/how-golang-can-save-you-days-of-web-scraping-72f019a6de87

## Parse
- RSS: Feed parse to get page links.
- Page: 
  - Store: raw rendered page and convert it to plain text with all tags removed/simplified.
  - Parse: Parse to generate the recipe object. I think the easiest way to do this is to use a library like BeautifulSoup to parse the HTML and extract the data we need. Or just make our own targeting the HTML tags (div > id, ul, etc.) we deem useful.
    - https://github.com/anaskhan96/soup

# Datastructure
## Notes: (T: Thought, Q: Question)
- T: Maybe store the raw page content in the database for future reference.
- T: Some parts of the page are not needed, such as the extended descriptions or further explanation of certain ingredients. Maybe these parts can be condensed or removed.
- Q: Is there a way to get the page as plain text? This would make parsing easier. The HTML tags are not needed for the data we are looking for.
```
{
    name:,
    imageUrls:[],
    description:,
    feeds:,
    ingredients: [
        {
            name: "water",
            quantity: 2,
            unit: cups,
            description: "Even though watermelon is mostly water..."
        }
    ]
    instructions: [
        {
            step: <maybe we just store this as an array and infer the number from the order of the array. So all we would need is an array of strings.>,
            description:,
        }
    ]
}
```