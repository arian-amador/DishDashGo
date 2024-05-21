# Search
- This is the recipe search page. This URL will return all results and show breadcrumb of next page. 
-- First: https://icecreamfromscratch.com/?s=
-- Next: https://icecreamfromscratch.com/page/2/?s=
-- Last: https://icecreamfromscratch.com/page/23/?s

# Feeds URLs:
- Content: https://icecreamfromscratch.com/feed/
- Comments: https://icecreamfromscratch.com/comments/feed/
- RSS: https://icecreamfromscratch.com/search/feed/rss2/

# Pipeline
- Get page data then parse it.

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