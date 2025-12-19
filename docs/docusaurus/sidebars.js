/**
 * Creating a sidebar enables you to:
 - create an ordered group of docs
 - render a sidebar for each doc of that group
 - provide next/previous navigation

 The sidebars can be generated from the filesystem, or explicitly defined here.

 Create as many sidebars as you want.
 */

// @ts-check

/** @type {import('@docusaurus/plugin-content-docs').SidebarsConfig} */
const sidebars = {
  tutorialSidebar: [
    'intro',
    {
      type: 'category',
      label: 'Getting Started',
      collapsed: false,
      items: [
        'getting-started/installation',
        'getting-started/quickstart',
        'getting-started/shell-setup',
      ],
    },
    {
      type: 'category',
      label: 'User Guide',
      collapsed: false,
      items: [
        'user-guide/basic-usage',
        'user-guide/managing-versions',
        'user-guide/version-files',
        'user-guide/workflows',
      ],
    },
    {
      type: 'category',
      label: 'Advanced Features',
      collapsed: false,
      items: [
        'advanced/cluster-detection',
        'advanced/version-normalization',
      ],
    },
    {
      type: 'category',
      label: 'Reference',
      collapsed: false,
      items: [
        'reference/cli',
        'reference/commands',
        'reference/configuration',
        'reference/troubleshooting',
      ],
    },
    {
      type: 'category',
      label: 'Developers',
      collapsed: true,
      items: [
        'developers/architecture',
        'developers/contributing',
        'developers/development',
      ],
    },
  ],
};

export default sidebars;
