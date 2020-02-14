function inspectAxiosError(error) {
  if (error === undefined || error === null) {
    console.log("=== ERROR.UNDEFINED OR NULL ===");
    console.log(error);
    return;
  }
  if (error.response) {
    // The request was made and the server responded with a status code
    // that falls out of the range of 2xx
    console.log("=== ERROR.RESPONSE ===");
    console.log(error.response.data);
    console.log(error.response.status);
    console.log(error.response.headers);
  } else if (error.request) {
    // The request was made but no response was received
    // `error.request` is an instance of XMLHttpRequest in the browser and an instance of
    // http.ClientRequest in node.js
    console.log("=== ERROR.REQUEST ===");
    console.log(error.request);
  } else {
    console.log("=== ERROR.ELSE ===");
    // Something happened in setting up the request that triggered an Error
    console.log("Error", error.message);
  }
  console.log("=== ERROR.CONFIG ===");
  console.log(error.config);
}

export function isAxiosError(error) {
  if (error === undefined || error === null) {
    return false;
  }

  return true;
}

export default inspectAxiosError;
