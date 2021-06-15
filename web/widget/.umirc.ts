import { defineConfig } from 'umi';

export default defineConfig({
  proxy: {
    'http://localhost:8002/api': {
      target: 'http://localhost:8199',
      changeOrigin: true,
      pathRewrite: { '^http://localhost:8002': '' },
    },
    'ws://localhost:8002/api': {
      target: 'ws://localhost:8199',
      changeOrigin: true,
      pathRewrite: { '^ws://localhost:8002': '' },
    },
  },
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
