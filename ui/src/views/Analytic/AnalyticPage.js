import React from "react";
import CounterCard from "./CounterCard";

function AnalyticPage() {
  return (
    <div class="container">
      <div class="row">
        <div class="col-sm">
          <CounterCard counter={234} label="Hit" />
        </div>
        <div class="col-sm">
          <CounterCard counter={45} label="Unique Page" />
        </div>
        <div class="col-sm"></div>
      </div>
    </div>
  );
}

export default AnalyticPage;

// TODO: Big question: How to combine multiple reducers (and stores) to compose a global store ? i.e. RootStore = AuthStore + LocaleStore + ...
// TODO: Have to use new way of React: useState, useReducer, useContext
// Bookmarks:
// - https://leewarrick.com/blog/a-guide-to-usestate-and-usereducer/
// - https://medium.com/crowdbotics/how-to-use-usereducer-in-react-hooks-for-performance-optimization-ecafca9e7bf5
// - https://dev.to/ramsay/build-a-redux-like-global-store-using-react-hooks-4a7n
