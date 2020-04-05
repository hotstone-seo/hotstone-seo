import serialize from 'serialize-javascript';
import config from 'config';
import { Helmet } from 'react-helmet';
import getClientAssets from '../assets';

export default function end(initState = {}) {
  const helmet = Helmet.renderStatic();
  const assets = getClientAssets();
  const devServer = process.env.DEV_SERVER;

  return `</div>
     <script type="text/javascript">
      window.__INITIAL_STATE__ = ${serialize(initState)};
     </script>
    
     <script type="text/javascript">
      ${!devServer ?
      `"serviceWorker" in window.navigator && window.addEventListener("load", function() {
            window.navigator.serviceWorker
              .register('/assets/service-worker.js')
              .then(function(registration) {
                console.log('ServiceWorker registration successful with scope: ', registration.scope)
              })
              .catch(function(ex) {
                console.warn(ex)
                console.warn('(This warning can be safely ignored outside of the production build.)')
              })
          })` : '\n'}
    </script>


    ${assets.vendor && assets.vendor.js ? ` <script defer src="${assets.vendor.js}" type="text/javascript"></script>` : ''}
    ${assets.client && assets.client.js ? `<script defer src="${assets.client.js}" type="text/javascript"></script>` : ''}
    
    <script type="text/javascript">
      window.CONFIG = ${serialize(config.globals)}
    </script>
    
    ${helmet.script.toString()}
  </body>
</html>`;
}
