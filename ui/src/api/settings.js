import client from "./client";

function fetchSettings(cfg = {}) {
  return client
    .get("/settings", cfg)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

export { fetchSettings };
