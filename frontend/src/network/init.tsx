import {baseApiClient} from './client';


export const init = () => {
  return baseApiClient.get("/");
};
