# Unsplash API

Basic API connection written in Go.

Create a config.toml and add your client ID.

To search for a specific term, add a query flag, for example `-query="mountain"`.

## Available flags

| key       | Value  | Example    | Required |
|-----------|--------|------------|----------|
| -query    | String | "mountain" | Yes      |
| -per_page | String | "3"        | No       |
