import { defineConfig } from 'umi';

export default defineConfig({
  nodeModulesTransform: {
    type: 'none',
  },
  routes: [
    { path: '/', component: '@/pages/index' },
  ],
  fastRefresh: {},
  headScripts: ['//g.alicdn.com/chatui/icons/0.2.7/index.js'],
  metas: [
    {
      name: 'viewport',
      content: 'width=device-width, initial-scale=1',
    }
  ]
});
