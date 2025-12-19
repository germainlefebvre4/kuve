// @ts-check
// `@type` JSDoc annotations allow editor autocompletion and type checking
// (when paired with `@ts-check`).
// There are various equivalent ways to declare your Docusaurus config.
// See: https://docusaurus.io/docs/api/docusaurus-config

import {themes as prismThemes} from 'prism-react-renderer';

/** @type {import('@docusaurus/types').Config} */
const config = {
  title: 'Kuve',
  tagline: 'Kubernetes Client Switcher - Manage multiple kubectl versions effortlessly',
  favicon: 'img/favicon.ico',

  // Set the production url of your site here
  url: 'https://germainlefebvre4.github.io',
  // Set the /<baseUrl>/ pathname under which your site is served
  // For GitHub pages deployment, it is often '/<projectName>/'
  baseUrl: '/kuve/',

  // GitHub pages deployment config.
  organizationName: 'germainlefebvre4',
  projectName: 'kuve',

  onBrokenLinks: 'throw',
  onBrokenMarkdownLinks: 'warn',

  // Even if you don't use internationalization, you can use this field to set
  // useful metadata like html lang. For example, if your site is Chinese, you
  // may want to replace "en" with "zh-Hans".
  i18n: {
    defaultLocale: 'en',
    locales: ['en'],
  },

  presets: [
    [
      'classic',
      /** @type {import('@docusaurus/preset-classic').Options} */
      ({
        docs: {
          sidebarPath: './sidebars.js',
          editUrl: 'https://github.com/germainlefebvre4/kuve/tree/main/docs/docusaurus/',
        },
        blog: false,
        theme: {
          customCss: './src/css/custom.css',
        },
      }),
    ],
  ],

  themeConfig:
    /** @type {import('@docusaurus/preset-classic').ThemeConfig} */
    ({
      // Replace with your project's social card
      image: 'img/kuve-social-card.jpg',
      navbar: {
        title: 'Kuve',
        logo: {
          alt: 'Kuve Logo',
          src: 'img/kuve_logo-transparent.png',
        },
        items: [
          {
            type: 'docSidebar',
            sidebarId: 'tutorialSidebar',
            position: 'left',
            label: 'Documentation',
          },
          {
            href: 'https://github.com/germainlefebvre4/kuve',
            label: 'GitHub',
            position: 'right',
          },
        ],
      },
      footer: {
        style: 'dark',
        links: [
          {
            title: 'Documentation',
            items: [
              {
                label: 'Getting Started',
                to: '/docs',
              },
              {
                label: 'Installation',
                to: '/docs/getting-started/installation',
              },
              {
                label: 'CLI Reference',
                to: '/docs/reference/cli',
              },
            ],
          },
          {
            title: 'Community',
            items: [
              {
                label: 'GitHub',
                href: 'https://github.com/germainlefebvre4/kuve',
              },
              {
                label: 'Issues',
                href: 'https://github.com/germainlefebvre4/kuve/issues',
              },
              {
                label: 'Discussions',
                href: 'https://github.com/germainlefebvre4/kuve/discussions',
              },
            ],
          },
          {
            title: 'More',
            items: [
              {
                label: 'Contributing',
                to: '/docs/developers/contributing',
              },
              {
                label: 'License',
                href: 'https://github.com/germainlefebvre4/kuve/blob/main/LICENSE',
              },
            ],
          },
        ],
        copyright: `Copyright © ${new Date().getFullYear()} Kuve Project. Built with Docusaurus.`,
      },
      prism: {
        theme: prismThemes.github,
        darkTheme: prismThemes.dracula,
        additionalLanguages: ['bash', 'go', 'json', 'yaml'],
      },
      algolia: {
        // The application ID provided by Algolia
        appId: 'YOUR_APP_ID',
        // Public API key: it is safe to commit it
        apiKey: 'YOUR_SEARCH_API_KEY',
        indexName: 'kuve',
        // Optional: see doc section below
        contextualSearch: true,
      },
    }),
};

export default config;
