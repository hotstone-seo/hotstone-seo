// === WORKING ===
// don't forget to rename this file from 'axios_disabled.js' to 'axios.js'

// const mockAxios = jest.genMockFromModule('axios')

// // this is the key to fix the axios.create() undefined error!
// mockAxios.create = jest.fn(() => mockAxios)
// export default mockAxios

// === TRY adapter (NOT WORKING) ===

// const axios = require("axios");
// const MockAdapter = require("axios-mock-adapter");

// const mockAdapter = new MockAdapter(axios);

// const mockAxios = jest.genMockFromModule('axios')

// mockAxios.create = jest.fn(() => mockAdapter)
// export default mockAxios