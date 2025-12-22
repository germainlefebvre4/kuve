# Kuve Documentation

This directory contains the Docusaurus-based documentation for Kuve.

## Prerequisites

- Node.js 18.0 or higher
- npm or yarn

## Installation

```bash
npm install
```

## Local Development

```bash
npm start
```

This command starts a local development server and opens up a browser window. Most changes are reflected live without having to restart the server.

## Build

```bash
npm run build
```

This command generates static content into the `build` directory and can be served using any static contents hosting service.

## Deployment

### GitHub Pages

```bash
npm run deploy
```

This command builds the website and pushes to the `gh-pages` branch.

### Manual Deployment

```bash
npm run build
# Then deploy the 'build' directory to your hosting service
```

## Documentation Structure

```
docs/
├── intro.md                    # Introduction
├── getting-started/            # Getting Started guides
│   ├── installation.md
│   ├── quickstart.md
│   └── shell-setup.md
├── user-guide/                 # User guides
│   ├── basic-usage.md
│   ├── managing-versions.md
│   ├── version-files.md
│   └── workflows.md
├── advanced/                   # Advanced features
│   ├── cluster-detection.md
│   ├── shell-integration.md
│   └── version-normalization.md
├── reference/                  # Reference documentation
│   ├── cli.md
│   ├── commands.md
│   ├── configuration.md
│   └── troubleshooting.md
└── developers/                 # Developer documentation
    ├── architecture.md
    ├── contributing.md
    └── development.md
```

## Configuration

### Site Configuration

Edit `docusaurus.config.js` to:
- Change site title and tagline
- Update URL and base path
- Configure GitHub integration
- Modify theme settings
- Add plugins

### Sidebar Configuration

Edit `sidebars.js` to:
- Organize documentation structure
- Add/remove sections
- Reorder pages

### Custom CSS

Edit `src/css/custom.css` to customize:
- Color scheme
- Fonts
- Component styles

## Writing Documentation

### Frontmatter

Add frontmatter to the top of each markdown file:

```markdown
---
sidebar_position: 1
title: Page Title
---

# Page Content
```

### Admonitions

Use admonitions for notes, tips, warnings:

```markdown
:::note
This is a note
:::

:::tip
This is a tip
:::

:::warning
This is a warning
:::

:::danger
This is danger
:::
```

### Code Blocks

Use syntax highlighting:

````markdown
```bash
kuve install v1.28.0
```

```go
func main() {
    fmt.Println("Hello")
}
```
````

### Links

Link to other pages:

```markdown
[Installation](./getting-started/installation)
[CLI Reference](./reference/cli)
```

## Maintenance

### Update Dependencies

```bash
npm update
```

### Check for Outdated Packages

```bash
npm outdated
```

### Search Integration

To enable search:

1. Create Algolia account
2. Get API keys
3. Update `docusaurus.config.js`:

```js
algolia: {
  appId: 'YOUR_APP_ID',
  apiKey: 'YOUR_SEARCH_API_KEY',
  indexName: 'kuve',
},
```

## Contributing

When contributing documentation:

1. **Follow the style guide**: Use clear, concise language
2. **Add examples**: Include code examples where relevant
3. **Test locally**: Run `npm start` to preview changes
4. **Check links**: Ensure all internal links work
5. **Update navigation**: Modify `sidebars.js` if adding new pages

## Resources

- [Docusaurus Documentation](https://docusaurus.io/docs)
- [Markdown Guide](https://www.markdownguide.org/)
- [MDX Documentation](https://mdxjs.com/)

## Support

- Report documentation issues: [GitHub Issues](https://github.com/germainlefebvre4/kuve/issues)
- Suggest improvements: [GitHub Discussions](https://github.com/germainlefebvre4/kuve/discussions)
