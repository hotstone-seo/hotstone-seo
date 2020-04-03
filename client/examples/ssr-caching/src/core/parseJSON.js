export default function parseJSON(context) {
  return function(response) {
    if (!response) {
      throw new Error(`Non JSON string given : ${JSON.stringify(response)}`, context);
    }

    if (typeof response === 'object') {
      return response;
    }

    return JSON.parse(response);
  };
}
