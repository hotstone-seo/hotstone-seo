import Layout from '../components/Layout';
import Landing from './Landing';
import NotFound from './NotFound';

export default [
  {
    component: Layout,
    routes: [
      {
        path: "/",
        exact: true,
        component: Landing
      },
      {
        path: "*",
        component: NotFound
      }
    ]
  }
];