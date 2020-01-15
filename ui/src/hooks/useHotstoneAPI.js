import { makeUseAxios } from "axios-hooks";
import axios from "axios";

const useHotstoneAPI = makeUseAxios({
  axios: axios.create({ baseURL: process.env.REACT_APP_API_URL })
});

export default useHotstoneAPI;
