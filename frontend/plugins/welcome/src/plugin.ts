import { createPlugin } from '@backstage/core';
import NurseSystem from './components/NurseSystem';
import Regeister from './components/Patient_Register';
import Rent from './components/Rent_Room';
import SignIn from './components/SignIn'
import Covered from './components/CoveredPerson'
//import login from './components/login'

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/NurseSystem', NurseSystem);
    router.registerRoute('/reg', Regeister);
    router.registerRoute('/rent', Rent);
    router.registerRoute('/covered', Covered);
    router.registerRoute('/', SignIn);
    //router.registerRoute('/', login);
    
  },
});