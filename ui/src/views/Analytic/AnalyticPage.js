import React from "react";
import HitCounterCard from "./HitCounterCard";
import UniquePageCounterCard from "./UniquePageCounterCard";
import HitChartCard from "./HitChartCard";

function AnalyticPage() {
  return (
    <div className="container">
      <div className="row">
        <div className="col">
          <div className="card">
            <div className="card-header">Rule Analytics</div>
          </div>
        </div>
      </div>
      <div className="row">
        <div className="col">
          <HitCounterCard />
        </div>
        <div className="col">
          <UniquePageCounterCard />
        </div>
        <div className="col"></div>
      </div>
      <div className="row">
        <div className="col">
          <HitChartCard />
        </div>
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
