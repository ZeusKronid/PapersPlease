# PapersPlease
### Important – Still under development, check back soon for updates.

Are you tired of sifting through hundreds of research papers to find the real gems? PapersPlease is the perfect solution for you! This command-line interface (CLI) helps you stay up-to-date with cutting-edge scientific research in your area of interest with minimal effort.

Powered by an API that analyzes tens of thousands of papers each week, PapersPlease helps you discover the best research based on a variety of metrics. You'll always be in the know, effortlessly.


## Features:

    Discover top research papers with just a few commands.
    Stay updated with the latest research in your field.
    Choose from different sorting methods to find papers that matter to you.
    Explore the scientific frontier without wasting time on irrelevant papers.

## Key Notes:

    Metadata is kept for one week only. Historical data is not accessible.
    Designed to help you stay up-to-date with the latest research.
    Focused on providing the most impactful papers, tailored to your needs.

## Use Cases:

    **paperspls ML -w 10 -m hi**: Shows the top 10 papers in Machine Learning based on the authors' overall h-index.
    **paperspls**: Displays top papers from all areas of research.
    **paperspls -d "Physics implementation in Biology"**: Returns top papers based on your text description.

## Available Flags:

    -w / --window: Defines the window size (default: 10, max: 100).
    -s / --sortby: Specifies the sorting type.
    -m / --method: Selects the method for finding best papers (default: oi for overall impact).

