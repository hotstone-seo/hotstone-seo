import React from "react";
import { useAPI } from "@umijs/hooks";
import urljoin from "url-join";

function useHotstoneAPI(props) {
  props.url = urljoin(process.env.REACT_APP_API_URL, props.url);
  return useAPI(props);
}

export default useHotstoneAPI;
