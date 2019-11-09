import { useContext, createContext } from 'react';

// const AppContext = React.createContext({});
// export default AppContext;

export const AppContext = createContext();

export function useAuth() {
  return useContext(AppContext);
}
