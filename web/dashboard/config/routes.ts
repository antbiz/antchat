export default [
  {
    name: '404',
    path: '/404',
    layout: false,
    hideInMenu: true,
    component: './exception/404',
  },
  {
    name: 'signin',
    path: '/signin',
    layout: false,
    hideInMenu: true,
    component: './signin',
  },
  {
    path: '/',
    redirect: '/dashboard',
  },
  {
    path: '/dashboard',
    name: 'dashboard',
    icon: 'smile',
    component: './Welcome',
  },
];
