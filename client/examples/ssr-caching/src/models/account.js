const accountEndpoint = `${CONFIG.API}/endpointaccount`;

// Orders
export const getAccount = (client) => () => {
  return client(accountEndpoint, {
    method: 'GET'
  });
};
