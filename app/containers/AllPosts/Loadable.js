/**
 *
 * Asynchronously loads the component for AllPosts
 *
 */

import loadable from 'utils/loadable';

export default loadable(() => import('./index'));
